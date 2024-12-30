#! /usr/bin/env python3

import base64
import csv
import os
import subprocess
import sys
from glob import glob

INCOMPATIBLE_WITH_LEGACY_VERSIONS = ["cira-family", "cira-private", "cira-protected"]
CURRENT_DIR = "v3"
LEGACY_DIR = "v2"
HISTORIC_DIR = "v1"
MINISIGN_PK = "RWQf6LRCGA9i53mlYecO4IzT51TGPpvWucNSCh1CBM0QTaLn73Y7GFO3"


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

    def format_legacy(self):
        out = "## " + self.name + "\n\n"
        out = out + self.description + "\n\n"
        out = out + self.stamps[0] + "\n"

        return out

    def csv_entry(self):
        parsed = DNSCryptStamp.parse(self.stamps[0])
        if parsed is None:
            return None
        version = 2
        if self.name.find("cisco") >= 0:
            version = 1
        dnssec = "no"
        nolog = "no"
        namecoin = "no"
        if parsed.dnssec:
            dnssec = "yes"
        if parsed.nolog:
            nolog = "yes"
        csv_entry = [
            self.name,
            self.name,
            self.description.splitlines(False)[0],
            None,
            None,
            None,
            version,
            dnssec,
            nolog,
            namecoin,
            parsed.addr,
            parsed.provider,
            parsed.pk,
            None,
        ]

        return csv_entry


class DNSCryptStamp:
    dnssec = False
    nolog = False
    nofilter = False
    addr = None
    pk = None
    provider = None

    @staticmethod
    def parse(stamp):
        bin = base64.urlsafe_b64decode(stamp.removeprefix("sdns://") + "==")
        i = 0
        if bin[i] != 0x01:
            return None
        i = i + 1
        parsed = DNSCryptStamp()
        props = bin[i]
        parsed.dnssec = not not ((props >> 0) & 1)
        parsed.nolog = not not ((props >> 1) & 1)
        parsed.nofilter = not not ((props >> 2) & 1)
        i = i + 8
        addr_len = bin[i]
        i = i + 1
        parsed.addr = bin[i : i + addr_len].decode("utf-8")
        i = i + addr_len
        pk_len = bin[i]
        i = i + 1
        if pk_len != 32:
            return None
        hpk = bin[i : i + pk_len].hex().upper()
        hpks = []
        for j in range(0, 16):
            hpks.append(hpk[j * 4 : j * 4 + 4])
        parsed.pk = ":".join(hpks)
        i = i + pk_len
        provider_len = bin[i]
        i = i + 1
        parsed.provider = bin[i : i + provider_len].decode("utf-8")
        i = i + provider_len

        return parsed


def process(md_path, signatures_to_update):
    md_legacy_path = LEGACY_DIR + "/" + os.path.basename(md_path)
    csv_historic_path = HISTORIC_DIR + "/" + "dnscrypt-resolvers.csv"
    print("\n[" + md_path + "]")
    entries = {}
    previous_content = ""
    out = ""
    out_legacy = """
# *** THIS LIST IS FOR OLD DNSCRYPT-PROXY VERSIONS ***

Version 2 of the list is for dnscrypt-proxy <= 2.0.42 users.

If you are running up-to-date software, replace `/v2/` with `/v3/` in the sources URLs
of the `dnscrypt-proxy.toml` file (relevant lines start with `urls = ['https://...']`
and are present in the `[sources]` section).

THIS LIST IS AUTOMATICALLY GENERATED AS A SUBSET OF THE V3 LIST. DO NOT EDIT IT MANUALLY.

If you want to contribute changes to a resolvers list, only edit files from the `v3` directory.

--
"""
    csv_entries = []

    with open(md_path) as f:
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
            entries[entry.name] = entry

    for name in sorted(entries.keys()):
        entry = entries[name]
        out = out + "\n" + entry.format() + "\n"
        if name not in INCOMPATIBLE_WITH_LEGACY_VERSIONS:
            out_legacy = out_legacy + "\n" + entry.format_legacy() + "\n"

    if os.path.basename(md_path) == "public-resolvers.md":
        for name in sorted(entries.keys()):
            entry = entries[name]
            csv_entry = entry.csv_entry()
            if csv_entry:
                csv_entries.append(entry.csv_entry())

    if out == previous_content:
        print("No changes")
    else:
        with open(md_path + ".tmp", "wt") as f:
            f.write(out)
        os.unlink(md_path)
        os.rename(md_path + ".tmp", md_path)

    # Legacy

    if (
        os.path.basename(md_path) == "odoh-relays.md"
        or os.path.basename(md_path) == "odoh-servers.md"
    ):
        md_legacy_path = md_path
    else:
        with open(md_legacy_path) as f:
            previous_content = f.read()
        if out_legacy == previous_content:
            print("No changes to the legacy version")
        else:
            with open(md_legacy_path + ".tmp", "wt") as f:
                f.write(out_legacy)
            os.unlink(md_legacy_path)
            os.rename(md_legacy_path + ".tmp", md_legacy_path)

    # Historic

    if len(csv_entries) != 0:
        with open(csv_historic_path, "wt") as f:
            w = csv.writer(f, dialect="unix", quoting=csv.QUOTE_MINIMAL)
            w.writerow(
                [
                    "Name",
                    "Full name",
                    "Description",
                    "Location",
                    "Coordinates",
                    "URL",
                    "Version",
                    "DNSSEC validation",
                    "No logs",
                    "Namecoin",
                    "Resolver address",
                    "Provider name",
                    "Provider public key",
                    "Provider public key TXT record",
                ]
            )
            for csv_entry in csv_entries:
                w.writerow(csv_entry)

    # Signatures

    if signatures_to_update is not None:
        for path in [md_path, md_legacy_path, csv_historic_path]:
            try:
                subprocess.run(
                    ["minisign", "-V", "-P", MINISIGN_PK, "-m", path], check=True
                )
            except subprocess.CalledProcessError:
                signatures_to_update.append(path)


# Change to None to skip signature updates, e.g. during development.
signatures_to_update = []

for md_path in glob(CURRENT_DIR + "/*.md"):
    process(md_path, signatures_to_update)

if signatures_to_update:
    subprocess.run(["minisign", "-Slm", *signatures_to_update])
