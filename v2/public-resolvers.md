# public-resolvers

This is an extensive list of public DNS resolvers supporting the
DNSCrypt and DNS-over-HTTP2 protocols.

This list is maintained by Frank Denis <j @ dnscrypt [.] info>

Warning: it includes servers that may censor content, servers that don't
verify DNSSEC records, and servers that will collect and monetize your
queries.

Adjust the `require_*` options in dnscrypt-proxy to filter that list
according to your needs.

To use that list, add this to the `[sources]` section of your
`dnscrypt-proxy.toml` configuration file:

    [sources.'public-resolvers']
    urls = ['https://raw.githubusercontent.com/DNSCrypt/dnscrypt-resolvers/master/v2/public-resolvers.md', 'https://download.dnscrypt.info/resolvers-list/v2/public-resolvers.md']
    minisign_key = 'RWQf6LRCGA9i53mlYecO4IzT51TGPpvWucNSCh1CBM0QTaLn73Y7GFO3'
    cache_file = 'public-resolvers.md'

--


## aaflalo-me-gcp

DNS-over-HTTPS proxy of aaflalo-me hosted in Google Cloud Platform.

Non-logging, AD-filtering, supports DNSSEC.

sdns://AgMAAAAAAAAADDM1LjIzMS42OS43NyA-GhoPbFPz6XpJLVcIS1uYBwWe4FerFQWHb9g_2j24OBJkbnMtZ2NwLmFhZmxhbG8ubWUKL2Rucy1xdWVyeQ

## aaflalo-me

DNS-over-HTTPS server running rust-doh with PiHole for Adblocking.

Non-logging, AD-filtering, supports DNSSEC.
Hosted in Netherlands on a RamNode VPS.

sdns://AgMAAAAAAAAADjE3Ni41Ni4yMzYuMTc1ID4aGg9sU_PpekktVwhLW5gHBZ7gV6sVBYdv2D_aPbg4DmRucy5hYWZsYWxvLm1lCi9kbnMtcXVlcnk

## adguard-dns-family

Adguard DNS with safesearch and adult content blocking

sdns://AQMAAAAAAAAAFDE3Ni4xMDMuMTMwLjEzMjo1NDQzILgxXdexS27jIKRw3C7Wsao5jMnlhvhdRUXWuMm1AFq6ITIuZG5zY3J5cHQuZmFtaWx5Lm5zMS5hZGd1YXJkLmNvbQ

## adguard-dns

Remove ads and protect your computer from malware

sdns://AQMAAAAAAAAAFDE3Ni4xMDMuMTMwLjEzMDo1NDQzINErR_JS3PLCu_iZEIbq95zkSV2LFsigxDIuUso_OQhzIjIuZG5zY3J5cHQuZGVmYXVsdC5uczEuYWRndWFyZC5jb20

## adguard-dns-family-doh

Adguard DNS with safesearch and adult content blocking (over DoH)

sdns://AgMAAAAAAAAADzE3Ni4xMDMuMTMwLjEzMiD5_zfwLmMstzhwJcB-V5CKPTcbfJXYzdA5DeIx7ZQ6EhZkbnMtZmFtaWx5LmFkZ3VhcmQuY29tCi9kbnMtcXVlcnk

## adguard-dns-doh

Remove ads and protect your computer from malware (over DoH)

sdns://AgMAAAAAAAAADzE3Ni4xMDMuMTMwLjEzMCD5_zfwLmMstzhwJcB-V5CKPTcbfJXYzdA5DeIx7ZQ6Eg9kbnMuYWRndWFyZC5jb20KL2Rucy1xdWVyeQ

## adguard-dns-family-ipv6

Adguard DNS with safesearch and adult content blocking

sdns://AQMAAAAAAAAAGlsyYTAwOjVhNjA6OmJhZDI6MGZmXTo1NDQzIIwhF6nrwVfW-2QFbwrbwRxdg2c0c8RuJY2bL1fU7jUfITIuZG5zY3J5cHQuZmFtaWx5Lm5zMi5hZGd1YXJkLmNvbQ

## adguard-dns-ipv6

Remove ads and protect your computer from malware

sdns://AQMAAAAAAAAAGVsyYTAwOjVhNjA6OmFkMjowZmZdOjU0NDMggdAC02pMpQxHO3R5ZQ_hLgKzIcthOFYqII5APf3FXpQiMi5kbnNjcnlwdC5kZWZhdWx0Lm5zMi5hZGd1YXJkLmNvbQ

## ads-securedns-doh

Filter Ads and no logging (DoH protocol)

sdns://AgMAAAAAAAAADjE0Ni4xODUuMTY3LjQzID4aGg9sU_PpekktVwhLW5gHBZ7gV6sVBYdv2D_aPbg4FGFkcy1kb2guc2VjdXJlZG5zLmV1Ci9kbnMtcXVlcnk

## ads-securedns-ipv6-doh

Filter Ads and no logging (IPv6, DoH protocol)

sdns://AgMAAAAAAAAAHFsyYTAzOmIwYzA6MDoxMDEwOjplOWE6MzAwMV0gPhoaD2xT8-l6SS1XCEtbmAcFnuBXqxUFh2_YP9o9uDgUYWRzLWRvaC5zZWN1cmVkbnMuZXUKL2Rucy1xdWVyeQ

## doh.appliedprivacy.net

Public DoH resolver operated by the Foundation for Applied Privacy (https://appliedprivacy.net).
Hosted in Vienna, Austria.

Non-logging, non-filtering, supports DNSSEC.

sdns://AgcAAAAAAAAAAKA-GhoPbFPz6XpJLVcIS1uYBwWe4FerFQWHb9g_2j24OKBoo-sB-l8CxNNfOhHQBMrwiyJL7qfXnFiMfxPIYTNgLqDvR4Wu5wydV1_nM4MG2T6nlhHl_tzvU2LdZsmLYLstvSAcVDa2UaK1QVwWz9ltGpcJ_ZyPJ-73XPlz2YL_5o5Y8BZkb2guYXBwbGllZHByaXZhY3kubmV0Bi9xdWVyeQ

## arvind-io

Public resolver by EnKrypt (https://arvind.io).
Hosted in Bangalore, India.

Non-logging, non-filtering, supports DNSSEC.

sdns://AQcAAAAAAAAAEjEzOS41OS4xNi4xMzA6NTM1MyCORifHOIOoUQMIIbpa5-XQQfSq75W3gpAWy2Udh8MoyRkyLmRuc2NyeXB0LWNlcnQuYXJ2aW5kLmlv

## bottlepost-dns-nl

Provided by bottlepost.me
Hosted in The Netherlands, DNSSEC / No Logs / No Filter

sdns://AQcAAAAAAAAAEzE3OC4xMjguMjU1LjI4OjUzNTMgkr1k-Lp2d9IXiFlXoBAgFGZUCJSPW_x81Ec6ShkPsJYdMi5kbnNjcnlwdC1jZXJ0LmJvdHRsZXBvc3QubWU

## brasil.dnscrypt-tupi.org-doh

DNSSEC validation, caching, no IPv6, non-logging, non-filtering, uncensored DNS server in Brazil. DNSCrypt, DoH, DoT protocols. More info https://dnscrypt-tupi.org/

sdns://AgcAAAAAAAAADjE5MS4yNTIuMTAwLjM1ABVkbnMuZG5zY3J5cHQtdHVwaS5vcmcKL2Rucy1xdWVyeQ

## captnemo-in

Server running out of a Digital Ocean droplet in BLR1 region.
Maintained by Abhay Rana aka Nemo.

If you are within India, this might be a nice DNS server to use.

sdns://AQQAAAAAAAAAEjEzOS41OS40OC4yMjI6NDQzNCAFOt_yxaMpFtga2IpneSwwK6rV0oAyleham9IvhoceEBsyLmRuc2NyeXB0LWNlcnQuY2FwdG5lbW8uaW4

## charis

Public DNSCrypt server in Amsterdam. DNSSEC, no logs, no filter.
Maintained by @lucenera

sdns://AQcAAAAAAAAAETUxLjE1LjEwNi4xNzY6NDQzIHoBq-npX9jLHUm1YOumKcmBtXIPQSxVHB8ngzhYAgIRHjIuZG5zY3J5cHQtY2VydC5hbXMuY2hhcmlzLmNvbQ

## cisco

Remove your DNS blind spot

Warning: modifies your queries to include a copy of your network
address when forwarding them to a selection of companies and organizations.

sdns://AQAAAAAAAAAADjIwOC42Ny4yMjAuMjIwILc1EUAgbyJdPivYItf9aR6hwzzI1maNDL4Ev6vKQ_t5GzIuZG5zY3J5cHQtY2VydC5vcGVuZG5zLmNvbQ

## cisco-familyshield

Block websites not suitable for children

Warning: modifies your queries to include a copy of your network
address when forwarding them to a selection of companies and organizations.

sdns://AQAAAAAAAAAADjIwOC42Ny4yMjAuMTIzILc1EUAgbyJdPivYItf9aR6hwzzI1maNDL4Ev6vKQ_t5GzIuZG5zY3J5cHQtY2VydC5vcGVuZG5zLmNvbQ

## cisco-ipv6

Cisco OpenDNS IPv6 sandbox

Warning: modifies your queries to include a copy of your network
address when forwarding them to a selection of companies and organizations.

sdns://AQAAAAAAAAAAD1syNjIwOjA6Y2NjOjoyXSC3NRFAIG8iXT4r2CLX_WkeocM8yNZmjQy-BL-rykP7eRsyLmRuc2NyeXB0LWNlcnQub3BlbmRucy5jb20

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

## doh-cleanbrowsing-security

Block access to phishing, malware and malicious domains. It does not block adult content.
By https://cleanbrowsing.org/

sdns://AQMAAAAAAAAAEjE4NS4yMjguMTY4Ljk6ODQ0MyC8rDL61UNpFx8IMtYHUCfDIIzu8Ojpn5QY3HdgZdSPKRFjbGVhbmJyb3dzaW5nLm9yZw

## cloudflare

Cloudflare DNS (anycast) - aka 1.1.1.1 / 1.0.0.1

sdns://AgcAAAAAAAAABzEuMC4wLjGgENk8mGSlIfMGXMOlIlCcKvq7AVgcrZxtjon911-ep0cg63Ul-I8NlFj4GplQGb_TTLiczclX57DvMV8Q-JdjgRgSZG5zLmNsb3VkZmxhcmUuY29tCi9kbnMtcXVlcnk

## cloudflare-ipv6

Cloudflare DNS over IPv6 (anycast)

sdns://AgcAAAAAAAAAGVsyNjA2OjQ3MDA6NDcwMDo6MTExMV06NTOgENk8mGSlIfMGXMOlIlCcKvq7AVgcrZxtjon911-ep0cg63Ul-I8NlFj4GplQGb_TTLiczclX57DvMV8Q-JdjgRgSZG5zLmNsb3VkZmxhcmUuY29tCi9kbnMtcXVlcnk

## commons-host

DoH server by the Commons Host CDN

sdns://AgUAAAAAAAAAACA-GhoPbFPz6XpJLVcIS1uYBwWe4FerFQWHb9g_2j24OAxjb21tb25zLmhvc3QKL2Rucy1xdWVyeQ

## comodo-02

Comodo Dome Shield (anycast) - https://cdome.comodo.com/shield/

sdns://AQAAAAAAAAAACjguMjAuMjQ3LjIg0sJUqpYcHsoXmZb1X7yAHwg2xyN5q1J-zaiGG-Dgs7AoMi5kbnNjcnlwdC1jZXJ0LnNoaWVsZC0yLmRuc2J5Y29tb2RvLmNvbQ

## cpunks-ru

Cypherpunks.ru public DNS server

sdns://AQYAAAAAAAAAEjc3LjUxLjE4MS4yMDk6NTM1MyAYOMyj2VMKZjQzXVAFvTdYROOXfuhoK2xVKBK9p40umR4yLmRuc2NyeXB0LWNlcnQuY3lwaGVycHVua3MucnU

## cs-ch

Switzerland DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAACzgxLjE3LjMxLjM0IDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw

## cs-swe

Sweden DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADzEyOC4xMjcuMTA0LjEwOCAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM

## cs-nl

Netherlands DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADjIxMy4xNjMuNjQuMjA4IDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw

## cs-nl2

Secondary Netherlands DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADTE4NS4xMDcuODAuODQgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz

## cs-fi

Finland DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADjE4NS4xMTcuMTE4LjIwIDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw

## cs-pl

Poland DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAACzUuMTMzLjguMTg3IDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw

## cs-dk

Denmark DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADzE4NS4yMTIuMTY5LjEzOSAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM

## cs-it

Italy DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADjE4NS45NC4xOTMuMjM0IDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw

## cs-fr

France DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADTIxMi4xMjkuNDYuMzIgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz

## cs-fr2

Secondary France DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADTE5NS4xNTQuNDAuNDggMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz

## cs-pt

Portugal DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADTEwOS43MS40Mi4yMjggMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz

## cs-hk

Hong Kong DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADDEwMy4xNi4yNy41MyAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM

## cs-ro

Romania DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADDUuMjU0Ljk2LjE5NSAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM

## cs-mo

Moldova DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADzE3OC4xNzUuMTM5LjIxMSAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM

## cs-lv

Latvia DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADzEwOS4yNDguMTQ5LjEzMyAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM

## cs-uk

England DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADTgyLjE2My43Mi4xMjMgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz

## cs-de

Germany DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADDg0LjE2LjI0MC40MyAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM

## cs-de2

Secondary Germany DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADjg5LjE2My4yMTQuMTc0IDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw

## cs-ca

Canada DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADzE2Mi4yMjEuMjA3LjIyOCAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM

## cs-ca2

Secondary Canada DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADjE2Ny4xMTQuODQuMTMyIDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw

## cs-usny

US - NY DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADzE3My4yMzQuMTU5LjIzNSAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM

## cs-usil

US - IL DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADjE3My4yMzQuNTYuMTE1IDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw

## cs-usnv

US - NV DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADzEwNC4yMzguMTk1LjEzOSAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM

## cs-uswa

US - WA DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADDY0LjEyMC41LjI1MSAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM

## cs-usdc

US - DC DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADDE5OC43LjU4LjIyNyAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM

## cs-ustx

US - TX DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADTIwOS41OC4xNDcuMzYgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz

## cs-usga

US - GA DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADTY0LjQyLjE4MS4yMjcgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz

## cs-usnc

US - NC DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADjE1NS4yNTQuMjkuMTEzIDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw

## cs-usca

US - CA DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADDIzLjE5LjY3LjExNiAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM

## cs-usor

US - OR DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADTEwNC4yNTUuMTc1LjIgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz

## d0wn-is-ns2

Server provided by Martin 'd0wn' Albus

sdns://AQcAAAAAAAAADTkzLjk1LjIyNi4xNjUghGA0qcYwyjwErEqQFiXxeoeyrLlBgKxIHiwQ6M7eGm8cMi5kbnNjcnlwdC1jZXJ0LmlzMi5kMHduLmJpeg

## d0wn-tz-ns1

Server provided by Martin 'd0wn' Albus

sdns://AQcAAAAAAAAACzQxLjc5LjY5LjEzINYGFfvRRTuhTnaKPlxcs6wXRhMxRj2gr4z33wTaTXVtGzIuZG5zY3J5cHQtY2VydC50ei5kMHduLmJpeg

## d0wn-tz-ns1-ipv6

Server provided by Martin 'd0wn' Albus

sdns://AQcAAAAAAAAAGFsyYzBmOmZkYTg6NTo6MmVkMTpkMmVjXSDWBhX70UU7oU52ij5cXLOsF0YTMUY9oK-M998E2k11bRsyLmRuc2NyeXB0LWNlcnQudHouZDB3bi5iaXo

## de.dnsmaschine.net

DNSSEC/Non-logged/Uncensored
Hosted by vultr.com (Frankfurt Germany)

sdns://AQcAAAAAAAAAEzIwOS4yNTAuMjM1LjE3MDo0NDMgz0wbvISl_NVCSe0wDJMS79BAFZoWth1djmhuzv_n3KAiMi5kbnNjcnlwdC1jZXJ0LmRlLmRuc21hc2NoaW5lLm5ldA

## dnscrypt.ca-1

Uncensored DNSSEC validating and log-free

sdns://AQcAAAAAAAAAFDE5OS4xNjcuMTMwLjExODo1MzUzIHT3RVUXvCb3EXflbXKTJ4hscpFbP0YoMD-RDEfDjoJ5HTIuZG5zY3J5cHQtY2VydC5kbnNjcnlwdC5jYS0x

## dnscrypt.ca-1-ipv6

Uncensored DNSSEC validating and log-free

sdns://AQcAAAAAAAAAH1syNjA1OjIxMDA6MDoxOjo3MzRkOjc4NzZdOjUzNTMgie_Aik8Gbx0Yhl3AXGNrjkhIIuR2hdxG8wSccOyE5podMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeXB0LmNhLTE

## dnscrypt.ca-2

Uncensored DNSSEC validating and log-free

sdns://AQcAAAAAAAAAFDE5OS4xNjcuMTI4LjExMjo1MzUzIEPVLIJZIpbC22-NSM4iT9zHJibhBvbjiGGT-gCQKWMbHTIuZG5zY3J5cHQtY2VydC5kbnNjcnlwdC5jYS0y

## dnscrypt.ca-2-ipv6

Uncensored DNSSEC validating and log-free

sdns://AQcAAAAAAAAAH1syNjA1OjIxMDA6MDoxOjpiNWFkOjE4ZTJdOjUzNTMg5DtuKuW1dRp0BBgQ97rtLa9wScW38wTZSLyEgVkXmowdMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeXB0LmNhLTI

## dnscrypt.eu-dk

Free, non-logged, uncensored. Hosted by Netgroup.

sdns://AQcAAAAAAAAADDc3LjY2Ljg0LjIzMyA3SFWF47nQiP0lrTawNwH1UgzWSJ6a3VIUV0lVnwqZVSUyLmRuc2NyeXB0LWNlcnQucmVzb2x2ZXIyLmRuc2NyeXB0LmV1

## dnscrypt.eu-dk-ipv6

Free, non-logged, uncensored. Hosted by Netgroup.

sdns://AQcAAAAAAAAAFFsyMDAxOjE0NDg6MjQzOjpkYzJdIDdIVYXjudCI_SWtNrA3AfVSDNZInprdUhRXSVWfCplVJTIuZG5zY3J5cHQtY2VydC5yZXNvbHZlcjIuZG5zY3J5cHQuZXU

## dnscrypt.eu-nl

Free, non-logged, uncensored. Hosted by RamNode.

sdns://AQcAAAAAAAAADjE3Ni41Ni4yMzcuMTcxIGfADywhxVSBRd18tGonGvLrlpkxQKMJtiuNFlMRhZxmJTIuZG5zY3J5cHQtY2VydC5yZXNvbHZlcjEuZG5zY3J5cHQuZXU

## dnscrypt.me

DNSSEC / no logs / uncensored, Germany
https://dnscrypt.me

sdns://AQcAAAAAAAAADTE2Ny44Ni45MC4xMDMgMl9Yy8VDXd8tCI8yLU4p0i_BmnLn4H77n36t4BqPm8UbMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeXB0Lm1l

## dnscrypt.me-ipv6

DNSSEC / no logs / uncensored, Germany
https://dnscrypt.me

sdns://AQcAAAAAAAAAGFsyYTAyOmMyMDc6MzAwMzo3MzM0OjoxXSAyX1jLxUNd3y0IjzItTinSL8Gacufgfvuffq3gGo-bxRsyLmRuc2NyeXB0LWNlcnQuZG5zY3J5cHQubWU

## dnscrypt.nl-ns0

DNSCrypt v2 server in Amsterdam, the Netherlands. DNSSEC, no logs, uncensored, recursive DNS. https://dnscrypt.nl

sdns://AQcAAAAAAAAADDQ1Ljc2LjM1LjIxMiBMhPuMBRFd-l-Xxe0DKRNwx4q81k4V3VOrCN5y-4RKyh8yLmRuc2NyeXB0LWNlcnQubnMwLmRuc2NyeXB0Lm5s

## dnscrypt.nl-ns0-ipv6

DNSCrypt v2 server in Amsterdam, the Netherlands. DNSSEC, no logs, uncensored, recursive DNS. https://dnscrypt.nl

sdns://AQcAAAAAAAAAJlsyMDAxOjE5ZjA6NTAwMTozMGE6NTQwMDpmZjpmZTU4OjcxNDBdIEyE-4wFEV36X5fF7QMpE3DHirzWThXdU6sI3nL7hErKHzIuZG5zY3J5cHQtY2VydC5uczAuZG5zY3J5cHQubmw

## dnscrypt.nl-ns0-doh

DNS-over-HTTPS server in Amsterdam, the Netherlands. DNSSEC, no logs, uncensored, recursive DNS. https://dnscrypt.nl

sdns://AgcAAAAAAAAADjEwOC42MS4xOTkuMTcwID4aGg9sU_PpekktVwhLW5gHBZ7gV6sVBYdv2D_aPbg4D2RvaC5kbnNjcnlwdC5ubAovZG5zLXF1ZXJ5

## dnscrypt.uk-ipv4

DNSCrypt v2, no logs, uncensored, DNSSEC. Hosted in London UK by Digital Ocean
https://www.dnscrypt.uk

sdns://AQcAAAAAAAAAEjEzOS41OS4yMDAuMTE2OjQ0MyAmJwT-OXZ9NntZ2eu_HtZeXARhCdiAynbBYcu6bArCdxsyLmRuc2NyeXB0LWNlcnQuZG5zY3J5cHQudWs

## dnscrypt.uk-ipv6

DNSCrypt v2, no logs, uncensored, DNSSEC. Hosted in London UK by Digital Ocean
https://www.dnscrypt.uk

sdns://AQcAAAAAAAAAHlsyYTAzOmIwYzA6MTplMDo6MmUzOmUwMDFdOjQ0MyAmJwT-OXZ9NntZ2eu_HtZeXARhCdiAynbBYcu6bArCdxsyLmRuc2NyeXB0LWNlcnQuZG5zY3J5cHQudWs

## dnscrypt-jp-blahdns-ipv4

Blocks ad and Tracking, no Logging, DNSSEC, Hosted in Japan. By https://blahdns.com/

sdns://AQMAAAAAAAAAEzEwOC42MS4yMDEuMTE5Ojg0NDMgyJjbSS4IgTY_2KH3NVGG0DNIgBPzLEqf8r00nAbcUxQbMi5kbnNjcnlwdC1jZXJ0LmJsYWhkbnMuY29t

## dnscrypt-jp-blahdns-ipv6

Blocks ad and Tracking, no Logging, DNSSEC, Hosted in Japan. By https://blahdns.com/

sdns://AQMAAAAAAAAALlsyMDAxOjE5ZjA6NzAwMToxZGVkOjU0MDA6MDFmZjpmZTkwOjk0NWJdOjg0NDMgyJjbSS4IgTY_2KH3NVGG0DNIgBPzLEqf8r00nAbcUxQbMi5kbnNjcnlwdC1jZXJ0LmJsYWhkbnMuY29t

## dnscrypt-de-blahdns-ipv4

Blocks ad and Tracking, no Logging, DNSSEC, Hosted in Germany. By https://blahdns.com/

sdns://AQMAAAAAAAAAEzE1OS42OS4xOTguMTAxOjg0NDMgyJjbSS4IgTY_2KH3NVGG0DNIgBPzLEqf8r00nAbcUxQbMi5kbnNjcnlwdC1jZXJ0LmJsYWhkbnMuY29t

## dnscrypt-de-blahdns-ipv6

Blocks ad and Tracking, no Logging, DNSSEC, Hosted in Germany. By https://blahdns.com/

sdns://AQMAAAAAAAAAHFsyYTAxOjRmODoxYzFjOjZiNGI6OjFdOjg0NDMgyJjbSS4IgTY_2KH3NVGG0DNIgBPzLEqf8r00nAbcUxQbMi5kbnNjcnlwdC1jZXJ0LmJsYWhkbnMuY29t

## dnswarden-adult-filter-ipv4

Blocks adult content, ads, trackers, phishing and malware sites. Enforces forcesafesearch on widely used search engines and youtube. DNSSEC enabled and no query logging. Hosted in Germany. By https://dnswarden.com

sdns://AgMAAAAAAAAADDE1OS42OS4xNi41OAASZG9oMi5kbnN3YXJkZW4uY29tAS8

## dnswarden-adult-filter-ipv6

Blocks adult content, ads, trackers, phishing and malware sites. Enforces forcesafesearch on widely used search engines and youtube. DNSSEC enabled and no query logging. Hosted in Germany. By https://dnswarden.com

sdns://AgMAAAAAAAAAF1syYTAxOjRmODoxYzFjOjc1NDk6OjFdABJkb2gyLmRuc3dhcmRlbi5jb20BLw

## dnswarden-normal-ipv4

Blocks ads, trackers, phishing and malware sites. DNSSEC enabled and no query logging. Hosted in Germany. By https://dnswarden.com

sdns://AgMAAAAAAAAADTk0LjEzMC4xODMuMTgAEmRvaDEuZG5zd2FyZGVuLmNvbQEv

## dnswarden-normal-ipv6

Blocks ads, trackers, phishing and malware sites. DNSSEC enabled and no query logging. Hosted in Germany. By https://dnswarden.com

sdns://AgMAAAAAAAAAF1syYTAxOjRmODoxYzBjOjQyYjI6OjFdABJkb2gxLmRuc3dhcmRlbi5jb20BLw

## doh-de-blahdns

Blocks ad and Tracking, no Logging, DNSSEC, Hosted in Germany. By https://blahdns.com/

sdns://AgMAAAAAAAAADjE1OS42OS4xOTguMTAxABJkb2gtZGUuYmxhaGRucy5jb20KL2Rucy1xdWVyeQ

## doh-de-blahdns-v6

Blocks ad and Tracking, no Logging, DNSSEC, Hosted in Germany. By https://blahdns.com/

sdns://AgMAAAAAAAAAF1syYTAxOjRmODoxYzFjOjZiNGI6OjFdABJkb2gtZGUuYmxhaGRucy5jb20KL2Rucy1xdWVyeQ

## doh-jp-blahdns

Blocks ad and Tracking, no Logging, DNSSEC, Hosted in Japan. By https://blahdns.com/

sdns://AgMAAAAAAAAADjEwOC42MS4yMDEuMTE5ABJkb2gtanAuYmxhaGRucy5jb20KL2Rucy1xdWVyeQ

## doh-jp-blahdns-v6

Blocks ad and Tracking, no Logging, DNSSEC, Hosted in Japan. By https://blahdns.com/

sdns://AgMAAAAAAAAAKVsyMDAxOjE5ZjA6NzAwMToxZGVkOjU0MDA6MDFmZjpmZTkwOjk0NWJdABJkb2gtanAuYmxhaGRucy5jb20KL2Rucy1xdWVyeQ

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

## doh-cleanbrowsing-security

Block access to phishing, malware and malicious domains. It does not block adult content.
By https://cleanbrowsing.org/

sdns://AgMAAAAAAAAAAAAVZG9oLmNsZWFuYnJvd3Npbmcub3JnFS9kb2gvc2VjdXJpdHktZmlsdGVyLw

## doh-crypto-sx

DNS-over-HTTPS server. Anycast, no logs, no censorship, DNSSEC.
Backend hosted by Scaleway, globally cached via Cloudflare.
Maintained by Frank Denis.

sdns://AgcAAAAAAAAADDEwNC4yOC4wLjEwNiAd2FCKjFZZBDl8eGRR4I9XYTzzyKcj9vN5_Uw4WLbznw1kb2guY3J5cHRvLnN4Ci9kbnMtcXVlcnk

## doh-crypto-sx-ipv6

DNS-over-HTTPS server accessible over IPv6. Anycast, no logs, no censorship, DNSSEC.
Backend hosted by Scaleway, globally cached via Cloudflare.
Maintained by Frank Denis.

sdns://AgcAAAAAAAAAF1syNjA2OjQ3MDA6MzA6OjY4MWM6NmFdIB3YUIqMVlkEOXx4ZFHgj1dhPPPIpyP283n9TDhYtvOfEmRvaC1pcHY2LmNyeXB0by5zeAovZG5zLXF1ZXJ5

## doh-ibksturm

doh-server (nginx - doh-httpproxy - unbound backend), DNSSEC / Non-Logged / Uncensored, OpenNIC and Root DNS-Zone Copy
Hosted in Switzerland by ibksturm, aka Andreas Ziegler

sdns://AgcAAAAAAAAAACA-GhoPbFPz6XpJLVcIS1uYBwWe4FerFQWHb9g_2j24OBRpYmtzdHVybS5zeW5vbG9neS5tZQovZG5zLXF1ZXJ5

## ev-va

Non-logging, uncensored DNS resolver provided by evilvibes.com
Location: Vancouver, Canada

sdns://AQcAAAAAAAAADTIzLjExMS43NC4yMTYg3_NERwhF2C4tPlnR0CSeIAmRC3wwXtucNhKMVoW_prQjMi5kbnNjcnlwdC1jZXJ0LmV2LXZhLmV2aWx2aWJlcy5jb20

## ev-to

Non-logging, uncensored DNS resolver provided by evilvibes.com
Location: Toronto, Canada

sdns://AQcAAAAAAAAADTIzLjExMS42OS4xMjYgTvUHn_uZRDf0m10w-HJPpAVY7_UcoDKNuXpXcWJSVwsjMi5kbnNjcnlwdC1jZXJ0LmV2LXRvLmV2aWx2aWJlcy5jb20

## freetsa.org

Non-logged/Uncensored provided by freetsa.org

sdns://AQcAAAAAAAAAEzIwNS4xODUuMTE2LjExNjo1NTMg2P-7QuAxvnp5cwtFVo1Jak6Ky1mqg2b9arkeJyp9FuQbMi5kbnNjcnlwdC1jZXJ0LmZyZWV0c2Eub3Jn

## geekdns-doh-east

GeekDNS in eastern China (Shanghai).

GeekDNS is a non-logging public DNS service located in mainland china,
that also blocks ads and trackers. Queries are cached locally, and,
for some domains, resolved by servers located in Taiwan.

https://dns.233py.com/

sdns://AgMAAAAAAAAADTQ3LjEwMS4xMzYuMzcgPhoaD2xT8-l6SS1XCEtbmAcFnuBXqxUFh2_YP9o9uDgOZWRucy4yMzNweS5jb20KL2Rucy1xdWVyeQ

## geekdns-doh-west

GeekDNS in western China (Chongqing).

GeekDNS is a non-logging public DNS service located in mainland china,
that also blocks ads and trackers. Queries are cached locally, and,
for some domains, resolved by servers located in Taiwan.

https://dns.233py.com/

sdns://AgMAAAAAAAAADjExOC4yNC4yMDguMTk3ID4aGg9sU_PpekktVwhLW5gHBZ7gV6sVBYdv2D_aPbg4DndkbnMuMjMzcHkuY29tCi9kbnMtcXVlcnk

## geekdns-doh-south

GeekDNS in southern China (Guangzhou).

GeekDNS is a non-logging public DNS service located in mainland china,
that also blocks ads and trackers. Queries are cached locally, and,
for some domains, resolved by servers located in Taiwan.

https://dns.233py.com/

sdns://AgMAAAAAAAAADTExOS4yOS4xMDcuODUgPhoaD2xT8-l6SS1XCEtbmAcFnuBXqxUFh2_YP9o9uDgOc2Rucy4yMzNweS5jb20KL2Rucy1xdWVyeQ

## geekdns-doh-north

GeekDNS in northern China (Beijing).

GeekDNS is a non-logging public DNS service located in mainland china,
that also blocks ads and trackers. Queries are cached locally, and,
for some domains, resolved by servers located in Taiwan.

https://dns.233py.com/

sdns://AgMAAAAAAAAADzExNC4xMTUuMjQwLjE3NSA-GhoPbFPz6XpJLVcIS1uYBwWe4FerFQWHb9g_2j24OA5uZG5zLjIzM3B5LmNvbQovZG5zLXF1ZXJ5

## google

Google DNS (anycast)

sdns://AgUAAAAAAAAAACAe9iTP_15r07rd8_3b_epWVGfjdymdx-5mdRZvMAzBuQ5kbnMuZ29vZ2xlLmNvbQ0vZXhwZXJpbWVudGFs

## gridns-jp

Gridth's public filtering non-logging DNS-over-HTTPS server. Block ads and tracking.
Hosted at Linode in Tokyo. Upstream to 1.1.1.1. No EDNS Client Subnet.

sdns://AgcAAAAAAAAADjE3Mi4xMDUuMjQxLjkzAA1qcC5ncmlkbnMueHl6Ci9kbnMtcXVlcnk

## gridns-jp-ipv6

Gridth's public filtering non-logging DNS-over-HTTPS server. Block ads and tracking.
Hosted at Linode in Tokyo. Upstream to 1.1.1.1. No EDNS Client Subnet. IPv6-enabled.

sdns://AgcAAAAAAAAAIFsyNDAwOjg5MDI6OmYwM2M6OTFmZjpmZWVkOjIyMGJdAA1qcC5ncmlkbnMueHl6Ci9kbnMtcXVlcnk

## gridns-sg

Gridth's public filtering non-logging DNS-over-HTTPS server. Block ads and tracking.
Hosted at Linode in Singapore. Upstream to 1.1.1.1. No EDNS Client Subnet.

sdns://AgcAAAAAAAAADTEzOS4xNjIuMy4xMjMADXNnLmdyaWRucy54eXoKL2Rucy1xdWVyeQ

## gridns-sg-ipv6

Gridth's public filtering non-logging DNS-over-HTTPS server. Block ads and tracking.
Hosted at Linode in Singapore. Upstream to 1.1.1.1. No EDNS Client Subnet. IPv6-enabled.

sdns://AgcAAAAAAAAAIFsyNDAwOjg5MDE6OmYwM2M6OTFmZjpmZWVkOjhkNDddAA1zZy5ncmlkbnMueHl6Ci9kbnMtcXVlcnk

## ibksturm

dnscrypt-server (nginx - dnscrypt-wrapper - unbound backend), DNSSEC / Non-Logged / Uncensored, OpenNIC and Root DNS-Zone Copy
Hosted in Switzerland by ibksturm, aka Andreas Ziegler

sdns://AQcAAAAAAAAADzIxNy4xNjIuMjA2LjIyMCAbkeo7E6QrYUfz_2_40sQSRfXIf4wu2U1aexB1dmIC2hgyLmRuc2NyeXB0LWNlcnQuaWJrc3R1cm0

## id-gmail

Non-Logging DNSCrypt server located in Singapore.
Filters out ads, trackers and malware, supports DNSSEC, provided by id-gmail.

sdns://AQMAAAAAAAAADTE0OS4yOC4xNTIuODEg75aAZujZlPBl2D7d0xru7fVthldGPkrKR83X_pfD1PYcMi5kbnNjcnlwdC1jZXJ0LmRucy50aWFyLmFwcA

## id-gmail-ipv6

Non-Logging DNSCrypt (IPv6) server located in Singapore.
Filters out ads, trackers and malware, supports DNSSEC, provided by id-gmail.

sdns://AQMAAAAAAAAAKFsyNDAxOmMwODA6MTQwMDo0ZmU2OjU0MDA6MWZmOmZlZDA6MmFlM10g75aAZujZlPBl2D7d0xru7fVthldGPkrKR83X_pfD1PYcMi5kbnNjcnlwdC1jZXJ0LmRucy50aWFyLmFwcA

## id-gmail-doh

Non-Logging DNS-over-HTTPS server located in Singapore.
Filters out ads, trackers and malware, supports DNSSEC, provided by id-gmail.

sdns://AgMAAAAAAAAACzQ1LjMyLjEwNS40ID4aGg9sU_PpekktVwhLW5gHBZ7gV6sVBYdv2D_aPbg4DGRvaC50aWFyLmFwcAovZG5zLXF1ZXJ5

## id-gmail-doh-ipv6

Non-Logging DNS-over-HTTPS (IPv6) server located in Singapore.
Filters out ads, trackers and malware, supports DNSSEC, provided by id-gmail.

sdns://AgMAAAAAAAAAKFsyMDAxOjE5ZjA6NDQwMDo3YmNjOjU0MDA6MWZmOmZlZDE6ODU5OV0gPhoaD2xT8-l6SS1XCEtbmAcFnuBXqxUFh2_YP9o9uDgMZG9oLnRpYXIuYXBwCi9kbnMtcXVlcnk

## ipredator

Public DNSCrypt server in Sweden provided by Ipredator.se

sdns://AQcAAAAAAAAADTE5NC4xMzIuMzIuMzIgxExWaqjWRsQysQT1PQCWGzLccc8cBL2esBPkgOekeCgcMi5kbnNjcnlwdC1jZXJ0LmlwcmVkYXRvci5zZQ

## nawala-childprotection

Internet filtering system (anycast), protecting child from inappropriate websites and abusive contents.
By http://nawala.id in Indonesia.

sdns://AQAAAAAAAAAADzE4MC4xMzEuMTQ0LjE0NCDGC-b_38Dj4-ikI477AO1GXcLPfETOFpE36KZIHdOzLhkyLmRuc2NyeXB0LWNlcnQubmF3YWxhLmlk

## opennic-ethservices

OpenNIC • DNSSEC • 24-hour Logs • AnonymousLogs • NoFilters
Location: Frankfurt, Germany
By ethservices.

sdns://AQcAAAAAAAAAEzE5NS4xMC4xOTUuMTk1OjUzNTMg8hbE05QkH0WdwNiGcxtcLvFewNj3USVp1A-VL0P77HIoMi5kbnNjcnlwdC1jZXJ0Lm9wZW5uaWMyLmV0aC1zZXJ2aWNlcy5kZQ

## opennic-ethservices2

OpenNIC • DNSSEC • 24-hour Logs • AnonymousLogs • NoFilters
Location: Frankfurt, Germany
By ethservices.

sdns://AQYAAAAAAAAAEzE5NS4xMC4xOTUuMTk1OjUzNTMg8hbE05QkH0WdwNiGcxtcLvFewNj3USVp1A-VL0P77HIoMi5kbnNjcnlwdC1jZXJ0Lm9wZW5uaWMyLmV0aC1zZXJ2aWNlcy5kZQ

## opennic-luggs

Public DNS server in Canada operated by Luggs

sdns://AQYAAAAAAAAADTE0Mi40LjIwNC4xMTEgHBl5MxvoI8zPCJp5BpN-XDQQKlasf2Jw4EYlsu3bBOMfMi5kbnNjcnlwdC1jZXJ0Lm5zMy5jYS5sdWdncy5jbw

## opennic-luggs-ipv6

Public DNS server in Canada operated by Luggs

sdns://AQYAAAAAAAAAIVsyNjA3OjUzMDA6MTIwOmE4YToxNDI6NDoyMDQ6MTExXSAcGXkzG-gjzM8ImnkGk35cNBAqVqx_YnDgRiWy7dsE4x8yLmRuc2NyeXB0LWNlcnQubnMzLmNhLmx1Z2dzLmNv

## opennic-luggs2

Second public DNS server in Canada operated by Luggs

sdns://AQYAAAAAAAAAEDE0Mi40LjIwNS40Nzo0NDMgvL-34FDBPaJCLACwsaya1kjFwmS8thcLiD1xishuugkfMi5kbnNjcnlwdC1jZXJ0Lm5zNC5jYS5sdWdncy5jbw

## opennic-luggs2-ipv6

Second public DNS server in Canada operated by Luggs (IPv6)

sdns://AQYAAAAAAAAAJFsyNjA3OjUzMDA6MTIwOmE4YToxNDI6NDoyMDU6NDddOjQ0MyC8v7fgUME9okIsALCxrJrWSMXCZLy2FwuIPXGKyG66CR8yLmRuc2NyeXB0LWNlcnQubnM0LmNhLmx1Z2dzLmNv

## powerdns-doh

By PowerDNS/Open-Xchange https://powerdns.org

sdns://AgcAAAAAAAAAACA-GhoPbFPz6XpJLVcIS1uYBwWe4FerFQWHb9g_2j24OBBkb2gucG93ZXJkbnMub3JnAS8

## publicarray-au

DNSSEC • OpenNic • Non-logging • Uncensored - hosted at vultr.com
Maintained by publicarray - https://dns.seby.io

sdns://AQcAAAAAAAAADDQ1Ljc2LjExMy4zMSAIVGh4i6eKXqlF6o9Fg92cgD2WcDvKQJ7v_Wq4XrQsVhsyLmRuc2NyeXB0LWNlcnQuZG5zLnNlYnkuaW8

## publicarray-au-doh

DNSSEC • OpenNic • Non-logging • Uncensored - hosted on vultr.com
Maintained by publicarray - https://dns.seby.io

sdns://AgcAAAAAAAAADDQ1Ljc2LjExMy4zMSA-GhoPbFPz6XpJLVcIS1uYBwWe4FerFQWHb9g_2j24OBBkb2guc2VieS5pbzo4NDQzCi9kbnMtcXVlcnk

## qag.me

Plain Vanilla setup of dnscrypt-server-docker on a headless box. Upstream to 1.1.1.1.
Home Server running on a static IP in Bangalore / Bengaluru, INDIA.
Maintained by Cruisemaniac (https://cruisemaniac.com) aka Ashwin Murali.

sdns://AQcAAAAAAAAAEjEwNi41MS4xMjguNzg6NDQzNCDrpsCqF14emkVAo_yJi9T2xxp5KmXhlGtbTL1R-5vVLhYyLmRuc2NyeXB0LWNlcnQucWFnLm1l

## quad9-dnscrypt-ip4-filter-pri
Quad9 (anycast) dnssec/no-log/filter 9.9.9.9
sdns://AQMAAAAAAAAADDkuOS45Ljk6ODQ0MyBnyEe4yHWM0SAkVUO-dWdG3zTfHYTAC4xHA2jfgh2GPhkyLmRuc2NyeXB0LWNlcnQucXVhZDkubmV0

## quad9-dnscrypt-ip4-filter-alt
Quad9 (anycast) dnssec/no-log/filter 149.112.112.9
sdns://AQMAAAAAAAAAEjE0OS4xMTIuMTEyLjk6ODQ0MyBnyEe4yHWM0SAkVUO-dWdG3zTfHYTAC4xHA2jfgh2GPhkyLmRuc2NyeXB0LWNlcnQucXVhZDkubmV0

## quad9-dnscrypt-ip4-nofilter-pri
Quad9 (anycast) no-dnssec/no-log/no-filter 9.9.9.10
sdns://AQYAAAAAAAAADTkuOS45LjEwOjg0NDMgZ8hHuMh1jNEgJFVDvnVnRt803x2EwAuMRwNo34Idhj4ZMi5kbnNjcnlwdC1jZXJ0LnF1YWQ5Lm5ldA

## quad9-dnscrypt-ip4-nofilter-alt
Quad9 (anycast) no-dnssec/no-log/no-filter 149.112.112.10
sdns://AQYAAAAAAAAAEzE0OS4xMTIuMTEyLjEwOjg0NDMgZ8hHuMh1jNEgJFVDvnVnRt803x2EwAuMRwNo34Idhj4ZMi5kbnNjcnlwdC1jZXJ0LnF1YWQ5Lm5ldA

## quad9-dnscrypt-ip6-filter-alt
Quad9 (anycast) dnssec/no-log/filter 2620:fe::9
sdns://AQMAAAAAAAAAEVsyNjIwOmZlOjo5XTo4NDQzIGfIR7jIdYzRICRVQ751Z0bfNN8dhMALjEcDaN-CHYY-GTIuZG5zY3J5cHQtY2VydC5xdWFkOS5uZXQ

## quad9-dnscrypt-ip6-filter-pri
Quad9 (anycast) dnssec/no-log/filter 2620:fe::fe:9
sdns://AQMAAAAAAAAAFFsyNjIwOmZlOjpmZTo5XTo4NDQzIGfIR7jIdYzRICRVQ751Z0bfNN8dhMALjEcDaN-CHYY-GTIuZG5zY3J5cHQtY2VydC5xdWFkOS5uZXQ

## quad9-dnscrypt-ip6-nofilter-pri
Quad9 (anycast) no-dnssec/no-log/no-filter 2620:fe::10
sdns://AQYAAAAAAAAAElsyNjIwOmZlOjoxMF06ODQ0MyBnyEe4yHWM0SAkVUO-dWdG3zTfHYTAC4xHA2jfgh2GPhkyLmRuc2NyeXB0LWNlcnQucXVhZDkubmV0

## quad9-dnscrypt-ip6-nofilter-alt
Quad9 (anycast) no-dnssec/no-log/no-filter 2620:fe::fe:10
sdns://AQYAAAAAAAAAFVsyNjIwOmZlOjpmZToxMF06ODQ0MyBnyEe4yHWM0SAkVUO-dWdG3zTfHYTAC4xHA2jfgh2GPhkyLmRuc2NyeXB0LWNlcnQucXVhZDkubmV0

## quad9-doh-ip4-filter-pri
Quad9 (anycast) dnssec/no-log/filter 9.9.9.9
sdns://AgMAAAAAAAAABzkuOS45LjmAABJkbnM5LnF1YWQ5Lm5ldDo0NDMKL2Rucy1xdWVyeQ

## quad9-doh-ip4-filter-alt
Quad9 (anycast) dnssec/no-log/filter 149.112.112.9
sdns://AgMAAAAAAAAADTE0OS4xMTIuMTEyLjmAABJkbnM5LnF1YWQ5Lm5ldDo0NDMKL2Rucy1xdWVyeQ

## quad9-doh-ip4-nofilter-pri
Quad9 (anycast) no-dnssec/no-log/no-filter 9.9.9.10
sdns://AgYAAAAAAAAACDkuOS45LjEwgAASZG5zOS5xdWFkOS5uZXQ6NDQzCi9kbnMtcXVlcnk

## quad9-doh-ip4-nofilter-alt
Quad9 (anycast) no-dnssec/no-log/no-filter 149.112.112.10
sdns://AgYAAAAAAAAADjE0OS4xMTIuMTEyLjEwgAASZG5zOS5xdWFkOS5uZXQ6NDQzCi9kbnMtcXVlcnk

## quad9-doh-ip6-filter-pri
Quad9 (anycast) dnssec/no-log/filter 2620:fe::9
sdns://AgMAAAAAAAAADFsyNjIwOmZlOjo5XYAAEmRuczkucXVhZDkubmV0OjQ0MwovZG5zLXF1ZXJ5

## quad9-doh-ip6-filter-alt
Quad9 (anycast) dnssec/no-log/filter 2620:fe::fe:9
sdns://AgMAAAAAAAAAD1syNjIwOmZlOjpmZTo5XYAAEmRuczkucXVhZDkubmV0OjQ0MwovZG5zLXF1ZXJ5

## quad9-doh-ip6-nofilter-pri
Quad9 (anycast) no-dnssec/no-log/no-filter 2620:fe::10
sdns://AgYAAAAAAAAADVsyNjIwOmZlOjoxMF2AABJkbnM5LnF1YWQ5Lm5ldDo0NDMKL2Rucy1xdWVyeQ

## quad9-doh-ip6-nofilter-alt
Quad9 (anycast) no-dnssec/no-log/no-filter 2620:fe::fe:10
sdns://AgYAAAAAAAAAEFsyNjIwOmZlOjpmZToxMF2AABJkbnM5LnF1YWQ5Lm5ldDo0NDMKL2Rucy1xdWVyeQ

## qualityology.com

Non-logging, non-filtering, DNSSEC validating server in Los Angeles, California.
Maintained by Evan Xu (@ex-git)

sdns://AQcAAAAAAAAAEjE3My44Mi4yMzIuMjMyOjg1MyCPlK_22Cu9WRVyKgl-CZp2GXezsRDWizG-BHIzChok4iAyLmRuc2NyeXB0LWNlcnQucXVhbGl0eW9sb2d5LmNvbQ

## rubyfish-ea

Resolver in mainland China, forwarding queries for non-Chinese domains
to upstream servers in East Asia.

https://www.rubyfish.cn/

sdns://AgUAAAAAAAAAACA-GhoPbFPz6XpJLVcIS1uYBwWe4FerFQWHb9g_2j24OBJlYS1kbnMucnVieWZpc2guY24KL2Rucy1xdWVyeQ

## rubyfish-uw

Resolver in mainland China, forwarding queries for non-Chinese domains
to US-West.

https://www.rubyfish.cn/

sdns://AgUAAAAAAAAAACA-GhoPbFPz6XpJLVcIS1uYBwWe4FerFQWHb9g_2j24OBJ1dy1kbnMucnVieWZpc2guY24KL2Rucy1xdWVyeQ

## scaleway-fr

DNSSEC/Non-logged/Uncensored - ARM server donated by Scaleway.com
Maintained by Frank Denis - https://fr.dnscrypt.info
This server used to be called `dnscrypt.org-fr`.

sdns://AQcAAAAAAAAADjIxMi40Ny4yMjguMTM2IOgBuE6mBr-wusDOQ0RbsV66ZLAvo8SqMa4QY2oHkDJNHzIuZG5zY3J5cHQtY2VydC5mci5kbnNjcnlwdC5vcmc

## securedns

Uncensored and no logging (DNSCrypt protocol)

sdns://AQcAAAAAAAAAEzE0Ni4xODUuMTY3LjQzOjUzNTMg9J8sc01itoYxntB-aRlDOy8ThfQe-8ovF21ZCy5FPoYcMi5kbnNjcnlwdC1jZXJ0LnNlY3VyZWRucy5ldQ

## securedns-ipv6

Uncensored and no logging (IPv6, DNSCrypt protocol)

sdns://AQcAAAAAAAAAIVsyYTAzOmIwYzA6MDoxMDEwOjplOWE6MzAwMV06NTM1MyD0nyxzTWK2hjGe0H5pGUM7LxOF9B77yi8XbVkLLkU-hhwyLmRuc2NyeXB0LWNlcnQuc2VjdXJlZG5zLmV1

## securedns-doh

Uncensored and no logging (DoH protocol)

sdns://AgcAAAAAAAAADjE0Ni4xODUuMTY3LjQzID4aGg9sU_PpekktVwhLW5gHBZ7gV6sVBYdv2D_aPbg4EGRvaC5zZWN1cmVkbnMuZXUKL2Rucy1xdWVyeQ

## securedns-ipv6-doh

Uncensored and no logging (IPv6, DoH protocol)

sdns://AgcAAAAAAAAAHFsyYTAzOmIwYzA6MDoxMDEwOjplOWE6MzAwMV0gPhoaD2xT8-l6SS1XCEtbmAcFnuBXqxUFh2_YP9o9uDgQZG9oLnNlY3VyZWRucy5ldQovZG5zLXF1ZXJ5

## sfw.scaleway-fr

Uses deep learning to block adult websites. Free, DNSSEC, no logs.
Hosted in Paris, running on a 1-XS server donated by Scaleway.com
Maintained by Frank Denis - https://fr.dnscrypt.info/sfw.html

sdns://AQMAAAAAAAAADzE2My4xNzIuMTgwLjEyNSDfYnO_x1IZKotaObwMhaw_-WRF1zZE9mJygl01WPGh_x8yLmRuc2NyeXB0LWNlcnQuc2Z3LnNjYWxld2F5LWZy

## soltysiak

Public DNSCrypt server in Poland

sdns://AQcAAAAAAAAAFDE3OC4yMTYuMjAxLjIyMjoyMDUzICXE4YgpFUaXj5wrvbanr6QB7aBRBQhdUwPnGSjAZo8hHTIuZG5zY3J5cHQtY2VydC5zb2x0eXNpYWsuY29t

## suami

A non-censoring, non-logging, non-censoring, DNSSEC-capable,
DNSCrypt-enabled DNS resolver in France, using the official Docker image.
Maintained by @lucenera

sdns://AQcAAAAAAAAAETUxLjE1OC4xMDYuNDI6NDQzIF_Fc4XhzA13UNWqYsy2Uc5kQPn_oJ_N0lfFPVHeicOrHDIuZG5zY3J5cHQtY2VydC5mci5zdWFtaS5jb20

## trashvpn.de

dnscrypt-server Docker image : DNSSEC/Non-logged/Uncensored
Hosted in Germany

sdns://AQcAAAAAAAAAEjM3LjIyMS4xOTUuMTgxOjQ0MyAl_sppDIKYr4Er_QKZ1ee96Xy_f5ZZs5Dxo0EvV22IoBsyLmRuc2NyeXB0LWNlcnQudHJhc2h2cG4uZGU

## ventricle.us

Public DNSCrypt resolver provided by Jacob Henner

sdns://AQcAAAAAAAAADTEwNy4xNzAuNTcuMzQg6YXxGK1OPMZf8iUgGJDG9Vi3W1pS9WsXz-rBAFyLm6olMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeXB0LnZlbnRyaWNsZS51cw

## yandex

Yandex public DNS server (anycast)

sdns://AQQAAAAAAAAAEDc3Ljg4LjguNzg6MTUzNTMg04TAccn3RmKvKszVe13MlxTUB7atNgHhrtwG1W1JYyciMi5kbnNjcnlwdC1jZXJ0LmJyb3dzZXIueWFuZGV4Lm5ldA

## zeroaim-ipv6

dnscrypt-server Docker image : DNSSEC/Non-logged/Uncensored
Hosted in Germany

sdns://AQcAAAAAAAAAGVsyYTAzOjQwMDA6YjoyMjM6OjFdOjg0NDMgcrQcuGXx2fhX6rmtaP6aPXj8gumVIrn4GIrn6aTB1fUfMi5kbnNjcnlwdC1jZXJ0Lnplcm9haW0uZGUtaXB2Ng

## opennic-bongobow

Non-logging OpenNIC resolver in Munich, Germany

sdns://AQYAAAAAAAAAEjUuMTg5LjE3MC4xOTY6NTM1MyBUNSxVQDuC7pPEB_3CNESXDZpW7yK_z_nskJzNMiQyaygyLmRuc2NyeXB0LWNlcnQubnMxNi5kZS5kbnMub3Blbm5pYy5nbHVl

## opennic-bongobow-ipv6

Non-logging OpenNIC resolver in Munich, Germany

sdns://AQYAAAAAAAAAIFsyYTAyOmMyMDc6MjAwODoyNTIwOjUzOjoxXTo1MzUzIFQ1LFVAO4Luk8QH_cI0RJcNmlbvIr_P-eyQnM0yJDJrKDIuZG5zY3J5cHQtY2VydC5uczE2LmRlLmRucy5vcGVubmljLmdsdWU

## opennic-R4SAS

Non-logging OpenNIC resolver in France

sdns://AQYAAAAAAAAAETE1MS44MC4yMjIuNzk6NDQzIO4Y9lZnORlvodxu39dnm6mFruwTRnlmovbEga4Fyw3TIDIuZG5zY3J5cHQtY2VydC5vcGVubmljLmkycGQueHl6

## opennic-R4SAS-ipv6

Non-logging OpenNIC resolver in France

sdns://AQYAAAAAAAAAG1syMDAxOjQ3MDoxZjE1OmI4MDo6NTNdOjQ0MyDuGPZWZzkZb6Hcbt_XZ5upha7sE0Z5ZqL2xIGuBcsN0yAyLmRuc2NyeXB0LWNlcnQub3Blbm5pYy5pMnBkLnh5eg

## dnsforfamily

Block adult websites, gambling websites and advertisements. No DNS queries are logged. As of March 2019 2.1million websites are blocked and new websites are added to blacklist daily.

Provided by: https://dnsforfamily.com

sdns://AQIAAAAAAAAADDc4LjQ3LjY0LjE2MSATJeLOABXNSYcSJIoqR5_iUYz87Y4OecMLB84aEAKPrRBkbnNmb3JmYW1pbHkuY29t

## dnsforfamily-v6

Block adult websites, gambling websites and advertisements. No DNS queries are logged. As of March 2019 2.1million websites are blocked and new websites are added to blacklist daily.

Provided by: https://dnsforfamily.com

sdns://AQIAAAAAAAAAF1syYTAxOjRmODoxYzE3OjRkZjg6OjFdIGN4CrSY4fb2hK8voFJL3GKiM7xQNwkKGH4b0k7LmMPxEGRuc2ZvcmZhbWlseS5jb20
