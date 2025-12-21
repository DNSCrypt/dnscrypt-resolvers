#!/usr/bin/env python3
"""
Update outdated certificate hashes in DNS stamps.

Usage:
    uv run --with pyopenssl python utils/update-cert-hashes.py v3/public-resolvers.md [--dry-run]

Requires: pyopenssl

This script finds all DoH/DoT stamps with embedded certificate hashes,
checks if the hashes are still valid, and updates them with the root CA
TBS hash if they've changed.
"""

import argparse
import base64
import hashlib
import re
import socket
import ssl
import struct
import sys
from concurrent.futures import ThreadPoolExecutor, as_completed
from dataclasses import dataclass
from typing import Optional


class DNSStamp:
    """Parser and serializer for DNS stamps (sdns:// format)."""

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
        self.pk: bytes = b""
        self.provider_name: str = ""

    @staticmethod
    def parse(stamp_str: str) -> "DNSStamp":
        """Parse a DNS stamp string into a DNSStamp object."""
        if not stamp_str.startswith("sdns://"):
            raise ValueError("Invalid stamp: must start with sdns://")

        b64 = stamp_str[7:]
        padding = 4 - (len(b64) % 4)
        if padding != 4:
            b64 += "=" * padding
        data = base64.urlsafe_b64decode(b64)

        stamp = DNSStamp()
        i = 0

        stamp.proto = data[i]
        i += 1

        stamp.props = struct.unpack("<Q", data[i : i + 8])[0]
        i += 8

        if stamp.proto == DNSStamp.DNSCRYPT:
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

        elif stamp.proto == DNSStamp.DNSCRYPT_RELAY:
            addr_len = data[i]
            i += 1
            stamp.addr = data[i : i + addr_len].decode("utf-8")
            i += addr_len

        elif stamp.proto == DNSStamp.ODOH_RELAY:
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
                length |= 0x80
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

        b64 = base64.urlsafe_b64encode(data).decode("ascii").rstrip("=")
        return f"sdns://{b64}"

    def get_connection_address(self) -> tuple[str, int]:
        """Extract the host and port to connect to for TLS."""
        addr = self.addr or self.hostname

        if addr.startswith("["):
            match = re.match(r"\[([^\]]+)\](?::(\d+))?", addr)
            if match:
                host = match.group(1)
                port = int(match.group(2)) if match.group(2) else 443
                return host, port

        if ":" in addr:
            parts = addr.rsplit(":", 1)
            if parts[1].isdigit():
                return parts[0], int(parts[1])

        return addr, 443


def parse_der_length(data: bytes, offset: int) -> tuple[int, int]:
    """Parse DER length encoding at the given offset."""
    if data[offset] < 0x80:
        return data[offset], 1
    else:
        num_bytes = data[offset] & 0x7F
        length = 0
        for i in range(num_bytes):
            length = (length << 8) | data[offset + 1 + i]
        return length, 1 + num_bytes


def get_tbs_hash_from_der(cert_der: bytes) -> bytes:
    """Extract and hash the TBS (To Be Signed) certificate data."""
    if cert_der[0] != 0x30:
        raise ValueError("Certificate must start with SEQUENCE tag")

    _, outer_len_size = parse_der_length(cert_der, 1)
    outer_header_len = 1 + outer_len_size

    tbs_start = outer_header_len
    if cert_der[tbs_start] != 0x30:
        raise ValueError("TBS certificate must be a SEQUENCE")

    tbs_content_len, tbs_len_size = parse_der_length(cert_der, tbs_start + 1)
    tbs_header_len = 1 + tbs_len_size
    tbs_total_len = tbs_header_len + tbs_content_len

    tbs_data = cert_der[tbs_start : tbs_start + tbs_total_len]
    return hashlib.sha256(tbs_data).digest()


def get_cert_hashes(host: str, port: int, server_hostname: str, timeout: float = 10) -> list[tuple[str, bytes]]:
    """Connect to server via TLS and return TBS certificate hashes for full chain.

    Returns list of (CN, TBS hash) tuples ordered from leaf to root CA.
    """
    import select
    from OpenSSL import SSL, crypto

    ctx = SSL.Context(SSL.TLS_CLIENT_METHOD)
    ctx.set_default_verify_paths()

    # Store certs with depth during verification (captures root CA)
    # Depth 0 = leaf, higher = closer to root
    chain_by_depth = {}

    def verify_callback(conn, cert, errno, depth, ok):
        chain_by_depth[depth] = cert
        return ok

    ctx.set_verify(SSL.VERIFY_PEER, verify_callback)

    sock = socket.create_connection((host, port), timeout=timeout)
    sock.setblocking(True)

    try:
        conn = SSL.Connection(ctx, sock)
        conn.set_tlsext_host_name(server_hostname.encode())
        conn.set_connect_state()

        # Handle handshake with retries for non-blocking I/O
        deadline = socket.getdefaulttimeout() or timeout
        while True:
            try:
                conn.do_handshake()
                break
            except SSL.WantReadError:
                select.select([sock], [], [], deadline)
            except SSL.WantWriteError:
                select.select([], [sock], [], deadline)

        try:
            conn.shutdown()
        except SSL.Error:
            pass
    finally:
        sock.close()

    if not chain_by_depth:
        raise RuntimeError("No certificate chain received")

    # Build results ordered by depth (leaf first, root last)
    results = []
    for depth in sorted(chain_by_depth.keys()):
        cert = chain_by_depth[depth]
        cert_der = crypto.dump_certificate(crypto.FILETYPE_ASN1, cert)
        tbs_hash = get_tbs_hash_from_der(cert_der)
        subject = cert.get_subject()
        cn = subject.CN if subject.CN else str(subject)
        results.append((cn, tbs_hash))

    return results


@dataclass
class StampUpdate:
    """Represents a stamp that needs updating."""
    resolver_name: str
    line_num: int
    old_stamp: str
    new_stamp: str
    old_hashes: list[str]
    new_hashes: list[str]
    cert_info: str


def check_and_update_stamp(resolver_name: str, line_num: int, stamp_str: str) -> Optional[StampUpdate]:
    """Check a stamp and return update info if hashes need updating."""
    try:
        stamp = DNSStamp.parse(stamp_str)
    except Exception as e:
        return None

    if stamp.proto not in (DNSStamp.DOH, DNSStamp.DOT, DNSStamp.DOQ,
                           DNSStamp.ODOH_TARGET, DNSStamp.ODOH_RELAY):
        return None

    if not stamp.hashes:
        return None

    try:
        host, port = stamp.get_connection_address()

        # Determine SNI hostname
        sni = stamp.hostname if stamp.hostname else host
        # Extract hostname without port for IPv6 addresses like [::1]:443
        if sni.startswith("["):
            match = re.match(r"\[([^\]]+)\](?::\d+)?", sni)
            if match:
                sni = match.group(1)
        elif ":" in sni and sni.rsplit(":", 1)[1].isdigit():
            sni = sni.rsplit(":", 1)[0]

        # Skip stamps where SNI is an IP address (no hostname available)
        # These can't be updated without knowing the real hostname
        if re.match(r"^[\d.:a-fA-F]+$", sni):
            return None

        cert_hashes = get_cert_hashes(host, port, sni, timeout=15)

        current_hash_set = {h for _, h in cert_hashes}
        stamp_hash_set = set(stamp.hashes)

        if stamp_hash_set & current_hash_set:
            return None

        # Prefer root CA (last in chain) as it's the most stable
        selected_cn, new_hash = cert_hashes[-1]

        # Verify the hash works by reconnecting
        verify_hashes = get_cert_hashes(host, port, sni, timeout=15)
        verify_hash_set = {h for _, h in verify_hashes}
        if new_hash not in verify_hash_set:
            raise RuntimeError(f"Verification failed: hash not found in certificate chain")

        old_hashes = [h.hex() for h in stamp.hashes]
        stamp.hashes = [new_hash]
        new_stamp = stamp.serialize()

        return StampUpdate(
            resolver_name=resolver_name,
            line_num=line_num,
            old_stamp=stamp_str,
            new_stamp=new_stamp,
            old_hashes=old_hashes,
            new_hashes=[new_hash.hex()],
            cert_info=selected_cn
        )

    except Exception as e:
        print(f"  [{resolver_name}] Error: {e}", file=sys.stderr)
        return None


def process_file(md_path: str, dry_run: bool = False, max_workers: int = 10) -> int:
    """Process the markdown file and update outdated certificate hashes."""
    with open(md_path, "r") as f:
        lines = f.readlines()

    stamps_to_check = []
    current_resolver = None

    for i, line in enumerate(lines):
        if line.startswith("## "):
            current_resolver = line[3:].strip()
        elif line.startswith("sdns://") and current_resolver:
            stamp_str = line.strip()
            try:
                stamp = DNSStamp.parse(stamp_str)
                if stamp.hashes:
                    stamps_to_check.append((current_resolver, i, stamp_str))
            except Exception:
                pass

    if not stamps_to_check:
        print("No stamps with certificate hashes found.")
        return 0

    print(f"Found {len(stamps_to_check)} stamps with certificate hashes to check...")

    updates: list[StampUpdate] = []
    errors = 0

    with ThreadPoolExecutor(max_workers=max_workers) as executor:
        futures = {
            executor.submit(check_and_update_stamp, name, line_num, stamp): (name, line_num)
            for name, line_num, stamp in stamps_to_check
        }

        for i, future in enumerate(as_completed(futures), 1):
            name, line_num = futures[future]
            print(f"\r[{i}/{len(stamps_to_check)}] Checking {name}...", end="", flush=True)

            try:
                result = future.result()
                if result:
                    updates.append(result)
            except Exception as e:
                errors += 1

    print()

    if not updates:
        print(f"\nAll certificate hashes are up to date! ({errors} connection errors)")
        return 0

    print(f"\nFound {len(updates)} stamps with outdated hashes:")
    for update in updates:
        print(f"\n  {update.resolver_name} (line {update.line_num + 1}):")
        print(f"    Old hash: {update.old_hashes[0][:16]}...")
        print(f"    New hash: {update.new_hashes[0][:16]}... ({update.cert_info})")

    if dry_run:
        print(f"\nDry run - no changes made. Use without --dry-run to apply updates.")
        return len(updates)

    for update in updates:
        lines[update.line_num] = update.new_stamp + "\n"

    with open(md_path, "w") as f:
        f.writelines(lines)

    print(f"\nUpdated {len(updates)} stamps in {md_path}")
    return len(updates)


def main():
    parser = argparse.ArgumentParser(
        description="Update outdated certificate hashes in DNS stamps."
    )
    parser.add_argument("file", help="Markdown file containing resolver stamps")
    parser.add_argument("--dry-run", action="store_true",
                        help="Check for outdated hashes without modifying the file")
    parser.add_argument("--workers", type=int, default=10,
                        help="Number of concurrent connections (default: 10)")

    args = parser.parse_args()

    try:
        count = process_file(args.file, dry_run=args.dry_run, max_workers=args.workers)
        sys.exit(0 if count == 0 else 1)
    except FileNotFoundError:
        print(f"Error: File not found: {args.file}", file=sys.stderr)
        sys.exit(2)
    except Exception as e:
        print(f"Error: {e}", file=sys.stderr)
        sys.exit(2)


if __name__ == "__main__":
    main()
