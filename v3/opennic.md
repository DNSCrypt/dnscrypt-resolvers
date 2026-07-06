# opennic

Resolvers from the [OpenNIC](https://www.opennic.org/) project.

To use that list, add this to the `[sources]` section of your
`dnscrypt-proxy.toml` configuration file:

    [sources.'opennic']
    urls = ['https://raw.githubusercontent.com/DNSCrypt/dnscrypt-resolvers/master/v3/opennic.md', 'https://download.dnscrypt.info/resolvers-list/v3/opennic.md', 'https://cdn.jsdelivr.net/gh/DNSCrypt/dnscrypt-resolvers@master/v3/opennic.md']
    minisign_key = 'RWQf6LRCGA9i53mlYecO4IzT51TGPpvWucNSCh1CBM0QTaLn73Y7GFO3'
    cache_file = 'opennic.md'

--


## libredns

LibreDNS public resolver in Germany.
No logging, but no DNS padding and no DNSSEC support. https://libredns.gr/
Operated by LibreOps. Service page: https://libredns.gr/

sdns://AgIAAAAAAAAADjExNi4yMDIuMTc2LjI2IDLtuxHMJY--M8LNJKFJJw8L5vRG9XJks70cJbmFMycuD2RvaC5saWJyZWRucy5ncgovZG5zLXF1ZXJ5


## libredns-noads

LibreDNS ad-blocking resolver in Germany.
No logging, but no DNS padding and no DNSSEC support. Uses StevenBlack's hosts list: https://github.com/StevenBlack/hosts
Operated by LibreOps. Service page: https://libredns.gr/

sdns://AgIAAAAAAAAADjExNi4yMDIuMTc2LjI2IDLtuxHMJY--M8LNJKFJJw8L5vRG9XJks70cJbmFMycuD2RvaC5saWJyZWRucy5ncgYvbm9hZHM

