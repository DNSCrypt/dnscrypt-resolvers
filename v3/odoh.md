# *** THIS LIST IS DEPRECATED ***

It was only used by 2.0.46 beta versions, and has since be split into
the `odoh-servers` and `odoh-relays` lists.

Please update your configuration to use these two new lists instead,
as the combined list is going to be removed soon.

~~~

Oblivious DNS-over-HTTPS servers and relays.

By convention, entries whose name start with `odohrelay-` prefix are *relays* and names with an `odoh-` prefix are servers that can be reached through ODoH relays.

Note that ODoH relays cannot be used with DNSCrypt servers, and DNSCrypt relays cannot be used to connect to ODoH servers.

ODoH relays can also only connect to servers supporting the ODoH protocol, not regular DoH servers.

In other words, only combine elements from that list together.

To use that list, add this to the `[sources]` section of your `dnscrypt-proxy.toml` configuration file:

    [sources.'odoh']
    urls = ['https://raw.githubusercontent.com/DNSCrypt/dnscrypt-resolvers/master/v3/odoh.md', 'https://download.dnscrypt.info/resolvers-list/v3/odoh.md']
    minisign_key = 'RWQf6LRCGA9i53mlYecO4IzT51TGPpvWucNSCh1CBM0QTaLn73Y7GFO3'
    cache_file = 'odoh.md'

--


## odoh-cloudflare

Cloudflare ODoH server.
https://cloudflare.com

sdns://BQcAAAAAAAAAF29kb2guY2xvdWRmbGFyZS1kbnMuY29tCi9kbnMtcXVlcnk


## odoh-crypto-sx

ODoH target server. Anycast, no logs.
Backend hosted by Scaleway. Maintained by Frank Denis.

sdns://BQcAAAAAAAAADm9kb2guY3J5cHRvLnN4Ci9kbnMtcXVlcnk


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
Maintained by @kokial

sdns://BQcAAAAAAAAAGG9kb2gtdGFyZ2V0LmFsZWtiZXJnLm5ldAovZG5zLXF1ZXJ5


## odoh-resolver4.dns.openinternet.io

ODoH target server. no logs, no filter, DNSSEC.
Running on dedicated hardware colocated at Sonic.net in Santa Rosa, CA in the United States.

Uses Sonic's recusrive DNS servers as upstream resolvers (but is not affiliated with Sonic
in any way). Provided by https://openinternet.io

sdns://BQcAAAAAAAAAHXJlc29sdmVyNC5kbnMub3BlbmludGVybmV0LmlvCi9kbnMtcXVlcnk


## odoh-tiarap.org

ODoH target server via Cloudflare, no logs.
Filter ads, trackers and malware.

sdns://BQMAAAAAAAAADmRvaC50aWFyYXAub3JnBS9vZG9o


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

