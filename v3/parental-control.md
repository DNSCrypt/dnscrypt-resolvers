# parental-control

A set of resolvers blocking popular websites that may not be appropriate
for children.

This is not bulletproof. In particular, websites in languages that are
not English will require additional, local rules.

To use that list, add this to the `[sources]` section of your
`dnscrypt-proxy.toml` configuration file:

    [sources.'parental-control']
    urls = ['https://raw.githubusercontent.com/DNSCrypt/dnscrypt-resolvers/master/v3/parental-control.md', 'https://download.dnscrypt.info/resolvers-list/v3/parental-control.md', 'https://cdn.jsdelivr.net/gh/DNSCrypt/dnscrypt-resolvers@master/v3/parental-control.md']
    minisign_key = 'RWQf6LRCGA9i53mlYecO4IzT51TGPpvWucNSCh1CBM0QTaLn73Y7GFO3'
    cache_file = 'parental-control.md'

In order to enforce safe search results from Google and Youtube, you may
also want to enable cloaking (`cloaking_rules` in the configuration file).

--


## adguard-dns-family-doh

AdGuard DNS Family Protection resolver.
Blocks ads, trackers, phishing, malicious and adult domains. Enforces safe search where available.
Operated by AdGuard. Service page: https://adguard-dns.io/en/public-dns.html

sdns://AgMAAAAAAAAADDk0LjE0MC4xNC4xNSCaOjT3J965vKUQA9nOnDn48n3ZxSQpAcK6saROY1oCGQw5NC4xNDAuMTQuMTUKL2Rucy1xdWVyeQ
sdns://AgMAAAAAAAAADDk0LjE0MC4xNS4xNiCaOjT3J965vKUQA9nOnDn48n3ZxSQpAcK6saROY1oCGQw5NC4xNDAuMTUuMTYKL2Rucy1xdWVyeQ


## adguard-dns-family-doh-ipv6

AdGuard DNS Family Protection resolver.
IPv6 endpoint. Blocks ads, trackers, phishing, malicious and adult domains. Enforces safe search where available.
Operated by AdGuard. Service page: https://adguard-dns.io/en/public-dns.html

sdns://AgMAAAAAAAAAFFsyYTEwOjUwYzA6OmJhZDE6ZmZdIJo6NPcn3rm8pRAD2c6cOfjyfdnFJCkBwrqxpE5jWgIZFmZhbWlseS5hZGd1YXJkLWRucy5jb20KL2Rucy1xdWVyeQ
sdns://AgMAAAAAAAAAFFsyYTEwOjUwYzA6OmJhZDI6ZmZdIJo6NPcn3rm8pRAD2c6cOfjyfdnFJCkBwrqxpE5jWgIZFmZhbWlseS5hZGd1YXJkLWRucy5jb20KL2Rucy1xdWVyeQ


## cira-family

CIRA Canadian Shield Family resolver.
Operated by the Canadian Internet Registration Authority for Canadian users on Canadian-hosted servers. Adds sexual-content blocking to the Protected malware, phishing, botnet and scam protection.

Operated by the Canadian Internet Registration Authority. Service page: https://www.cira.ca/en/canadian-shield/
Info: Not anonymous, but CIRA says it does not sell user data or use it for advertising. The 'Family' name is one that CIRA uses for this resolver.

sdns://AgEAAAAAAAAADjE0OS4xMTIuMTIxLjMwIMlfeyD2_NOf06B6LkQlJCO2NP2-NeHgRdlk3upiYRXLHWZhbWlseS5jYW5hZGlhbnNoaWVsZC5jaXJhLmNhCi9kbnMtcXVlcnk
sdns://AgEAAAAAAAAADjE0OS4xMTIuMTIyLjMwIMlfeyD2_NOf06B6LkQlJCO2NP2-NeHgRdlk3upiYRXLHWZhbWlseS5jYW5hZGlhbnNoaWVsZC5jaXJhLmNhCi9kbnMtcXVlcnk


## cira-family-ipv6

CIRA Canadian Shield Family resolver.
IPv6 endpoint. Operated by the Canadian Internet Registration Authority for Canadian users on Canadian-hosted servers. Adds sexual-content blocking to the Protected malware, phishing, botnet and scam protection.

Operated by the Canadian Internet Registration Authority. Service page: https://www.cira.ca/en/canadian-shield/
Info: Not anonymous, but CIRA says it does not sell user data or use it for advertising. The 'Family' name is one that CIRA uses for this resolver.

sdns://AgEAAAAAAAAAE1syNjIwOjEwQTo4MEJCOjozMF0gyV97IPb805_ToHouRCUkI7Y0_b414eBF2WTe6mJhFcsdZmFtaWx5LmNhbmFkaWFuc2hpZWxkLmNpcmEuY2EKL2Rucy1xdWVyeQ
sdns://AgEAAAAAAAAAE1syNjIwOjEwQTo4MEJDOjozMF0gyV97IPb805_ToHouRCUkI7Y0_b414eBF2WTe6mJhFcsdZmFtaWx5LmNhbmFkaWFuc2hpZWxkLmNpcmEuY2EKL2Rucy1xdWVyeQ


## cleanbrowsing-adult-doh

CleanBrowsing Adult Filter.
Blocks adult, pornographic and explicit sites. Allows proxy and VPN domains and mixed-content sites. Google and Bing are set to Safe Mode.
Operated by CleanBrowsing. Service page: https://cleanbrowsing.org/filters/

sdns://AgMAAAAAAAAADjE4NS4yMjguMTY4LjEwIIxQM5c__k_x7wqpKYT8kygbp6knjvi_ZBImuJjRDU50EWNsZWFuYnJvd3Npbmcub3JnES9kb2gvYWR1bHQtZmlsdGVy
sdns://AgMAAAAAAAAADzE4NS4yMjguMTY4LjE2OCCMUDOXP_5P8e8KqSmE_JMoG6epJ474v2QSJriY0Q1OdBFjbGVhbmJyb3dzaW5nLm9yZxEvZG9oL2FkdWx0LWZpbHRlcg


## cleanbrowsing-family-doh

CleanBrowsing Family Filter.
Blocks adult and explicit sites, proxy and VPN bypass domains, and mixed-content sites. Google, Bing and YouTube are set to Safe Mode.
Operated by CleanBrowsing. Service page: https://cleanbrowsing.org/filters/

sdns://AgMAAAAAAAAADjE4NS4yMjguMTY4LjEwIIxQM5c__k_x7wqpKYT8kygbp6knjvi_ZBImuJjRDU50EWNsZWFuYnJvd3Npbmcub3JnEi9kb2gvZmFtaWx5LWZpbHRlcg
sdns://AgMAAAAAAAAADzE4NS4yMjguMTY4LjE2OCCMUDOXP_5P8e8KqSmE_JMoG6epJ474v2QSJriY0Q1OdBFjbGVhbmJyb3dzaW5nLm9yZxIvZG9oL2ZhbWlseS1maWx0ZXI


## cloudflare-family

Cloudflare 1.1.1.1 for Families resolver.
Global anycast. Blocks malware and adult-content domains.
Operated by Cloudflare. Service page: https://developers.cloudflare.com/1.1.1.1/

sdns://AgMAAAAAAAAABzEuMC4wLjMABzEuMC4wLjMKL2Rucy1xdWVyeQ


## cloudflare-family-ipv6

Cloudflare 1.1.1.1 for Families resolver.
IPv6 endpoint. Global anycast. Blocks malware and adult-content domains.
Operated by Cloudflare. Service page: https://developers.cloudflare.com/1.1.1.1/

sdns://AgMAAAAAAAAAFlsyNjA2OjQ3MDA6NDcwMDo6MTExM10AGlsyNjA2OjQ3MDA6NDcwMDo6MTExM106NDQzCi9kbnMtcXVlcnk
sdns://AgMAAAAAAAAAFlsyNjA2OjQ3MDA6NDcwMDo6MTAwM10AGlsyNjA2OjQ3MDA6NDcwMDo6MTAwM106NDQzCi9kbnMtcXVlcnk


## dnsforfamily

DNS for Family safe-search resolver.
Blocks adult, gambling, drug, malware, scam/phishing, proxy/anonymizer and advertising domains; its blacklist is updated daily. Enforces safe search in Google, YouTube, Brave, Ecosia, Bing, DuckDuckGo and Yandex. Social sites such as Facebook and Instagram are not blocked. No DNS queries are logged. Free service.

Warning: This server is incompatible with anonymization.

Provided by: https://dnsforfamily.com

sdns://AQMAAAAAAAAADjk0LjEzMC4xODAuMjI1ILtn1Ada3rLi6VNcj4pB-I5eHBqFzFbs_XFRHG-6KenTEGRuc2ZvcmZhbWlseS5jb20
sdns://AQMAAAAAAAAADDc4LjQ3LjY0LjE2MSATJeLOABXNSYcSJIoqR5_iUYz87Y4OecMLB84aEAKPrRBkbnNmb3JmYW1pbHkuY29t


## dnsforfamily-doh

DNS for Family safe-search resolver.
Blocks adult, gambling, drug, malware, scam/phishing, proxy/anonymizer and advertising domains; its blacklist is updated daily. Enforces safe search in Google, YouTube, Brave, Ecosia, Bing, DuckDuckGo and Yandex. Social sites such as Facebook and Instagram are not blocked. No DNS queries are logged. Free service.

Provided by: https://dnsforfamily.com

sdns://AgMAAAAAAAAADzE2Ny4yMzUuMjM2LjEwNyAy7bsRzCWPvjPCzSShSScPC-b0RvVyZLO9HCW5hTMnLhhkbnMtZG9oLmRuc2ZvcmZhbWlseS5jb20KL2Rucy1xdWVyeQ


## dnsforfamily-doh-no-safe-search

DNS for Family no-safe-search resolver.
Blocks adult, gambling, drug, malware, scam/phishing, proxy/anonymizer and advertising domains; its blacklist is updated daily. Unlike the default DNS for Family resolvers, this variant does not enforce safe search in Google, YouTube, Brave, Ecosia, Bing, DuckDuckGo or Yandex. Social sites such as Facebook and Instagram are not blocked. No DNS queries are logged. Free service.

Warning: This server is incompatible with anonymization.

Provided by: https://dnsforfamily.com

sdns://AgMAAAAAAAAADzE2Ny4yMzUuMjM2LjEwNyAy7bsRzCWPvjPCzSShSScPC-b0RvVyZLO9HCW5hTMnLidkbnMtZG9oLW5vLXNhZmUtc2VhcmNoLmRuc2ZvcmZhbWlseS5jb20KL2Rucy1xdWVyeQ


## dnsforfamily-no-safe-search

DNS for Family no-safe-search resolver.
Blocks adult, gambling, drug, malware, scam/phishing, proxy/anonymizer and advertising domains; its blacklist is updated daily. Unlike the default DNS for Family resolvers, this variant does not enforce safe search in Google, YouTube, Brave, Ecosia, Bing, DuckDuckGo or Yandex. Social sites such as Facebook and Instagram are not blocked. No DNS queries are logged. Free service.

Warning: This server is incompatible with anonymization.

Provided by: https://dnsforfamily.com

sdns://AQMAAAAAAAAADzEzNS4xODEuMTkzLjIyMiDrxcZ_hFtGE6tfATvQZYjxgl5pTY_e2cRH_ms8bEWofBBkbnNmb3JmYW1pbHkuY29t


## doh-cleanbrowsing-adult

CleanBrowsing Adult Filter.
Blocks adult, pornographic and explicit sites. Allows proxy and VPN domains and mixed-content sites. Google and Bing are set to Safe Mode.

Operated by CleanBrowsing. Service page: https://cleanbrowsing.org/filters/

sdns://AgMAAAAAAAAAAAAVZG9oLmNsZWFuYnJvd3Npbmcub3JnEi9kb2gvYWR1bHQtZmlsdGVyLw


## doh-cleanbrowsing-family

CleanBrowsing Family Filter.
Blocks adult and explicit sites, proxy and VPN bypass domains, and mixed-content sites. Google, Bing and YouTube are set to Safe Mode.

Operated by CleanBrowsing. Service page: https://cleanbrowsing.org/filters/

sdns://AgMAAAAAAAAAAAAVZG9oLmNsZWFuYnJvd3Npbmcub3JnEy9kb2gvZmFtaWx5LWZpbHRlci8


## mullvad-family-doh

Mullvad family filtering resolver.
No logging, DNSSEC-capable, global anycast network. Blocks ads, trackers, malware, adult content and gambling.
Operated by Mullvad. Service page: https://mullvad.net/en/help/dns-over-https-and-dns-over-tls/

sdns://AgMAAAAAAAAACzE5NC4yNDIuMi42ABZmYW1pbHkuZG5zLm11bGx2YWQubmV0Ci9kbnMtcXVlcnk


## sfw.scaleway-fr

Scaleway family-safety resolver.
Uses deep learning to block adult websites. Free, DNSSEC, no logs.
Maintained by Frank Denis (@jedisct1) and hosted by Scaleway. Provider: https://www.scaleway.com/

sdns://AQMAAAAAAAAADzE2My4xNzIuMTgwLjEyNSDfYnO_x1IZKotaObwMhaw_-WRF1zZE9mJygl01WPGh_x8yLmRuc2NyeXB0LWNlcnQuc2Z3LnNjYWxld2F5LWZy

