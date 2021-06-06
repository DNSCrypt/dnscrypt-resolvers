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

sdns://AgcAAAAAAAAADTIxNy4xNjkuMjAuMjOgMob_ZaZfrzIIXuoTiMNzi6fjeHPJBszjxKKLTMKliYigRE69Z7uD-IB7OSHpOKyReLiCvVCq2xEjHwRM9fCN984gzBBg05yDKbYrb7x9DW35MJhpuYHn5jktXNj6QI9NgOYNZG5zLmFhLm5ldC51awovZG5zLXF1ZXJ5


## acsacsar-ams-ipv4

Public non-censoring, non-logging, DNSSEC-capable, DNSCrypt-enabled DNS resolver hosted on Scaleway by @acsacsar (twitter)

sdns://AQcAAAAAAAAADTUxLjE1OC4xNjYuOTcgAyfzz5J-mV9G-yOB4Hwcdk7yX12EQs5Iva7kV3oGtlEgMi5kbnNjcnlwdC1jZXJ0LmFjc2Fjc2FyLWFtcy5jb20


## acsacsar-ams-ipv6

Public non-censoring, non-logging, DNSSEC-capable, DNSCrypt-enabled DNS resolver hosted on Scaleway by @acsacsar (twitter)

sdns://AQcAAAAAAAAAFlsyMDAxOmJjODoxODI0OjczODo6MV0gAyfzz5J-mV9G-yOB4Hwcdk7yX12EQs5Iva7kV3oGtlEgMi5kbnNjcnlwdC1jZXJ0LmFjc2Fjc2FyLWFtcy5jb20


## adfree.usableprivacy.net

Public updns DoH service with advertising, tracker and malware filters.

Hosted in Europe by @usableprivacy, details see: https://docs.usableprivacy.com

sdns://AgMAAAAAAAAADzE0OS4xNTQuMTUzLjE1M6Ayhv9lpl-vMghe6hOIw3OLp-N4c8kGzOPEootMwqWJiKBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33ziDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5hhhZGZyZWUudXNhYmxlcHJpdmFjeS5uZXQKL2Rucy1xdWVyeQ


## adguard-dns

Remove ads and protect your computer from malware

Warning: This server is incompatible with anonymization.

sdns://AQMAAAAAAAAAETk0LjE0MC4xNC4xNDo1NDQzINErR_JS3PLCu_iZEIbq95zkSV2LFsigxDIuUso_OQhzIjIuZG5zY3J5cHQuZGVmYXVsdC5uczEuYWRndWFyZC5jb20


## adguard-dns-doh

Remove ads and protect your computer from malware (over DoH)

sdns://AgMAAAAAAAAADzE3Ni4xMDMuMTMwLjEzMCCsFdIhxY-VWoedpSrEKWAhaBEVj-8L-p_FJl6wMpPufg9kbnMuYWRndWFyZC5jb20KL2Rucy1xdWVyeQ


## adguard-dns-family

Adguard DNS with safesearch and adult content blocking

Warning: This server is incompatible with anonymization.

sdns://AQMAAAAAAAAAETk0LjE0MC4xNC4xNTo1NDQzILgxXdexS27jIKRw3C7Wsao5jMnlhvhdRUXWuMm1AFq6ITIuZG5zY3J5cHQuZmFtaWx5Lm5zMS5hZGd1YXJkLmNvbQ


## adguard-dns-family-doh

Adguard DNS with safesearch and adult content blocking (over DoH)

sdns://AgMAAAAAAAAADzE3Ni4xMDMuMTMwLjEzMiCsFdIhxY-VWoedpSrEKWAhaBEVj-8L-p_FJl6wMpPufhZkbnMtZmFtaWx5LmFkZ3VhcmQuY29tCi9kbnMtcXVlcnk


## adguard-dns-family-ipv6

Adguard DNS with safesearch and adult content blocking

Warning: This server is incompatible with anonymization.

sdns://AQMAAAAAAAAAGVsyYTEwOjUwYzA6OmJhZDE6ZmZdOjU0NDMguDFd17FLbuMgpHDcLtaxqjmMyeWG-F1FRda4ybUAWrohMi5kbnNjcnlwdC5mYW1pbHkubnMxLmFkZ3VhcmQuY29t


## adguard-dns-ipv6

Remove ads and protect your computer from malware

Warning: This server is incompatible with anonymization.

sdns://AQMAAAAAAAAAGFsyYTEwOjUwYzA6OmFkMTpmZl06NTQ0MyDRK0fyUtzywrv4mRCG6vec5EldixbIoMQyLlLKPzkIcyIyLmRuc2NyeXB0LmRlZmF1bHQubnMxLmFkZ3VhcmQuY29t


## adguard-dns-unfiltered

AdGuard public DNS servers without filters

Warning: This server is incompatible with anonymization.

sdns://AQMAAAAAAAAAEjk0LjE0MC4xNC4xNDA6NTQ0MyC16ETWuDo-PhJo62gfvqcN48X6aNvWiBQdvy7AZrLa-iUyLmRuc2NyeXB0LnVuZmlsdGVyZWQubnMxLmFkZ3VhcmQuY29t


## adguard-dns-unfiltered-ipv6

AdGuard public DNS servers without filters

Warning: This server is incompatible with anonymization.

sdns://AQMAAAAAAAAAFlsyYTEwOjUwYzA6OjE6ZmZdOjU0NDMgtehE1rg6Pj4SaOtoH76nDePF-mjb1ogUHb8uwGay2volMi5kbnNjcnlwdC51bmZpbHRlcmVkLm5zMS5hZGd1YXJkLmNvbQ


## ahadns-doh-in

A zero logging DNS with support for DNS-over-HTTPS (DoH) & DNS-over-TLS (DoT). Blocks ads, malware, trackers, viruses, ransomware, telemetry and more. No persistent logs. DNSSEC. Hosted in Mumbai, India. By https://ahadns.com/
Server statistics can be seen at: https://statistics.ahadns.com/?server=in

sdns://AgMAAAAAAAAADTQ1Ljc5LjEyMC4yMzOgMob_ZaZfrzIIXuoTiMNzi6fjeHPJBszjxKKLTMKliYigRE69Z7uD-IB7OSHpOKyReLiCvVCq2xEjHwRM9fCN984gzBBg05yDKbYrb7x9DW35MJhpuYHn5jktXNj6QI9NgOYRZG9oLmluLmFoYWRucy5uZXQKL2Rucy1xdWVyeQ


## ahadns-doh-in-ipv6

A zero logging DNS with support for DNS-over-HTTPS (DoH) & DNS-over-TLS (DoT). Blocks ads, malware, trackers, viruses, ransomware, telemetry and more. No persistent logs. DNSSEC. Hosted in Mumbai, India. By https://ahadns.com/
Server statistics can be seen at: https://statistics.ahadns.com/?server=in

sdns://AgMAAAAAAAAAF1syNDAwOjg5MDQ6ZTAwMTo0Mzo6NDNdoDKG_2WmX68yCF7qE4jDc4un43hzyQbM48Sii0zCpYmIoEROvWe7g_iAezkh6TiskXi4gr1QqtsRIx8ETPXwjffOIMwQYNOcgym2K2-8fQ1t-TCYabmB5-Y5LVzY-kCPTYDmEWRvaC5pbi5haGFkbnMubmV0Ci9kbnMtcXVlcnk


## ahadns-doh-la

A zero logging DNS with support for DNS-over-HTTPS (DoH) & DNS-over-TLS (DoT). Blocks ads, malware, trackers, viruses, ransomware, telemetry and more. No persistent logs. DNSSEC. Hosted in Los Angeles, USA. By https://ahadns.com/
Server statistics can be seen at: https://statistics.ahadns.com/?server=la

sdns://AgMAAAAAAAAADTQ1LjY3LjIxOS4yMDigMob_ZaZfrzIIXuoTiMNzi6fjeHPJBszjxKKLTMKliYigRE69Z7uD-IB7OSHpOKyReLiCvVCq2xEjHwRM9fCN984gzBBg05yDKbYrb7x9DW35MJhpuYHn5jktXNj6QI9NgOYRZG9oLmxhLmFoYWRucy5uZXQKL2Rucy1xdWVyeQ


## ahadns-doh-la-ipv6

A zero logging DNS with support for DNS-over-HTTPS (DoH) & DNS-over-TLS (DoT). Blocks ads, malware, trackers, viruses, ransomware, telemetry and more. No persistent logs. DNSSEC. Hosted in Los Angeles, USA. By https://ahadns.com/
Server statistics can be seen at: https://statistics.ahadns.com/?server=la

sdns://AgMAAAAAAAAAFlsyYTA0OmJkYzc6MTAwOjcwOjo3MF2gMob_ZaZfrzIIXuoTiMNzi6fjeHPJBszjxKKLTMKliYigRE69Z7uD-IB7OSHpOKyReLiCvVCq2xEjHwRM9fCN984gzBBg05yDKbYrb7x9DW35MJhpuYHn5jktXNj6QI9NgOYRZG9oLmxhLmFoYWRucy5uZXQKL2Rucy1xdWVyeQ


## ahadns-doh-nl

A zero logging DNS with support for DNS-over-HTTPS (DoH) & DNS-over-TLS (DoT). Blocks ads, malware, trackers, viruses, ransomware, telemetry and more. No persistent logs. DNSSEC. Hosted in Amsterdam, Netherlands. By https://ahadns.com/
Server statistics can be seen at: https://statistics.ahadns.com/?server=nl

sdns://AgMAAAAAAAAACTUuMi43NS43NaAyhv9lpl-vMghe6hOIw3OLp-N4c8kGzOPEootMwqWJiKBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33ziDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5hFkb2gubmwuYWhhZG5zLm5ldAovZG5zLXF1ZXJ5


## ahadns-doh-nl-ipv6

A zero logging DNS with support for DNS-over-HTTPS (DoH) & DNS-over-TLS (DoT). Blocks ads, malware, trackers, viruses, ransomware, telemetry and more. No persistent logs. DNSSEC. Hosted in Amsterdam, Netherlands. By https://ahadns.com/
Server statistics can be seen at: https://statistics.ahadns.com/?server=nl

sdns://AgMAAAAAAAAAFlsyYTA0OjUyYzA6MTAxOjc1Ojo3NV2gMob_ZaZfrzIIXuoTiMNzi6fjeHPJBszjxKKLTMKliYigRE69Z7uD-IB7OSHpOKyReLiCvVCq2xEjHwRM9fCN984gzBBg05yDKbYrb7x9DW35MJhpuYHn5jktXNj6QI9NgOYRZG9oLm5sLmFoYWRucy5uZXQKL2Rucy1xdWVyeQ


## alidns-doh

A public DNS resolver that supports DoH/DoT in mainland China, provided by Alibaba-Cloud.

Warning: GFW filtering rules are applied by that resolver.

Homepage: https://alidns.com/

sdns://AgAAAAAAAAAACTIyMy41LjUuNSCoF6cUD2dwqtorNi96I2e3nkHPSJH1ka3xbdOglmOVkQ5kbnMuYWxpZG5zLmNvbQovZG5zLXF1ZXJ5


## altername

Protocol: DNSCrypt IPv4 | Features: Non-logging, Non-filtering, DNSSEC, OpenNIC, EmerDNS | Current status: Beta | Location: Moscow, Russia

sdns://AQcAAAAAAAAAEzk1LjE4MS4xNTUuMTQwOjg0NDMgvZmXVeT0DhM8O1F9vIC-0KIjH11eoON1lQxZKczVYxIZMi5kbnNjcnlwdC1jZXJ0LmFsdGVybmFtZQ


## ams-ads-doh-nl

Resolver in Amsterdam. DoH protocol. Non-logging. Blocks ads, malware and trackers. DNSSEC enabled.

sdns://AgMAAAAAAAAAEjUxLjE1LjEyNC4yMDg6NDQ0M6Ayhv9lpl-vMghe6hOIw3OLp-N4c8kGzOPEootMwqWJiKBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33ziDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5hdkbnNubC5hbGVrYmVyZy5uZXQ6NDQ0MwovZG5zLXF1ZXJ5


## ams-ads-doh-nl-ipv6

Resolver in Amsterdam. DoH protocol. Non-logging. Blocks ads, malware and trackers. DNSSEC enabled.

sdns://AgMAAAAAAAAAHFsyMDAxOmJjODoxODMwOjIwMTg6OjFdOjQ0NDOgMob_ZaZfrzIIXuoTiMNzi6fjeHPJBszjxKKLTMKliYigRE69Z7uD-IB7OSHpOKyReLiCvVCq2xEjHwRM9fCN984gzBBg05yDKbYrb7x9DW35MJhpuYHn5jktXNj6QI9NgOYXZG5zbmwuYWxla2JlcmcubmV0OjQ0NDMKL2Rucy1xdWVyeQ


## ams-dnscrypt-nl

Resolver in Amsterdam. Dnscrypt protocol. Non-logging, non-filtering, DNSSEC.

sdns://AQcAAAAAAAAAEjUxLjE1LjEyNC4yMDg6NDM0MyC8E4j1dj497HXxyQ_JFb-2iurf6xxF9phRgGOcYOfxYh8yLmRuc2NyeXB0LWNlcnQuYW1zLWRuc2NyeXB0LW5s


## ams-dnscrypt-nl-ipv6

Resolver in Amsterdam. Dnscrypt protocol. Non-logging, non-filtering, DNSSEC.

sdns://AQcAAAAAAAAAHFsyMDAxOmJjODoxODMwOjIwMTg6OjFdOjQzNDMgvBOI9XY-Pex18ckPyRW_torq3-scRfaYUYBjnGDn8WIfMi5kbnNjcnlwdC1jZXJ0LmFtcy1kbnNjcnlwdC1ubA


## ams-doh-nl

Resolver in Amsterdam. DoH protocol. Non-logging, non-filtering, DNSSEC.

sdns://AgcAAAAAAAAADTUxLjE1LjEyNC4yMDigMob_ZaZfrzIIXuoTiMNzi6fjeHPJBszjxKKLTMKliYigRE69Z7uD-IB7OSHpOKyReLiCvVCq2xEjHwRM9fCN984gzBBg05yDKbYrb7x9DW35MJhpuYHn5jktXNj6QI9NgOYSZG5zbmwuYWxla2JlcmcubmV0Ci9kbnMtcXVlcnk


## ams-doh-nl-ipv6

Resolver in Amsterdam. DoH protocol. Non-logging, non-filtering, DNSSEC.

sdns://AgcAAAAAAAAAF1syMDAxOmJjODoxODMwOjIwMTg6OjFdoDKG_2WmX68yCF7qE4jDc4un43hzyQbM48Sii0zCpYmIoEROvWe7g_iAezkh6TiskXi4gr1QqtsRIx8ETPXwjffOIMwQYNOcgym2K2-8fQ1t-TCYabmB5-Y5LVzY-kCPTYDmEmRuc25sLmFsZWtiZXJnLm5ldAovZG5zLXF1ZXJ5


## arapurayil-doh-in

DoH server(ipv4/ipv6). Also supports DNSCrypt,DoT,DoQ protocols. Located in Mumbai, India.  Visit https://www.dns.arapurayil.com for details.
Non-logging | Filtering | DNSSEC | anti-CNAME cloaking | QNAME Minimization | No EDNS Client-Subnet

sdns://AgMAAAAAAAAAAKAyhv9lpl-vMghe6hOIw3OLp-N4c8kGzOPEootMwqWJiKBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33ziDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5hJkbnMuYXJhcHVyYXlpbC5jb20KL2Rucy1xdWVyeQ


## att

AT&T test DoH server.

sdns://AgQAAAAAAAAAAKBLTrSwdCmLgotcADCVoQtFI_uVHAyINIsJxT5bq6QIoyD2Hldod9qWUClMzLX5bHX8txvaG7xGRjZ8Tr7aidcxjxBkb2h0cmlhbC5hdHQubmV0Ci9kbnMtcXVlcnk


## bcn-dnscrypt

Resolver in Barcelona, Spain. DNSCrypt protocol. Non-logging, non-filtering, DNSSEC.

sdns://AQcAAAAAAAAAEzE4NS4yNTMuMTU0LjY2OjQzNDMg4Bb-Ip19x31rL17BHy3FBLBP4s_ZtyjEYe_jDFPIa00cMi5kbnNjcnlwdC1jZXJ0LmJjbi1kbnNjcnlwdA


## bcn-doh

Resolver in Barcelona, Spain. DoH protocol. Non-logging, non-filtering, DNSSEC.

sdns://AgcAAAAAAAAADjE4NS4yNTMuMTU0LjY2oDKG_2WmX68yCF7qE4jDc4un43hzyQbM48Sii0zCpYmIoEROvWe7g_iAezkh6TiskXi4gr1QqtsRIx8ETPXwjffOIMwQYNOcgym2K2-8fQ1t-TCYabmB5-Y5LVzY-kCPTYDmEmRuc2VzLmFsZWtiZXJnLm5ldAovZG5zLXF1ZXJ5


## bortzmeyer

Non-logging DoH server in France operated by Stéphane Bortzmeyer.

https://www.bortzmeyer.org/doh-bortzmeyer-fr-policy.html

sdns://AgcAAAAAAAAADDE5My43MC44NS4xMaAyhv9lpl-vMghe6hOIw3OLp-N4c8kGzOPEootMwqWJiKBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33ziDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5hFkb2guYm9ydHptZXllci5mcgEv


## bortzmeyer-ipv6

Non-logging DoH server in France operated by Stéphane Bortzmeyer (IPv6 only).

https://www.bortzmeyer.org/doh-bortzmeyer-fr-policy.html

sdns://AgcAAAAAAAAAGVsyMDAxOjQxZDA6MzAyOjIyMDA6OjE4MF2gMob_ZaZfrzIIXuoTiMNzi6fjeHPJBszjxKKLTMKliYigRE69Z7uD-IB7OSHpOKyReLiCvVCq2xEjHwRM9fCN984gzBBg05yDKbYrb7x9DW35MJhpuYHn5jktXNj6QI9NgOYRZG9oLmJvcnR6bWV5ZXIuZnIBLw


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


## comodo-02

Comodo Dome Shield (anycast) - https://cdome.comodo.com/shield/

sdns://AQUAAAAAAAAACjguMjAuMjQ3LjIg0sJUqpYcHsoXmZb1X7yAHwg2xyN5q1J-zaiGG-Dgs7AoMi5kbnNjcnlwdC1jZXJ0LnNoaWVsZC0yLmRuc2J5Y29tb2RvLmNvbQ


## comss.one

DNS server in Lithuania filtering phishing and ads.
https://www.comss.ru/page.php?id=7315

sdns://AQMAAAAAAAAADTk0LjE3Ni4yMzMuOTMgFWnIA4ZtJKvnIs9g6yZT4WIyPb-rQEB27paIxM_OkxMdMi5kbnNjcnlwdC1jZXJ0LmRucy5jb21zcy5vbmU


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


## cs-rome

Rome, Italy DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADjM3LjEyMC4yMDcuMTMxIDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


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

sdns://AgcAAAAAAAAADDE4NS40My4xMzUuMaAyhv9lpl-vMghe6hOIw3OLp-N4c8kGzOPEootMwqWJiKBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33ziDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5gtvZHZyLm5pYy5jegQvZG9o


## cz.nic-ipv6

CZ.NIC's open DNSSEC validating resolvers in Prague, Czech Republic (IPv6 only).
CZ.NIC resolvers neither collect any personal data nor gather
information on pages where your computer sends personal data.
https://www.nic.cz/odvr/

sdns://AgcAAAAAAAAAE1syMDAxOjE0OGY6ZmZmZTo6MV2gMob_ZaZfrzIIXuoTiMNzi6fjeHPJBszjxKKLTMKliYigRE69Z7uD-IB7OSHpOKyReLiCvVCq2xEjHwRM9fCN984gzBBg05yDKbYrb7x9DW35MJhpuYHn5jktXNj6QI9NgOYLb2R2ci5uaWMuY3oEL2RvaA


## d0wn-is-ns2

Server provided by Martin 'd0wn' Albus

sdns://AQcAAAAAAAAADTkzLjk1LjIyNi4xNjUghGA0qcYwyjwErEqQFiXxeoeyrLlBgKxIHiwQ6M7eGm8cMi5kbnNjcnlwdC1jZXJ0LmlzMi5kMHduLmJpeg


## d0wn-tz-ns1

Server provided by Martin 'd0wn' Albus

sdns://AQcAAAAAAAAACzQxLjc5LjY5LjEzINYGFfvRRTuhTnaKPlxcs6wXRhMxRj2gr4z33wTaTXVtGzIuZG5zY3J5cHQtY2VydC50ei5kMHduLmJpeg


## d0wn-tz-ns1-ipv6

Server provided by Martin 'd0wn' Albus

sdns://AQcAAAAAAAAAGFsyYzBmOmZkYTg6NTo6MmVkMTpkMmVjXSDWBhX70UU7oU52ij5cXLOsF0YTMUY9oK-M998E2k11bRsyLmRuc2NyeXB0LWNlcnQudHouZDB3bi5iaXo


## dama.no osl-s04

DNSCrypt server located in Oslo/Norway. Link-speed 100 Mbit/s.
Non-censoring, non-logging, DNSSEC-capable.

sdns://AQcAAAAAAAAAEzIxNy4xNzAuMjA1LjEwNDo0NDMg7XzHuxk9EXOSK6lTQexoqlCRRdvD7TuNzBTlb9udIRMfMi5kbnNjcnlwdC1jZXJ0Lm9zbC1zMDQuZGFtYS5ubw


## dama.no sa-a80

DNSCrypt server located in Sandefjord/Norway. Link-speed 2.5 Gbit/s and low latency from Northern Europe (north of Hamburg).
Non-censoring, non-logging, DNSSEC-capable.

Try `ping -c 100 193.200.238.80` to measure RTT.

sdns://AQcAAAAAAAAADjE5My4yMDAuMjM4LjgwIIZwYLY60l2SdIfAz_N0-AmQM8gjVwLWNly83S3XVcSaHjIuZG5zY3J5cHQtY2VydC5zYS1hODAuZGFtYS5ubw


## dct-ru1

DNSCrypt | IPv4 only | Non-logging | Non-filtering | DNSSEC | Saint Petersburg, Russia.

sdns://AQcAAAAAAAAAEDk0LjI0Mi41OS4zNTo0NDMgTg8e1Qx5q3cMqZ8MkxmG6850fFHXcEwdd_7gAcWcp6wXMi5kbnNjcnlwdC1jZXJ0LmRjdC1ydTE


## dct-ru2

DNSCrypt | IPv4 only | Non-logging | Non-filtering | DNSSEC | Moscow, Russia.

sdns://AQcAAAAAAAAAETE4NS4yMi4xNTQuMTk6NDQzIB9z-95JSBqhYTv8UIQqUs65Ym2DYNXzBf5cN-kH8Pt9FzIuZG5zY3J5cHQtY2VydC5kY3QtcnUy


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

sdns://AgcAAAAAAAAADTE4NS45NS4yMTguNDKgMob_ZaZfrzIIXuoTiMNzi6fjeHPJBszjxKKLTMKliYigRE69Z7uD-IB7OSHpOKyReLiCvVCq2xEjHwRM9fCN984gzBBg05yDKbYrb7x9DW35MJhpuYHn5jktXNj6QI9NgOYcZG5zLmRpZ2l0YWxlLWdlc2VsbHNjaGFmdC5jaAovZG5zLXF1ZXJ5


## dns.digitale-gesellschaft.ch-2

Public DoH resolver operated by the Digital Society (https://www.digitale-gesellschaft.ch).
Hosted in Zurich, Switzerland.

Non-logging, non-filtering, supports DNSSEC.

sdns://AgcAAAAAAAAADTE4NS45NS4yMTguNDOgMob_ZaZfrzIIXuoTiMNzi6fjeHPJBszjxKKLTMKliYigRE69Z7uD-IB7OSHpOKyReLiCvVCq2xEjHwRM9fCN984gzBBg05yDKbYrb7x9DW35MJhpuYHn5jktXNj6QI9NgOYcZG5zLmRpZ2l0YWxlLWdlc2VsbHNjaGFmdC5jaAovZG5zLXF1ZXJ5


## dns.digitale-gesellschaft.ch-ipv6

Public IPv6 DoH resolver operated by the Digital Society (https://www.digitale-gesellschaft.ch).
Hosted in Zurich, Switzerland.

Non-logging, non-filtering, supports DNSSEC.

sdns://AgcAAAAAAAAAD1syYTA1OmZjODQ6OjQyXaAyhv9lpl-vMghe6hOIw3OLp-N4c8kGzOPEootMwqWJiKBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33ziDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5hxkbnMuZGlnaXRhbGUtZ2VzZWxsc2NoYWZ0LmNoCi9kbnMtcXVlcnk


## dns.digitale-gesellschaft.ch-ipv6-2

Public IPv6 DoH resolver operated by the Digital Society (https://www.digitale-gesellschaft.ch).
Hosted in Zurich, Switzerland.

Non-logging, non-filtering, supports DNSSEC.

sdns://AgcAAAAAAAAAD1syYTA1OmZjODQ6OjQzXaAyhv9lpl-vMghe6hOIw3OLp-N4c8kGzOPEootMwqWJiKBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33ziDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5hxkbnMuZGlnaXRhbGUtZ2VzZWxsc2NoYWZ0LmNoCi9kbnMtcXVlcnk


## dns.ryan-palmer

Non-logging, non-filtering, DNSSEC DoH Server. Hosted in the UK.

sdns://AgcAAAAAAAAADjY4LjE4My4yNTMuMjAwoDKG_2WmX68yCF7qE4jDc4un43hzyQbM48Sii0zCpYmIoEROvWe7g_iAezkh6TiskXi4gr1QqtsRIx8ETPXwjffOIMwQYNOcgym2K2-8fQ1t-TCYabmB5-Y5LVzY-kCPTYDmFGRuczEucnlhbi1wYWxtZXIuY29tCi9kbnMtcXVlcnk


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

sdns://AgMAAAAAAAAADzE3Mi4xMDQuMjA2LjE3NKAyhv9lpl-vMghe6hOIw3OLp-N4c8kGzOPEootMwqWJiKBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33ziDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5hRkbnMudGhlcmlmbGVtYW4ubmFtZQovZG5zLXF1ZXJ5


## dnscrypt-ch-blahdns-ipv4

Blocks ad and Tracking, no Logging, DNSSEC, Hosted in Switzerland. By https://blahdns.com/

sdns://AQMAAAAAAAAAETQ1LjkxLjkyLjEyMTo4NDQzIIUTqWvE4INLlQ1jtQLo0-LVTQSWayYCkzTpBx26I_wrGzIuZG5zY3J5cHQtY2VydC5ibGFoZG5zLmNvbQ


## dnscrypt-ch-blahdns-ipv6

Blocks ad and Tracking, no Logging, DNSSEC, Hosted in Switzerland. By https://blahdns.com/

sdns://AQMAAAAAAAAAF1syYTBlOmRjMDo2OjIzOjoyXTo4NDQzIIUTqWvE4INLlQ1jtQLo0-LVTQSWayYCkzTpBx26I_wrGzIuZG5zY3J5cHQtY2VydC5ibGFoZG5zLmNvbQ


## dnscrypt-de-blahdns-ipv4

Blocks ad and Tracking, no Logging, DNSSEC, Hosted in Germany. By https://blahdns.com/

sdns://AQMAAAAAAAAAEjc4LjQ2LjI0NC4xNDM6ODQ0MyCFE6lrxOCDS5UNY7UC6NPi1U0ElmsmApM06QcduiP8KxsyLmRuc2NyeXB0LWNlcnQuYmxhaGRucy5jb20


## dnscrypt-de-blahdns-ipv6

Blocks ad and Tracking, no Logging, DNSSEC, Hosted in Germany. By https://blahdns.com/

sdns://AQMAAAAAAAAAG1syYTAxOjRmODpjMTc6ZWM2Nzo6MV06ODQ0MyCFE6lrxOCDS5UNY7UC6NPi1U0ElmsmApM06QcduiP8KxsyLmRuc2NyeXB0LWNlcnQuYmxhaGRucy5jb20


## dnscrypt-fi-blahdns-ipv4

Blocks ad and Tracking, no Logging, DNSSEC, Hosted in Finland. By https://blahdns.com/

sdns://AQMAAAAAAAAAEzk1LjIxNi4yMTIuMTc3Ojg0NDMghROpa8Tgg0uVDWO1AujT4tVNBJZrJgKTNOkHHboj_CsbMi5kbnNjcnlwdC1jZXJ0LmJsYWhkbnMuY29t


## dnscrypt-fi-blahdns-ipv6

Blocks ad and Tracking, no Logging, DNSSEC, Hosted in Finland. By https://blahdns.com/

sdns://AQMAAAAAAAAAHFsyYTAxOjRmOTpjMDEwOjQzY2U6OjFdOjg0NDMghROpa8Tgg0uVDWO1AujT4tVNBJZrJgKTNOkHHboj_CsbMi5kbnNjcnlwdC1jZXJ0LmJsYWhkbnMuY29t


## dnscrypt-jp-blahdns-ipv4

Blocks ad and Tracking, no Logging, DNSSEC, Hosted in Japan. By https://blahdns.com/

sdns://AQMAAAAAAAAAEzEzOS4xNjIuMTEyLjQ3Ojg0NDMghROpa8Tgg0uVDWO1AujT4tVNBJZrJgKTNOkHHboj_CsbMi5kbnNjcnlwdC1jZXJ0LmJsYWhkbnMuY29t


## dnscrypt-jp-blahdns-ipv6

Blocks ad and Tracking, no Logging, DNSSEC, Hosted in Japan. By https://blahdns.com/

sdns://AQMAAAAAAAAAJVsyNDAwOjg5MDI6OmYwM2M6OTJmZjpmZTI3OjM0NGJdOjg0NDMghROpa8Tgg0uVDWO1AujT4tVNBJZrJgKTNOkHHboj_CsbMi5kbnNjcnlwdC1jZXJ0LmJsYWhkbnMuY29t


## dnscrypt-sg-blahdns-ipv4

Blocks ad and Tracking, no Logging, DNSSEC, Hosted in Singapore. By https://blahdns.com/

sdns://AQMAAAAAAAAAEzE5Mi41My4xNzUuMTQ5Ojg0NDMghROpa8Tgg0uVDWO1AujT4tVNBJZrJgKTNOkHHboj_CsbMi5kbnNjcnlwdC1jZXJ0LmJsYWhkbnMuY29t


## dnscrypt-sg-blahdns-ipv6

Blocks ad and Tracking, no Logging, DNSSEC, Hosted in Singapore. By https://blahdns.com/

sdns://AQMAAAAAAAAAJVsyNDAwOjg5MDE6OmYwM2M6OTJmZjpmZTI3Ojg3MGFdOjg0NDMghROpa8Tgg0uVDWO1AujT4tVNBJZrJgKTNOkHHboj_CsbMi5kbnNjcnlwdC1jZXJ0LmJsYWhkbnMuY29t


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

sdns://AgcAAAAAAAAADzE2Ny4xMTQuMjIwLjEyNaAyhv9lpl-vMghe6hOIw3OLp-N4c8kGzOPEootMwqWJiKBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33ziDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5hRkbnMxLmRuc2NyeXB0LmNhOjQ1MwovZG5zLXF1ZXJ5


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

sdns://AgcAAAAAAAAADTE0OS41Ni4yMjguNDWgMob_ZaZfrzIIXuoTiMNzi6fjeHPJBszjxKKLTMKliYigRE69Z7uD-IB7OSHpOKyReLiCvVCq2xEjHwRM9fCN984gzBBg05yDKbYrb7x9DW35MJhpuYHn5jktXNj6QI9NgOYUZG5zMi5kbnNjcnlwdC5jYTo0NTMKL2Rucy1xdWVyeQ


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

sdns://AgIAAAAAAAAADTk1LjIxNy4yMTMuOTSgMob_ZaZfrzIIXuoTiMNzi6fjeHPJBszjxKKLTMKliYigRE69Z7uD-IB7OSHpOKyReLiCvVCq2xEjHwRM9fCN984gzBBg05yDKbYrb7x9DW35MJhpuYHn5jktXNj6QI9NgOYYZG5zLWRvaC5kbnNmb3JmYW1pbHkuY29tCi9kbnMtcXVlcnk


## dnsforfamily-doh-no-safe-search

(DoH Protocol) Block adult websites, gambling websites, malwares and advertisements.
Unlike other dnsforfamily DNSCrypt servers, this one does not enforces safe search. So Google, YouTube, Bing, DuckDuckGo and Yandex are completely accessible without any restriction.

Social websites like Facebook and Instagram are not blocked. No DNS queries are logged.

As of December 2020 2.7 million websites are blocked and new websites are added to blacklist daily.
Completely free, no ads or any commercial motive. Operating for 3 years now.

Warning: This server is incompatible with anonymization.

Provided by: https://dnsforfamily.com

sdns://AgIAAAAAAAAADTk1LjIxNy4yMTMuOTSgMob_ZaZfrzIIXuoTiMNzi6fjeHPJBszjxKKLTMKliYigRE69Z7uD-IB7OSHpOKyReLiCvVCq2xEjHwRM9fCN984gzBBg05yDKbYrb7x9DW35MJhpuYHn5jktXNj6QI9NgOYnZG5zLWRvaC1uby1zYWZlLXNlYXJjaC5kbnNmb3JmYW1pbHkuY29tCi9kbnMtcXVlcnk


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

sdns://AgMAAAAAAAAADDE3Ni45LjkzLjE5OKAyhv9lpl-vMghe6hOIw3OLp-N4c8kGzOPEootMwqWJiKBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33ziDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5gtkbnNmb3JnZS5kZQovZG5zLXF1ZXJ5


## dnshome-doh

https://www.dnshome.de/ public resolver in Germany

sdns://AgYAAAAAAAAAAKAyhv9lpl-vMghe6hOIw3OLp-N4c8kGzOPEootMwqWJiKBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33ziDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5g5kbnMuZG5zaG9tZS5kZQovZG5zLXF1ZXJ5


## dnspod-doh

A public DNS resolver in mainland China provided by DNSPod (Tencent Cloud).

https://www.dnspod.cn/Products/Public.DNS?lang=en

sdns://AgAAAAAAAAAAACDrdSX4jw2UWPgamVAZv9NMuJzNyVfnsO8xXxD4l2OBGAdkb2gucHViCi9kbnMtcXVlcnk


## doh-ch-blahdns

Blocks ad and Tracking, no Logging, DNSSEC, Hosted in Switzerland. By https://blahdns.com/

sdns://AgMAAAAAAAAADDQ1LjkxLjkyLjEyMQASZG9oLWNoLmJsYWhkbnMuY29tCi9kbnMtcXVlcnk


## doh-ch-blahdns-v6

Blocks ad and Tracking, no Logging, DNSSEC, Hosted in Switzerland. By https://blahdns.com/

sdns://AgMAAAAAAAAAElsyYTBlOmRjMDo2OjIzOjoyXQASZG9oLWNoLmJsYWhkbnMuY29tCi9kbnMtcXVlcnk


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

sdns://AgcAAAAAAAAACzEwNC4yMS42Ljc4AA1kb2guY3J5cHRvLnN4Ci9kbnMtcXVlcnk
sdns://AgcAAAAAAAAADjE3Mi42Ny4xMzQuMTU3AA1kb2guY3J5cHRvLnN4Ci9kbnMtcXVlcnk


## doh-crypto-sx-ipv6

DNS-over-HTTPS server accessible over IPv6. Anycast, no logs, no censorship, DNSSEC.
Backend hosted by Scaleway, globally cached via Cloudflare.
Maintained by Frank Denis.

sdns://AgcAAAAAAAAAGlsyNjA2OjQ3MDA6MzAzNzo6NjgxNTo2NGVdABJkb2gtaXB2Ni5jcnlwdG8uc3gKL2Rucy1xdWVyeQ
sdns://AgcAAAAAAAAAG1syNjA2OjQ3MDA6MzAzNjo6YWM0Mzo4NjlkXQASZG9oLWlwdjYuY3J5cHRvLnN4Ci9kbnMtcXVlcnk


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


## doh-ibksturm

DoH Server, No Logging, No Filters, DNSSEC

Running privately by ibksturm in Thurgau, Switzerland

sdns://AgcAAAAAAAAAAKAyhv9lpl-vMghe6hOIw3OLp-N4c8kGzOPEootMwqWJiKBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33ziDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5hRpYmtzdHVybS5zeW5vbG9neS5tZQovZG5zLXF1ZXJ5


## doh-jp-blahdns

Blocks ad and Tracking, no Logging, DNSSEC, Hosted in Japan. By https://blahdns.com/

sdns://AgMAAAAAAAAADjEzOS4xNjIuMTEyLjQ3ABJkb2gtanAuYmxhaGRucy5jb20KL2Rucy1xdWVyeQ


## doh-jp-blahdns-v6

Blocks ad and Tracking, no Logging, DNSSEC, Hosted in Japan. By https://blahdns.com/

sdns://AgMAAAAAAAAAIFsyNDAwOjg5MDI6OmYwM2M6OTJmZjpmZTI3OjM0NGJdABJkb2gtanAuYmxhaGRucy5jb20KL2Rucy1xdWVyeQ


## doh-sg-blahdns

Blocks ad and Tracking, no Logging, DNSSEC, Hosted in Singapore. By https://blahdns.com/

sdns://AgMAAAAAAAAADjE5Mi41My4xNzUuMTQ5ABJkb2gtc2cuYmxhaGRucy5jb20KL2Rucy1xdWVyeQ


## doh-sg-blahdns-v6

Blocks ad and Tracking, no Logging, DNSSEC, Hosted in Singapore. By https://blahdns.com/

sdns://AgMAAAAAAAAAIFsyNDAwOjg5MDE6OmYwM2M6OTJmZjpmZTI3Ojg3MGFdABJkb2gtc2cuYmxhaGRucy5jb20KL2Rucy1xdWVyeQ


## doh.appliedprivacy.net

Public DoH resolver operated by the Foundation for Applied Privacy (https://appliedprivacy.net).
Hosted in Vienna, Austria.

Non-logging, non-filtering, supports DNSSEC.

sdns://AgcAAAAAAAAAAKAyhv9lpl-vMghe6hOIw3OLp-N4c8kGzOPEootMwqWJiKBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33ziDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5hZkb2guYXBwbGllZHByaXZhY3kubmV0Bi9xdWVyeQ


## doh.ffmuc.net

An open (non-logging, non-filtering, non-censoring) DoH resolver operated by Freifunk Munich with nodes in DE.
https://ffmuc.net/

sdns://AgcAAAAAAAAACjUuMS42Ni4yNTWgMob_ZaZfrzIIXuoTiMNzi6fjeHPJBszjxKKLTMKliYigRE69Z7uD-IB7OSHpOKyReLiCvVCq2xEjHwRM9fCN984gzBBg05yDKbYrb7x9DW35MJhpuYHn5jktXNj6QI9NgOYNZG9oLmZmbXVjLm5ldAovZG5zLXF1ZXJ5


## doh.ffmuc.net-v6

An open (non-logging, non-filtering, non-censoring) DoH resolver operated by Freifunk Munich with nodes in DE.
https://ffmuc.net/

sdns://AgcAAAAAAAAAFVsyMDAxOjY3ODplNjg6ZjAwMDo6XaAyhv9lpl-vMghe6hOIw3OLp-N4c8kGzOPEootMwqWJiKBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33ziDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5g1kb2guZmZtdWMubmV0Ci9kbnMtcXVlcnk


## doh.tiarap.org

Non-Logging DNS-over-HTTPS server, cached via Cloudflare.
Filters out ads, trackers and malware, NO ECS, supports DNSSEC.

sdns://AgMAAAAAAAAADDEwNC4yMS42NS42MAAOZG9oLnRpYXJhcC5vcmcKL2Rucy1xdWVyeQ


## doh.tiarap.org-ipv6

Non-Logging DNS-over-HTTPS server (IPv6), cached via Cloudflare.
Filters out ads, trackers and malware, NO ECS, supports DNSSEC.

sdns://AgMAAAAAAAAAG1syNjA2OjQ3MDA6MzAzNDo6NjgxNTo0MTNjXQAOZG9oLnRpYXJhcC5vcmcKL2Rucy1xdWVyeQ


## east.comss.one

DNS with adblock filters and antiphishing, gaining popularity among russian-speaking users.

sdns://AQMAAAAAAAAAEjkxLjIzMC4yMTEuNjc6NTQ0MyAVacgDhm0kq-ciz2DrJlPhYjI9v6tAQHbulojEz86TEyIyLmRuc2NyeXB0LWNlcnQuZG5zLmVhc3QuY29tc3Mub25l


## emeraldonion-doh

Non-logging, anycast DoH server with multiple nodes in the US.

Emerald Onion is a 501(c)(3) nonprofit organization and transit
internet service provider (ISP) based in Seattle, WA.

https://emeraldonion.org/faq/

This public DNS service is shared by Emerald Onion's Tor exit relays,
meaning that Tor user's queries are blended with non-Tor exit
user's queries, further enhancing DNS privacy.

sdns://AgcAAAAAAAAADTEwMy4yMzIuMjA3LjOgMob_ZaZfrzIIXuoTiMNzi6fjeHPJBszjxKKLTMKliYigRE69Z7uD-IB7OSHpOKyReLiCvVCq2xEjHwRM9fCN984gzBBg05yDKbYrb7x9DW35MJhpuYHn5jktXNj6QI9NgOYUZG5zLmVtZXJhbGRvbmlvbi5vcmcKL2Rucy1xdWVyeQ


## faelix-ch-ipv4

An open (non-logging, non-filtering, no ECS) DNSCrypt resolver operated by https://faelix.net/ with IPv4 nodes anycast within AS41495 in Switzerland.

sdns://AQcAAAAAAAAAEzE4NS4xMzQuMTk2LjU0Ojg0NDMgfsvvPi8BgDKNYODh0ewj5Oh32OoJoZNwGgTWs8C-i-EfMi5kbnNjcnlwdC1jZXJ0LnJkbnMuZmFlbGl4Lm5ldA


## faelix-ch-ipv6

An open (non-logging, non-filtering, no ECS) DNSCrypt resolver operated by https://faelix.net/ with IPv6 nodes anycast within AS41495 in Switzerland.

sdns://AQcAAAAAAAAAFFsyYTAxOjllMDE6OjU0XTo4NDQzIH7L7z4vAYAyjWDg4dHsI-Tod9jqCaGTcBoE1rPAvovhHzIuZG5zY3J5cHQtY2VydC5yZG5zLmZhZWxpeC5uZXQ


## faelix-uk-ipv4

An open (non-logging, non-filtering, no ECS) DNSCrypt resolver operated by https://faelix.net/ with IPv4 nodes anycast within AS41495 in the UK.

sdns://AQcAAAAAAAAAEjQ2LjIyNy4yMDAuNTQ6ODQ0MyB-y-8-LwGAMo1g4OHR7CPk6HfY6gmhk3AaBNazwL6L4R8yLmRuc2NyeXB0LWNlcnQucmRucy5mYWVsaXgubmV0


## faelix-uk-ipv6

An open (non-logging, non-filtering, no ECS) DNSCrypt resolver operated by https://faelix.net/ with IPv6 nodes anycast within AS41495 in the UK.

sdns://AQcAAAAAAAAAFFsyYTAxOjllMDA6OjU0XTo4NDQzIH7L7z4vAYAyjWDg4dHsI-Tod9jqCaGTcBoE1rPAvovhHzIuZG5zY3J5cHQtY2VydC5yZG5zLmZhZWxpeC5uZXQ


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

sdns://AgUAAAAAAAAABzguOC44LjigHvYkz_9ea9O63fP92_3qVlRn43cpncfuZnUWbzAMwbmgdoAkR6AZkxo_AEMExT_cbBssN43Evo9zs5_ZyWnftEUgalBisNF41VbxY7E7Gw8ZQ10CWIKRzHVYnf7m6xHI1cMKZG5zLmdvb2dsZQovZG5zLXF1ZXJ5


## google-ipv6

Google DNS (anycast)

sdns://AgUAAAAAAAAAFlsyMDAxOjQ4NjA6NDg2MDo6ODg4OF2gHvYkz_9ea9O63fP92_3qVlRn43cpncfuZnUWbzAMwbmgdoAkR6AZkxo_AEMExT_cbBssN43Evo9zs5_ZyWnftEUgalBisNF41VbxY7E7Gw8ZQ10CWIKRzHVYnf7m6xHI1cMKZG5zLmdvb2dsZQovZG5zLXF1ZXJ5


## he

Hurricane Electric DoH server (anycast)

Unknown logging policy.

sdns://AgUAAAAAAAAACzc0LjgyLjQyLjQyoDKG_2WmX68yCF7qE4jDc4un43hzyQbM48Sii0zCpYmIoEROvWe7g_iAezkh6TiskXi4gr1QqtsRIx8ETPXwjffOIMwQYNOcgym2K2-8fQ1t-TCYabmB5-Y5LVzY-kCPTYDmDG9yZG5zLmhlLm5ldAovZG5zLXF1ZXJ5


## ibksturm

DNSCRYPT V2 Server, No Logging, No Filters, DNSSEC

Running privately by ibksturm in Thurgau, Switzerland

sdns://AQcAAAAAAAAAEzIxMy4xOTYuMTkxLjk2Ojg0NDMgHK0AUhqiLSuBFR07jpBhKvko_oyqyWnot8z4cce7cKkYMi5kbnNjcnlwdC1jZXJ0Lmlia3N0dXJt


## id-gmail

Non-Logging DNSCrypt server located in Singapore.
Filters out ads, trackers and malware, supports DNSSEC, provided by id-gmail.

sdns://AQMAAAAAAAAADjE3NC4xMzguMjEuMTI4IO-WgGbo2ZTwZdg-3dMa7u31bYZXRj5KykfN1_6Xw9T2HDIuZG5zY3J5cHQtY2VydC5kbnMudGlhci5hcHA


## id-gmail-doh

Non-Logging DNS-over-HTTPS server located in Singapore.
Filters out ads, trackers and malware, supports DNSSEC, provided by id-gmail.

sdns://AgMAAAAAAAAADjE3NC4xMzguMjkuMTc1IEROvWe7g_iAezkh6TiskXi4gr1QqtsRIx8ETPXwjffODGRvaC50aWFyLmFwcAovZG5zLXF1ZXJ5


## id-gmail-doh-ipv6

Non-Logging DNS-over-HTTPS (IPv6) server located in Singapore.
Filters out ads, trackers and malware, supports DNSSEC, provided by id-gmail.

sdns://AgMAAAAAAAAAG1syNDAwOjYxODA6MDpkMDo6NWY3Mzo0MDAxXSBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33zgxkb2gudGlhci5hcHAKL2Rucy1xdWVyeQ


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

sdns://AgYAAAAAAAAACjExNi4wLjIuODagMob_ZaZfrzIIXuoTiMNzi6fjeHPJBszjxKKLTMKliYigRE69Z7uD-IB7OSHpOKyReLiCvVCq2xEjHwRM9fCN984gzBBg05yDKbYrb7x9DW35MJhpuYHn5jktXNj6QI9NgOYPZG9oLnRoaXMud2ViLmlkCi9kbnMtcXVlcnk


## jp.tiar.app

Non-Logging, Non-Filtering DNSCrypt server in Japan.
No ECS, Support DNSSEC

sdns://AQcAAAAAAAAAEjE3Mi4xMDQuOTMuODA6MTQ0MyAyuHY-8b9lNqHeahPAzW9IoXnjiLaZpTeNbVs8TN9UUxsyLmRuc2NyeXB0LWNlcnQuanAudGlhci5hcHA


## jp.tiar.app-doh

Non-Logging, Non-Filtering DNS-over-HTTPS server in Japan.
No ECS, Support DNSSEC

sdns://AgcAAAAAAAAADTE3Mi4xMDQuOTMuODCgMob_ZaZfrzIIXuoTiMNzi6fjeHPJBszjxKKLTMKliYigRE69Z7uD-IB7OSHpOKyReLiCvVCq2xEjHwRM9fCN984gzBBg05yDKbYrb7x9DW35MJhpuYHn5jktXNj6QI9NgOYLanAudGlhci5hcHAKL2Rucy1xdWVyeQ


## jp.tiar.app-doh-ipv6

Non-Logging, Non-Filtering DNS-over-HTTPS (IPv6) server in Japan.
No ECS, Support DNSSEC

sdns://AgcAAAAAAAAAIFsyNDAwOjg5MDI6OmYwM2M6OTFmZjpmZWRhOmM1MTRdoDKG_2WmX68yCF7qE4jDc4un43hzyQbM48Sii0zCpYmIoEROvWe7g_iAezkh6TiskXi4gr1QqtsRIx8ETPXwjffOIMwQYNOcgym2K2-8fQ1t-TCYabmB5-Y5LVzY-kCPTYDmC2pwLnRpYXIuYXBwCi9kbnMtcXVlcnk


## jp.tiar.app-ipv6

Non-Logging, Non-Filtering DNSCrypt (IPv6) server in Japan.
No ECS, Support DNSSEC

sdns://AQcAAAAAAAAAJVsyNDAwOjg5MDI6OmYwM2M6OTFmZjpmZWRhOmM1MTRdOjE0NDMgMrh2PvG_ZTah3moTwM1vSKF544i2maU3jW1bPEzfVFMbMi5kbnNjcnlwdC1jZXJ0LmpwLnRpYXIuYXBw


## jp.tiarap.org

DNS-over-HTTPS Server. Non-Logging, Non-Filtering, No ECS, Support DNSSEC.
Cached via Cloudflare.

sdns://AgcAAAAAAAAAACDlfOUtFRBtpOz9nhH9pf0dHgpr4BIkGITTYuodSvRk9w1qcC50aWFyYXAub3JnCi9kbnMtcXVlcnk


## jp.tiarap.org-ipv6

DNS-over-HTTPS Server (IPv6). Non-Logging, Non-Filtering, No ECS, Support DNSSEC.
Cached via Cloudflare.

sdns://AgcAAAAAAAAAG1syNjA2OjQ3MDA6MzAzNjo6NjgxYjo5NmFhXSDlfOUtFRBtpOz9nhH9pf0dHgpr4BIkGITTYuodSvRk9w1qcC50aWFyYXAub3JnCi9kbnMtcXVlcnk


## libredns

DoH server in Germany. No logging, but no DNS padding and no DNSSEC support.
https://libredns.gr/

sdns://AgYAAAAAAAAAAKAyhv9lpl-vMghe6hOIw3OLp-N4c8kGzOPEootMwqWJiKBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33ziDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5g9kb2gubGlicmVkbnMuZ3IKL2Rucy1xdWVyeQ


## libredns-noads

DoH server in Germany. No logging, but no DNS padding and no DNSSEC support.
no ads version, uses StevenBlack's host list: https://github.com/StevenBlack/hosts

sdns://AgIAAAAAAAAAAKAyhv9lpl-vMghe6hOIw3OLp-N4c8kGzOPEootMwqWJiKBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33ziDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5g9kb2gubGlicmVkbnMuZ3IEL2Fkcw


## linuxsec-adguard

DoH server with AdGuard Home for ads blocking.
Operated by LinuxSec.

Non-logging, AD-filtering, supports DNSSEC.

https://doh.linuxsec.org

sdns://AgMAAAAAAAAADDEwMy44NS4xNS42MKAyhv9lpl-vMghe6hOIw3OLp-N4c8kGzOPEootMwqWJiKBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33ziDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5hBkb2gubGludXhzZWMub3JnCC9hZGd1YXJk


## linuxsec-doh

DoH server operated by LinuxSec.
Server located in the Indonesia.

Non-logging, supports DNSSEC.

https://doh.linuxsec.org

sdns://AgcAAAAAAAAADDEwMy44NS4xNS42MKAyhv9lpl-vMghe6hOIw3OLp-N4c8kGzOPEootMwqWJiKBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33ziDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5hBkb2gubGludXhzZWMub3JnCi9kbnMtcXVlcnk


## meganerd

DNSCrypt server by MegaNerd.nl (IPv4) - https://meganerd.nl/encrypted-dns-server
Hosted in Amsterdam (AMS1), The Netherlands.

Non-logging, non-filtering, supports DNSSEC.

sdns://AQcAAAAAAAAADjEzNi4yNDQuOTcuMTE0IPyq3HBOXuNgu6FO4pU71Si6CTV6kPD85NA6AThr_6tiGDIuZG5zY3J5cHQtY2VydC5tZWdhbmVyZA


## meganerd-doh-ipv6

DoH server by MegaNerd.nl (IPv6) - https://meganerd.nl/encrypted-dns-server
Hosted in Amsterdam (AMS1), The Netherlands.

Non-logging, non-filtering, supports DNSSEC.

sdns://AgcAAAAAAAAAKFsyMDAxOjE5ZjA6NTAwMTpjYmI6NTQwMDowM2ZmOmZlMDc6ZjcwZF2gMob_ZaZfrzIIXuoTiMNzi6fjeHPJBszjxKKLTMKliYigRE69Z7uD-IB7OSHpOKyReLiCvVCq2xEjHwRM9fCN984gzBBg05yDKbYrb7x9DW35MJhpuYHn5jktXNj6QI9NgOYZY2hld2JhY2NhLm1lZ2FuZXJkLm5sOjQ0MwQvZG9o


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


## moulticast-fr-ipv4

Public | Non-filtering | Non-logging | DNSSEC aware | Hosted in France | Operated by @herver (Github) | https://moulticast.net/dnscrypt/

sdns://AQcAAAAAAAAADTE5My4yMDAuNDMuMjUgdahEBxLVMuwDPeuVlJbOFrrDypHdrAlyIm7FrDgPyMgdMi5kbnNjcnlwdC1jZXJ0Lm1vdWx0aWNhc3QtZnI


## moulticast-sg-ipv4

Public | Non-filtering | Non-logging | DNSSEC aware | Hosted in Singapore | Operated by @herver (Github) | https://moulticast.net/dnscrypt/

sdns://AQcAAAAAAAAADTUxLjc5LjE1OC4xODMgL4fe2yQRI0ljCMe05zhSmeKGYz-izKVjyE040Ovq-gUdMi5kbnNjcnlwdC1jZXJ0Lm1vdWx0aWNhc3Qtc2c


## moulticast-sg-ipv6

Public | Non-filtering | Non-logging | DNSSEC aware | Hosted in Singapore | Operated by @herver (Github) | https://moulticast.net/dnscrypt/

sdns://AQcAAAAAAAAAGlsyNDAyOjFmMDA6ODAwMDo4MDA6OjEyYmFdIC-H3tskESNJYwjHtOc4UpnihmM_osylY8hNONDr6voFHTIuZG5zY3J5cHQtY2VydC5tb3VsdGljYXN0LXNn


## moulticast-uk-ipv4

Public | Non-filtering | Non-logging | DNSSEC aware | Hosted in UK | Operated by @herver (Github) | https://moulticast.net/dnscrypt/

sdns://AQcAAAAAAAAADjUxLjE5NS4yMDAuMTgyIOdFe4OLrR_kCKC-8omGs5my5qxIyBgkldWZoSUmYvCNHTIuZG5zY3J5cHQtY2VydC5tb3VsdGljYXN0LXVr


## mullvad-adblock-doh

Public ad-blocking, non-logging (audited), DNSSEC-capable, DNS-over-HTTPS resolver hosted by VPN provider Mullvad
Anycast IPv4/IPv6 with servers in SE, DE, UK, US, AU and SG
https://mullvad.net/en/help/dns-over-https-and-dns-over-tls/

sdns://AgMAAAAAAAAAACD5_zfwLmMstzhwJcB-V5CKPTcbfJXYzdA5DeIx7ZQ6EhdhZGJsb2NrLmRvaC5tdWxsdmFkLm5ldAovZG5zLXF1ZXJ5


## mullvad-doh

Public non-filtering, non-logging (audited), DNSSEC-capable, DNS-over-HTTPS resolver hosted by VPN provider Mullvad
Anycast IPv4/IPv6 with servers in SE, DE, UK, US, AU and SG
https://mullvad.net/en/help/dns-over-https-and-dns-over-tls/

sdns://AgcAAAAAAAAAACD5_zfwLmMstzhwJcB-V5CKPTcbfJXYzdA5DeIx7ZQ6Eg9kb2gubXVsbHZhZC5uZXQKL2Rucy1xdWVyeQ


## nextdns

NextDNS is a cloud-based private DNS service that gives you full control
over what is allowed and what is blocked on the Internet.

DNSSEC, Anycast, Non-logging, NoFilters

https://www.nextdns.io/

sdns://AgcAAAAAAAAACjQ1LjkwLjI4LjCgMob_ZaZfrzIIXuoTiMNzi6fjeHPJBszjxKKLTMKliYigRE69Z7uD-IB7OSHpOKyReLiCvVCq2xEjHwRM9fCN984gzBBg05yDKbYrb7x9DW35MJhpuYHn5jktXNj6QI9NgOYOZG5zLm5leHRkbnMuaW8PL2Ruc2NyeXB0LXByb3h5


## nextdns-ipv6

Connects to NextDNS over IPv6.

NextDNS is a cloud-based private DNS service that gives you full control
over what is allowed and what is blocked on the Internet.

DNSSEC, Anycast, Non-logging, NoFilters

https://www.nextdns.io/

sdns://AgcAAAAAAAAADVsyYTA3OmE4YzA6Ol2gMob_ZaZfrzIIXuoTiMNzi6fjeHPJBszjxKKLTMKliYigRE69Z7uD-IB7OSHpOKyReLiCvVCq2xEjHwRM9fCN984gzBBg05yDKbYrb7x9DW35MJhpuYHn5jktXNj6QI9NgOYOZG5zLm5leHRkbnMuaW8PL2Ruc2NyeXB0LXByb3h5


## njalla-doh

Non-logging DoH server in Sweden operated by Njalla.

https://dns.njal.la/

sdns://AgcAAAAAAAAADDk1LjIxNS4xOS41M6Ayhv9lpl-vMghe6hOIw3OLp-N4c8kGzOPEootMwqWJiKBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33ziDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5gtkbnMubmphbC5sYQovZG5zLXF1ZXJ5


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


## plan9-ns1

DNSCrypt server in New Jersey, USA. Non-logging, non-filtering, DNSSEC, anonymized.
Running the official Docker image, info - https://jlongua.github.io/plan9-dns/

sdns://AQcAAAAAAAAADjE3My4xOTkuMTI2LjM1IJLwH3z8-G6TDyEk6yXMmGxuss7nGfFnkHrtm4VOSc9SGTIuZG5zY3J5cHQtY2VydC5wbGFuOS1kbnM


## plan9-ns2

DNSCrypt server in Florida, USA. Non-logging, non-filtering, DNSSEC, anonymized.
info - https://jlongua.github.io/plan9-dns/

sdns://AQcAAAAAAAAAEjQ1LjYzLjExMC4xODc6ODQ0MyCcjeRhPcJTsKhZ8iViALPd39CussG6SnprFT9z_1f03x0yLmRuc2NyeXB0LWNlcnQucGxhbjktbnMyLmNvbQ


## plan9-ns2-doh

doh server in Florida, USA. Non-logging, non-filtering, DNSSEC.
info - https://jlongua.github.io/plan9-dns/

sdns://AgcAAAAAAAAADTQ1LjYzLjExMC4xODegMob_ZaZfrzIIXuoTiMNzi6fjeHPJBszjxKKLTMKliYigRE69Z7uD-IB7OSHpOKyReLiCvVCq2xEjHwRM9fCN984gzBBg05yDKbYrb7x9DW35MJhpuYHn5jktXNj6QI9NgOYTZHJhY28ucGxhbjktbnMyLmNvbQovZG5zLXF1ZXJ5


## plan9-ns2-doh-ipv6

doh server in Florida, USA. Non-logging, non-filtering, DNSSEC.
info - https://jlongua.github.io/plan9-dns/

sdns://AgcAAAAAAAAAHVsyMDAxOjE5ZjA6OTAwMjoxZDc0OjU0MDA6OjFdoDKG_2WmX68yCF7qE4jDc4un43hzyQbM48Sii0zCpYmIoEROvWe7g_iAezkh6TiskXi4gr1QqtsRIx8ETPXwjffOIMwQYNOcgym2K2-8fQ1t-TCYabmB5-Y5LVzY-kCPTYDmE2RyYWNvLnBsYW45LW5zMi5jb20KL2Rucy1xdWVyeQ


## powerdns-doh

By PowerDNS/Open-Xchange https://powerdns.org

sdns://AgcAAAAAAAAAAKAyhv9lpl-vMghe6hOIw3OLp-N4c8kGzOPEootMwqWJiKBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33ziDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5hBkb2gucG93ZXJkbnMub3JnAS8


## pryv8boi

By pryv8, non Logging, uncensored, DNSSEC - hosted on contabo servers

sdns://AQcAAAAAAAAAEzE2NC42OC4xMjEuMTYyOjQ0NDMg0I13MMoiCKduiHKph6yaHFtoNLhPFroCyPOffUqqcsocMi5kbnNjcnlwdC1jZXJ0LnByeXY4Ym9pLm9yZw


## publicarray-au

DNSSEC • OpenNIC • Non-logging • Uncensored - hosted on vultr.com
Maintained by publicarray - https://dns.seby.io

sdns://AQcAAAAAAAAADDQ1Ljc2LjExMy4zMSAIVGh4i6eKXqlF6o9Fg92cgD2WcDvKQJ7v_Wq4XrQsVhsyLmRuc2NyeXB0LWNlcnQuZG5zLnNlYnkuaW8


## publicarray-au-doh

DNSSEC • OpenNIC • Non-logging • Uncensored - hosted on vultr.com
Maintained by publicarray - https://dns.seby.io

sdns://AgcAAAAAAAAADDQ1Ljc2LjExMy4zMaAyhv9lpl-vMghe6hOIw3OLp-N4c8kGzOPEootMwqWJiKBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33ziDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5hBkb2guc2VieS5pbzo4NDQzCi9kbnMtcXVlcnk


## publicarray-au2

DNSSEC • OpenNIC • Non-logging • Uncensored - hosted on ovh.com.au
Maintained by publicarray - https://dns.seby.io

sdns://AQcAAAAAAAAAEjEzOS45OS4yMjIuNzI6ODQ0MyB3NPVeil3tZfA7hH0fz-MBp7VSd6pprpccj78c5dOjFRsyLmRuc2NyeXB0LWNlcnQuZG5zLnNlYnkuaW8


## publicarray-au2-doh

DNSSEC • OpenNIC • Non-logging • Uncensored - hosted on ovh.com.au
Maintained by publicarray - https://dns.seby.io

sdns://AgcAAAAAAAAADTEzOS45OS4yMjIuNzKgMob_ZaZfrzIIXuoTiMNzi6fjeHPJBszjxKKLTMKliYigRE69Z7uD-IB7OSHpOKyReLiCvVCq2xEjHwRM9fCN984gzBBg05yDKbYrb7x9DW35MJhpuYHn5jktXNj6QI9NgOYRZG9oLTIuc2VieS5pbzo0NDMKL2Rucy1xdWVyeQ


## pwoss.org-dnscrypt

No filter | No logs | DNSSEC | Nuremberg, Germany (netcup) | Maintained by https://pwoss.org/ (Dan)

sdns://AQcAAAAAAAAAEzQ1LjE0Mi4xNzYuMTcwOjQ0MzQgZF8FTkuFVFw6YP5x8ydebnTb8ONP7Ti6P0K1REyhUHgZMi5kbnNjcnlwdC1jZXJ0LnB3b3NzLm9yZw


## quad101

DNSSEC-aware public resolver by the Taiwan Network Information Center (TWNIC)
https://101.101.101.101/index_en.html

sdns://AgcAAAAAAAAAACC2vD25TAYM7EnyCH8Xw1-0g5OccnTsGH9vQUUH0njRtAxkbnMudHduaWMudHcKL2Rucy1xdWVyeQ


## quad9-dnscrypt-ip4-filter-pri

Quad9 (anycast) dnssec/no-log/filter 9.9.9.9 / 149.112.112.9

sdns://AQMAAAAAAAAADDkuOS45Ljk6ODQ0MyBnyEe4yHWM0SAkVUO-dWdG3zTfHYTAC4xHA2jfgh2GPhkyLmRuc2NyeXB0LWNlcnQucXVhZDkubmV0
sdns://AQMAAAAAAAAAEjE0OS4xMTIuMTEyLjk6ODQ0MyBnyEe4yHWM0SAkVUO-dWdG3zTfHYTAC4xHA2jfgh2GPhkyLmRuc2NyeXB0LWNlcnQucXVhZDkubmV0


## quad9-dnscrypt-ip4-nofilter-pri

Quad9 (anycast) no-dnssec/no-log/no-filter 9.9.9.10 / 149.112.112.10

sdns://AQYAAAAAAAAADTkuOS45LjEwOjg0NDMgZ8hHuMh1jNEgJFVDvnVnRt803x2EwAuMRwNo34Idhj4ZMi5kbnNjcnlwdC1jZXJ0LnF1YWQ5Lm5ldA
sdns://AQYAAAAAAAAAEzE0OS4xMTIuMTEyLjEwOjg0NDMgZ8hHuMh1jNEgJFVDvnVnRt803x2EwAuMRwNo34Idhj4ZMi5kbnNjcnlwdC1jZXJ0LnF1YWQ5Lm5ldA


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


## safesurfer

Family safety focused blocklist for over 2 million adult sites, as well as phishing and malware and more.
Free to use, paid for customizing blocking for more categories+sites and viewing usage at my.safesurfer.io. Logs taken for viewing
usage, data never sold - https://safesurfer.io

Warning: this server is incompatible with DNS anonymization.

sdns://AQMAAAAAAAAADzEwNC4xNTUuMjM3LjIyNSAnIH_VEgToNntINABd-f_R0wu-KpwzY55u2_iu2R1A2CAyLmRuc2NyeXB0LWNlcnQuc2FmZXN1cmZlci5jby5ueg


## saldns01-conoha-ipv4

Hosted on ConoHa VPS Tokyo region. No log. No filter.

sdns://AQcAAAAAAAAAFDExOC4yNy4xMDguMTQwOjUwNDQzIJ6mPj1aeYrSINYeUzVU2n9JMTUgalCmDR6OhSlzrFtfIjIuZG5zY3J5cHQtY2VydC5zYWxkbnMwMS50eXBlcS5vcmc


## saldns02-conoha-ipv4

Hosted on ConoHa VPS Tokyo region. No log. No filter.

sdns://AQcAAAAAAAAAFTEzMy4xMzAuMTE4LjEwMzo1MDQ0MyCo72atOPI72UkQcaypJBmUUd1btQTwm1VRxHLioOo19SIyLmRuc2NyeXB0LWNlcnQuc2FsZG5zMDIudHlwZXEub3Jn


## saldnssg01-conoha-ipv4

Hosted on ConoHa VPS Singapore region. No log. No filter.

sdns://AQcAAAAAAAAAFDE2My40NC4xNTQuMTQ0OjUwNDQzIGri6V6DYmNp01gdIzQyrACcU7_tbcUKEN8-P4B5NATAJTIuZG5zY3J5cHQtY2VydC5zYWxkbnMtc2cwMS50eXBlcS5vcmc


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


## sicher-surfen-de

Resolver in Frankfurt, Germany and New York, United States. Anycast. More locations following. DNSCrypt server. Non-logging, non-filtering, DNSSEC.

sdns://AQcAAAAAAAAAEDE4NS4yMzUuNjAuMTo0NDMgA-0Um9P9C564j6S8u0ij_DVITfFpSmvk7KIAfPFr3NcrMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeXB0LnNpY2hlci1zdXJmZW4udGVjaA


## sth-dnscrypt-se

Resolver in Stockholm, Sweden. DNSCrypt server. Non-logging, non-filtering, DNSSEC.

sdns://AQcAAAAAAAAAEjQ1LjE1My4xODcuOTY6NDM0MyB5eVqmG8yEeu7KgiIFlYUJCqBps33UPRtAQFvecn3JRh8yLmRuc2NyeXB0LWNlcnQuc3RoLWRuc2NyeXB0LXNl


## sth-doh-se

Resolver in Stockholm, Sweden. DoH server. Non-logging, non-filtering, DNSSEC.

sdns://AgcAAAAAAAAADTQ1LjE1My4xODcuOTagMob_ZaZfrzIIXuoTiMNzi6fjeHPJBszjxKKLTMKliYigRE69Z7uD-IB7OSHpOKyReLiCvVCq2xEjHwRM9fCN984gzBBg05yDKbYrb7x9DW35MJhpuYHn5jktXNj6QI9NgOYSZG5zc2UuYWxla2JlcmcubmV0Ci9kbnMtcXVlcnk


## switch

Public DoH service provided by SWITCH in Switzerland

https://www.switch.ch

Provides protection against malware, but does not block ads.

sdns://AgMAAAAAAAAAACA7FAZ73R-Sz_NJdsl8yiP0L0H3yntGnmJAVtqSpC3Nmg1kbnMuc3dpdGNoLmNoCi9kbnMtcXVlcnk


## t53

Deutsche Telekom experimental DoH server.

sdns://AgUAAAAAAAAAAKAyhv9lpl-vMghe6hOIw3OLp-N4c8kGzOPEootMwqWJiKBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33ziDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5gpkbnMudDUzLmRlCi9kbnMtcXVlcnk


## userspace-australia

Public DNSCrypt in Australia (Brisbane & Melbourne) by UserSpace. No logging, currently uncensored, IPv4.

sdns://AQcAAAAAAAAAEjEwMy4yMzYuMTYyLjExOTo1NCC82yDYAafAl4ht8sAbOgn3TjyhAbdTakx92VaUFLBwux0yLmRuc2NyeXB0LWNlcnQudXNlcnNwYWNlLW1lbA
sdns://AQcAAAAAAAAAEDEwMy4xNi4xMzEuNzc6NTQgEWOI99_3M0Rk_Jsm7KQtUjC1612HdCwSWA7_ftojHG4dMi5kbnNjcnlwdC1jZXJ0LnVzZXJzcGFjZS1ibmU


## userspace-australia-ipv6

Public DNSCrypt in Australia (Brisbane & Melbourne) by UserSpace. No logging, currently uncensored, IPv6.

sdns://AQcAAAAAAAAAJVsyNDA0Ojk0MDA6MzowOjIxNjozZWZmOmZlZTA6N2Y2OV06NTQgvNsg2AGnwJeIbfLAGzoJ9048oQG3U2pMfdlWlBSwcLsdMi5kbnNjcnlwdC1jZXJ0LnVzZXJzcGFjZS1tZWw
sdns://AQcAAAAAAAAAJVsyNDA0Ojk0MDA6MTowOjIxNjozZWZmOmZlZjA6MTgwYV06NTQgEWOI99_3M0Rk_Jsm7KQtUjC1612HdCwSWA7_ftojHG4dMi5kbnNjcnlwdC1jZXJ0LnVzZXJzcGFjZS1ibmU


## v.dnscrypt.uk-ipv4

DNSCrypt v2, no logs, uncensored, DNSSEC. Hosted in London UK on Vultr
https://www.dnscrypt.uk

sdns://AQcAAAAAAAAADzEwNC4yMzguMTg2LjE5MiDtST2M6teQZk8GPEe3lZojaS18kDY8nkPMtZF75bQe5R0yLmRuc2NyeXB0LWNlcnQudi5kbnNjcnlwdC51aw


## v.dnscrypt.uk-ipv6

DNSCrypt v2, no logs, uncensored, DNSSEC. Hosted in London UK on Vultr
https://www.dnscrypt.uk

sdns://AQcAAAAAAAAAKFsyMDAxOjE5ZjA6NzQwMjoxNTc0OjU0MDA6MmZmOmZlNjY6MmNmZl0g7Uk9jOrXkGZPBjxHt5WaI2ktfJA2PJ5DzLWRe-W0HuUdMi5kbnNjcnlwdC1jZXJ0LnYuZG5zY3J5cHQudWs


## wevpn-adblock-singapore

Private DNSCrypt service by WeVPN in Singapore - https://wevpn.com/dns
With ad blocking.

sdns://AQMAAAAAAAAAEzE0My4yNDQuMzMuOTA6MTUzNTMgufWZBK6y8zr6mDW3z47IASY0jQoaoRFSBaB8i1GHng8fMi5kbnNjcnlwdC1jZXJ0LnNlY3VyZS5kbnMudGVzdA


## wevpn-adblock-uk

Private DNSCrypt service by WeVPN in London, UK - https://wevpn.com/dns
With ad blocking.

sdns://AQMAAAAAAAAAFDIxNi4xMTkuMTU1LjQ5OjE1MzUzIF2QDfN8-gL0x3IyXgqgrawgSxKm7A80BgexEH7WghsBHzIuZG5zY3J5cHQtY2VydC5zZWN1cmUuZG5zLnRlc3Q


## wevpn-adblock-useast

Private DNSCrypt service by WeVPN in New Jersey, USA - https://wevpn.com/dns
With ad blocking.

sdns://AQMAAAAAAAAAFDIzLjIyNi4xMzQuMjQzOjE1MzUzIBf1guPReaoN_V7w-UelcO4YnvqoXnDGsxRPIRc6uSjsHzIuZG5zY3J5cHQtY2VydC5zZWN1cmUuZG5zLnRlc3Q


## wevpn-adblock-uswest

Private DNSCrypt service by WeVPN in Los Angeles, USA - https://wevpn.com/dns
With ad blocking.

sdns://AQMAAAAAAAAAEjcyLjExLjEzNC45MToxNTM1MyBV2R9ORyXQIRO7dlT0F2mYeEdvXEuBOWMdFkripZcN9B8yLmRuc2NyeXB0LWNlcnQuc2VjdXJlLmRucy50ZXN0


## wevpn-singapore

Private DNSCrypt service by WeVPN in Singapore - https://wevpn.com/dns
With ad blocking.

sdns://AQcAAAAAAAAAEzE0My4yNDQuMzMuNzQ6MTUzNTMgFTXwu5MfYkBOrRpDeoB-yOWEjCnf-l3yixhtuzuPBskfMi5kbnNjcnlwdC1jZXJ0LnNlY3VyZS5kbnMudGVzdA


## wevpn-uk

Private DNSCrypt service by WeVPN in London, UK - https://wevpn.com/dns
With ad blocking.

sdns://AQcAAAAAAAAAEjIxMi43OC45NC40MDoxNTM1MyCAw5p4sJ073gZnQ5jy00DHU3r7Y9mopz-_idDV_HHuuR8yLmRuc2NyeXB0LWNlcnQuc2VjdXJlLmRucy50ZXN0


## wevpn-useast

Private DNSCrypt service by WeVPN in New Jersey, USA - https://wevpn.com/dns
With ad blocking.

sdns://AQcAAAAAAAAAFDIzLjIyNi4xMzQuMjQyOjE1MzUzII_Le5DiGa3AfdRxR7DRt52ZaexL_22aLfjDJwp5saIsHzIuZG5zY3J5cHQtY2VydC5zZWN1cmUuZG5zLnRlc3Q


## wevpn-uswest

Private DNSCrypt service by WeVPN in Los Angeles, USA - https://wevpn.com/dns
With ad blocking.

sdns://AQcAAAAAAAAAEjcyLjExLjEzNC45MDoxNTM1MyAKLsInrJLgKMxBqSL1VvH74T3wwp1bn5wkvPwUlea3Kh8yLmRuc2NyeXB0LWNlcnQuc2VjdXJlLmRucy50ZXN0


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


## zackptg5-us-pit-ipv4

DNSSEC/unfiltered/non-logged. Hosted on TeraSwitch in Pittsburgh, PA. Running the official Docker image by @zackptg5

sdns://AQcAAAAAAAAAETc0LjExOC4xNDEuNDc6NDQyILLrUp55NwetKjM1XaPs19NIKDIbD1w0K-lCK33_gKcgHzIuZG5zY3J5cHQtY2VydC56YWNrcHRnNS11cy1waXQ


## zackptg5-us-pit-ipv6

DNSSEC/unfiltered/non-logged. Hosted on TeraSwitch in Pittsburgh, PA. Running the official Docker image by @zackptg5

sdns://AQcAAAAAAAAAJ1syNjA3OmZkYzA6MjowOmY4MTY6M2VmZjpmZTMyOmM3YjRdOjQ0MiCy61KeeTcHrSozNV2j7NfTSCgyGw9cNCvpQit9_4CnIB8yLmRuc2NyeXB0LWNlcnQuemFja3B0ZzUtdXMtcGl0

