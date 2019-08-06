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

## aaflalo-me-nyc

DNS-over-HTTPS server running dns-over-https with PiHole for Adblocking in NYC, USA.

Non-logging, AD-filtering, supports DNSSEC.
Hosted in New York on a RamNode Cloud Instance.

sdns://AgMAAAAAAAAADjE2OC4yMzUuODEuMTY3ID4aGg9sU_PpekktVwhLW5gHBZ7gV6sVBYdv2D_aPbg4EmRucy1ueWMuYWFmbGFsby5tZQovZG5zLXF1ZXJ5

## aaflalo-me-gcp

Same as aaflalo-me-nyc. Use aaflalo-me-nyc.

Kept for backward compatibility with people use this server.

sdns://AgMAAAAAAAAADjE2OC4yMzUuODEuMTY3ID4aGg9sU_PpekktVwhLW5gHBZ7gV6sVBYdv2D_aPbg4EmRucy1ueWMuYWFmbGFsby5tZQovZG5zLXF1ZXJ5


## aaflalo-me

DNS-over-HTTPS server running dns-over-https with PiHole for Adblocking in NL.

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


## ads-dnswarden-dc1

DnsCrypt protocol . Non-logging, AD-filtering, supports DNSSEC. By https://dnswarden.com

sdns://AQMAAAAAAAAAEzExNi4yMDMuNzAuMTU2OjQ0NDMgenKjVeH-LU7Opsyq1ljKZz14fHsngOK8OOeQ-cR2mAsjMi5kbnNjcnlwdC1jZXJ0LmRuc3dhcmRlbi0xLWFkYmxvY2s


## ads-dnswarden-dc1-ipv6

DnsCrypt protocol . Non-logging, AD-filtering, supports DNSSEC. By https://dnswarden.com

sdns://AQMAAAAAAAAAHFsyYTAxOjRmODoxYzFjOjc1YjQ6OjFdOjQ0NDMgenKjVeH-LU7Opsyq1ljKZz14fHsngOK8OOeQ-cR2mAsjMi5kbnNjcnlwdC1jZXJ0LmRuc3dhcmRlbi0xLWFkYmxvY2s


## ads-dnswarden-dc2

DnsCrypt protocol . Non-logging, AD-filtering, supports DNSSEC. By https://dnswarden.com

sdns://AQMAAAAAAAAAEzExNi4yMDMuMzUuMjU1OjQ0NDMg-IlTTFFgMuUntnNV78COzbhJN9_OEOVWNgHhdg4BNXwjMi5kbnNjcnlwdC1jZXJ0LmRuc3dhcmRlbi0yLWFkYmxvY2s

## ads-dnswarden-dc2-ipv6

DnsCrypt protocol . Non-logging, AD-filtering, supports DNSSEC. By https://dnswarden.com

sdns://AQMAAAAAAAAAHFsyYTAxOjRmODoxYzFjOjVlNzc6OjFdOjQ0NDMg-IlTTFFgMuUntnNV78COzbhJN9_OEOVWNgHhdg4BNXwjMi5kbnNjcnlwdC1jZXJ0LmRuc3dhcmRlbi0yLWFkYmxvY2s

## ads-dnswarden-doh1

Dns-over-HTTPS protocol . Non-logging, AD-filtering, supports DNSSEC. By https://dnswarden.com

sdns://AgMAAAAAAAAADjExNi4yMDMuNzAuMTU2ID4aGg9sU_PpekktVwhLW5gHBZ7gV6sVBYdv2D_aPbg4EWRvaC5kbnN3YXJkZW4uY29tCC9hZGJsb2Nr

## ads-dnswarden-doh1-ipv6

Dns-over-HTTPS protocol . Non-logging, AD-filtering, supports DNSSEC. By https://dnswarden.com

sdns://AgMAAAAAAAAAF1syYTAxOjRmODoxYzFjOjc1YjQ6OjFdID4aGg9sU_PpekktVwhLW5gHBZ7gV6sVBYdv2D_aPbg4EWRvaC5kbnN3YXJkZW4uY29tCC9hZGJsb2Nr

## ads-dnswarden-doh2

Dns-over-HTTPS protocol . Non-logging, AD-filtering, supports DNSSEC. By https://dnswarden.com

sdns://AgMAAAAAAAAADjExNi4yMDMuMzUuMjU1ID4aGg9sU_PpekktVwhLW5gHBZ7gV6sVBYdv2D_aPbg4EWRvaC5kbnN3YXJkZW4uY29tCC9hZGJsb2Nr

## ads-dnswarden-doh2-ipv6

Dns-over-HTTPS protocol . Non-logging, AD-filtering, supports DNSSEC. By https://dnswarden.com

sdns://AgMAAAAAAAAAF1syYTAxOjRmODoxYzFjOjVlNzc6OjFdID4aGg9sU_PpekktVwhLW5gHBZ7gV6sVBYdv2D_aPbg4EWRvaC5kbnN3YXJkZW4uY29tCC9hZGJsb2Nr

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

sdns://AQcAAAAAAAAAFDIwNi4xODkuMTQyLjE3OTo1MzUzII5GJ8c4g6hRAwghulrn5dBB9KrvlbeCkBbLZR2HwyjJGTIuZG5zY3J5cHQtY2VydC5hcnZpbmQuaW8

## bottlepost-dns-nl

Provided by bottlepost.me
Hosted in The Netherlands, DNSSEC / No Logs / No Filter

sdns://AQcAAAAAAAAAEzE3OC4xMjguMjU1LjI4OjUzNTMgkr1k-Lp2d9IXiFlXoBAgFGZUCJSPW_x81Ec6ShkPsJYdMi5kbnNjcnlwdC1jZXJ0LmJvdHRsZXBvc3QubWU

## brahma-world

DNS-over-HTTPS / DNS over TLS server with PiHole. Filters ads, trackers and malware.
Forwards to Google servers.

Hosted in Bengaluru, India. (https://dns.brahma.world)

sdns://AgEAAAAAAAAADjEzNC4yMDkuMTQ2LjE2oD4aGg9sU_PpekktVwhLW5gHBZ7gV6sVBYdv2D_aPbg4ID4aGg9sU_PpekktVwhLW5gHBZ7gV6sVBYdv2D_aPbg4EGRucy5icmFobWEud29ybGQKL2Rucy1xdWVyeQ

## captnemo-in

Server running out of a Digital Ocean droplet in BLR1 region.
Maintained by Abhay Rana aka Nemo.

If you are within India, this might be a nice DNS server to use.

sdns://AQQAAAAAAAAAEjEzOS41OS40OC4yMjI6NDQzNCAFOt_yxaMpFtga2IpneSwwK6rV0oAyleham9IvhoceEBsyLmRuc2NyeXB0LWNlcnQuY2FwdG5lbW8uaW4

## charis

Public DNSCrypt server in Germany. DNSSEC, no logs, no filter.
Maintained by @lucenera

sdns://AQcAAAAAAAAAETUxLjE1LjEwNi4xNzY6NDQzIGcUiAnFqewnNLjh8DUYpcePX07pXc3sDOf2U-vpI55WHjIuZG5zY3J5cHQtY2VydC5hbXMuY2hhcmlzLmNvbQ

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

## dns.digitale-gesellschaft.ch

Public DoH resolver operated by the Digital Society (https://www.digitale-gesellschaft.ch).
Hosted in Zurich, Switzerland.

Non-logging, non-filtering, supports DNSSEC.

sdns://AgcAAAAAAAAAAKA-GhoPbFPz6XpJLVcIS1uYBwWe4FerFQWHb9g_2j24OKBoo-sB-l8CxNNfOhHQBMrwiyJL7qfXnFiMfxPIYTNgLqDvR4Wu5wydV1_nM4MG2T6nlhHl_tzvU2LdZsmLYLstvSAcVDa2UaK1QVwWz9ltGpcJ_ZyPJ-73XPlz2YL_5o5Y8BxkbnMuZGlnaXRhbGUtZ2VzZWxsc2NoYWZ0LmNoCi9kbnMtcXVlcnk

## dnscrypt.ca-1

Free, Canadian, uncensored, no-logs, encrypted, and DNSSEC validated
DNS service for your pleasure.

sdns://AQcAAAAAAAAAEjE5Mi45OS4xODMuMTMyOjQ0MyAaU6PJUHicvdELGTOkaJtshGpA8bc9F1KuysmCnst84h0yLmRuc2NyeXB0LWNlcnQuZG5zY3J5cHQuY2EtMQ

## dnscrypt.ca-2

Free, Canadian, uncensored, no-logs, encrypted, and DNSSEC validated
DNS service for your pleasure.

sdns://AQcAAAAAAAAAEjE5Mi45OS4xODMuMTMzOjQ0MyABCFSrO1an7vnTkVj-9oIL_5OiNXyJFgjbnhXTu-ARhR0yLmRuc2NyeXB0LWNlcnQuZG5zY3J5cHQuY2EtMg

## dnscrypt.ca-1-ipv6

Free, Canadian, uncensored, no-logs, encrypted, and DNSSEC validated
DNS service for your pleasure.

sdns://AQcAAAAAAAAAHFsyNjA3OjUzMDA6NjA6NGFhODo6NjAwXTo0NDMgINkZ1Ow8UAjNxwpR6itPy_6KmTxkMdDsaB1epzhFq4AiMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeXB0LmNhLTEtaXB2Ng

## dnscrypt.ca-2-ipv6

Free, Canadian, uncensored, no-logs, encrypted, and DNSSEC validated
DNS service for your pleasure.

sdns://AQcAAAAAAAAAHFsyNjA3OjUzMDA6NjA6NGFhODo6NzAwXTo0NDMgmHxwqJfb2hUaNK1LVeqADvVhzASq1cV90sPYYfwX9CkiMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeXB0LmNhLTItaXB2Ng

## dnswarden-dc1

DnsCrypt protocol . Non-logging, supports DNSSEC. By https://dnswarden.com

sdns://AQcAAAAAAAAAEzExNi4yMDMuNzAuMTU2Ojg0NDMgenKjVeH-LU7Opsyq1ljKZz14fHsngOK8OOeQ-cR2mAskMi5kbnNjcnlwdC1jZXJ0LmRuc3dhcmRlbi0xLXVuY2Vuc29y

## dnswarden-dc1-ipv6

DnsCrypt protocol . Non-logging,supports DNSSEC. By https://dnswarden.com

sdns://AQcAAAAAAAAAHFsyYTAxOjRmODoxYzFjOjc1YjQ6OjFdOjg0NDMgenKjVeH-LU7Opsyq1ljKZz14fHsngOK8OOeQ-cR2mAskMi5kbnNjcnlwdC1jZXJ0LmRuc3dhcmRlbi0xLXVuY2Vuc29y


## dnswarden-dc2

DnsCrypt protocol . Non-logging, supports DNSSEC. By https://dnswarden.com

sdns://AQcAAAAAAAAAEzExNi4yMDMuMzUuMjU1Ojg0NDMg-IlTTFFgMuUntnNV78COzbhJN9_OEOVWNgHhdg4BNXwkMi5kbnNjcnlwdC1jZXJ0LmRuc3dhcmRlbi0yLXVuY2Vuc29y

## dnswarden-dc2-ipv6

DnsCrypt protocol . Non-logging,supports DNSSEC. By https://dnswarden.com

sdns://AQcAAAAAAAAAHFsyYTAxOjRmODoxYzFjOjVlNzc6OjFdOjg0NDMg-IlTTFFgMuUntnNV78COzbhJN9_OEOVWNgHhdg4BNXwkMi5kbnNjcnlwdC1jZXJ0LmRuc3dhcmRlbi0yLXVuY2Vuc29y


## dnswarden-doh1

Dns-over-HTTPS protocol . Non-logging, supports DNSSEC. By https://dnswarden.com

sdns://AgcAAAAAAAAADjExNi4yMDMuNzAuMTU2ID4aGg9sU_PpekktVwhLW5gHBZ7gV6sVBYdv2D_aPbg4EWRvaC5kbnN3YXJkZW4uY29tCy91bmNlbnNvcmVk

## dnswarden-doh1-ipv6

Dns-over-HTTPS protocol . Non-logging, supports DNSSEC. By https://dnswarden.com

sdns://AgcAAAAAAAAAF1syYTAxOjRmODoxYzFjOjc1YjQ6OjFdID4aGg9sU_PpekktVwhLW5gHBZ7gV6sVBYdv2D_aPbg4EWRvaC5kbnN3YXJkZW4uY29tCy91bmNlbnNvcmVk

## dnswarden-doh2

Dns-over-HTTPS protocol . Non-logging, supports DNSSEC. By https://dnswarden.com

sdns://AgcAAAAAAAAADjExNi4yMDMuMzUuMjU1ID4aGg9sU_PpekktVwhLW5gHBZ7gV6sVBYdv2D_aPbg4EWRvaC5kbnN3YXJkZW4uY29tCy91bmNlbnNvcmVk

## dnswarden-doh2-ipv6

Dns-over-HTTPS protocol . Non-logging, supports DNSSEC. By https://dnswarden.com

sdns://AgcAAAAAAAAAF1syYTAxOjRmODoxYzFjOjVlNzc6OjFdID4aGg9sU_PpekktVwhLW5gHBZ7gV6sVBYdv2D_aPbg4EWRvaC5kbnN3YXJkZW4uY29tCy91bmNlbnNvcmVk



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

## dnscrypt.eu-dk

Free, non-logged, uncensored. Hosted by Netgroup.

sdns://AQcAAAAAAAAADDc3LjY2Ljg0LjIzMyA3SFWF47nQiP0lrTawNwH1UgzWSJ6a3VIUV0lVnwqZVSUyLmRuc2NyeXB0LWNlcnQucmVzb2x2ZXIyLmRuc2NyeXB0LmV1

## dnscrypt.eu-dk-ipv6

Free, non-logged, uncensored. Hosted by Netgroup.

sdns://AQcAAAAAAAAAFFsyMDAxOjE0NDg6MjQzOjpkYzJdIDdIVYXjudCI_SWtNrA3AfVSDNZInprdUhRXSVWfCplVJTIuZG5zY3J5cHQtY2VydC5yZXNvbHZlcjIuZG5zY3J5cHQuZXU

## dnscrypt.eu-nl

Free, non-logged, uncensored. Hosted by RamNode.

sdns://AQcAAAAAAAAADjE3Ni41Ni4yMzcuMTcxIGfADywhxVSBRd18tGonGvLrlpkxQKMJtiuNFlMRhZxmJTIuZG5zY3J5cHQtY2VydC5yZXNvbHZlcjEuZG5zY3J5cHQuZXU

## dnscrypt.nl-ns0

DNSCrypt v2 server in Amsterdam, the Netherlands. DNSSEC, no logs, uncensored, recursive DNS. https://dnscrypt.nl

sdns://AQcAAAAAAAAADDQ1Ljc2LjM1LjIxMiBMhPuMBRFd-l-Xxe0DKRNwx4q81k4V3VOrCN5y-4RKyh8yLmRuc2NyeXB0LWNlcnQubnMwLmRuc2NyeXB0Lm5s

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

## dnscrypt-01.adsnomore.io

DNSCrypt server located in Nuremberg, Germany.
No logging, DNSSEC, disk encryption, Canonical Livepatch and monitored 24/7.
Uses the official Docker image.

sdns://AQcAAAAAAAAAETk0LjEzMC4xNzguNTY6NDQzIIxpj-7XPoT_79rA9pnvVGz0bIQRuEL-eI-0NlYJaGcpJjIuZG5zY3J5cHQtY2VydC5kbnNjcnlwdC0wMS5tYWRwb255Lmlv

## dnscrypt-02.adsnomore.io

DNSCrypt server located in Miami, USA.
No logging, DNSSEC, disk encryption, Canonical Livepatch and monitored 24/7.
Uses the official Docker image.

sdns://AQcAAAAAAAAAETE0MC44Mi4yNi4xMDM6NDQzIE15px_otxEaCZ20DybtbfMu92IH3Ritg83ibv6LeizTKTIuZG5zY3J5cHQtY2VydC5kbnNjcnlwdC0wMi5tYWRwb255LnNwYWNl

## dnscrypt-03.adsnomore.io

DNSCrypt server located in São Paulo, Brazil.
No logging, DNSSEC, disk encryption, Canonical Livepatch and monitored 24/7.
Uses the official Docker image.

sdns://AQcAAAAAAAAAEDUuMTg4LjIzOC42ODo0NDMg1uv1UTjfRdCF1XDI3T10v4EXWcdK6x8qM5Qut7bwb_gpMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeXB0LTAzLm1hZHBvbnkuc3BhY2U

## dnscrypt-04.adsnomore.io

DNSCrypt server located in Tokyo, Japan.
No logging, DNSSEC, disk encryption, Canonical Livepatch and monitored 24/7.
Uses the official Docker image.

sdns://AQcAAAAAAAAAEDQ1LjMyLjMxLjIzMTo0NDMgmk18Se_bsOdRNFJ64Lrw5MJ83y_au6WNrh3lZOceiqgpMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeXB0LTA0Lm1hZHBvbnkuc3BhY2U

## dnscrypt-05.adsnomore.io

DNSCrypt server located in Sydney, Australia.
No logging, DNSSEC, disk encryption, Canonical Livepatch and monitored 24/7.
Uses the official Docker image.

sdns://AQcAAAAAAAAAETE0OS4yOC4xNjguNjI6NDQzIENfI6UCxKdNccBA9YW-OhkV-HB_b_Yj5nQbq-gM1TAMKTIuZG5zY3J5cHQtY2VydC5kbnNjcnlwdC0wNS5tYWRwb255LnNwYWNl

## dnscrypt-06.adsnomore.io

DNSCrypt server located in Amsterdam, Netherlands.
No logging, DNSSEC, disk encryption, Canonical Livepatch and monitored 24/7.
Uses the official Docker image.

sdns://AQcAAAAAAAAAEjk1LjE3OS4xNzguMTAwOjQ0MyCzDTlSDfD9-UOciubW46-f6tsh8o60Rt1m4i7XH5hBqykyLmRuc2NyeXB0LWNlcnQuZG5zY3J5cHQtMDYubWFkcG9ueS5zcGFjZQ

## dnscrypt-07.adsnomore.io

DNSCrypt server located in Singapore.
No logging, DNSSEC, disk encryption, Canonical Livepatch and monitored 24/7.
Uses the official Docker image.

sdns://AQcAAAAAAAAAEjEzOS4xODAuMjE2LjgzOjQ0MyBPxDlEgU5vJPp0n-Zh505hgFMSBQj8CQc7p9uUaIWigSkyLmRuc2NyeXB0LWNlcnQuZG5zY3J5cHQtMDcubWFkcG9ueS5zcGFjZQ

## dnscrypt-ch-blahdns-ipv6

Blocks ad and Tracking, no Logging, DNSSEC, Hosted in Switzerland. By https://blahdns.com/

sdns://AQMAAAAAAAAAJVsyYTBhOmU1YzA6MjoyOjA6YzhmZjpmZTY4OmJmNDhdOjg0NDMgyJjbSS4IgTY_2KH3NVGG0DNIgBPzLEqf8r00nAbcUxQbMi5kbnNjcnlwdC1jZXJ0LmJsYWhkbnMuY29t

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

## dns.sb

DNSSEC-enabled DoH server by https://xtom.com/
Using Cloudflare as a frontend.

https://dns.sb

sdns://AgUAAAAAAAAAACA9pLcWNKQTwc7zSqltJaQBY01M82w7Ezx0KU5I5jcBKgpkb2guZG5zLnNiCi9kbnMtcXVlcnk

## doh-ch-blahdns-v6

Blocks ad and Tracking, no Logging, DNSSEC, Hosted in Switzerland. By https://blahdns.com/

sdns://AgMAAAAAAAAAIFsyYTBhOmU1YzA6MjoyOjA6YzhmZjpmZTY4OmJmNDhdABJkb2gtY2guYmxhaGRucy5jb20KL2Rucy1xdWVyeQ

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

## encrypt-town

Non-logging, non-censoring, DNSSEC-capable DNSCrypt server located in Northern California.
By https://encrypt.town

sdns://AQcAAAAAAAAAETEzLjUyLjEzOC4yMTc6ODg4IBBFTcFTuQTe-sR5uokPA7FcM00RRnHKBZylyElKcwT2HDIuZG5zY3J5cHQtY2VydC5lbmNyeXB0LnRvd24

## ev-va

Non-logging, uncensored DNS resolver provided by evilvibes.com
Location: Vancouver, Canada

sdns://AQcAAAAAAAAADTIzLjExMS43NC4yMTYg3_NERwhF2C4tPlnR0CSeIAmRC3wwXtucNhKMVoW_prQjMi5kbnNjcnlwdC1jZXJ0LmV2LXZhLmV2aWx2aWJlcy5jb20

## ev-to

Non-logging, uncensored DNS resolver provided by evilvibes.com
Location: Toronto, Canada

sdns://AQcAAAAAAAAADTIzLjExMS42OS4xMjYgTvUHn_uZRDf0m10w-HJPpAVY7_UcoDKNuXpXcWJSVwsjMi5kbnNjcnlwdC1jZXJ0LmV2LXRvLmV2aWx2aWJlcy5jb20

## faelix

An open DoH resolver operated by https://faelix.net/ with nodes in UK and CH.

sdns://AgcAAAAAAAAAACA-GhoPbFPz6XpJLVcIS1uYBwWe4FerFQWHb9g_2j24OA9yZG5zLmZhZWxpeC5uZXQBLw

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

sdns://AgUAAAAAAAAABzguOC44LjigHvYkz_9ea9O63fP92_3qVlRn43cpncfuZnUWbzAMwbkgdoAkR6AZkxo_AEMExT_cbBssN43Evo9zs5_ZyWnftEUKZG5zLmdvb2dsZQovZG5zLXF1ZXJ5

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

sdns://AQcAAAAAAAAADjE3OC44Mi4xMDIuMTkwIBuR6jsTpCthR_P_b_jSxBJF9ch_jC7ZTVp7EHV2YgLaGDIuZG5zY3J5cHQtY2VydC5pYmtzdHVybQ

## id-gmail

Non-Logging DNSCrypt server located in Singapore.
Filters out ads, trackers and malware, supports DNSSEC, provided by id-gmail.

sdns://AQMAAAAAAAAADjE3NC4xMzguMjEuMTI4IO-WgGbo2ZTwZdg-3dMa7u31bYZXRj5KykfN1_6Xw9T2HDIuZG5zY3J5cHQtY2VydC5kbnMudGlhci5hcHA

## id-gmail-ipv6

Non-Logging DNSCrypt (IPv6) server located in Singapore.
Filters out ads, trackers and malware, supports DNSSEC, provided by id-gmail.

sdns://AQMAAAAAAAAAG1syNDAwOjYxODA6MDpkMDo6NWY2ZTo0MDAxXSDvloBm6NmU8GXYPt3TGu7t9W2GV0Y-SspHzdf-l8PU9hwyLmRuc2NyeXB0LWNlcnQuZG5zLnRpYXIuYXBw

## id-gmail-doh

Non-Logging DNS-over-HTTPS server located in Singapore.
Filters out ads, trackers and malware, supports DNSSEC, provided by id-gmail.

sdns://AgMAAAAAAAAADjE3NC4xMzguMjkuMTc1ID4aGg9sU_PpekktVwhLW5gHBZ7gV6sVBYdv2D_aPbg4DGRvaC50aWFyLmFwcAovZG5zLXF1ZXJ5

## id-gmail-doh-ipv6

Non-Logging DNS-over-HTTPS (IPv6) server located in Singapore.
Filters out ads, trackers and malware, supports DNSSEC, provided by id-gmail.

sdns://AgMAAAAAAAAAG1syNDAwOjYxODA6MDpkMDo6NWY3Mzo0MDAxXSA-GhoPbFPz6XpJLVcIS1uYBwWe4FerFQWHb9g_2j24OAxkb2gudGlhci5hcHAKL2Rucy1xdWVyeQ

## iij

DoH server operated by Internet Initiative Japan in Tokyo.
https://www.iij.ad.jp/

sdns://AgUAAAAAAAAACjEwMy4yLjU3LjUgs2lfGAQCPrV0DPQqOkPci0Jei0GhMK_ne-QHwPbfn4oRcHVibGljLmRucy5paWouanAKL2Rucy1xdWVyeQ

## ipredator

Public DNSCrypt server in Sweden provided by Ipredator.se

sdns://AQcAAAAAAAAADTE5NC4xMzIuMzIuMzIgxExWaqjWRsQysQT1PQCWGzLccc8cBL2esBPkgOekeCgcMi5kbnNjcnlwdC1jZXJ0LmlwcmVkYXRvci5zZQ

## jp.tiar.app

Non-Logging, Non-Filtering DNSCrypt server in Japan.
No ECS, Support DNSSEC

sdns://AQcAAAAAAAAAEjE3Mi4xMDQuOTMuODA6MTQ0MyAyuHY-8b9lNqHeahPAzW9IoXnjiLaZpTeNbVs8TN9UUxsyLmRuc2NyeXB0LWNlcnQuanAudGlhci5hcHA

## jp.tiar.app-ipv6

Non-Logging, Non-Filtering DNSCrypt (IPv6) server in Japan.
No ECS, Support DNSSEC

sdns://AQcAAAAAAAAAJVsyNDAwOjg5MDI6OmYwM2M6OTFmZjpmZWRhOmM1MTRdOjE0NDMgMrh2PvG_ZTah3moTwM1vSKF544i2maU3jW1bPEzfVFMbMi5kbnNjcnlwdC1jZXJ0LmpwLnRpYXIuYXBw

## jp.tiar.app-doh

Non-Logging, Non-Filtering DNS-over-HTTPS server in Japan.
No ECS, Support DNSSEC

sdns://AgcAAAAAAAAADTE3Mi4xMDQuOTMuODAgPhoaD2xT8-l6SS1XCEtbmAcFnuBXqxUFh2_YP9o9uDgLanAudGlhci5hcHAKL2Rucy1xdWVyeQ

## jp.tiar.app-doh-ipv6

Non-Logging, Non-Filtering DNS-over-HTTPS (IPv6) server in Japan.
No ECS, Support DNSSEC

sdns://AgcAAAAAAAAAIFsyNDAwOjg5MDI6OmYwM2M6OTFmZjpmZWRhOmM1MTRdID4aGg9sU_PpekktVwhLW5gHBZ7gV6sVBYdv2D_aPbg4C2pwLnRpYXIuYXBwCi9kbnMtcXVlcnk

## nawala-childprotection

Internet filtering system (anycast), protecting child from inappropriate websites and abusive contents.
By http://nawala.id in Indonesia.

sdns://AQAAAAAAAAAADzE4MC4xMzEuMTQ0LjE0NCDGC-b_38Dj4-ikI477AO1GXcLPfETOFpE36KZIHdOzLhkyLmRuc2NyeXB0LWNlcnQubmF3YWxhLmlk

## nextdns

NextDNS is a cloud-based private DNS service that gives you full control
over what is allowed and what is blocked on the Internet.

DNSSEC, Anycast.

https://www.nextdns.io/

sdns://AgUAAAAAAAAACzUuMTgyLjIwOC4wID4aGg9sU_PpekktVwhLW5gHBZ7gV6sVBYdv2D_aPbg4DmRucy5uZXh0ZG5zLmlvDy9kbnNjcnlwdC1wcm94eQ

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

## oszx

Secure DNS Project by PumpleX - Hosted in the UK by OVH
No Logging / Ad-Blocking
Information at https://dns.oszx.co

sdns://AQIAAAAAAAAAETUxLjM4LjgzLjE0MTo1MzUzIMwm9_oYw26P4JIVoDhJ_5kFDdNxX1ke4fEzL1V5bwEjFzIuZG5zY3J5cHQtY2VydC5vc3p4LmNv

## powerdns-doh

By PowerDNS/Open-Xchange https://powerdns.org

sdns://AgcAAAAAAAAAACA-GhoPbFPz6XpJLVcIS1uYBwWe4FerFQWHb9g_2j24OBBkb2gucG93ZXJkbnMub3JnAS8

## publicarray-au

DNSSEC • OpenNIC • Non-logging • Uncensored - hosted on vultr.com
Maintained by publicarray - https://dns.seby.io

sdns://AQcAAAAAAAAADDQ1Ljc2LjExMy4zMSAIVGh4i6eKXqlF6o9Fg92cgD2WcDvKQJ7v_Wq4XrQsVhsyLmRuc2NyeXB0LWNlcnQuZG5zLnNlYnkuaW8

## publicarray-au-doh

DNSSEC • OpenNIC • Non-logging • Uncensored - hosted on vultr.com
Maintained by publicarray - https://dns.seby.io

sdns://AgcAAAAAAAAADDQ1Ljc2LjExMy4zMSA-GhoPbFPz6XpJLVcIS1uYBwWe4FerFQWHb9g_2j24OBBkb2guc2VieS5pbzo4NDQzCi9kbnMtcXVlcnk

## publicarray-au2

DNSSEC • OpenNIC • Non-logging • Uncensored - hosted on ovh.com.au
Maintained by publicarray - https://dns.seby.io

sdns://AQcAAAAAAAAADTEzOS45OS4yMjIuNzIgCwVoTw0L4dgal5LC1FbZUtHtLR_rVuV6rVnxO95e4GUbMi5kbnNjcnlwdC1jZXJ0LmRucy5zZWJ5Lmlv

## publicarray-au2-doh

DNSSEC • OpenNIC • Non-logging • Uncensored - hosted on ovh.com.au
Maintained by publicarray - https://dns.seby.io

sdns://AgcAAAAAAAAADTEzOS45OS4yMjIuNzIgPhoaD2xT8-l6SS1XCEtbmAcFnuBXqxUFh2_YP9o9uDgRZG9oLTIuc2VieS5pbzo0NDMKL2Rucy1xdWVyeQ

## qag.me

Ad-blocking DNS cache.
Home Server running on a static IP in Bangalore / Bengaluru, INDIA.
Maintained by Cruisemaniac (https://cruisemaniac.com) aka Ashwin Murali.

sdns://AQMAAAAAAAAAETEwNi41MS40NC41MTo0NDM0ING7Nd6ynAGtsN3sIyHYjkWg7pcKTGvRz4dOe0sbDZjcFjIuZG5zY3J5cHQtY2VydC5xYWcubWU

## quad101

DNSSEC-aware public resolver by the Taiwan Network Information Center (TWNIC)
https://101.101.101.101/index_en.html

sdns://AgcAAAAAAAAAACAoPxWWFWiOuUdTdn7SvYpzbNqr_iDmmJrktihy4wca5gxkbnMudHduaWMudHcKL2Rucy1xdWVyeQ

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

sdns://AQcAAAAAAAAADjIxMi40Ny4yMjguMTM2IOgBuE6mBr-wusDOQ0RbsV66ZLAvo8SqMa4QY2oHkDJNHzIuZG5zY3J5cHQtY2VydC5mci5kbnNjcnlwdC5vcmc

## securedns

Uncensored and no logging (DNSCrypt protocol)

sdns://AQcAAAAAAAAAEzE0Ni4xODUuMTY3LjQzOjUzNTMg9J8sc01itoYxntB-aRlDOy8ThfQe-8ovF21ZCy5FPoYcMi5kbnNjcnlwdC1jZXJ0LnNlY3VyZWRucy5ldQ

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

sdns://AQcAAAAAAAAAEjE2My4xNzIuMTM1LjkxOjQ0MyAa1c2tko5PA1_fBzSXfz3pIB-3QKK04gA99gF7D_ePqBwyLmRuc2NyeXB0LWNlcnQuZnIuc3VhbWkuY29t

## ventricle.us

Public DNSCrypt resolver provided by Jacob Henner

sdns://AQcAAAAAAAAADTEwNy4xNzAuNTcuMzQg6YXxGK1OPMZf8iUgGJDG9Vi3W1pS9WsXz-rBAFyLm6olMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeXB0LnZlbnRyaWNsZS51cw

## yandex

Yandex public DNS server (anycast)

sdns://AQQAAAAAAAAAEDc3Ljg4LjguNzg6MTUzNTMg04TAccn3RmKvKszVe13MlxTUB7atNgHhrtwG1W1JYyciMi5kbnNjcnlwdC1jZXJ0LmJyb3dzZXIueWFuZGV4Lm5ldA

## opennic-R4SAS

Non-logging OpenNIC resolver in France

sdns://AQYAAAAAAAAAETE1MS44MC4yMjIuNzk6NDQzIO4Y9lZnORlvodxu39dnm6mFruwTRnlmovbEga4Fyw3TIDIuZG5zY3J5cHQtY2VydC5vcGVubmljLmkycGQueHl6

## dnsforfamily

Block adult websites, gambling websites and advertisements. No DNS queries are logged. As of March 2019 2.1million websites are blocked and new websites are added to blacklist daily.

Provided by: https://dnsforfamily.com

sdns://AQIAAAAAAAAADDc4LjQ3LjY0LjE2MSATJeLOABXNSYcSJIoqR5_iUYz87Y4OecMLB84aEAKPrRBkbnNmb3JmYW1pbHkuY29t

## dnsforfamily-v6

Block adult websites, gambling websites and advertisements. No DNS queries are logged. As of March 2019 2.1million websites are blocked and new websites are added to blacklist daily.

Provided by: https://dnsforfamily.com

sdns://AQIAAAAAAAAAF1syYTAxOjRmODoxYzE3OjRkZjg6OjFdIGN4CrSY4fb2hK8voFJL3GKiM7xQNwkKGH4b0k7LmMPxEGRuc2ZvcmZhbWlseS5jb20
