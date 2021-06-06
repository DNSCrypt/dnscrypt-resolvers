# ODoH

Oblivious DoH DNS servers and relays.

To use that list, add this to the `[sources]` section of your `dnscrypt-proxy.toml` configuration file:

[sources.'odoh']
urls = ['https://raw.githubusercontent.com/DNSCrypt/dnscrypt-resolvers/master/v3/odoh.md', 'https://download.dnscrypt.info/resolvers-list/v3/odoh.md']
minisign_key = 'RWQf6LRCGA9i53mlYecO4IzT51TGPpvWucNSCh1CBM0QTaLn73Y7GFO3'
cache_file = 'odoh.md'

--


## odoh-crypto-sx

ODoH target server. Anycast, no logs, no censorship, DNSSEC.
Backend hosted by Scaleway, globally cached via Cloudflare.
Maintained by Frank Denis.

sdns://BQcAAAAAAAAADWRvaC5jcnlwdG8uc3gKL2Rucy1xdWVyeQ

