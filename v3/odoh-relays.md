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

Oblivious DoH relay running on Fastly Compute. https://fastly.com
Maintained by Frank Denis.

sdns://hQcAAAAAAAAAAAAab2RvaC1yZWxheS5lZGdlY29tcHV0ZS5hcHABLw

