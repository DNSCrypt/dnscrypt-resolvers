#! /usr/bin/env python3

import sys


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


def process(names_path, md_path):
    names_set = set()
    in_header = True
    header = ""
    with open(names_path) as f:
        for line in f.readlines():
            line = line.rstrip()
            if in_header:
                header = header + line + "\n"
                if line == "--":
                    in_header = False
            elif line != "":
                names_set.add(line)

    entries = {}

    with open(md_path) as f:
        previous_content = f.read()
        c = previous_content.split("\n## ")
        raw_entries = c[1:]
        for i in range(0, len(raw_entries)):
            entry = Entry.parse(raw_entries[i])
            if not entry:
                print("Invalid entry: [" + raw_entries[i] + "]", file=sys.stderr)
                continue
            if entry.name in entries:
                print("Duplicate entry: [" + entry.name + "]", file=sys.stderr)
            entries[entry.name] = entry

    print(header)

    for name in entries.keys():
        if name not in names_set:
            continue

        entry = entries[name]

        print("## " + name)
        print()
        print(entry.description)
        for stamp in entry.stamps:
            print(stamp)
        print()
        print()


process(sys.argv[1], sys.argv[2])
