#! /usr/bin/env python3

import json
import os
import sys
import unicodedata
import urllib.request

CURRENT_DIR = "v3"


class Entry:
    name = None
    description = None
    stamps = None

    def __init__(self, name, description, stamps):
        self.name = name
        self.description = description
        self.stamps = stamps

    @staticmethod
    def parse(raw_entry):
        description = ""
        stamps = []
        lines = raw_entry.strip().splitlines()
        if len(lines) < 2:
            return None
        name = lines[0].strip()
        previous_was_blank = False
        for line in lines[1:]:
            line = line.strip()
            if previous_was_blank is True and line == "":
                continue
            previous_was_blank = False
            if line.startswith("sdns://"):
                stamps.append(line)
            else:
                description = description + line + "\n"

        description = description.strip()
        if len(name) < 2 or len(description) < 10 or len(stamps) < 1:
            return None

        return Entry(name, description, stamps)

    def format(self):
        out = "## " + self.name + "\n\n"
        out = out + self.description + "\n\n"
        for stamp in self.stamps:
            out = out + stamp + "\n"

        return out


class DNSCryDotPTEntry:
    raw = None
    name = None
    description = None

    def __init__(self, raw, name, description):
        self.raw = raw
        self.name = name
        self.description = description

    @staticmethod
    def parse(raw_entry):
        name = (
            raw_entry["location"]
            .lower()
            .replace(" ", "")
            .replace("-", "")
            .replace("'", "")
        )
        # The hand-written entries were ASCII-only, so this is done to reduce churn.
        name = unicodedata.normalize("NFKD", name).encode("ASCII", "ignore").decode()
        # NOTE: Just trusting this is true for all dnscry.pt resolvers/relays, rather than trying
        # to process the stamp to confirm this.
        description = f"DNSCry.pt {raw_entry['location']} - DNSCrypt, no filter, no logs, DNSSEC support"
        return DNSCryDotPTEntry(raw_entry, name, description)

    def get_ipv4_resolver_entry(self):
        return Entry(
            f"dnscry.pt-{self.name}-ipv4",
            f"{self.description} (IPv4 server)\n\nhttps://www.dnscry.pt",
            [self.raw["ipv4_stamp"]],
        )

    def get_ipv6_resolver_entry(self):
        return Entry(
            f"dnscry.pt-{self.name}-ipv6",
            f"{self.description} (IPv6 server)\n\nhttps://www.dnscry.pt",
            [self.raw["ipv6_stamp"]],
        )

    def get_ipv4_relay_entry(self):
        return Entry(
            f"dnscry.pt-anon-{self.name}-ipv4",
            f"{self.description} (IPv4 server)\n\nhttps://www.dnscry.pt",
            [self.raw["anon_ipv4_stamp"]],
        )

    def get_ipv6_relay_entry(self):
        return Entry(
            f"dnscry.pt-anon-{self.name}-ipv6",
            f"{self.description} (IPv6 server)\n\nhttps://www.dnscry.pt",
            [self.raw["anon_ipv6_stamp"]],
        )


def process(md_path, resolvers, is_resolvers):
    print("\n[" + md_path + "]")
    entries = {}
    previous_content = ""
    out = ""

    with open(md_path, encoding="utf8") as f:
        previous_content = f.read()
        c = previous_content.split("\n## ")
        out = out + c[0].strip() + "\n\n"
        raw_entries = c[1:]
        for i in range(0, len(raw_entries)):
            entry = Entry.parse(raw_entries[i])
            if not entry:
                print("Invalid entry: [" + raw_entries[i] + "]", file=sys.stderr)
                continue
            if entry.name in entries:
                print("Duplicate entry: [" + entry.name + "]", file=sys.stderr)
            if entry.name.startswith("dnscry.pt"):
                continue
            entries[entry.name] = entry

    for resolver in resolvers:
        if is_resolvers:
            v4 = resolver.get_ipv4_resolver_entry()
            v6 = resolver.get_ipv6_resolver_entry()
        else:
            v4 = resolver.get_ipv4_relay_entry()
            v6 = resolver.get_ipv6_relay_entry()

        # Some entries in the dnscry.pt database used to have repeated location names.
        # Right now, all entries are unique, so this code shouldn't be needed, but
        # it'll catch any future oversights rather than silently losing entries.
        if v4.name in entries:
            v4name = v4.name
            v4suffix = resolver.raw["host"].split(".")[0][3:]
            v4name = v4name.replace(resolver.name, f"{resolver.name}{v4suffix}")
            print(
                "Duplicate entry: [" + v4.name + "] => [" + v4name + "]",
                file=sys.stderr,
            )
            v4.name = v4name
        entries[v4.name] = v4

        if v6.name in entries:
            v6name = v6.name
            v6suffix = resolver.raw["host"].split(".")[0][3:]
            v6name = v6name.replace(resolver.name, f"{resolver.name}{v6suffix}")
            print(
                "Duplicate entry: [" + v6.name + "] => [" + v6name + "]",
                file=sys.stderr,
            )
            v6.name = v6name
        entries[v6.name] = v6

    for name in sorted(entries.keys()):
        entry = entries[name]
        out = out + "\n" + entry.format() + "\n"

    if out == previous_content:
        print("No changes")
    else:
        with open(md_path + ".tmp", "wt", encoding="utf8") as f:
            f.write(out)
        os.unlink(md_path)
        os.rename(md_path + ".tmp", md_path)


with urllib.request.urlopen("https://www.dnscry.pt/resolvers.json") as response:
    resolverdata = json.load(response)

resolvers = []
for data in resolverdata:
    resolvers.append(DNSCryDotPTEntry.parse(data))

resolvers.sort(key=lambda r: r.name)

process(f"{CURRENT_DIR}/public-resolvers.md", resolvers, is_resolvers=True)
process(f"{CURRENT_DIR}/relays.md", resolvers, is_resolvers=False)
