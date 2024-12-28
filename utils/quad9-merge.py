#! /usr/bin/env python3

# Origin: https://github.com/Quad9DNS/dnscrypt-settings

import sys


class Entry:
    name = None
    description = None
    stamps = None

    def __init__(self, name, description, stamps):
        self.name = name
        self.description = description
        self.stamps = stamps


servers = {}

with sys.stdin as f:
    while True:
        line = f.readline()
        if line == "":
            break
        parts = line.strip().split("## ")
        name = parts[1].strip()
        description = f.readline().strip()
        stamp = f.readline().strip()
        f.readline()
        server = Entry(name, description, [stamp])
        servers[name] = server

servers2 = {}

for name in servers:
    server = servers[name]
    parts = name.split("-")
    base_name = "-".join(parts[:-1])
    ext = parts[-1]
    if ext != "pri":
        continue
    description = server.description
    stamps = server.stamps
    for extz in ["alt", "alt1", "alt2", "alt3", "alt4"]:
        alt_name = base_name + "-" + extz
        if alt_name in servers:
            alt_ip = servers[alt_name].description.split(" ")[-1]
            description += " - " + alt_ip
            stamps += servers[alt_name].stamps
    print("## quad9-" + name)
    print(description)
    print("\n".join(stamps))
    print()
