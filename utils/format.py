#! /usr/bin/env python3

import sys
import os
import subprocess
from glob import glob

INCOMPATIBLE_WITH_LEGACY_VERSIONS = [
    "cira-family", "cira-private", "cira-protected"
]
CURRENT_DIR = "v3"
LEGACY_DIR = "v2"


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
        lines = raw_entry.strip().split("\n")
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


def process(md_path, signatures_to_update):
    md_legacy_path = LEGACY_DIR + "/" + os.path.basename(md_path)
    print("\n[" + md_path + "]")
    entries = {}
    previous_content = ""
    out = ""
    out_legacy = """
# *** THIS IS A LEGACY LIST ***

This is a temporary, legacy list, for dnscrypt-proxy <= 2.0.42 users.

If you are running up-to-date software, replace `/v2/` with `/v3/` in the sources URLs
of the `dnscrypt-proxy.toml` file (relevant lines start with `urls = ['https://...']`
and are present in the `[sources]` section).

THIS LIST IS AUTOMATICALLY GENERATED AS A SUBSET OF THE V3 LIST. DO NOT EDIT IT MANUALLY.

If you want to contribute changes to a resolvers list, only edit files from the `v3` directory.

--
"""

    with open(md_path) as f:
        previous_content = f.read()
        c = previous_content.split("\n## ")
        out = out + c[0].strip() + "\n\n"
        raw_entries = c[1:]
        for i in range(0, len(raw_entries)):
            entry = Entry.parse(raw_entries[i])
            if not entry:
                print(
                    "Invalid entry: [" + raw_entries[i] + "]", file=sys.stderr)
                continue
            if entry.name in entries:
                print("Duplicate entry: [" + entry.name + "]", file=sys.stderr)
            entries[entry.name] = entry

    for name in sorted(entries.keys()):
        entry = entries[name]
        out = out + "\n" + entry.format() + "\n"
        if not name in INCOMPATIBLE_WITH_LEGACY_VERSIONS:
            out_legacy = out_legacy + "\n" + entry.format_legacy() + "\n"

    if out == previous_content:
        print("No changes")
    else:
        with open(md_path + ".tmp", "wt") as f:
            f.write(out)
        os.rename(md_path + ".tmp", md_path)
        signatures_to_update.append(md_path)

    with open(md_legacy_path) as f:
        previous_content = f.read()
    if out_legacy == previous_content:
        print("No changes to the legacy version")
    else:
        with open(md_legacy_path + ".tmp", "wt") as f:
            f.write(out_legacy)
            os.rename(md_legacy_path + ".tmp", md_legacy_path)
            signatures_to_update.append(md_legacy_path)


signatures_to_update = []

for md_path in glob(CURRENT_DIR + "/*.md"):
    process(md_path, signatures_to_update)

if signatures_to_update:
    subprocess.run(["minisign", "-Sm", *signatures_to_update])
