# Oblivious DoH servers list

Oblivious DNS-over-HTTPS servers.

Connecting to these can only be done via a relay from the odoh-relays list.

To use that list, add this to the `[sources]` section of your `dnscrypt-proxy.toml` configuration file:

    [sources.'odoh-servers']
    urls = ['https://raw.githubusercontent.com/DNSCrypt/dnscrypt-resolvers/master/v3/odoh-servers.md', 'https://download.dnscrypt.info/resolvers-list/v3/odoh-servers.md']
    minisign_key = 'RWQf6LRCGA9i53mlYecO4IzT51TGPpvWucNSCh1CBM0QTaLn73Y7GFO3'
    cache_file = 'odoh-servers.md'

--


## odoh-cloudflare

Cloudflare ODoH server.
https://cloudflare.com

sdns://BQcAAAAAAAAAF29kb2guY2xvdWRmbGFyZS1kbnMuY29tCi9kbnMtcXVlcnk


## odoh-crypto-sx

ODoH target server. Anycast, no logs.
Backend hosted by Scaleway. Maintained by Frank Denis.

sdns://BQcAAAAAAAAADm9kb2guY3J5cHRvLnN4Ci9kbnMtcXVlcnk


## odoh-ibksturm

ODoH target server hosted by Ibksturm. No logs, No Filter, DNSSEC.

sdns://BQcAAAAAAAAAFGlia3N0dXJtLnN5bm9sb2d5Lm1lCi9kbnMtcXVlcnk


## odoh-id-gmail

ODoH target server. Based in Singapore, no logs.
Filter ads, trackers and malware.

sdns://BQMAAAAAAAAADGRvaC50aWFyLmFwcAUvb2RvaA


## odoh-jp.tiar.app

ODoH target server. no logs.

sdns://BQcAAAAAAAAAC2pwLnRpYXIuYXBwBS9vZG9o


## odoh-jp.tiarap.org

ODoH target server via Cloudflare, no logs.

sdns://BQcAAAAAAAAADWpwLnRpYXJhcC5vcmcFL29kb2g


## odoh-koki-ams

Oblivious DoH target server in The Netherlands. No logs, No filter, DNSSEC.

sdns://BQcAAAAAAAAAFG9kb2gtbmwuYWxla2JlcmcubmV0Ci9kbnMtcXVlcnk


## odoh-koki-noads-ams

Oblivious DoH target server in The Netherlands. No logs, filter ads and malware, DNSSEC.

sdns://BQMAAAAAAAAAGm9kb2gtbm9hZHMtbmwuYWxla2JlcmcubmV0Ci9kbnMtcXVlcnk


## odoh-koki-noads-se

Oblivious DoH target server in Sweden. No logs, filter ads and malware, DNSSEC.

sdns://BQMAAAAAAAAAGm9kb2gtbm9hZHMtc2UuYWxla2JlcmcubmV0Ci9kbnMtcXVlcnk


## odoh-koki-se

Oblivious DoH target server in Sweden. No logs, No filter, DNSSEC.

sdns://BQcAAAAAAAAAFG9kb2gtc2UuYWxla2JlcmcubmV0Ci9kbnMtcXVlcnk


## odoh-marco.cx

ODoH target server by marco.cx.
Warning: uses Cloudflare resolver.

sdns://BQcAAAAAAAAAE29kb2gtcmVsYXkubWFyY28uY3gKL2Rucy1xdWVyeQ


## odoh-tiarap.org

ODoH target server via Cloudflare, no logs.
Filter ads, trackers and malware.

sdns://BQMAAAAAAAAADmRvaC50aWFyYXAub3JnBS9vZG9o

