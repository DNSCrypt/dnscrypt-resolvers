#! /usr/bin/env python3

# Origin: https://www.dnscry.pt/anon-relays.md

import sys


in_header = True

with sys.stdin as f:
    while True:
        line = f.readline()
        if not line:
            break
        if in_header == True:
            if line.startswith("## "):
                in_header = False
            else:
                continue
        if line.startswith("## "):
            name = line.split("## ")[1]
            name = f"dnscry.pt-{name}".replace(" ", "")
            print(f"## {name}")
        else:
            print(line)
