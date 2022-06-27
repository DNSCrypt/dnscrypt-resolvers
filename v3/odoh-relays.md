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


## odohrelay-ibksturm

Oblivious DoH relay hosted by Ibksturm. No Logs

sdns://hQcAAAAAAAAADjIxMy4xOTYuMTkxLjk2ABRpYmtzdHVybS5zeW5vbG9neS5tZQovZG5zLXF1ZXJ5


## odohrelay-koki-ams

Oblivious DoH relay in The Netherlands. No logs.
Maintained by @kokial

sdns://hQcAAAAAAAAADDg5LjM4LjEzMS4zOAAYb2RvaC1ubC5hbGVrYmVyZy5uZXQ6NDQzBi9wcm94eQ


## odohrelay-koki-se

Oblivious DoH relay in Sweden. No logs.
Maintained by @kokial

sdns://hQcAAAAAAAAADTQ1LjE1My4xODcuOTYAGG9kb2gtc2UuYWxla2JlcmcubmV0OjQ0MwYvcHJveHk


## odohrelay-surf

SURFdomeinen oblivious DoH relay.
https://www.surf.nl

sdns://hQcAAAAAAAAACjE0NS4wLjYuNTMgEbEC5rH2PlKJhNYCXzKxOCQfyIu9dRlXTXDJgy1T4egVb2RvaDEuc3VyZmRvbWVpbmVuLm5sBi9wcm94eQ

