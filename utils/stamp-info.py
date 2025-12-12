#! /usr/bin/env python3

"""
DNS Stamp decoder - displays DNS stamp details in human-readable form.

Usage:
    python3 utils/stamp-info.py <stamp>
    python3 utils/stamp-info.py sdns://...

DNS stamp specification: https://dnscrypt.info/stamps-specifications/
"""

import base64
import sys


STAMP_TYPES = {
    0x00: "Plain DNS",
    0x01: "DNSCrypt",
    0x02: "DNS-over-HTTPS",
    0x03: "DNS-over-TLS",
    0x04: "DNS-over-HTTPS (path only)",
    0x05: "Oblivious DoH target",
    0x06: "Oblivious DoH relay",
    0x81: "Anonymized DNSCrypt relay",
    0x85: "Oblivious DoH relay (legacy)",
}


def decode_stamp(stamp: str) -> bytes:
    """Decode a DNS stamp from its sdns:// URL format."""
    if stamp.startswith("sdns://"):
        stamp = stamp[7:]
    # Add padding if needed
    padding = (4 - len(stamp) % 4) % 4
    stamp += "=" * padding
    return base64.urlsafe_b64decode(stamp)


def read_lp_string(data: bytes, offset: int) -> tuple[str, int]:
    """Read a length-prefixed string. Returns (string, new_offset)."""
    length = data[offset]
    value = data[offset + 1 : offset + 1 + length].decode("utf-8")
    return value, offset + 1 + length


def read_lp_bytes(data: bytes, offset: int) -> tuple[bytes, int]:
    """Read length-prefixed bytes. Returns (bytes, new_offset)."""
    length = data[offset]
    value = data[offset + 1 : offset + 1 + length]
    return value, offset + 1 + length


def read_vlp_list(data: bytes, offset: int) -> tuple[list[bytes], int]:
    """Read a variable-length-prefixed list of items. Returns (list, new_offset)."""
    items = []
    while offset < len(data):
        length = data[offset]
        if length == 0:
            offset += 1
            break
        # High bit indicates more items follow
        has_more = length & 0x80
        length &= 0x7F
        items.append(data[offset + 1 : offset + 1 + length])
        offset += 1 + length
        if not has_more:
            break
    return items, offset


def format_props(props: int) -> dict:
    """Parse the properties byte into flags."""
    return {
        "dnssec": bool(props & 0x01),
        "no_logs": bool(props & 0x02),
        "no_filter": bool(props & 0x04),
    }


def format_pk(pk: bytes) -> str:
    """Format a public key as colon-separated hex groups."""
    hex_str = pk.hex().upper()
    groups = [hex_str[i : i + 4] for i in range(0, len(hex_str), 4)]
    return ":".join(groups)


def parse_plain_dns(data: bytes) -> dict:
    """Parse a Plain DNS stamp (type 0x00)."""
    offset = 1
    props = int.from_bytes(data[offset : offset + 8], "little")
    offset += 8
    addr, offset = read_lp_string(data, offset)
    return {
        "type": "Plain DNS",
        "properties": format_props(props),
        "address": addr or "default",
    }


def parse_dnscrypt(data: bytes) -> dict:
    """Parse a DNSCrypt stamp (type 0x01)."""
    offset = 1
    props = int.from_bytes(data[offset : offset + 8], "little")
    offset += 8
    addr, offset = read_lp_string(data, offset)
    pk, offset = read_lp_bytes(data, offset)
    provider, offset = read_lp_string(data, offset)
    return {
        "type": "DNSCrypt",
        "properties": format_props(props),
        "address": addr,
        "public_key": format_pk(pk),
        "provider_name": provider,
    }


def parse_doh(data: bytes) -> dict:
    """Parse a DNS-over-HTTPS stamp (type 0x02)."""
    offset = 1
    props = int.from_bytes(data[offset : offset + 8], "little")
    offset += 8
    addr, offset = read_lp_string(data, offset)
    hashes, offset = read_vlp_list(data, offset)
    hostname, offset = read_lp_string(data, offset)
    path, offset = read_lp_string(data, offset)
    bootstrap_ips = []
    if offset < len(data):
        bootstrap_ips, offset = read_vlp_list(data, offset)

    result = {
        "type": "DNS-over-HTTPS",
        "properties": format_props(props),
        "address": addr or "default",
        "hostname": hostname,
        "path": path or "/dns-query",
    }
    if hashes:
        result["tls_pin_hashes"] = [h.hex() for h in hashes]
    if bootstrap_ips:
        result["bootstrap_ips"] = [ip.decode("utf-8") for ip in bootstrap_ips]
    return result


def parse_dot(data: bytes) -> dict:
    """Parse a DNS-over-TLS stamp (type 0x03)."""
    offset = 1
    props = int.from_bytes(data[offset : offset + 8], "little")
    offset += 8
    addr, offset = read_lp_string(data, offset)
    hashes, offset = read_vlp_list(data, offset)
    hostname, offset = read_lp_string(data, offset)
    bootstrap_ips = []
    if offset < len(data):
        bootstrap_ips, offset = read_vlp_list(data, offset)

    result = {
        "type": "DNS-over-TLS",
        "properties": format_props(props),
        "address": addr or "default",
        "hostname": hostname,
    }
    if hashes:
        result["tls_pin_hashes"] = [h.hex() for h in hashes]
    if bootstrap_ips:
        result["bootstrap_ips"] = [ip.decode("utf-8") for ip in bootstrap_ips]
    return result


def parse_doh_path(data: bytes) -> dict:
    """Parse a DNS-over-HTTPS path-only stamp (type 0x04)."""
    offset = 1
    props = int.from_bytes(data[offset : offset + 8], "little")
    offset += 8
    path, offset = read_lp_string(data, offset)
    return {
        "type": "DNS-over-HTTPS (path only)",
        "properties": format_props(props),
        "path": path,
    }


def parse_odoh_target(data: bytes) -> dict:
    """Parse an Oblivious DoH target stamp (type 0x05)."""
    offset = 1
    props = int.from_bytes(data[offset : offset + 8], "little")
    offset += 8
    hostname, offset = read_lp_string(data, offset)
    path, offset = read_lp_string(data, offset)
    return {
        "type": "Oblivious DoH target",
        "properties": format_props(props),
        "hostname": hostname,
        "path": path or "/dns-query",
    }


def parse_odoh_relay(data: bytes) -> dict:
    """Parse an Oblivious DoH relay stamp (type 0x06 or 0x85)."""
    offset = 1
    props = int.from_bytes(data[offset : offset + 8], "little")
    offset += 8
    addr, offset = read_lp_string(data, offset)
    hashes, offset = read_vlp_list(data, offset)
    hostname, offset = read_lp_string(data, offset)
    path, offset = read_lp_string(data, offset)
    bootstrap_ips = []
    if offset < len(data):
        bootstrap_ips, offset = read_vlp_list(data, offset)

    result = {
        "type": "Oblivious DoH relay",
        "properties": format_props(props),
        "address": addr or "default",
        "hostname": hostname,
        "path": path or "/proxy",
    }
    if hashes:
        result["tls_pin_hashes"] = [h.hex() for h in hashes]
    if bootstrap_ips:
        result["bootstrap_ips"] = [ip.decode("utf-8") for ip in bootstrap_ips]
    return result


def parse_dnscrypt_relay(data: bytes) -> dict:
    """Parse an Anonymized DNSCrypt relay stamp (type 0x81)."""
    offset = 1
    addr, offset = read_lp_string(data, offset)
    return {
        "type": "Anonymized DNSCrypt relay",
        "address": addr,
    }


def parse_stamp(stamp: str) -> dict:
    """Parse a DNS stamp and return its details as a dictionary."""
    data = decode_stamp(stamp)
    stamp_type = data[0]

    parsers = {
        0x00: parse_plain_dns,
        0x01: parse_dnscrypt,
        0x02: parse_doh,
        0x03: parse_dot,
        0x04: parse_doh_path,
        0x05: parse_odoh_target,
        0x06: parse_odoh_relay,
        0x81: parse_dnscrypt_relay,
        0x85: parse_odoh_relay,
    }

    parser = parsers.get(stamp_type)
    if parser is None:
        return {"error": f"Unknown stamp type: 0x{stamp_type:02X}"}

    try:
        return parser(data)
    except Exception as e:
        return {"error": f"Failed to parse stamp: {e}"}


def print_stamp_info(info: dict, indent: int = 0) -> None:
    """Print stamp information in a readable format."""
    prefix = "  " * indent
    for key, value in info.items():
        if key == "properties":
            print(f"{prefix}Properties:")
            props = value
            print(f"{prefix}  DNSSEC: {'yes' if props['dnssec'] else 'no'}")
            print(f"{prefix}  No logs: {'yes' if props['no_logs'] else 'no'}")
            print(f"{prefix}  No filter: {'yes' if props['no_filter'] else 'no'}")
        elif isinstance(value, list):
            print(f"{prefix}{key.replace('_', ' ').title()}:")
            for item in value:
                print(f"{prefix}  - {item}")
        elif isinstance(value, dict):
            print(f"{prefix}{key.replace('_', ' ').title()}:")
            print_stamp_info(value, indent + 1)
        else:
            label = key.replace("_", " ").title()
            print(f"{prefix}{label}: {value}")


def main():
    if len(sys.argv) < 2:
        print("Usage: stamp-info.py <stamp>", file=sys.stderr)
        print("       stamp-info.py sdns://...", file=sys.stderr)
        sys.exit(1)

    stamp = sys.argv[1]
    info = parse_stamp(stamp)

    if "error" in info:
        print(f"Error: {info['error']}", file=sys.stderr)
        sys.exit(1)

    print_stamp_info(info)


if __name__ == "__main__":
    main()
