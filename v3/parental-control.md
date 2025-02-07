# parental-control

A set of resolvers blocking popular websites that may not be appropriate
for children.

This is not bulletproof. In particular, websites in languages that are
not English will require additional, local rules.

To use that list, add this to the `[sources]` section of your
`dnscrypt-proxy.toml` configuration file:

    [sources.'parental-control']
    urls = ['https://raw.githubusercontent.com/DNSCrypt/dnscrypt-resolvers/master/v3/parental-control.md', 'https://download.dnscrypt.info/resolvers-list/v3/parental-control.md']
    minisign_key = 'RWQf6LRCGA9i53mlYecO4IzT51TGPpvWucNSCh1CBM0QTaLn73Y7GFO3'
    cache_file = 'parental-control.md'

In order to enforce safe search results from Google and Youtube, you may
also want to enable cloaking (`cloaking_rules` in the configuration file).

--


## adguard-dns-family

AdGuard DNS with safesearch and adult content blocking

Warning: This server is incompatible with anonymization.

sdns://AQMAAAAAAAAAETk0LjE0MC4xNC4xNTo1NDQzILgxXdexS27jIKRw3C7Wsao5jMnlhvhdRUXWuMm1AFq6ITIuZG5zY3J5cHQuZmFtaWx5Lm5zMS5hZGd1YXJkLmNvbQ
sdns://AQMAAAAAAAAAETk0LjE0MC4xNS4xNjo1NDQzILgxXdexS27jIKRw3C7Wsao5jMnlhvhdRUXWuMm1AFq6ITIuZG5zY3J5cHQuZmFtaWx5Lm5zMS5hZGd1YXJkLmNvbQ


## adguard-dns-family-doh

AdGuard DNS with safesearch and adult content blocking (over DoH)

sdns://AgMAAAAAAAAADDk0LjE0MC4xNC4xNSCaOjT3J965vKUQA9nOnDn48n3ZxSQpAcK6saROY1oCGQw5NC4xNDAuMTQuMTUKL2Rucy1xdWVyeQ
sdns://AgMAAAAAAAAADDk0LjE0MC4xNS4xNiCaOjT3J965vKUQA9nOnDn48n3ZxSQpAcK6saROY1oCGQw5NC4xNDAuMTUuMTYKL2Rucy1xdWVyeQ


## adguard-dns-family-doh-ipv6

AdGuard DNS with safesearch and adult content blocking (over DoH, over IPv6)

sdns://AgMAAAAAAAAAFFsyYTEwOjUwYzA6OmJhZDE6ZmZdIJo6NPcn3rm8pRAD2c6cOfjyfdnFJCkBwrqxpE5jWgIZFmZhbWlseS5hZGd1YXJkLWRucy5jb20KL2Rucy1xdWVyeQ
sdns://AgMAAAAAAAAAFFsyYTEwOjUwYzA6OmJhZDI6ZmZdIJo6NPcn3rm8pRAD2c6cOfjyfdnFJCkBwrqxpE5jWgIZFmZhbWlseS5hZGd1YXJkLWRucy5jb20KL2Rucy1xdWVyeQ


## adguard-dns-family-ipv6

AdGuard DNS with safesearch and adult content blocking (over IPv6)

Warning: This server is incompatible with anonymization.

sdns://AQMAAAAAAAAAGVsyYTEwOjUwYzA6OmJhZDE6ZmZdOjU0NDMguDFd17FLbuMgpHDcLtaxqjmMyeWG-F1FRda4ybUAWrohMi5kbnNjcnlwdC5mYW1pbHkubnMxLmFkZ3VhcmQuY29t
sdns://AQMAAAAAAAAAGVsyYTEwOjUwYzA6OmJhZDI6ZmZdOjU0NDMguDFd17FLbuMgpHDcLtaxqjmMyeWG-F1FRda4ybUAWrohMi5kbnNjcnlwdC5mYW1pbHkubnMxLmFkZ3VhcmQuY29t


## cisco-familyshield

Block websites not suitable for children (DNSCrypt protocol)

Warning: Doesn't work any more in some countries such as France and Portugal.

Warning: modifies your queries to include a copy of your network
address when forwarding them to a selection of companies and organizations.

Currently incompatible with DNS anonymization.

sdns://AQEAAAAAAAAADjIwOC42Ny4yMjAuMTIzILc1EUAgbyJdPivYItf9aR6hwzzI1maNDL4Ev6vKQ_t5GzIuZG5zY3J5cHQtY2VydC5vcGVuZG5zLmNvbQ


## cisco-familyshield-ipv6

Block websites not suitable for children (IPv6)

Warning: Doesn't work any more in some countries such as France and Portugal.

Warning: This server is incompatible with anonymization.

Warning: modifies your queries to include a copy of your network
address when forwarding them to a selection of companies and organizations.

sdns://AQEAAAAAAAAAEVsyNjIwOjExOTozNTo6MzVdILc1EUAgbyJdPivYItf9aR6hwzzI1maNDL4Ev6vKQ_t5GzIuZG5zY3J5cHQtY2VydC5vcGVuZG5zLmNvbQ


## cleanbrowsing-adult

Blocks access to adult, pornographic and explicit sites. It does
not block proxy or VPNs, nor mixed-content sites. Sites like Reddit
are allowed. Google and Bing are set to the Safe Mode.

Warning: This server is incompatible with anonymization.

By https://cleanbrowsing.org/

sdns://AQMAAAAAAAAAEzE4NS4yMjguMTY4LjEwOjg0NDMgvKwy-tVDaRcfCDLWB1AnwyCM7vDo6Z-UGNx3YGXUjykRY2xlYW5icm93c2luZy5vcmc
sdns://AQMAAAAAAAAAEzE4NS4yMjguMTY5LjExOjg0NDMgvKwy-tVDaRcfCDLWB1AnwyCM7vDo6Z-UGNx3YGXUjykRY2xlYW5icm93c2luZy5vcmc


## cleanbrowsing-adult-doh

Blocks access to adult, pornographic and explicit sites over DoH. It does
not block proxy or VPNs, nor mixed-content sites. Sites like Reddit
are allowed. Google and Bing are set to the Safe Mode.

sdns://AgMAAAAAAAAADjE4NS4yMjguMTY4LjEwoPn_N_AuYyy3OHAlwH5XkIo9Nxt8ldjN0DkN4jHtlDoSoCso0AXN1mJZ2xEYZeoXy7YLPI9UcGhjjZAqZL54Sv34IOaSTdvwPj_u_RiUGT7gQuBqadbySK2eIW2kKyiPLBAZEWNsZWFuYnJvd3Npbmcub3JnES9kb2gvYWR1bHQtZmlsdGVy
sdns://AgMAAAAAAAAADzE4NS4yMjguMTY4LjE2OKD5_zfwLmMstzhwJcB-V5CKPTcbfJXYzdA5DeIx7ZQ6EqArKNAFzdZiWdsRGGXqF8u2CzyPVHBoY42QKmS-eEr9-CDmkk3b8D4_7v0YlBk-4ELgamnW8kitniFtpCsojywQGRFjbGVhbmJyb3dzaW5nLm9yZxEvZG9oL2FkdWx0LWZpbHRlcg


## cleanbrowsing-adult-ipv6

Blocks access to adult, pornographic and explicit sites over IPv6. It does
not block proxy or VPNs, nor mixed-content sites. Sites like Reddit
are allowed. Google and Bing are set to the Safe Mode.

Warning: This server is incompatible with anonymization.

sdns://AQMAAAAAAAAAFVsyYTBkOjJhMDA6MTo6MV06ODQ0MyC8rDL61UNpFx8IMtYHUCfDIIzu8Ojpn5QY3HdgZdSPKRFjbGVhbmJyb3dzaW5nLm9yZw
sdns://AQMAAAAAAAAAFVsyYTBkOjJhMDA6Mjo6MV06ODQ0MyC8rDL61UNpFx8IMtYHUCfDIIzu8Ojpn5QY3HdgZdSPKRFjbGVhbmJyb3dzaW5nLm9yZw


## cleanbrowsing-family

Blocks access to adult, pornographic and explicit sites. It also
blocks proxy and VPN domains that are used to bypass the filters.
Mixed content sites (like Reddit) are also blocked. Google, Bing and
Youtube are set to the Safe Mode.

Warning: This server is incompatible with anonymization.

By https://cleanbrowsing.org/

sdns://AQMAAAAAAAAAFDE4NS4yMjguMTY4LjE2ODo4NDQzILysMvrVQ2kXHwgy1gdQJ8MgjO7w6OmflBjcd2Bl1I8pEWNsZWFuYnJvd3Npbmcub3Jn
sdns://AQMAAAAAAAAAFDE4NS4yMjguMTY5LjE2ODo4NDQzILysMvrVQ2kXHwgy1gdQJ8MgjO7w6OmflBjcd2Bl1I8pEWNsZWFuYnJvd3Npbmcub3Jn


## cleanbrowsing-family-doh

Blocks access to adult, pornographic and explicit sites over DoH. It also
blocks proxy and VPN domains that are used to bypass the filters.
Mixed content sites (like Reddit) are also blocked. Google, Bing and
Youtube are set to the Safe Mode.

sdns://AgMAAAAAAAAADjE4NS4yMjguMTY4LjEwoPn_N_AuYyy3OHAlwH5XkIo9Nxt8ldjN0DkN4jHtlDoSoCso0AXN1mJZ2xEYZeoXy7YLPI9UcGhjjZAqZL54Sv34IOaSTdvwPj_u_RiUGT7gQuBqadbySK2eIW2kKyiPLBAZEWNsZWFuYnJvd3Npbmcub3JnEi9kb2gvZmFtaWx5LWZpbHRlcg
sdns://AgMAAAAAAAAADzE4NS4yMjguMTY4LjE2OKD5_zfwLmMstzhwJcB-V5CKPTcbfJXYzdA5DeIx7ZQ6EqArKNAFzdZiWdsRGGXqF8u2CzyPVHBoY42QKmS-eEr9-CDmkk3b8D4_7v0YlBk-4ELgamnW8kitniFtpCsojywQGRFjbGVhbmJyb3dzaW5nLm9yZxIvZG9oL2ZhbWlseS1maWx0ZXI


## cleanbrowsing-family-ipv6

Blocks access to adult, pornographic and explicit sites over IPv6. It also
blocks proxy and VPN domains that are used to bypass the filters.
Mixed content sites (like Reddit) are also blocked. Google, Bing and
Youtube are set to the Safe Mode.

Warning: This server is incompatible with anonymization.

sdns://AQMAAAAAAAAAFFsyYTBkOjJhMDA6MTo6XTo4NDQzILysMvrVQ2kXHwgy1gdQJ8MgjO7w6OmflBjcd2Bl1I8pEWNsZWFuYnJvd3Npbmcub3Jn
sdns://AQMAAAAAAAAAFFsyYTBkOjJhMDA6Mjo6XTo4NDQzILysMvrVQ2kXHwgy1gdQJ8MgjO7w6OmflBjcd2Bl1I8pEWNsZWFuYnJvd3Npbmcub3Jn


## cloudflare-family

Cloudflare DNS (anycast) with malware protection and parental control - aka 1.1.1.3 / 1.0.0.3

sdns://AgMAAAAAAAAABzEuMS4xLjMABzEuMS4xLjMKL2Rucy1xdWVyeQ
sdns://AgMAAAAAAAAABzEuMC4wLjMABzEuMC4wLjMKL2Rucy1xdWVyeQ


## cloudflare-family-ipv6

Cloudflare DNS over IPv6 (anycast) with malware protection and parental control

sdns://AgMAAAAAAAAAAAAWWzI2MDY6NDcwMDo0NzAwOjoxMTEzXQovZG5zLXF1ZXJ5
sdns://AgMAAAAAAAAAAAAWWzI2MDY6NDcwMDo0NzAwOjoxMDAzXQovZG5zLXF1ZXJ5


## dns0-kids

A free, sovereign and GDPR-compliant recursive DNS resolver with a
strong focus on security to protect the citizens and organizations of
the European Union.

This version blocks content not suitable for children.

Also blocks new domains, dynamic DNS, parked domains, uncommon TLDs, etc.

https://www.dns0.eu/

sdns://AgMAAAAAAAAAACCaOjT3J965vKUQA9nOnDn48n3ZxSQpAcK6saROY1oCGQxraWRzLmRuczAuZXUKL2Rucy1xdWVyeQ


## dnsforfamily

(DNSCrypt Protocol) (Now supports DNSSEC). Block adult websites, gambling websites, malwares, trackers and advertisements.
It also enforces safe search in: Google, YouTube, Bing, DuckDuckGo and Yandex.

Social websites like Facebook and Instagram are not blocked. No DNS queries are logged.

As of 26-May-2022 5.9 million websites are blocked and new websites are added to blacklist daily.
Completely free, no ads or any commercial motive. Operating for 4 years now.

Warning: This server is incompatible with anonymization.

Provided by: https://dnsforfamily.com

sdns://AQMAAAAAAAAADDc4LjQ3LjY0LjE2MSATJeLOABXNSYcSJIoqR5_iUYz87Y4OecMLB84aEAKPrRBkbnNmb3JmYW1pbHkuY29t


## dnsforfamily-doh

(DoH Protocol) (Now supports DNSSEC). Block adult websites, gambling websites, malwares, trackers and advertisements.
It also enforces safe search in: Google, YouTube, Bing, DuckDuckGo and Yandex.

Social websites like Facebook and Instagram are not blocked. No DNS queries are logged.

As of 26-May-2022 5.9 million websites are blocked and new websites are added to blacklist daily.
Completely free, no ads or any commercial motive. Operating for 4 years now.

Provided by: https://dnsforfamily.com

sdns://AgMAAAAAAAAADzE2Ny4yMzUuMjM2LjEwN6DMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5qD39a8ogB6HfX3zor4o_hSfMnSn7SnxNv1QPOOYy6gLDaCzXgWVDPvYJZ-TdboQGHC4w8TA7FpSl1fmMWfp3iap6KBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33zqDmRLppY-M1_nZcuZdrErEOtUKUtCR3dkzLOjrMo6yy_CAqrj-3vwXkyBxBlNykRRHU-a8wR4bsGuchhAnPYqCDVRhkbnMtZG9oLmRuc2ZvcmZhbWlseS5jb20KL2Rucy1xdWVyeQ


## dnsforfamily-doh-no-safe-search

(DoH Protocol) (Now supports DNSSEC) Block adult websites, gambling websites, malwares, trackers and advertisements.
Unlike other dnsforfamily servers, this one does not enforces safe search. So Google, YouTube, Bing, DuckDuckGo and Yandex are completely accessible without any restriction.

Social websites like Facebook and Instagram are not blocked. No DNS queries are logged.

As of 26-May-2022 5.9 million websites are blocked and new websites are added to blacklist daily.
Completely free, no ads or any commercial motive. Operating for 4 years now.

Warning: This server is incompatible with anonymization.

Provided by: https://dnsforfamily.com

sdns://AgMAAAAAAAAADzE2Ny4yMzUuMjM2LjEwN6DMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5qD39a8ogB6HfX3zor4o_hSfMnSn7SnxNv1QPOOYy6gLDaCzXgWVDPvYJZ-TdboQGHC4w8TA7FpSl1fmMWfp3iap6KBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33zqDmRLppY-M1_nZcuZdrErEOtUKUtCR3dkzLOjrMo6yy_CAqrj-3vwXkyBxBlNykRRHU-a8wR4bsGuchhAnPYqCDVSdkbnMtZG9oLW5vLXNhZmUtc2VhcmNoLmRuc2ZvcmZhbWlseS5jb20KL2Rucy1xdWVyeQ


## dnsforfamily-no-safe-search

(DNSCrypt Protocol) (Now supports DNSSEC) Block adult websites, gambling websites, malwares, trackers and advertisements.
Unlike other dnsforfamily servers, this one does not enforces safe search. So Google, YouTube, Bing, DuckDuckGo and Yandex are completely accessible without any restriction.

Social websites like Facebook and Instagram are not blocked. No DNS queries are logged.

As of 26-May-2022 5.9 million websites are blocked and new websites are added to blacklist daily.
Completely free, no ads or any commercial motive. Operating for 4 years now.

Warning: This server is incompatible with anonymization.

Provided by: https://dnsforfamily.com

sdns://AQMAAAAAAAAADzEzNS4xODEuMTkzLjIyMiDrxcZ_hFtGE6tfATvQZYjxgl5pTY_e2cRH_ms8bEWofBBkbnNmb3JmYW1pbHkuY29t


## dnsforfamily-v6

(DNSCrypt Protocol) (Now supports DNSSEC) Block adult websites, gambling websites, malwares, trackers and advertisements.
It also enforces safe search in: Google, YouTube, Bing, DuckDuckGo and Yandex.

Social websites like Facebook and Instagram are not blocked. No DNS queries are logged.

As of 26-May-2022 5.9 million websites are blocked and new websites are added to blacklist daily.
Completely free, no ads or any commercial motive. Operating for 4 years now.

Provided by: https://dnsforfamily.com

sdns://AQMAAAAAAAAAF1syYTAxOjRmODoxYzE3OjRkZjg6OjFdIBMl4s4AFc1JhxIkiipHn-JRjPztjg55wwsHzhoQAo-tEGRuc2ZvcmZhbWlseS5jb20


## doh-cleanbrowsing-adult

Blocks access to adult, pornographic and explicit sites. It does
not block proxy or VPNs, nor mixed-content sites. Sites like Reddit
are allowed. Google and Bing are set to the Safe Mode.

By https://cleanbrowsing.org/

sdns://AgMAAAAAAAAAAAAVZG9oLmNsZWFuYnJvd3Npbmcub3JnEi9kb2gvYWR1bHQtZmlsdGVyLw


## doh-cleanbrowsing-family

Blocks access to adult, pornographic and explicit sites. It also
blocks proxy and VPN domains that are used to bypass the filters.
Mixed content sites (like Reddit) are also blocked. Google, Bing and
Youtube are set to the Safe Mode.

By https://cleanbrowsing.org/

sdns://AgMAAAAAAAAAAAAVZG9oLmNsZWFuYnJvd3Npbmcub3JnEy9kb2gvZmFtaWx5LWZpbHRlci8


## mullvad-family-doh

Same as mullvad-doh but blocks ads, trackers, malware, adult content, and gambling.

sdns://AgMAAAAAAAAACzE5NC4yNDIuMi42ABZmYW1pbHkuZG5zLm11bGx2YWQubmV0Ci9kbnMtcXVlcnk


## safesurfer

Family safety focused blocklist for over 2 million adult sites, as well as phishing and malware and more.
Free to use, paid for customizing blocking for more categories+sites and viewing usage at my.safesurfer.io. Logs taken for viewing
usage, data never sold - https://safesurfer.io

Warning: this server is incompatible with DNS anonymization.

sdns://AQIAAAAAAAAADzEwNC4xNTUuMjM3LjIyNSAnIH_VEgToNntINABd-f_R0wu-KpwzY55u2_iu2R1A2CAyLmRuc2NyeXB0LWNlcnQuc2FmZXN1cmZlci5jby5ueg


## sfw.scaleway-fr

Uses deep learning to block adult websites. Free, DNSSEC, no logs.
Hosted in Paris, running on a 1-XS server donated by Scaleway.com

Maintained by Frank Denis - https://fr.dnscrypt.info/sfw.html

sdns://AQMAAAAAAAAADzE2My4xNzIuMTgwLjEyNSDfYnO_x1IZKotaObwMhaw_-WRF1zZE9mJygl01WPGh_x8yLmRuc2NyeXB0LWNlcnQuc2Z3LnNjYWxld2F5LWZy

