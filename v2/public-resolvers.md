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


## a-and-a

Non-logging DoH server in the UK operated by Andrews & Arnold Ltd, a
company providing Internet connectivity and VoIP in the UK.

https://www.aa.net.uk/legal/dohdot-disclaimer/

sdns://AgcAAAAAAAAADTIxNy4xNjkuMjAuMjMgPhoaD2xT8-l6SS1XCEtbmAcFnuBXqxUFh2_YP9o9uDgNZG5zLmFhLm5ldC51awovZG5zLXF1ZXJ5


## aaflalo-me

DNS-over-HTTPS server running dns-over-https with PiHole for Adblocking in NL.

Non-logging, AD-filtering, supports DNSSEC.
Hosted in Netherlands on a RamNode VPS.

sdns://AgMAAAAAAAAADjE3Ni41Ni4yMzYuMTc1ID4aGg9sU_PpekktVwhLW5gHBZ7gV6sVBYdv2D_aPbg4DmRucy5hYWZsYWxvLm1lCi9kbnMtcXVlcnk


## aaflalo-me-gcp

Same as aaflalo-me-nyc. Use aaflalo-me-nyc.

Kept for backward compatibility with people use this server.

sdns://AgMAAAAAAAAADjE2OC4yMzUuODEuMTY3ID4aGg9sU_PpekktVwhLW5gHBZ7gV6sVBYdv2D_aPbg4EmRucy1ueWMuYWFmbGFsby5tZQovZG5zLXF1ZXJ5


## aaflalo-me-nyc

DNS-over-HTTPS server running dns-over-https with PiHole for Adblocking in NYC, USA.

Non-logging, AD-filtering, supports DNSSEC.
Hosted in New York on a RamNode Cloud Instance.

sdns://AgMAAAAAAAAADjE2OC4yMzUuODEuMTY3ID4aGg9sU_PpekktVwhLW5gHBZ7gV6sVBYdv2D_aPbg4EmRucy1ueWMuYWFmbGFsby5tZQovZG5zLXF1ZXJ5


## adguard-dns

Remove ads and protect your computer from malware

sdns://AQMAAAAAAAAAFDE3Ni4xMDMuMTMwLjEzMDo1NDQzINErR_JS3PLCu_iZEIbq95zkSV2LFsigxDIuUso_OQhzIjIuZG5zY3J5cHQuZGVmYXVsdC5uczEuYWRndWFyZC5jb20


## adguard-dns-doh

Remove ads and protect your computer from malware (over DoH)

sdns://AgMAAAAAAAAADzE3Ni4xMDMuMTMwLjEzMCD5_zfwLmMstzhwJcB-V5CKPTcbfJXYzdA5DeIx7ZQ6Eg9kbnMuYWRndWFyZC5jb20KL2Rucy1xdWVyeQ


## adguard-dns-family

Adguard DNS with safesearch and adult content blocking

sdns://AQMAAAAAAAAAFDE3Ni4xMDMuMTMwLjEzMjo1NDQzILgxXdexS27jIKRw3C7Wsao5jMnlhvhdRUXWuMm1AFq6ITIuZG5zY3J5cHQuZmFtaWx5Lm5zMS5hZGd1YXJkLmNvbQ


## adguard-dns-family-doh

Adguard DNS with safesearch and adult content blocking (over DoH)

sdns://AgMAAAAAAAAADzE3Ni4xMDMuMTMwLjEzMiD5_zfwLmMstzhwJcB-V5CKPTcbfJXYzdA5DeIx7ZQ6EhZkbnMtZmFtaWx5LmFkZ3VhcmQuY29tCi9kbnMtcXVlcnk


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


## ams-dnscrypt-nl

Resolver in Amsterdam. DNSCrypt protocol. Non-logging, non-filtering, DNSSEC.
Forward DNS to anoymized DNS servers.

sdns://AQcAAAAAAAAAEjUxLjE1LjEyNC4yMDg6NDM0MyC8E4j1dj497HXxyQ_JFb-2iurf6xxF9phRgGOcYOfxYh8yLmRuc2NyeXB0LWNlcnQuYW1zLWRuc2NyeXB0LW5s


## ams-dnscrypt-noads-nl

Resolver in Amsterdam. DNSCrypt protocol. Non-logging, AD-filtering, DNSSEC.
Forward DNS to anoymized DNS servers.

sdns://AQMAAAAAAAAAEjUxLjE1LjEyNC4yMDg6NDQ0MyC8E4j1dj497HXxyQ_JFb-2iurf6xxF9phRgGOcYOfxYiUyLmRuc2NyeXB0LWNlcnQuYW1zLWRuc2NyeXB0LW5vYWRzLW5s


## ams-doh-nl

Resolver in Amsterdam. DoH protocol. Non-logging, non-filtering, DNSSEC.
Forward DNS to anoymized DNS servers.

sdns://AgcAAAAAAAAAETUxLjE1LjEyNC4yMDg6NDQzABRkbnMuYWxla2JlcmcubmV0OjQ0MwEv


## ams-doh-noads-nl

Resolver in Amsterdam. DoH protocol. Non-logging, ads-filtering, DNSSEC.
Forward DNS to anoymized DNS servers.

sdns://AgMAAAAAAAAAETUxLjE1LjEyNC4yMDg6NDQzABRkbnMuYWxla2JlcmcubmV0OjQ0MwYvbm9hZHM


## arvind-io

Public resolver by EnKrypt (https://arvind.io).
Hosted in Bangalore, India.

Non-logging, non-filtering, supports DNSSEC.

sdns://AQcAAAAAAAAAFDIwNi4xODkuMTQyLjE3OTo1MzUzII5GJ8c4g6hRAwghulrn5dBB9KrvlbeCkBbLZR2HwyjJGTIuZG5zY3J5cHQtY2VydC5hcnZpbmQuaW8


## att

AT&T test DoH server.

sdns://AgQAAAAAAAAAAKBLTrSwdCmLgotcADCVoQtFI_uVHAyINIsJxT5bq6QIoyD2Hldod9qWUClMzLX5bHX8txvaG7xGRjZ8Tr7aidcxjxBkb2h0cmlhbC5hdHQubmV0Ci9kbnMtcXVlcnk


## brahma-world

DNS-over-HTTPS / DNS over TLS server with PiHole. Filters ads, trackers and malware.

Hosted in Bengaluru, India. (https://dns.brahma.world)

sdns://AgMAAAAAAAAADjEzNC4yMDkuMTQ2LjE2oD4aGg9sU_PpekktVwhLW5gHBZ7gV6sVBYdv2D_aPbg4ID4aGg9sU_PpekktVwhLW5gHBZ7gV6sVBYdv2D_aPbg4EGRucy5icmFobWEud29ybGQKL2Rucy1xdWVyeQ


## captnemo-in

Server running out of a Digital Ocean droplet in BLR1 region.
Maintained by Abhay Rana aka Nemo.

If you are within India, this might be a nice DNS server to use.

sdns://AQQAAAAAAAAAEjEzOS41OS40OC4yMjI6NDQzNCAFOt_yxaMpFtga2IpneSwwK6rV0oAyleham9IvhoceEBsyLmRuc2NyeXB0LWNlcnQuY2FwdG5lbW8uaW4


## cisco

Remove your DNS blind spot (DNSCrypt protocol)

Warning: modifies your queries to include a copy of your network
address when forwarding them to a selection of companies and organizations.

Currently incompatible with DNS anonymization.

sdns://AQEAAAAAAAAADDIwOC42Ny4yMjAuMiC3NRFAIG8iXT4r2CLX_WkeocM8yNZmjQy-BL-rykP7eRsyLmRuc2NyeXB0LWNlcnQub3BlbmRucy5jb20


## cisco-doh

Remove your DNS blind spot (DoH protocol)

Warning: modifies your queries to include a copy of your network
address when forwarding them to a selection of companies and organizations.

sdns://AgAAAAAAAAAADDE0Ni4xMTIuNDEuMiBUDrXp92r0ml9Aq9cu3mXf2w_ugmc61w74ZllxOxR-Vw9kb2gub3BlbmRucy5jb20KL2Rucy1xdWVyeQ


## cisco-familyshield

Block websites not suitable for children (DNSCrypt protocol)

Warning: modifies your queries to include a copy of your network
address when forwarding them to a selection of companies and organizations.

Currently incompatible with DNS anonymization.

sdns://AQAAAAAAAAAADjIwOC42Ny4yMjAuMTIzILc1EUAgbyJdPivYItf9aR6hwzzI1maNDL4Ev6vKQ_t5GzIuZG5zY3J5cHQtY2VydC5vcGVuZG5zLmNvbQ


## cisco-familyshield-doh

Block websites not suitable for children (DoH protocol)

Warning: modifies your queries to include a copy of your network
address when forwarding them to a selection of companies and organizations.

sdns://AgAAAAAAAAAADDE0Ni4xMTIuNDEuMyBUDrXp92r0ml9Aq9cu3mXf2w_ugmc61w74ZllxOxR-Vxxkb2guZmFtaWx5c2hpZWxkLm9wZW5kbnMuY29tCi9kbnMtcXVlcnk


## cisco-ipv6

Cisco OpenDNS over IPv6 (DNSCrypt protocol)

Warning: modifies your queries to include a copy of your network
address when forwarding them to a selection of companies and organizations.

Currently incompatible with DNS anonymization.

sdns://AQEAAAAAAAAAD1syNjIwOjA6Y2NjOjoyXSC3NRFAIG8iXT4r2CLX_WkeocM8yNZmjQy-BL-rykP7eRsyLmRuc2NyeXB0LWNlcnQub3BlbmRucy5jb20


## cisco-ipv6-doh

Cisco OpenDNS over IPv6 (DoH protocol)

Warning: modifies your queries to include a copy of your network
address when forwarding them to a selection of companies and organizations.

sdns://AgAAAAAAAAAAEFsyNjIwOjExOTpmYzo6Ml0gVA616fdq9JpfQKvXLt5l39sP7oJnOtcO-GZZcTsUflcPZG9oLm9wZW5kbnMuY29tCi9kbnMtcXVlcnk


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


## cleanbrowsing-security

Block access to phishing, malware and malicious domains. It does not block adult content.
By https://cleanbrowsing.org/

sdns://AQMAAAAAAAAAEjE4NS4yMjguMTY4Ljk6ODQ0MyC8rDL61UNpFx8IMtYHUCfDIIzu8Ojpn5QY3HdgZdSPKRFjbGVhbmJyb3dzaW5nLm9yZw


## cloudflare

Cloudflare DNS (anycast) - aka 1.1.1.1 / 1.0.0.1

sdns://AgcAAAAAAAAABzEuMC4wLjEAEmRucy5jbG91ZGZsYXJlLmNvbQovZG5zLXF1ZXJ5


## cloudflare-ipv6

Cloudflare DNS over IPv6 (anycast)

sdns://AgcAAAAAAAAAGVsyNjA2OjQ3MDA6NDcwMDo6MTExMV06NTMAEmRucy5jbG91ZGZsYXJlLmNvbQovZG5zLXF1ZXJ5


## commons-host

DoH server by the Commons Host CDN

sdns://AgUAAAAAAAAAACA-GhoPbFPz6XpJLVcIS1uYBwWe4FerFQWHb9g_2j24OAxjb21tb25zLmhvc3QKL2Rucy1xdWVyeQ


## comodo-02

Comodo Dome Shield (anycast) - https://cdome.comodo.com/shield/

sdns://AQUAAAAAAAAACjguMjAuMjQ3LjIg0sJUqpYcHsoXmZb1X7yAHwg2xyN5q1J-zaiGG-Dgs7AoMi5kbnNjcnlwdC1jZXJ0LnNoaWVsZC0yLmRuc2J5Y29tb2RvLmNvbQ


## containerpi

Non-logging, non-filtering, DNSSEC validating server, EDNS Client Subnet enabled.
Multiple nodes in China Mainland, Japan and Germany.
Maintained by CPI-tech-Group

sdns://AgMAAAAAAAAADDQ1Ljc3LjE4MC4xMCBsA2QQ3lR1Nl9Ygfr8FdBIpL-doxmHECRx3T5NIXYYtxNkbnMuY29udGFpbmVycGkuY29tCi9kbnMtcXVlcnk


## containerpi-ipv6

Non-logging, non-filtering, DNSSEC validating server, EDNS Client Subnet enabled.
Multiple nodes in China Mainland, Japan and Germany.
Maintained by CPI-tech-Group

sdns://AgMAAAAAAAAALVsyMDAxOjE5ZjA6NzAwMTo1NTU0OjU0MDA6MDJmZjpmZTU3OjMwNzddOjQ0MyBsA2QQ3lR1Nl9Ygfr8FdBIpL-doxmHECRx3T5NIXYYtxNkbnMuY29udGFpbmVycGkuY29tCi9kbnMtcXVlcnk


## cs-ca

Canada DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADzE2Mi4yMjEuMjA3LjIyOCAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM


## cs-ca2

Secondary Canada DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADjE2Ny4xMTQuODQuMTMyIDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-ch

Switzerland DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAACzgxLjE3LjMxLjM0IDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-de

Germany DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADDg0LjE2LjI0MC40MyAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM


## cs-fi

Finland DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADjE4NS4xMTcuMTE4LjIwIDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-fr

France DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADTIxMi4xMjkuNDYuMzIgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz


## cs-it

Italy DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADjE4NS45NC4xOTMuMjM0IDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-lv

Latvia DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADzEwOS4yNDguMTQ5LjEzMyAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM


## cs-mo

Moldova DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADzE3OC4xNzUuMTM5LjIxMSAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM


## cs-nl2

Secondary Netherlands DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADTE4NS4xMDcuODAuODQgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz


## cs-ro

Romania DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADDUuMjU0Ljk2LjE5NSAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM


## cs-swe

Sweden DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADzEyOC4xMjcuMTA0LjEwOCAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM


## cs-usca

US - CA DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADDIzLjE5LjY3LjExNiAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM


## cs-usga

US - GA DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADTY0LjQyLjE4MS4yMjcgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz


## cs-usil

US - IL DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADjE3My4yMzQuNTYuMTE1IDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-usnc

US - NC DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADjE1NS4yNTQuMjkuMTEzIDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-usor

US - OR DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADTEwNC4yNTUuMTc1LjIgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz


## cs-ustx

US - TX DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADTIwOS41OC4xNDcuMzYgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz


## cs-uswa

US - WA DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADDY0LjEyMC41LjI1MSAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM


## cz.nic

CZ.NIC's open DNSSEC validating resolvers in Prague, Czech Republic.
CZ.NIC resolvers neither collect any personal data nor gather
information on pages where your computer sends personal data.
https://www.nic.cz/odvr/

sdns://AgcAAAAAAAAADDE4NS40My4xMzUuMSA_4zhjTgUQYz3kU8o1CxXOwzmz3Li6nyot0k0QqDj-6wtvZHZyLm5pYy5jegQvZG9o


## d0wn-is-ns2

Server provided by Martin 'd0wn' Albus

sdns://AQcAAAAAAAAADTkzLjk1LjIyNi4xNjUghGA0qcYwyjwErEqQFiXxeoeyrLlBgKxIHiwQ6M7eGm8cMi5kbnNjcnlwdC1jZXJ0LmlzMi5kMHduLmJpeg


## d0wn-tz-ns1

Server provided by Martin 'd0wn' Albus

sdns://AQcAAAAAAAAACzQxLjc5LjY5LjEzINYGFfvRRTuhTnaKPlxcs6wXRhMxRj2gr4z33wTaTXVtGzIuZG5zY3J5cHQtY2VydC50ei5kMHduLmJpeg


## d0wn-tz-ns1-ipv6

Server provided by Martin 'd0wn' Albus

sdns://AQcAAAAAAAAAGFsyYzBmOmZkYTg6NTo6MmVkMTpkMmVjXSDWBhX70UU7oU52ij5cXLOsF0YTMUY9oK-M998E2k11bRsyLmRuc2NyeXB0LWNlcnQudHouZDB3bi5iaXo


## dns.digitale-gesellschaft.ch

Public DoH resolver operated by the Digital Society (https://www.digitale-gesellschaft.ch).
Hosted in Zurich, Switzerland.

Non-logging, non-filtering, supports DNSSEC.

sdns://AgcAAAAAAAAAAKA-GhoPbFPz6XpJLVcIS1uYBwWe4FerFQWHb9g_2j24OKBoo-sB-l8CxNNfOhHQBMrwiyJL7qfXnFiMfxPIYTNgLqDvR4Wu5wydV1_nM4MG2T6nlhHl_tzvU2LdZsmLYLstvSAcVDa2UaK1QVwWz9ltGpcJ_ZyPJ-73XPlz2YL_5o5Y8BxkbnMuZGlnaXRhbGUtZ2VzZWxsc2NoYWZ0LmNoCi9kbnMtcXVlcnk


## dns.sb

DNSSEC-enabled DoH server by https://xtom.com/
Using Cloudflare as a frontend.

https://dns.sb

sdns://AgUAAAAAAAAAACA9pLcWNKQTwc7zSqltJaQBY01M82w7Ezx0KU5I5jcBKgpkb2guZG5zLnNiCi9kbnMtcXVlcnk


## dnscrypt-de-blahdns-ipv4

Blocks ad and Tracking, no Logging, DNSSEC, Hosted in Germany. By https://blahdns.com/

sdns://AQMAAAAAAAAAEzE1OS42OS4xOTguMTAxOjg0NDMgU4ToFEMUKT5W3RsUCh7xcq1HvboXmciVcpSVPQNOtccbMi5kbnNjcnlwdC1jZXJ0LmJsYWhkbnMuY29t


## dnscrypt-de-blahdns-ipv6

Blocks ad and Tracking, no Logging, DNSSEC, Hosted in Germany. By https://blahdns.com/

sdns://AQMAAAAAAAAAHFsyYTAxOjRmODoxYzFjOjZiNGI6OjFdOjg0NDMgU4ToFEMUKT5W3RsUCh7xcq1HvboXmciVcpSVPQNOtccbMi5kbnNjcnlwdC1jZXJ0LmJsYWhkbnMuY29t


## dnscrypt-fi-blahdns-ipv4

Blocks ad and Tracking, no Logging, DNSSEC, Hosted in Finland. By https://blahdns.com/

sdns://AQMAAAAAAAAAEzk1LjIxNi4yMTIuMTc3Ojg0NDMgU4ToFEMUKT5W3RsUCh7xcq1HvboXmciVcpSVPQNOtccbMi5kbnNjcnlwdC1jZXJ0LmJsYWhkbnMuY29t


## dnscrypt-fi-blahdns-ipv6

Blocks ad and Tracking, no Logging, DNSSEC, Hosted in Finland. By https://blahdns.com/

sdns://AQMAAAAAAAAAHFsyYTAxOjRmOTpjMDEwOjQzY2U6OjFdOjg0NDMgU4ToFEMUKT5W3RsUCh7xcq1HvboXmciVcpSVPQNOtccbMi5kbnNjcnlwdC1jZXJ0LmJsYWhkbnMuY29t


## dnscrypt-jp-blahdns-ipv4

Blocks ad and Tracking, no Logging, DNSSEC, Hosted in Japan. By https://blahdns.com/

sdns://AQMAAAAAAAAAEDQ1LjMyLjU1Ljk0Ojg0NDMgU4ToFEMUKT5W3RsUCh7xcq1HvboXmciVcpSVPQNOtccbMi5kbnNjcnlwdC1jZXJ0LmJsYWhkbnMuY29t


## dnscrypt-jp-blahdns-ipv6

Blocks ad and Tracking, no Logging, DNSSEC, Hosted in Japan. By https://blahdns.com/

sdns://AQMAAAAAAAAALlsyMDAxOjE5ZjA6NzAwMTozMjU5OjU0MDA6MDJmZjpmZTcxOjBiYzldOjg0NDMgU4ToFEMUKT5W3RsUCh7xcq1HvboXmciVcpSVPQNOtccbMi5kbnNjcnlwdC1jZXJ0LmJsYWhkbnMuY29t


## dnscrypt.ca-1

Free, Canadian, uncensored, no-logs, encrypted, and DNSSEC validated
DNS service for your pleasure.

sdns://AQcAAAAAAAAAEzE2Ny4xMTQuMjIwLjEyNTo0NDMgGlOjyVB4nL3RCxkzpGibbIRqQPG3PRdSrsrJgp7LfOIdMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeXB0LmNhLTE


## dnscrypt.ca-1-doh

Free, Canadian, uncensored, no-logs, encrypted, and DNSSEC validated DNS service for your pleasure.

sdns://AgcAAAAAAAAADzE2Ny4xMTQuMjIwLjEyNSA-GhoPbFPz6XpJLVcIS1uYBwWe4FerFQWHb9g_2j24OBRkbnMxLmRuc2NyeXB0LmNhOjQ1MwovZG5zLXF1ZXJ5


## dnscrypt.ca-1-ipv6

Free, Canadian, uncensored, no-logs, encrypted, and DNSSEC validated
DNS service for your pleasure.

sdns://AQcAAAAAAAAAKlsyNjA3OjUzMDA6NjE6OTVmOjcyODM6MTFkOTowZjg2OmU2ODldOjQ0MyAg2RnU7DxQCM3HClHqK0_L_oqZPGQx0OxoHV6nOEWrgCIyLmRuc2NyeXB0LWNlcnQuZG5zY3J5cHQuY2EtMS1pcHY2


## dnscrypt.ca-1-ipv6-doh

Free, Canadian, uncensored, no-logs, encrypted, and DNSSEC validated DNS service for your pleasure.

sdns://AgcAAAAAAAAAJlsyNjA3OjUzMDA6NjE6OTVmOjcyODM6MTFkOTowZjg2OmU2ODldID4aGg9sU_PpekktVwhLW5gHBZ7gV6sVBYdv2D_aPbg4FGRuczEuZG5zY3J5cHQuY2E6NDUzCi9kbnMtcXVlcnk


## dnscrypt.ca-2

Free, Canadian, uncensored, no-logs, encrypted, and DNSSEC validated
DNS service for your pleasure.

sdns://AQcAAAAAAAAAETE0OS41Ni4yMjguNDU6NDQzIAEIVKs7Vqfu-dORWP72ggv_k6I1fIkWCNueFdO74BGFHTIuZG5zY3J5cHQtY2VydC5kbnNjcnlwdC5jYS0y


## dnscrypt.ca-2-doh

Free, Canadian, uncensored, no-logs, encrypted, and DNSSEC validated DNS service for your pleasure.

sdns://AgcAAAAAAAAADTE0OS41Ni4yMjguNDUgPhoaD2xT8-l6SS1XCEtbmAcFnuBXqxUFh2_YP9o9uDgUZG5zMi5kbnNjcnlwdC5jYTo0NTMKL2Rucy1xdWVyeQ


## dnscrypt.ca-2-ipv6

Free, Canadian, uncensored, no-logs, encrypted, and DNSSEC validated
DNS service for your pleasure.

sdns://AQcAAAAAAAAAHFsyNjA3OjUzMDA6MTIwOmI5Yjo6MjAwXTo0NDMgmHxwqJfb2hUaNK1LVeqADvVhzASq1cV90sPYYfwX9CkiMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeXB0LmNhLTItaXB2Ng


## dnscrypt.ca-2-ipv6-doh

Free, Canadian, uncensored, no-logs, encrypted, and DNSSEC validated DNS service for your pleasure.

sdns://AgcAAAAAAAAAGFsyNjA3OjUzMDA6MTIwOmI5Yjo6MjAwXSA-GhoPbFPz6XpJLVcIS1uYBwWe4FerFQWHb9g_2j24OBRkbnMyLmRuc2NyeXB0LmNhOjQ1MwovZG5zLXF1ZXJ5


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


## dnscrypt.one

non-logging, non-filtering and non-censoring, DNSSEC. Germany (Nuremberg), https://dnscrypt.one

sdns://AQcAAAAAAAAAEjE0NC45MS4xMDYuMjI3OjQ0MyBCc4nfqRWGkg4g1vk6yPqtIdsD5Ub4y3sDEKYnQ2vN3BwyLmRuc2NyeXB0LWNlcnQuZG5zY3J5cHQub25l


## dnscrypt.uk-ipv4

DNSCrypt v2, no logs, uncensored, DNSSEC. Hosted in London UK on Digital Ocean
https://www.dnscrypt.uk

sdns://AQcAAAAAAAAAEjEzOS41OS4yMDAuMTE2OjQ0MyBxkKwJmxhDAMvj-aF2TcFOK16cSI5EpMxBJG3Ze0lRvBsyLmRuc2NyeXB0LWNlcnQuZG5zY3J5cHQudWs


## dnscrypt.uk-ipv6

DNSCrypt v2, no logs, uncensored, DNSSEC. Hosted in London UK on Digital Ocean
https://www.dnscrypt.uk

sdns://AQcAAAAAAAAAHlsyYTAzOmIwYzA6MTplMDo6MmUzOmUwMDFdOjQ0MyBxkKwJmxhDAMvj-aF2TcFOK16cSI5EpMxBJG3Ze0lRvBsyLmRuc2NyeXB0LWNlcnQuZG5zY3J5cHQudWs


## dnsforfamily

Block adult websites, gambling websites and advertisements. No DNS queries are logged. As of March 2019 2.1million websites are blocked and new websites are added to blacklist daily.

Provided by: https://dnsforfamily.com

sdns://AQIAAAAAAAAADDc4LjQ3LjY0LjE2MSATJeLOABXNSYcSJIoqR5_iUYz87Y4OecMLB84aEAKPrRBkbnNmb3JmYW1pbHkuY29t


## dnsforfamily-v6

Block adult websites, gambling websites and advertisements. No DNS queries are logged. As of March 2019 2.1million websites are blocked and new websites are added to blacklist daily.

Provided by: https://dnsforfamily.com

sdns://AQIAAAAAAAAAF1syYTAxOjRmODoxYzE3OjRkZjg6OjFdIGN4CrSY4fb2hK8voFJL3GKiM7xQNwkKGH4b0k7LmMPxEGRuc2ZvcmZhbWlseS5jb20


## dnsforge.de

Public DoH resolver running with Pihole for Adblocking (https://dnsforge.de).

Non-logging, AD-filtering, supports DNSSEC. Hosted in Germany.

sdns://AgMAAAAAAAAADDE3Ni45LjkzLjE5OCA-GhoPbFPz6XpJLVcIS1uYBwWe4FerFQWHb9g_2j24OAtkbnNmb3JnZS5kZQovZG5zLXF1ZXJ5


## doh-centraleu-pi-dns

Public Pi-hole DNS with DoH support. Blocks ads and trackers. No persistent logs, DNSSEC, Hosted in Nuremberg, Germany. By https://pi-dns.com/
Server status can be seen at: https://pi-dns.com/

sdns://AgMAAAAAAAAADTg4LjE5OC45MS4xODcgPhoaD2xT8-l6SS1XCEtbmAcFnuBXqxUFh2_YP9o9uDgYZG9oLmNlbnRyYWxldS5waS1kbnMuY29tCi9kbnMtcXVlcnk


## doh-centraleu-pi-dns-ipv6

Public Pi-hole DNS with DoH support. Blocks ads and trackers. No persistent logs, DNSSEC, Hosted in Nuremberg, Germany. By https://pi-dns.com/
Server status can be seen at: https://pi-dns.com/

sdns://AgMAAAAAAAAAF1syYTAxOjRmODoxYzBjOjgyMzM6OjFdID4aGg9sU_PpekktVwhLW5gHBZ7gV6sVBYdv2D_aPbg4GGRvaC5jZW50cmFsZXUucGktZG5zLmNvbQovZG5zLXF1ZXJ5


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

sdns://AgcAAAAAAAAADDEwNC4yOC4wLjEwNiA9pLcWNKQTwc7zSqltJaQBY01M82w7Ezx0KU5I5jcBKg1kb2guY3J5cHRvLnN4Ci9kbnMtcXVlcnk


## doh-crypto-sx-ipv6

DNS-over-HTTPS server accessible over IPv6. Anycast, no logs, no censorship, DNSSEC.
Backend hosted by Scaleway, globally cached via Cloudflare.
Maintained by Frank Denis.

sdns://AgcAAAAAAAAAF1syNjA2OjQ3MDA6MzA6OjY4MWM6NmFdIB3YUIqMVlkEOXx4ZFHgj1dhPPPIpyP283n9TDhYtvOfEmRvaC1pcHY2LmNyeXB0by5zeAovZG5zLXF1ZXJ5


## doh-de-blahdns

Blocks ad and Tracking, no Logging, DNSSEC, Hosted in Germany. By https://blahdns.com/

sdns://AgMAAAAAAAAADjE1OS42OS4xOTguMTAxABJkb2gtZGUuYmxhaGRucy5jb20KL2Rucy1xdWVyeQ


## doh-de-blahdns-v6

Blocks ad and Tracking, no Logging, DNSSEC, Hosted in Germany. By https://blahdns.com/

sdns://AgMAAAAAAAAAF1syYTAxOjRmODoxYzFjOjZiNGI6OjFdABJkb2gtZGUuYmxhaGRucy5jb20KL2Rucy1xdWVyeQ


## doh-eastus-pi-dns

Public Pi-hole DNS with DoH support. Blocks ads and trackers. No persistent logs, DNSSEC, Hosted in New York, USA. By https://pi-dns.com/
Server status can be seen at: https://pi-dns.com/

sdns://AgMAAAAAAAAADjE4NS4yMTMuMjYuMTg3ID4aGg9sU_PpekktVwhLW5gHBZ7gV6sVBYdv2D_aPbg4FWRvaC5lYXN0dXMucGktZG5zLmNvbQovZG5zLXF1ZXJ5


## doh-eastus-pi-dns-ipv6

Public Pi-hole DNS with DoH support. Blocks ads and trackers. No persistent logs, DNSSEC, Hosted in New York, USA. By https://pi-dns.com/
Server status can be seen at: https://pi-dns.com/

sdns://AgMAAAAAAAAAFlsyYTBkOjU2MDA6MzM6Mzo6YWJjZF0gPhoaD2xT8-l6SS1XCEtbmAcFnuBXqxUFh2_YP9o9uDgVZG9oLmVhc3R1cy5waS1kbnMuY29tCi9kbnMtcXVlcnk


## doh-fi-blahdns

Blocks ad and Tracking, no Logging, DNSSEC, Hosted in Finland. By https://blahdns.com/

sdns://AQMAAAAAAAAAEzk1LjIxNi4yMTIuMTc3Ojg0NDMgU4ToFEMUKT5W3RsUCh7xcq1HvboXmciVcpSVPQNOtccbMi5kbnNjcnlwdC1jZXJ0LmJsYWhkbnMuY29t


## doh-fi-blahdns-v6

Blocks ad and Tracking, no Logging, DNSSEC, Hosted in Finland. By https://blahdns.com/

sdns://AgMAAAAAAAAAF1syYTAxOjRmOTpjMDEwOjQzY2U6OjFdABJkb2gtZmkuYmxhaGRucy5jb20KL2Rucy1xdWVyeQ


## doh-fi-snopyta

Snopyta non-logging DoH Server in Finland
run by Noah Seefried - https://snopyta.org

sdns://AgcAAAAAAAAADjk1LjIxNi4yMjkuMTUzID4aGg9sU_PpekktVwhLW5gHBZ7gV6sVBYdv2D_aPbg4FmZpLmRvaC5kbnMuc25vcHl0YS5vcmcKL2Rucy1xdWVyeQ


## doh-fi-snopyta-ipv6

Snopyta non-logging DoH Server in Finland - IPv6 only
run by Noah Seefried - https://snopyta.org

sdns://AgcAAAAAAAAAFlsyYTAxOjRmOToyYToxOTE5OjoyMV0gPhoaD2xT8-l6SS1XCEtbmAcFnuBXqxUFh2_YP9o9uDgWZmkuZG9oLmRucy5zbm9weXRhLm9yZwovZG5zLXF1ZXJ5


## doh-ibksturm

doh-server (nginx - dnsproxy - unbound backend), DNSSEC / Non-Logged / Uncensored, OpenNIC and Root DNS-Zone Copy
Hosted in Switzerland by ibksturm, aka Andreas Ziegler

sdns://AgcAAAAAAAAAACA-GhoPbFPz6XpJLVcIS1uYBwWe4FerFQWHb9g_2j24OBRpYmtzdHVybS5zeW5vbG9neS5tZQovZG5zLXF1ZXJ5


## doh-jp-blahdns

Blocks ad and Tracking, no Logging, DNSSEC, Hosted in Japan. By https://blahdns.com/

sdns://AgMAAAAAAAAACzQ1LjMyLjU1Ljk0ABJkb2gtanAuYmxhaGRucy5jb20KL2Rucy1xdWVyeQ


## doh-jp-blahdns-v6

Blocks ad and Tracking, no Logging, DNSSEC, Hosted in Japan. By https://blahdns.com/

sdns://AgMAAAAAAAAAKVsyMDAxOjE5ZjA6NzAwMTozMjU5OjU0MDA6MDJmZjpmZTcxOjBiYzldABJkb2gtanAuYmxhaGRucy5jb20KL2Rucy1xdWVyeQ


## doh-northeu-pi-dns

Public Pi-hole DNS with DoH support. Blocks ads and trackers. No persistent logs, DNSSEC, Hosted in Helsinki, Finland. By https://pi-dns.com/
Server status can be seen at: https://pi-dns.com/

sdns://AgMAAAAAAAAADjk1LjIxNi4xODEuMjI4ID4aGg9sU_PpekktVwhLW5gHBZ7gV6sVBYdv2D_aPbg4FmRvaC5ub3J0aGV1LnBpLWRucy5jb20KL2Rucy1xdWVyeQ


## doh-northeu-pi-dns-ipv6

Public Pi-hole DNS with DoH support. Blocks ads and trackers. No persistent logs, DNSSEC, Hosted in Helsinki, Finland. By https://pi-dns.com/
Server status can be seen at: https://pi-dns.com/

sdns://AgMAAAAAAAAAF1syYTAxOjRmOTpjMDFmOjQ6OmFiY2RdID4aGg9sU_PpekktVwhLW5gHBZ7gV6sVBYdv2D_aPbg4FmRvaC5ub3J0aGV1LnBpLWRucy5jb20KL2Rucy1xdWVyeQ


## doh-westus-pi-dns

Public Pi-hole DNS with DoH support. Blocks ads and trackers. No persistent logs, DNSSEC, Hosted in Los Angeles, USA. By https://pi-dns.com/
Server status can be seen at: https://pi-dns.com/

sdns://AgMAAAAAAAAADTQ1LjY3LjIxOS4yMDggPhoaD2xT8-l6SS1XCEtbmAcFnuBXqxUFh2_YP9o9uDgVZG9oLndlc3R1cy5waS1kbnMuY29tCi9kbnMtcXVlcnk


## doh-westus-pi-dns-ipv6

Public Pi-hole DNS with DoH support. Blocks ads and trackers. No persistent logs, DNSSEC, Hosted in Los Angeles, USA. By https://pi-dns.com/
Server status can be seen at: https://pi-dns.com/

sdns://AgMAAAAAAAAAGFsyYTA0OmJkYzc6MTAwOjcwOjphYmNkXSA-GhoPbFPz6XpJLVcIS1uYBwWe4FerFQWHb9g_2j24OBVkb2gud2VzdHVzLnBpLWRucy5jb20KL2Rucy1xdWVyeQ


## doh.appliedprivacy.net

Public DoH resolver operated by the Foundation for Applied Privacy (https://appliedprivacy.net).
Hosted in Vienna, Austria.

Non-logging, non-filtering, supports DNSSEC.

sdns://AgcAAAAAAAAAAKA-GhoPbFPz6XpJLVcIS1uYBwWe4FerFQWHb9g_2j24OKBoo-sB-l8CxNNfOhHQBMrwiyJL7qfXnFiMfxPIYTNgLqDvR4Wu5wydV1_nM4MG2T6nlhHl_tzvU2LdZsmLYLstvSAcVDa2UaK1QVwWz9ltGpcJ_ZyPJ-73XPlz2YL_5o5Y8BZkb2guYXBwbGllZHByaXZhY3kubmV0Bi9xdWVyeQ


## doh.ffmuc.net

An open DoH resolver operated by Freifunk Munich with nodes in DE.
https://ffmuc.net/

sdns://AgcAAAAAAAAADDE5NS4zMC45NC4yOAANZG9oLmZmbXVjLm5ldAovZG5zLXF1ZXJ5


## doh.ffmuc.net-v6

An open DoH resolver operated by Freifunk Munich with nodes in DE.
https://ffmuc.net/

sdns://AgcAAAAAAAAAFVsyMDAxOjYwODphMDE6OjNdOjQ0MyA-GhoPbFPz6XpJLVcIS1uYBwWe4FerFQWHb9g_2j24OA1kb2guZmZtdWMubmV0Ci9kbnMtcXVlcnk


## doh.li

Public DoH server in London, UK.

Warning: provides EDNS Client Subnet information to upstream servers.

sdns://AgUAAAAAAAAAACA-GhoPbFPz6XpJLVcIS1uYBwWe4FerFQWHb9g_2j24OAZkb2gubGkKL2Rucy1xdWVyeQ


## doh.tiarap.org

Non-Logging DNS-over-HTTPS server, cached via Cloudflare.
Filters out ads, trackers and malware, NO ECS, supports DNSSEC.

sdns://AgMAAAAAAAAADDEwNC4yOC4yOC4zNCBPtWwTIp4-T40ZbjCdyCfeStS1-WkKW8w_WWEQubJpyQ5kb2gudGlhcmFwLm9yZwovZG5zLXF1ZXJ5


## doh.tiarap.org-ipv6

Non-Logging DNS-over-HTTPS server (IPv6), cached via Cloudflare.
Filters out ads, trackers and malware, NO ECS, supports DNSSEC.

sdns://AgMAAAAAAAAAGVsyNjA2OjQ3MDA6MzA6OjY4MWM6MWQyMl0gT7VsEyKePk-NGW4wncgn3krUtflpClvMP1lhELmyackOZG9oLnRpYXJhcC5vcmcKL2Rucy1xdWVyeQ


## ev-to

Non-logging, uncensored DNS resolver provided by evilvibes.com
Location: Toronto, Canada

sdns://AQcAAAAAAAAADTIzLjExMS42OS4yMzggdOCmaDjs4MgSiTXowm5ae9FetLivEu9dl9QbbruylMcdMi5kbnNjcnlwdC1jZXJ0LmV2aWx2aWJlcy5jb20


## ev-va

Non-logging, uncensored DNS resolver provided by evilvibes.com
Location: Vancouver, Canada

sdns://AQcAAAAAAAAADjEwNC4zNi4xNDkuMTc3INRzo8srd84NzJJs_OGLkPMBC2jBNVH_U4US4ca7DXRXHTIuZG5zY3J5cHQtY2VydC5ldmlsdmliZXMuY29t


## faelix

An open DoH resolver operated by https://faelix.net/ with nodes in UK and CH.

sdns://AgcAAAAAAAAADDE5NS4zMC45NC4yOCA-GhoPbFPz6XpJLVcIS1uYBwWe4FerFQWHb9g_2j24OA1kb2guZmZtdWMubmV0Ci9kbnMtcXVlcnk


## ffmuc.net

An open DNSCrypt resolver operated by Freifunk Munich with nodes in DE.
https://ffmuc.net/

sdns://AQcAAAAAAAAAETE5NS4zMC45NC4yODo4NDQzIAfQevHP3F2Zdp0_AmaQpwRJZcJ0J2x5HK71rvO5DEb6GTIuZG5zY3J5cHQtY2VydC5mZm11Yy5uZXQ


## ffmuc.net-v6

An open DNSCrypt resolver operated by Freifunk Munich with nodes in DE.
https://ffmuc.net/

sdns://AQcAAAAAAAAAFlsyMDAxOjYwODphMDE6OjNdOjg0NDMgB9B68c_cXZl2nT8CZpCnBEllwnQnbHkcrvWu87kMRvoZMi5kbnNjcnlwdC1jZXJ0LmZmbXVjLm5ldA


## freetsa.org

Non-logged/Uncensored provided by freetsa.org

sdns://AQcAAAAAAAAAEzIwNS4xODUuMTE2LjExNjo1NTMg2P-7QuAxvnp5cwtFVo1Jak6Ky1mqg2b9arkeJyp9FuQbMi5kbnNjcnlwdC1jZXJ0LmZyZWV0c2Eub3Jn


## geekdns-doh-east

GeekDNS in eastern China (Shanghai).

GeekDNS is a non-logging public DNS service located in mainland China.
Queries are cached locally, and, for some domains, resolved by servers located
in Taiwan.

https://dns.233py.com/

sdns://AgcAAAAAAAAADTQ3LjEwMS4xMzYuMzcgot4zSQxHbTVuLbxzfCd5aSJJUmtlq4-6mjQoBIHIvfwOZWRucy4yMzNweS5jb20KL2Rucy1xdWVyeQ


## geekdns-doh-north

GeekDNS in northern China (Beijing).

GeekDNS is a non-logging public DNS service located in mainland China.
Queries are cached locally, and, for some domains, resolved by servers located
in Taiwan.

https://dns.233py.com/

sdns://AgcAAAAAAAAADzExNC4xMTUuMjQwLjE3NSCi3jNJDEdtNW4tvHN8J3lpIklSa2Wrj7qaNCgEgci9_A5uZG5zLjIzM3B5LmNvbQovZG5zLXF1ZXJ5


## geekdns-doh-south

GeekDNS in southern China (Guangzhou).

GeekDNS is a non-logging public DNS service located in mainland China.
Queries are cached locally, and, for some domains, resolved by servers located
in Taiwan.

https://dns.233py.com/

sdns://AgcAAAAAAAAADTExOS4yOS4xMDcuODUgot4zSQxHbTVuLbxzfCd5aSJJUmtlq4-6mjQoBIHIvfwOc2Rucy4yMzNweS5jb20KL2Rucy1xdWVyeQ


## geekdns-doh-west

GeekDNS in western China (Chongqing).

GeekDNS is a non-logging public DNS service located in mainland China.
Queries are cached locally, and, for some domains, resolved by servers located
in Taiwan.

https://dns.233py.com/

sdns://AgMAAAAAAAAADjExOC4yNC4yMDguMTk3ID4aGg9sU_PpekktVwhLW5gHBZ7gV6sVBYdv2D_aPbg4DndkbnMuMjMzcHkuY29tCi9kbnMtcXVlcnk


## geekdns-hk

GeekDNS in Hong Kong.

GeekDNS is a non-logging public DNS service located in Hong Kong..
Queries are cached locally, and, for some domains, resolved by servers located
in Taiwan.

https://dns.233py.com/

sdns://AQcAAAAAAAAAETE1MC4xMDkuNzQuMTY0OjIyIPcEBXJOU2jQys6br08P8yyn132SuDixQ8Oek3lhoRQoGTIuZG5zY3J5cHQtY2VydC4yMzNweS5jb20


## geekdns-north

GeekDNS in northern China (Beijing).

GeekDNS is a non-logging public DNS service located in mainland China.
Queries are cached locally, and, for some domains, resolved by servers located
in Taiwan.

https://dns.233py.com/

sdns://AQMAAAAAAAAAEjExNC4xMTUuMjQwLjE3NToyMiCLntDYEK0AwismFtMCM0YMkflJGNZiJnINFtDLcCLLwBkyLmRuc2NyeXB0LWNlcnQuMjMzcHkuY29t


## geekdns-south

GeekDNS in southern China (Guangzhou).

GeekDNS is a non-logging public DNS service located in mainland China.
Queries are cached locally, and, for some domains, resolved by servers located
in Taiwan.

https://dns.233py.com/

sdns://AQMAAAAAAAAAEDExOS4yOS4xMDcuODU6MjIgCsUpkBu6ujbEDZ8PQI2wa3DCkDzxLlczLOwTyQTQM70ZMi5kbnNjcnlwdC1jZXJ0LjIzM3B5LmNvbQ


## geekdns-west

GeekDNS in western China (Chongqing).

GeekDNS is a non-logging public DNS service located in mainland China.
Queries are cached locally, and, for some domains, resolved by servers located
in Taiwan.

https://dns.233py.com/

sdns://AQMAAAAAAAAAETExOC4yNC4yMDguMTk3OjIyIL66dzE0aNGlvYsF3RnukLk4AI3lQqSxeEo6PxHme-qZGTIuZG5zY3J5cHQtY2VydC4yMzNweS5jb20


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


## ibksturm

dnscrypt-server (nginx - encrypted-dns - unbound backend), DNSSEC / Non-Logged / Uncensored, OpenNIC and Root DNS-Zone Copy
Hosted in Switzerland by ibksturm, aka Andreas Ziegler

sdns://AQcAAAAAAAAADzgzLjc3Ljg1Ljc6ODQ0MyDBz1dQALBbwmxiH17PmqJWCs6_AH6-yzp_9LIN4LQ57hgyLmRuc2NyeXB0LWNlcnQuaWJrc3R1cm0


## ibksturm-ipv6

dnscrypt-server (nginx - encrypted-dns - unbound backend), DNSSEC / Non-Logged / Uncensored, OpenNIC and Root DNS-Zone Copy
Hosted in Switzerland by ibksturm, aka Andreas Ziegler

sdns://AQcAAAAAAAAALlsyYTAyOjEyMDU6MzRkNTo1MDcwOmIyNmU6YmZmZjpmZTFkOmUxOWJdOjg0NDMgwc9XUACwW8JsYh9ez5qiVgrOvwB-vss6f_SyDeC0Oe4YMi5kbnNjcnlwdC1jZXJ0Lmlia3N0dXJt


## id-gmail

Non-Logging DNSCrypt server located in Singapore.
Filters out ads, trackers and malware, supports DNSSEC, provided by id-gmail.

sdns://AQMAAAAAAAAADjE3NC4xMzguMjEuMTI4IO-WgGbo2ZTwZdg-3dMa7u31bYZXRj5KykfN1_6Xw9T2HDIuZG5zY3J5cHQtY2VydC5kbnMudGlhci5hcHA


## id-gmail-doh

Non-Logging DNS-over-HTTPS server located in Singapore.
Filters out ads, trackers and malware, supports DNSSEC, provided by id-gmail.

sdns://AgMAAAAAAAAADjE3NC4xMzguMjkuMTc1ID4aGg9sU_PpekktVwhLW5gHBZ7gV6sVBYdv2D_aPbg4DGRvaC50aWFyLmFwcAovZG5zLXF1ZXJ5


## id-gmail-doh-ipv6

Non-Logging DNS-over-HTTPS (IPv6) server located in Singapore.
Filters out ads, trackers and malware, supports DNSSEC, provided by id-gmail.

sdns://AgMAAAAAAAAAG1syNDAwOjYxODA6MDpkMDo6NWY3Mzo0MDAxXSA-GhoPbFPz6XpJLVcIS1uYBwWe4FerFQWHb9g_2j24OAxkb2gudGlhci5hcHAKL2Rucy1xdWVyeQ


## id-gmail-ipv6

Non-Logging DNSCrypt (IPv6) server located in Singapore.
Filters out ads, trackers and malware, supports DNSSEC, provided by id-gmail.

sdns://AQMAAAAAAAAAG1syNDAwOjYxODA6MDpkMDo6NWY2ZTo0MDAxXSDvloBm6NmU8GXYPt3TGu7t9W2GV0Y-SspHzdf-l8PU9hwyLmRuc2NyeXB0LWNlcnQuZG5zLnRpYXIuYXBw


## iij

DoH server operated by Internet Initiative Japan in Tokyo.
https://www.iij.ad.jp/

sdns://AgUAAAAAAAAACjEwMy4yLjU3LjUgs2lfGAQCPrV0DPQqOkPci0Jei0GhMK_ne-QHwPbfn4oRcHVibGljLmRucy5paWouanAKL2Rucy1xdWVyeQ


## jp.tiar.app

Non-Logging, Non-Filtering DNSCrypt server in Japan.
No ECS, Support DNSSEC

sdns://AQcAAAAAAAAAEjE3Mi4xMDQuOTMuODA6MTQ0MyAyuHY-8b9lNqHeahPAzW9IoXnjiLaZpTeNbVs8TN9UUxsyLmRuc2NyeXB0LWNlcnQuanAudGlhci5hcHA


## jp.tiar.app-doh

Non-Logging, Non-Filtering DNS-over-HTTPS server in Japan.
No ECS, Support DNSSEC

sdns://AgcAAAAAAAAADTE3Mi4xMDQuOTMuODAgPhoaD2xT8-l6SS1XCEtbmAcFnuBXqxUFh2_YP9o9uDgLanAudGlhci5hcHAKL2Rucy1xdWVyeQ


## jp.tiar.app-doh-ipv6

Non-Logging, Non-Filtering DNS-over-HTTPS (IPv6) server in Japan.
No ECS, Support DNSSEC

sdns://AgcAAAAAAAAAIFsyNDAwOjg5MDI6OmYwM2M6OTFmZjpmZWRhOmM1MTRdID4aGg9sU_PpekktVwhLW5gHBZ7gV6sVBYdv2D_aPbg4C2pwLnRpYXIuYXBwCi9kbnMtcXVlcnk


## jp.tiar.app-ipv6

Non-Logging, Non-Filtering DNSCrypt (IPv6) server in Japan.
No ECS, Support DNSSEC

sdns://AQcAAAAAAAAAJVsyNDAwOjg5MDI6OmYwM2M6OTFmZjpmZWRhOmM1MTRdOjE0NDMgMrh2PvG_ZTah3moTwM1vSKF544i2maU3jW1bPEzfVFMbMi5kbnNjcnlwdC1jZXJ0LmpwLnRpYXIuYXBw


## jp.tiarap.org

DNS-over-HTTPS Server. Non-Logging, Non-Filtering, No ECS, Support DNSSEC.
Cached via Cloudflare.

sdns://AgcAAAAAAAAADDEwNC4yOC4yOC4zNCBPtWwTIp4-T40ZbjCdyCfeStS1-WkKW8w_WWEQubJpyQ1qcC50aWFyYXAub3JnCi9kbnMtcXVlcnk


## jp.tiarap.org-ipv6

DNS-over-HTTPS Server (IPv6). Non-Logging, Non-Filtering, No ECS, Support DNSSEC.
Cached via Cloudflare.

sdns://AgcAAAAAAAAAGVsyNjA2OjQ3MDA6MzA6OjY4MWM6MWQyMl0gT7VsEyKePk-NGW4wncgn3krUtflpClvMP1lhELmyackNanAudGlhcmFwLm9yZwovZG5zLXF1ZXJ5


## lelux.fi

DoH server in France. No logging, no filtering, DNSSEC support.
https://lelux.fi/resolver/

sdns://AgcAAAAAAAAADTUxLjE1OC4xNDcuNTAgPhoaD2xT8-l6SS1XCEtbmAcFnuBXqxUFh2_YP9o9uDgUcmVzb2x2ZXItZXUubGVsdXguZmkKL2Rucy1xdWVyeQ


## libredns

DoH server in Germany. No logging, but no DNS padding and no DNSSEC support.
https://libredns.gr/

sdns://AgYAAAAAAAAAACA-GhoPbFPz6XpJLVcIS1uYBwWe4FerFQWHb9g_2j24OA9kb2gubGlicmVkbnMuZ3IKL2Rucy1xdWVyeQ


## libredns-noads

DoH server in Germany. No logging, but no DNS padding and no DNSSEC support.
no ads version, uses StevenBlack's host list: https://github.com/StevenBlack/hosts

sdns://AgIAAAAAAAAAACA-GhoPbFPz6XpJLVcIS1uYBwWe4FerFQWHb9g_2j24OA9kb2gubGlicmVkbnMuZ3IEL2Fkcw


## mrkaran

Public DNSCrypt Server hosted in Bengaluru, India on a tiny Digital
Ocean droplet. It supports DNSSEC validation, the DNSCrypt protocol
and has caching enabled for faster responses.

https://mrkaran.dev/dnscrypt/

sdns://AQcAAAAAAAAAETEzOS41OS41NS4xMzo0MzQzIMI_cHfgQzHFYUiS8m2ghR8Ij9nb88EQZXAYDlOhBGhmHzIuZG5zY3J5cHQtY2VydC5kbnMubXJrYXJhbi5kZXY


## nextdns

NextDNS is a cloud-based private DNS service that gives you full control
over what is allowed and what is blocked on the Internet.

DNSSEC, Anycast, Non-logging, NoFilters

https://www.nextdns.io/

sdns://AgcAAAAAAAAACjQ1LjkwLjI4LjAgPhoaD2xT8-l6SS1XCEtbmAcFnuBXqxUFh2_YP9o9uDgOZG5zLm5leHRkbnMuaW8PL2Ruc2NyeXB0LXByb3h5


## opennic-R4SAS

DNSSEC - OpenNIC - Non-logging - Uncensored - hosted on ovh.com
Location: Gravelines, France.
Maintained by R4SAS.

sdns://AQcAAAAAAAAAETE1MS44MC4yMjIuNzk6NDQzIKnWMjpPJYAJJhl1FQLOIx4fdtned2yHxruyig7_2w5OIDIuZG5zY3J5cHQtY2VydC5vcGVubmljLmkycGQueHl6


## opennic-R4SAS-ipv6

DNSSEC - OpenNIC - Non-logging - Uncensored - hosted on ovh.com
Location: Gravelines, France.
Maintained by R4SAS.

sdns://AQcAAAAAAAAAG1syMDAxOjQ3MDoxZjE1OmI4MDo6NTNdOjQ0MyCp1jI6TyWACSYZdRUCziMeH3bZ3ndsh8a7sooO_9sOTiAyLmRuc2NyeXB0LWNlcnQub3Blbm5pYy5pMnBkLnh5eg


## opennic-bongobow

OpenNIC • Non-logging • No DNSSEC
Location: Munich, Germany

sdns://AQYAAAAAAAAAETUuMTg5LjE3MC4xOTY6NDY1IFQ1LFVAO4Luk8QH_cI0RJcNmlbvIr_P-eyQnM0yJDJrKDIuZG5zY3J5cHQtY2VydC5uczE2LmRlLmRucy5vcGVubmljLmdsdWU


## opennic-eleix

OpenNIC • 24-hour Logs • No DNSSEC
Location: Montreal, QC, Canada

sdns://AQQAAAAAAAAAEzE5OC4xMDAuMTQ4LjIyNDo0NDMg-I9_7aXj72-YfDvHJRG_Q3zhfU3m33yuKPcfqYpr_VYkMi5kbnNjcnlwdC1jZXJ0LmpvbHRlb24uYm9vdGhsYWJzLm1l


## opennic-luggs

Public DNS server in Canada operated by Luggs

sdns://AQYAAAAAAAAADTE0Mi40LjIwNC4xMTEgHBl5MxvoI8zPCJp5BpN-XDQQKlasf2Jw4EYlsu3bBOMfMi5kbnNjcnlwdC1jZXJ0Lm5zMy5jYS5sdWdncy5jbw


## opennic-luggs-ipv6

Public DNS server in Canada operated by Luggs (IPv6)

sdns://AQYAAAAAAAAAIVsyNjA3OjUzMDA6MTIwOmE4YToxNDI6NDoyMDQ6MTExXSAcGXkzG-gjzM8ImnkGk35cNBAqVqx_YnDgRiWy7dsE4x8yLmRuc2NyeXB0LWNlcnQubnMzLmNhLmx1Z2dzLmNv


## opennic-luggs2

Second public DNS server in Canada operated by Luggs

sdns://AQYAAAAAAAAAEDE0Mi40LjIwNS40Nzo0NDMgvL-34FDBPaJCLACwsaya1kjFwmS8thcLiD1xishuugkfMi5kbnNjcnlwdC1jZXJ0Lm5zNC5jYS5sdWdncy5jbw


## opennic-luggs2-ipv6

Second public DNS server in Canada operated by Luggs (IPv6)

sdns://AQYAAAAAAAAAJFsyNjA3OjUzMDA6MTIwOmE4YToxNDI6NDoyMDU6NDddOjQ0MyC8v7fgUME9okIsALCxrJrWSMXCZLy2FwuIPXGKyG66CR8yLmRuc2NyeXB0LWNlcnQubnM0LmNhLmx1Z2dzLmNv


## opennic-rico4514

OpenNIC • Non-logging • No DNSSEC
Location: Texas, 13, MX

sdns://AQYAAAAAAAAAETE0Mi40LjIwNC4xMTE6NDQzIBwZeTMb6CPMzwiaeQaTflw0ECpWrH9icOBGJbLt2wTjHzIuZG5zY3J5cHQtY2VydC5uczMuY2EubHVnZ3MuY28


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

sdns://AQcAAAAAAAAAEjEzOS45OS4yMjIuNzI6ODQ0MyALBWhPDQvh2BqXksLUVtlS0e0tH-tW5XqtWfE73l7gZRsyLmRuc2NyeXB0LWNlcnQuZG5zLnNlYnkuaW8


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


## quad9-dnscrypt-ip4-filter-alt
Quad9 (anycast) dnssec/no-log/filter 149.112.112.9
sdns://AQMAAAAAAAAAEjE0OS4xMTIuMTEyLjk6ODQ0MyBnyEe4yHWM0SAkVUO-dWdG3zTfHYTAC4xHA2jfgh2GPhkyLmRuc2NyeXB0LWNlcnQucXVhZDkubmV0


## quad9-dnscrypt-ip4-filter-pri
Quad9 (anycast) dnssec/no-log/filter 9.9.9.9
sdns://AQMAAAAAAAAADDkuOS45Ljk6ODQ0MyBnyEe4yHWM0SAkVUO-dWdG3zTfHYTAC4xHA2jfgh2GPhkyLmRuc2NyeXB0LWNlcnQucXVhZDkubmV0


## quad9-dnscrypt-ip4-nofilter-alt
Quad9 (anycast) no-dnssec/no-log/no-filter 149.112.112.10
sdns://AQYAAAAAAAAAEzE0OS4xMTIuMTEyLjEwOjg0NDMgZ8hHuMh1jNEgJFVDvnVnRt803x2EwAuMRwNo34Idhj4ZMi5kbnNjcnlwdC1jZXJ0LnF1YWQ5Lm5ldA


## quad9-dnscrypt-ip4-nofilter-pri
Quad9 (anycast) no-dnssec/no-log/no-filter 9.9.9.10
sdns://AQYAAAAAAAAADTkuOS45LjEwOjg0NDMgZ8hHuMh1jNEgJFVDvnVnRt803x2EwAuMRwNo34Idhj4ZMi5kbnNjcnlwdC1jZXJ0LnF1YWQ5Lm5ldA


## quad9-dnscrypt-ip6-filter-alt
Quad9 (anycast) dnssec/no-log/filter 2620:fe::9
sdns://AQMAAAAAAAAAEVsyNjIwOmZlOjo5XTo4NDQzIGfIR7jIdYzRICRVQ751Z0bfNN8dhMALjEcDaN-CHYY-GTIuZG5zY3J5cHQtY2VydC5xdWFkOS5uZXQ


## quad9-dnscrypt-ip6-filter-pri
Quad9 (anycast) dnssec/no-log/filter 2620:fe::fe:9
sdns://AQMAAAAAAAAAFFsyNjIwOmZlOjpmZTo5XTo4NDQzIGfIR7jIdYzRICRVQ751Z0bfNN8dhMALjEcDaN-CHYY-GTIuZG5zY3J5cHQtY2VydC5xdWFkOS5uZXQ


## quad9-dnscrypt-ip6-nofilter-alt
Quad9 (anycast) no-dnssec/no-log/no-filter 2620:fe::fe:10
sdns://AQYAAAAAAAAAFVsyNjIwOmZlOjpmZToxMF06ODQ0MyBnyEe4yHWM0SAkVUO-dWdG3zTfHYTAC4xHA2jfgh2GPhkyLmRuc2NyeXB0LWNlcnQucXVhZDkubmV0


## quad9-dnscrypt-ip6-nofilter-pri
Quad9 (anycast) no-dnssec/no-log/no-filter 2620:fe::10
sdns://AQYAAAAAAAAAElsyNjIwOmZlOjoxMF06ODQ0MyBnyEe4yHWM0SAkVUO-dWdG3zTfHYTAC4xHA2jfgh2GPhkyLmRuc2NyeXB0LWNlcnQucXVhZDkubmV0


## quad9-doh-ip4-filter-alt
Quad9 (anycast) dnssec/no-log/filter 149.112.112.9
sdns://AgMAAAAAAAAADTE0OS4xMTIuMTEyLjmAABJkbnM5LnF1YWQ5Lm5ldDo0NDMKL2Rucy1xdWVyeQ


## quad9-doh-ip4-filter-pri
Quad9 (anycast) dnssec/no-log/filter 9.9.9.9
sdns://AgMAAAAAAAAABzkuOS45LjmAABJkbnM5LnF1YWQ5Lm5ldDo0NDMKL2Rucy1xdWVyeQ


## quad9-doh-ip4-nofilter-alt
Quad9 (anycast) no-dnssec/no-log/no-filter 149.112.112.10
sdns://AgYAAAAAAAAADjE0OS4xMTIuMTEyLjEwgAASZG5zOS5xdWFkOS5uZXQ6NDQzCi9kbnMtcXVlcnk


## quad9-doh-ip4-nofilter-pri
Quad9 (anycast) no-dnssec/no-log/no-filter 9.9.9.10
sdns://AgYAAAAAAAAACDkuOS45LjEwgAASZG5zOS5xdWFkOS5uZXQ6NDQzCi9kbnMtcXVlcnk


## quad9-doh-ip6-filter-alt
Quad9 (anycast) dnssec/no-log/filter 2620:fe::fe:9
sdns://AgMAAAAAAAAAD1syNjIwOmZlOjpmZTo5XYAAEmRuczkucXVhZDkubmV0OjQ0MwovZG5zLXF1ZXJ5


## quad9-doh-ip6-filter-pri
Quad9 (anycast) dnssec/no-log/filter 2620:fe::9
sdns://AgMAAAAAAAAADFsyNjIwOmZlOjo5XYAAEmRuczkucXVhZDkubmV0OjQ0MwovZG5zLXF1ZXJ5


## quad9-doh-ip6-nofilter-alt
Quad9 (anycast) no-dnssec/no-log/no-filter 2620:fe::fe:10
sdns://AgYAAAAAAAAAEFsyNjIwOmZlOjpmZToxMF2AABJkbnM5LnF1YWQ5Lm5ldDo0NDMKL2Rucy1xdWVyeQ


## quad9-doh-ip6-nofilter-pri
Quad9 (anycast) no-dnssec/no-log/no-filter 2620:fe::10
sdns://AgYAAAAAAAAADVsyNjIwOmZlOjoxMF2AABJkbnM5LnF1YWQ5Lm5ldDo0NDMKL2Rucy1xdWVyeQ


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


## scaleway-ams

DNSSEC/Non-logged/Uncensored in Amsterdam - ARM server donated by Scaleway.com
Maintained by Frank Denis - https://fr.dnscrypt.info

sdns://AQcAAAAAAAAAETUxLjE1LjEyMi4yNTA6NDQzIOkN2X2qXGx4Ihyixe0BaJYtE8tVbM90V7NUVG8VK8GQHDIuZG5zY3J5cHQtY2VydC5zY2FsZXdheS1hbXM


## scaleway-ams-ipv6

DNSSEC/Non-logged/Uncensored in Amsterdam - IPv6 only - ARM server donated by Scaleway.com
Maintained by Frank Denis - https://fr.dnscrypt.info

sdns://AQcAAAAAAAAAHlsyMDAxOmJjODo0NzAwOjI1MDA6OjEzN2RdOjQ0MyDpDdl9qlxseCIcosXtAWiWLRPLVWzPdFezVFRvFSvBkBwyLmRuc2NyeXB0LWNlcnQuc2NhbGV3YXktYW1z


## scaleway-fr

DNSSEC/Non-logged/Uncensored in Paris - ARM server donated by Scaleway.com
Maintained by Frank Denis - https://fr.dnscrypt.info

sdns://AQcAAAAAAAAADjIxMi40Ny4yMjguMTM2IOgBuE6mBr-wusDOQ0RbsV66ZLAvo8SqMa4QY2oHkDJNHzIuZG5zY3J5cHQtY2VydC5mci5kbnNjcnlwdC5vcmc


## sea-dnscrypt-us

Resolver in Seattle (USA). DNSCrypt protocol. Non-logging, DNSSEC.
Remove ads and protect your computer from malware, phishing, ransomware and spam.

sdns://AQMAAAAAAAAAFDEwNC4xNjguMjQ3LjEzODo0MzQzIAOMTiaSmi-tS_UQJMeWX_BbeSaI4sBteYF81b-NFTOXHzIuZG5zY3J5cHQtY2VydC5zZWEtZG5zY3J5cHQtdXM


## sea-doh-us

Resolver in Seattle (USA). DoH protocol. Non-logging, DNSSEC, No-filter.

sdns://AgcAAAAAAAAAEzEwNC4xNjguMjQ3LjEzODo0NDMAFWRuczIuYWxla2JlcmcubmV0OjQ0MwEv


## securedns

Uncensored and no logging (DNSCrypt protocol)

sdns://AQcAAAAAAAAAEzE0Ni4xODUuMTY3LjQzOjUzNTMg9J8sc01itoYxntB-aRlDOy8ThfQe-8ovF21ZCy5FPoYcMi5kbnNjcnlwdC1jZXJ0LnNlY3VyZWRucy5ldQ


## securedns-doh

Uncensored and no logging (DoH protocol)

sdns://AgcAAAAAAAAADjE0Ni4xODUuMTY3LjQzID4aGg9sU_PpekktVwhLW5gHBZ7gV6sVBYdv2D_aPbg4EGRvaC5zZWN1cmVkbnMuZXUKL2Rucy1xdWVyeQ


## securedns-ipv6

Uncensored and no logging (DNSCrypt protocol)

sdns://AQcAAAAAAAAAIVsyYTAzOmIwYzA6MDoxMDEwOjplOWE6MzAwMV06NTM1MyD0nyxzTWK2hjGe0H5pGUM7LxOF9B77yi8XbVkLLkU-hhwyLmRuc2NyeXB0LWNlcnQuc2VjdXJlZG5zLmV1


## securedns-ipv6-doh

Uncensored and no logging (IPv6, DoH protocol)

sdns://AgcAAAAAAAAAHFsyYTAzOmIwYzA6MDoxMDEwOjplOWE6MzAwMV0gPhoaD2xT8-l6SS1XCEtbmAcFnuBXqxUFh2_YP9o9uDgQZG9oLnNlY3VyZWRucy5ldQovZG5zLXF1ZXJ5


## sfw.scaleway-fr

Uses deep learning to block adult websites. Free, DNSSEC, no logs.
Hosted in Paris, running on a 1-XS server donated by Scaleway.com
Maintained by Frank Denis - https://fr.dnscrypt.info/sfw.html

sdns://AQMAAAAAAAAAEzE2My4xNzIuMTgwLjEyNTo0NDMg32Jzv8dSGSqLWjm8DIWsP_lkRdc2RPZicoJdNVjxof8fMi5kbnNjcnlwdC1jZXJ0LnNmdy5zY2FsZXdheS1mcg


## skyfighter-dns

Uncensored, DNSSEC, no logging DNSCrypt server in Netherlands by Scaleway.com
Maintained by tuttimann

sdns://AQcAAAAAAAAADzUxLjE1LjYyLjY1OjQ0MyDRa9y0Ck5_qrri1JzXpTlOfz7vsOpqOvgepY6nL5iXmyIyLmRuc2NyeXB0LWNlcnQuc2t5ZmlnaHRlci1kbnMuY29t


## skyfighter-dns-ipv6

Uncensored, DNSSEC, no logging DNSCrypt server (IPv6) in Netherlands by Scaleway.com
Maintained by tuttimann

sdns://AQcAAAAAAAAAG1syMDAxOmJjODoxODI0OjE3MGY6OjFdOjQ0MyDRa9y0Ck5_qrri1JzXpTlOfz7vsOpqOvgepY6nL5iXmyIyLmRuc2NyeXB0LWNlcnQuc2t5ZmlnaHRlci1kbnMuY29t


## soltysiak

Public DNSCrypt server in Poland

sdns://AQcAAAAAAAAAFDE3OC4yMTYuMjAxLjIyMjoyMDUzICXE4YgpFUaXj5wrvbanr6QB7aBRBQhdUwPnGSjAZo8hHTIuZG5zY3J5cHQtY2VydC5zb2x0eXNpYWsuY29t


## sth-dnscrypt-se
Resolver in Stockholm, Sweden. DNSCrypt server. Non-logging, non-filtering, DNSSEC. Forward DNS to anoymized DNS servers

sdns://AQcAAAAAAAAAETQ1LjE1My4xODcuOTY6NDQzILwTiPV2Pj3sdfHJD8kVv7aK6t_rHEX2mFGAY5xg5_FiHzIuZG5zY3J5cHQtY2VydC5zdGgtZG5zY3J5cHQtc2U


## sth-doh-se
Resolver in Stockholm, Sweden. DoH server. Non-logging, non-filtering, DNSSEC. Forward DNS to anoymized DNS servers

sdns://AgcAAAAAAAAAETQ1LjE1My4xODcuOTY6NDQzABZkbnNzZS5hbGVrYmVyZy5uZXQ6NDQzBC9kbnM


## t53

Deutsche Telekom experimental DoH server.

sdns://AgUAAAAAAAAAACA-GhoPbFPz6XpJLVcIS1uYBwWe4FerFQWHb9g_2j24OApkbnMudDUzLmRlCi9kbnMtcXVlcnk


## v.dnscrypt.uk-ipv4

DNSCrypt v2, no logs, uncensored, DNSSEC. Hosted in London UK on Vultr
https://www.dnscrypt.uk

sdns://AQcAAAAAAAAAEzEwNC4yMzguMTg2LjE5Mjo0NDMg7Uk9jOrXkGZPBjxHt5WaI2ktfJA2PJ5DzLWRe-W0HuUdMi5kbnNjcnlwdC1jZXJ0LnYuZG5zY3J5cHQudWs


## v.dnscrypt.uk-ipv6

DNSCrypt v2, no logs, uncensored, DNSSEC. Hosted in London UK on Vultr
https://www.dnscrypt.uk

sdns://AQcAAAAAAAAALFsyMDAxOjE5ZjA6NzQwMjoxNTc0OjU0MDA6MmZmOmZlNjY6MmNmZl06NDQzIO1JPYzq15BmTwY8R7eVmiNpLXyQNjyeQ8y1kXvltB7lHTIuZG5zY3J5cHQtY2VydC52LmRuc2NyeXB0LnVr


## ventricle.us

Public DNSCrypt resolver provided by Jacob Henner

sdns://AQcAAAAAAAAADTEwNy4xNzAuNTcuMzQg6YXxGK1OPMZf8iUgGJDG9Vi3W1pS9WsXz-rBAFyLm6olMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeXB0LnZlbnRyaWNsZS51cw


## xfinity

Comcast DOH server

sdns://AgUAAAAAAAAAEjk2LjExMy4xNTEuMTQzOjQ0M6Ax_Wo8PCx8I5Gkl1qfoqes0mp4xMrk1W5GIynLLGg2syANc6YU7vdZbPWjRzP3Ta8sz-Tfe0pABpv0PEPkKCZBdw9kb2gueGZpbml0eS5jb20KL2Rucy1xdWVyeQ


## yandex

Yandex public DNS server (anycast)

sdns://AQQAAAAAAAAAEDc3Ljg4LjguNzg6MTUzNTMg04TAccn3RmKvKszVe13MlxTUB7atNgHhrtwG1W1JYyciMi5kbnNjcnlwdC1jZXJ0LmJyb3dzZXIueWFuZGV4Lm5ldA
