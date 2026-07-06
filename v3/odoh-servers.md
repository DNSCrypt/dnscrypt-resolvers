# Oblivious DoH servers list

Oblivious DNS-over-HTTPS servers.

Connecting to these can only be done via a relay from the odoh-relays list.

To use that list, add this to the `[sources]` section of your `dnscrypt-proxy.toml` configuration file:

    [sources.'odoh-servers']
    urls = ['https://raw.githubusercontent.com/DNSCrypt/dnscrypt-resolvers/master/v3/odoh-servers.md', 'https://download.dnscrypt.info/resolvers-list/v3/odoh-servers.md', 'https://cdn.jsdelivr.net/gh/DNSCrypt/dnscrypt-resolvers@master/v3/odoh-servers.md']
    minisign_key = 'RWQf6LRCGA9i53mlYecO4IzT51TGPpvWucNSCh1CBM0QTaLn73Y7GFO3'
    cache_file = 'odoh-servers.md'

--


## odoh-cloudflare

Cloudflare ODoH target.
Operated by Cloudflare. https://cloudflare.com

sdns://BQcAAAAAAAAAF29kb2guY2xvdWRmbGFyZS1kbnMuY29tCi9kbnMtcXVlcnk


## odoh-crypto-sx

crypto.sx ODoH target.
Anycast, no logs. Backend hosted by Scaleway. Maintained by Frank Denis.
Maintained by Frank Denis. Service page: https://crypto.sx

sdns://BQcAAAAAAAAADm9kb2guY3J5cHRvLnN4Ci9kbnMtcXVlcnk


## odoh-id-gmail

Tiarap Singapore ODoH target.
Maintained by id-gmail. Based in Singapore, no logs. Filters ads, trackers and malware.
Maintained by id-gmail / Tiarap. Service page: https://tiarap.org/

sdns://BQMAAAAAAAAADGRvaC50aWFyLmFwcAUvb2RvaA


## odoh-snowstorm

Snowstorm ODoH target.
Hosted by Snowstorm. No logs, no filtering, DNSSEC.
Operated by Snowstorm. Service page: https://snowstorm.net/

sdns://BQcAAAAAAAAAE2RvcGUuc25vd3N0b3JtLmxvdmUKL2Rucy1xdWVyeQ


## odoh-tiarap.org

Tiarap Cloudflare-cached ODoH target.
Maintained by id-gmail. No logs. Filters ads, trackers and malware.
Maintained by id-gmail / Tiarap. Service page: https://tiarap.org/

sdns://BQMAAAAAAAAADmRvaC50aWFyYXAub3JnBS9vZG9o

