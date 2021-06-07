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


## odohrelay-fastly

Oblivious DoH relay, maintained by Frank Denis.
Backend written in Zig, running on Fastly Compute@Edge.

sdns://hQcAAAAAAAAAAAAab2RvaC1yZWxheS5lZGdlY29tcHV0ZS5hcHABLw


## odohrelay-surf

SURFdomeinen oblivious DoH relay.
https://www.surf.nl/

sdns://hQcAAAAAAAAACjE0NS4wLjYuNTMgEbEC5rH2PlKJhNYCXzKxOCQfyIu9dRlXTXDJgy1T4egVb2RvaDEuc3VyZmRvbWVpbmVuLm5sBi9wcm94eQ

