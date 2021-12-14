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

sdns://AgcAAAAAAAAADTIxNy4xNjkuMjAuMjOgzBBg05yDKbYrb7x9DW35MJhpuYHn5jktXNj6QI9NgOYgRE69Z7uD-IB7OSHpOKyReLiCvVCq2xEjHwRM9fCN984NZG5zLmFhLm5ldC51awovZG5zLXF1ZXJ5


## acsacsar-ams-ipv4

Public non-censoring, non-logging, DNSSEC-capable, DNSCrypt-enabled DNS resolver hosted on Scaleway by @acsacsar (twitter)

sdns://AQcAAAAAAAAADTUxLjE1OC4xNjYuOTcgAyfzz5J-mV9G-yOB4Hwcdk7yX12EQs5Iva7kV3oGtlEgMi5kbnNjcnlwdC1jZXJ0LmFjc2Fjc2FyLWFtcy5jb20


## acsacsar-ams-ipv6

Public non-censoring, non-logging, DNSSEC-capable, DNSCrypt-enabled DNS resolver hosted on Scaleway by @acsacsar (twitter)

sdns://AQcAAAAAAAAAFlsyMDAxOmJjODoxODI0OjczODo6MV0gAyfzz5J-mV9G-yOB4Hwcdk7yX12EQs5Iva7kV3oGtlEgMi5kbnNjcnlwdC1jZXJ0LmFjc2Fjc2FyLWFtcy5jb20


## adfree.usableprivacy.net

Public updns DoH service with advertising, tracker and malware filters.

Hosted in Europe by @usableprivacy, details see: https://docs.usableprivacy.com

sdns://AgMAAAAAAAAADzE0OS4xNTQuMTUzLjE1M6DMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5iBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33zhhhZGZyZWUudXNhYmxlcHJpdmFjeS5uZXQKL2Rucy1xdWVyeQ


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

sdns://AgMAAAAAAAAADTQ1Ljc5LjEyMC4yMzOgzBBg05yDKbYrb7x9DW35MJhpuYHn5jktXNj6QI9NgOYgRE69Z7uD-IB7OSHpOKyReLiCvVCq2xEjHwRM9fCN984RZG9oLmluLmFoYWRucy5uZXQKL2Rucy1xdWVyeQ


## ahadns-doh-in-ipv6

A zero logging DNS with support for DNS-over-HTTPS (DoH) & DNS-over-TLS (DoT). Blocks ads, malware, trackers, viruses, ransomware, telemetry and more. No persistent logs. DNSSEC. Hosted in Mumbai, India. By https://ahadns.com/
Server statistics can be seen at: https://statistics.ahadns.com/?server=in

sdns://AgMAAAAAAAAAF1syNDAwOjg5MDQ6ZTAwMTo0Mzo6NDNdoMwQYNOcgym2K2-8fQ1t-TCYabmB5-Y5LVzY-kCPTYDmIEROvWe7g_iAezkh6TiskXi4gr1QqtsRIx8ETPXwjffOEWRvaC5pbi5haGFkbnMubmV0Ci9kbnMtcXVlcnk


## ahadns-doh-la

A zero logging DNS with support for DNS-over-HTTPS (DoH) & DNS-over-TLS (DoT). Blocks ads, malware, trackers, viruses, ransomware, telemetry and more. No persistent logs. DNSSEC. Hosted in Los Angeles, USA. By https://ahadns.com/
Server statistics can be seen at: https://statistics.ahadns.com/?server=la

sdns://AgMAAAAAAAAADTQ1LjY3LjIxOS4yMDigzBBg05yDKbYrb7x9DW35MJhpuYHn5jktXNj6QI9NgOYgRE69Z7uD-IB7OSHpOKyReLiCvVCq2xEjHwRM9fCN984RZG9oLmxhLmFoYWRucy5uZXQKL2Rucy1xdWVyeQ


## ahadns-doh-la-ipv6

A zero logging DNS with support for DNS-over-HTTPS (DoH) & DNS-over-TLS (DoT). Blocks ads, malware, trackers, viruses, ransomware, telemetry and more. No persistent logs. DNSSEC. Hosted in Los Angeles, USA. By https://ahadns.com/
Server statistics can be seen at: https://statistics.ahadns.com/?server=la

sdns://AgMAAAAAAAAAFlsyYTA0OmJkYzc6MTAwOjcwOjo3MF2gzBBg05yDKbYrb7x9DW35MJhpuYHn5jktXNj6QI9NgOYgRE69Z7uD-IB7OSHpOKyReLiCvVCq2xEjHwRM9fCN984RZG9oLmxhLmFoYWRucy5uZXQKL2Rucy1xdWVyeQ


## ahadns-doh-nl

A zero logging DNS with support for DNS-over-HTTPS (DoH) & DNS-over-TLS (DoT). Blocks ads, malware, trackers, viruses, ransomware, telemetry and more. No persistent logs. DNSSEC. Hosted in Amsterdam, Netherlands. By https://ahadns.com/
Server statistics can be seen at: https://statistics.ahadns.com/?server=nl

sdns://AgMAAAAAAAAACTUuMi43NS43NaDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5iBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33zhFkb2gubmwuYWhhZG5zLm5ldAovZG5zLXF1ZXJ5


## ahadns-doh-nl-ipv6

A zero logging DNS with support for DNS-over-HTTPS (DoH) & DNS-over-TLS (DoT). Blocks ads, malware, trackers, viruses, ransomware, telemetry and more. No persistent logs. DNSSEC. Hosted in Amsterdam, Netherlands. By https://ahadns.com/
Server statistics can be seen at: https://statistics.ahadns.com/?server=nl

sdns://AgMAAAAAAAAAFlsyYTA0OjUyYzA6MTAxOjc1Ojo3NV2gzBBg05yDKbYrb7x9DW35MJhpuYHn5jktXNj6QI9NgOYgRE69Z7uD-IB7OSHpOKyReLiCvVCq2xEjHwRM9fCN984RZG9oLm5sLmFoYWRucy5uZXQKL2Rucy1xdWVyeQ


## alidns-doh

A public DNS resolver that supports DoH/DoT in mainland China, provided by Alibaba-Cloud.

Warning: GFW filtering rules are applied by that resolver.

Homepage: https://alidns.com/

sdns://AgAAAAAAAAAACTIyMy41LjUuNSCoF6cUD2dwqtorNi96I2e3nkHPSJH1ka3xbdOglmOVkQ5kbnMuYWxpZG5zLmNvbQovZG5zLXF1ZXJ5


## altername

Protocol: DNSCrypt | Features: Non-logging, Non-filtering, DNSSEC, EmerDNS | Location: Moscow, Russia

sdns://AQcAAAAAAAAAEDE4NS4yMDQuMi44Ojg0NDMgmfGTecriYs8VCfSDIUo6U6yQ5g98itBT7FWH3yYVB1MZMi5kbnNjcnlwdC1jZXJ0LmFsdGVybmFtZQ


## altername-ipv6

Protocol: DNSCrypt IPv6 | Features: Non-logging, Non-filtering, DNSSEC, EmerDNS | Location: Moscow, Russia

sdns://AQcAAAAAAAAAG1syYTA0OjUyMDA6ZmZmNDo6MTNmZl06ODQ0MyCZ8ZN5yuJizxUJ9IMhSjpTrJDmD3yK0FPsVYffJhUHUxkyLmRuc2NyeXB0LWNlcnQuYWx0ZXJuYW1l


## ams-ads-doh-nl

Resolver in Amsterdam. DoH protocol. Non-logging. Blocks ads, malware and trackers. DNSSEC enabled.

sdns://AgMAAAAAAAAADTUxLjE1LjEyNC4yMDgAGGRuc25sLW5vYWRzLmFsZWtiZXJnLm5ldAovZG5zLXF1ZXJ5


## ams-dnscrypt-nl

Resolver in Amsterdam. Dnscrypt protocol. Non-logging, non-filtering, DNSSEC.

sdns://AQcAAAAAAAAADTUxLjE1LjEyNC4yMDggvBOI9XY-Pex18ckPyRW_torq3-scRfaYUYBjnGDn8WIfMi5kbnNjcnlwdC1jZXJ0LmFtcy1kbnNjcnlwdC1ubA


## ams-dnscrypt-nl-ipv6

Resolver in Amsterdam. Dnscrypt protocol. Non-logging, non-filtering, DNSSEC.

sdns://AQcAAAAAAAAAF1syMDAxOmJjODoxODMwOjIwMTg6OjFdILwTiPV2Pj3sdfHJD8kVv7aK6t_rHEX2mFGAY5xg5_FiHzIuZG5zY3J5cHQtY2VydC5hbXMtZG5zY3J5cHQtbmw


## ams-doh-nl

Resolver in Amsterdam. DoH protocol. Non-logging, non-filtering, DNSSEC.

sdns://AgcAAAAAAAAADTUxLjE1LjEyNC4yMDigzBBg05yDKbYrb7x9DW35MJhpuYHn5jktXNj6QI9NgOYgRE69Z7uD-IB7OSHpOKyReLiCvVCq2xEjHwRM9fCN984SZG5zbmwuYWxla2JlcmcubmV0Ci9kbnMtcXVlcnk


## ams-doh-nl-ipv6

Resolver in Amsterdam. DoH protocol. Non-logging, non-filtering, DNSSEC.

sdns://AgcAAAAAAAAAF1syMDAxOmJjODoxODMwOjIwMTg6OjFdoMwQYNOcgym2K2-8fQ1t-TCYabmB5-Y5LVzY-kCPTYDmIEROvWe7g_iAezkh6TiskXi4gr1QqtsRIx8ETPXwjffOEmRuc25sLmFsZWtiZXJnLm5ldAovZG5zLXF1ZXJ5


## att

AT&T test DoH server.

sdns://AgQAAAAAAAAAAKBLTrSwdCmLgotcADCVoQtFI_uVHAyINIsJxT5bq6QIoyD2Hldod9qWUClMzLX5bHX8txvaG7xGRjZ8Tr7aidcxjxBkb2h0cmlhbC5hdHQubmV0Ci9kbnMtcXVlcnk


## bcn-dnscrypt

Resolver in Barcelona, Spain. DNSCrypt protocol. Non-logging, non-filtering, DNSSEC.

sdns://AQcAAAAAAAAAEjE4NS4yNTMuMTU0LjY2OjQ0MyDgFv4inX3HfWsvXsEfLcUEsE_iz9m3KMRh7-MMU8hrTRwyLmRuc2NyeXB0LWNlcnQuYmNuLWRuc2NyeXB0


## bcn-doh

Resolver in Barcelona, Spain. DoH protocol. Non-logging, non-filtering, DNSSEC.

sdns://AgcAAAAAAAAADjE4NS4yNTMuMTU0LjY2oMwQYNOcgym2K2-8fQ1t-TCYabmB5-Y5LVzY-kCPTYDmIEROvWe7g_iAezkh6TiskXi4gr1QqtsRIx8ETPXwjffOEmRuc2VzLmFsZWtiZXJnLm5ldAovZG5zLXF1ZXJ5


## bortzmeyer

Non-logging DoH server in France operated by Stéphane Bortzmeyer.

https://www.bortzmeyer.org/doh-bortzmeyer-fr-policy.html

sdns://AgcAAAAAAAAADDE5My43MC44NS4xMaDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5iBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33zhFkb2guYm9ydHptZXllci5mcgEv


## bortzmeyer-ipv6

Non-logging DoH server in France operated by Stéphane Bortzmeyer (IPv6 only).

https://www.bortzmeyer.org/doh-bortzmeyer-fr-policy.html

sdns://AgcAAAAAAAAAGVsyMDAxOjQxZDA6MzAyOjIyMDA6OjE4MF2gzBBg05yDKbYrb7x9DW35MJhpuYHn5jktXNj6QI9NgOYgRE69Z7uD-IB7OSHpOKyReLiCvVCq2xEjHwRM9fCN984RZG9oLmJvcnR6bWV5ZXIuZnIBLw


## brahma-world

DNS-over-HTTPS server. Non Logging, filters ads, trackers and malware. DNSSEC ready, QNAME Minimization, No EDNS Client-Subnet.

Hosted in Stockholm, Sweden. (https://dns.brahma.world)

sdns://AgMAAAAAAAAADDEzLjUzLjE4Ni41MqDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5iBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33zhBkbnMuYnJhaG1hLndvcmxkCi9kbnMtcXVlcnk


## brahma-world-ipv6

DNS-over-HTTPS server. Non Logging, filters ads, trackers and malware. DNSSEC ready, QNAME Minimization, No EDNS Client-Subnet.

Hosted in Stockholm, Sweden. (https://dns.brahma.world)

sdns://AgMAAAAAAAAAJ1syYTA1OmQwMTY6YWY4OjQwMDA6NzcxMDo2ZmM6YmRlMzpmZTBlXaDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5iBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33zhBkbnMuYnJhaG1hLndvcmxkCi9kbnMtcXVlcnk


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

sdns://AQIAAAAAAAAADTk0LjE5OC40MS4yMzUgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz


## cs-barcelona

Barcelona, Spain DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADjM3LjEyMC4xNDIuMTE1IDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-belgium

Brussels, Belgium DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADTM3LjEyMC4yMzYuMTEgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz


## cs-berlin

Berlin, Germany DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADTM3LjEyMC4yMTcuNzUgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz


## cs-brazil

Sao Paulo, Brazil DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADjE3Ny41NC4xNDUuMTMxIDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-bulgaria

Sofia, Bulgaria DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADjM3LjEyMC4xNTIuMjM1IDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-ch

Switzerland DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAACzgxLjE3LjMxLjM0IDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-czech

Prague, Czech Republic DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADzIxNy4xMzguMjIwLjI0MyAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM


## cs-dc

US - Washington, DC DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADDE5OC43LjU4LjIyNyAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM


## cs-de

Frankfurt, Germany DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAACzE0Ni43MC44Mi4zIDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-dk

Denmark DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADTM3LjEyMC4yMzIuNDMgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz


## cs-dus1

Dusseldorf, Germany 1 DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADjIxMy4yMDIuMjE2LjEyIDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-dus2

Dusseldorf, Germany 2 DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADzIxMy4yMDIuMjE2LjIzNiAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM


## cs-dus3

Dusseldorf, Germany 3 DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADjg5LjE2My4yMjEuMTgxIDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-dus4

Dusseldorf, Germany 4 DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADjg1LjExNC4xMzguMTE5IDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-finland

Finland DNSCrypt server provided by https://cryptostorm.is

sdns://AQIAAAAAAAAADjE4NS4xMTcuMTE4LjIwIDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-fl

US - Miami, FL DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADDIzLjE5LjExNy41NSAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM


## cs-fl2

US - Miami, FL 2 DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADDIzLjE5LjExNy41MSAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM


## cs-fr

France DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADTE2My4xNzIuMzQuNTYgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz


## cs-free

France DNSCrypt server, running on our free server, provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADTE5NS4xNTQuNDAuNDggMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz


## cs-ga

US - Atlanta, GA DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADTY0LjQyLjE4MS4yMjcgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz


## cs-hungary

Budapest, Hungary DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADTg2LjEwNi43NC4yMTkgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz


## cs-il

US - Chicago, IL DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADjE3My4yMzQuNTYuMTE1IDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-il2

US - Chicago, IL 2 DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADjE3Mi4xMDcuMTk5LjE5IDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-india

India DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADzE2NS4yMzEuMjUzLjE2MyAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM


## cs-ireland

Dublin, Ireland DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADjM3LjEyMC4yMzUuMTg3IDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-is

Iceland DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADTgyLjIyMS4xMjguNDQgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz


## cs-la

US - Los Angeles, CA DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADDIzLjE5LjY3LjExNiAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM


## cs-london

London, England DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADTc4LjEyOS4xNDAuNjUgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz


## cs-lv

Latvia DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADzEwOS4yNDguMTQ5LjEzMyAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM


## cs-madrid

Madrid, Spain DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADjE4NS4xODMuMTA2LjgzIDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-manchester

Manchester, England DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADTE5NS4xMi40OC4xNzEgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz


## cs-md

Moldova DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADjE3OC4xNzUuMTI5LjE2IDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-mexico

Mexico City, Mexico DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADTEwMy4xNC4yNi4xOTAgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz


## cs-milan

Milan, Italy DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADzIxNy4xMzguMjE5LjIxOSAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM


## cs-montreal

Montreal, Canada DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADjE2Ny4xMTQuODQuMTMyIDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-montreal2

Montreal, Canada 2 DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADzE0NC4xNjguMTY1LjE2MyAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM


## cs-nc

US - North Carolina DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADjE1NS4yNTQuMjEuMjUwIDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-nl

Netherlands DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADTE4NS4xMDcuODAuODQgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz


## cs-nl2

Netherlands 2 DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADzE3Mi4xMDcuMjI2LjEwNyAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM


## cs-norway

Oslo, Norway DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADjkxLjIxOS4yMTUuMjI3IDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-nv

US - Las Vegas, NV DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADDM3LjEyMC4xNDcuMiAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM


## cs-nyc1

US - New York City, NY 1 DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADDIzLjE5LjI0NS44OCAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM


## cs-nyc2

US - New York City, NY 2 DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADDIzLjE5LjI0NS44NCAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM


## cs-ore

US - Oregon DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADTEwNC4yNTUuMTc1LjIgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz


## cs-poland

Warsaw, Poland DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADTM3LjEyMC4yMTEuOTEgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz


## cs-pt

Portugal DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADjkxLjIwNS4yMzAuMjI0IDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-ro

Romania DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADTE0Ni43MC42Ni4yMjcgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz


## cs-rome

Rome, Italy DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADjM3LjEyMC4yMDcuMTMxIDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-sea

US - Seattle, WA DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADDY0LjEyMC41LjI1MSAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM


## cs-serbia

Belgrade, Serbia DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADjM3LjEyMC4xOTMuMjE5IDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-sk

South Korea DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADzE3Mi4xMDcuMTU0LjIxMSAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM


## cs-slovakia

Bratislava, Slovakia DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADjE5My4zNy4yNTUuMjI3IDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-swe

Sweden DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADzEyOC4xMjcuMTA0LjEwOCAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM


## cs-sydney

Sydney, Australia DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADjM3LjEyMC4yMzQuMjUxIDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-tokyo

Tokyo, Japan DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADDE0Ni43MC4zMS40MyAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM


## cs-tx

US - Dallas, TX DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADTIwOS41OC4xNDcuMzYgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz


## cs-tx2

US - Dallas, TX 2 DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAACzQ1LjM1LjM1Ljk5IDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-tx3

US - Dallas, TX 3 DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAACzQ1LjM1LjcyLjQzIDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-vancouver

Vancouver, Canada DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADDcxLjE5LjI1MS4zNCAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM


## cz.nic

CZ.NIC's open DNSSEC validating resolvers in Prague, Czech Republic.
CZ.NIC resolvers neither collect any personal data nor gather
information on pages where your computer sends personal data.
https://www.nic.cz/odvr/

sdns://AgcAAAAAAAAADDE4NS40My4xMzUuMaDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5iBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33zgtvZHZyLm5pYy5jegQvZG9o


## cz.nic-ipv6

CZ.NIC's open DNSSEC validating resolvers in Prague, Czech Republic (IPv6 only).
CZ.NIC resolvers neither collect any personal data nor gather
information on pages where your computer sends personal data.
https://www.nic.cz/odvr/

sdns://AgcAAAAAAAAAE1syMDAxOjE0OGY6ZmZmZTo6MV2gzBBg05yDKbYrb7x9DW35MJhpuYHn5jktXNj6QI9NgOYgRE69Z7uD-IB7OSHpOKyReLiCvVCq2xEjHwRM9fCN984Lb2R2ci5uaWMuY3oEL2RvaA


## d0wn-is-ns2

Server provided by Martin 'd0wn' Albus

sdns://AQcAAAAAAAAADTkzLjk1LjIyNi4xNjUghGA0qcYwyjwErEqQFiXxeoeyrLlBgKxIHiwQ6M7eGm8cMi5kbnNjcnlwdC1jZXJ0LmlzMi5kMHduLmJpeg


## d0wn-tz-ns1

Server provided by Martin 'd0wn' Albus

sdns://AQcAAAAAAAAACzQxLjc5LjY5LjEzINYGFfvRRTuhTnaKPlxcs6wXRhMxRj2gr4z33wTaTXVtGzIuZG5zY3J5cHQtY2VydC50ei5kMHduLmJpeg


## d0wn-tz-ns1-ipv6

Server provided by Martin 'd0wn' Albus

sdns://AQcAAAAAAAAAGFsyYzBmOmZkYTg6NTo6MmVkMTpkMmVjXSDWBhX70UU7oU52ij5cXLOsF0YTMUY9oK-M998E2k11bRsyLmRuc2NyeXB0LWNlcnQudHouZDB3bi5iaXo


## dct-de1

DNSCrypt | IPv4 only | Non-logging | Non-filtering | DNSSEC | Nuremberg, Germany.

sdns://AQcAAAAAAAAACzE1OS42OS44Mi41IE4PHtUMeat3DKmfDJMZhuvOdHxR13BMHXf-4AHFnKesFzIuZG5zY3J5cHQtY2VydC5kY3QtZGUx


## dct-ru1

DNSCrypt | IPv4 only | Non-logging | Non-filtering | DNSSEC | Saint Petersburg, Russia.

sdns://AQcAAAAAAAAADDk0LjI0Mi41OS4zNSBODx7VDHmrdwypnwyTGYbrznR8UddwTB13_uABxZynrBcyLmRuc2NyeXB0LWNlcnQuZGN0LXJ1MQ


## dct-ru2

DNSCrypt | IPv4 only | Non-logging | Non-filtering | DNSSEC | Moscow, Russia.

sdns://AQcAAAAAAAAAEjE4NS4yMi4xNTQuMTk6ODQ0MyAfc_veSUgaoWE7_FCEKlLOuWJtg2DV8wX-XDfpB_D7fRcyLmRuc2NyeXB0LWNlcnQuZGN0LXJ1Mg


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

sdns://AgcAAAAAAAAADTE4NS45NS4yMTguNDKgzBBg05yDKbYrb7x9DW35MJhpuYHn5jktXNj6QI9NgOYgRE69Z7uD-IB7OSHpOKyReLiCvVCq2xEjHwRM9fCN984cZG5zLmRpZ2l0YWxlLWdlc2VsbHNjaGFmdC5jaAovZG5zLXF1ZXJ5


## dns.digitale-gesellschaft.ch-2

Public DoH resolver operated by the Digital Society (https://www.digitale-gesellschaft.ch).
Hosted in Zurich, Switzerland.

Non-logging, non-filtering, supports DNSSEC.

sdns://AgcAAAAAAAAADTE4NS45NS4yMTguNDOgzBBg05yDKbYrb7x9DW35MJhpuYHn5jktXNj6QI9NgOYgRE69Z7uD-IB7OSHpOKyReLiCvVCq2xEjHwRM9fCN984cZG5zLmRpZ2l0YWxlLWdlc2VsbHNjaGFmdC5jaAovZG5zLXF1ZXJ5


## dns.digitale-gesellschaft.ch-ipv6

Public IPv6 DoH resolver operated by the Digital Society (https://www.digitale-gesellschaft.ch).
Hosted in Zurich, Switzerland.

Non-logging, non-filtering, supports DNSSEC.

sdns://AgcAAAAAAAAAD1syYTA1OmZjODQ6OjQyXaDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5iBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33zhxkbnMuZGlnaXRhbGUtZ2VzZWxsc2NoYWZ0LmNoCi9kbnMtcXVlcnk


## dns.digitale-gesellschaft.ch-ipv6-2

Public IPv6 DoH resolver operated by the Digital Society (https://www.digitale-gesellschaft.ch).
Hosted in Zurich, Switzerland.

Non-logging, non-filtering, supports DNSSEC.

sdns://AgcAAAAAAAAAD1syYTA1OmZjODQ6OjQzXaDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5iBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33zhxkbnMuZGlnaXRhbGUtZ2VzZWxsc2NoYWZ0LmNoCi9kbnMtcXVlcnk


## dns.ryan-palmer

Non-logging, non-filtering, DNSSEC DoH Server. Hosted in the UK.

sdns://AgcAAAAAAAAADjY4LjE4My4yNTMuMjAwoMwQYNOcgym2K2-8fQ1t-TCYabmB5-Y5LVzY-kCPTYDmIEROvWe7g_iAezkh6TiskXi4gr1QqtsRIx8ETPXwjffOFGRuczEucnlhbi1wYWxtZXIuY29tCi9kbnMtcXVlcnk


## dns.sb

DNSSEC-enabled DoH server by https://xtom.com/
https://dns.sb/doh/

sdns://AgcAAAAAAAAADzE4NS4yMjIuMjIyLjIyMiACKC69_nqGQOLLyjCQJNWgwRXTVfrgMllhwr12J14JbAZkb2guc2IKL2Rucy1xdWVyeQ


## dns.sb-ipv6

DNSSEC-enabled DoH server by https://xtom.com/ - IPv6
https://dns.sb/doh/

sdns://AgcAAAAAAAAACFsyYTA5OjpdIAIoLr3-eoZA4svKMJAk1aDBFdNV-uAyWWHCvXYnXglsBmRvaC5zYgovZG5zLXF1ZXJ5


## dns.therifleman.name

DNS-over-HTTPS DNS forwarder from Mumbai, India. Blocks web and Android trackers and ads.
IP addresses are not logged, but queries are logged for 24 hours for debugging.
Report issues, send suggestions @ joker349 at protonmail.com.
Also supports DoT (for android) @ dns.therifleman.name and plain DNS @ 172.104.206.174

sdns://AgMAAAAAAAAADzE3Mi4xMDQuMjA2LjE3NKDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5iBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33zhRkbnMudGhlcmlmbGVtYW4ubmFtZQovZG5zLXF1ZXJ5


## dns.watch

Free, uncensored, non-logging server in Germany.
https://dns.watch

sdns://AQcAAAAAAAAADDg0LjIwMC43MC40MCBATVoA32LgIUTsC0hxmLlKGa_o5PRngHyldrEei5T5sCMyLmRuc2NyeXB0LWNlcnQucmVzb2x2ZXIyLmRucy53YXRjaA


## dns.watch-ipv6

Free, uncensored, non-logging server in Germany. IPv6 only.
https://dns.watch

sdns://AQcAAAAAAAAAHFsyMDAxOjE2MDg6MTA6MjU6OjkyNDk6ZDY5Yl0gQE1aAN9i4CFE7AtIcZi5Shmv6OT0Z4B8pXaxHouU-bAjMi5kbnNjcnlwdC1jZXJ0LnJlc29sdmVyMi5kbnMud2F0Y2g


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

Free, Canadian, uncensored, no-logs, encrypted, and DNSSEC validated DNS service for your pleasure. https://dnscrypt.ca/

sdns://AQcAAAAAAAAAEzE2Ny4xMTQuMjIwLjEyNTo0NDMgGlOjyVB4nL3RCxkzpGibbIRqQPG3PRdSrsrJgp7LfOIdMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeXB0LmNhLTE


## dnscrypt.ca-1-doh

Free, Canadian, uncensored, no-logs, encrypted, and DNSSEC validated DNS service for your pleasure. https://dnscrypt.ca/

sdns://AgcAAAAAAAAAEzE2Ny4xMTQuMjIwLjEyNTo0NTMgRE69Z7uD-IB7OSHpOKyReLiCvVCq2xEjHwRM9fCN984UZG5zMS5kbnNjcnlwdC5jYTo0NTMKL2Rucy1xdWVyeQ


## dnscrypt.ca-1-doh-ipv6

Free, Canadian, uncensored, no-logs, encrypted, and DNSSEC validated DNS service for your pleasure. https://dnscrypt.ca/

sdns://AgcAAAAAAAAAKVsyNjA3OjUzMDA6NjE6OTVmOjcyODM6MTFkOTpmODY6ZTY4OV06NDUzIEROvWe7g_iAezkh6TiskXi4gr1QqtsRIx8ETPXwjffOGWRuczEuaXB2Ni5kbnNjcnlwdC5jYTo0NTMKL2Rucy1xdWVyeQ


## dnscrypt.ca-1-ipv6

Free, Canadian, uncensored, no-logs, encrypted, and DNSSEC validated DNS service for your pleasure. https://dnscrypt.ca/

sdns://AQcAAAAAAAAAKVsyNjA3OjUzMDA6NjE6OTVmOjcyODM6MTFkOTpmODY6ZTY4OV06NDQzICDZGdTsPFAIzccKUeorT8v-ipk8ZDHQ7GgdXqc4RauAIjIuZG5zY3J5cHQtY2VydC5kbnNjcnlwdC5jYS0xLWlwdjY


## dnscrypt.ca-2

Free, Canadian, uncensored, no-logs, encrypted, and DNSSEC validated DNS service for your pleasure. https://dnscrypt.ca/

sdns://AQcAAAAAAAAAETE0OS41Ni4yMjguNDU6NDQzIAEIVKs7Vqfu-dORWP72ggv_k6I1fIkWCNueFdO74BGFHTIuZG5zY3J5cHQtY2VydC5kbnNjcnlwdC5jYS0y


## dnscrypt.ca-2-doh

Free, Canadian, uncensored, no-logs, encrypted, and DNSSEC validated DNS service for your pleasure. https://dnscrypt.ca/

sdns://AgcAAAAAAAAAETE0OS41Ni4yMjguNDU6NDUzIEROvWe7g_iAezkh6TiskXi4gr1QqtsRIx8ETPXwjffOFGRuczIuZG5zY3J5cHQuY2E6NDUzCi9kbnMtcXVlcnk


## dnscrypt.ca-2-doh-ipv6

Free, Canadian, uncensored, no-logs, encrypted, and DNSSEC validated DNS service for your pleasure. https://dnscrypt.ca/

sdns://AgcAAAAAAAAAKVsyNjA3OjUzMDA6NjE6OTVmOjcyODM6MTFkOTpmODY6ZTY5MF06NDUzIEROvWe7g_iAezkh6TiskXi4gr1QqtsRIx8ETPXwjffOGWRuczIuaXB2Ni5kbnNjcnlwdC5jYTo0NTMKL2Rucy1xdWVyeQ


## dnscrypt.ca-2-ipv6

Free, Canadian, uncensored, no-logs, encrypted, and DNSSEC validated DNS service for your pleasure. https://dnscrypt.ca/

sdns://AQcAAAAAAAAAKVsyNjA3OjUzMDA6NjE6OTVmOjcyODM6MTFkOTpmODY6ZTY5MF06NDQzIJh8cKiX29oVGjStS1XqgA71YcwEqtXFfdLD2GH8F_QpIjIuZG5zY3J5cHQtY2VydC5kbnNjcnlwdC5jYS0yLWlwdjY


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


## dnscrypt.uk-doh-ipv4

DoH, no logs, uncensored, DNSSEC. Hosted in London UK on Digital Ocean
https://www.dnscrypt.uk

sdns://AgcAAAAAAAAADzE2NS4yMjcuMjMzLjIwMAATZG9oLmRuc2NyeXB0LnVrOjQ0MwovZG5zLXF1ZXJ5


## dnscrypt.uk-doh-ipv6

DoH, no logs, uncensored, DNSSEC. Hosted in London UK on Digital Ocean
https://www.dnscrypt.uk

sdns://AgcAAAAAAAAAGlsyYTAzOmIwYzA6MTpkMDo6ZWJlOmQwMDFdABNkb2guZG5zY3J5cHQudWs6NDQzCi9kbnMtcXVlcnk


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

sdns://AgIAAAAAAAAADTk1LjIxNy4yMTMuOTSgzBBg05yDKbYrb7x9DW35MJhpuYHn5jktXNj6QI9NgOYgRE69Z7uD-IB7OSHpOKyReLiCvVCq2xEjHwRM9fCN984YZG5zLWRvaC5kbnNmb3JmYW1pbHkuY29tCi9kbnMtcXVlcnk


## dnsforfamily-doh-no-safe-search

(DoH Protocol) Block adult websites, gambling websites, malwares and advertisements.
Unlike other dnsforfamily DNSCrypt servers, this one does not enforces safe search. So Google, YouTube, Bing, DuckDuckGo and Yandex are completely accessible without any restriction.

Social websites like Facebook and Instagram are not blocked. No DNS queries are logged.

As of December 2020 2.7 million websites are blocked and new websites are added to blacklist daily.
Completely free, no ads or any commercial motive. Operating for 3 years now.

Warning: This server is incompatible with anonymization.

Provided by: https://dnsforfamily.com

sdns://AgIAAAAAAAAADTk1LjIxNy4yMTMuOTSgzBBg05yDKbYrb7x9DW35MJhpuYHn5jktXNj6QI9NgOYgRE69Z7uD-IB7OSHpOKyReLiCvVCq2xEjHwRM9fCN984nZG5zLWRvaC1uby1zYWZlLXNlYXJjaC5kbnNmb3JmYW1pbHkuY29tCi9kbnMtcXVlcnk


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

sdns://AgMAAAAAAAAADDE3Ni45LjkzLjE5OKDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5iBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33zgtkbnNmb3JnZS5kZQovZG5zLXF1ZXJ5


## dnshome-doh

https://www.dnshome.de/ public resolver in Germany

sdns://AgYAAAAAAAAAAKDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5iBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33zg5kbnMuZG5zaG9tZS5kZQovZG5zLXF1ZXJ5


## dnspod-doh

A public DNS resolver in mainland China provided by DNSPod (Tencent Cloud).

https://www.dnspod.cn/Products/Public.DNS?lang=en

sdns://AgAAAAAAAAAAACDBlAfoWsQD52fP2oOBh_Ag-lY6yBaIr1EMIqd559RaVgdkb2gucHViCi9kbnMtcXVlcnk


## dnswarden-asia-adblock-dcv4

Hosted in Singapore. For more information look [here](https://github.com/bhanupratapys/dnswarden) or [here](https://dnswarden.com).

sdns://AQMAAAAAAAAAFDIwNy4xNDguMTIwLjI0NDoyMDUzIKaotzxLPX3tRvpH4lDhbTo1cHCL4X-ZhW_Ji6jU8LwyITIuZG5zY3J5cHQtY2VydC5kbnN3YXJkZW4tYWRibG9jaw


## dnswarden-asia-adblock-dcv6

Hosted in Singapore. For more information look [here](https://github.com/bhanupratapys/dnswarden) or [here](https://dnswarden.com).

sdns://AQMAAAAAAAAALVsyMDAxOjE5ZjA6NDQwMDo0ZGUyOjU0MDA6M2ZmOmZlYTk6ZDQxOF06MjA1MyCmqLc8Sz197Ub6R-JQ4W06NXBwi-F_mYVvyYuo1PC8MiEyLmRuc2NyeXB0LWNlcnQuZG5zd2FyZGVuLWFkYmxvY2s


## dnswarden-asia-adblock-dohv4

Hosted in Singapore. For more information look [here](https://github.com/bhanupratapys/dnswarden) or [here](https://dnswarden.com).

sdns://AgMAAAAAAAAADzIwNy4xNDguMTIwLjI0NKDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5iBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33zhZkb2guYXNpYS5kbnN3YXJkZW4uY29tCC9hZGJsb2Nr


## dnswarden-asia-adblock-dohv6

Hosted in Singapore. For more information look [here](https://github.com/bhanupratapys/dnswarden) or [here](https://dnswarden.com).

sdns://AgMAAAAAAAAAKFsyMDAxOjE5ZjA6NDQwMDo0ZGUyOjU0MDA6M2ZmOmZlYTk6ZDQxOF2gzBBg05yDKbYrb7x9DW35MJhpuYHn5jktXNj6QI9NgOYgRE69Z7uD-IB7OSHpOKyReLiCvVCq2xEjHwRM9fCN984WZG9oLmFzaWEuZG5zd2FyZGVuLmNvbQgvYWRibG9jaw


## dnswarden-asia-adultfilter-dcv4

Hosted in Singapore. For more information look [here](https://github.com/bhanupratapys/dnswarden) or [here](https://dnswarden.com).

sdns://AQMAAAAAAAAAFDIwNy4xNDguMTIwLjI0NDo1MzUzIKaotzxLPX3tRvpH4lDhbTo1cHCL4X-ZhW_Ji6jU8LwyJTIuZG5zY3J5cHQtY2VydC5kbnN3YXJkZW4tYWR1bHRmaWx0ZXI


## dnswarden-asia-adultfilter-dcv6

Hosted in Singapore. For more information look [here](https://github.com/bhanupratapys/dnswarden) or [here](https://dnswarden.com).

sdns://AQMAAAAAAAAALVsyMDAxOjE5ZjA6NDQwMDo0ZGUyOjU0MDA6M2ZmOmZlYTk6ZDQxOF06NTM1MyCmqLc8Sz197Ub6R-JQ4W06NXBwi-F_mYVvyYuo1PC8MiUyLmRuc2NyeXB0LWNlcnQuZG5zd2FyZGVuLWFkdWx0ZmlsdGVy


## dnswarden-asia-adultfilter-dohv4

Hosted in Singapore. For more information look [here](https://github.com/bhanupratapys/dnswarden) or [here](https://dnswarden.com).

sdns://AgMAAAAAAAAADzIwNy4xNDguMTIwLjI0NKDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5iBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33zhZkb2guYXNpYS5kbnN3YXJkZW4uY29tDC9hZHVsdGZpbHRlcg


## dnswarden-asia-adultfilter-dohv6

Hosted in Singapore. For more information look [here](https://github.com/bhanupratapys/dnswarden) or [here](https://dnswarden.com).

sdns://AgMAAAAAAAAAKFsyMDAxOjE5ZjA6NDQwMDo0ZGUyOjU0MDA6M2ZmOmZlYTk6ZDQxOF2gzBBg05yDKbYrb7x9DW35MJhpuYHn5jktXNj6QI9NgOYgRE69Z7uD-IB7OSHpOKyReLiCvVCq2xEjHwRM9fCN984WZG9oLmFzaWEuZG5zd2FyZGVuLmNvbQwvYWR1bHRmaWx0ZXI


## dnswarden-asia-uncensor-dcv4

Hosted in Singapore. For more information look [here](https://github.com/bhanupratapys/dnswarden) or [here](https://dnswarden.com).

sdns://AQcAAAAAAAAAFTIwNy4xNDguMTIwLjI0NDoxNTM1MyCmqLc8Sz197Ub6R-JQ4W06NXBwi-F_mYVvyYuo1PC8MiIyLmRuc2NyeXB0LWNlcnQuZG5zd2FyZGVuLXVuY2Vuc29y


## dnswarden-asia-uncensor-dcv6

Hosted in Singapore. For more information look [here](https://github.com/bhanupratapys/dnswarden) or [here](https://dnswarden.com).

sdns://AQcAAAAAAAAALlsyMDAxOjE5ZjA6NDQwMDo0ZGUyOjU0MDA6M2ZmOmZlYTk6ZDQxOF06MTUzNTMgpqi3PEs9fe1G-kfiUOFtOjVwcIvhf5mFb8mLqNTwvDIiMi5kbnNjcnlwdC1jZXJ0LmRuc3dhcmRlbi11bmNlbnNvcg


## dnswarden-asia-uncensor-dohv4

Hosted in Singapore. For more information look [here](https://github.com/bhanupratapys/dnswarden) or [here](https://dnswarden.com).

sdns://AgcAAAAAAAAADzIwNy4xNDguMTIwLjI0NKDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5iBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33zhZkb2guYXNpYS5kbnN3YXJkZW4uY29tCy91bmNlbnNvcmVk


## dnswarden-asia-uncensor-dohv6

Hosted in Singapore. For more information look [here](https://github.com/bhanupratapys/dnswarden) or [here](https://dnswarden.com).

sdns://AgcAAAAAAAAAKFsyMDAxOjE5ZjA6NDQwMDo0ZGUyOjU0MDA6M2ZmOmZlYTk6ZDQxOF2gzBBg05yDKbYrb7x9DW35MJhpuYHn5jktXNj6QI9NgOYgRE69Z7uD-IB7OSHpOKyReLiCvVCq2xEjHwRM9fCN984WZG9oLmFzaWEuZG5zd2FyZGVuLmNvbQsvdW5jZW5zb3JlZA


## dnswarden-eu-adblock-dcv4

Hosted in Germany. For more information look [here](https://github.com/bhanupratapys/dnswarden) or [here](https://dnswarden.com).

sdns://AQMAAAAAAAAAEDQ1Ljc2Ljg4LjIwOjIwNTMg1s4Ar-3awN0OmfAQzlZywvR2ZGVMg0aeNdwrLHspdpMhMi5kbnNjcnlwdC1jZXJ0LmRuc3dhcmRlbi1hZGJsb2Nr


## dnswarden-eu-adblock-dcv6

Hosted in Germany. For more information look [here](https://github.com/bhanupratapys/dnswarden) or [here](https://dnswarden.com).

sdns://AQMAAAAAAAAALVsyMDAxOjE5ZjA6NmMwMToyYjE3OjU0MDA6M2ZmOmZlYTk6ZDQyZl06MjA1MyDWzgCv7drA3Q6Z8BDOVnLC9HZkZUyDRp413Csseyl2kyEyLmRuc2NyeXB0LWNlcnQuZG5zd2FyZGVuLWFkYmxvY2s


## dnswarden-eu-adblock-dohv4

Hosted in Germany. For more information look [here](https://github.com/bhanupratapys/dnswarden) or [here](https://dnswarden.com).

sdns://AgMAAAAAAAAACzQ1Ljc2Ljg4LjIwoMwQYNOcgym2K2-8fQ1t-TCYabmB5-Y5LVzY-kCPTYDmIEROvWe7g_iAezkh6TiskXi4gr1QqtsRIx8ETPXwjffOFGRvaC5ldS5kbnN3YXJkZW4uY29tCC9hZGJsb2Nr


## dnswarden-eu-adblock-dohv6

Hosted in Germany. For more information look [here](https://github.com/bhanupratapys/dnswarden) or [here](https://dnswarden.com).

sdns://AgMAAAAAAAAAKFsyMDAxOjE5ZjA6NmMwMToyYjE3OjU0MDA6M2ZmOmZlYTk6ZDQyZl2gzBBg05yDKbYrb7x9DW35MJhpuYHn5jktXNj6QI9NgOYgRE69Z7uD-IB7OSHpOKyReLiCvVCq2xEjHwRM9fCN984UZG9oLmV1LmRuc3dhcmRlbi5jb20IL2FkYmxvY2s


## dnswarden-eu-adultfilter-dcv4

Hosted in Germany. For more information look [here](https://github.com/bhanupratapys/dnswarden) or [here](https://dnswarden.com).

sdns://AQMAAAAAAAAAEDQ1Ljc2Ljg4LjIwOjUzNTMg1s4Ar-3awN0OmfAQzlZywvR2ZGVMg0aeNdwrLHspdpMlMi5kbnNjcnlwdC1jZXJ0LmRuc3dhcmRlbi1hZHVsdGZpbHRlcg


## dnswarden-eu-adultfilter-dcv6

Hosted in Germany. For more information look [here](https://github.com/bhanupratapys/dnswarden) or [here](https://dnswarden.com).

sdns://AQMAAAAAAAAALVsyMDAxOjE5ZjA6NmMwMToyYjE3OjU0MDA6M2ZmOmZlYTk6ZDQyZl06NTM1MyDWzgCv7drA3Q6Z8BDOVnLC9HZkZUyDRp413Csseyl2kyUyLmRuc2NyeXB0LWNlcnQuZG5zd2FyZGVuLWFkdWx0ZmlsdGVy


## dnswarden-eu-adultfilter-dohv4

Hosted in Germany. For more information look [here](https://github.com/bhanupratapys/dnswarden) or [here](https://dnswarden.com).

sdns://AgMAAAAAAAAACzQ1Ljc2Ljg4LjIwoMwQYNOcgym2K2-8fQ1t-TCYabmB5-Y5LVzY-kCPTYDmIEROvWe7g_iAezkh6TiskXi4gr1QqtsRIx8ETPXwjffOFGRvaC5ldS5kbnN3YXJkZW4uY29tDC9hZHVsdGZpbHRlcg


## dnswarden-eu-adultfilter-dohv6

Hosted in Germany. For more information look [here](https://github.com/bhanupratapys/dnswarden) or [here](https://dnswarden.com).

sdns://AgMAAAAAAAAAKFsyMDAxOjE5ZjA6NmMwMToyYjE3OjU0MDA6M2ZmOmZlYTk6ZDQyZl2gzBBg05yDKbYrb7x9DW35MJhpuYHn5jktXNj6QI9NgOYgRE69Z7uD-IB7OSHpOKyReLiCvVCq2xEjHwRM9fCN984UZG9oLmV1LmRuc3dhcmRlbi5jb20ML2FkdWx0ZmlsdGVy


## dnswarden-eu-uncensor-dcv4

Hosted in Germany. For more information look [here](https://github.com/bhanupratapys/dnswarden) or [here](https://dnswarden.com).

sdns://AQcAAAAAAAAAETQ1Ljc2Ljg4LjIwOjE1MzUzINbOAK_t2sDdDpnwEM5WcsL0dmRlTINGnjXcKyx7KXaTIjIuZG5zY3J5cHQtY2VydC5kbnN3YXJkZW4tdW5jZW5zb3I


## dnswarden-eu-uncensor-dcv6

Hosted in Germany. For more information look [here](https://github.com/bhanupratapys/dnswarden) or [here](https://dnswarden.com).

sdns://AQcAAAAAAAAALlsyMDAxOjE5ZjA6NmMwMToyYjE3OjU0MDA6M2ZmOmZlYTk6ZDQyZl06MTUzNTMg1s4Ar-3awN0OmfAQzlZywvR2ZGVMg0aeNdwrLHspdpMiMi5kbnNjcnlwdC1jZXJ0LmRuc3dhcmRlbi11bmNlbnNvcg


## dnswarden-eu-uncensor-dohv4

Hosted in Germany. For more information look [here](https://github.com/bhanupratapys/dnswarden) or [here](https://dnswarden.com).

sdns://AgcAAAAAAAAACzQ1Ljc2Ljg4LjIwoMwQYNOcgym2K2-8fQ1t-TCYabmB5-Y5LVzY-kCPTYDmIEROvWe7g_iAezkh6TiskXi4gr1QqtsRIx8ETPXwjffOFGRvaC5ldS5kbnN3YXJkZW4uY29tCy91bmNlbnNvcmVk


## dnswarden-eu-uncensor-dohv6

Hosted in Germany. For more information look [here](https://github.com/bhanupratapys/dnswarden) or [here](https://dnswarden.com).

sdns://AgcAAAAAAAAAKFsyMDAxOjE5ZjA6NmMwMToyYjE3OjU0MDA6M2ZmOmZlYTk6ZDQyZl2gzBBg05yDKbYrb7x9DW35MJhpuYHn5jktXNj6QI9NgOYgRE69Z7uD-IB7OSHpOKyReLiCvVCq2xEjHwRM9fCN984UZG9oLmV1LmRuc3dhcmRlbi5jb20LL3VuY2Vuc29yZWQ


## dnswarden-us-adblock-dcv4

Hosted in USA (Dallas) . For more information look [here](https://github.com/bhanupratapys/dnswarden) or [here](https://dnswarden.com).

sdns://AQMAAAAAAAAAEzE0NC4yMDIuNjkuMTQ5OjIwNTMgc_eiadnb3wpUqznTHbI2NziLpELYc4uuDM-s3v3CsSQhMi5kbnNjcnlwdC1jZXJ0LmRuc3dhcmRlbi1hZGJsb2Nr


## dnswarden-us-adblock-dcv6

Hosted in USA (Dallas) . For more information look [here](https://github.com/bhanupratapys/dnswarden) or [here](https://dnswarden.com).

sdns://AQMAAAAAAAAALFsyMDAxOjE5ZjA6NjQwMTpkNmU6NTQwMDozZmY6ZmVhOTpkNDVlXToyMDUzIHP3omnZ298KVKs50x2yNjc4i6RC2HOLrgzPrN79wrEkITIuZG5zY3J5cHQtY2VydC5kbnN3YXJkZW4tYWRibG9jaw


## dnswarden-us-adblock-dohv4

Hosted in USA (Dallas) . For more information look [here](https://github.com/bhanupratapys/dnswarden) or [here](https://dnswarden.com).

sdns://AgMAAAAAAAAADjE0NC4yMDIuNjkuMTQ5oMwQYNOcgym2K2-8fQ1t-TCYabmB5-Y5LVzY-kCPTYDmIEROvWe7g_iAezkh6TiskXi4gr1QqtsRIx8ETPXwjffOFGRvaC51cy5kbnN3YXJkZW4uY29tCC9hZGJsb2Nr


## dnswarden-us-adblock-dohv6

Hosted in USA (Dallas) . For more information look [here](https://github.com/bhanupratapys/dnswarden) or [here](https://dnswarden.com).

sdns://AgMAAAAAAAAAJ1syMDAxOjE5ZjA6NjQwMTpkNmU6NTQwMDozZmY6ZmVhOTpkNDVlXaDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5iBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33zhRkb2gudXMuZG5zd2FyZGVuLmNvbQgvYWRibG9jaw


## dnswarden-us-adultfilter-dcv4

Hosted in USA (Dallas) . For more information look [here](https://github.com/bhanupratapys/dnswarden) or [here](https://dnswarden.com).

sdns://AQMAAAAAAAAAEzE0NC4yMDIuNjkuMTQ5OjUzNTMgc_eiadnb3wpUqznTHbI2NziLpELYc4uuDM-s3v3CsSQlMi5kbnNjcnlwdC1jZXJ0LmRuc3dhcmRlbi1hZHVsdGZpbHRlcg


## dnswarden-us-adultfilter-dcv6

Hosted in USA (Dallas) . For more information look [here](https://github.com/bhanupratapys/dnswarden) or [here](https://dnswarden.com).

sdns://AQMAAAAAAAAALFsyMDAxOjE5ZjA6NjQwMTpkNmU6NTQwMDozZmY6ZmVhOTpkNDVlXTo1MzUzIHP3omnZ298KVKs50x2yNjc4i6RC2HOLrgzPrN79wrEkJTIuZG5zY3J5cHQtY2VydC5kbnN3YXJkZW4tYWR1bHRmaWx0ZXI


## dnswarden-us-adultfilter-dohv4

Hosted in USA (Dallas) . For more information look [here](https://github.com/bhanupratapys/dnswarden) or [here](https://dnswarden.com).

sdns://AgMAAAAAAAAADjE0NC4yMDIuNjkuMTQ5oMwQYNOcgym2K2-8fQ1t-TCYabmB5-Y5LVzY-kCPTYDmIEROvWe7g_iAezkh6TiskXi4gr1QqtsRIx8ETPXwjffOFGRvaC51cy5kbnN3YXJkZW4uY29tDC9hZHVsdGZpbHRlcg


## dnswarden-us-adultfilter-dohv6

Hosted in USA (Dallas) . For more information look [here](https://github.com/bhanupratapys/dnswarden) or [here](https://dnswarden.com).

sdns://AgMAAAAAAAAAJ1syMDAxOjE5ZjA6NjQwMTpkNmU6NTQwMDozZmY6ZmVhOTpkNDVlXaDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5iBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33zhRkb2gudXMuZG5zd2FyZGVuLmNvbQwvYWR1bHRmaWx0ZXI


## dnswarden-us-uncensor-dcv4

Hosted in USA (Dallas) . For more information look [here](https://github.com/bhanupratapys/dnswarden) or [here](https://dnswarden.com).

sdns://AQcAAAAAAAAAFDE0NC4yMDIuNjkuMTQ5OjE1MzUzIHP3omnZ298KVKs50x2yNjc4i6RC2HOLrgzPrN79wrEkIjIuZG5zY3J5cHQtY2VydC5kbnN3YXJkZW4tdW5jZW5zb3I


## dnswarden-us-uncensor-dcv6

Hosted in USA (Dallas) . For more information look [here](https://github.com/bhanupratapys/dnswarden) or [here](https://dnswarden.com).

sdns://AQcAAAAAAAAALVsyMDAxOjE5ZjA6NjQwMTpkNmU6NTQwMDozZmY6ZmVhOTpkNDVlXToxNTM1MyBz96Jp2dvfClSrOdMdsjY3OIukQthzi64Mz6ze_cKxJCIyLmRuc2NyeXB0LWNlcnQuZG5zd2FyZGVuLXVuY2Vuc29y


## dnswarden-us-uncensor-dohv4

Hosted in USA (Dallas) . For more information look [here](https://github.com/bhanupratapys/dnswarden) or [here](https://dnswarden.com).

sdns://AgcAAAAAAAAADjE0NC4yMDIuNjkuMTQ5oMwQYNOcgym2K2-8fQ1t-TCYabmB5-Y5LVzY-kCPTYDmIEROvWe7g_iAezkh6TiskXi4gr1QqtsRIx8ETPXwjffOFGRvaC51cy5kbnN3YXJkZW4uY29tCy91bmNlbnNvcmVk


## dnswarden-us-uncensor-dohv6

Hosted in USA (Dallas) . For more information look [here](https://github.com/bhanupratapys/dnswarden) or [here](https://dnswarden.com).

sdns://AgcAAAAAAAAAJ1syMDAxOjE5ZjA6NjQwMTpkNmU6NTQwMDozZmY6ZmVhOTpkNDVlXaDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5iBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33zhRkb2gudXMuZG5zd2FyZGVuLmNvbQsvdW5jZW5zb3JlZA


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


## doh-jp-blahdns

Blocks ad and Tracking, no Logging, DNSSEC, Hosted in Japan. By https://blahdns.com/

sdns://AgMAAAAAAAAADjEzOS4xNjIuMTEyLjQ3ABJkb2gtanAuYmxhaGRucy5jb20KL2Rucy1xdWVyeQ


## doh-jp-blahdns-v6

Blocks ad and Tracking, no Logging, DNSSEC, Hosted in Japan. By https://blahdns.com/

sdns://AgMAAAAAAAAAIFsyNDAwOjg5MDI6OmYwM2M6OTJmZjpmZTI3OjM0NGJdABJkb2gtanAuYmxhaGRucy5jb20KL2Rucy1xdWVyeQ


## doh.appliedprivacy.net

Public DoH resolver operated by the Foundation for Applied Privacy (https://appliedprivacy.net).
Hosted in Vienna, Austria.

Non-logging, non-filtering, supports DNSSEC.

sdns://AgcAAAAAAAAAAKDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5iBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33zhZkb2guYXBwbGllZHByaXZhY3kubmV0Bi9xdWVyeQ


## doh.ffmuc.net

An open (non-logging, non-filtering, non-censoring) DoH resolver operated by Freifunk Munich with nodes in DE.
https://ffmuc.net/

sdns://AgcAAAAAAAAACjUuMS42Ni4yNTWgzBBg05yDKbYrb7x9DW35MJhpuYHn5jktXNj6QI9NgOYgRE69Z7uD-IB7OSHpOKyReLiCvVCq2xEjHwRM9fCN984NZG9oLmZmbXVjLm5ldAovZG5zLXF1ZXJ5


## doh.ffmuc.net-2

An open (non-logging, non-filtering, non-censoring) DoH resolver operated by Freifunk Munich with nodes in DE.
https://ffmuc.net/

sdns://AgcAAAAAAAAADjE4NS4xNTAuOTkuMjU1oMwQYNOcgym2K2-8fQ1t-TCYabmB5-Y5LVzY-kCPTYDmIEROvWe7g_iAezkh6TiskXi4gr1QqtsRIx8ETPXwjffODWRvaC5mZm11Yy5uZXQKL2Rucy1xdWVyeQ


## doh.ffmuc.net-v6

An open (non-logging, non-filtering, non-censoring) DoH resolver operated by Freifunk Munich with nodes in DE.
https://ffmuc.net/

sdns://AgcAAAAAAAAAFVsyMDAxOjY3ODplNjg6ZjAwMDo6XaDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5iBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33zg1kb2guZmZtdWMubmV0Ci9kbnMtcXVlcnk


## doh.ffmuc.net-v6-2

An open (non-logging, non-filtering, non-censoring) DoH resolver operated by Freifunk Munich with nodes in DE.
https://ffmuc.net/

sdns://AgcAAAAAAAAAFVsyMDAxOjY3ODplZDA6ZjAwMDo6XaDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5iBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33zg1kb2guZmZtdWMubmV0Ci9kbnMtcXVlcnk


## doh.tiarap.org

Non-Logging DNS-over-HTTPS server, cached via Cloudflare.
Filters out ads, trackers and malware, NO ECS, supports DNSSEC.

sdns://AgMAAAAAAAAADDEwNC4yMS42NS42MAAOZG9oLnRpYXJhcC5vcmcKL2Rucy1xdWVyeQ


## doh.tiarap.org-ipv6

Non-Logging DNS-over-HTTPS server (IPv6), cached via Cloudflare.
Filters out ads, trackers and malware, NO ECS, supports DNSSEC.

sdns://AgMAAAAAAAAAG1syNjA2OjQ3MDA6MzAzNDo6NjgxNTo0MTNjXQAOZG9oLnRpYXJhcC5vcmcKL2Rucy1xdWVyeQ


## e-utp.net-ipv6-doh

e-utp.net DNS Cache service, only accessible over IPv6
https://fido.e-utp.net/display/EUTPNET/Recursive+DNS

sdns://AgcAAAAAAAAAE1syMDAxOjY3YzoyMWVjOjo1M10gXMJdnIoy7PLLu5Hr9wvZVsUgOfnyDoQJ2mFUcVI21x0SZG5zY2FjaGUuZS11dHAubmV0Ci9kbnMtcXVlcnk


## east.comss.one

DNS with adblock filters and antiphishing, gaining popularity among russian-speaking users.

sdns://AQMAAAAAAAAAEjkxLjIzMC4yMTEuNjc6NTQ0MyAVacgDhm0kq-ciz2DrJlPhYjI9v6tAQHbulojEz86TEyIyLmRuc2NyeXB0LWNlcnQuZG5zLmVhc3QuY29tc3Mub25l


## faelix-ch-ipv4

An open (non-logging, non-filtering, no ECS) DNSCrypt resolver operated by https://faelix.net/ with IPv4 nodes anycast within AS41495 in Switzerland.

sdns://AQYAAAAAAAAAEzE4NS4xMzQuMTk2LjU0Ojg0NDMgfsvvPi8BgDKNYODh0ewj5Oh32OoJoZNwGgTWs8C-i-EfMi5kbnNjcnlwdC1jZXJ0LnJkbnMuZmFlbGl4Lm5ldA
sdns://AQYAAAAAAAAAEzE4NS4xMzQuMTk2LjU1Ojg0NDMgfsvvPi8BgDKNYODh0ewj5Oh32OoJoZNwGgTWs8C-i-EfMi5kbnNjcnlwdC1jZXJ0LnJkbnMuZmFlbGl4Lm5ldA


## faelix-uk-ipv4

An open (non-logging, non-filtering, no ECS) DNSCrypt resolver operated by https://faelix.net/ with IPv4 nodes anycast within AS41495 in the UK.

sdns://AQYAAAAAAAAAEjQ2LjIyNy4yMDAuNTQ6ODQ0MyB-y-8-LwGAMo1g4OHR7CPk6HfY6gmhk3AaBNazwL6L4R8yLmRuc2NyeXB0LWNlcnQucmRucy5mYWVsaXgubmV0
sdns://AQYAAAAAAAAAEjQ2LjIyNy4yMDAuNTU6ODQ0MyB-y-8-LwGAMo1g4OHR7CPk6HfY6gmhk3AaBNazwL6L4R8yLmRuc2NyeXB0LWNlcnQucmRucy5mYWVsaXgubmV0


## faelix-uk-ipv6

An open (non-logging, non-filtering, no ECS) DNSCrypt resolver operated by https://faelix.net/ with IPv6 nodes anycast within AS41495 in the UK.

sdns://AQYAAAAAAAAAFFsyYTAxOjllMDA6OjU0XTo4NDQzIH7L7z4vAYAyjWDg4dHsI-Tod9jqCaGTcBoE1rPAvovhHzIuZG5zY3J5cHQtY2VydC5yZG5zLmZhZWxpeC5uZXQ
sdns://AQYAAAAAAAAAFFsyYTAxOjllMDA6OjU1XTo4NDQzIH7L7z4vAYAyjWDg4dHsI-Tod9jqCaGTcBoE1rPAvovhHzIuZG5zY3J5cHQtY2VydC5yZG5zLmZhZWxpeC5uZXQ


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


## gombadi-syd

Protocol: DNSCrypt IPv4 | Features: Non-logging, Non-filtering, DNSSEC, Unbound | Location: Sydney, AU

sdns://AQUAAAAAAAAADzE3Mi4xMDUuMTY2LjE4OSCqZ1fiHV9W_d5oL0UH7Cc1fIohJG9lej3-8EW2ND_xbBwyLmRuc2NyeXB0LWNlcnQuZ29tYmFkaS5zeWQy


## gombadi-syd-ipv6

Protocol: DNSCrypt IPv6 | Features: Non-logging, Non-filtering, DNSSEC, Unbound | Location: Sydney, AU

sdns://AQUAAAAAAAAAIFsyNDAwOjg5MDc6OmYwM2M6OTJmZjpmZTYzOmU3NDhdIKpnV-IdX1b93mgvRQfsJzV8iiEkb2V6Pf7wRbY0P_FsHDIuZG5zY3J5cHQtY2VydC5nb21iYWRpLnN5ZDI


## google

Google DNS (anycast)

sdns://AgUAAAAAAAAABzguOC44LjigHvYkz_9ea9O63fP92_3qVlRn43cpncfuZnUWbzAMwbmgdoAkR6AZkxo_AEMExT_cbBssN43Evo9zs5_ZyWnftEUgalBisNF41VbxY7E7Gw8ZQ10CWIKRzHVYnf7m6xHI1cMKZG5zLmdvb2dsZQovZG5zLXF1ZXJ5


## google-ipv6

Google DNS (anycast)

sdns://AgUAAAAAAAAAFlsyMDAxOjQ4NjA6NDg2MDo6ODg4OF2gHvYkz_9ea9O63fP92_3qVlRn43cpncfuZnUWbzAMwbmgdoAkR6AZkxo_AEMExT_cbBssN43Evo9zs5_ZyWnftEUgalBisNF41VbxY7E7Gw8ZQ10CWIKRzHVYnf7m6xHI1cMKZG5zLmdvb2dsZQovZG5zLXF1ZXJ5


## hdns

HDNS is a public DNS resolver that supports Handshake domains.
https://www.hdns.io

sdns://AgYAAAAAAAAADjEwMy4xOTYuMzguMjAwIBGxAuax9j5SiYTWAl8ysTgkH8iLvXUZV01wyYMtU-HoDXF1ZXJ5LmhkbnMuaW8KL2Rucy1xdWVyeQ


## he

Hurricane Electric DoH server (anycast)

Unknown logging policy.

sdns://AgUAAAAAAAAACzc0LjgyLjQyLjQyoMwQYNOcgym2K2-8fQ1t-TCYabmB5-Y5LVzY-kCPTYDmIEROvWe7g_iAezkh6TiskXi4gr1QqtsRIx8ETPXwjffODG9yZG5zLmhlLm5ldAovZG5zLXF1ZXJ5


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

sdns://AgMAAAAAAAAADjE3NC4xMzguMjkuMTc1oMwQYNOcgym2K2-8fQ1t-TCYabmB5-Y5LVzY-kCPTYDmIEROvWe7g_iAezkh6TiskXi4gr1QqtsRIx8ETPXwjffODGRvaC50aWFyLmFwcAovZG5zLXF1ZXJ5


## id-gmail-doh-ipv6

Non-Logging DNS-over-HTTPS (IPv6) server located in Singapore.
Filters out ads, trackers and malware, supports DNSSEC, provided by id-gmail.

sdns://AgMAAAAAAAAAG1syNDAwOjYxODA6MDpkMDo6NWY3Mzo0MDAxXaDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5iBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33zgxkb2gudGlhci5hcHAKL2Rucy1xdWVyeQ


## id-gmail-ipv6

Non-Logging DNSCrypt (IPv6) server located in Singapore.
Filters out ads, trackers and malware, supports DNSSEC, provided by id-gmail.

sdns://AQMAAAAAAAAAG1syNDAwOjYxODA6MDpkMDo6NWY2ZTo0MDAxXSDvloBm6NmU8GXYPt3TGu7t9W2GV0Y-SspHzdf-l8PU9hwyLmRuc2NyeXB0LWNlcnQuZG5zLnRpYXIuYXBw


## iij

DoH server operated by Internet Initiative Japan in Tokyo.
https://www.iij.ad.jp/

sdns://AgUAAAAAAAAACjEwMy4yLjU3LjUgmOPV5TavKVjNL38U9wTvSidtJeM81l8uZfXk8nJ8EzARcHVibGljLmRucy5paWouanAKL2Rucy1xdWVyeQ


## jp.tiar.app

Non-Logging, Non-Filtering DNSCrypt server in Japan.
No ECS, Support DNSSEC

sdns://AQcAAAAAAAAAEjE3Mi4xMDQuOTMuODA6MTQ0MyAyuHY-8b9lNqHeahPAzW9IoXnjiLaZpTeNbVs8TN9UUxsyLmRuc2NyeXB0LWNlcnQuanAudGlhci5hcHA


## jp.tiar.app-doh

Non-Logging, Non-Filtering DNS-over-HTTPS server in Japan.
No ECS, Support DNSSEC

sdns://AgcAAAAAAAAADTE3Mi4xMDQuOTMuODCgzBBg05yDKbYrb7x9DW35MJhpuYHn5jktXNj6QI9NgOYgRE69Z7uD-IB7OSHpOKyReLiCvVCq2xEjHwRM9fCN984LanAudGlhci5hcHAKL2Rucy1xdWVyeQ


## jp.tiar.app-doh-ipv6

Non-Logging, Non-Filtering DNS-over-HTTPS (IPv6) server in Japan.
No ECS, Support DNSSEC

sdns://AgcAAAAAAAAAIFsyNDAwOjg5MDI6OmYwM2M6OTFmZjpmZWRhOmM1MTRdoMwQYNOcgym2K2-8fQ1t-TCYabmB5-Y5LVzY-kCPTYDmIEROvWe7g_iAezkh6TiskXi4gr1QqtsRIx8ETPXwjffOC2pwLnRpYXIuYXBwCi9kbnMtcXVlcnk


## jp.tiar.app-ipv6

Non-Logging, Non-Filtering DNSCrypt (IPv6) server in Japan.
No ECS, Support DNSSEC

sdns://AQcAAAAAAAAAJVsyNDAwOjg5MDI6OmYwM2M6OTFmZjpmZWRhOmM1MTRdOjE0NDMgMrh2PvG_ZTah3moTwM1vSKF544i2maU3jW1bPEzfVFMbMi5kbnNjcnlwdC1jZXJ0LmpwLnRpYXIuYXBw


## jp.tiarap.org

DNS-over-HTTPS Server. Non-Logging, Non-Filtering, No ECS, Support DNSSEC.
Cached via Cloudflare.

sdns://AgcAAAAAAAAAAAANanAudGlhcmFwLm9yZwovZG5zLXF1ZXJ5


## jp.tiarap.org-ipv6

DNS-over-HTTPS Server (IPv6). Non-Logging, Non-Filtering, No ECS, Support DNSSEC.
Cached via Cloudflare.

sdns://AgcAAAAAAAAAG1syNjA2OjQ3MDA6MzAzNjo6NjgxYjo5NmFhXQANanAudGlhcmFwLm9yZwovZG5zLXF1ZXJ5


## libredns

DoH server in Germany. No logging, but no DNS padding and no DNSSEC support.
https://libredns.gr/

sdns://AgYAAAAAAAAADjExNi4yMDIuMTc2LjI2oMwQYNOcgym2K2-8fQ1t-TCYabmB5-Y5LVzY-kCPTYDmIEROvWe7g_iAezkh6TiskXi4gr1QqtsRIx8ETPXwjffOD2RvaC5saWJyZWRucy5ncgovZG5zLXF1ZXJ5


## libredns-noads

DoH server in Germany. No logging, but no DNS padding and no DNSSEC support.
no ads version, uses StevenBlack's host list: https://github.com/StevenBlack/hosts

sdns://AgIAAAAAAAAADjExNi4yMDIuMTc2LjI2oMwQYNOcgym2K2-8fQ1t-TCYabmB5-Y5LVzY-kCPTYDmIEROvWe7g_iAezkh6TiskXi4gr1QqtsRIx8ETPXwjffOD2RvaC5saWJyZWRucy5ncgQvYWRz


## meganerd

DNSCrypt server by MegaNerd.nl (IPv4) - https://meganerd.nl/encrypted-dns-server
Hosted in Amsterdam (AMS1), The Netherlands.

Non-logging, non-filtering, supports DNSSEC.

sdns://AQcAAAAAAAAADjEzNi4yNDQuOTcuMTE0IPyq3HBOXuNgu6FO4pU71Si6CTV6kPD85NA6AThr_6tiGDIuZG5zY3J5cHQtY2VydC5tZWdhbmVyZA


## meganerd-doh-ipv4

DoH server by MegaNerd.nl (IPv4) - https://meganerd.nl/encrypted-dns-server
Hosted in Amsterdam (AMS1), The Netherlands.

Non-logging, non-filtering, supports DNSSEC.

sdns://AgcAAAAAAAAADjEzNi4yNDQuOTcuMTE0oMwQYNOcgym2K2-8fQ1t-TCYabmB5-Y5LVzY-kCPTYDmIEROvWe7g_iAezkh6TiskXi4gr1QqtsRIx8ETPXwjffOGWNoZXdiYWNjYS5tZWdhbmVyZC5ubDo0NDMEL2RvaA


## meganerd-doh-ipv6

DoH server by MegaNerd.nl (IPv6) - https://meganerd.nl/encrypted-dns-server
Hosted in Amsterdam (AMS1), The Netherlands.

Non-logging, non-filtering, supports DNSSEC.

sdns://AgcAAAAAAAAAKFsyMDAxOjE5ZjA6NTAwMTpjYmI6NTQwMDowM2ZmOmZlMDc6ZjcwZF2gzBBg05yDKbYrb7x9DW35MJhpuYHn5jktXNj6QI9NgOYgRE69Z7uD-IB7OSHpOKyReLiCvVCq2xEjHwRM9fCN984VY2hld2JhY2NhLm1lZ2FuZXJkLm5sBC9kb2g


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

sdns://AgcAAAAAAAAACjQ1LjkwLjI4LjAgmjo09yfeubylEAPZzpw5-PJ92cUkKQHCurGkTmNaAhkOZG5zLm5leHRkbnMuaW8PL2Ruc2NyeXB0LXByb3h5


## nextdns-ipv6

Connects to NextDNS over IPv6.

NextDNS is a cloud-based private DNS service that gives you full control
over what is allowed and what is blocked on the Internet.

DNSSEC, Anycast, Non-logging, NoFilters

https://www.nextdns.io/

sdns://AgcAAAAAAAAADVsyYTA3OmE4YzA6Ol0gmjo09yfeubylEAPZzpw5-PJ92cUkKQHCurGkTmNaAhkOZG5zLm5leHRkbnMuaW8PL2Ruc2NyeXB0LXByb3h5


## njalla-doh

Non-logging DoH server in Sweden operated by Njalla.

https://dns.njal.la/

sdns://AgcAAAAAAAAADDk1LjIxNS4xOS41M6DMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5iBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33zgtkbnMubmphbC5sYQovZG5zLXF1ZXJ5


## opennic-R4SAS

DNSSEC • OpenNIC • Non-logging • Uncensored - hosted on OVH - https://opennic.i2pd.xyz/
Location: Paris, France
Maintained by R4SAS - https://github.com/r4sas

sdns://AQcAAAAAAAAADTE1MS44MC4yMjIuNzkgqdYyOk8lgAkmGXUVAs4jHh922d53bIfGu7KKDv_bDk4gMi5kbnNjcnlwdC1jZXJ0Lm9wZW5uaWMuaTJwZC54eXo


## opennic-R4SAS-doh-ipv6

DOH • DNSSEC • OpenNIC • Non-logging • Uncensored - hosted on OVH - https://opennic.i2pd.xyz/
Location: Paris, France
Maintained by R4SAS - https://github.com/r4sas

sdns://AgYAAAAAAAAAF1syMDAxOjQ3MDoxZjE1OmI4MDo6NTNdICP03s8txO4Hwg9YFuWhLmcco1yifyxoI-p67-Np-MuyEG9wZW5uaWMuaTJwZC54eXoKL2Rucy1xdWVyeQ


## opennic-R4SAS-ipv6

DNSSEC • OpenNIC • Non-logging • Uncensored - hosted on OVH - https://opennic.i2pd.xyz/
Location: Paris, France
Maintained by R4SAS - https://github.com/r4sas

sdns://AQcAAAAAAAAAF1syMDAxOjQ3MDoxZjE1OmI4MDo6NTNdIKnWMjpPJYAJJhl1FQLOIx4fdtned2yHxruyig7_2w5OIDIuZG5zY3J5cHQtY2VydC5vcGVubmljLmkycGQueHl6


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

sdns://AgcAAAAAAAAADTQ1LjYzLjExMC4xODegzBBg05yDKbYrb7x9DW35MJhpuYHn5jktXNj6QI9NgOYgRE69Z7uD-IB7OSHpOKyReLiCvVCq2xEjHwRM9fCN984TZHJhY28ucGxhbjktbnMyLmNvbQovZG5zLXF1ZXJ5


## plan9-ns2-doh-ipv6

doh server in Florida, USA. Non-logging, non-filtering, DNSSEC.
info - https://jlongua.github.io/plan9-dns/

sdns://AgcAAAAAAAAAHVsyMDAxOjE5ZjA6OTAwMjoxZDc0OjU0MDA6OjFdoMwQYNOcgym2K2-8fQ1t-TCYabmB5-Y5LVzY-kCPTYDmIEROvWe7g_iAezkh6TiskXi4gr1QqtsRIx8ETPXwjffOE2RyYWNvLnBsYW45LW5zMi5jb20KL2Rucy1xdWVyeQ


## pryv8boi

By pryv8, non Logging, uncensored, DNSSEC - hosted on contabo servers

sdns://AQcAAAAAAAAAEzE2NC42OC4xMjEuMTYyOjQ0NDMgHtKNfXpUMzPyLnXK8DauHpWm1Rqhz7LqwBBmSzdY9BIcMi5kbnNjcnlwdC1jZXJ0LnByeXY4Ym9pLm9yZw


## publicarray-au-doh

DNSSEC • OpenNIC • Non-logging • Uncensored - hosted on vultr.com
Maintained by publicarray - https://dns.seby.io

sdns://AgcAAAAAAAAADDQ1Ljc2LjExMy4zMaDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5iBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33zhBkb2guc2VieS5pbzo4NDQzCi9kbnMtcXVlcnk


## publicarray-au2-doh

DNSSEC • OpenNIC • Non-logging • Uncensored - hosted on ovh.com.au
Maintained by publicarray - https://dns.seby.io

sdns://AgcAAAAAAAAADTEzOS45OS4yMjIuNzIgmjo09yfeubylEAPZzpw5-PJ92cUkKQHCurGkTmNaAhkNZG9oLTIuc2VieS5pbwovZG5zLXF1ZXJ5


## pwoss.org-dnscrypt

No filter | No logs | DNSSEC | Nuremberg, Germany (netcup) | Maintained by https://pwoss.org/ (Dan)

sdns://AQcAAAAAAAAAEzQ1LjE0Mi4xNzYuMTcwOjQ0MzQgZF8FTkuFVFw6YP5x8ydebnTb8ONP7Ti6P0K1REyhUHgZMi5kbnNjcnlwdC1jZXJ0LnB3b3NzLm9yZw


## quad101

DNSSEC-aware public resolver by the Taiwan Network Information Center (TWNIC)
https://101.101.101.101/index_en.html

sdns://AgcAAAAAAAAAACC2vD25TAYM7EnyCH8Xw1-0g5OccnTsGH9vQUUH0njRtAxkbnMudHduaWMudHcKL2Rucy1xdWVyeQ


## quad9-dnscrypt-ip4-filter-ecs-pri

Quad9 (anycast) dnssec/no-log/filter/ecs 9.9.9.11 - 149.112.112.11

sdns://AQMAAAAAAAAADTkuOS45LjExOjg0NDMgZ8hHuMh1jNEgJFVDvnVnRt803x2EwAuMRwNo34Idhj4ZMi5kbnNjcnlwdC1jZXJ0LnF1YWQ5Lm5ldA
sdns://AQMAAAAAAAAAEzE0OS4xMTIuMTEyLjExOjg0NDMgZ8hHuMh1jNEgJFVDvnVnRt803x2EwAuMRwNo34Idhj4ZMi5kbnNjcnlwdC1jZXJ0LnF1YWQ5Lm5ldA


## quad9-dnscrypt-ip4-filter-pri

Quad9 (anycast) dnssec/no-log/filter 9.9.9.9 - 149.112.112.9 - 149.112.112.112

sdns://AQMAAAAAAAAADDkuOS45Ljk6ODQ0MyBnyEe4yHWM0SAkVUO-dWdG3zTfHYTAC4xHA2jfgh2GPhkyLmRuc2NyeXB0LWNlcnQucXVhZDkubmV0
sdns://AQMAAAAAAAAAEjE0OS4xMTIuMTEyLjk6ODQ0MyBnyEe4yHWM0SAkVUO-dWdG3zTfHYTAC4xHA2jfgh2GPhkyLmRuc2NyeXB0LWNlcnQucXVhZDkubmV0
sdns://AQMAAAAAAAAAFDE0OS4xMTIuMTEyLjExMjo4NDQzIGfIR7jIdYzRICRVQ751Z0bfNN8dhMALjEcDaN-CHYY-GTIuZG5zY3J5cHQtY2VydC5xdWFkOS5uZXQ


## quad9-dnscrypt-ip4-nofilter-ecs-pri

Quad9 (anycast) no-dnssec/no-log/no-filter/ecs 9.9.9.12 - 149.112.112.12

sdns://AQYAAAAAAAAADTkuOS45LjEyOjg0NDMgZ8hHuMh1jNEgJFVDvnVnRt803x2EwAuMRwNo34Idhj4ZMi5kbnNjcnlwdC1jZXJ0LnF1YWQ5Lm5ldA
sdns://AQYAAAAAAAAAEzE0OS4xMTIuMTEyLjEyOjg0NDMgZ8hHuMh1jNEgJFVDvnVnRt803x2EwAuMRwNo34Idhj4ZMi5kbnNjcnlwdC1jZXJ0LnF1YWQ5Lm5ldA


## quad9-dnscrypt-ip4-nofilter-pri

Quad9 (anycast) no-dnssec/no-log/no-filter 149.112.112.10

sdns://AQYAAAAAAAAAEzE0OS4xMTIuMTEyLjEwOjg0NDMgZ8hHuMh1jNEgJFVDvnVnRt803x2EwAuMRwNo34Idhj4ZMi5kbnNjcnlwdC1jZXJ0LnF1YWQ5Lm5ldA


## quad9-doh-ip4-port443-filter-ecs-pri

Quad9 (anycast) dnssec/no-log/filter/ecs 9.9.9.11 - 149.112.112.11

sdns://AgMAAAAAAAAACDkuOS45LjExICoV9dastufAkBreTrvHQ7LM1IkDK0bhZC8Gk2gwASWKE2RuczExLnF1YWQ5Lm5ldDo0NDMKL2Rucy1xdWVyeQ
sdns://AgMAAAAAAAAADjE0OS4xMTIuMTEyLjExICoV9dastufAkBreTrvHQ7LM1IkDK0bhZC8Gk2gwASWKE2RuczExLnF1YWQ5Lm5ldDo0NDMKL2Rucy1xdWVyeQ


## quad9-doh-ip4-port443-filter-pri

Quad9 (anycast) dnssec/no-log/filter 9.9.9.9 - 149.112.112.9 - 149.112.112.112

sdns://AgMAAAAAAAAABzkuOS45LjkgKhX11qy258CQGt5Ou8dDsszUiQMrRuFkLwaTaDABJYoSZG5zOS5xdWFkOS5uZXQ6NDQzCi9kbnMtcXVlcnk
sdns://AgMAAAAAAAAADTE0OS4xMTIuMTEyLjkgKhX11qy258CQGt5Ou8dDsszUiQMrRuFkLwaTaDABJYoSZG5zOS5xdWFkOS5uZXQ6NDQzCi9kbnMtcXVlcnk
sdns://AgMAAAAAAAAADzE0OS4xMTIuMTEyLjExMiAqFfXWrLbnwJAa3k67x0OyzNSJAytG4WQvBpNoMAElihFkbnMucXVhZDkubmV0OjQ0MwovZG5zLXF1ZXJ5


## quad9-doh-ip4-port443-nofilter-ecs-pri

Quad9 (anycast) no-dnssec/no-log/no-filter/ecs 9.9.9.12 - 149.112.112.12

sdns://AgYAAAAAAAAACDkuOS45LjEyICoV9dastufAkBreTrvHQ7LM1IkDK0bhZC8Gk2gwASWKE2RuczEyLnF1YWQ5Lm5ldDo0NDMKL2Rucy1xdWVyeQ
sdns://AgYAAAAAAAAADjE0OS4xMTIuMTEyLjEyICoV9dastufAkBreTrvHQ7LM1IkDK0bhZC8Gk2gwASWKE2RuczEyLnF1YWQ5Lm5ldDo0NDMKL2Rucy1xdWVyeQ


## quad9-doh-ip4-port443-nofilter-pri

Quad9 (anycast) no-dnssec/no-log/no-filter 9.9.9.10 - 149.112.112.10

sdns://AgYAAAAAAAAACDkuOS45LjEwICoV9dastufAkBreTrvHQ7LM1IkDK0bhZC8Gk2gwASWKE2RuczEwLnF1YWQ5Lm5ldDo0NDMKL2Rucy1xdWVyeQ
sdns://AgYAAAAAAAAADjE0OS4xMTIuMTEyLjEwICoV9dastufAkBreTrvHQ7LM1IkDK0bhZC8Gk2gwASWKE2RuczEwLnF1YWQ5Lm5ldDo0NDMKL2Rucy1xdWVyeQ


## quad9-doh-ip4-port5053-filter-ecs-pri

Quad9 (anycast) dnssec/no-log/filter/ecs 9.9.9.11 - 149.112.112.11

sdns://AgMAAAAAAAAACDkuOS45LjExICoV9dastufAkBreTrvHQ7LM1IkDK0bhZC8Gk2gwASWKFGRuczExLnF1YWQ5Lm5ldDo1MDUzCi9kbnMtcXVlcnk
sdns://AgMAAAAAAAAADjE0OS4xMTIuMTEyLjExICoV9dastufAkBreTrvHQ7LM1IkDK0bhZC8Gk2gwASWKFGRuczExLnF1YWQ5Lm5ldDo1MDUzCi9kbnMtcXVlcnk


## quad9-doh-ip4-port5053-filter-pri

Quad9 (anycast) dnssec/no-log/filter 9.9.9.9 - 149.112.112.9 - 149.112.112.112

sdns://AgMAAAAAAAAABzkuOS45LjkgKhX11qy258CQGt5Ou8dDsszUiQMrRuFkLwaTaDABJYoTZG5zOS5xdWFkOS5uZXQ6NTA1MwovZG5zLXF1ZXJ5
sdns://AgMAAAAAAAAADTE0OS4xMTIuMTEyLjkgKhX11qy258CQGt5Ou8dDsszUiQMrRuFkLwaTaDABJYoTZG5zOS5xdWFkOS5uZXQ6NTA1MwovZG5zLXF1ZXJ5
sdns://AgMAAAAAAAAADzE0OS4xMTIuMTEyLjExMiAqFfXWrLbnwJAa3k67x0OyzNSJAytG4WQvBpNoMAElihJkbnMucXVhZDkubmV0OjUwNTMKL2Rucy1xdWVyeQ


## quad9-doh-ip4-port5053-nofilter-ecs-pri

Quad9 (anycast) no-dnssec/no-log/no-filter/ecs 9.9.9.12 - 149.112.112.12

sdns://AgYAAAAAAAAACDkuOS45LjEyICoV9dastufAkBreTrvHQ7LM1IkDK0bhZC8Gk2gwASWKFGRuczEyLnF1YWQ5Lm5ldDo1MDUzCi9kbnMtcXVlcnk
sdns://AgYAAAAAAAAADjE0OS4xMTIuMTEyLjEyICoV9dastufAkBreTrvHQ7LM1IkDK0bhZC8Gk2gwASWKFGRuczEyLnF1YWQ5Lm5ldDo1MDUzCi9kbnMtcXVlcnk


## quad9-doh-ip4-port5053-nofilter-pri

Quad9 (anycast) no-dnssec/no-log/no-filter 9.9.9.10 - 149.112.112.10

sdns://AgYAAAAAAAAACDkuOS45LjEwICoV9dastufAkBreTrvHQ7LM1IkDK0bhZC8Gk2gwASWKFGRuczEwLnF1YWQ5Lm5ldDo1MDUzCi9kbnMtcXVlcnk
sdns://AgYAAAAAAAAADjE0OS4xMTIuMTEyLjEwICoV9dastufAkBreTrvHQ7LM1IkDK0bhZC8Gk2gwASWKFGRuczEwLnF1YWQ5Lm5ldDo1MDUzCi9kbnMtcXVlcnk


## quad9-doh-ip6-port443-filter-ecs-pri

Quad9 (anycast) dnssec/no-log/filter/ecs 2620:fe::11 - 2620:fe::fe:11

sdns://AgMAAAAAAAAADVsyNjIwOmZlOjoxMV0gKhX11qy258CQGt5Ou8dDsszUiQMrRuFkLwaTaDABJYoTZG5zMTEucXVhZDkubmV0OjQ0MwovZG5zLXF1ZXJ5
sdns://AgMAAAAAAAAAEFsyNjIwOmZlOjpmZToxMV0gKhX11qy258CQGt5Ou8dDsszUiQMrRuFkLwaTaDABJYoTZG5zMTEucXVhZDkubmV0OjQ0MwovZG5zLXF1ZXJ5


## quad9-doh-ip6-port443-filter-pri

Quad9 (anycast) dnssec/no-log/filter 2620:fe::fe - 2620:fe::9 - 2620:fe::fe:9

sdns://AgMAAAAAAAAADVsyNjIwOmZlOjpmZV0gKhX11qy258CQGt5Ou8dDsszUiQMrRuFkLwaTaDABJYoRZG5zLnF1YWQ5Lm5ldDo0NDMKL2Rucy1xdWVyeQ
sdns://AgMAAAAAAAAADFsyNjIwOmZlOjo5XSAqFfXWrLbnwJAa3k67x0OyzNSJAytG4WQvBpNoMAElihFkbnMucXVhZDkubmV0OjQ0MwovZG5zLXF1ZXJ5
sdns://AgMAAAAAAAAAD1syNjIwOmZlOjpmZTo5XSAqFfXWrLbnwJAa3k67x0OyzNSJAytG4WQvBpNoMAElihJkbnM5LnF1YWQ5Lm5ldDo0NDMKL2Rucy1xdWVyeQ


## quad9-doh-ip6-port443-nofilter-ecs-pri

Quad9 (anycast) no-dnssec/no-log/no-filter/ecs 2620:fe::12 - 2620:fe::fe:12

sdns://AgYAAAAAAAAADVsyNjIwOmZlOjoxMl0gKhX11qy258CQGt5Ou8dDsszUiQMrRuFkLwaTaDABJYoTZG5zMTIucXVhZDkubmV0OjQ0MwovZG5zLXF1ZXJ5
sdns://AgYAAAAAAAAAEFsyNjIwOmZlOjpmZToxMl0gKhX11qy258CQGt5Ou8dDsszUiQMrRuFkLwaTaDABJYoTZG5zMTIucXVhZDkubmV0OjQ0MwovZG5zLXF1ZXJ5


## quad9-doh-ip6-port443-nofilter-pri

Quad9 (anycast) no-dnssec/no-log/no-filter 2620:fe::10 - 2620:fe::fe:10

sdns://AgYAAAAAAAAADVsyNjIwOmZlOjoxMF0gKhX11qy258CQGt5Ou8dDsszUiQMrRuFkLwaTaDABJYoTZG5zMTAucXVhZDkubmV0OjQ0MwovZG5zLXF1ZXJ5
sdns://AgYAAAAAAAAAEFsyNjIwOmZlOjpmZToxMF0gKhX11qy258CQGt5Ou8dDsszUiQMrRuFkLwaTaDABJYoTZG5zMTAucXVhZDkubmV0OjQ0MwovZG5zLXF1ZXJ5


## quad9-doh-ip6-port5053-filter-ecs-pri

Quad9 (anycast) dnssec/no-log/filter/ecs 2620:fe::11 - 2620:fe::fe:11

sdns://AgMAAAAAAAAADVsyNjIwOmZlOjoxMV0gKhX11qy258CQGt5Ou8dDsszUiQMrRuFkLwaTaDABJYoUZG5zMTEucXVhZDkubmV0OjUwNTMKL2Rucy1xdWVyeQ
sdns://AgMAAAAAAAAAEFsyNjIwOmZlOjpmZToxMV0gKhX11qy258CQGt5Ou8dDsszUiQMrRuFkLwaTaDABJYoUZG5zMTEucXVhZDkubmV0OjUwNTMKL2Rucy1xdWVyeQ


## quad9-doh-ip6-port5053-filter-pri

Quad9 (anycast) dnssec/no-log/filter 2620:fe::fe - 2620:fe::9 - 2620:fe::fe:9

sdns://AgMAAAAAAAAADVsyNjIwOmZlOjpmZV0gKhX11qy258CQGt5Ou8dDsszUiQMrRuFkLwaTaDABJYoSZG5zLnF1YWQ5Lm5ldDo1MDUzCi9kbnMtcXVlcnk
sdns://AgMAAAAAAAAADFsyNjIwOmZlOjo5XSAqFfXWrLbnwJAa3k67x0OyzNSJAytG4WQvBpNoMAElihJkbnMucXVhZDkubmV0OjUwNTMKL2Rucy1xdWVyeQ
sdns://AgMAAAAAAAAAD1syNjIwOmZlOjpmZTo5XSAqFfXWrLbnwJAa3k67x0OyzNSJAytG4WQvBpNoMAElihNkbnM5LnF1YWQ5Lm5ldDo1MDUzCi9kbnMtcXVlcnk


## quad9-doh-ip6-port5053-nofilter-ecs-pri

Quad9 (anycast) no-dnssec/no-log/no-filter/ecs 2620:fe::12 - 2620:fe::fe:12

sdns://AgYAAAAAAAAADVsyNjIwOmZlOjoxMl0gKhX11qy258CQGt5Ou8dDsszUiQMrRuFkLwaTaDABJYoUZG5zMTIucXVhZDkubmV0OjUwNTMKL2Rucy1xdWVyeQ
sdns://AgYAAAAAAAAAEFsyNjIwOmZlOjpmZToxMl0gKhX11qy258CQGt5Ou8dDsszUiQMrRuFkLwaTaDABJYoUZG5zMTIucXVhZDkubmV0OjUwNTMKL2Rucy1xdWVyeQ


## quad9-doh-ip6-port5053-nofilter-pri

Quad9 (anycast) no-dnssec/no-log/no-filter 2620:fe::10 - 2620:fe::fe:10

sdns://AgYAAAAAAAAADVsyNjIwOmZlOjoxMF0gKhX11qy258CQGt5Ou8dDsszUiQMrRuFkLwaTaDABJYoUZG5zMTAucXVhZDkubmV0OjUwNTMKL2Rucy1xdWVyeQ
sdns://AgYAAAAAAAAAEFsyNjIwOmZlOjpmZToxMF0gKhX11qy258CQGt5Ou8dDsszUiQMrRuFkLwaTaDABJYoUZG5zMTAucXVhZDkubmV0OjUwNTMKL2Rucy1xdWVyeQ


## resolver4.dns.openinternet.io

DNSCrypt & DoH resolver on dedicated hardware colocated at Sonic.net in Santa Rosa, CA in the United States.

No log, no filter, DNSSEC. Uses Sonic's recusrive DNS servers as upstream resolvers (but is not affiliated with Sonic
in any way). Provided by https://openinternet.io

sdns://AQcAAAAAAAAADTcwLjM2LjE3MC4xMjYgIMqRyWkPSPlDAmN_2ne3A6A6EQtQRyEJcXoCWPHbX5EtMi5kbnNjcnlwdC1jZXJ0LnJlc29sdmVyNC5kbnMub3BlbmludGVybmV0Lmlv


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


## sgp-dn53

non-censoring, non-logging, DNSSEC-capable based on official docker image hosted in Singapore

sdns://AQcAAAAAAAAADDIwLjIxMi41Mi43NiDjxysoCLj_kqq1dtoYCpB9_eZ3vMihC0hAJx7wyc5aRBgyLmRuc2NyeXB0LWNlcnQuc2dwLWRuNTM


## sgp-dn53-ipv6

non-censoring, non-logging, DNSSEC-capable based on official docker image hosted in Singapore

sdns://AQcAAAAAAAAAH1syYTAxOjExMTpmMTAwOjcwMDA6OjZmZGQ6NTNjNl0g48crKAi4_5KqtXbaGAqQff3md7zIoQtIQCce8MnOWkQYMi5kbnNjcnlwdC1jZXJ0LnNncC1kbjUz


## sth-dnscrypt-se

Resolver in Stockholm, Sweden. DNSCrypt server. Non-logging, non-filtering, DNSSEC.

sdns://AQcAAAAAAAAADTQ1LjE1My4xODcuOTYgeXlaphvMhHruyoIiBZWFCQqgabN91D0bQEBb3nJ9yUYfMi5kbnNjcnlwdC1jZXJ0LnN0aC1kbnNjcnlwdC1zZQ


## sth-doh-se

Resolver in Stockholm, Sweden. DoH server. Non-logging, non-filtering, DNSSEC.

sdns://AgcAAAAAAAAADTQ1LjE1My4xODcuOTagzBBg05yDKbYrb7x9DW35MJhpuYHn5jktXNj6QI9NgOYgRE69Z7uD-IB7OSHpOKyReLiCvVCq2xEjHwRM9fCN984SZG5zc2UuYWxla2JlcmcubmV0Ci9kbnMtcXVlcnk


## sth-doh-se-ipv6

Resolver in Stockholm, Sweden. DoH server. Non-logging, non-filtering, DNSSEC.

sdns://AgcAAAAAAAAAFVsyYTA5OmNkNDI6Zjo0MjViOjoxXaDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5iBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33zhJkbnNzZS5hbGVrYmVyZy5uZXQKL2Rucy1xdWVyeQ


## switch

Public DoH service provided by SWITCH in Switzerland

https://www.switch.ch

Provides protection against malware, but does not block ads.

sdns://AgMAAAAAAAAAACA7FAZ73R-Sz_NJdsl8yiP0L0H3yntGnmJAVtqSpC3Nmg1kbnMuc3dpdGNoLmNoCi9kbnMtcXVlcnk


## tuna-doh-ipv4

DoH server provided by Tsinghua University TUNA Association, located in mainland China, no GFW poisoning yet it has a manual blacklist.

sdns://AgEAAAAAAAAACTEwMS42LjYuNiCtaqQXa1OG2FEfYZ3bQLmrut2k3X_Bc2TbYUSD_cyoRw4xMDEuNi42LjY6ODQ0MwovZG5zLXF1ZXJ5


## uncensoreddns-dk-ipv4

Also known as censurfridns.
DoH, no logs, no filter, DNSSEC, unicast hosted in Denmark - https://blog.uncensoreddns.org

sdns://AgcAAAAAAAAADDg5LjIzMy40My43MaDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5iBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33zhl1bmljYXN0LnVuY2Vuc29yZWRkbnMub3JnCi9kbnMtcXVlcnk


## uncensoreddns-dk-ipv6

Also known as censurfridns.
DoH, no logs, no filter, DNSSEC, unicast hosted in Denmark - https://blog.uncensoreddns.org

sdns://AgcAAAAAAAAAElsyYTAxOjNhMDo1Mzo1Mzo6XaDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5iBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33zhl1bmljYXN0LnVuY2Vuc29yZWRkbnMub3JnCi9kbnMtcXVlcnk


## uncensoreddns-ipv4

Also known as censurfridns.
DoH, no logs, no filter, DNSSEC, anycast - https://blog.uncensoreddns.org

sdns://AgcAAAAAAAAADjkxLjIzOS4xMDAuMTAwoMwQYNOcgym2K2-8fQ1t-TCYabmB5-Y5LVzY-kCPTYDmIEROvWe7g_iAezkh6TiskXi4gr1QqtsRIx8ETPXwjffOGWFueWNhc3QudW5jZW5zb3JlZGRucy5vcmcKL2Rucy1xdWVyeQ


## uncensoreddns-ipv6

Also known as censurfridns.
DoH, no logs, no filter, DNSSEC, anycast - https://blog.uncensoreddns.org

sdns://AgcAAAAAAAAAEVsyMDAxOjY3YzoyOGE0OjpdoMwQYNOcgym2K2-8fQ1t-TCYabmB5-Y5LVzY-kCPTYDmIEROvWe7g_iAezkh6TiskXi4gr1QqtsRIx8ETPXwjffOGWFueWNhc3QudW5jZW5zb3JlZGRucy5vcmcKL2Rucy1xdWVyeQ


## userspace-australia

DNSCrypt in Australia (Brisbane & Melbourne) by UserSpace.
No logs | IPv4 | Filtered

sdns://AQMAAAAAAAAAEDEwMy4xNi4xMzEuNzc6NTQgnRNtOLv4IzxEfkbLFOaHa-ncLImdQiP-pS1jaFY5jlUdMi5kbnNjcnlwdC1jZXJ0LnVzZXJzcGFjZS1ibmU
sdns://AQMAAAAAAAAAEjEwMy4yMzYuMTYyLjExOTo1NCBPr5jCD_2geOTMmS5LgQg_v79pgppTm3vLZhe_oahbgR0yLmRuc2NyeXB0LWNlcnQudXNlcnNwYWNlLW1lbA


## userspace-australia-ipv6

DNSCrypt in Australia (Brisbane & Melbourne) by UserSpace.
No logs | IPv6 | Filtered

sdns://AQMAAAAAAAAAJVsyNDA0Ojk0MDA6MTowOjIxNjozZWZmOmZlZjA6MTgwYV06NTQgnRNtOLv4IzxEfkbLFOaHa-ncLImdQiP-pS1jaFY5jlUdMi5kbnNjcnlwdC1jZXJ0LnVzZXJzcGFjZS1ibmU
sdns://AQMAAAAAAAAAJVsyNDA0Ojk0MDA6MzowOjIxNjozZWZmOmZlZTA6N2Y2OV06NTQgT6-Ywg_9oHjkzJkuS4EIP7-_aYKaU5t7y2YXv6GoW4EdMi5kbnNjcnlwdC1jZXJ0LnVzZXJzcGFjZS1tZWw


## v.dnscrypt.uk-doh-ipv4

DoH, no logs, uncensored, DNSSEC. Hosted in London UK on Digital Ocean
https://www.dnscrypt.uk

sdns://AgcAAAAAAAAADTEzNi4yNDQuNjUuMjAAEXYuZG5zY3J5cHQudWs6NDQzCi9kbnMtcXVlcnk


## v.dnscrypt.uk-doh-ipv6

DoH, no logs, uncensored, DNSSEC. Hosted in London UK on Digital Ocean
https://www.dnscrypt.uk

sdns://AgcAAAAAAAAAJ1syYTA1OmY0ODA6MTAwMDo5Njc6NTQwMDozZmY6ZmU3ZTo1NWQzXQARdi5kbnNjcnlwdC51azo0NDMKL2Rucy1xdWVyeQ


## v.dnscrypt.uk-ipv4

DNSCrypt, no logs, uncensored, DNSSEC. Hosted in London UK on Digital Ocean
https://www.dnscrypt.uk

sdns://AQcAAAAAAAAADzEwNC4yMzguMTg2LjE5MiDtST2M6teQZk8GPEe3lZojaS18kDY8nkPMtZF75bQe5R0yLmRuc2NyeXB0LWNlcnQudi5kbnNjcnlwdC51aw


## v.dnscrypt.uk-ipv6

DNSCrypt, no logs, uncensored, DNSSEC. Hosted in London UK on Digital Ocean
https://www.dnscrypt.uk

sdns://AQcAAAAAAAAAKFsyMDAxOjE5ZjA6NzQwMjoxNTc0OjU0MDA6MmZmOmZlNjY6MmNmZl0g7Uk9jOrXkGZPBjxHt5WaI2ktfJA2PJ5DzLWRe-W0HuUdMi5kbnNjcnlwdC1jZXJ0LnYuZG5zY3J5cHQudWs


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

