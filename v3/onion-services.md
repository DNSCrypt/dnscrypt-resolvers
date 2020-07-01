# DNS servers as .onion services

All DNSCrypt and DoH servers are accessible over Tor, via exit nodes.
This is safe as all the transactions are encrypted and authenticated.

However, it may be faster to directly access a server as an onion
service. This requires specifically configured servers.

The servers below are not accessible without Tor.

To use that list, add this to the `[sources]` section of your
`dnscrypt-proxy.toml` configuration file:

    [sources.'onion-services']
    urls = ['https://raw.githubusercontent.com/DNSCrypt/dnscrypt-resolvers/master/v3/onion-services.md', 'https://download.dnscrypt.info/resolvers-list/v3/onion-services.md']
    minisign_key = 'RWQf6LRCGA9i53mlYecO4IzT51TGPpvWucNSCh1CBM0QTaLn73Y7GFO3'
    cache_file = 'onion-services.md'

--


## onion-cloudflare

Cloudflare Onion Service

sdns://AgcAAAAAAAAAACC0WWFtenR5met-s8i0oiShMtYstulWSybPBq-zBUEMNT5kbnM0dG9ycG5sZnMyaWZ1ejJzMnlmM2ZjN3JkbXNiaG02cnc3NWV1ajM1cGFjNmFwMjV6Z3FhZC5vbmlvbgovZG5zLXF1ZXJ5

