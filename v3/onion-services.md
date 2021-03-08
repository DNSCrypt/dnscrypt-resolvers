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

sdns://AgcAAAAAAAAAAAA-ZG5zNHRvcnBubGZzMmlmdXoyczJ5ZjNmYzdyZG1zYmhtNnJ3NzVldWozNXBhYzZhcDI1emdxYWQub25pb24KL2Rucy1xdWVyeQ

