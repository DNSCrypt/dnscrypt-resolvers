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


## odohrelay-ams

Oblivious DoH relay in The Netherlands. No logs.

sdns://hQcAAAAAAAAADDg5LjM4LjEzMS4zOAAYb2RvaC1ubC5hbGVrYmVyZy5uZXQ6NDQzBi9wcm94eQ


## odohrelay-crypto-sx

Oblivious DoH relay running on Fastly Compute. https://fastly.com
Maintained by Frank Denis.

sdns://hQcAAAAAAAAAAAAab2RvaC1yZWxheS5lZGdlY29tcHV0ZS5hcHABLw


## odohrelay-ibksturm

Maintained by ibksturm, Thurgau Switzerland

sdns://hQcAAAAAAAAADjIxMy4xOTYuMTkxLjk2ABhpYmtzdHVybS5zeW5vbG9neS5tZTo0NDMGL3Byb3h5


## odohrelay-se

Oblivious DoH relay in Sweden. No logs.

sdns://hQcAAAAAAAAADTQ1LjE1My4xODcuOTYAGG9kb2gtc2UuYWxla2JlcmcubmV0OjQ0MwYvcHJveHk

