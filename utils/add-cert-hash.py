#!/usr/bin/env python3
# /// script
# dependencies = ["pyopenssl"]
# ///
"""
Add certificate hashes to DNS stamps for a specific resolver.

Usage:
    python3 utils/add-cert-hash.py v3/public-resolvers.md resolver-name

This tool connects to the resolver's server via TLS, extracts the highest
intermediate CA certificate hash (SHA256 of the TBS certificate), and adds it
to the stamps.

The highest intermediate CA (closest to root) is used because:
- Root CAs are NOT sent by servers (only leaf + intermediates are sent)
- dnscrypt-proxy only verifies against certificates actually in the TLS chain
- Intermediate CAs are more stable than leaf certs (years vs months)

Note: The hash is computed from the TBS (To Be Signed) portion of the certificate,
which is what dnscrypt-proxy verifies against.
"""

import base64
import hashlib
import re
import select
import socket
import struct
import sys
from typing import Optional

from OpenSSL import SSL, crypto


class DNSStamp:
    """Parser and serializer for DNS stamps (sdns:// format)."""

    # Protocol types
    DNSCRYPT = 0x01
    DOH = 0x02
    DOT = 0x03
    DOQ = 0x04
    ODOH_TARGET = 0x05
    DNSCRYPT_RELAY = 0x81
    ODOH_RELAY = 0x85

    def __init__(self):
        self.proto: int = 0
        self.props: int = 0
        self.addr: str = ""
        self.hashes: list[bytes] = []
        self.hostname: str = ""
        self.path: str = ""
        self.bootstrap: list[str] = []
        # DNSCrypt specific
        self.pk: bytes = b""
        self.provider_name: str = ""

    @staticmethod
    def parse(stamp_str: str) -> "DNSStamp":
        """Parse a DNS stamp string into a DNSStamp object."""
        if not stamp_str.startswith("sdns://"):
            raise ValueError("Invalid stamp: must start with sdns://")

        # Base64url decode (add padding if needed)
        b64 = stamp_str[7:]
        padding = 4 - (len(b64) % 4)
        if padding != 4:
            b64 += "=" * padding
        data = base64.urlsafe_b64decode(b64)

        stamp = DNSStamp()
        i = 0

        # Protocol type
        stamp.proto = data[i]
        i += 1

        # Properties (8 bytes, little-endian)
        stamp.props = struct.unpack("<Q", data[i : i + 8])[0]
        i += 8

        if stamp.proto == DNSStamp.DNSCRYPT:
            # DNSCrypt: props ‖ LP(addr) ‖ LP(pk) ‖ LP(provider_name)
            addr_len = data[i]
            i += 1
            stamp.addr = data[i : i + addr_len].decode("utf-8")
            i += addr_len

            pk_len = data[i]
            i += 1
            stamp.pk = data[i : i + pk_len]
            i += pk_len

            provider_len = data[i]
            i += 1
            stamp.provider_name = data[i : i + provider_len].decode("utf-8")
            i += provider_len

        elif stamp.proto in (DNSStamp.DOH, DNSStamp.DOT, DNSStamp.DOQ, DNSStamp.ODOH_TARGET):
            # DoH/DoT/DoQ: props ‖ LP(addr) ‖ VLP(hashes) ‖ LP(hostname) ‖ LP(path) [‖ VLP(bootstrap)]
            addr_len = data[i]
            i += 1
            stamp.addr = data[i : i + addr_len].decode("utf-8")
            i += addr_len

            # VLP hashes
            stamp.hashes, i = DNSStamp._parse_vlp(data, i)

            hostname_len = data[i]
            i += 1
            stamp.hostname = data[i : i + hostname_len].decode("utf-8")
            i += hostname_len

            path_len = data[i]
            i += 1
            stamp.path = data[i : i + path_len].decode("utf-8")
            i += path_len

            # Optional bootstrap IPs
            if i < len(data):
                bootstrap_bytes, i = DNSStamp._parse_vlp(data, i)
                stamp.bootstrap = [b.decode("utf-8") for b in bootstrap_bytes]

        elif stamp.proto == DNSStamp.DNSCRYPT_RELAY:
            # Relay: props ‖ LP(addr)
            addr_len = data[i]
            i += 1
            stamp.addr = data[i : i + addr_len].decode("utf-8")
            i += addr_len

        elif stamp.proto == DNSStamp.ODOH_RELAY:
            # ODoH relay: props ‖ LP(addr) ‖ VLP(hashes) ‖ LP(hostname) ‖ LP(path) [‖ VLP(bootstrap)]
            addr_len = data[i]
            i += 1
            stamp.addr = data[i : i + addr_len].decode("utf-8")
            i += addr_len

            stamp.hashes, i = DNSStamp._parse_vlp(data, i)

            hostname_len = data[i]
            i += 1
            stamp.hostname = data[i : i + hostname_len].decode("utf-8")
            i += hostname_len

            path_len = data[i]
            i += 1
            stamp.path = data[i : i + path_len].decode("utf-8")
            i += path_len

            if i < len(data):
                bootstrap_bytes, i = DNSStamp._parse_vlp(data, i)
                stamp.bootstrap = [b.decode("utf-8") for b in bootstrap_bytes]

        return stamp

    @staticmethod
    def _parse_vlp(data: bytes, start: int) -> tuple[list[bytes], int]:
        """Parse VLP (variable-length-prefixed) data."""
        items = []
        i = start
        while i < len(data):
            length = data[i]
            has_more = length & 0x80
            length = length & 0x7F
            i += 1
            if length > 0:
                items.append(data[i : i + length])
                i += length
            if not has_more:
                break
        return items, i

    @staticmethod
    def _encode_vlp(items: list[bytes]) -> bytes:
        """Encode items using VLP format."""
        if not items:
            return b"\x00"
        result = b""
        for idx, item in enumerate(items):
            length = len(item)
            if idx < len(items) - 1:
                length |= 0x80  # More items follow
            result += bytes([length]) + item
        return result

    @staticmethod
    def _encode_lp(s: str) -> bytes:
        """Encode a string using LP (length-prefixed) format."""
        encoded = s.encode("utf-8")
        return bytes([len(encoded)]) + encoded

    def serialize(self) -> str:
        """Serialize the stamp back to sdns:// string format."""
        data = bytes([self.proto])
        data += struct.pack("<Q", self.props)

        if self.proto == DNSStamp.DNSCRYPT:
            data += self._encode_lp(self.addr)
            data += bytes([len(self.pk)]) + self.pk
            data += self._encode_lp(self.provider_name)

        elif self.proto in (DNSStamp.DOH, DNSStamp.DOT, DNSStamp.DOQ, DNSStamp.ODOH_TARGET):
            data += self._encode_lp(self.addr)
            data += self._encode_vlp(self.hashes)
            data += self._encode_lp(self.hostname)
            data += self._encode_lp(self.path)
            if self.bootstrap:
                data += self._encode_vlp([b.encode("utf-8") for b in self.bootstrap])

        elif self.proto == DNSStamp.DNSCRYPT_RELAY:
            data += self._encode_lp(self.addr)

        elif self.proto == DNSStamp.ODOH_RELAY:
            data += self._encode_lp(self.addr)
            data += self._encode_vlp(self.hashes)
            data += self._encode_lp(self.hostname)
            data += self._encode_lp(self.path)
            if self.bootstrap:
                data += self._encode_vlp([b.encode("utf-8") for b in self.bootstrap])

        # Base64url encode without padding
        b64 = base64.urlsafe_b64encode(data).decode("ascii").rstrip("=")
        return f"sdns://{b64}"

    def get_connection_address(self) -> tuple[str, int]:
        """Extract the host and port to connect to for TLS."""
        # Prefer explicit address, fall back to hostname
        addr = self.addr or self.hostname

        # Parse IPv6 address with optional port: [addr]:port
        if addr.startswith("["):
            match = re.match(r"\[([^\]]+)\](?::(\d+))?", addr)
            if match:
                host = match.group(1)
                port = int(match.group(2)) if match.group(2) else 443
                return host, port

        # Parse IPv4 or hostname with optional port
        if ":" in addr:
            # Could be IPv4:port or just hostname:port
            parts = addr.rsplit(":", 1)
            if parts[1].isdigit():
                return parts[0], int(parts[1])

        return addr, 443


def get_cert_hashes(host: str, port: int, server_hostname: str) -> list[tuple[str, bytes]]:
    """
    Connect to server via TLS and return TBS certificate hashes for certs sent by server.

    Args:
        host: IP address or hostname to connect to
        port: Port number
        server_hostname: SNI hostname for TLS

    Returns:
        List of (subject_cn, tbs_hash) tuples for each certificate sent by the server.
        The list is ordered: [leaf_cert, intermediate_ca1, intermediate_ca2, ...]
        Note: Root CA is not included since it's not sent by the server.
    """
    context = SSL.Context(SSL.TLS_CLIENT_METHOD)
    context.set_default_verify_paths()
    context.set_verify(SSL.VERIFY_PEER, lambda *args: True)

    sock = socket.create_connection((host, port), timeout=10)
    sock.setblocking(False)
    conn = SSL.Connection(context, sock)
    conn.set_tlsext_host_name(server_hostname.encode())
    conn.set_connect_state()

    while True:
        try:
            conn.do_handshake()
            break
        except SSL.WantReadError:
            select.select([sock], [], [], 10)
        except SSL.WantWriteError:
            select.select([], [sock], [], 10)

    cert_chain = conn.get_peer_cert_chain()
    if not cert_chain:
        conn.shutdown()
        sock.close()
        raise RuntimeError("No certificate chain received")

    results = []
    for cert in cert_chain:
        cert_der = crypto.dump_certificate(crypto.FILETYPE_ASN1, cert)
        tbs_hash = get_tbs_hash_from_der(cert_der)
        subject = cert.get_subject().CN or "unknown"
        results.append((subject, tbs_hash))

    conn.shutdown()
    sock.close()

    if not results:
        raise RuntimeError("No server certificates found in chain")

    return results


def get_tbs_hash_from_der(cert_der: bytes) -> bytes:
    """
    Extract and hash the TBS (To Be Signed) certificate data.

    The TBS certificate is the portion of the certificate that is signed,
    excluding the signature algorithm and signature itself.

    X.509 Certificate structure (DER):
    SEQUENCE {
        tbsCertificate       TBSCertificate,      <- This is what we hash
        signatureAlgorithm   AlgorithmIdentifier,
        signatureValue       BIT STRING
    }
    """
    # Parse the outer SEQUENCE
    if cert_der[0] != 0x30:
        raise ValueError("Certificate must start with SEQUENCE tag")

    # Get length of outer sequence (1 byte tag + length field)
    _, outer_len_size = parse_der_length(cert_der, 1)
    outer_header_len = 1 + outer_len_size  # tag + length bytes

    # The TBS certificate starts right after the outer header
    tbs_start = outer_header_len
    if cert_der[tbs_start] != 0x30:
        raise ValueError("TBS certificate must be a SEQUENCE")

    # Get length of TBS certificate content and its header size
    tbs_content_len, tbs_len_size = parse_der_length(cert_der, tbs_start + 1)
    tbs_header_len = 1 + tbs_len_size  # tag + length bytes
    tbs_total_len = tbs_header_len + tbs_content_len

    # Extract TBS certificate bytes (header + content)
    tbs_data = cert_der[tbs_start : tbs_start + tbs_total_len]

    # Hash it
    return hashlib.sha256(tbs_data).digest()


def parse_der_length(data: bytes, offset: int) -> tuple[int, int]:
    """
    Parse DER length encoding at the given offset.

    Returns (content_length, length_field_size)
    where length_field_size is the number of bytes used to encode the length
    (not including the tag byte).
    """
    if data[offset] < 0x80:
        # Short form: length is directly encoded in one byte
        return data[offset], 1
    else:
        # Long form: first byte indicates number of length bytes
        num_bytes = data[offset] & 0x7F
        length = 0
        for i in range(num_bytes):
            length = (length << 8) | data[offset + 1 + i]
        return length, 1 + num_bytes


def process_resolver(md_path: str, resolver_name: str) -> None:
    """
    Read the markdown file, find the resolver, add cert hashes to its stamps.
    """
    with open(md_path, "r") as f:
        content = f.read()

    # Split by resolver entries
    parts = content.split("\n## ")
    header = parts[0]
    entries = parts[1:]

    modified = False
    new_entries = []

    for entry in entries:
        lines = entry.split("\n")
        name = lines[0].strip()

        if name == resolver_name:
            print(f"Found resolver: {name}")
            new_lines = [lines[0]]

            for line in lines[1:]:
                if line.startswith("sdns://"):
                    try:
                        stamp = DNSStamp.parse(line)

                        # Only DoH, DoT, DoQ, ODoH support certificate hashes
                        if stamp.proto in (DNSStamp.DOH, DNSStamp.DOT, DNSStamp.DOQ,
                                          DNSStamp.ODOH_TARGET, DNSStamp.ODOH_RELAY):
                            host, port = stamp.get_connection_address()
                            # Use hostname for SNI
                            sni = stamp.hostname.split(":")[0] if stamp.hostname else host

                            print(f"  Connecting to {host}:{port} (SNI: {sni})...")

                            try:
                                cert_hashes = get_cert_hashes(host, port, sni)
                                print(f"  Certificates sent by server:")
                                for i, (cn, h) in enumerate(cert_hashes):
                                    if i == 0:
                                        cert_type = "Leaf"
                                    else:
                                        cert_type = f"Intermediate CA {i}"
                                    print(f"    [{cert_type}] {cn}: {h.hex()}")

                                # Use highest intermediate CA (last sent by server) - most stable
                                selected_cn, cert_hash = cert_hashes[-1]
                                print(f"  Using: {selected_cn}")

                                # Update hash if it differs from the expected one
                                if stamp.hashes != [cert_hash]:
                                    stamp.hashes = [cert_hash]
                                    new_stamp = stamp.serialize()
                                    print(f"  Old: {line}")
                                    print(f"  New: {new_stamp}")
                                    line = new_stamp
                                    modified = True
                                else:
                                    print("  Hash already present")

                            except Exception as e:
                                print(f"  Error getting cert: {e}")
                        else:
                            print(f"  Skipping non-DoH/DoT stamp (proto={stamp.proto:#x})")

                    except Exception as e:
                        print(f"  Error parsing stamp: {e}")

                new_lines.append(line)

            new_entries.append("\n".join(new_lines))
        else:
            new_entries.append(entry)

    if modified:
        new_content = header + "\n## " + "\n## ".join(new_entries)
        with open(md_path, "w") as f:
            f.write(new_content)
        print(f"\nFile updated: {md_path}")
    else:
        print("\nNo changes made")


def main():
    if len(sys.argv) != 3:
        print(f"Usage: {sys.argv[0]} <markdown-file> <resolver-name>")
        print(f"Example: {sys.argv[0]} v3/public-resolvers.md cira-family")
        sys.exit(1)

    md_path = sys.argv[1]
    resolver_name = sys.argv[2]

    process_resolver(md_path, resolver_name)


if __name__ == "__main__":
    main()
