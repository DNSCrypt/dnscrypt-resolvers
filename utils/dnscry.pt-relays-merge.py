#! /usr/bin/env python3

# Origin: https://www.dnscry.pt/anon-relays.md

import sys


in_header = True

with sys.stdin as f:
    while True:
        head = f.readline()
        if not head:
            break
        if in_header == True:
            if head.startswith("## "):
                in_header = False
            else:
                continue
        if head.startswith("## "):
            name = head.split("## ")[1]
            name = f"dnscry.pt-{name}".replace(" ", "")
            print(f"## {name}")
        else:
            print(head)
