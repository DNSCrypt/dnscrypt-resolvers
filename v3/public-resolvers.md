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
    urls = ['https://raw.githubusercontent.com/DNSCrypt/dnscrypt-resolvers/master/v3/public-resolvers.md', 'https://download.dnscrypt.info/resolvers-list/v3/public-resolvers.md']
    minisign_key = 'RWQf6LRCGA9i53mlYecO4IzT51TGPpvWucNSCh1CBM0QTaLn73Y7GFO3'
    cache_file = 'public-resolvers.md'

--


## a-and-a

Non-logging DoH server in the UK operated by Andrews & Arnold Ltd, a
company providing Internet connectivity and VoIP in the UK.

https://www.aa.net.uk/legal/dohdot-disclaimer/

sdns://AgcAAAAAAAAADTIxNy4xNjkuMjAuMjOgPhoaD2xT8-l6SS1XCEtbmAcFnuBXqxUFh2_YP9o9uDggMob_ZaZfrzIIXuoTiMNzi6fjeHPJBszjxKKLTMKliYgNZG5zLmFhLm5ldC51awovZG5zLXF1ZXJ5


## acsacsar-ams-ipv4

Public non-censoring, non-logging, DNSSEC-capable, DNSCrypt-enabled DNS resolver hosted on Scaleway by @acsacsar (twitter)

sdns://AQcAAAAAAAAADTUxLjE1OC4xNjYuOTcgAyfzz5J-mV9G-yOB4Hwcdk7yX12EQs5Iva7kV3oGtlEgMi5kbnNjcnlwdC1jZXJ0LmFjc2Fjc2FyLWFtcy5jb20


## acsacsar-ams-ipv6

Public non-censoring, non-logging, DNSSEC-capable, DNSCrypt-enabled DNS resolver hosted on Scaleway by @acsacsar (twitter)

sdns://AQcAAAAAAAAAFlsyMDAxOmJjODoxODI0OjczODo6MV0gAyfzz5J-mV9G-yOB4Hwcdk7yX12EQs5Iva7kV3oGtlEgMi5kbnNjcnlwdC1jZXJ0LmFjc2Fjc2FyLWFtcy5jb20


## adfree.usableprivacy.net

Public non-logging DoH server with advertising and tracker filtering.

Hosted in Austria/Europe, details see: [docs.usableprivacy.com](https://docs.usableprivacy.com)

sdns://AgMAAAAAAAAADzE0OS4xNTQuMTUzLjE1M6A-GhoPbFPz6XpJLVcIS1uYBwWe4FerFQWHb9g_2j24OCAyhv9lpl-vMghe6hOIw3OLp-N4c8kGzOPEootMwqWJiBhhZGZyZWUudXNhYmxlcHJpdmFjeS5uZXQKL2Rucy1xdWVyeQ


## adguard-dns

Remove ads and protect your computer from malware

Warning: This server is incompatible with anonymization.

sdns://AQIAAAAAAAAAFDE3Ni4xMDMuMTMwLjEzMDo1NDQzINErR_JS3PLCu_iZEIbq95zkSV2LFsigxDIuUso_OQhzIjIuZG5zY3J5cHQuZGVmYXVsdC5uczEuYWRndWFyZC5jb20


## adguard-dns-doh

Remove ads and protect your computer from malware (over DoH)

sdns://AgMAAAAAAAAADzE3Ni4xMDMuMTMwLjEzMCCsFdIhxY-VWoedpSrEKWAhaBEVj-8L-p_FJl6wMpPufg9kbnMuYWRndWFyZC5jb20KL2Rucy1xdWVyeQ


## adguard-dns-family

Adguard DNS with safesearch and adult content blocking

Warning: This server is incompatible with anonymization.

sdns://AQIAAAAAAAAAFDE3Ni4xMDMuMTMwLjEzMjo1NDQzILgxXdexS27jIKRw3C7Wsao5jMnlhvhdRUXWuMm1AFq6ITIuZG5zY3J5cHQuZmFtaWx5Lm5zMS5hZGd1YXJkLmNvbQ


## adguard-dns-family-doh

Adguard DNS with safesearch and adult content blocking (over DoH)

sdns://AgMAAAAAAAAADzE3Ni4xMDMuMTMwLjEzMiCsFdIhxY-VWoedpSrEKWAhaBEVj-8L-p_FJl6wMpPufhZkbnMtZmFtaWx5LmFkZ3VhcmQuY29tCi9kbnMtcXVlcnk


## adguard-dns-family-ipv6

Adguard DNS with safesearch and adult content blocking

Warning: This server is incompatible with anonymization.

sdns://AQMAAAAAAAAAGlsyYTAwOjVhNjA6OmJhZDI6MGZmXTo1NDQzIIwhF6nrwVfW-2QFbwrbwRxdg2c0c8RuJY2bL1fU7jUfITIuZG5zY3J5cHQuZmFtaWx5Lm5zMi5hZGd1YXJkLmNvbQ


## adguard-dns-ipv6

Remove ads and protect your computer from malware

Warning: This server is incompatible with anonymization.

sdns://AQMAAAAAAAAAGVsyYTAwOjVhNjA6OmFkMjowZmZdOjU0NDMggdAC02pMpQxHO3R5ZQ_hLgKzIcthOFYqII5APf3FXpQiMi5kbnNjcnlwdC5kZWZhdWx0Lm5zMi5hZGd1YXJkLmNvbQ


## adguard-dns-unfiltered

AdGuard public DNS servers without filters

Warning: This server is incompatible with anonymization.

sdns://AQcAAAAAAAAAFDE3Ni4xMDMuMTMwLjEzNjo1NDQzILXoRNa4Oj4-EmjraB--pw3jxfpo29aIFB2_LsBmstr6JTIuZG5zY3J5cHQudW5maWx0ZXJlZC5uczEuYWRndWFyZC5jb20


## ahadns-doh-in

A zero logging DNS with support for DNS-over-HTTPS (DoH) & DNS-over-TLS (DoT). Blocks ads, malware, trackers, viruses, ransomware, telemetry and more. No persistent logs. DNSSEC. Hosted in Mumbai, India. By https://ahadns.com/
Server statistics can be seen at: https://statistics.ahadns.com/?server=in

sdns://AgMAAAAAAAAADTQ1Ljc5LjEyMC4yMzOgPhoaD2xT8-l6SS1XCEtbmAcFnuBXqxUFh2_YP9o9uDggMob_ZaZfrzIIXuoTiMNzi6fjeHPJBszjxKKLTMKliYgRZG9oLmluLmFoYWRucy5uZXQKL2Rucy1xdWVyeQ


## ahadns-doh-in-ipv6

A zero logging DNS with support for DNS-over-HTTPS (DoH) & DNS-over-TLS (DoT). Blocks ads, malware, trackers, viruses, ransomware, telemetry and more. No persistent logs. DNSSEC. Hosted in Mumbai, India. By https://ahadns.com/
Server statistics can be seen at: https://statistics.ahadns.com/?server=in

sdns://AgMAAAAAAAAAF1syNDAwOjg5MDQ6ZTAwMTo0Mzo6NDNdoD4aGg9sU_PpekktVwhLW5gHBZ7gV6sVBYdv2D_aPbg4IDKG_2WmX68yCF7qE4jDc4un43hzyQbM48Sii0zCpYmIEWRvaC5pbi5haGFkbnMubmV0Ci9kbnMtcXVlcnk


## ahadns-doh-la

A zero logging DNS with support for DNS-over-HTTPS (DoH) & DNS-over-TLS (DoT). Blocks ads, malware, trackers, viruses, ransomware, telemetry and more. No persistent logs. DNSSEC. Hosted in Los Angeles, USA. By https://ahadns.com/
Server statistics can be seen at: https://statistics.ahadns.com/?server=la

sdns://AgMAAAAAAAAADTQ1LjY3LjIxOS4yMDigPhoaD2xT8-l6SS1XCEtbmAcFnuBXqxUFh2_YP9o9uDggMob_ZaZfrzIIXuoTiMNzi6fjeHPJBszjxKKLTMKliYgRZG9oLmxhLmFoYWRucy5uZXQKL2Rucy1xdWVyeQ


## ahadns-doh-la-ipv6

A zero logging DNS with support for DNS-over-HTTPS (DoH) & DNS-over-TLS (DoT). Blocks ads, malware, trackers, viruses, ransomware, telemetry and more. No persistent logs. DNSSEC. Hosted in Los Angeles, USA. By https://ahadns.com/
Server statistics can be seen at: https://statistics.ahadns.com/?server=la

sdns://AgMAAAAAAAAAFlsyYTA0OmJkYzc6MTAwOjcwOjo3MF2gPhoaD2xT8-l6SS1XCEtbmAcFnuBXqxUFh2_YP9o9uDggMob_ZaZfrzIIXuoTiMNzi6fjeHPJBszjxKKLTMKliYgRZG9oLmxhLmFoYWRucy5uZXQKL2Rucy1xdWVyeQ


## ahadns-doh-nl

A zero logging DNS with support for DNS-over-HTTPS (DoH) & DNS-over-TLS (DoT). Blocks ads, malware, trackers, viruses, ransomware, telemetry and more. No persistent logs. DNSSEC. Hosted in Amsterdam, Netherlands. By https://ahadns.com/
Server statistics can be seen at: https://statistics.ahadns.com/?server=nl

sdns://AgMAAAAAAAAACTUuMi43NS43NaA-GhoPbFPz6XpJLVcIS1uYBwWe4FerFQWHb9g_2j24OCAyhv9lpl-vMghe6hOIw3OLp-N4c8kGzOPEootMwqWJiBFkb2gubmwuYWhhZG5zLm5ldAovZG5zLXF1ZXJ5


## ahadns-doh-nl-ipv6

A zero logging DNS with support for DNS-over-HTTPS (DoH) & DNS-over-TLS (DoT). Blocks ads, malware, trackers, viruses, ransomware, telemetry and more. No persistent logs. DNSSEC. Hosted in Amsterdam, Netherlands. By https://ahadns.com/
Server statistics can be seen at: https://statistics.ahadns.com/?server=nl

sdns://AgMAAAAAAAAAFlsyYTA0OjUyYzA6MTAxOjc1Ojo3NV2gPhoaD2xT8-l6SS1XCEtbmAcFnuBXqxUFh2_YP9o9uDggMob_ZaZfrzIIXuoTiMNzi6fjeHPJBszjxKKLTMKliYgRZG9oLm5sLmFoYWRucy5uZXQKL2Rucy1xdWVyeQ


## alidns-doh

A public DNS resolver that supports DoH/DoT in mainland China, provided by Alibaba-Cloud.

Warning: GFW filtering rules are applied by that resolver.

Homepage: https://alidns.com/

sdns://AgAAAAAAAAAACTIyMy41LjUuNSCoF6cUD2dwqtorNi96I2e3nkHPSJH1ka3xbdOglmOVkQ5kbnMuYWxpZG5zLmNvbQovZG5zLXF1ZXJ5


## ams-ads-doh-nl

Resolver in Amsterdam. DoH protocol. Non-logging. Blocks ads, malware and trackers. DNSSEC enabled.

sdns://AgMAAAAAAAAAEjUxLjE1LjEyNC4yMDg6NDQ0M6A-GhoPbFPz6XpJLVcIS1uYBwWe4FerFQWHb9g_2j24OCAyhv9lpl-vMghe6hOIw3OLp-N4c8kGzOPEootMwqWJiBdkbnNubC5hbGVrYmVyZy5uZXQ6NDQ0MwovZG5zLXF1ZXJ5


## ams-ads-doh-nl-ipv6

Resolver in Amsterdam. DoH protocol. Non-logging. Blocks ads, malware and trackers. DNSSEC enabled.

sdns://AgMAAAAAAAAAHFsyMDAxOmJjODoxODMwOjIwMTg6OjFdOjQ0NDOgPhoaD2xT8-l6SS1XCEtbmAcFnuBXqxUFh2_YP9o9uDggMob_ZaZfrzIIXuoTiMNzi6fjeHPJBszjxKKLTMKliYgXZG5zbmwuYWxla2JlcmcubmV0OjQ0NDMKL2Rucy1xdWVyeQ


## ams-dnscrypt-nl-ipv6

Resolver in Amsterdam. Dnscrypt protocol. Non-logging, non-filtering, DNSSEC.

sdns://AQcAAAAAAAAAHFsyMDAxOmJjODoxODMwOjIwMTg6OjFdOjQzNDMgvBOI9XY-Pex18ckPyRW_torq3-scRfaYUYBjnGDn8WIfMi5kbnNjcnlwdC1jZXJ0LmFtcy1kbnNjcnlwdC1ubA


## ams-doh-nl

Resolver in Amsterdam. DoH protocol. Non-logging, non-filtering, DNSSEC.

sdns://AgcAAAAAAAAADTUxLjE1LjEyNC4yMDigPhoaD2xT8-l6SS1XCEtbmAcFnuBXqxUFh2_YP9o9uDggMob_ZaZfrzIIXuoTiMNzi6fjeHPJBszjxKKLTMKliYgSZG5zbmwuYWxla2JlcmcubmV0Ci9kbnMtcXVlcnk


## ams-doh-nl-ipv6

Resolver in Amsterdam. DoH protocol. Non-logging, non-filtering, DNSSEC.

sdns://AgcAAAAAAAAAF1syMDAxOmJjODoxODMwOjIwMTg6OjFdoD4aGg9sU_PpekktVwhLW5gHBZ7gV6sVBYdv2D_aPbg4IDKG_2WmX68yCF7qE4jDc4un43hzyQbM48Sii0zCpYmIEmRuc25sLmFsZWtiZXJnLm5ldAovZG5zLXF1ZXJ5


## arapurayil-doh-in

DoH server(ipv4/ipv6). Also supports DNSCrypt,DoT,DoQ protocols. Located in Mumbai, India.  Visit https://www.dns.arapurayil.com for details.
Non-logging | Filtering | DNSSEC | anti-CNAME cloaking | QNAME Minimization | No EDNS Client-Subnet

sdns://AgMAAAAAAAAAACAyhv9lpl-vMghe6hOIw3OLp-N4c8kGzOPEootMwqWJiBJkbnMuYXJhcHVyYXlpbC5jb20KL2Rucy1xdWVyeQ


## arvind-io

Public resolver by EnKrypt (https://arvind.io).
Hosted in Bangalore, India.

Non-logging, non-filtering, supports DNSSEC.

sdns://AQcAAAAAAAAAFDIwNi4xODkuMTQyLjE3OTo1MzUzII5GJ8c4g6hRAwghulrn5dBB9KrvlbeCkBbLZR2HwyjJGTIuZG5zY3J5cHQtY2VydC5hcnZpbmQuaW8


## att

AT&T test DoH server.

sdns://AgQAAAAAAAAAAKBLTrSwdCmLgotcADCVoQtFI_uVHAyINIsJxT5bq6QIoyD2Hldod9qWUClMzLX5bHX8txvaG7xGRjZ8Tr7aidcxjxBkb2h0cmlhbC5hdHQubmV0Ci9kbnMtcXVlcnk


## bcn-dnscrypt

Resolver in Barcelona, Spain. DNSCrypt protocol. Non-logging, non-filtering, DNSSEC.

sdns://AQcAAAAAAAAAEzE4NS4yNTMuMTU0LjY2OjQzNDMg_IlDWhFp1tL1VycXVO3QTEiQIKG_1OwG4tNTAD-nLj0cMi5kbnNjcnlwdC1jZXJ0LmJjbi1kbnNjcnlwdA


## bcn-doh

Resolver in Barcelona, Spain. DoH protocol. Non-logging, non-filtering, DNSSEC.

sdns://AgcAAAAAAAAADjE4NS4yNTMuMTU0LjY2oD4aGg9sU_PpekktVwhLW5gHBZ7gV6sVBYdv2D_aPbg4IDKG_2WmX68yCF7qE4jDc4un43hzyQbM48Sii0zCpYmIEmRuc2VzLmFsZWtiZXJnLm5ldAovZG5zLXF1ZXJ5


## bortzmeyer

Non-logging DoH server in France operated by Stéphane Bortzmeyer.

https://www.bortzmeyer.org/doh-bortzmeyer-fr-policy.html

sdns://AgcAAAAAAAAADDE5My43MC44NS4xMaA-GhoPbFPz6XpJLVcIS1uYBwWe4FerFQWHb9g_2j24OCAyhv9lpl-vMghe6hOIw3OLp-N4c8kGzOPEootMwqWJiBFkb2guYm9ydHptZXllci5mcgEv


## bortzmeyer-ipv6

Non-logging DoH server in France operated by Stéphane Bortzmeyer (IPv6 only).

https://www.bortzmeyer.org/doh-bortzmeyer-fr-policy.html

sdns://AgcAAAAAAAAAGVsyMDAxOjQxZDA6MzAyOjIyMDA6OjE4MF2gPhoaD2xT8-l6SS1XCEtbmAcFnuBXqxUFh2_YP9o9uDggMob_ZaZfrzIIXuoTiMNzi6fjeHPJBszjxKKLTMKliYgRZG9oLmJvcnR6bWV5ZXIuZnIBLw


## captnemo-in

Server running out of a Digital Ocean droplet in BLR1 region.
Maintained by Abhay Rana aka Nemo.

If you are within India, this might be a nice DNS server to use.

sdns://AQQAAAAAAAAAEjEzOS41OS40OC4yMjI6NDQzNCAFOt_yxaMpFtga2IpneSwwK6rV0oAyleham9IvhoceEBsyLmRuc2NyeXB0LWNlcnQuY2FwdG5lbW8uaW4


## circl-doh

DoH server operated by CIRCL, Computer Incident Response Center Luxembourg.

Hosted in Bettembourg, Luxembourg.

sdns://AgYAAAAAAAAADTE4NS4xOTQuOTQuNzEADGRucy5jaXJjbC5sdQovZG5zLXF1ZXJ5


## circl-doh-ipv6

DoH server operated by CIRCL, Computer Incident Response Center Luxembourg.

Hosted in Bettembourg, Luxembourg.

sdns://AgYAAAAAAAAAElsyYTAwOjU5ODA6OTQ6OjcxXQAMZG5zLmNpcmNsLmx1Ci9kbnMtcXVlcnk


## cisco

Remove your DNS blind spot (DNSCrypt protocol)

Warning: This server is incompatible with anonymization.

Warning: modifies your queries to include a copy of your network
address when forwarding them to a selection of companies and organizations.

sdns://AQEAAAAAAAAADjIwOC42Ny4yMjAuMjIwILc1EUAgbyJdPivYItf9aR6hwzzI1maNDL4Ev6vKQ_t5GzIuZG5zY3J5cHQtY2VydC5vcGVuZG5zLmNvbQ


## cisco-doh

Remove your DNS blind spot (DoH protocol)

Warning: modifies your queries to include a copy of your network
address when forwarding them to a selection of companies and organizations.

sdns://AgAAAAAAAAAADDE0Ni4xMTIuNDEuMiBoU4_HgY6B0kIqkGBjb6UoKkP2Dc4bumDC1_Orq2YAlw9kb2gub3BlbmRucy5jb20KL2Rucy1xdWVyeQ


## cisco-familyshield

Block websites not suitable for children (DNSCrypt protocol)

Warning: modifies your queries to include a copy of your network
address when forwarding them to a selection of companies and organizations.

Currently incompatible with DNS anonymization.

sdns://AQEAAAAAAAAADjIwOC42Ny4yMjAuMTIzILc1EUAgbyJdPivYItf9aR6hwzzI1maNDL4Ev6vKQ_t5GzIuZG5zY3J5cHQtY2VydC5vcGVuZG5zLmNvbQ


## cisco-familyshield-ipv6

Block websites not suitable for children (IPv6)

Warning: This server is incompatible with anonymization.

Warning: modifies your queries to include a copy of your network
address when forwarding them to a selection of companies and organizations.

sdns://AQEAAAAAAAAAEVsyNjIwOjExOTozNTo6MzVdILc1EUAgbyJdPivYItf9aR6hwzzI1maNDL4Ev6vKQ_t5GzIuZG5zY3J5cHQtY2VydC5vcGVuZG5zLmNvbQ


## cisco-ipv6

Cisco OpenDNS over IPv6 (DNSCrypt protocol)

Warning: This server is incompatible with anonymization.

Warning: modifies your queries to include a copy of your network
address when forwarding them to a selection of companies and organizations.

sdns://AQEAAAAAAAAAD1syNjIwOjA6Y2NjOjoyXSC3NRFAIG8iXT4r2CLX_WkeocM8yNZmjQy-BL-rykP7eRsyLmRuc2NyeXB0LWNlcnQub3BlbmRucy5jb20


## cisco-ipv6-doh

Cisco OpenDNS over IPv6 (DoH protocol)

Warning: modifies your queries to include a copy of your network
address when forwarding them to a selection of companies and organizations.

sdns://AgAAAAAAAAAAEFsyNjIwOjExOTpmYzo6Ml0gaFOPx4GOgdJCKpBgY2-lKCpD9g3OG7pgwtfzq6tmAJcPZG9oLm9wZW5kbnMuY29tCi9kbnMtcXVlcnk


## cleanbrowsing-adult

Blocks access to all adult, pornographic and explicit sites. It does
not block proxy or VPNs, nor mixed-content sites. Sites like Reddit
are allowed. Google and Bing are set to the Safe Mode.

Warning: This server is incompatible with anonymization.

By https://cleanbrowsing.org/

sdns://AQMAAAAAAAAAEzE4NS4yMjguMTY4LjEwOjg0NDMgvKwy-tVDaRcfCDLWB1AnwyCM7vDo6Z-UGNx3YGXUjykRY2xlYW5icm93c2luZy5vcmc


## cleanbrowsing-adult-ipv6

Blocks access to all adult, pornographic and explicit sites. It does
not block proxy or VPNs, nor mixed-content sites. Sites like Reddit
are allowed. Google and Bing are set to the Safe Mode.

Warning: This server is incompatible with anonymization.

By https://cleanbrowsing.org/

sdns://AQMAAAAAAAAAFVsyYTBkOjJhMDA6MTo6MV06ODQ0MyC8rDL61UNpFx8IMtYHUCfDIIzu8Ojpn5QY3HdgZdSPKRFjbGVhbmJyb3dzaW5nLm9yZw


## cleanbrowsing-family

Blocks access to all adult, pornographic and explicit sites. It also
blocks proxy and VPN domains that are used to bypass the filters.
Mixed content sites (like Reddit) are also blocked. Google, Bing and
Youtube are set to the Safe Mode.

Warning: This server is incompatible with anonymization.

By https://cleanbrowsing.org/

sdns://AQMAAAAAAAAAFDE4NS4yMjguMTY4LjE2ODo4NDQzILysMvrVQ2kXHwgy1gdQJ8MgjO7w6OmflBjcd2Bl1I8pEWNsZWFuYnJvd3Npbmcub3Jn


## cleanbrowsing-family-ipv6

Blocks access to all adult, pornographic and explicit sites. It also
blocks proxy and VPN domains that are used to bypass the filters.
Mixed content sites (like Reddit) are also blocked. Google, Bing and
Youtube are set to the Safe Mode.

Warning: This server is incompatible with anonymization.

By https://cleanbrowsing.org/

sdns://AQMAAAAAAAAAFFsyYTBkOjJhMDA6MTo6XTo4NDQzILysMvrVQ2kXHwgy1gdQJ8MgjO7w6OmflBjcd2Bl1I8pEWNsZWFuYnJvd3Npbmcub3Jn


## cleanbrowsing-security

Block access to phishing, malware and malicious domains. It does not block adult content.
By https://cleanbrowsing.org/

Warning: This server is incompatible with anonymization.

sdns://AQMAAAAAAAAAEjE4NS4yMjguMTY4Ljk6ODQ0MyC8rDL61UNpFx8IMtYHUCfDIIzu8Ojpn5QY3HdgZdSPKRFjbGVhbmJyb3dzaW5nLm9yZw


## cloudflare

Cloudflare DNS (anycast) - aka 1.1.1.1 / 1.0.0.1

sdns://AgcAAAAAAAAABzEuMC4wLjEAEmRucy5jbG91ZGZsYXJlLmNvbQovZG5zLXF1ZXJ5


## cloudflare-family

Cloudflare DNS (anycast) with malware protection and parental control - aka 1.1.1.3 / 1.0.0.3

sdns://AgMAAAAAAAAABzEuMC4wLjMAGWZhbWlseS5jbG91ZGZsYXJlLWRucy5jb20KL2Rucy1xdWVyeQ


## cloudflare-family-ipv6

Cloudflare DNS over IPv6 (anycast) with malware protection and parental control

sdns://AgMAAAAAAAAAFlsyNjA2OjQ3MDA6NDcwMDo6MTExM10AGWZhbWlseS5jbG91ZGZsYXJlLWRucy5jb20KL2Rucy1xdWVyeQ
sdns://AgMAAAAAAAAAFlsyNjA2OjQ3MDA6NDcwMDo6MTAwM10AGWZhbWlseS5jbG91ZGZsYXJlLWRucy5jb20KL2Rucy1xdWVyeQ


## cloudflare-ipv6

Cloudflare DNS over IPv6 (anycast)

sdns://AgcAAAAAAAAAFlsyNjA2OjQ3MDA6NDcwMDo6MTExMV0AIDFkb3QxZG90MWRvdDEuY2xvdWRmbGFyZS1kbnMuY29tCi9kbnMtcXVlcnk
sdns://AgcAAAAAAAAAFlsyNjA2OjQ3MDA6NDcwMDo6MTAwMV0AIDFkb3QxZG90MWRvdDEuY2xvdWRmbGFyZS1kbnMuY29tCi9kbnMtcXVlcnk


## cloudflare-security

Cloudflare DNS (anycast) with malware blocking - aka 1.1.1.2 / 1.0.0.2

sdns://AgMAAAAAAAAABzEuMC4wLjIAG3NlY3VyaXR5LmNsb3VkZmxhcmUtZG5zLmNvbQovZG5zLXF1ZXJ5


## cloudflare-security-ipv6

Cloudflare DNS over IPv6 (anycast) with malware blocking

sdns://AgMAAAAAAAAAFlsyNjA2OjQ3MDA6NDcwMDo6MTExMl0AG3NlY3VyaXR5LmNsb3VkZmxhcmUtZG5zLmNvbQovZG5zLXF1ZXJ5
sdns://AgMAAAAAAAAAFlsyNjA2OjQ3MDA6NDcwMDo6MTAwMl0AG3NlY3VyaXR5LmNsb3VkZmxhcmUtZG5zLmNvbQovZG5zLXF1ZXJ5


## commons-host

DoH server by the Commons Host CDN

sdns://AgUAAAAAAAAAAKA-GhoPbFPz6XpJLVcIS1uYBwWe4FerFQWHb9g_2j24OCAyhv9lpl-vMghe6hOIw3OLp-N4c8kGzOPEootMwqWJiAxjb21tb25zLmhvc3QKL2Rucy1xdWVyeQ


## comodo-02

Comodo Dome Shield (anycast) - https://cdome.comodo.com/shield/

sdns://AQUAAAAAAAAACjguMjAuMjQ3LjIg0sJUqpYcHsoXmZb1X7yAHwg2xyN5q1J-zaiGG-Dgs7AoMi5kbnNjcnlwdC1jZXJ0LnNoaWVsZC0yLmRuc2J5Y29tb2RvLmNvbQ


## comss.one

DNS server in Lithuania filtering phishing and ads.
https://www.comss.ru/page.php?id=7315

sdns://AQMAAAAAAAAADTk0LjE3Ni4yMzMuOTMgFWnIA4ZtJKvnIs9g6yZT4WIyPb-rQEB27paIxM_OkxMdMi5kbnNjcnlwdC1jZXJ0LmRucy5jb21zcy5vbmU


## containerpi

Non-logging, non-filtering, DNSSEC validating server, EDNS Client Subnet enabled.
Multiple nodes in China Mainland, Japan and Germany.
Maintained by CPI-tech-Group

sdns://AgMAAAAAAAAADDQ1Ljc3LjE4MC4xMCA59q74zeUV4xVJXPXvLw0G1pi3YdoaF4FskwSxGJjS1BNkbnMuY29udGFpbmVycGkuY29tCi9kbnMtcXVlcnk


## containerpi-ipv6

Non-logging, non-filtering, DNSSEC validating server, EDNS Client Subnet enabled.
Multiple nodes in China Mainland, Japan and Germany.
Maintained by CPI-tech-Group

sdns://AgMAAAAAAAAAKVsyMDAxOjE5ZjA6NzAwMTo1NTU0OjU0MDA6MDJmZjpmZTU3OjMwNzddIDn2rvjN5RXjFUlc9e8vDQbWmLdh2hoXgWyTBLEYmNLUE2Rucy5jb250YWluZXJwaS5jb20KL2Rucy1xdWVyeQ


## cs-austria

Wien, Austria DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADTk0LjE5OC40MS4yMzUgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz


## cs-barcelona

Barcelona, Spain DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADjM3LjEyMC4xNDIuMTE1IDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-belgium

Brussels, Belgium DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADTM3LjEyMC4yMzYuMTEgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz


## cs-bulgaria

Sofia, Bulgaria DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADjM3LjEyMC4xNTIuMjM1IDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-ch

Switzerland DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAACzgxLjE3LjMxLjM0IDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-czech

Prague, Czech Republic DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADzIxNy4xMzguMjIwLjI0MyAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM


## cs-dc

US - Washington, DC DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADDE5OC43LjU4LjIyNyAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM


## cs-de

Frankfurt, Germany DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADDg0LjE2LjI0MC40MyAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM


## cs-dk

Denmark DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADTM3LjEyMC4yMzIuNDMgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz


## cs-dus1

Dusseldorf, Germany 1 DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADjIxMy4yMDIuMjE2LjEyIDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-dus2

Dusseldorf, Germany 2 DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADzIxMy4yMDIuMjE2LjIzNiAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM


## cs-dus3

Dusseldorf, Germany 3 DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADjg5LjE2My4yMjEuMTgxIDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-dus4

Dusseldorf, Germany 4 DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADjg1LjExNC4xMzguMTE5IDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-finland

Finland DNSCrypt server provided by https://cryptostorm.is

sdns://AQYAAAAAAAAADjE4NS4xMTcuMTE4LjIwIDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-fr

France DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADTIxMi4xMjkuNDYuMzIgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz


## cs-ga

US - Atlanta, GA DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADTY0LjQyLjE4MS4yMjcgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz


## cs-hungary

Budapest, Hungary DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADTg2LjEwNi43NC4yMTkgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz


## cs-il

US - Chicago, IL DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADjE3My4yMzQuNTYuMTE1IDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-ireland

Dublin, Ireland DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADjM3LjEyMC4yMzUuMTg3IDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-is

Iceland DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADTgyLjIyMS4xMjguNDQgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz


## cs-la

US - Los Angeles, CA DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADDIzLjE5LjY3LjExNiAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM


## cs-london

London, England DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADTc4LjEyOS4xNDAuNjUgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz


## cs-lv

Latvia DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADzEwOS4yNDguMTQ5LjEzMyAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM


## cs-madrid

Madrid, Spain DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADjE4NS4xODMuMTA2LjgzIDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-manchester

Manchester, England DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADTE5NS4xMi40OC4xNzEgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz


## cs-md

Moldova DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADzE3OC4xNzUuMTM5LjIxMSAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM


## cs-milan

Milan, Italy DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADzIxNy4xMzguMjE5LjIxOSAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM


## cs-montreal

Montreal, Canada DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADjE2Ny4xMTQuODQuMTMyIDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-nc

US - North Carolina DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADjE1NS4yNTQuMjkuMTEzIDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-nl

Netherlands DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADTE4NS4xMDcuODAuODQgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz


## cs-nl2

Netherlands 2 DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADjIxMy4xNjMuNjQuMjA4IDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-norway

Oslo, Norway DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADjkxLjIxOS4yMTUuMjI3IDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-nv

US - Las Vegas, NV DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADDM3LjEyMC4xNDcuMiAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM


## cs-nyc1

US - New York City, NY 1 DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADDIzLjE5LjI0NS44OCAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM


## cs-nyc2

US - New York City, NY 2 DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADDIzLjE5LjI0NS44NCAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM


## cs-ore

US - Oregon DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADTEwNC4yNTUuMTc1LjIgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz


## cs-poland

Warsaw, Poland DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADTM3LjEyMC4yMTEuOTEgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz


## cs-pt

Portugal DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADTEwOS43MS40Mi4yMjggMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz


## cs-ro

Romania DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADDUuMjU0Ljk2LjE5NSAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM


## cs-rome

Rome, Italy DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADjM3LjEyMC4yMDcuMTMxIDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-sea

US - Seattle, WA DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADDY0LjEyMC41LjI1MSAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM


## cs-serbia

Belgrade, Serbia DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADjM3LjEyMC4xOTMuMjE5IDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-sk

South Korea DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADDI3LjI1NS43Ny41NiAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM


## cs-slovakia

Bratislava, Slovakia DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADjE5My4zNy4yNTUuMjI3IDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-swe

Sweden DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADzEyOC4xMjcuMTA0LjEwOCAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM


## cs-tx

US - Dallas, TX DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADTIwOS41OC4xNDcuMzYgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz


## cs-vancouver

Vancouver, Canada DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADzE2Mi4yMjEuMjA3LjIyOCAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM


## cz.nic

CZ.NIC's open DNSSEC validating resolvers in Prague, Czech Republic.
CZ.NIC resolvers neither collect any personal data nor gather
information on pages where your computer sends personal data.
https://www.nic.cz/odvr/

sdns://AgcAAAAAAAAADDE4NS40My4xMzUuMaA-GhoPbFPz6XpJLVcIS1uYBwWe4FerFQWHb9g_2j24OCAyhv9lpl-vMghe6hOIw3OLp-N4c8kGzOPEootMwqWJiAtvZHZyLm5pYy5jegQvZG9o


## cz.nic-ipv6

CZ.NIC's open DNSSEC validating resolvers in Prague, Czech Republic (IPv6 only).
CZ.NIC resolvers neither collect any personal data nor gather
information on pages where your computer sends personal data.
https://www.nic.cz/odvr/

sdns://AgcAAAAAAAAAE1syMDAxOjE0OGY6ZmZmZTo6MV2gPhoaD2xT8-l6SS1XCEtbmAcFnuBXqxUFh2_YP9o9uDggMob_ZaZfrzIIXuoTiMNzi6fjeHPJBszjxKKLTMKliYgLb2R2ci5uaWMuY3oEL2RvaA


## d0wn-is-ns2

Server provided by Martin 'd0wn' Albus

sdns://AQcAAAAAAAAADTkzLjk1LjIyNi4xNjUghGA0qcYwyjwErEqQFiXxeoeyrLlBgKxIHiwQ6M7eGm8cMi5kbnNjcnlwdC1jZXJ0LmlzMi5kMHduLmJpeg


## d0wn-tz-ns1

Server provided by Martin 'd0wn' Albus

sdns://AQcAAAAAAAAACzQxLjc5LjY5LjEzINYGFfvRRTuhTnaKPlxcs6wXRhMxRj2gr4z33wTaTXVtGzIuZG5zY3J5cHQtY2VydC50ei5kMHduLmJpeg


## d0wn-tz-ns1-ipv6

Server provided by Martin 'd0wn' Albus

sdns://AQcAAAAAAAAAGFsyYzBmOmZkYTg6NTo6MmVkMTpkMmVjXSDWBhX70UU7oU52ij5cXLOsF0YTMUY9oK-M998E2k11bRsyLmRuc2NyeXB0LWNlcnQudHouZDB3bi5iaXo


## decloudus-nogoogle-tst

Servers helps you deGoogle and unGoogle by completely blocking Google tracking in addition to annoying ads, online trackers, and malware. Supports DNSSEC. No Logs.

Contributed by: https://decloudus.com

sdns://AQMAAAAAAAAAEjE3Ni45LjE5OS4xNTg6ODQ0MyD73Ye9XeCsS7TdFu9fRP7s5k-0aL91yygulGVmeOAKLh4yLmRuc2NyeXB0LWNlcnQuRGVDbG91ZFVzLXRlc3Q


## decloudus-nogoogle-tstipv6

Servers helps you deGoogle and unGoogle by completely blocking Google tracking in addition to annoying ads, online trackers, and malware. Supports DNSSEC. No Logs. For IPv6.

Contributed by: https://decloudus.com

sdns://AQMAAAAAAAAAG1syYTAxOjRmODoxNTE6MTFiMDo6M106ODQ0MyD73Ye9XeCsS7TdFu9fRP7s5k-0aL91yygulGVmeOAKLh4yLmRuc2NyeXB0LWNlcnQuRGVDbG91ZFVzLXRlc3Q


## deffer-dns.au

DNSSEC/Non-logged/Uncensored in Sydney (AWS).
Encrypted DNS Server image by jedisct1, maintaned by DeffeR.

sdns://AQcAAAAAAAAADTUyLjY1LjIzNS4xMjkg5Q00RDDBkwx3fUaa0_etjz4iH3lLBOqsg95bYDmV07MdMi5kbnNjcnlwdC1jZXJ0LmRlZmZlci1kbnMuYXU


## dns.digitale-gesellschaft.ch

Public DoH resolver operated by the Digital Society (https://www.digitale-gesellschaft.ch).
Hosted in Zurich, Switzerland.

Non-logging, non-filtering, supports DNSSEC.

sdns://AgcAAAAAAAAADTE4NS45NS4yMTguNDKgPhoaD2xT8-l6SS1XCEtbmAcFnuBXqxUFh2_YP9o9uDigaKPrAfpfAsTTXzoR0ATK8IsiS-6n15xYjH8TyGEzYC6g70eFrucMnVdf5zODBtk-p5YR5f7c71Ni3WbJi2C7Lb2gHFQ2tlGitUFcFs_ZbRqXCf2cjyfu91z5c9mC_-aOWPAgMob_ZaZfrzIIXuoTiMNzi6fjeHPJBszjxKKLTMKliYgcZG5zLmRpZ2l0YWxlLWdlc2VsbHNjaGFmdC5jaAovZG5zLXF1ZXJ5


## dns.digitale-gesellschaft.ch-2

Public DoH resolver operated by the Digital Society (https://www.digitale-gesellschaft.ch).
Hosted in Zurich, Switzerland.

Non-logging, non-filtering, supports DNSSEC.

sdns://AgcAAAAAAAAADTE4NS45NS4yMTguNDOgPhoaD2xT8-l6SS1XCEtbmAcFnuBXqxUFh2_YP9o9uDigaKPrAfpfAsTTXzoR0ATK8IsiS-6n15xYjH8TyGEzYC6g70eFrucMnVdf5zODBtk-p5YR5f7c71Ni3WbJi2C7Lb2gHFQ2tlGitUFcFs_ZbRqXCf2cjyfu91z5c9mC_-aOWPAgMob_ZaZfrzIIXuoTiMNzi6fjeHPJBszjxKKLTMKliYgcZG5zLmRpZ2l0YWxlLWdlc2VsbHNjaGFmdC5jaAovZG5zLXF1ZXJ5


## dns.digitale-gesellschaft.ch-ipv6

Public IPv6 DoH resolver operated by the Digital Society (https://www.digitale-gesellschaft.ch).
Hosted in Zurich, Switzerland.

Non-logging, non-filtering, supports DNSSEC.

sdns://AgcAAAAAAAAAD1syYTA1OmZjODQ6OjQyXaA-GhoPbFPz6XpJLVcIS1uYBwWe4FerFQWHb9g_2j24OKBoo-sB-l8CxNNfOhHQBMrwiyJL7qfXnFiMfxPIYTNgLqDvR4Wu5wydV1_nM4MG2T6nlhHl_tzvU2LdZsmLYLstvaAcVDa2UaK1QVwWz9ltGpcJ_ZyPJ-73XPlz2YL_5o5Y8CAyhv9lpl-vMghe6hOIw3OLp-N4c8kGzOPEootMwqWJiBxkbnMuZGlnaXRhbGUtZ2VzZWxsc2NoYWZ0LmNoCi9kbnMtcXVlcnk


## dns.digitale-gesellschaft.ch-ipv6-2

Public IPv6 DoH resolver operated by the Digital Society (https://www.digitale-gesellschaft.ch).
Hosted in Zurich, Switzerland.

Non-logging, non-filtering, supports DNSSEC.

sdns://AgcAAAAAAAAAD1syYTA1OmZjODQ6OjQzXaA-GhoPbFPz6XpJLVcIS1uYBwWe4FerFQWHb9g_2j24OKBoo-sB-l8CxNNfOhHQBMrwiyJL7qfXnFiMfxPIYTNgLqDvR4Wu5wydV1_nM4MG2T6nlhHl_tzvU2LdZsmLYLstvaAcVDa2UaK1QVwWz9ltGpcJ_ZyPJ-73XPlz2YL_5o5Y8CAyhv9lpl-vMghe6hOIw3OLp-N4c8kGzOPEootMwqWJiBxkbnMuZGlnaXRhbGUtZ2VzZWxsc2NoYWZ0LmNoCi9kbnMtcXVlcnk


## dns.ryan-palmer

Non-logging, non-filtering, DNSSEC DoH Server. Hosted in the UK.

sdns://AgcAAAAAAAAADjY4LjE4My4yNTMuMjAwoD4aGg9sU_PpekktVwhLW5gHBZ7gV6sVBYdv2D_aPbg4IDKG_2WmX68yCF7qE4jDc4un43hzyQbM48Sii0zCpYmIFGRuczEucnlhbi1wYWxtZXIuY29tCi9kbnMtcXVlcnk


## dns.sb

DNSSEC-enabled DoH server by https://xtom.com/
Using Cloudflare as a frontend.

https://dns.sb

sdns://AgUAAAAAAAAAAAAKZG9oLmRucy5zYgovZG5zLXF1ZXJ5


## dns.therifleman.name

DNS-over-HTTPS DNS forwarder from Mumbai, India. Blocks web and Android trackers and ads.
IP addresses are not logged, but queries are logged for 24 hours for debugging.
Report issues, send suggestions @ joker349 at protonmail.com.
Also supports DoT (for android) @ dns.therifleman.name and plain DNS @ 172.104.206.174

sdns://AgMAAAAAAAAADzE3Mi4xMDQuMjA2LjE3NCAyhv9lpl-vMghe6hOIw3OLp-N4c8kGzOPEootMwqWJiBRkbnMudGhlcmlmbGVtYW4ubmFtZQovZG5zLXF1ZXJ5


## dnscrypt-ch-blahdns-ipv4

Blocks ad and Tracking, no Logging, DNSSEC, Hosted in Switzerland. By https://blahdns.com/

sdns://AQMAAAAAAAAAETQ1LjkwLjU3LjEyMTo4NDQzIGwtSBHT3HesNLSJKRuzyT4LBhtDvQWV_oH9ISdBP0i1GzIuZG5zY3J5cHQtY2VydC5ibGFoZG5zLmNvbQ


## dnscrypt-jp-blahdns-ipv4

Blocks ad and Tracking, no Logging, DNSSEC, Hosted in Japan. By https://blahdns.com/

sdns://AQMAAAAAAAAAEzEzOS4xNjIuMTEyLjQ3Ojg0NDMgbC1IEdPcd6w0tIkpG7PJPgsGG0O9BZX-gf0hJ0E_SLUbMi5kbnNjcnlwdC1jZXJ0LmJsYWhkbnMuY29t


## dnscrypt-jp-blahdns-ipv6

Blocks ad and Tracking, no Logging, DNSSEC, Hosted in Japan. By https://blahdns.com/

sdns://AQMAAAAAAAAAJVsyNDAwOjg5MDI6OmYwM2M6OTJmZjpmZTI3OjM0NGJdOjg0NDMgbC1IEdPcd6w0tIkpG7PJPgsGG0O9BZX-gf0hJ0E_SLUbMi5kbnNjcnlwdC1jZXJ0LmJsYWhkbnMuY29t


## dnscrypt-sg-blahdns-ipv4

Blocks ad and Tracking, no Logging, DNSSEC, Hosted in Singapore. By https://blahdns.com/

sdns://AQMAAAAAAAAAEzE5Mi41My4xNzUuMTQ5Ojg0NDMgbC1IEdPcd6w0tIkpG7PJPgsGG0O9BZX-gf0hJ0E_SLUbMi5kbnNjcnlwdC1jZXJ0LmJsYWhkbnMuY29t


## dnscrypt.be

Resolver in Leuven, Belgium (UCLL Campus Proximus). Non-logging/DNSSEC/Uncensored. https://dnscrypt.be
Maintained by Sigfried (https://sigfried.be) hosted by ISW Leuven (https://iswleuven.be).

sdns://AQcAAAAAAAAADzE5My4xOTEuMTg3LjEwNyAzWmXOT_I8k2BKJzxIJ_iYoXRQRWcR0Q1FFyrJWtvogxsyLmRuc2NyeXB0LWNlcnQuZG5zY3J5cHQuYmU


## dnscrypt.ca-1

Free, Canadian, uncensored, no-logs, encrypted, and DNSSEC validated
DNS service for your pleasure.

sdns://AQcAAAAAAAAADzE2Ny4xMTQuMjIwLjEyNSAaU6PJUHicvdELGTOkaJtshGpA8bc9F1KuysmCnst84h0yLmRuc2NyeXB0LWNlcnQuZG5zY3J5cHQuY2EtMQ


## dnscrypt.ca-1-doh

Free, Canadian, uncensored, no-logs, encrypted, and DNSSEC validated DNS service for your pleasure.

sdns://AgcAAAAAAAAADzE2Ny4xMTQuMjIwLjEyNaA-GhoPbFPz6XpJLVcIS1uYBwWe4FerFQWHb9g_2j24OCAyhv9lpl-vMghe6hOIw3OLp-N4c8kGzOPEootMwqWJiBRkbnMxLmRuc2NyeXB0LmNhOjQ1MwovZG5zLXF1ZXJ5


## dnscrypt.ca-1-ipv6

Free, Canadian, uncensored, no-logs, encrypted, and DNSSEC validated
DNS service for your pleasure.

Warning: This server is incompatible with anonymization.

sdns://AQcAAAAAAAAAJlsyNjA3OjUzMDA6NjE6OTVmOjcyODM6MTFkOTowZjg2OmU2ODldICDZGdTsPFAIzccKUeorT8v-ipk8ZDHQ7GgdXqc4RauAIjIuZG5zY3J5cHQtY2VydC5kbnNjcnlwdC5jYS0xLWlwdjY


## dnscrypt.ca-2

Free, Canadian, uncensored, no-logs, encrypted, and DNSSEC validated
DNS service for your pleasure.

sdns://AQcAAAAAAAAADTE0OS41Ni4yMjguNDUgAQhUqztWp-7505FY_vaCC_-TojV8iRYI254V07vgEYUdMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeXB0LmNhLTI


## dnscrypt.ca-2-doh

Free, Canadian, uncensored, no-logs, encrypted, and DNSSEC validated DNS service for your pleasure.

sdns://AgcAAAAAAAAADTE0OS41Ni4yMjguNDWgPhoaD2xT8-l6SS1XCEtbmAcFnuBXqxUFh2_YP9o9uDggMob_ZaZfrzIIXuoTiMNzi6fjeHPJBszjxKKLTMKliYgUZG5zMi5kbnNjcnlwdC5jYTo0NTMKL2Rucy1xdWVyeQ


## dnscrypt.ca-2-ipv6

Free, Canadian, uncensored, no-logs, encrypted, and DNSSEC validated
DNS service for your pleasure.

Warning: This server is incompatible with anonymization.

sdns://AQcAAAAAAAAAJVsyNjA3OjUzMDA6NjE6OTVmOjcyODM6MTFkOTpmODY6ZTY5MF0gmHxwqJfb2hUaNK1LVeqADvVhzASq1cV90sPYYfwX9CkiMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeXB0LmNhLTItaXB2Ng


## dnscrypt.eu-dk

Free, non-logged, uncensored. Hosted by Netgroup.

sdns://AQcAAAAAAAAADDc3LjY2Ljg0LjIzMyA3SFWF47nQiP0lrTawNwH1UgzWSJ6a3VIUV0lVnwqZVSUyLmRuc2NyeXB0LWNlcnQucmVzb2x2ZXIyLmRuc2NyeXB0LmV1


## dnscrypt.eu-dk-ipv6

Free, non-logged, uncensored. Hosted by Netgroup.

sdns://AQcAAAAAAAAAFFsyMDAxOjE0NDg6MjQzOjpkYzJdIDdIVYXjudCI_SWtNrA3AfVSDNZInprdUhRXSVWfCplVJTIuZG5zY3J5cHQtY2VydC5yZXNvbHZlcjIuZG5zY3J5cHQuZXU


## dnscrypt.eu-nl

Free, non-logged, uncensored. Hosted by RamNode.

sdns://AQcAAAAAAAAADjE3Ni41Ni4yMzcuMTcxIGfADywhxVSBRd18tGonGvLrlpkxQKMJtiuNFlMRhZxmJTIuZG5zY3J5cHQtY2VydC5yZXNvbHZlcjEuZG5zY3J5cHQuZXU


## dnscrypt.eu-nl-ipv6

Free, non-logged, uncensored. Hosted by RamNode.

sdns://AQcAAAAAAAAAGlsyYTAwOmQ4ODA6MzoxOjphNmMxOjJlODldIGfADywhxVSBRd18tGonGvLrlpkxQKMJtiuNFlMRhZxmJTIuZG5zY3J5cHQtY2VydC5yZXNvbHZlcjEuZG5zY3J5cHQuZXU


## dnscrypt.one

Non-logging, non-censoring, DNSSEC-capable DNSCrypt resolver hosted in Germany (Nuremberg), https://dnscrypt.one

sdns://AQcAAAAAAAAADTE2MS45Ny4xMDEuNTEgCjK6MsxRVt6dNByNcgmcutGvdW71yaYVnUlRM4DN6h4cMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeXB0Lm9uZQ


## dnscrypt.pl

Free | No filtering | Zero logs | DNSSEC | Poland | https://dnscrypt.pl/ | @dnscryptpl

sdns://AQcAAAAAAAAAFDE3OC4yMTYuMjAxLjEyODoyMDUzIH9hfLgepVPSNMSbwnnHT3tUmAUNHb8RGv7mmWPGR6FpGzIuZG5zY3J5cHQtY2VydC5kbnNjcnlwdC5wbA


## dnscrypt.pl-guardian

Free | Malware and phishing filtering | Zero logs | DNSSEC | Poland | https://dnscrypt.pl/ | @dnscryptpl

sdns://AQMAAAAAAAAAFDE3OC4yMTYuMjAxLjEyODoyMDU0IH9hfLgepVPSNMSbwnnHT3tUmAUNHb8RGv7mmWPGR6FpGzIuZG5zY3J5cHQtY2VydC5kbnNjcnlwdC5wbA


## dnscrypt.uk-ipv4

DNSCrypt, no logs, uncensored, DNSSEC. Hosted in London UK on Digital Ocean
https://www.dnscrypt.uk

sdns://AQcAAAAAAAAADTE2NS4yMzIuMzIuOTUgAdMO6A4gTdoYcdOhDyhRVGdlvxS02kQEpP8EbX2paDwbMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeXB0LnVr


## dnscrypt.uk-ipv6

DNSCrypt, no logs, uncensored, DNSSEC. Hosted in London UK on Digital Ocean
https://www.dnscrypt.uk

sdns://AQcAAAAAAAAAGlsyYTAzOmIwYzA6MTplMDo6NDg3OjEwMDFdIAHTDugOIE3aGHHToQ8oUVRnZb8UtNpEBKT_BG19qWg8GzIuZG5zY3J5cHQtY2VydC5kbnNjcnlwdC51aw


## dnsforfamily

(DNSCrypt Protocol) Block adult websites, gambling websites, malwares and advertisements.
It also enforces safe search in: Google, YouTube, Bing, DuckDuckGo and Yandex.

Social websites like Facebook and Instagram are not blocked. No DNS queries are logged.

As of December 2020 2.7 million websites are blocked and new websites are added to blacklist daily.
Completely free, no ads or any commercial motive. Operating for 3 years now.

Warning: This server is incompatible with anonymization.

Provided by: https://dnsforfamily.com

sdns://AQIAAAAAAAAADDc4LjQ3LjY0LjE2MSATJeLOABXNSYcSJIoqR5_iUYz87Y4OecMLB84aEAKPrRBkbnNmb3JmYW1pbHkuY29t


## dnsforfamily-doh

(DoH Protocol) Block adult websites, gambling websites, malwares and advertisements.
It also enforces safe search in: Google, YouTube, Bing, DuckDuckGo and Yandex.
Social websites like Facebook and Instagram are not blocked. No DNS queries are logged.
As of December 2020 2.7 million websites are blocked and new websites are added to blacklist daily.
Completely free, no ads or any commercial motive. Operating for 3 years now.

Provided by: https://dnsforfamily.com

sdns://AgIAAAAAAAAADTk1LjIxNy4yMTMuOTSgPhoaD2xT8-l6SS1XCEtbmAcFnuBXqxUFh2_YP9o9uDggMob_ZaZfrzIIXuoTiMNzi6fjeHPJBszjxKKLTMKliYgYZG5zLWRvaC5kbnNmb3JmYW1pbHkuY29tCi9kbnMtcXVlcnk


## dnsforfamily-doh-no-safe-search

(DoH Protocol) Block adult websites, gambling websites, malwares and advertisements.
Unlike other dnsforfamily DNSCrypt servers, this one does not enforces safe search. So Google, YouTube, Bing, DuckDuckGo and Yandex are completely accessible without any restriction.

Social websites like Facebook and Instagram are not blocked. No DNS queries are logged.

As of December 2020 2.7 million websites are blocked and new websites are added to blacklist daily.
Completely free, no ads or any commercial motive. Operating for 3 years now.

Warning: This server is incompatible with anonymization.

Provided by: https://dnsforfamily.com

sdns://AgIAAAAAAAAADTk1LjIxNy4yMTMuOTQgMob_ZaZfrzIIXuoTiMNzi6fjeHPJBszjxKKLTMKliYgnZG5zLWRvaC1uby1zYWZlLXNlYXJjaC5kbnNmb3JmYW1pbHkuY29tCi9kbnMtcXVlcnk


## dnsforfamily-no-safe-search

(DNSCrypt Protocol) Block adult websites, gambling websites, malwares and advertisements.
Unlike other dnsforfamily DNSCrypt servers, this one does not enforces safe search. So Google, YouTube, Bing, DuckDuckGo and Yandex are completely accessible without any restriction.

Social websites like Facebook and Instagram are not blocked. No DNS queries are logged.

As of December 2020 2.7 million websites are blocked and new websites are added to blacklist daily.
Completely free, no ads or any commercial motive. Operating for 3 years now.

Warning: This server is incompatible with anonymization.

Provided by: https://dnsforfamily.com

sdns://AQIAAAAAAAAADzEzNS4xODEuMTkzLjIyMiBHFKrWl_Swzwd8Mcwa8ZhdLGFgC94SpKo_g57e_49DthBkbnNmb3JmYW1pbHkuY29t


## dnsforfamily-v6

(DNSCrypt Protocol) Block adult websites, gambling websites, malwares and advertisements.
It also enforces safe search in: Google, YouTube, Bing, DuckDuckGo and Yandex.

Social websites like Facebook and Instagram are not blocked. No DNS queries are logged.
As of December 2020 2.7 million websites are blocked and new websites are added to blacklist daily.
Completely free, no ads or any commercial motive. Operating for 3 years now.

Provided by: https://dnsforfamily.com

sdns://AQIAAAAAAAAAF1syYTAxOjRmODoxYzE3OjRkZjg6OjFdIGN4CrSY4fb2hK8voFJL3GKiM7xQNwkKGH4b0k7LmMPxEGRuc2ZvcmZhbWlseS5jb20


## dnsforge.de

Public DoH resolver running with Pihole for Adblocking (https://dnsforge.de).

Non-logging, AD-filtering, supports DNSSEC. Hosted in Germany.

sdns://AgMAAAAAAAAADDE3Ni45LjkzLjE5OKA-GhoPbFPz6XpJLVcIS1uYBwWe4FerFQWHb9g_2j24OCAyhv9lpl-vMghe6hOIw3OLp-N4c8kGzOPEootMwqWJiAtkbnNmb3JnZS5kZQovZG5zLXF1ZXJ5


## dnshome-doh

https://www.dnshome.de/ public resolver in Germany

sdns://AgYAAAAAAAAAAKA-GhoPbFPz6XpJLVcIS1uYBwWe4FerFQWHb9g_2j24OCAyhv9lpl-vMghe6hOIw3OLp-N4c8kGzOPEootMwqWJiA5kbnMuZG5zaG9tZS5kZQovZG5zLXF1ZXJ5


## dnspod-doh

A public DNS resolver in mainland China provided by DNSPod (Tencent Cloud).

https://www.dnspod.cn/Products/Public.DNS?lang=en

sdns://AgAAAAAAAAAAACDrdSX4jw2UWPgamVAZv9NMuJzNyVfnsO8xXxD4l2OBGAdkb2gucHViCi9kbnMtcXVlcnk


## doh-ch-blahdns

Blocks ad and Tracking, no Logging, DNSSEC, Hosted in Switzerland. By https://blahdns.com/

sdns://AgMAAAAAAAAADDQ1LjkwLjU3LjEyMQAXZG9oLWNoLmJsYWhkbnMuY29tOjQ0NDMKL2Rucy1xdWVyeQ


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

sdns://AgcAAAAAAAAADDEwNC4yOC4wLjEwNgANZG9oLmNyeXB0by5zeAovZG5zLXF1ZXJ5
sdns://AgcAAAAAAAAADDEwNC4yOC4xLjEwNgANZG9oLmNyeXB0by5zeAovZG5zLXF1ZXJ5


## doh-crypto-sx-ipv6

DNS-over-HTTPS server accessible over IPv6. Anycast, no logs, no censorship, DNSSEC.
Backend hosted by Scaleway, globally cached via Cloudflare.
Maintained by Frank Denis.

sdns://AgcAAAAAAAAAGVsyNjA2OjQ3MDA6MzAzNjo6NjgxYzo2YV0AEmRvaC1pcHY2LmNyeXB0by5zeAovZG5zLXF1ZXJ5
sdns://AgcAAAAAAAAAGlsyNjA2OjQ3MDA6MzAzNDo6NjgxYzoxNmFdABJkb2gtaXB2Ni5jcnlwdG8uc3gKL2Rucy1xdWVyeQ


## doh-de-blahdns

Blocks ad and Tracking, no Logging, DNSSEC, Hosted in Germany. By https://blahdns.com/

sdns://AgMAAAAAAAAADTc4LjQ2LjI0NC4xNDMAEmRvaC1kZS5ibGFoZG5zLmNvbQovZG5zLXF1ZXJ5


## doh-de-blahdns-v6

Blocks ad and Tracking, no Logging, DNSSEC, Hosted in Germany. By https://blahdns.com/

sdns://AgMAAAAAAAAAFlsyYTAxOjRmODpjMTc6ZWM2Nzo6MV0AEmRvaC1kZS5ibGFoZG5zLmNvbQovZG5zLXF1ZXJ5


## doh-fi-blahdns

Blocks ad and Tracking, no Logging, DNSSEC, Hosted in Finland. By https://blahdns.com/

sdns://AgMAAAAAAAAADjk1LjIxNi4yMTIuMTc3ABJkb2gtZmkuYmxhaGRucy5jb20KL2Rucy1xdWVyeQ


## doh-fi-blahdns-v6

Blocks ad and Tracking, no Logging, DNSSEC, Hosted in Finland. By https://blahdns.com/

sdns://AgMAAAAAAAAAF1syYTAxOjRmOTpjMDEwOjQzY2U6OjFdABJkb2gtZmkuYmxhaGRucy5jb20KL2Rucy1xdWVyeQ


## doh-jp-blahdns

Blocks ad and Tracking, no Logging, DNSSEC, Hosted in Japan. By https://blahdns.com/

sdns://AgMAAAAAAAAADjEzOS4xNjIuMTEyLjQ3ABJkb2gtanAuYmxhaGRucy5jb20KL2Rucy1xdWVyeQ


## doh-jp-blahdns-v6

Blocks ad and Tracking, no Logging, DNSSEC, Hosted in Japan. By https://blahdns.com/

sdns://AgMAAAAAAAAAIFsyNDAwOjg5MDI6OmYwM2M6OTJmZjpmZTI3OjM0NGJdABJkb2gtanAuYmxhaGRucy5jb20KL2Rucy1xdWVyeQ


## doh-sg-blahdns

Blocks ad and Tracking, no Logging, DNSSEC, Hosted in Singapore. By https://blahdns.com/

sdns://AgMAAAAAAAAADjE5Mi41My4xNzUuMTQ5ABJkb2gtc2cuYmxhaGRucy5jb20KL2Rucy1xdWVyeQ


## doh.appliedprivacy.net

Public DoH resolver operated by the Foundation for Applied Privacy (https://appliedprivacy.net).
Hosted in Vienna, Austria.

Non-logging, non-filtering, supports DNSSEC.

sdns://AgcAAAAAAAAAAKA-GhoPbFPz6XpJLVcIS1uYBwWe4FerFQWHb9g_2j24OKBoo-sB-l8CxNNfOhHQBMrwiyJL7qfXnFiMfxPIYTNgLqDvR4Wu5wydV1_nM4MG2T6nlhHl_tzvU2LdZsmLYLstvaAcVDa2UaK1QVwWz9ltGpcJ_ZyPJ-73XPlz2YL_5o5Y8CAyhv9lpl-vMghe6hOIw3OLp-N4c8kGzOPEootMwqWJiBZkb2guYXBwbGllZHByaXZhY3kubmV0Bi9xdWVyeQ


## doh.ffmuc.net

An open (non-logging, non-filtering, non-censoring) DoH resolver operated by Freifunk Munich with nodes in DE.
https://ffmuc.net/

sdns://AgcAAAAAAAAACjUuMS42Ni4yNTWgPhoaD2xT8-l6SS1XCEtbmAcFnuBXqxUFh2_YP9o9uDggMob_ZaZfrzIIXuoTiMNzi6fjeHPJBszjxKKLTMKliYgNZG9oLmZmbXVjLm5ldAovZG5zLXF1ZXJ5


## doh.ffmuc.net-v6

An open (non-logging, non-filtering, non-censoring) DoH resolver operated by Freifunk Munich with nodes in DE.
https://ffmuc.net/

sdns://AgcAAAAAAAAAFVsyMDAxOjY3ODplNjg6ZjAwMDo6XaA-GhoPbFPz6XpJLVcIS1uYBwWe4FerFQWHb9g_2j24OCAyhv9lpl-vMghe6hOIw3OLp-N4c8kGzOPEootMwqWJiA1kb2guZmZtdWMubmV0Ci9kbnMtcXVlcnk


## doh.li

Public DoH server in London, UK.

Warning: provides EDNS Client Subnet information to upstream servers.

sdns://AgUAAAAAAAAAAKA-GhoPbFPz6XpJLVcIS1uYBwWe4FerFQWHb9g_2j24OCAyhv9lpl-vMghe6hOIw3OLp-N4c8kGzOPEootMwqWJiAZkb2gubGkKL2Rucy1xdWVyeQ


## doh.tiarap.org

Non-Logging DNS-over-HTTPS server, cached via Cloudflare.
Filters out ads, trackers and malware, NO ECS, supports DNSSEC.

sdns://AgMAAAAAAAAADDEwNC4yOC4yOC4zNCDlfOUtFRBtpOz9nhH9pf0dHgpr4BIkGITTYuodSvRk9w5kb2gudGlhcmFwLm9yZwovZG5zLXF1ZXJ5


## doh.tiarap.org-ipv6

Non-Logging DNS-over-HTTPS server (IPv6), cached via Cloudflare.
Filters out ads, trackers and malware, NO ECS, supports DNSSEC.

sdns://AgMAAAAAAAAAG1syNjA2OjQ3MDA6MzAzNDo6YWM0MzpiNjZhXSDlfOUtFRBtpOz9nhH9pf0dHgpr4BIkGITTYuodSvRk9w5kb2gudGlhcmFwLm9yZwovZG5zLXF1ZXJ5


## east.comss.one

DNS with adblock filters and antiphishing, gaining popularity among russian-speaking users.

sdns://AQMAAAAAAAAAEjkxLjIzMC4yMTEuNjc6NTQ0MyAVacgDhm0kq-ciz2DrJlPhYjI9v6tAQHbulojEz86TEyIyLmRuc2NyeXB0LWNlcnQuZG5zLmVhc3QuY29tc3Mub25l


## faelix-ch-ipv4

An open (non-logging, non-filtering, no ECS) DNSCrypt resolver operated by https://faelix.net/ with IPv4 nodes anycast within AS41495 in Switzerland.

sdns://AQcAAAAAAAAAEzE4NS4xMzQuMTk2LjU0Ojg0NDMgfsvvPi8BgDKNYODh0ewj5Oh32OoJoZNwGgTWs8C-i-EfMi5kbnNjcnlwdC1jZXJ0LnJkbnMuZmFlbGl4Lm5ldA


## faelix-ch-ipv4-doh

An open (non-logging, non-filtering, no ECS) DoH resolver operated by https://faelix.net/ with IPv4 nodes anycast within AS41495 in Switzerland.

sdns://AgcAAAAAAAAADjE4NS4xMzQuMTk2LjU0oD4aGg9sU_PpekktVwhLW5gHBZ7gV6sVBYdv2D_aPbg4IDKG_2WmX68yCF7qE4jDc4un43hzyQbM48Sii0zCpYmID3JkbnMuZmFlbGl4Lm5ldAEv


## faelix-ch-ipv6

An open (non-logging, non-filtering, no ECS) DNSCrypt resolver operated by https://faelix.net/ with IPv6 nodes anycast within AS41495 in Switzerland.

sdns://AQcAAAAAAAAAFFsyYTAxOjllMDE6OjU0XTo4NDQzIH7L7z4vAYAyjWDg4dHsI-Tod9jqCaGTcBoE1rPAvovhHzIuZG5zY3J5cHQtY2VydC5yZG5zLmZhZWxpeC5uZXQ


## faelix-ch-ipv6-doh

An open (non-logging, non-filtering, no ECS) DoH resolver operated by https://faelix.net/ with IPv6 nodes anycast within AS41495 in Switzerland.

sdns://AgcAAAAAAAAAD1syYTAxOjllMDE6OjU0XaA-GhoPbFPz6XpJLVcIS1uYBwWe4FerFQWHb9g_2j24OCAyhv9lpl-vMghe6hOIw3OLp-N4c8kGzOPEootMwqWJiA9yZG5zLmZhZWxpeC5uZXQBLw


## faelix-uk-ipv4

An open (non-logging, non-filtering, no ECS) DNSCrypt resolver operated by https://faelix.net/ with IPv4 nodes anycast within AS41495 in the UK.

sdns://AQcAAAAAAAAAEjQ2LjIyNy4yMDAuNTQ6ODQ0MyB-y-8-LwGAMo1g4OHR7CPk6HfY6gmhk3AaBNazwL6L4R8yLmRuc2NyeXB0LWNlcnQucmRucy5mYWVsaXgubmV0


## faelix-uk-ipv4-doh

An open (non-logging, non-filtering, no ECS) DoH resolver operated by https://faelix.net/ with IPv4 nodes anycast within AS41495 in the UK.

sdns://AgcAAAAAAAAADTQ2LjIyNy4yMDAuNTSgPhoaD2xT8-l6SS1XCEtbmAcFnuBXqxUFh2_YP9o9uDggMob_ZaZfrzIIXuoTiMNzi6fjeHPJBszjxKKLTMKliYgPcmRucy5mYWVsaXgubmV0AS8


## faelix-uk-ipv6

An open (non-logging, non-filtering, no ECS) DNSCrypt resolver operated by https://faelix.net/ with IPv6 nodes anycast within AS41495 in the UK.

sdns://AQcAAAAAAAAAFFsyYTAxOjllMDA6OjU0XTo4NDQzIH7L7z4vAYAyjWDg4dHsI-Tod9jqCaGTcBoE1rPAvovhHzIuZG5zY3J5cHQtY2VydC5yZG5zLmZhZWxpeC5uZXQ


## faelix-uk-ipv6-doh

An open (non-logging, non-filtering, no ECS) DoH resolver operated by https://faelix.net/ with IPv6 nodes anycast within AS41495 in the UK.

sdns://AgcAAAAAAAAAD1syYTAxOjllMDA6OjU0XaA-GhoPbFPz6XpJLVcIS1uYBwWe4FerFQWHb9g_2j24OCAyhv9lpl-vMghe6hOIw3OLp-N4c8kGzOPEootMwqWJiA9yZG5zLmZhZWxpeC5uZXQBLw


## ffmuc.net

An open (non-logging, non-filtering, non-censoring) DNSCrypt resolver operated by Freifunk Munich with nodes in DE.
https://ffmuc.net/

sdns://AQcAAAAAAAAADzUuMS42Ni4yNTU6ODQ0MyAH0Hrxz9xdmXadPwJmkKcESWXCdCdseRyu9a7zuQxG-hkyLmRuc2NyeXB0LWNlcnQuZmZtdWMubmV0


## ffmuc.net-v6

An open (non-logging, non-filtering, non-censoring) DNSCrypt resolver operated by Freifunk Munich with nodes in DE.
https://ffmuc.net/

sdns://AQcAAAAAAAAAGlsyMDAxOjY3ODplNjg6ZjAwMDo6XTo4NDQzIAfQevHP3F2Zdp0_AmaQpwRJZcJ0J2x5HK71rvO5DEb6GTIuZG5zY3J5cHQtY2VydC5mZm11Yy5uZXQ


## freetsa.org-ipv6

Non-logged/Uncensored provided by www.freetsa.org. Support for DNS and DNS-over-TLS (DoT)

sdns://AQcAAAAAAAAAH1syNjA3OmYxMzA6MDpmODo6MzA4NTplOTYxXTo1NTMg2P-7QuAxvnp5cwtFVo1Jak6Ky1mqg2b9arkeJyp9FuQbMi5kbnNjcnlwdC1jZXJ0LmZyZWV0c2Eub3Jn


## google

Google DNS (anycast)

sdns://AgUAAAAAAAAABzguOC44LjigHvYkz_9ea9O63fP92_3qVlRn43cpncfuZnUWbzAMwbkgdoAkR6AZkxo_AEMExT_cbBssN43Evo9zs5_ZyWnftEUKZG5zLmdvb2dsZQovZG5zLXF1ZXJ5


## google-ipv6

Google DNS (anycast)

sdns://AgUAAAAAAAAAFlsyMDAxOjQ4NjA6NDg2MDo6ODg4OF2gHvYkz_9ea9O63fP92_3qVlRn43cpncfuZnUWbzAMwbkgdoAkR6AZkxo_AEMExT_cbBssN43Evo9zs5_ZyWnftEUKZG5zLmdvb2dsZQovZG5zLXF1ZXJ5


## he

Hurricane Electric DoH server (anycast)

Unknown logging policy.

sdns://AgUAAAAAAAAACzc0LjgyLjQyLjQyoD4aGg9sU_PpekktVwhLW5gHBZ7gV6sVBYdv2D_aPbg4IDKG_2WmX68yCF7qE4jDc4un43hzyQbM48Sii0zCpYmIDG9yZG5zLmhlLm5ldAovZG5zLXF1ZXJ5


## id-gmail

Non-Logging DNSCrypt server located in Singapore.
Filters out ads, trackers and malware, supports DNSSEC, provided by id-gmail.

sdns://AQMAAAAAAAAADjE3NC4xMzguMjEuMTI4IO-WgGbo2ZTwZdg-3dMa7u31bYZXRj5KykfN1_6Xw9T2HDIuZG5zY3J5cHQtY2VydC5kbnMudGlhci5hcHA


## id-gmail-doh

Non-Logging DNS-over-HTTPS server located in Singapore.
Filters out ads, trackers and malware, supports DNSSEC, provided by id-gmail.

sdns://AgMAAAAAAAAADjE3NC4xMzguMjkuMTc1oD4aGg9sU_PpekktVwhLW5gHBZ7gV6sVBYdv2D_aPbg4IDKG_2WmX68yCF7qE4jDc4un43hzyQbM48Sii0zCpYmIDGRvaC50aWFyLmFwcAovZG5zLXF1ZXJ5


## id-gmail-doh-ipv6

Non-Logging DNS-over-HTTPS (IPv6) server located in Singapore.
Filters out ads, trackers and malware, supports DNSSEC, provided by id-gmail.

sdns://AgMAAAAAAAAAG1syNDAwOjYxODA6MDpkMDo6NWY3Mzo0MDAxXaA-GhoPbFPz6XpJLVcIS1uYBwWe4FerFQWHb9g_2j24OCAyhv9lpl-vMghe6hOIw3OLp-N4c8kGzOPEootMwqWJiAxkb2gudGlhci5hcHAKL2Rucy1xdWVyeQ


## id-gmail-ipv6

Non-Logging DNSCrypt (IPv6) server located in Singapore.
Filters out ads, trackers and malware, supports DNSSEC, provided by id-gmail.

sdns://AQMAAAAAAAAAG1syNDAwOjYxODA6MDpkMDo6NWY2ZTo0MDAxXSDvloBm6NmU8GXYPt3TGu7t9W2GV0Y-SspHzdf-l8PU9hwyLmRuc2NyeXB0LWNlcnQuZG5zLnRpYXIuYXBw


## iij

DoH server operated by Internet Initiative Japan in Tokyo.
https://www.iij.ad.jp/

sdns://AgUAAAAAAAAACjEwMy4yLjU3LjUgmOPV5TavKVjNL38U9wTvSidtJeM81l8uZfXk8nJ8EzARcHVibGljLmRucy5paWouanAKL2Rucy1xdWVyeQ


## indonesia-unfilter-doh

Non-logging and non-Filtering DoH forwarder in Indonesia.

sdns://AgYAAAAAAAAACjExNi4wLjIuODYgMob_ZaZfrzIIXuoTiMNzi6fjeHPJBszjxKKLTMKliYgPZG9oLnRoaXMud2ViLmlkCi9kbnMtcXVlcnk


## jitender

DNSCrypt resolver in Hyderabad, India, Oracle Cloud

Warning: this server is incompatible with DNS anonymization.

Operated by Jitender (@coolquasar) in Hyderabad, India.
Blocks ads/trackers/analytics domains and provides a clean browsing experience.

sdns://AQMAAAAAAAAADTE1Mi42Ny4xNjUuMjYgBQbmyjgl5vyzNCe7s6WXJck9ZZR425-NzDAP63Xij5AcMi5kbnNjcnlwdC1jZXJ0LmppdC5kZG5zLm5ldA


## jp.tiar.app

Non-Logging, Non-Filtering DNSCrypt server in Japan.
No ECS, Support DNSSEC

sdns://AQcAAAAAAAAAEjE3Mi4xMDQuOTMuODA6MTQ0MyAyuHY-8b9lNqHeahPAzW9IoXnjiLaZpTeNbVs8TN9UUxsyLmRuc2NyeXB0LWNlcnQuanAudGlhci5hcHA


## jp.tiar.app-doh

Non-Logging, Non-Filtering DNS-over-HTTPS server in Japan.
No ECS, Support DNSSEC

sdns://AgcAAAAAAAAADTE3Mi4xMDQuOTMuODCgPhoaD2xT8-l6SS1XCEtbmAcFnuBXqxUFh2_YP9o9uDggMob_ZaZfrzIIXuoTiMNzi6fjeHPJBszjxKKLTMKliYgLanAudGlhci5hcHAKL2Rucy1xdWVyeQ


## jp.tiar.app-doh-ipv6

Non-Logging, Non-Filtering DNS-over-HTTPS (IPv6) server in Japan.
No ECS, Support DNSSEC

sdns://AgcAAAAAAAAAIFsyNDAwOjg5MDI6OmYwM2M6OTFmZjpmZWRhOmM1MTRdoD4aGg9sU_PpekktVwhLW5gHBZ7gV6sVBYdv2D_aPbg4IDKG_2WmX68yCF7qE4jDc4un43hzyQbM48Sii0zCpYmIC2pwLnRpYXIuYXBwCi9kbnMtcXVlcnk


## jp.tiar.app-ipv6

Non-Logging, Non-Filtering DNSCrypt (IPv6) server in Japan.
No ECS, Support DNSSEC

sdns://AQcAAAAAAAAAJVsyNDAwOjg5MDI6OmYwM2M6OTFmZjpmZWRhOmM1MTRdOjE0NDMgMrh2PvG_ZTah3moTwM1vSKF544i2maU3jW1bPEzfVFMbMi5kbnNjcnlwdC1jZXJ0LmpwLnRpYXIuYXBw


## jp.tiarap.org

DNS-over-HTTPS Server. Non-Logging, Non-Filtering, No ECS, Support DNSSEC.
Cached via Cloudflare.

sdns://AgcAAAAAAAAADDEwNC4yOC4yOC4zNCDlfOUtFRBtpOz9nhH9pf0dHgpr4BIkGITTYuodSvRk9w1qcC50aWFyYXAub3JnCi9kbnMtcXVlcnk


## jp.tiarap.org-ipv6

DNS-over-HTTPS Server (IPv6). Non-Logging, Non-Filtering, No ECS, Support DNSSEC.
Cached via Cloudflare.

sdns://AgcAAAAAAAAAG1syNjA2OjQ3MDA6MzAzNjo6NjgxYjo5NmFhXSDlfOUtFRBtpOz9nhH9pf0dHgpr4BIkGITTYuodSvRk9w1qcC50aWFyYXAub3JnCi9kbnMtcXVlcnk


## libredns

DoH server in Germany. No logging, but no DNS padding and no DNSSEC support.
https://libredns.gr/

sdns://AgYAAAAAAAAAAKA-GhoPbFPz6XpJLVcIS1uYBwWe4FerFQWHb9g_2j24OCAyhv9lpl-vMghe6hOIw3OLp-N4c8kGzOPEootMwqWJiA9kb2gubGlicmVkbnMuZ3IKL2Rucy1xdWVyeQ


## libredns-noads

DoH server in Germany. No logging, but no DNS padding and no DNSSEC support.
no ads version, uses StevenBlack's host list: https://github.com/StevenBlack/hosts

sdns://AgIAAAAAAAAAAKA-GhoPbFPz6XpJLVcIS1uYBwWe4FerFQWHb9g_2j24OCAyhv9lpl-vMghe6hOIw3OLp-N4c8kGzOPEootMwqWJiA9kb2gubGlicmVkbnMuZ3IEL2Fkcw


## linuxsec-adguard

DoH server with AdGuard Home for ads blocking.
Operated by LinuxSec.

Non-logging, AD-filtering, supports DNSSEC.

https://doh.linuxsec.org

sdns://AgMAAAAAAAAADDEwMy44NS4xNS42MCAyhv9lpl-vMghe6hOIw3OLp-N4c8kGzOPEootMwqWJiBBkb2gubGludXhzZWMub3JnCC9hZGd1YXJk


## linuxsec-doh

DoH server operated by LinuxSec.
Server located in the Indonesia.

Non-logging, supports DNSSEC.

https://doh.linuxsec.org

sdns://AgcAAAAAAAAADDEwMy44NS4xNS42MCAyhv9lpl-vMghe6hOIw3OLp-N4c8kGzOPEootMwqWJiBBkb2gubGludXhzZWMub3JnCi9kbnMtcXVlcnk


## meganerd

DNSCrypt server by MegaNerd.nl (IPv4) - https://meganerd.nl/encrypted-dns-server
Hosted in Amsterdam (AMS1), The Netherlands.

Non-logging, non-filtering, supports DNSSEC.

sdns://AQcAAAAAAAAADjEzNi4yNDQuOTcuMTE0IPyq3HBOXuNgu6FO4pU71Si6CTV6kPD85NA6AThr_6tiGDIuZG5zY3J5cHQtY2VydC5tZWdhbmVyZA


## meganerd-doh-ipv6

DoH server by MegaNerd.nl (IPv6) - https://meganerd.nl/encrypted-dns-server
Hosted in Amsterdam (AMS1), The Netherlands.

Non-logging, non-filtering, supports DNSSEC.

sdns://AgcAAAAAAAAAKFsyMDAxOjE5ZjA6NTAwMTpjYmI6NTQwMDowM2ZmOmZlMDc6ZjcwZF0gMob_ZaZfrzIIXuoTiMNzi6fjeHPJBszjxKKLTMKliYgZY2hld2JhY2NhLm1lZ2FuZXJkLm5sOjQ0MwQvZG9o


## meganerd-ipv6

DNSCrypt server by MegaNerd.nl (IPv6) - https://meganerd.nl/encrypted-dns-server
Hosted in Amsterdam (AMS1), The Netherlands.

Non-logging, non-filtering, supports DNSSEC.

sdns://AQcAAAAAAAAAJ1syMDAxOjE5ZjA6NTAwMTpjYmI6NTQwMDozZmY6ZmUwNzpmNzBkXSD8qtxwTl7jYLuhTuKVO9Uougk1epDw_OTQOgE4a_-rYhgyLmRuc2NyeXB0LWNlcnQubWVnYW5lcmQ


## moulticast-ca-ipv4

Public | Non-filtering | Non-logging | DNSSEC aware | Hosted in Canada | Operated by @herver (Github) | https://moulticast.net/dnscrypt/

sdns://AQcAAAAAAAAADTE0OS41Ni4xNC4xNTkggqVNQrswzoCc-93hdnomI0UwlYC3q80lR9-79MEC3xUdMi5kbnNjcnlwdC1jZXJ0Lm1vdWx0aWNhc3QtY2E


## moulticast-ca-ipv6

Public | Non-filtering | Non-logging | DNSSEC aware | Hosted in Canada | Operated by @herver (Github) | https://moulticast.net/dnscrypt/

sdns://AQcAAAAAAAAAGlsyNjA3OjUzMDA6MjAxOjMxMDA6OjNkZmJdIIKlTUK7MM6AnPvd4XZ6JiNFMJWAt6vNJUffu_TBAt8VHTIuZG5zY3J5cHQtY2VydC5tb3VsdGljYXN0LWNh


## moulticast-de-ipv4

Public | Non-filtering | Non-logging | DNSSEC aware | Hosted in Germany | Operated by @herver (Github) | https://moulticast.net/dnscrypt/

sdns://AQcAAAAAAAAADTUxLjE5NS4xMTcuMjYgI0N2y9-Xq_-5VbF43rxF6u10RV5LphBpRaTRNZuAOnAdMi5kbnNjcnlwdC1jZXJ0Lm1vdWx0aWNhc3QtZGU


## moulticast-doh-ipv6

Public DoH | Non-filtering | Non-logging | DNSSEC aware | Anycast | Operated by @herver (Github) | https://moulticast.net/dnscrypt/

sdns://AgcAAAAAAAAAFVsyMDAxOjY3ODpkMDg6ZmY6OjUzXQASZG5zLm1vdWx0aWNhc3QubmV0Ci9kbnMtcXVlcnk


## moulticast-fr-ipv4

Public | Non-filtering | Non-logging | DNSSEC aware | Hosted in France | Operated by @herver (Github) | https://moulticast.net/dnscrypt/

sdns://AQcAAAAAAAAADTE5My4yMDAuNDMuMjUgdahEBxLVMuwDPeuVlJbOFrrDypHdrAlyIm7FrDgPyMgdMi5kbnNjcnlwdC1jZXJ0Lm1vdWx0aWNhc3QtZnI


## moulticast-ipv6

Public | Non-filtering | Non-logging | DNSSEC aware | Anycast | Operated by @herver (Github) | https://moulticast.net/dnscrypt/

sdns://AQcAAAAAAAAAFVsyMDAxOjY3ODpkMDg6ZmY6OjU0XSA71VxH3ImIQdj36bGmC0PNbtpWJwAvNiYqBqtFG9XvfCcyLmRuc2NyeXB0LWNlcnQuZG5zY3J5cHQubW91bHRpY2FzdC5uZXQ


## moulticast-sg-ipv4

Public | Non-filtering | Non-logging | DNSSEC aware | Hosted in Singapore | Operated by @herver (Github) | https://moulticast.net/dnscrypt/

sdns://AQcAAAAAAAAADTUxLjc5LjE1OC4xODMgL4fe2yQRI0ljCMe05zhSmeKGYz-izKVjyE040Ovq-gUdMi5kbnNjcnlwdC1jZXJ0Lm1vdWx0aWNhc3Qtc2c


## moulticast-sg-ipv6

Public | Non-filtering | Non-logging | DNSSEC aware | Hosted in Singapore | Operated by @herver (Github) | https://moulticast.net/dnscrypt/

sdns://AQcAAAAAAAAAGlsyNDAyOjFmMDA6ODAwMDo4MDA6OjEyYmFdIC-H3tskESNJYwjHtOc4UpnihmM_osylY8hNONDr6voFHTIuZG5zY3J5cHQtY2VydC5tb3VsdGljYXN0LXNn


## moulticast-uk-ipv4

Public | Non-filtering | Non-logging | DNSSEC aware | Hosted in UK | Operated by @herver (Github) | https://moulticast.net/dnscrypt/

sdns://AQcAAAAAAAAADjUxLjE5NS4yMDAuMTgyIOdFe4OLrR_kCKC-8omGs5my5qxIyBgkldWZoSUmYvCNHTIuZG5zY3J5cHQtY2VydC5tb3VsdGljYXN0LXVr


## nextdns

NextDNS is a cloud-based private DNS service that gives you full control
over what is allowed and what is blocked on the Internet.

DNSSEC, Anycast, Non-logging, NoFilters

https://www.nextdns.io/

sdns://AgcAAAAAAAAACjQ1LjkwLjI4LjCgPhoaD2xT8-l6SS1XCEtbmAcFnuBXqxUFh2_YP9o9uDggMob_ZaZfrzIIXuoTiMNzi6fjeHPJBszjxKKLTMKliYgOZG5zLm5leHRkbnMuaW8PL2Ruc2NyeXB0LXByb3h5


## nextdns-ipv6

Connects to NextDNS over IPv6.

NextDNS is a cloud-based private DNS service that gives you full control
over what is allowed and what is blocked on the Internet.

DNSSEC, Anycast, Non-logging, NoFilters

https://www.nextdns.io/

sdns://AgcAAAAAAAAADVsyYTA3OmE4YzA6Ol2gPhoaD2xT8-l6SS1XCEtbmAcFnuBXqxUFh2_YP9o9uDggMob_ZaZfrzIIXuoTiMNzi6fjeHPJBszjxKKLTMKliYgOZG5zLm5leHRkbnMuaW8PL2Ruc2NyeXB0LXByb3h5


## njalla-doh

Non-logging DoH server in Sweden operated by Njalla.

https://dns.njal.la/

sdns://AgcAAAAAAAAADDk1LjIxNS4xOS41M6A-GhoPbFPz6XpJLVcIS1uYBwWe4FerFQWHb9g_2j24OCAyhv9lpl-vMghe6hOIw3OLp-N4c8kGzOPEootMwqWJiAtkbnMubmphbC5sYQovZG5zLXF1ZXJ5


## opennic-fische

OpenNIC • Non-logging • DNSSEC
Location: Nurnberg, Germany

sdns://AQcAAAAAAAAAEDc4LjQ3LjI0My4zOjEwNTMgN4CAbUDR-b3uJJMVzfCdL9ivVV7s8wRhifLRPWBfSmQdMi5kbnNjcnlwdC1jZXJ0Lm5zMS5maXNjaGUuaW8


## opennic-luggs

Public DNS server in Canada operated by Luggs

sdns://AQYAAAAAAAAADTE0Mi40LjIwNC4xMTEgHBl5MxvoI8zPCJp5BpN-XDQQKlasf2Jw4EYlsu3bBOMfMi5kbnNjcnlwdC1jZXJ0Lm5zMy5jYS5sdWdncy5jbw


## opennic-luggs2

Second public DNS server in Canada operated by Luggs

sdns://AQYAAAAAAAAADDE0Mi40LjIwNS40NyC8v7fgUME9okIsALCxrJrWSMXCZLy2FwuIPXGKyG66CR8yLmRuc2NyeXB0LWNlcnQubnM0LmNhLmx1Z2dzLmNv


## oszx

Secure DNS Project by PumpleX - Hosted in the UK by OVH
No Logging / Ad-Blocking
Information at https://dns.oszx.co

sdns://AQIAAAAAAAAAETUxLjM4LjgzLjE0MTo1MzUzIMwm9_oYw26P4JIVoDhJ_5kFDdNxX1ke4fEzL1V5bwEjFzIuZG5zY3J5cHQtY2VydC5vc3p4LmNv


## pf-dnscrypt

by post-factum | Zürich, Switzerland | Non-logging | Non-filtering | DNSSEC | https://dns.post-factum.tk

sdns://AQcAAAAAAAAAFDE0MC4yMzguMjE1LjE5Mjo4NDQzIH2l4fL6H6BQcKWfdf9ZnrvWxZL_vxKUtQMcWDdZwB6bHjIuZG5zY3J5cHQtY2VydC5wb3N0LWZhY3R1bS50aw


## pf-doh

by post-factum | Zürich, Switzerland | Non-logging | Non-filtering | DNSSEC | Cached globally via Cloudflare | https://dns.post-factum.tk

sdns://AgcAAAAAAAAAACCcu6D8li96KzHGKxsXWi3kxQqDlXJ-MLYmqAAJp4Dj2BJkb2gucG9zdC1mYWN0dW0udGsKL2Rucy1xdWVyeQ


## plan9-dns

Resolver in New Jersey, USA. DNSCrypt protocol. Non-logging, non-filtering, DNSSEC, anonymized.
Running the official Docker image on Vultr by @jlongua1

sdns://AQcAAAAAAAAADjE3My4xOTkuMTI2LjM1IJLwH3z8-G6TDyEk6yXMmGxuss7nGfFnkHrtm4VOSc9SGTIuZG5zY3J5cHQtY2VydC5wbGFuOS1kbnM


## powerdns-doh

By PowerDNS/Open-Xchange https://powerdns.org

sdns://AgcAAAAAAAAAAKA-GhoPbFPz6XpJLVcIS1uYBwWe4FerFQWHb9g_2j24OCAyhv9lpl-vMghe6hOIw3OLp-N4c8kGzOPEootMwqWJiBBkb2gucG93ZXJkbnMub3JnAS8


## publicarray-au

DNSSEC • OpenNIC • Non-logging • Uncensored - hosted on vultr.com
Maintained by publicarray - https://dns.seby.io

sdns://AQcAAAAAAAAADDQ1Ljc2LjExMy4zMSAIVGh4i6eKXqlF6o9Fg92cgD2WcDvKQJ7v_Wq4XrQsVhsyLmRuc2NyeXB0LWNlcnQuZG5zLnNlYnkuaW8


## publicarray-au-doh

DNSSEC • OpenNIC • Non-logging • Uncensored - hosted on vultr.com
Maintained by publicarray - https://dns.seby.io

sdns://AgcAAAAAAAAADDQ1Ljc2LjExMy4zMaA-GhoPbFPz6XpJLVcIS1uYBwWe4FerFQWHb9g_2j24OCAyhv9lpl-vMghe6hOIw3OLp-N4c8kGzOPEootMwqWJiBBkb2guc2VieS5pbzo4NDQzCi9kbnMtcXVlcnk


## publicarray-au2

DNSSEC • OpenNIC • Non-logging • Uncensored - hosted on ovh.com.au
Maintained by publicarray - https://dns.seby.io

sdns://AQcAAAAAAAAAEjEzOS45OS4yMjIuNzI6ODQ0MyCH8q9RBc6SiI6_tBYIuyfkh13C3UaDhV-k_d3XspWMNBsyLmRuc2NyeXB0LWNlcnQuZG5zLnNlYnkuaW8


## publicarray-au2-doh

DNSSEC • OpenNIC • Non-logging • Uncensored - hosted on ovh.com.au
Maintained by publicarray - https://dns.seby.io

sdns://AgcAAAAAAAAADTEzOS45OS4yMjIuNzKgPhoaD2xT8-l6SS1XCEtbmAcFnuBXqxUFh2_YP9o9uDggMob_ZaZfrzIIXuoTiMNzi6fjeHPJBszjxKKLTMKliYgRZG9oLTIuc2VieS5pbzo0NDMKL2Rucy1xdWVyeQ


## pwoss.org-dnscrypt

No filter | No logs | DNSSEC | Nuremberg, Germany (netcup) | Maintained by https://pwoss.org/ (Dan)

sdns://AQcAAAAAAAAAEzQ1LjE0Mi4xNzYuMTcwOjQ0MzQgZF8FTkuFVFw6YP5x8ydebnTb8ONP7Ti6P0K1REyhUHgZMi5kbnNjcnlwdC1jZXJ0LnB3b3NzLm9yZw


## quad9-dnscrypt-ip4-filter-pri

Quad9 (anycast) dnssec/no-log/filter 9.9.9.9 / 149.112.112.9

sdns://AQMAAAAAAAAADDkuOS45Ljk6ODQ0MyBnyEe4yHWM0SAkVUO-dWdG3zTfHYTAC4xHA2jfgh2GPhkyLmRuc2NyeXB0LWNlcnQucXVhZDkubmV0
sdns://AQMAAAAAAAAAEjE0OS4xMTIuMTEyLjk6ODQ0MyBnyEe4yHWM0SAkVUO-dWdG3zTfHYTAC4xHA2jfgh2GPhkyLmRuc2NyeXB0LWNlcnQucXVhZDkubmV0


## quad9-dnscrypt-ip4-nofilter-pri

Quad9 (anycast) no-dnssec/no-log/no-filter 9.9.9.10 / 149.112.112.10

sdns://AQYAAAAAAAAADTkuOS45LjEwOjg0NDMgZ8hHuMh1jNEgJFVDvnVnRt803x2EwAuMRwNo34Idhj4ZMi5kbnNjcnlwdC1jZXJ0LnF1YWQ5Lm5ldA
sdns://AQYAAAAAAAAAEzE0OS4xMTIuMTEyLjEwOjg0NDMgZ8hHuMh1jNEgJFVDvnVnRt803x2EwAuMRwNo34Idhj4ZMi5kbnNjcnlwdC1jZXJ0LnF1YWQ5Lm5ldA


## quad9-dnscrypt-ip6-filter-pri

Quad9 (anycast) dnssec/no-log/filter 2620:fe::fe:9 / 2620:fe::9

Warning: this server is incompatible with DNS anonymization.

sdns://AQMAAAAAAAAAFFsyNjIwOmZlOjpmZTo5XTo4NDQzIGfIR7jIdYzRICRVQ751Z0bfNN8dhMALjEcDaN-CHYY-GTIuZG5zY3J5cHQtY2VydC5xdWFkOS5uZXQ


## quad9-dnscrypt-ip6-nofilter-pri

Quad9 (anycast) no-dnssec/no-log/no-filter 2620:fe::10 / 2620:fe::fe:10

Warning: this server is incompatible with DNS anonymization.

sdns://AQYAAAAAAAAAElsyNjIwOmZlOjoxMF06ODQ0MyBnyEe4yHWM0SAkVUO-dWdG3zTfHYTAC4xHA2jfgh2GPhkyLmRuc2NyeXB0LWNlcnQucXVhZDkubmV0


## quad9-doh-ip4-filter-pri

Quad9 (anycast) dnssec/no-log/filter 9.9.9.9 / 149.112.112.9

sdns://AgMAAAAAAAAABzkuOS45LjkAEmRuczkucXVhZDkubmV0OjQ0MwovZG5zLXF1ZXJ5
sdns://AgMAAAAAAAAADTE0OS4xMTIuMTEyLjkAEmRuczkucXVhZDkubmV0OjQ0MwovZG5zLXF1ZXJ5


## quad9-doh-ip4-nofilter-pri

Quad9 (anycast) no-dnssec/no-log/no-filter 9.9.9.10 / 149.112.112.10

sdns://AgYAAAAAAAAACDkuOS45LjEwABJkbnM5LnF1YWQ5Lm5ldDo0NDMKL2Rucy1xdWVyeQ
sdns://AgYAAAAAAAAADjE0OS4xMTIuMTEyLjEwABJkbnM5LnF1YWQ5Lm5ldDo0NDMKL2Rucy1xdWVyeQ


## quad9-doh-ip6-filter-pri

Quad9 (anycast) dnssec/no-log/filter 2620:fe::9 / 2620:fe::fe:9

sdns://AgMAAAAAAAAADFsyNjIwOmZlOjo5XQASZG5zOS5xdWFkOS5uZXQ6NDQzCi9kbnMtcXVlcnk
sdns://AgMAAAAAAAAAD1syNjIwOmZlOjpmZTo5XQASZG5zOS5xdWFkOS5uZXQ6NDQzCi9kbnMtcXVlcnk


## quad9-doh-ip6-nofilter-pri

Quad9 (anycast) no-dnssec/no-log/no-filter 2620:fe::10 / 2620:fe::fe:10

sdns://AgYAAAAAAAAADVsyNjIwOmZlOjoxMF0AEmRuczkucXVhZDkubmV0OjQ0MwovZG5zLXF1ZXJ5
sdns://AgYAAAAAAAAAEFsyNjIwOmZlOjpmZToxMF0AEmRuczkucXVhZDkubmV0OjQ0MwovZG5zLXF1ZXJ5


## rubyfish-ea

Resolver in mainland China, forwarding queries for non-Chinese domains
to upstream servers in East Asia.

https://www.rubyfish.cn/

sdns://AgEAAAAAAAAAAKA-GhoPbFPz6XpJLVcIS1uYBwWe4FerFQWHb9g_2j24OCAyhv9lpl-vMghe6hOIw3OLp-N4c8kGzOPEootMwqWJiBJlYS1kbnMucnVieWZpc2guY24KL2Rucy1xdWVyeQ


## rubyfish-uw

Resolver in mainland China, forwarding queries for non-Chinese domains
to US-West.

https://www.rubyfish.cn/

sdns://AgEAAAAAAAAAAKA-GhoPbFPz6XpJLVcIS1uYBwWe4FerFQWHb9g_2j24OCAyhv9lpl-vMghe6hOIw3OLp-N4c8kGzOPEootMwqWJiBJ1dy1kbnMucnVieWZpc2guY24KL2Rucy1xdWVyeQ


## safesurfer

Family safety focused blocklist for over 2 million adult sites, as well as phishing and malware and more.
Free to use, paid for customizing blocking for more categories+sites and viewing usage at my.safesurfer.io. Logs taken for viewing
usage, data never sold - https://safesurfer.io

Warning: this server is incompatible with DNS anonymization.

sdns://AQMAAAAAAAAADzEwNC4xNTUuMjM3LjIyNSAnIH_VEgToNntINABd-f_R0wu-KpwzY55u2_iu2R1A2CAyLmRuc2NyeXB0LWNlcnQuc2FmZXN1cmZlci5jby5ueg


## sarpel-dns-istanbul

No-filter | No-logs | Uncensored | Hosted in Istanbul(Turkey) on Cloudeos

sdns://AQcAAAAAAAAADzE4NS4xNTMuMjQ5LjIzOSDkXGvqcMO6TaLJUSkr49jGcXvjT3q1ZyiO-Oy6uE-40jAyLmRuc2NyeXB0LWNlcnQuMTg1LTE1My0yNDktMjM5LmluZi5jbG91ZGVvcy5uZXQ


## scaleway-ams

DNSSEC/Non-logged/Uncensored in Amsterdam - DEV1-S instance donated by Scaleway.com
Maintained by Frank Denis - https://fr.dnscrypt.info

sdns://AQcAAAAAAAAADTUxLjE1LjEyMi4yNTAg6Q3ZfapcbHgiHKLF7QFoli0Ty1Vsz3RXs1RUbxUrwZAcMi5kbnNjcnlwdC1jZXJ0LnNjYWxld2F5LWFtcw


## scaleway-ams-ipv6

DNSSEC/Non-logged/Uncensored in Amsterdam - IPv6 only - DEV1-S instance donated by Scaleway.com
Maintained by Frank Denis - https://fr.dnscrypt.info

sdns://AQcAAAAAAAAAFlsyMDAxOmJjODoxODIwOjUwZDo6MV0g6Q3ZfapcbHgiHKLF7QFoli0Ty1Vsz3RXs1RUbxUrwZAcMi5kbnNjcnlwdC1jZXJ0LnNjYWxld2F5LWFtcw


## scaleway-fr

DNSSEC/Non-logged/Uncensored in Paris - DEV1-S instance donated by Scaleway.com
Maintained by Frank Denis - https://fr.dnscrypt.info

sdns://AQcAAAAAAAAADjIxMi40Ny4yMjguMTM2IOgBuE6mBr-wusDOQ0RbsV66ZLAvo8SqMa4QY2oHkDJNHzIuZG5zY3J5cHQtY2VydC5mci5kbnNjcnlwdC5vcmc


## scaleway-fr-ipv6

DNSSEC/Non-logged/Uncensored in Paris - IPv6 only - DEV1-S instance donated by Scaleway.com
Maintained by Frank Denis - https://fr.dnscrypt.info

sdns://AQcAAAAAAAAAF1syMDAxOmJjODo0N2IwOjFhMDE6OjFdIOgBuE6mBr-wusDOQ0RbsV66ZLAvo8SqMa4QY2oHkDJNHzIuZG5zY3J5cHQtY2VydC5mci5kbnNjcnlwdC5vcmc


## serbica

Public DNSCrypt server in the Netherlands by https://litepay.ch

sdns://AQcAAAAAAAAAEzE4NS42Ni4xNDMuMTc4OjUzNTMg-Y2MQmGOXiggAEKulN-ITGEn_Kj3TIP1UK1X2wh3o7wXMi5kbnNjcnlwdC1jZXJ0LnNlcmJpY2E


## sfw.scaleway-fr

Uses deep learning to block adult websites. Free, DNSSEC, no logs.
Hosted in Paris, running on a 1-XS server donated by Scaleway.com
Maintained by Frank Denis - https://fr.dnscrypt.info/sfw.html

sdns://AQMAAAAAAAAADzE2My4xNzIuMTgwLjEyNSDfYnO_x1IZKotaObwMhaw_-WRF1zZE9mJygl01WPGh_x8yLmRuc2NyeXB0LWNlcnQuc2Z3LnNjYWxld2F5LWZy


## sth-dnscrypt-se

Resolver in Stockholm, Sweden. DNSCrypt server. Non-logging, non-filtering, DNSSEC.

sdns://AQcAAAAAAAAAEjQ1LjE1My4xODcuOTY6NDM0MyC8E4j1dj497HXxyQ_JFb-2iurf6xxF9phRgGOcYOfxYh8yLmRuc2NyeXB0LWNlcnQuc3RoLWRuc2NyeXB0LXNl


## sth-doh-se

Resolver in Stockholm, Sweden. DoH server. Non-logging, non-filtering, DNSSEC.

sdns://AgcAAAAAAAAADTQ1LjE1My4xODcuOTagPhoaD2xT8-l6SS1XCEtbmAcFnuBXqxUFh2_YP9o9uDggMob_ZaZfrzIIXuoTiMNzi6fjeHPJBszjxKKLTMKliYgSZG5zc2UuYWxla2JlcmcubmV0Ci9kbnMtcXVlcnk


## t53

Deutsche Telekom experimental DoH server.

sdns://AgUAAAAAAAAAAKA-GhoPbFPz6XpJLVcIS1uYBwWe4FerFQWHb9g_2j24OCAyhv9lpl-vMghe6hOIw3OLp-N4c8kGzOPEootMwqWJiApkbnMudDUzLmRlCi9kbnMtcXVlcnk


## v.dnscrypt.uk-ipv4

DNSCrypt v2, no logs, uncensored, DNSSEC. Hosted in London UK on Vultr
https://www.dnscrypt.uk

sdns://AQcAAAAAAAAADzEwNC4yMzguMTg2LjE5MiDtST2M6teQZk8GPEe3lZojaS18kDY8nkPMtZF75bQe5R0yLmRuc2NyeXB0LWNlcnQudi5kbnNjcnlwdC51aw


## v.dnscrypt.uk-ipv6

DNSCrypt v2, no logs, uncensored, DNSSEC. Hosted in London UK on Vultr
https://www.dnscrypt.uk

sdns://AQcAAAAAAAAAKFsyMDAxOjE5ZjA6NzQwMjoxNTc0OjU0MDA6MmZmOmZlNjY6MmNmZl0g7Uk9jOrXkGZPBjxHt5WaI2ktfJA2PJ5DzLWRe-W0HuUdMi5kbnNjcnlwdC1jZXJ0LnYuZG5zY3J5cHQudWs


## yandex

Yandex public DNS server (anycast)

sdns://AQQAAAAAAAAAEDc3Ljg4LjguNzg6MTUzNTMg04TAccn3RmKvKszVe13MlxTUB7atNgHhrtwG1W1JYyciMi5kbnNjcnlwdC1jZXJ0LmJyb3dzZXIueWFuZGV4Lm5ldA


## yepdns-doh-sg

Filtering | DNSSEC | No-logs | Singapore | https://get.yepdns.com/

sdns://AgMAAAAAAAAAAAANc2cueWVwZG5zLmNvbQovZG5zLXF1ZXJ5


## zackptg5-us-il-ipv4

DNSSEC/unfiltered/non-logged. Hosted on Vultr in Chicago, IL. Running the official Docker image by @zackptg5

sdns://AQcAAAAAAAAADTEzNy4yMjAuNTkuNjIgTNMjCh27ODMeT_zUDR4yV_L9iWnqXTyyMuyLeAQW0eceMi5kbnNjcnlwdC1jZXJ0LnphY2twdGc1LXVzLWls


## zackptg5-us-il-ipv6

DNSSEC/unfiltered/non-logged. Hosted on Vultr in Chicago, IL. Running the official Docker image by @zackptg5

sdns://AQcAAAAAAAAAKFsyMDAxOjE5ZjA6NWMwMToxNDdlOjU0MDA6M2ZmOmZlMmE6M2IxZl0gTNMjCh27ODMeT_zUDR4yV_L9iWnqXTyyMuyLeAQW0eceMi5kbnNjcnlwdC1jZXJ0LnphY2twdGc1LXVzLWls

