#! /usr/bin/env python3

"""
Add an IP address to DoH stamps in a resolver entry.

Usage:
    python3 utils/add-doh-ip.py <file> <resolver-name> <ip-address>

Examples:
    python3 utils/add-doh-ip.py v3/public-resolvers.md cloudflare 1.1.1.1
    python3 utils/add-doh-ip.py v3/public-resolvers.md cloudflare-ipv6 [2606:4700:4700::1111]

IPv6 addresses must be enclosed in square brackets.
"""

import base64
import re
import sys


def decode_stamp(stamp: str) -> bytes:
    """Decode a DNS stamp from its sdns:// URL format."""
    if stamp.startswith("sdns://"):
        stamp = stamp[7:]
    padding = (4 - len(stamp) % 4) % 4
    stamp += "=" * padding
    return base64.urlsafe_b64decode(stamp)


def encode_stamp(data: bytes) -> str:
    """Encode binary data to a DNS stamp URL."""
    encoded = base64.urlsafe_b64encode(data).decode("ascii")
    return "sdns://" + encoded.rstrip("=")


def read_lp_string(data: bytes, offset: int) -> tuple[str, int]:
    """Read a length-prefixed string."""
    length = data[offset]
    value = data[offset + 1 : offset + 1 + length].decode("utf-8")
    return value, offset + 1 + length


def write_lp_string(value: str) -> bytes:
    """Write a length-prefixed string."""
    encoded = value.encode("utf-8")
    return bytes([len(encoded)]) + encoded


def read_vlp_list(data: bytes, offset: int) -> tuple[list[bytes], int]:
    """Read a variable-length-prefixed list."""
    items = []
    while offset < len(data):
        length = data[offset]
        if length == 0:
            offset += 1
            break
        has_more = length & 0x80
        length &= 0x7F
        items.append(data[offset + 1 : offset + 1 + length])
        offset += 1 + length
        if not has_more:
            break
    return items, offset


def write_vlp_list(items: list[bytes]) -> bytes:
    """Write a variable-length-prefixed list."""
    if not items:
        return b"\x00"
    result = b""
    for i, item in enumerate(items):
        length = len(item)
        if i < len(items) - 1:
            length |= 0x80  # More items follow
        result += bytes([length]) + item
    return result


def add_ip_to_doh_stamp(stamp: str, ip_address: str) -> str:
    """Add or replace the IP address in a DoH stamp."""
    data = decode_stamp(stamp)

    # Check stamp type
    stamp_type = data[0]
    if stamp_type != 0x02:
        raise ValueError(f"Not a DoH stamp (type 0x{stamp_type:02X})")

    # Parse the stamp
    offset = 1
    props = data[offset : offset + 8]
    offset += 8

    # Skip old address
    old_addr, offset = read_lp_string(data, offset)

    # Read hashes
    hashes, offset = read_vlp_list(data, offset)

    # Read hostname
    hostname, offset = read_lp_string(data, offset)

    # Read path
    path, offset = read_lp_string(data, offset)

    # Read bootstrap IPs if present
    bootstrap_ips = []
    if offset < len(data):
        bootstrap_ips, offset = read_vlp_list(data, offset)

    # Rebuild the stamp with the new address
    new_data = bytes([0x02])  # DoH type
    new_data += props
    new_data += write_lp_string(ip_address)
    new_data += write_vlp_list(hashes)
    new_data += write_lp_string(hostname)
    new_data += write_lp_string(path)
    if bootstrap_ips:
        new_data += write_vlp_list(bootstrap_ips)

    return encode_stamp(new_data)


def format_ip(ip: str) -> str:
    """Format IP address for stamp: IPv6 must be in brackets."""
    ip = ip.strip()
    # If it looks like IPv6 and not already in brackets, add them
    if ":" in ip and not ip.startswith("["):
        return f"[{ip}]"
    return ip


def process_file(filepath: str, resolver_name: str, ip_address: str) -> bool:
    """Process a resolver file and update DoH stamps for a specific resolver."""
    with open(filepath) as f:
        content = f.read()

    # Find the resolver entry
    pattern = rf"(## {re.escape(resolver_name)}\n)(.*?)(\n## |\Z)"
    match = re.search(pattern, content, re.DOTALL)

    if not match:
        print(f"Error: Resolver '{resolver_name}' not found in {filepath}", file=sys.stderr)
        return False

    entry_start = match.start()
    entry_header = match.group(1)
    entry_body = match.group(2)
    entry_end_marker = match.group(3)

    # Find and update DoH stamps in the entry
    stamp_pattern = r"(sdns://Ag[A-Za-z0-9_-]+)"
    stamps_found = 0
    stamps_updated = 0

    def replace_stamp(m):
        nonlocal stamps_found, stamps_updated
        stamps_found += 1
        old_stamp = m.group(1)
        try:
            new_stamp = add_ip_to_doh_stamp(old_stamp, ip_address)
            stamps_updated += 1
            print(f"Updated: {old_stamp[:50]}...")
            print(f"     to: {new_stamp[:50]}...")
            return new_stamp
        except Exception as e:
            print(f"Warning: Could not update stamp: {e}", file=sys.stderr)
            return old_stamp

    new_entry_body = re.sub(stamp_pattern, replace_stamp, entry_body)

    if stamps_found == 0:
        print(f"Error: No DoH stamps found in resolver '{resolver_name}'", file=sys.stderr)
        return False

    if stamps_updated == 0:
        print("No stamps were updated", file=sys.stderr)
        return False

    # Reconstruct the file
    new_content = (
        content[:entry_start]
        + entry_header
        + new_entry_body
        + (entry_end_marker if entry_end_marker != "\n## " else "")
    )
    # If we removed the separator, need to add it back for next entry
    if entry_end_marker == "\n## ":
        new_content += "\n## "
    new_content += content[match.end() :]

    with open(filepath, "w") as f:
        f.write(new_content)

    print(f"\nUpdated {stamps_updated} stamp(s) in '{resolver_name}'")
    return True


def main():
    if len(sys.argv) != 4:
        print("Usage: add-doh-ip.py <file> <resolver-name> <ip-address>", file=sys.stderr)
        print("", file=sys.stderr)
        print("Examples:", file=sys.stderr)
        print("  add-doh-ip.py v3/public-resolvers.md cloudflare 1.1.1.1", file=sys.stderr)
        print("  add-doh-ip.py v3/public-resolvers.md my-resolver-ipv6 2606:4700::1", file=sys.stderr)
        sys.exit(1)

    filepath = sys.argv[1]
    resolver_name = sys.argv[2]
    ip_address = format_ip(sys.argv[3])

    if not process_file(filepath, resolver_name, ip_address):
        sys.exit(1)


if __name__ == "__main__":
    main()
