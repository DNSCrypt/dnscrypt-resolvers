# opennic

Resolvers from the [OpenNIC](https://www.opennic.org/) project.

To use that list, add this to the `[sources]` section of your
`dnscrypt-proxy.toml` configuration file:

    [sources.'opennic']
    urls = ['https://raw.githubusercontent.com/DNSCrypt/dnscrypt-resolvers/master/v3/opennic.md', 'https://download.dnscrypt.info/resolvers-list/v3/opennic.md']
    minisign_key = 'RWQf6LRCGA9i53mlYecO4IzT51TGPpvWucNSCh1CBM0QTaLn73Y7GFO3'
    cache_file = 'opennic.md'

--


## libredns

DoH server in Germany. No logging, but no DNS padding and no DNSSEC support.
https://libredns.gr/

sdns://AgIAAAAAAAAADjExNi4yMDIuMTc2LjI2INYdluv7P6GQluLAAuNgO_f9qu-yJwTahmgGK76G7UuTD2RvaC5saWJyZWRucy5ncgovZG5zLXF1ZXJ5


## libredns-noads

DoH server in Germany. No logging, but no DNS padding and no DNSSEC support.
no ads version, uses StevenBlack's host list: https://github.com/StevenBlack/hosts

sdns://AgIAAAAAAAAADjExNi4yMDIuMTc2LjI2INYdluv7P6GQluLAAuNgO_f9qu-yJwTahmgGK76G7UuTD2RvaC5saWJyZWRucy5ncgYvbm9hZHM

