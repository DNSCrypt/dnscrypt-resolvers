#! /usr/bin/env python

import sys

with open(sys.argv[1]) as f:
    c = f.read().split("\n## ")
    print(c[0].strip())
    entries = c[1:]
    for i in range(0, len(entries)):
        entries[i] = entries[i].strip()
    for entry in sorted(entries):
        print("\n\n## " + entry)
