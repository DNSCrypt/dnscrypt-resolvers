# parental-control

A set of resolvers blocking popular websites that may not be appropriate
for children.

This is not bulletproof. In particular, websites in languages that are
not English will require additional, local rules.

To use that list, add this to the `[sources]` section of your
`dnscrypt-proxy.toml` configuration file:

    [sources.'parental-control']
    urls = ['https://raw.githubusercontent.com/DNSCrypt/dnscrypt-resolvers/master/v2/parental-control.md', 'https://download.dnscrypt.info/resolvers-list/v2/parental-control.md']
    minisign_key = 'RWQf6LRCGA9i53mlYecO4IzT51TGPpvWucNSCh1CBM0QTaLn73Y7GFO3'
    cache_file = 'parental-control.md'

In order to enforce safe search results from Google and Youtube, you may
also want to enable cloaking (`cloaking_rules` in the configuration file).

--


## adguard-dns-family

Adguard DNS with safesearch and adult content blocking

sdns://AQIAAAAAAAAAFDE3Ni4xMDMuMTMwLjEzMjo1NDQzILgxXdexS27jIKRw3C7Wsao5jMnlhvhdRUXWuMm1AFq6ITIuZG5zY3J5cHQuZmFtaWx5Lm5zMS5hZGd1YXJkLmNvbQ


## af-dnswarden-dc1

DnsCrypt protocol . Adult content filtering, Non-logging, AD-filtering, enforces force safe search for search engines and youtube, supports DNSSEC. By https://dnswarden.com

sdns://AQMAAAAAAAAAEzExNi4yMDMuNzAuMTU2OjE0NDMgenKjVeH-LU7Opsyq1ljKZz14fHsngOK8OOeQ-cR2mAsoMi5kbnNjcnlwdC1jZXJ0LmRuc3dhcmRlbi0xLWFkdWx0LWZpbHRlcg


## af-dnswarden-dc1-ecs
DnsCrypt protocol . Adult content filtering, Non-logging, AD-filtering, enforces force safe search for search engines and youtube, supports DNSSEC, ECS is enabled and self resolves as of now. By https://dnswarden.com

sdns://AQMAAAAAAAAAEzExNi4yMDMuNzAuMTU2OjU0NDMgenKjVeH-LU7Opsyq1ljKZz14fHsngOK8OOeQ-cR2mAssMi5kbnNjcnlwdC1jZXJ0LmRuc3dhcmRlbi0xLWFkdWx0LWZpbHRlci1lY3M


## af-dnswarden-dc1-ipv6

DnsCrypt protocol . Adult content filtering, Non-logging, AD-filtering, enforces force safe search for search engines and youtube, supports DNSSEC. By https://dnswarden.com

sdns://AQMAAAAAAAAAHFsyYTAxOjRmODoxYzFjOjc1YjQ6OjFdOjE0NDMgenKjVeH-LU7Opsyq1ljKZz14fHsngOK8OOeQ-cR2mAsoMi5kbnNjcnlwdC1jZXJ0LmRuc3dhcmRlbi0xLWFkdWx0LWZpbHRlcg


## af-dnswarden-dc1-ipv6-ecs
DnsCrypt protocol . Adult content filtering, Non-logging, AD-filtering, enforces force safe search for search engines and youtube, supports DNSSEC, ECS is enabled and self resolves as of now. By https://dnswarden.com

sdns://AQMAAAAAAAAAHFsyYTAxOjRmODoxYzFjOjc1YjQ6OjFdOjU0NDMgenKjVeH-LU7Opsyq1ljKZz14fHsngOK8OOeQ-cR2mAssMi5kbnNjcnlwdC1jZXJ0LmRuc3dhcmRlbi0xLWFkdWx0LWZpbHRlci1lY3M


## af-dnswarden-dc2

DnsCrypt protocol . Adult content filtering, Non-logging, AD-filtering, enforces force safe search for search engines and youtube, supports DNSSEC. By https://dnswarden.com

sdns://AQMAAAAAAAAAEzExNi4yMDMuMzUuMjU1OjE0NDMg-IlTTFFgMuUntnNV78COzbhJN9_OEOVWNgHhdg4BNXwoMi5kbnNjcnlwdC1jZXJ0LmRuc3dhcmRlbi0yLWFkdWx0LWZpbHRlcg


## af-dnswarden-dc2-ecs
DnsCrypt protocol . Adult content filtering, Non-logging, AD-filtering, enforces force safe search for search engines and youtube, supports DNSSEC, ECS is enabled and self resolves as of now. By https://dnswarden.com

sdns://AQMAAAAAAAAAEzExNi4yMDMuMzUuMjU1OjU0NDMg-IlTTFFgMuUntnNV78COzbhJN9_OEOVWNgHhdg4BNXwsMi5kbnNjcnlwdC1jZXJ0LmRuc3dhcmRlbi0yLWFkdWx0LWZpbHRlci1lY3M


## af-dnswarden-dc2-ipv6

DnsCrypt protocol . Adult content filtering, Non-logging, AD-filtering, enforces force safe search for search engines and youtube, supports DNSSEC. By https://dnswarden.com

sdns://AQMAAAAAAAAAHFsyYTAxOjRmODoxYzFjOjVlNzc6OjFdOjE0NDMg-IlTTFFgMuUntnNV78COzbhJN9_OEOVWNgHhdg4BNXwoMi5kbnNjcnlwdC1jZXJ0LmRuc3dhcmRlbi0yLWFkdWx0LWZpbHRlcg


## af-dnswarden-dc2-ipv6-ecs
DnsCrypt protocol . Adult content filtering, Non-logging, AD-filtering, enforces force safe search for search engines and youtube, supports DNSSEC, ECS is enabled and self resolves as of now. By https://dnswarden.com

sdns://AQMAAAAAAAAAHFsyYTAxOjRmODoxYzFjOjVlNzc6OjFdOjU0NDMg-IlTTFFgMuUntnNV78COzbhJN9_OEOVWNgHhdg4BNXwsMi5kbnNjcnlwdC1jZXJ0LmRuc3dhcmRlbi0yLWFkdWx0LWZpbHRlci1lY3M


## af-dnswarden-dc3

DnsCrypt protocol . Adult content filtering, Non-logging, AD-filtering, enforces force safe search for search engines and youtube, supports DNSSEC. By https://dnswarden.com

sdns://AQMAAAAAAAAAETg4LjE5OC4xNjEuODoxNDQzIBc_e3j6AbaHocJHqJVzfpZ7xi2puAm6tL5xZTRrMTPbKDIuZG5zY3J5cHQtY2VydC5kbnN3YXJkZW4tMy1hZHVsdC1maWx0ZXI


## af-dnswarden-dc3-ecs
DnsCrypt protocol . Adult content filtering, Non-logging, AD-filtering, enforces force safe search for search engines and youtube, supports DNSSEC, ECS is enabled and self resolves as of now. By https://dnswarden.com

sdns://AQMAAAAAAAAAETg4LjE5OC4xNjEuODo1NDQzIBc_e3j6AbaHocJHqJVzfpZ7xi2puAm6tL5xZTRrMTPbLDIuZG5zY3J5cHQtY2VydC5kbnN3YXJkZW4tMy1hZHVsdC1maWx0ZXItZWNz


## af-dnswarden-dc3-ipv6

DnsCrypt protocol . Adult content filtering, Non-logging, AD-filtering, enforces force safe search for search engines and youtube, supports DNSSEC. By https://dnswarden.com

sdns://AQMAAAAAAAAAG1syYTAxOjRmODpjMGM6NjJhNzo6MV06MTQ0MyAXP3t4-gG2h6HCR6iVc36We8YtqbgJurS-cWU0azEz2ygyLmRuc2NyeXB0LWNlcnQuZG5zd2FyZGVuLTMtYWR1bHQtZmlsdGVy


## af-dnswarden-dc3-ipv6-ecs
DnsCrypt protocol . Adult content filtering, Non-logging, AD-filtering, enforces force safe search for search engines and youtube, supports DNSSEC, ECS is enabled and self resolves as of now. By https://dnswarden.com

sdns://AQMAAAAAAAAAG1syYTAxOjRmODpjMGM6NjJhNzo6MV06NTQ0MyAXP3t4-gG2h6HCR6iVc36We8YtqbgJurS-cWU0azEz2ywyLmRuc2NyeXB0LWNlcnQuZG5zd2FyZGVuLTMtYWR1bHQtZmlsdGVyLWVjcw


## af-dnswarden-doh1

Dns-over-HTTPS protocol . Adult content filtering, Non-logging, AD-filtering, enforces force safe search for search engines and youtube, supports DNSSEC. By https://dnswarden.com

sdns://AgMAAAAAAAAADjExNi4yMDMuNzAuMTU2ID4aGg9sU_PpekktVwhLW5gHBZ7gV6sVBYdv2D_aPbg4EWRvaC5kbnN3YXJkZW4uY29tDS9hZHVsdC1maWx0ZXI


## af-dnswarden-doh1-ecs
Dns-over-HTTPS protocol . Adult content filtering, Non-logging, AD-filtering, enforces force safe search for search engines and youtube, supports DNSSEC, ECS is enabled and self resolves as of now. By https://dnswarden.com

sdns://AgMAAAAAAAAADjExNi4yMDMuNzAuMTU2ID4aGg9sU_PpekktVwhLW5gHBZ7gV6sVBYdv2D_aPbg4FWVjcy1kb2guZG5zd2FyZGVuLmNvbREvYWR1bHQtZmlsdGVyLWVjcw


## af-dnswarden-doh1-ipv6

Dns-over-HTTPS protocol . Adult content filtering, Non-logging, AD-filtering, enforces force safe search for search engines and youtube, supports DNSSEC. By https://dnswarden.com

sdns://AgMAAAAAAAAAF1syYTAxOjRmODoxYzFjOjc1YjQ6OjFdID4aGg9sU_PpekktVwhLW5gHBZ7gV6sVBYdv2D_aPbg4EWRvaC5kbnN3YXJkZW4uY29tDS9hZHVsdC1maWx0ZXI


## af-dnswarden-doh1-ipv6-ecs
Dns-over-HTTPS protocol . Adult content filtering, Non-logging, AD-filtering, enforces force safe search for search engines and youtube, supports DNSSEC, ECS is enabled and self resolves as of now. By https://dnswarden.com

sdns://AgMAAAAAAAAAF1syYTAxOjRmODoxYzFjOjc1YjQ6OjFdID4aGg9sU_PpekktVwhLW5gHBZ7gV6sVBYdv2D_aPbg4FWVjcy1kb2guZG5zd2FyZGVuLmNvbREvYWR1bHQtZmlsdGVyLWVjcw


## af-dnswarden-doh2

Dns-over-HTTPS protocol . Adult content filtering, Non-logging, AD-filtering, enforces force safe search for search engines and youtube, supports DNSSEC. By https://dnswarden.com

sdns://AgMAAAAAAAAADjExNi4yMDMuMzUuMjU1ID4aGg9sU_PpekktVwhLW5gHBZ7gV6sVBYdv2D_aPbg4EWRvaC5kbnN3YXJkZW4uY29tDS9hZHVsdC1maWx0ZXI


## af-dnswarden-doh2-ecs
Dns-over-HTTPS protocol . Adult content filtering, Non-logging, AD-filtering, enforces force safe search for search engines and youtube, supports DNSSEC, ECS is enabled and self resolves as of now. By https://dnswarden.com

sdns://AgMAAAAAAAAADjExNi4yMDMuMzUuMjU1ID4aGg9sU_PpekktVwhLW5gHBZ7gV6sVBYdv2D_aPbg4FWVjcy1kb2guZG5zd2FyZGVuLmNvbREvYWR1bHQtZmlsdGVyLWVjcw


## af-dnswarden-doh2-ipv6

Dns-over-HTTPS protocol . Adult content filtering, Non-logging, AD-filtering, enforces force safe search for search engines and youtube, supports DNSSEC. By https://dnswarden.com

sdns://AgMAAAAAAAAAF1syYTAxOjRmODoxYzFjOjVlNzc6OjFdID4aGg9sU_PpekktVwhLW5gHBZ7gV6sVBYdv2D_aPbg4EWRvaC5kbnN3YXJkZW4uY29tDS9hZHVsdC1maWx0ZXI


## af-dnswarden-doh2-ipv6-ecs
Dns-over-HTTPS protocol . Adult content filtering, Non-logging, AD-filtering, enforces force safe search for search engines and youtube, supports DNSSEC, ECS is enabled and self resolves as of now. By https://dnswarden.com

sdns://AgMAAAAAAAAAF1syYTAxOjRmODoxYzFjOjVlNzc6OjFdID4aGg9sU_PpekktVwhLW5gHBZ7gV6sVBYdv2D_aPbg4FWVjcy1kb2guZG5zd2FyZGVuLmNvbREvYWR1bHQtZmlsdGVyLWVjcw


## af-dnswarden-doh3

Dns-over-HTTPS protocol . Adult content filtering, Non-logging, AD-filtering, enforces force safe search for search engines and youtube, supports DNSSEC. By https://dnswarden.com

sdns://AgMAAAAAAAAADDg4LjE5OC4xNjEuOCA-GhoPbFPz6XpJLVcIS1uYBwWe4FerFQWHb9g_2j24OBFkb2guZG5zd2FyZGVuLmNvbQ0vYWR1bHQtZmlsdGVy


## af-dnswarden-doh3-ecs
Dns-over-HTTPS protocol . Adult content filtering, Non-logging, AD-filtering, enforces force safe search for search engines and youtube, supports DNSSEC, ECS is enabled and self resolves as of now. By https://dnswarden.com

sdns://AgMAAAAAAAAADDg4LjE5OC4xNjEuOCA-GhoPbFPz6XpJLVcIS1uYBwWe4FerFQWHb9g_2j24OBVlY3MtZG9oLmRuc3dhcmRlbi5jb20RL2FkdWx0LWZpbHRlci1lY3M


## af-dnswarden-doh3-ipv6

Dns-over-HTTPS protocol . Adult content filtering, Non-logging, AD-filtering, enforces force safe search for search engines and youtube, supports DNSSEC. By https://dnswarden.com

sdns://AgMAAAAAAAAAFlsyYTAxOjRmODpjMGM6NjJhNzo6MV0gPhoaD2xT8-l6SS1XCEtbmAcFnuBXqxUFh2_YP9o9uDgRZG9oLmRuc3dhcmRlbi5jb20NL2FkdWx0LWZpbHRlcg


## af-dnswarden-doh3-ipv6-ecs
Dns-over-HTTPS protocol . Adult content filtering, Non-logging, AD-filtering, enforces force safe search for search engines and youtube, supports DNSSEC, ECS is enabled and self resolves as of now. By https://dnswarden.com

sdns://AgMAAAAAAAAAFlsyYTAxOjRmODpjMGM6NjJhNzo6MV0gPhoaD2xT8-l6SS1XCEtbmAcFnuBXqxUFh2_YP9o9uDgVZWNzLWRvaC5kbnN3YXJkZW4uY29tES9hZHVsdC1maWx0ZXItZWNz


## cisco-familyshield

Block websites not suitable for children

sdns://AQAAAAAAAAAAEjIwOC42Ny4yMjAuMTIzOjQ0MyC3NRFAIG8iXT4r2CLX_WkeocM8yNZmjQy-BL-rykP7eRsyLmRuc2NyeXB0LWNlcnQub3BlbmRucy5jb20


## cleanbrowsing-adult

Blocks access to all adult, pornographic and explicit sites. It does
not block proxy or VPNs, nor mixed-content sites. Sites like Reddit
are allowed. Google and Bing are set to the Safe Mode.

By https://cleanbrowsing.org/

sdns://AQMAAAAAAAAAEzE4NS4yMjguMTY4LjEwOjg0NDMgvKwy-tVDaRcfCDLWB1AnwyCM7vDo6Z-UGNx3YGXUjykRY2xlYW5icm93c2luZy5vcmc


## cleanbrowsing-adult-ipv6

Blocks access to all adult, pornographic and explicit sites. It does
not block proxy or VPNs, nor mixed-content sites. Sites like Reddit
are allowed. Google and Bing are set to the Safe Mode.

By https://cleanbrowsing.org/

sdns://AQMAAAAAAAAAFVsyYTBkOjJhMDA6MTo6MV06ODQ0MyC8rDL61UNpFx8IMtYHUCfDIIzu8Ojpn5QY3HdgZdSPKRFjbGVhbmJyb3dzaW5nLm9yZw


## cleanbrowsing-family

Blocks access to all adult, pornographic and explicit sites. It also
blocks proxy and VPN domains that are used to bypass the filters.
Mixed content sites (like Reddit) are also blocked. Google, Bing and
Youtube are set to the Safe Mode.

By https://cleanbrowsing.org/

sdns://AQMAAAAAAAAAFDE4NS4yMjguMTY4LjE2ODo4NDQzILysMvrVQ2kXHwgy1gdQJ8MgjO7w6OmflBjcd2Bl1I8pEWNsZWFuYnJvd3Npbmcub3Jn


## cleanbrowsing-family-ipv6

Blocks access to all adult, pornographic and explicit sites. It also
blocks proxy and VPN domains that are used to bypass the filters.
Mixed content sites (like Reddit) are also blocked. Google, Bing and
Youtube are set to the Safe Mode.

By https://cleanbrowsing.org/

sdns://AQMAAAAAAAAAFFsyYTBkOjJhMDA6MTo6XTo4NDQzILysMvrVQ2kXHwgy1gdQJ8MgjO7w6OmflBjcd2Bl1I8pEWNsZWFuYnJvd3Npbmcub3Jn


## dnsforfamily

Block adult websites, porn websites, gambling websites and advertisements.
No DNS queries are logged. As of March 2019 2.1million websites are blocked and new websites are added to blacklist daily
Completely free, no ads or any commercial motive.

Provided by: https://dnsforfamily.com

sdns://AQIAAAAAAAAADDc4LjQ3LjY0LjE2MSATJeLOABXNSYcSJIoqR5_iUYz87Y4OecMLB84aEAKPrRBkbnNmb3JmYW1pbHkuY29t


## doh-cleanbrowsing-adult

Blocks access to all adult, pornographic and explicit sites. It does
not block proxy or VPNs, nor mixed-content sites. Sites like Reddit
are allowed. Google and Bing are set to the Safe Mode.

By https://cleanbrowsing.org/

sdns://AgMAAAAAAAAAAAAVZG9oLmNsZWFuYnJvd3Npbmcub3JnEi9kb2gvYWR1bHQtZmlsdGVyLw


## doh-cleanbrowsing-family

Blocks access to all adult, pornographic and explicit sites. It also
blocks proxy and VPN domains that are used to bypass the filters.
Mixed content sites (like Reddit) are also blocked. Google, Bing and
Youtube are set to the Safe Mode.

By https://cleanbrowsing.org/

sdns://AgMAAAAAAAAAAAAVZG9oLmNsZWFuYnJvd3Npbmcub3JnEy9kb2gvZmFtaWx5LWZpbHRlci8


## sfw.scaleway-fr

Uses deep learning to block adult websites. Free, DNSSEC, no logs.
Hosted in Paris, running on a 1-XS server donated by Scaleway.com
Maintained by Frank Denis - https://fr.dnscrypt.info/sfw.html

sdns://AQMAAAAAAAAAEzE2My4xNzIuMTgwLjEyNTo0NDMg32Jzv8dSGSqLWjm8DIWsP_lkRdc2RPZicoJdNVjxof8fMi5kbnNjcnlwdC1jZXJ0LnNmdy5zY2FsZXdheS1mcg
