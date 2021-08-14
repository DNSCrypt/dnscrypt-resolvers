# Oblivious DoH relays list

Oblivious DNS-over-HTTPS relays.

By convention, entries whose name start with `odohrelay-` prefix are *relays* and names with an `odoh-` prefix are servers that can be reached through ODoH relays.

Note that ODoH relays cannot be used with DNSCrypt servers, and DNSCrypt relays cannot be used to connect to ODoH servers.

ODoH relays can also only connect to servers supporting the ODoH protocol, not regular DoH servers.

In other words, only combine ODoH relays with ODoH servers.

To use that list, add this to the `[sources]` section of your `dnscrypt-proxy.toml` configuration file:

    [sources.'odoh-relays']
    urls = ['https://raw.githubusercontent.com/DNSCrypt/dnscrypt-resolvers/master/v3/odoh-relays.md', 'https://download.dnscrypt.info/resolvers-list/v3/odoh-relays.md']
    minisign_key = 'RWQf6LRCGA9i53mlYecO4IzT51TGPpvWucNSCh1CBM0QTaLn73Y7GFO3'
    cache_file = 'odoh-relays.md'

--


## odohrelay-crypto-sx

Oblivious DoH relay running on Fastly Compute@Edge.
Maintained by Frank Denis.

sdns://hQcAAAAAAAAAACCi3jNJDEdtNW4tvHN8J3lpIklSa2Wrj7qaNCgEgci9_BpvZG9oLXJlbGF5LmVkZ2Vjb21wdXRlLmFwcAEv


## odohrelay-koki-ams

Oblivious DoH relay in The Netherlands. No logs.
Maintained by @kokial

sdns://hQcAAAAAAAAADTUxLjE1LjEyNC4yMDgAGG9kb2gtbmwuYWxla2JlcmcubmV0OjQ0MwYvcHJveHk


## odohrelay-koki-bcn

Oblivious DoH relay in Spain. No logs.
Maintained by @kokial

sdns://hQcAAAAAAAAADjE4NS4yNTMuMTU0LjY2ABhvZG9oLWVzLmFsZWtiZXJnLm5ldDo0NDMGL3Byb3h5


## odohrelay-surf

SURFdomeinen oblivious DoH relay.
https://www.surf.nl

sdns://hQcAAAAAAAAACjE0NS4wLjYuNTMgEbEC5rH2PlKJhNYCXzKxOCQfyIu9dRlXTXDJgy1T4egVb2RvaDEuc3VyZmRvbWVpbmVuLm5sBi9wcm94eQ

