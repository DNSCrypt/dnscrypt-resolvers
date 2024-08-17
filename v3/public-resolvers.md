ri# public-resolvers

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

Non-filtering, No-logging, DNSSEC DoH operated by Andrews & Arnold LTD.
Homepage: https://www.aa.net.uk/dns/

sdns://AgcAAAAAAAAADTIxNy4xNjkuMjAuMjIADWRucy5hYS5uZXQudWsKL2Rucy1xdWVyeQ
sdns://AgcAAAAAAAAADTIxNy4xNjkuMjAuMjMADWRucy5hYS5uZXQudWsKL2Rucy1xdWVyeQ


## a-and-a-ipv6

Non-filtering, No-logging, DNSSEC DoH over IPv6 operated by Andrews & Arnold LTD.
Homepage: https://www.aa.net.uk/dns/

sdns://AgcAAAAAAAAAEFsyMDAxOjhiMDo6MjAyMl0ADWRucy5hYS5uZXQudWsKL2Rucy1xdWVyeQ
sdns://AgcAAAAAAAAAEFsyMDAxOjhiMDo6MjAyM10ADWRucy5hYS5uZXQudWsKL2Rucy1xdWVyeQ


## adfilter-adl

Hosted in Adelaide, Australia.

Blocks ads, malware, trackers and more. No persistent logs. DNSSEC. No EDNS Client-Subnet.

sdns://AgMAAAAAAAAADjE2My40Ny4xMTcuMTc2oMwQYNOcgym2K2-8fQ1t-TCYabmB5-Y5LVzY-kCPTYDmoPf1ryiAHod9ffOivij-FJ8ydKftKfE2_VA845jLqAsNoLNeBZUM-9gln5N1uhAYcLjDxMDsWlKXV-YxZ-neJqnooEROvWe7g_iAezkh6TiskXi4gr1QqtsRIx8ETPXwjffOoOZEumlj4zX-dly5l2sSsQ61QpS0JHd2TMs6OsyjrLL8ICquP7e_BeTIHEGU3KRFEdT5rzBHhuwa5yGECc9ioINVEGFkbC5hZGZpbHRlci5uZXQKL2Rucy1xdWVyeQ


## adfilter-adl-ipv6

Hosted in Adelaide, Australia.

Blocks ads, malware, trackers and more. No persistent logs. DNSSEC. No EDNS Client-Subnet.

sdns://AgMAAAAAAAAAHlsyNDAwOmM0MDE6OjUwNTQ6ZmY6ZmUxYjpiMDM2XaDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5qD39a8ogB6HfX3zor4o_hSfMnSn7SnxNv1QPOOYy6gLDaCzXgWVDPvYJZ-TdboQGHC4w8TA7FpSl1fmMWfp3iap6KBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33zqDmRLppY-M1_nZcuZdrErEOtUKUtCR3dkzLOjrMo6yy_CAqrj-3vwXkyBxBlNykRRHU-a8wR4bsGuchhAnPYqCDVRBhZGwuYWRmaWx0ZXIubmV0Ci9kbnMtcXVlcnk


## adfilter-per

Hosted in Perth, Australia.

Blocks ads, malware, trackers and more. No persistent logs. DNSSEC. No EDNS Client-Subnet.

sdns://AgMAAAAAAAAADTIwMy4yOS4yNDEuNzagzBBg05yDKbYrb7x9DW35MJhpuYHn5jktXNj6QI9NgOag9_WvKIAeh31986K-KP4UnzJ0p-0p8Tb9UDzjmMuoCw2gs14FlQz72CWfk3W6EBhwuMPEwOxaUpdX5jFn6d4mqeigRE69Z7uD-IB7OSHpOKyReLiCvVCq2xEjHwRM9fCN986g5kS6aWPjNf52XLmXaxKxDrVClLQkd3ZMyzo6zKOssvwgKq4_t78F5MgcQZTcpEUR1PmvMEeG7BrnIYQJz2Kgg1UQcGVyLmFkZmlsdGVyLm5ldAovZG5zLXF1ZXJ5


## adfilter-per-ipv6

Hosted in Perth, Australia.

Blocks ads, malware, trackers and more. No persistent logs. DNSSEC. No EDNS Client-Subnet.

sdns://AgMAAAAAAAAAGFsyNDA0Ojk0MDA6NDFhOTo0ODAwOjoxXaDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5qD39a8ogB6HfX3zor4o_hSfMnSn7SnxNv1QPOOYy6gLDaCzXgWVDPvYJZ-TdboQGHC4w8TA7FpSl1fmMWfp3iap6KBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33zqDmRLppY-M1_nZcuZdrErEOtUKUtCR3dkzLOjrMo6yy_CAqrj-3vwXkyBxBlNykRRHU-a8wR4bsGuchhAnPYqCDVRBwZXIuYWRmaWx0ZXIubmV0Ci9kbnMtcXVlcnk


## adfilter-syd

Hosted in Sydney, Australia.

Blocks ads, malware, trackers and more. No persistent logs. DNSSEC. No EDNS Client-Subnet.

sdns://AgMAAAAAAAAADjExMi4yMTMuMzIuMjE5oMwQYNOcgym2K2-8fQ1t-TCYabmB5-Y5LVzY-kCPTYDmoPf1ryiAHod9ffOivij-FJ8ydKftKfE2_VA845jLqAsNoLNeBZUM-9gln5N1uhAYcLjDxMDsWlKXV-YxZ-neJqnooEROvWe7g_iAezkh6TiskXi4gr1QqtsRIx8ETPXwjffOoOZEumlj4zX-dly5l2sSsQ61QpS0JHd2TMs6OsyjrLL8ICquP7e_BeTIHEGU3KRFEdT5rzBHhuwa5yGECc9ioINVEHN5ZC5hZGZpbHRlci5uZXQKL2Rucy1xdWVyeQ


## adfilter-syd-ipv6

Hosted in Sydney, Australia.

Blocks ads, malware, trackers and more. No persistent logs. DNSSEC. No EDNS Client-Subnet.

sdns://AgMAAAAAAAAAGFsyNDA0Ojk0MDA6MjE0ZTplYTAwOjoxXaDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5qD39a8ogB6HfX3zor4o_hSfMnSn7SnxNv1QPOOYy6gLDaCzXgWVDPvYJZ-TdboQGHC4w8TA7FpSl1fmMWfp3iap6KBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33zqDmRLppY-M1_nZcuZdrErEOtUKUtCR3dkzLOjrMo6yy_CAqrj-3vwXkyBxBlNykRRHU-a8wR4bsGuchhAnPYqCDVRBzeWQuYWRmaWx0ZXIubmV0Ci9kbnMtcXVlcnk


## adfree.usableprivacy.net-ipv4

Public updns DoH service with advertising, tracker and malware filters.

Hosted in Europe by @usableprivacy, details see: https://docs.usableprivacy.com

sdns://AgMAAAAAAAAADTc4LjQ3LjE2My4xNDGgJMDCcXhcb_MSalmIxClhGOov_5PNfxMeTDDB3xSfphsg5kS6aWPjNf52XLmXaxKxDrVClLQkd3ZMyzo6zKOssvwYYWRmcmVlLnVzYWJsZXByaXZhY3kubmV0Ci9kbnMtcXVlcnk


## adfree.usableprivacy.net-ipv6

Public updns IPv6 DoH service with advertising, tracker and malware filters.

Hosted in Europe by @usableprivacy, details see: https://docs.usableprivacy.com

sdns://AgMAAAAAAAAAFlsyYTAxOjRmODoxYzBjOjgzMmQ6Ol2gJMDCcXhcb_MSalmIxClhGOov_5PNfxMeTDDB3xSfphsg5kS6aWPjNf52XLmXaxKxDrVClLQkd3ZMyzo6zKOssvwYYWRmcmVlLnVzYWJsZXByaXZhY3kubmV0Ci9kbnMtcXVlcnk


## adguard-dns

Remove ads and protect your computer from malware

Warning: This server is incompatible with anonymization.

sdns://AQMAAAAAAAAAETk0LjE0MC4xNC4xNDo1NDQzINErR_JS3PLCu_iZEIbq95zkSV2LFsigxDIuUso_OQhzIjIuZG5zY3J5cHQuZGVmYXVsdC5uczEuYWRndWFyZC5jb20
sdns://AQMAAAAAAAAAETk0LjE0MC4xNS4xNTo1NDQzINErR_JS3PLCu_iZEIbq95zkSV2LFsigxDIuUso_OQhzIjIuZG5zY3J5cHQuZGVmYXVsdC5uczEuYWRndWFyZC5jb20


## adguard-dns-doh

Remove ads and protect your computer from malware (over DoH)

sdns://AgMAAAAAAAAADDk0LjE0MC4xNC4xNCCaOjT3J965vKUQA9nOnDn48n3ZxSQpAcK6saROY1oCGQw5NC4xNDAuMTQuMTQKL2Rucy1xdWVyeQ
sdns://AgMAAAAAAAAADDk0LjE0MC4xNS4xNSCaOjT3J965vKUQA9nOnDn48n3ZxSQpAcK6saROY1oCGQw5NC4xNDAuMTUuMTUKL2Rucy1xdWVyeQ


## adguard-dns-family

Adguard DNS with safesearch and adult content blocking

Warning: This server is incompatible with anonymization.

sdns://AQMAAAAAAAAAETk0LjE0MC4xNC4xNTo1NDQzILgxXdexS27jIKRw3C7Wsao5jMnlhvhdRUXWuMm1AFq6ITIuZG5zY3J5cHQuZmFtaWx5Lm5zMS5hZGd1YXJkLmNvbQ
sdns://AQMAAAAAAAAAETk0LjE0MC4xNS4xNjo1NDQzILgxXdexS27jIKRw3C7Wsao5jMnlhvhdRUXWuMm1AFq6ITIuZG5zY3J5cHQuZmFtaWx5Lm5zMS5hZGd1YXJkLmNvbQ


## adguard-dns-family-doh

Adguard DNS with safesearch and adult content blocking (over DoH)

sdns://AgMAAAAAAAAADDk0LjE0MC4xNC4xNSCaOjT3J965vKUQA9nOnDn48n3ZxSQpAcK6saROY1oCGQw5NC4xNDAuMTQuMTUKL2Rucy1xdWVyeQ
sdns://AgMAAAAAAAAADDk0LjE0MC4xNS4xNiCaOjT3J965vKUQA9nOnDn48n3ZxSQpAcK6saROY1oCGQw5NC4xNDAuMTUuMTYKL2Rucy1xdWVyeQ


## adguard-dns-family-ipv6

Adguard DNS with safesearch and adult content blocking (Over IPv6)

Warning: This server is incompatible with anonymization.

sdns://AQMAAAAAAAAAGVsyYTEwOjUwYzA6OmJhZDE6ZmZdOjU0NDMguDFd17FLbuMgpHDcLtaxqjmMyeWG-F1FRda4ybUAWrohMi5kbnNjcnlwdC5mYW1pbHkubnMxLmFkZ3VhcmQuY29t
sdns://AQMAAAAAAAAAGVsyYTEwOjUwYzA6OmJhZDI6ZmZdOjU0NDMguDFd17FLbuMgpHDcLtaxqjmMyeWG-F1FRda4ybUAWrohMi5kbnNjcnlwdC5mYW1pbHkubnMxLmFkZ3VhcmQuY29t


## adguard-dns-ipv6

Remove ads and protect your computer from malware (over IPv6)

Warning: This server is incompatible with anonymization.

sdns://AQMAAAAAAAAAGFsyYTEwOjUwYzA6OmFkMTpmZl06NTQ0MyDRK0fyUtzywrv4mRCG6vec5EldixbIoMQyLlLKPzkIcyIyLmRuc2NyeXB0LmRlZmF1bHQubnMxLmFkZ3VhcmQuY29t
sdns://AQMAAAAAAAAAGFsyYTEwOjUwYzA6OmFkMjpmZl06NTQ0MyDRK0fyUtzywrv4mRCG6vec5EldixbIoMQyLlLKPzkIcyIyLmRuc2NyeXB0LmRlZmF1bHQubnMxLmFkZ3VhcmQuY29t


## adguard-dns-unfiltered

AdGuard public DNS servers without filters

Warning: This server is incompatible with anonymization.

sdns://AQcAAAAAAAAAEjk0LjE0MC4xNC4xNDA6NTQ0MyC16ETWuDo-PhJo62gfvqcN48X6aNvWiBQdvy7AZrLa-iUyLmRuc2NyeXB0LnVuZmlsdGVyZWQubnMxLmFkZ3VhcmQuY29t
sdns://AQcAAAAAAAAAEjk0LjE0MC4xNC4xNDE6NTQ0MyC16ETWuDo-PhJo62gfvqcN48X6aNvWiBQdvy7AZrLa-iUyLmRuc2NyeXB0LnVuZmlsdGVyZWQubnMxLmFkZ3VhcmQuY29t


## adguard-dns-unfiltered-doh

AdGuard public DNS servers without filters (over DoH)

sdns://AgcAAAAAAAAADTk0LjE0MC4xNC4xNDAgmjo09yfeubylEAPZzpw5-PJ92cUkKQHCurGkTmNaAhkNOTQuMTQwLjE0LjE0MAovZG5zLXF1ZXJ5
sdns://AgcAAAAAAAAADTk0LjE0MC4xNC4xNDEgmjo09yfeubylEAPZzpw5-PJ92cUkKQHCurGkTmNaAhkNOTQuMTQwLjE0LjE0MQovZG5zLXF1ZXJ5


## adguard-dns-unfiltered-ipv6

AdGuard public DNS servers without filters (over IPv6)

Warning: This server is incompatible with anonymization.

sdns://AQcAAAAAAAAAFlsyYTEwOjUwYzA6OjE6ZmZdOjU0NDMgtehE1rg6Pj4SaOtoH76nDePF-mjb1ogUHb8uwGay2volMi5kbnNjcnlwdC51bmZpbHRlcmVkLm5zMS5hZGd1YXJkLmNvbQ
sdns://AQcAAAAAAAAAFlsyYTEwOjUwYzA6OjI6ZmZdOjU0NDMgtehE1rg6Pj4SaOtoH76nDePF-mjb1ogUHb8uwGay2volMi5kbnNjcnlwdC51bmZpbHRlcmVkLm5zMS5hZGd1YXJkLmNvbQ


## ahadns-doh-la

A zero logging DNS with support for DNS-over-HTTPS (DoH) & DNS-over-TLS (DoT). Blocks ads, malware, trackers, viruses, ransomware, telemetry and more. No persistent logs. DNSSEC. Hosted in Los Angeles, USA. By https://ahadns.com/
Server statistics can be seen at: https://statistics.ahadns.com/?server=la

sdns://AgMAAAAAAAAADTQ1LjY3LjIxOS4yMDigzBBg05yDKbYrb7x9DW35MJhpuYHn5jktXNj6QI9NgOag9_WvKIAeh31986K-KP4UnzJ0p-0p8Tb9UDzjmMuoCw2gs14FlQz72CWfk3W6EBhwuMPEwOxaUpdX5jFn6d4mqeigRE69Z7uD-IB7OSHpOKyReLiCvVCq2xEjHwRM9fCN986g5kS6aWPjNf52XLmXaxKxDrVClLQkd3ZMyzo6zKOssvwgKq4_t78F5MgcQZTcpEUR1PmvMEeG7BrnIYQJz2Kgg1URZG9oLmxhLmFoYWRucy5uZXQKL2Rucy1xdWVyeQ


## ahadns-doh-nl

A zero logging DNS with support for DNS-over-HTTPS (DoH) & DNS-over-TLS (DoT). Blocks ads, malware, trackers, viruses, ransomware, telemetry and more. No persistent logs. DNSSEC. Hosted in Amsterdam, Netherlands. By https://ahadns.com/
Server statistics can be seen at: https://statistics.ahadns.com/?server=nl

sdns://AgMAAAAAAAAACTUuMi43NS43NaDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5qD39a8ogB6HfX3zor4o_hSfMnSn7SnxNv1QPOOYy6gLDaCzXgWVDPvYJZ-TdboQGHC4w8TA7FpSl1fmMWfp3iap6KBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33zqDmRLppY-M1_nZcuZdrErEOtUKUtCR3dkzLOjrMo6yy_CAqrj-3vwXkyBxBlNykRRHU-a8wR4bsGuchhAnPYqCDVRFkb2gubmwuYWhhZG5zLm5ldAovZG5zLXF1ZXJ5


## alidns-doh

A public DNS resolver that supports DoH/DoT in mainland China, provided by Alibaba-Cloud.
Homepage: https://alidns.com

Warning: GFW filtering rules are applied by this resolver.

sdns://AgAAAAAAAAAACTIyMy41LjUuNSCY49XlNq8pWM0vfxT3BO9KJ20l4zzWXy5l9eTycnwTMAkyMjMuNS41LjUKL2Rucy1xdWVyeQ
sdns://AgAAAAAAAAAACTIyMy42LjYuNiCY49XlNq8pWM0vfxT3BO9KJ20l4zzWXy5l9eTycnwTMAkyMjMuNi42LjYKL2Rucy1xdWVyeQ
sdns://AgAAAAAAAAAADTEyMC41NS4yMDMuNDQgmOPV5TavKVjNL38U9wTvSidtJeM81l8uZfXk8nJ8EzANMTIwLjU1LjIwMy40NAovZG5zLXF1ZXJ5
sdns://AgAAAAAAAAAACzQ3LjEwOC4wLjYzIJjj1eU2rylYzS9_FPcE70onbSXjPNZfLmX15PJyfBMwCzQ3LjEwOC4wLjYzCi9kbnMtcXVlcnk
sdns://AgAAAAAAAAAADTM5LjEwMy4yNi4yMDQgmOPV5TavKVjNL38U9wTvSidtJeM81l8uZfXk8nJ8EzANMzkuMTAzLjI2LjIwNAovZG5zLXF1ZXJ5
sdns://AgAAAAAAAAAADzEzOS4xMjkuMTM3LjEzNyCY49XlNq8pWM0vfxT3BO9KJ20l4zzWXy5l9eTycnwTMA8xMzkuMTI5LjEzNy4xMzcKL2Rucy1xdWVyeQ
sdns://AgAAAAAAAAAACzQ3LjEyMi44LjExIJjj1eU2rylYzS9_FPcE70onbSXjPNZfLmX15PJyfBMwCzQ3LjEyMi44LjExCi9kbnMtcXVlcnk
sdns://AgAAAAAAAAAADjEyMy4xODQuMTk4LjIyIJjj1eU2rylYzS9_FPcE70onbSXjPNZfLmX15PJyfBMwDjEyMy4xODQuMTk4LjIyCi9kbnMtcXVlcnk
sdns://AgAAAAAAAAAADjExMy4xNDIuODMuMTMyIJjj1eU2rylYzS9_FPcE70onbSXjPNZfLmX15PJyfBMwDjExMy4xNDIuODMuMTMyCi9kbnMtcXVlcnk
sdns://AgAAAAAAAAAADjExMy4yNDAuODguMTcyIJjj1eU2rylYzS9_FPcE70onbSXjPNZfLmX15PJyfBMwDjExMy4yNDAuODguMTcyCi9kbnMtcXVlcnk
sdns://AgAAAAAAAAAADTEyMy4xNTEuMTA3LjUgmOPV5TavKVjNL38U9wTvSidtJeM81l8uZfXk8nJ8EzANMTIzLjE1MS4xMDcuNQovZG5zLXF1ZXJ5
sdns://AgAAAAAAAAAADDE4Mi40MC43MC4xMiCY49XlNq8pWM0vfxT3BO9KJ20l4zzWXy5l9eTycnwTMAwxODIuNDAuNzAuMTIKL2Rucy1xdWVyeQ
sdns://AgAAAAAAAAAADTguMTI5LjE1Mi4yMzAgmOPV5TavKVjNL38U9wTvSidtJeM81l8uZfXk8nJ8EzANOC4xMjkuMTUyLjIzMAovZG5zLXF1ZXJ5
sdns://AgAAAAAAAAAACjEuNzEuMjAuMzcgmOPV5TavKVjNL38U9wTvSidtJeM81l8uZfXk8nJ8EzAKMS43MS4yMC4zNwovZG5zLXF1ZXJ5


## alidns-doh-ipv6

A public DNS resolver over IPv6 that supports DoH/DoT in mainland China, provided by Alibaba-Cloud.
Homepage: https://alidns.com

Warning: GFW filtering rules are applied by this resolver.

sdns://AgAAAAAAAAAADlsyNDAwOjMyMDA6OjFdIJjj1eU2rylYzS9_FPcE70onbSXjPNZfLmX15PJyfBMwCTIyMy41LjUuNQovZG5zLXF1ZXJ5
sdns://AgAAAAAAAAAAE1syNDAwOjMyMDA6YmFiYTo6MV0gmOPV5TavKVjNL38U9wTvSidtJeM81l8uZfXk8nJ8EzAJMjIzLjUuNS41Ci9kbnMtcXVlcnk


## ams-ads-doh-nl

Resolver in Amsterdam. HTTP/3, DoH protocol. Non-logging. Blocks ads, malware and trackers. DNSSEC enabled.

sdns://AgMAAAAAAAAADDg5LjM4LjEzMS4zOAAYZG5zbmwtbm9hZHMuYWxla2JlcmcubmV0Ci9kbnMtcXVlcnk


## ams-dnscrypt-nl

Resolver in Amsterdam. Dnscrypt protocol. Non-logging, non-filtering, DNSSEC.

sdns://AQcAAAAAAAAAETg5LjM4LjEzMS4zODo0MzQzIKWHS9r0FoKY--wcnJl1Ar5aOUb91xsufvPUjid3rNRaHzIuZG5zY3J5cHQtY2VydC5hbXMtZG5zY3J5cHQtbmw


## ams-dnscrypt-nl-ipv6

Resolver in Amsterdam. Dnscrypt protocol. Non-logging, non-filtering, DNSSEC.

sdns://AQcAAAAAAAAAGlsyYTBjOmI5YzA6Zjo0NTFkOjoxXTo0MzQzIKWHS9r0FoKY--wcnJl1Ar5aOUb91xsufvPUjid3rNRaHzIuZG5zY3J5cHQtY2VydC5hbXMtZG5zY3J5cHQtbmw


## ams-doh-nl

Resolver in Amsterdam. HTTP/3, DoH protocol. Non-logging, non-filtering, DNSSEC.

sdns://AgcAAAAAAAAADDg5LjM4LjEzMS4zOKDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5qD39a8ogB6HfX3zor4o_hSfMnSn7SnxNv1QPOOYy6gLDaCzXgWVDPvYJZ-TdboQGHC4w8TA7FpSl1fmMWfp3iap6KBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33zqDmRLppY-M1_nZcuZdrErEOtUKUtCR3dkzLOjrMo6yy_CAqrj-3vwXkyBxBlNykRRHU-a8wR4bsGuchhAnPYqCDVRJkbnNubC5hbGVrYmVyZy5uZXQKL2Rucy1xdWVyeQ


## ams-doh-nl-ipv6

Resolver in Amsterdam. HTTP/3, DoH protocol. Non-logging, non-filtering, DNSSEC.

sdns://AgcAAAAAAAAAFVsyYTBjOmI5YzA6Zjo0NTFkOjoxXaDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5qD39a8ogB6HfX3zor4o_hSfMnSn7SnxNv1QPOOYy6gLDaCzXgWVDPvYJZ-TdboQGHC4w8TA7FpSl1fmMWfp3iap6KBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33zqDmRLppY-M1_nZcuZdrErEOtUKUtCR3dkzLOjrMo6yy_CAqrj-3vwXkyBxBlNykRRHU-a8wR4bsGuchhAnPYqCDVRJkbnNubC5hbGVrYmVyZy5uZXQKL2Rucy1xdWVyeQ


## artikel10-doh-ipv4

DNSSEC, No-filter and No-log DoH resolver operated by Artikel10 association. Homepage: https://artikel10.org

sdns://AgcAAAAAAAAADjIxNy4xOTcuOTEuMTUzoMwQYNOcgym2K2-8fQ1t-TCYabmB5-Y5LVzY-kCPTYDmoPf1ryiAHod9ffOivij-FJ8ydKftKfE2_VA845jLqAsNoLNeBZUM-9gln5N1uhAYcLjDxMDsWlKXV-YxZ-neJqnooEROvWe7g_iAezkh6TiskXi4gr1QqtsRIx8ETPXwjffOoOZEumlj4zX-dly5l2sSsQ61QpS0JHd2TMs6OsyjrLL8ICquP7e_BeTIHEGU3KRFEdT5rzBHhuwa5yGECc9ioINVEWRucy5hcnRpa2VsMTAub3JnCi9kbnMtcXVlcnk


## artikel10-doh-ipv6

DNSSEC, No-filter and No-log DoH resolver (IPv6) operated by Artikel10 association. Homepage: https://artikel10.org

sdns://AgcAAAAAAAAAF1syMDAxOjY3YzoxNDAxOjIxMjA6OjFdoMwQYNOcgym2K2-8fQ1t-TCYabmB5-Y5LVzY-kCPTYDmoPf1ryiAHod9ffOivij-FJ8ydKftKfE2_VA845jLqAsNoLNeBZUM-9gln5N1uhAYcLjDxMDsWlKXV-YxZ-neJqnooEROvWe7g_iAezkh6TiskXi4gr1QqtsRIx8ETPXwjffOoOZEumlj4zX-dly5l2sSsQ61QpS0JHd2TMs6OsyjrLL8ICquP7e_BeTIHEGU3KRFEdT5rzBHhuwa5yGECc9ioINVEWRucy5hcnRpa2VsMTAub3JnCi9kbnMtcXVlcnk


## bebasdns

BebasDNS default server by BebasID. DNSSEC and OpenNIC supported. Filters ads, tracker, and malware.

sdns://AQIAAAAAAAAAEjEwMy44Ny42OC4xOTQ6ODQ0MyAxXDKkdrOao8ZeLyu7vTnVrT0C7YlPNNf6trdMkje7QR8yLmRuc2NyeXB0LWNlcnQuZG5zLmJlYmFzaWQuY29t


## bebasdns-family

BebasDNS Family Variant by BebasID. DNSSEC and OpenNIC supported. Blocks malicious links, pornography, gambling, and hate websites.

sdns://AQIAAAAAAAAAEjEwMy44Ny42OC4xOTY6ODQ0MyD5k4vgIHmBCZ2DeLtmoDVu1C6nVrRNzSVgZ1T0m0-3rCkyLmRuc2NyeXB0LWNlcnQuaW50ZXJuZXRzZWhhdC5iZWJhc2lkLmNvbQ


## bebasdns-unfiltered-doh

BebasDNS by BebasID. DNSSEC and OpenNIC supported. This variant doesn't block anything

sdns://AgcAAAAAAAAADTEwMy44Ny42OC4xOTQAD2Rucy5iZWJhc2lkLmNvbQsvdW5maWx0ZXJlZA


## bortzmeyer

Non-logging DoH server in France operated by Stéphane Bortzmeyer.

https://www.bortzmeyer.org/doh-bortzmeyer-fr-policy.html

sdns://AgcAAAAAAAAADDE5My43MC44NS4xMaDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5qD39a8ogB6HfX3zor4o_hSfMnSn7SnxNv1QPOOYy6gLDaCzXgWVDPvYJZ-TdboQGHC4w8TA7FpSl1fmMWfp3iap6KBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33zqDmRLppY-M1_nZcuZdrErEOtUKUtCR3dkzLOjrMo6yy_CAqrj-3vwXkyBxBlNykRRHU-a8wR4bsGuchhAnPYqCDVRFkb2guYm9ydHptZXllci5mcgEv


## bortzmeyer-ipv6

Non-logging DoH server in France operated by Stéphane Bortzmeyer (IPv6 only).

https://www.bortzmeyer.org/doh-bortzmeyer-fr-policy.html

sdns://AgcAAAAAAAAAGVsyMDAxOjQxZDA6MzAyOjIyMDA6OjE4MF2gzBBg05yDKbYrb7x9DW35MJhpuYHn5jktXNj6QI9NgOag9_WvKIAeh31986K-KP4UnzJ0p-0p8Tb9UDzjmMuoCw2gs14FlQz72CWfk3W6EBhwuMPEwOxaUpdX5jFn6d4mqeigRE69Z7uD-IB7OSHpOKyReLiCvVCq2xEjHwRM9fCN986g5kS6aWPjNf52XLmXaxKxDrVClLQkd3ZMyzo6zKOssvwgKq4_t78F5MgcQZTcpEUR1PmvMEeG7BrnIYQJz2Kgg1URZG9oLmJvcnR6bWV5ZXIuZnIBLw


## brahma-world

DNS-over-HTTPS server. Non Logging, filters ads, trackers and malware. DNSSEC ready, QNAME Minimization, No EDNS Client-Subnet.

Hosted in Nuremberg, Germany. (https://dns.brahma.world)

sdns://AgMAAAAAAAAADTE1Ny45MC4xMjQuNjKgzBBg05yDKbYrb7x9DW35MJhpuYHn5jktXNj6QI9NgOag9_WvKIAeh31986K-KP4UnzJ0p-0p8Tb9UDzjmMuoCw2gs14FlQz72CWfk3W6EBhwuMPEwOxaUpdX5jFn6d4mqeigRE69Z7uD-IB7OSHpOKyReLiCvVCq2xEjHwRM9fCN986g5kS6aWPjNf52XLmXaxKxDrVClLQkd3ZMyzo6zKOssvwgKq4_t78F5MgcQZTcpEUR1PmvMEeG7BrnIYQJz2Kgg1UQZG5zLmJyYWhtYS53b3JsZAovZG5zLXF1ZXJ5


## brahma-world-ipv6

DNS-over-HTTPS server. Non Logging, filters ads, trackers and malware. DNSSEC ready, QNAME Minimization, No EDNS Client-Subnet.

Hosted in Nuremberg, Germany. (https://dns.brahma.world)

sdns://AgMAAAAAAAAAF1syYTAxOjRmODoxYzFjOmY1ZTE6OjFdoMwQYNOcgym2K2-8fQ1t-TCYabmB5-Y5LVzY-kCPTYDmoPf1ryiAHod9ffOivij-FJ8ydKftKfE2_VA845jLqAsNoLNeBZUM-9gln5N1uhAYcLjDxMDsWlKXV-YxZ-neJqnooEROvWe7g_iAezkh6TiskXi4gr1QqtsRIx8ETPXwjffOoOZEumlj4zX-dly5l2sSsQ61QpS0JHd2TMs6OsyjrLL8ICquP7e_BeTIHEGU3KRFEdT5rzBHhuwa5yGECc9ioINVEGRucy5icmFobWEud29ybGQKL2Rucy1xdWVyeQ


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

Warning: Doesn't work any more in some countries such as France and Portugal.

Warning: This server is incompatible with anonymization.

Warning: modifies your queries to include a copy of your network
address when forwarding them to a selection of companies and organizations.

sdns://AQEAAAAAAAAADjIwOC42Ny4yMjAuMjIwILc1EUAgbyJdPivYItf9aR6hwzzI1maNDL4Ev6vKQ_t5GzIuZG5zY3J5cHQtY2VydC5vcGVuZG5zLmNvbQ


## cisco-doh

Remove your DNS blind spot (DoH protocol)

Warning: Doesn't work any more in some countries such as France and Portugal.

Warning: modifies your queries to include a copy of your network
address when forwarding them to a selection of companies and organizations.

sdns://AgAAAAAAAAAADDE0Ni4xMTIuNDEuMiCYZO337qhZZ1J0sPrfvSaTZamrnrp3PahnSUxalKQ33w9kb2gub3BlbmRucy5jb20KL2Rucy1xdWVyeQ


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


## cisco-ipv6

Cisco OpenDNS over IPv6 (DNSCrypt protocol)

Warning: Doesn't work any more in some countries such as France and Portugal.

Warning: This server is incompatible with anonymization.

Warning: modifies your queries to include a copy of your network
address when forwarding them to a selection of companies and organizations.

sdns://AQEAAAAAAAAAD1syNjIwOjA6Y2NjOjoyXSC3NRFAIG8iXT4r2CLX_WkeocM8yNZmjQy-BL-rykP7eRsyLmRuc2NyeXB0LWNlcnQub3BlbmRucy5jb20


## cisco-ipv6-doh

Cisco OpenDNS over IPv6 (DoH protocol)

Warning: Doesn't work any more in some countries such as France and Portugal.

Warning: modifies your queries to include a copy of your network
address when forwarding them to a selection of companies and organizations.

sdns://AgAAAAAAAAAAEFsyNjIwOjExOTpmYzo6Ml0gmGTt9-6oWWdSdLD6370mk2Wpq566dz2oZ0lMWpSkN98PZG9oLm9wZW5kbnMuY29tCi9kbnMtcXVlcnk


## cisco-sandbox

Cisco OpenDNS Sandbox (anycast)

Warning: Doesn't work any more in some countries such as France and Portugal.

Warning: This server is incompatible with anonymization.

sdns://AQEAAAAAAAAADDE0Ni4xMTIuNDEuNCC3NRFAIG8iXT4r2CLX_WkeocM8yNZmjQy-BL-rykP7eRsyLmRuc2NyeXB0LWNlcnQub3BlbmRucy5jb20


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


## cleanbrowsing-security

Blocks only phishing, spam and malicious domains.

Warning: This server is incompatible with anonymization.

sdns://AQMAAAAAAAAAEjE4NS4yMjguMTY4Ljk6ODQ0MyC8rDL61UNpFx8IMtYHUCfDIIzu8Ojpn5QY3HdgZdSPKRFjbGVhbmJyb3dzaW5nLm9yZw
sdns://AQMAAAAAAAAAEjE4NS4yMjguMTY5Ljk6ODQ0MyC8rDL61UNpFx8IMtYHUCfDIIzu8Ojpn5QY3HdgZdSPKRFjbGVhbmJyb3dzaW5nLm9yZw


## cleanbrowsing-security-doh

Blocks only phishing, spam and malicious domains over DoH.

sdns://AgMAAAAAAAAADjE4NS4yMjguMTY4LjEwoPn_N_AuYyy3OHAlwH5XkIo9Nxt8ldjN0DkN4jHtlDoSoCso0AXN1mJZ2xEYZeoXy7YLPI9UcGhjjZAqZL54Sv34IOaSTdvwPj_u_RiUGT7gQuBqadbySK2eIW2kKyiPLBAZEWNsZWFuYnJvd3Npbmcub3JnFC9kb2gvc2VjdXJpdHktZmlsdGVy
sdns://AgMAAAAAAAAADzE4NS4yMjguMTY4LjE2OKD5_zfwLmMstzhwJcB-V5CKPTcbfJXYzdA5DeIx7ZQ6EqArKNAFzdZiWdsRGGXqF8u2CzyPVHBoY42QKmS-eEr9-CDmkk3b8D4_7v0YlBk-4ELgamnW8kitniFtpCsojywQGRFjbGVhbmJyb3dzaW5nLm9yZxQvZG9oL3NlY3VyaXR5LWZpbHRlcg


## cleanbrowsing-security-ipv6

Blocks only phishing, spam and malicious domains over IPv6.

Warning: This server is incompatible with anonymization.

sdns://AQMAAAAAAAAAFVsyYTBkOjJhMDA6MTo6Ml06ODQ0MyC8rDL61UNpFx8IMtYHUCfDIIzu8Ojpn5QY3HdgZdSPKRFjbGVhbmJyb3dzaW5nLm9yZw
sdns://AQMAAAAAAAAAFVsyYTBkOjJhMDA6Mjo6Ml06ODQ0MyC8rDL61UNpFx8IMtYHUCfDIIzu8Ojpn5QY3HdgZdSPKRFjbGVhbmJyb3dzaW5nLm9yZw


## cloudflare

Cloudflare DNS (anycast) - aka 1.1.1.1 / 1.0.0.1

sdns://AgcAAAAAAAAABzEuMC4wLjEAEmRucy5jbG91ZGZsYXJlLmNvbQovZG5zLXF1ZXJ5
sdns://AgcAAAAAAAAABzEuMS4xLjEABzEuMS4xLjEKL2Rucy1xdWVyeQ
sdns://AgcAAAAAAAAABzEuMC4wLjEABzEuMC4wLjEKL2Rucy1xdWVyeQ
sdns://AgcAAAAAAAAADDE2Mi4xNTkuMzYuMQAMMTYyLjE1OS4zNi4xCi9kbnMtcXVlcnk
sdns://AgcAAAAAAAAADDE2Mi4xNTkuNDYuMQAMMTYyLjE1OS40Ni4xCi9kbnMtcXVlcnk
sdns://AgcAAAAAAAAADjEwNC4xNi4xMzIuMjI5ABJkbnMuY2xvdWRmbGFyZS5jb20KL2Rucy1xdWVyeQ
sdns://AgcAAAAAAAAADjEwNC4xNi4xMzMuMjI5ABJkbnMuY2xvdWRmbGFyZS5jb20KL2Rucy1xdWVyeQ
sdns://AgcAAAAAAAAADjEwNC4xNi4yNDkuMjQ5ABJjbG91ZGZsYXJlLWRucy5jb20KL2Rucy1xdWVyeQ
sdns://AgcAAAAAAAAADjEwNC4xNi4yNDguMjQ5ABJjbG91ZGZsYXJlLWRucy5jb20KL2Rucy1xdWVyeQ


## cloudflare-family

Cloudflare DNS (anycast) with malware protection and parental control - aka 1.1.1.3 / 1.0.0.3

sdns://AgMAAAAAAAAABzEuMS4xLjMABzEuMS4xLjMKL2Rucy1xdWVyeQ
sdns://AgMAAAAAAAAABzEuMC4wLjMABzEuMC4wLjMKL2Rucy1xdWVyeQ


## cloudflare-family-ipv6

Cloudflare DNS over IPv6 (anycast) with malware protection and parental control

sdns://AgMAAAAAAAAAAAAWWzI2MDY6NDcwMDo0NzAwOjoxMTEzXQovZG5zLXF1ZXJ5
sdns://AgMAAAAAAAAAAAAWWzI2MDY6NDcwMDo0NzAwOjoxMDAzXQovZG5zLXF1ZXJ5


## cloudflare-ipv6

Cloudflare DNS over IPv6 (anycast)

sdns://AgcAAAAAAAAAAAAWWzI2MDY6NDcwMDo0NzAwOjoxMTExXQovZG5zLXF1ZXJ5
sdns://AgcAAAAAAAAAAAAWWzI2MDY6NDcwMDo0NzAwOjoxMDAxXQovZG5zLXF1ZXJ5
sdns://AgcAAAAAAAAAAAAUWzI2MDY6NDcwMDo0NzAwOjo2NF0KL2Rucy1xdWVyeQ
sdns://AgcAAAAAAAAAAAAWWzI2MDY6NDcwMDo0NzAwOjo2NDAwXQovZG5zLXF1ZXJ5
sdns://AgcAAAAAAAAAFlsyNjA2OjQ3MDA6OjY4MTA6ODRlNV0AEmRucy5jbG91ZGZsYXJlLmNvbQovZG5zLXF1ZXJ5
sdns://AgcAAAAAAAAAFlsyNjA2OjQ3MDA6OjY4MTA6ODVlNV0AEmRucy5jbG91ZGZsYXJlLmNvbQovZG5zLXF1ZXJ5
sdns://AgcAAAAAAAAAFlsyNjA2OjQ3MDA6OjY4MTA6ZjhmOV0AEmNsb3VkZmxhcmUtZG5zLmNvbQovZG5zLXF1ZXJ5
sdns://AgcAAAAAAAAAFlsyNjA2OjQ3MDA6OjY4MTA6ZjlmOV0AEmNsb3VkZmxhcmUtZG5zLmNvbQovZG5zLXF1ZXJ5


## cloudflare-security

Cloudflare DNS (anycast) with malware blocking - aka 1.1.1.2 / 1.0.0.2

sdns://AgMAAAAAAAAABzEuMS4xLjIABzEuMS4xLjIKL2Rucy1xdWVyeQ
sdns://AgMAAAAAAAAABzEuMC4wLjIABzEuMC4wLjIKL2Rucy1xdWVyeQ


## cloudflare-security-ipv6

Cloudflare DNS over IPv6 (anycast) with malware blocking

sdns://AgMAAAAAAAAAAAAWWzI2MDY6NDcwMDo0NzAwOjoxMTEyXQovZG5zLXF1ZXJ5
sdns://AgMAAAAAAAAAAAAWWzI2MDY6NDcwMDo0NzAwOjoxMDAyXQovZG5zLXF1ZXJ5


## comodo-02

Comodo Dome Shield (anycast) - https://cdome.comodo.com/shield/

sdns://AQUAAAAAAAAACjguMjAuMjQ3LjIg0sJUqpYcHsoXmZb1X7yAHwg2xyN5q1J-zaiGG-Dgs7AoMi5kbnNjcnlwdC1jZXJ0LnNoaWVsZC0yLmRuc2J5Y29tb2RvLmNvbQ
sdns://AQUAAAAAAAAACzguMjAuMjQ3LjIwINLCVKqWHB7KF5mW9V-8gB8INscjeatSfs2ohhvg4LOwKDIuZG5zY3J5cHQtY2VydC5zaGllbGQtMi5kbnNieWNvbW9kby5jb20
sdns://AQUAAAAAAAAACjguMjYuNTYuMjYg0sJUqpYcHsoXmZb1X7yAHwg2xyN5q1J-zaiGG-Dgs7AoMi5kbnNjcnlwdC1jZXJ0LnNoaWVsZC0yLmRuc2J5Y29tb2RvLmNvbQ


## controld-block-malware

ControlD Free DNS. Take CONTROL of your Internet. Block unwanted content, bypass geo-restrictions and be more productive. DoH protocol and No logging. - https://controld.com/free-dns

This DNS blocks Malware domains.

sdns://AgMAAAAAAAAACjc2Ljc2LjIuMTEAFGZyZWVkbnMuY29udHJvbGQuY29tAy9wMQ


## controld-block-malware-ad

ControlD Free DNS. Take CONTROL of your Internet. Block unwanted content, bypass geo-restrictions and be more productive. DoH protocol and No logging. - https://controld.com/free-dns

This DNS blocks Malware, Ads & Tracking domains.

sdns://AgMAAAAAAAAACjc2Ljc2LjIuMTEAFGZyZWVkbnMuY29udHJvbGQuY29tAy9wMg


## controld-block-malware-ad-social

ControlD Free DNS. Take CONTROL of your Internet. Block unwanted content, bypass geo-restrictions and be more productive. DoH protocol and No logging. - https://controld.com/free-dns

This DNS blocks Malware, Ads & Tracking and Social Networks domains.

sdns://AgMAAAAAAAAACjc2Ljc2LjIuMTEAFGZyZWVkbnMuY29udHJvbGQuY29tAy9wMw


## controld-family-friendly

ControlD Free DNS. Take CONTROL of your Internet. Block unwanted content, bypass geo-restrictions and be more productive. DoH protocol and No logging. - https://controld.com/free-dns

This DNS blocks Malware, Ads & Tracking, Adult Content and Drugs domains.

sdns://AgMAAAAAAAAACjc2Ljc2LjIuMTEAFGZyZWVkbnMuY29udHJvbGQuY29tBy9mYW1pbHk


## controld-uncensored

ControlD Free DNS. Take CONTROL of your Internet. Block unwanted content, bypass geo-restrictions and be more productive. DoH protocol and No logging. - https://controld.com/free-dns

This DNS unblocks censored domains from various countries.

sdns://AgcAAAAAAAAACjc2Ljc2LjIuMTEAFGZyZWVkbnMuY29udHJvbGQuY29tCy91bmNlbnNvcmVk


## controld-unfiltered

ControlD Free DNS. Take CONTROL of your Internet. Block unwanted content, bypass geo-restrictions and be more productive. DoH protocol and No logging. - https://controld.com/free-dns

This is a Unfiltered DNS, no DNS record blocking or manipulation here, if you want to block Malware, Ads & Tracking or Social Network domains, use the other ControlD DNS configs.

sdns://AgcAAAAAAAAACjc2Ljc2LjIuMTEAFGZyZWVkbnMuY29udHJvbGQuY29tAy9wMA


## cs-austria

Wien, Austria DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADTk0LjE5OC40MS4yMzUgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz


## cs-barcelona

Barcelona, Spain DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADjM3LjEyMC4xNDIuMTE1IDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-belgium

Brussels, Belgium DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADTM3LjEyMC4yMzYuMTEgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz


## cs-berlin

Berlin, Germany DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADTM3LjEyMC4yMTcuNzUgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz


## cs-brazil

Sao Paulo, Brazil DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADjE3Ny41NC4xNDUuMTMxIDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-bulgaria

Sofia, Bulgaria DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADjM3LjEyMC4xNTIuMjM1IDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-ch

Switzerland DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADzE5MC4yMTEuMjU1LjIyNyAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM


## cs-ch2

Switzerland 2 DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADTE0Ni43MC4xMzUuNTkgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz


## cs-czech

Prague, Czech Republic DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADzIxNy4xMzguMjIwLjI0MyAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM


## cs-dc

US - Washington, DC DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADDE5OC43LjU4LjIyNyAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM


## cs-de

Frankfurt, Germany DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAACzE0Ni43MC44Mi4zIDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


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


## cs-fl

US - Miami, FL DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADDIzLjE5LjExNy41NSAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM


## cs-fl2

US - Miami, FL 2 DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADDIzLjE5LjExNy41MSAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM


## cs-fr

France DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADTE2My4xNzIuMzQuNTYgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz


## cs-ga

US - Atlanta, GA DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADTY0LjQyLjE4MS4yMjcgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz


## cs-il

US - Chicago, IL DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADjE3My4yMzQuNTYuMTE1IDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-il2

US - Chicago, IL 2 DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADjEwOC4xODEuNjMuMTYzIDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-india

India DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADzE2NS4yMzEuMjUzLjE2MyAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM


## cs-ireland

Dublin, Ireland DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADjM3LjEyMC4yMzUuMTg3IDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-la

US - Los Angeles, CA DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADzE5NS4yMDYuMTA0LjIwMyAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM


## cs-london

London, England DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADTc4LjEyOS4yNDguNjcgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz


## cs-lv

Latvia DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADzEwOS4yNDguMTQ5LjEzMyAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM


## cs-madrid

Madrid, Spain DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADjE4NS4xODMuMTA2LjgzIDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-manchester

Manchester, England DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADTE5NS4xMi40OC4xNzEgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz


## cs-mexico

Mexico City, Mexico DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADTEwMy4xNC4yNi4xOTAgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz


## cs-milan

Milan, Italy DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADzIxNy4xMzguMjE5LjIxOSAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM


## cs-montreal

Montreal, Canada DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADTE3Ni4xMTMuNzQuMTkgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz


## cs-nc

US - North Carolina DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADjE1NS4yNTQuMjEuMjUwIDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-nl

Netherlands DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADTE4NS4xMDcuODAuODQgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz


## cs-nl2

Netherlands 2 DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADjEwOC4xODEuMTI0LjI3IDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


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

sdns://AQYAAAAAAAAADTE3OS42MS4yMjMuNDcgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz


## cs-poland

Warsaw, Poland DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADTM3LjEyMC4yMTEuOTEgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz


## cs-pt

Portugal DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADjkxLjIwNS4yMzAuMjI0IDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-ro

Romania DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADTE0Ni43MC42Ni4yMjcgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz


## cs-rome

Rome, Italy DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADjM3LjEyMC4yMDcuMTMxIDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-serbia

Belgrade, Serbia DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADjM3LjEyMC4xOTMuMjE5IDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-singapore

Singapore DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADTM3LjEyMC4xNTEuMTEgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz


## cs-sk

South Korea DNSCrypt server provided by https://cryptostorm.is/

sdns://AQIAAAAAAAAADjEwOC4xODEuNTAuMjE4IDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-slovakia

Bratislava, Slovakia DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADjE5My4zNy4yNTUuMjI3IDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-swe

Sweden DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADzEyOC4xMjcuMTA0LjEwOCAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM


## cs-sydney

Sydney, Australia DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADjM3LjEyMC4yMzQuMjUxIDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-tokyo

Tokyo, Japan DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADDE0Ni43MC4zMS40MyAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM


## cs-tx

US - Dallas, TX DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADTIwOS41OC4xNDcuMzYgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz


## cs-tx2

US - Dallas, TX 2 DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADjEwOC4xODEuMTAxLjI3IDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-tx3

US - Dallas, TX 3 DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADjEwOC4xODEuMTAxLjY3IDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-vancouver

Vancouver, Canada DNSCrypt server provided by https://cryptostorm.is/

sdns://AQYAAAAAAAAADjE5Ni4yNDAuNzkuMTYzIDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## dct-de

DNSCrypt | IPv4 only | Non-logging | Non-filtering | DNSSEC | Frankfurt, Germany.

sdns://AQcAAAAAAAAADTE4NS4xNy4xNDYuODAgj59Zfjt7-XPMmhsZtTS5HZP_oA9SLu8gqWl3Ve2M27EWMi5kbnNjcnlwdC1jZXJ0LmRjdC1kZQ


## dct-nl

DNSCrypt | IPv4 only | Non-logging | Non-filtering | DNSSEC | Dronten, Netherlands.

sdns://AQcAAAAAAAAADDUuMjU1LjEwNS4yNCBmA48agbqHYOdV4gbPDt9f9ZMTAQzZfznm0NcBa8e_QRYyLmRuc2NyeXB0LWNlcnQuZGN0LW5s


## decloudus-nogoogle-tstipv6

Servers helps you deGoogle and unGoogle by completely blocking Google tracking in addition to annoying ads, online trackers, and malware. No Logs. For IPv6.

Contributed by: https://decloudus.com

sdns://AQIAAAAAAAAAHFsyYTAxOjRmODoxM2E6MjUwYjo6MzBdOjk0NDMgTUTeE2lcp5HMJFQG0gawr6-F93NwsoM_-FQ1HEYYYWMeMi5kbnNjcnlwdC1jZXJ0LkRlQ2xvdWRVcy10ZXN0


## deffer-dns.au

DNSSEC/Non-logged/Uncensored in Sydney (AWS).
Encrypted DNS Server image by jedisct1, maintaned by DeffeR.

sdns://AQcAAAAAAAAADTUyLjY1LjIzNS4xMjkg5Q00RDDBkwx3fUaa0_etjz4iH3lLBOqsg95bYDmV07MdMi5kbnNjcnlwdC1jZXJ0LmRlZmZlci1kbnMuYXU


## digitalprivacy.diy-dnscrypt-ipv4

IPv4 server | No filter | No logs | DNSSEC | Nuremberg, Germany (netcup) | Maintained by https://digitalprivacy.diy/

sdns://AQcAAAAAAAAAEjM3LjIyMS4xOTQuODQ6NDQzNCCiyGRvm0TcyJmI7lTXstgh-8AoAAiFcTQQp7Od_brTYCIyLmRuc2NyeXB0LWNlcnQuZGlnaXRhbHByaXZhY3kuZGl5


## dns.digitale-gesellschaft.ch

Public DoH resolver operated by the Digital Society (https://www.digitale-gesellschaft.ch).
Hosted in Zurich, Switzerland.

Non-logging, non-filtering, supports DNSSEC.

sdns://AgcAAAAAAAAADTE4NS45NS4yMTguNDKgzBBg05yDKbYrb7x9DW35MJhpuYHn5jktXNj6QI9NgOag9_WvKIAeh31986K-KP4UnzJ0p-0p8Tb9UDzjmMuoCw2gs14FlQz72CWfk3W6EBhwuMPEwOxaUpdX5jFn6d4mqeigRE69Z7uD-IB7OSHpOKyReLiCvVCq2xEjHwRM9fCN986g5kS6aWPjNf52XLmXaxKxDrVClLQkd3ZMyzo6zKOssvwgKq4_t78F5MgcQZTcpEUR1PmvMEeG7BrnIYQJz2Kgg1UcZG5zLmRpZ2l0YWxlLWdlc2VsbHNjaGFmdC5jaAovZG5zLXF1ZXJ5
sdns://AgcAAAAAAAAADTE4NS45NS4yMTguNDOgzBBg05yDKbYrb7x9DW35MJhpuYHn5jktXNj6QI9NgOag9_WvKIAeh31986K-KP4UnzJ0p-0p8Tb9UDzjmMuoCw2gs14FlQz72CWfk3W6EBhwuMPEwOxaUpdX5jFn6d4mqeigRE69Z7uD-IB7OSHpOKyReLiCvVCq2xEjHwRM9fCN986g5kS6aWPjNf52XLmXaxKxDrVClLQkd3ZMyzo6zKOssvwgKq4_t78F5MgcQZTcpEUR1PmvMEeG7BrnIYQJz2Kgg1UcZG5zLmRpZ2l0YWxlLWdlc2VsbHNjaGFmdC5jaAovZG5zLXF1ZXJ5


## dns.digitale-gesellschaft.ch-ipv6

Public IPv6 DoH resolver operated by the Digital Society (https://www.digitale-gesellschaft.ch).
Hosted in Zurich, Switzerland.

Non-logging, non-filtering, supports DNSSEC.

sdns://AgcAAAAAAAAAD1syYTA1OmZjODQ6OjQyXaDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5qD39a8ogB6HfX3zor4o_hSfMnSn7SnxNv1QPOOYy6gLDaCzXgWVDPvYJZ-TdboQGHC4w8TA7FpSl1fmMWfp3iap6KBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33zqDmRLppY-M1_nZcuZdrErEOtUKUtCR3dkzLOjrMo6yy_CAqrj-3vwXkyBxBlNykRRHU-a8wR4bsGuchhAnPYqCDVRxkbnMuZGlnaXRhbGUtZ2VzZWxsc2NoYWZ0LmNoCi9kbnMtcXVlcnk
sdns://AgcAAAAAAAAAD1syYTA1OmZjODQ6OjQzXaDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5qD39a8ogB6HfX3zor4o_hSfMnSn7SnxNv1QPOOYy6gLDaCzXgWVDPvYJZ-TdboQGHC4w8TA7FpSl1fmMWfp3iap6KBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33zqDmRLppY-M1_nZcuZdrErEOtUKUtCR3dkzLOjrMo6yy_CAqrj-3vwXkyBxBlNykRRHU-a8wR4bsGuchhAnPYqCDVRxkbnMuZGlnaXRhbGUtZ2VzZWxsc2NoYWZ0LmNoCi9kbnMtcXVlcnk


## dns.digitalsize.net

A public, non-tracking, non-filtering DNS resolver with DNSSEC enabled and hosted in Germany (https://dns.digitalsize.net)

sdns://AgcAAAAAAAAADjk0LjEzMC4xMzUuMjAzoMwQYNOcgym2K2-8fQ1t-TCYabmB5-Y5LVzY-kCPTYDmoPf1ryiAHod9ffOivij-FJ8ydKftKfE2_VA845jLqAsNoLNeBZUM-9gln5N1uhAYcLjDxMDsWlKXV-YxZ-neJqnooEROvWe7g_iAezkh6TiskXi4gr1QqtsRIx8ETPXwjffOoOZEumlj4zX-dly5l2sSsQ61QpS0JHd2TMs6OsyjrLL8ICquP7e_BeTIHEGU3KRFEdT5rzBHhuwa5yGECc9ioINVE2Rucy5kaWdpdGFsc2l6ZS5uZXQKL2Rucy1xdWVyeQ


## dns.digitalsize.net-ipv6

A public, non-tracking, non-filtering DNS resolver with DNSSEC enabled and hosted in Germany (https://dns.digitalsize.net)

sdns://AgcAAAAAAAAAGVsyYTAxOjRmODoxM2I6MzQwNzo6ZmFjZV2gzBBg05yDKbYrb7x9DW35MJhpuYHn5jktXNj6QI9NgOag9_WvKIAeh31986K-KP4UnzJ0p-0p8Tb9UDzjmMuoCw2gs14FlQz72CWfk3W6EBhwuMPEwOxaUpdX5jFn6d4mqeigRE69Z7uD-IB7OSHpOKyReLiCvVCq2xEjHwRM9fCN986g5kS6aWPjNf52XLmXaxKxDrVClLQkd3ZMyzo6zKOssvwgKq4_t78F5MgcQZTcpEUR1PmvMEeG7BrnIYQJz2Kgg1UTZG5zLmRpZ2l0YWxzaXplLm5ldAovZG5zLXF1ZXJ5


## dns.ryan-palmer

Non-logging, non-filtering, DNSSEC DoH Server. Hosted in the UK.

sdns://AgcAAAAAAAAADjY4LjE4My4yNTMuMjAwoMwQYNOcgym2K2-8fQ1t-TCYabmB5-Y5LVzY-kCPTYDmoPf1ryiAHod9ffOivij-FJ8ydKftKfE2_VA845jLqAsNoLNeBZUM-9gln5N1uhAYcLjDxMDsWlKXV-YxZ-neJqnooEROvWe7g_iAezkh6TiskXi4gr1QqtsRIx8ETPXwjffOoOZEumlj4zX-dly5l2sSsQ61QpS0JHd2TMs6OsyjrLL8ICquP7e_BeTIHEGU3KRFEdT5rzBHhuwa5yGECc9ioINVFGRuczEucnlhbi1wYWxtZXIuY29tCi9kbnMtcXVlcnk


## dns.sb

DoH server runned by xTom.com. No logs, no filtering, supports DNSSEC.

Homepage: https://dns.sb

sdns://AgcAAAAAAAAADzE4NS4yMjIuMjIyLjIyMiAOp5Svj-oV-Fz-65-8H2VKHLKJ0egmfEgrdPeAQlUFFA8xODUuMjIyLjIyMi4yMjIKL2Rucy1xdWVyeQ
sdns://AgcAAAAAAAAACzQ1LjExLjQ1LjExIA6nlK-P6hX4XP7rn7wfZUocsonR6CZ8SCt094BCVQUUCzQ1LjExLjQ1LjExCi9kbnMtcXVlcnk


## dns0

A free, sovereign and GDPR-compliant recursive DNS resolver with a
strong focus on security to protect the citizens and organizations of
the European Union.

Blocks new domains, dynamic DNS, parked domains, uncommon TLDs, etc.

https://www.dns0.eu/

sdns://AgMAAAAAAAAAACCaOjT3J965vKUQA9nOnDn48n3ZxSQpAcK6saROY1oCGQdkbnMwLmV1Ci9kbnMtcXVlcnk


## dns0-kids

A free, sovereign and GDPR-compliant recursive DNS resolver with a
strong focus on security to protect the citizens and organizations of
the European Union.

This version blocks content not suitable for children.

Also blocks new domains, dynamic DNS, parked domains, uncommon TLDs, etc.

https://www.dns0.eu/

sdns://AgMAAAAAAAAAACCaOjT3J965vKUQA9nOnDn48n3ZxSQpAcK6saROY1oCGQxraWRzLmRuczAuZXUKL2Rucy1xdWVyeQ


## dns0-unfiltered

The unfiltered version of dns0.eu.
https://open.dns0.eu/

sdns://AgcAAAAAAAAAACCaOjT3J965vKUQA9nOnDn48n3ZxSQpAcK6saROY1oCGQxvcGVuLmRuczAuZXUKL2Rucy1xdWVyeQ


## dnscry.pt-adelaide-ipv4

DNSCry.pt Adelaide - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADjIzLjE0MC4yNDguMTAwIFa3zBQNs5jjEISHskpY7WSNK4sLj_qrbFiLk5tSBN1uGTIuZG5zY3J5cHQtY2VydC5kbnNjcnkucHQ


## dnscry.pt-adelaide-ipv6

DNSCry.pt Adelaide - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAFlsyNjAyOmZiYTE6YTAwOjoxMDA6MV0gVrfMFA2zmOMQhIeySljtZI0riwuP-qtsWIuTm1IE3W4ZMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeS5wdA


## dnscry.pt-allentown-ipv4

DNSCry.pt Allentown - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADTIzLjEzNy4yNTMuMjQg3Z0YI7udXIjKWcPC5GdTm4Uk6D1x2DuyYuj2OZz2cKQZMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeS5wdA


## dnscry.pt-allentown-ipv6

DNSCry.pt Allentown - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAHFsyNjAyOmZjMjQ6MTk6NzRiMDo1Mjg1OjoxMl0g3Z0YI7udXIjKWcPC5GdTm4Uk6D1x2DuyYuj2OZz2cKQZMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeS5wdA


## dnscry.pt-amsterdam-ipv4

DNSCry.pt Amsterdam - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADTQ1Ljg2LjE2Mi4xMTAgblxXPJozaH3d0T9h_69Op1nnYQYbW4yIWd8ypOORnK8ZMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeS5wdA


## dnscry.pt-ashburn-ipv4

DNSCry.pt Ashburn - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAACzQ1LjExLjIzMC44IMGyYyUUH-ohVO5gxPJoOoTQYe6WeqqivutZK9FR5v2eGTIuZG5zY3J5cHQtY2VydC5kbnNjcnkucHQ


## dnscry.pt-ashburn-ipv6

DNSCry.pt Ashburn - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAFVsyMDAxOjQ3MDo4OjE2OTo6MTAwXSDBsmMlFB_qIVTuYMTyaDqE0GHulnqqor7rWSvRUeb9nhkyLmRuc2NyeXB0LWNlcnQuZG5zY3J5LnB0


## dnscry.pt-athens-ipv4

DNSCry.pt Athens - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADTE4NS4yMzQuNTIuODcg7sJacnOa_EK646WTMceomii6ew1ZjD2YPZq6T3cbAZYZMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeS5wdA


## dnscry.pt-athens-ipv6

DNSCry.pt Athens - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAFVsyYTA5OmNkNDM6Zjo0MmExOjo1XSDuwlpyc5r8QrrjpZMxx6iaKLp7DVmMPZg9mrpPdxsBlhkyLmRuc2NyeXB0LWNlcnQuZG5zY3J5LnB0


## dnscry.pt-atlanta-ipv4

DNSCry.pt Atlanta - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADzE3MC4yNDkuMjM3LjE1NCDi7_UCIU8-uBI-dM7qpE0Y0Qo8GpJTDcSX578fvK7jOhkyLmRuc2NyeXB0LWNlcnQuZG5zY3J5LnB0


## dnscry.pt-atlanta-ipv6

DNSCry.pt Atlanta - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAE1syNjAwOjRjMDA6ODA6ODo6YV0g4u_1AiFPPrgSPnTO6qRNGNEKPBqSUw3El-e_H7yu4zoZMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeS5wdA


## dnscry.pt-auckland-ipv4

DNSCry.pt Auckland - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADjE4NS45OS4xMzMuMTEyIBWQZQSuMzmL_YANsdr26wFOHmJCYEtA2P2JI6w1-0ezGTIuZG5zY3J5cHQtY2VydC5kbnNjcnkucHQ


## dnscry.pt-auckland-ipv6

DNSCry.pt Auckland - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAHFsyYTA2OjEyODA6YmVlMToyOjplZTEyOjIwOF0gFZBlBK4zOYv9gA2x2vbrAU4eYkJgS0DY_YkjrDX7R7MZMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeS5wdA


## dnscry.pt-budapest-ipv4

DNSCry.pt Budapest - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADTE5NC4xMDIuMjI3LjUg_cf2xKCt4ZqfeHjW_2pukHH4EMfQZR_CVNkIVci8enMZMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeS5wdA


## dnscry.pt-budapest-ipv6

DNSCry.pt Budapest - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAFlsyYTBlOjFkODA6Mjc6NjI4OTo6MV0g_cf2xKCt4ZqfeHjW_2pukHH4EMfQZR_CVNkIVci8enMZMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeS5wdA


## dnscry.pt-chicago-ipv4

DNSCry.pt Chicago - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADTQ1LjQxLjIwNC4yMDQgbQ_3dUnLx_3R3UeHibflzQIDKCqMGcViiAPftt2eDbIZMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeS5wdA


## dnscry.pt-chicago-ipv6

DNSCry.pt Chicago - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAH1syNjAyOmZlYTc6ZTBjOmU6YmZmOjY6NzA6MTk0Y10gbQ_3dUnLx_3R3UeHibflzQIDKCqMGcViiAPftt2eDbIZMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeS5wdA


## dnscry.pt-chisinau-ipv4

DNSCry.pt Chișinău - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADjE3Ni4xMjMuMTAuMTA1IEJtkG567ZvN_tTXhVcSyywcrDRhziwxmbnyohp5u8gPGTIuZG5zY3J5cHQtY2VydC5kbnNjcnkucHQ


## dnscry.pt-chisinau-ipv6

DNSCry.pt Chișinău - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAHVsyMDAxOjY3ODo2ZDQ6NTA4MDo6M2RlYToxMDldIEJtkG567ZvN_tTXhVcSyywcrDRhziwxmbnyohp5u8gPGTIuZG5zY3J5cHQtY2VydC5kbnNjcnkucHQ


## dnscry.pt-coeurdalene-ipv4

DNSCry.pt Coeur d'Alene - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADzEwNC4xOTIuMTAyLjEzMiAh0DCEwDEN5foF_DicSr6yuabwtz2HckSxLlfHs3B5PRkyLmRuc2NyeXB0LWNlcnQuZG5zY3J5LnB0


## dnscry.pt-coeurdalene-ipv6

DNSCry.pt Coeur d'Alene - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAGlsyNjAyOmZlNTQ6MjI6NTc6OjViZDoxMzRdICHQMITAMQ3l-gX8OJxKvrK5pvC3PYdyRLEuV8ezcHk9GTIuZG5zY3J5cHQtY2VydC5kbnNjcnkucHQ


## dnscry.pt-copenhagen-ipv4

DNSCry.pt Copenhagen - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAACzg1LjEyMC44NC41IAQtkx4Ql90wruZmjw-U_klFVM9Mn7MHWgMR_Db140SrGTIuZG5zY3J5cHQtY2VydC5kbnNjcnkucHQ


## dnscry.pt-copenhagen-ipv6

DNSCry.pt Copenhagen - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAHlsyYTBlOjFkODA6MzE6OGE1NjowOmIwZTo1ZTowXSAELZMeEJfdMK7mZo8PlP5JRVTPTJ-zB1oDEfw29eNEqxkyLmRuc2NyeXB0LWNlcnQuZG5zY3J5LnB0


## dnscry.pt-coventry-ipv4

DNSCry.pt Coventry - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADzEwNC4xMjguMTkwLjEwOCAaJ1Ca-Hx6A91RK7871Z_-pdiQX85eaKtbbNCZRw9Z_hkyLmRuc2NyeXB0LWNlcnQuZG5zY3J5LnB0


## dnscry.pt-coventry-ipv6

DNSCry.pt Coventry - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAIFsyYTBjOmI4NDA6MjoxNjI6MTgwODowOjFjOjliNmNdIBonUJr4fHoD3VErvzvVn_6l2JBfzl5oq1ts0JlHD1n-GTIuZG5zY3J5cHQtY2VydC5kbnNjcnkucHQ


## dnscry.pt-dallas-ipv4

DNSCry.pt Dallas - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADTIzLjIzMC4yNTMuOTgg1OKRDMWAtnBoieTPNbjK-OrVjcuML2vQMc6gh-ZmYpAZMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeS5wdA


## dnscry.pt-dallas-ipv6

DNSCry.pt Dallas - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAE1syNjAyOmZiOTQ6MTozOTo6YV0g1OKRDMWAtnBoieTPNbjK-OrVjcuML2vQMc6gh-ZmYpAZMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeS5wdA


## dnscry.pt-denver-ipv4

DNSCry.pt Denver - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADTE5My44LjE3Mi4yNDggyvGozs1x9kpQ32QdDvmYikSft35xdA_amTUnAX5dI3gZMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeS5wdA


## dnscry.pt-denver-ipv6

DNSCry.pt Denver - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAGlsyNjA2OjY2ODA6NjoxOjozZWE5OjNjZTZdIMrxqM7NcfZKUN9kHQ75mIpEn7d-cXQP2pk1JwF-XSN4GTIuZG5zY3J5cHQtY2VydC5kbnNjcnkucHQ


## dnscry.pt-detroit-ipv4

DNSCry.pt Detroit - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADDY2LjE4Ny43LjE0MCBpn2OKcwbE01MLSkSXcaPKLf8IOmKbuE9GGZvAOBwaNRkyLmRuc2NyeXB0LWNlcnQuZG5zY3J5LnB0


## dnscry.pt-detroit-ipv6

DNSCry.pt Detroit - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAI1syNjA2OjY1YzA6NDA6NDo1ZjM6NTRjNDo4ZDEwOjliOThdIGmfY4pzBsTTUwtKRJdxo8ot_wg6Ypu4T0YZm8A4HBo1GTIuZG5zY3J5cHQtY2VydC5kbnNjcnkucHQ


## dnscry.pt-dublin-ipv4

DNSCry.pt Dublin - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADTE5NC4yNi4yMTMuMTUgEzWgsAQfbmA1ppXryEJ6vQ3Vvc2Kk2oRkdjodTEYvPQZMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeS5wdA


## dnscry.pt-dublin-ipv6

DNSCry.pt Dublin - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAFVsyYTA5OmNkNDY6Zjo0MjllOjo1XSATNaCwBB9uYDWmlevIQnq9DdW9zYqTahGR2Oh1MRi89BkyLmRuc2NyeXB0LWNlcnQuZG5zY3J5LnB0


## dnscry.pt-durham-ipv4

DNSCry.pt Durham - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADDM4LjQ1LjY0LjExNyAS3jjOGrb2p9i5bpMiO0WB-XlTLq7Ek3soP2xndELQ8xkyLmRuc2NyeXB0LWNlcnQuZG5zY3J5LnB0


## dnscry.pt-dusseldorf-ipv4

DNSCry.pt Düsseldorf - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADjE4NS4yNDQuMjcuMTM2IG5RCKZnWcBIWwMJ9wfdIkLhWRuNCczv-aVchrqwIzAmGTIuZG5zY3J5cHQtY2VydC5kbnNjcnkucHQ


## dnscry.pt-dusseldorf-ipv6

DNSCry.pt Düsseldorf - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAGFsyYTBmOjU3MDc6YWE4MTo1ZTNjOjoxXSBuUQimZ1nASFsDCfcH3SJC4VkbjQnM7_mlXIa6sCMwJhkyLmRuc2NyeXB0LWNlcnQuZG5zY3J5LnB0


## dnscry.pt-ebene-ipv4

DNSCry.pt Ebène - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADzEwMi4yMjIuMTA2LjE2NSBgatfPF03NfypARpsP4HbLOQbtu1RAh97lsjEVJVhR9BkyLmRuc2NyeXB0LWNlcnQuZG5zY3J5LnB0


## dnscry.pt-eygelshoven-ipv4

DNSCry.pt Eygelshoven - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADDkzLjk1LjExNS4yMSDit1FyUAiu0W-x936EJIC1keajbwu1pvb6yVKVj0KVYhkyLmRuc2NyeXB0LWNlcnQuZG5zY3J5LnB0


## dnscry.pt-flint-ipv4

DNSCry.pt Flint - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADzE0Ny4xODkuMTQwLjEzNiCL7wgLXnE-35sDhXk5N1RNpUfWmM2aUBcMFlst7FPdnRkyLmRuc2NyeXB0LWNlcnQuZG5zY3J5LnB0


## dnscry.pt-flint-ipv6

DNSCry.pt Flint - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAG1syNjA2OjY2ODA6Mjk6MTo6NTg1OTphMzdiXSCL7wgLXnE-35sDhXk5N1RNpUfWmM2aUBcMFlst7FPdnRkyLmRuc2NyeXB0LWNlcnQuZG5zY3J5LnB0


## dnscry.pt-frankfurt-ipv4

DNSCry.pt Frankfurt - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADTE5NC41MC4xOS4xNTAg-_TD5LiJYj-861zIGFSucHEg_7IT-3T3x8fYWhWrsekZMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeS5wdA


## dnscry.pt-frankfurt-ipv6

DNSCry.pt Frankfurt - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAI1syYTBmOjU3MDc6YWI4MDozMzRlOjI6MjoyY2QyOmE4YmNdIPv0w-S4iWI_vOtcyBhUrnBxIP-yE_t098fH2FoVq7HpGTIuZG5zY3J5cHQtY2VydC5kbnNjcnkucHQ


## dnscry.pt-fremont-ipv4

DNSCry.pt Fremont - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADDIzLjEzNC44OC43MSBywAheyGUw9-6xRRd9vJpQ9_so9hVXnOc-T50u1HjNlxkyLmRuc2NyeXB0LWNlcnQuZG5zY3J5LnB0


## dnscry.pt-fremont-ipv6

DNSCry.pt Fremont - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAFVsyNjAyOmZiYTE6MTAwOjo3MToxXSBywAheyGUw9-6xRRd9vJpQ9_so9hVXnOc-T50u1HjNlxkyLmRuc2NyeXB0LWNlcnQuZG5zY3J5LnB0


## dnscry.pt-fujairah-ipv4

DNSCry.pt Fujairah - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADTg5LjM2LjE2Mi4xODcgoviGbbki583h_iXBE6u-lLdJjyIsScLTYL107SA5Y0MZMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeS5wdA


## dnscry.pt-fujairah-ipv6

DNSCry.pt Fujairah - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAKFsyYTA2OmY5MDI6NDAwMToxMDA6OTAwMDo5MDAwOjM5YTQ6NWZlYl0goviGbbki583h_iXBE6u-lLdJjyIsScLTYL107SA5Y0MZMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeS5wdA


## dnscry.pt-geneva-ipv4

DNSCry.pt Geneva - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADDQ1LjkwLjU5LjE5MyApCKLNC-QxtyiyCC4AQIb36KxxFcalmSGG9V_CLDDyVxkyLmRuc2NyeXB0LWNlcnQuZG5zY3J5LnB0


## dnscry.pt-geneva-ipv6

DNSCry.pt Geneva - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAEFsyYTA1Ojk0MDY6OmFlMV0gKQiizQvkMbcosgguAECG9-iscRXGpZkhhvVfwiww8lcZMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeS5wdA


## dnscry.pt-hanoi-ipv4

DNSCry.pt Hanoi - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADTEwMy4xOTkuMTYuOTMg6iF-oJet7zyL2odP--IayA5Wrz6t94RPc7PXF53V82cZMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeS5wdA


## dnscry.pt-hanoi-ipv6

DNSCry.pt Hanoi - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAGFsyNDA0OmZiYzA6MDoxMWM4OjphMzI0XSDqIX6gl63vPIvah0_74hrIDlavPq33hE9zs9cXndXzZxkyLmRuc2NyeXB0LWNlcnQuZG5zY3J5LnB0


## dnscry.pt-helsinki-ipv4

DNSCry.pt Helsinki - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADTE4NS4yNTMuMTExLjYg0dTSOAcqwZKuSUDc2NFSGMgbgX-h0mEtcc_L2gCIcpYZMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeS5wdA


## dnscry.pt-helsinki-ipv6

DNSCry.pt Helsinki - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAElsyYTEyOjFmYzA6MDpjOjozXSDR1NI4ByrBkq5JQNzY0VIYyBuBf6HSYS1xz8vaAIhylhkyLmRuc2NyeXB0LWNlcnQuZG5zY3J5LnB0


## dnscry.pt-hochiminhcity-ipv4

DNSCry.pt Ho-Chi-Minh City - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADTE1Ny4yMC44My4xMzUgjO-Llfy6TK7yRAZaDd6xpzGDj51mcPeGAttNfFcI0ZYZMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeS5wdA


## dnscry.pt-hochiminhcity-ipv6

DNSCry.pt Ho-Chi-Minh City - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAGFsyNDAwOjZlYTA6MDoxMWFlOjphZGM0XSCM74uV_LpMrvJEBloN3rGnMYOPnWZw94YC2018VwjRlhkyLmRuc2NyeXB0LWNlcnQuZG5zY3J5LnB0


## dnscry.pt-hongkong-ipv4

DNSCry.pt Hong Kong - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADjQ1LjEyMy4xODguMTI5IAtxIfzDy0d11GLJHr7HPkdtPwGbimmNUM0gUa0gfjHIGTIuZG5zY3J5cHQtY2VydC5kbnNjcnkucHQ


## dnscry.pt-hongkong-ipv6

DNSCry.pt Hong Kong - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAF1syNDA2OjQzMDA6YmFlOjZiMDg6OjFdIAtxIfzDy0d11GLJHr7HPkdtPwGbimmNUM0gUa0gfjHIGTIuZG5zY3J5cHQtY2VydC5kbnNjcnkucHQ


## dnscry.pt-istanbul-ipv4

DNSCry.pt Istanbul - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADTE5My4yMjguMS4xMzAgXK0o7fAtGbbqsMWn6eFc8pQ3_Tf3YrD64C9ZuBjV-l0ZMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeS5wdA


## dnscry.pt-istanbul-ipv6

DNSCry.pt Istanbul - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAF1syYTEyOmUzNDI6MjAwOjoyOjE4MTldIFytKO3wLRm26rDFp-nhXPKUN_0392Kw-uAvWbgY1fpdGTIuZG5zY3J5cHQtY2VydC5kbnNjcnkucHQ


## dnscry.pt-jakarta-ipv4

DNSCry.pt Jakarta - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADTEwMy43Ni4xMjkuOTQgyn6S3ZBNV4fF8e7OwGbxj6P8udNOV2Epsd8DrUohOx8ZMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeS5wdA


## dnscry.pt-jakarta-ipv6

DNSCry.pt Jakarta - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAIFsyNDA3OjZhYzA6Mzo1OjEyMzQ6ZTM0ZTo3MmU0OjFdIMp-kt2QTVeHxfHuzsBm8Y-j_LnTTldhKbHfA61KITsfGTIuZG5zY3J5cHQtY2VydC5kbnNjcnkucHQ


## dnscry.pt-johannesburg-ipv4

DNSCry.pt Johannesburg - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADzE2OS4yMzkuMTI4LjEyNCDPBt-20rnrKqM3G3-ZKudPSvU9-zClzYY5-F2KRJSgsBkyLmRuc2NyeXB0LWNlcnQuZG5zY3J5LnB0


## dnscry.pt-johannesburg-ipv6

DNSCry.pt Johannesburg - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAFFsyYzBmOmY1MzA6OmQwMDoxODhdIM8G37bSuesqozcbf5kq509K9T37MKXNhjn4XYpElKCwGTIuZG5zY3J5cHQtY2VydC5kbnNjcnkucHQ


## dnscry.pt-kansascity-ipv4

DNSCry.pt Kansas City - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADTIzLjE1MC40MC4xMjEgQprQrFLF3Y2975ylDjnD8kdKAJLUvauubVrBGueEkcgZMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeS5wdA


## dnscry.pt-kansascity-ipv6

DNSCry.pt Kansas City - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAGVsyNjAyOjJiNzpkMDE6YzI5NTo6YjoxOF0gQprQrFLF3Y2975ylDjnD8kdKAJLUvauubVrBGueEkcgZMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeS5wdA


## dnscry.pt-kharkiv-ipv4

DNSCry.pt Kharkiv - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADjE5My4yMzguMTUzLjE3IG2aHSWgTHM0_Uy0coL03hWlRhQ23saGZAQFPzr3GQklGTIuZG5zY3J5cHQtY2VydC5kbnNjcnkucHQ


## dnscry.pt-kharkiv-ipv6

DNSCry.pt Kharkiv - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAFFsyYTAyOjI3YTg6ZmVlZDo6ODFdIG2aHSWgTHM0_Uy0coL03hWlRhQ23saGZAQFPzr3GQklGTIuZG5zY3J5cHQtY2VydC5kbnNjcnkucHQ


## dnscry.pt-kyiv-ipv4

DNSCry.pt Kyiv - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADTIxNy4xMi4yMjEuNjEgskgLubDTWs4bK9zH1IXKRYSylrG8XVPGWMJpUM37vwUZMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeS5wdA


## dnscry.pt-kyiv-ipv6

DNSCry.pt Kyiv - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAEFsyYTAyOjI3YWQ6OjIwMV0gskgLubDTWs4bK9zH1IXKRYSylrG8XVPGWMJpUM37vwUZMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeS5wdA


## dnscry.pt-lagos-ipv4

DNSCry.pt Lagos - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADTE3Ni45Ny4xOTIuMTIgcCgpSUHINZEdZRhbgwLZOUR6fOPJU5L4bY9g88TbNusZMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeS5wdA


## dnscry.pt-lagos-ipv6

DNSCry.pt Lagos - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAH1syYTA2OmY5MDE6NDAwMToxMDA6OjJkNmM6NzM2YV0gcCgpSUHINZEdZRhbgwLZOUR6fOPJU5L4bY9g88TbNusZMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeS5wdA


## dnscry.pt-lasvegas-ipv4

DNSCry.pt Las Vegas - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADTIwOS4xNDEuNDUuMjcgxObWYoxN9G0beY5ta20qYDsWjcrgoJsnpi7ILY0M9C4ZMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeS5wdA


## dnscry.pt-lasvegas-ipv6

DNSCry.pt Las Vegas - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAJVsyNjA1OjY0MDA6MjA6MjI1ODo3YWNiOjkxZmY6MjA5ODphOV0gxObWYoxN9G0beY5ta20qYDsWjcrgoJsnpi7ILY0M9C4ZMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeS5wdA


## dnscry.pt-libertylake-ipv4

DNSCry.pt Liberty Lake - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADDIzLjE4NC40OC4xOSCwg3q2XK6z70eHJhi0H7whWQ_ZWQylhMItvqKpd9GtzRkyLmRuc2NyeXB0LWNlcnQuZG5zY3J5LnB0


## dnscry.pt-libertylake-ipv6

DNSCry.pt Liberty Lake - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAGFsyNjAyOmZjMjQ6MTg6MzNmMjo6YWIxXSCwg3q2XK6z70eHJhi0H7whWQ_ZWQylhMItvqKpd9GtzRkyLmRuc2NyeXB0LWNlcnQuZG5zY3J5LnB0


## dnscry.pt-lima-ipv4

DNSCry.pt Lima - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADjE2MS4xMzIuNDcuMTg1IBzvQz4ky30VaNL_Y911uLyP-k1AgJGH5ryXuMlLPwd1GTIuZG5zY3J5cHQtY2VydC5kbnNjcnkucHQ


## dnscry.pt-lima-ipv6

DNSCry.pt Lima - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAFlsyMDAxOjQ3MDoxZjJhOjFkZTo6Ml0gHO9DPiTLfRVo0v9j3XW4vI_6TUCAkYfmvJe4yUs_B3UZMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeS5wdA


## dnscry.pt-london-ipv4

DNSCry.pt London - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADzE3OC4yMzkuMTc0LjI0NCCPZtxEvrtixgzqLZkrkl_-HL7-Cau2YUCEF2vb8sox7hkyLmRuc2NyeXB0LWNlcnQuZG5zY3J5LnB0


## dnscry.pt-london-ipv6

DNSCry.pt London - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAFFsyYTA1OjQxNDA6NzAwOmU6OmFdII9m3ES-u2LGDOotmSuSX_4cvv4Jq7ZhQIQXa9vyyjHuGTIuZG5zY3J5cHQtY2VydC5kbnNjcnkucHQ


## dnscry.pt-losangeles-ipv4

DNSCry.pt Los Angeles - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADjEwNC4yMDAuNjcuMTk0IIhxeSuGQHwchZdstQqcoKD_RAuV4w8Qr_1XmXFZucGEGTIuZG5zY3J5cHQtY2VydC5kbnNjcnkucHQ


## dnscry.pt-losangeles-ipv6

DNSCry.pt Los Angeles - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAF1syNjAyOmZmNzU6NzpiNzk6OmI0YjRdIIhxeSuGQHwchZdstQqcoKD_RAuV4w8Qr_1XmXFZucGEGTIuZG5zY3J5cHQtY2VydC5kbnNjcnkucHQ


## dnscry.pt-madrid-ipv4

DNSCry.pt Madrid - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADTUuMTM0LjExOC4xOTggF4pp6ab33hO4Nb9tp8zuU8Drkh2GcvzYZikut4DIHN8ZMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeS5wdA


## dnscry.pt-madrid-ipv6

DNSCry.pt Madrid - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAG1syYTAzOmM3YzA6NTI6MjY0MToxODA6OjEzXSAXimnppvfeE7g1v22nzO5TwOuSHYZy_NhmKS63gMgc3xkyLmRuc2NyeXB0LWNlcnQuZG5zY3J5LnB0


## dnscry.pt-miami-ipv4

DNSCry.pt Miami - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAACzg0LjMzLjE0LjEwIH3mGbFy3GVlTqL3Gxb0NAaMycTt6QbjaAogYiEsmf8mGTIuZG5zY3J5cHQtY2VydC5kbnNjcnkucHQ


## dnscry.pt-miami-ipv6

DNSCry.pt Miami - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAFlsyYTBjOjhmYzM6ODAwMjo6MjIxNl0gfeYZsXLcZWVOovcbFvQ0BozJxO3pBuNoCiBiISyZ_yYZMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeS5wdA


## dnscry.pt-milan-ipv4

DNSCry.pt Milan - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADTIxNy4xNTYuNTAuMjUgUjkX7nXh6Lfn1wuvsca7nm85e_I5RbNqLBNHkJU2CPgZMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeS5wdA


## dnscry.pt-milan-ipv6

DNSCry.pt Milan - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAFlsyYTBlOjFkODA6MjE6OWNjMjo6MV0gUjkX7nXh6Lfn1wuvsca7nm85e_I5RbNqLBNHkJU2CPgZMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeS5wdA


## dnscry.pt-mumbai-ipv4

DNSCry.pt Mumbai - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADjEwMy4xMTEuMTE0LjI1IENdCfc5GHRGIG-JtMeIw2cVTN1nHG4kan0vc_aonHWDGTIuZG5zY3J5cHQtY2VydC5kbnNjcnkucHQ


## dnscry.pt-mumbai-ipv6

DNSCry.pt Mumbai - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAH1syYTA2OmY5MDI6ODAwMToxMDA6OjE3NTc6ZTYxN10gQ10J9zkYdEYgb4m0x4jDZxVM3WccbiRqfS9z9qicdYMZMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeS5wdA


## dnscry.pt-munich-ipv4

DNSCry.pt Munich - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADTE5NC4zOS4yMDUuMTAgQtC7u79NGEO2MGscsRWQJwJZy8mvvDwc1gpY_VjEf2IZMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeS5wdA


## dnscry.pt-munich-ipv6

DNSCry.pt Munich - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAGlsyYTBjOjhmYzA6MTc0OTo2NjoxODo6MTZdIELQu7u_TRhDtjBrHLEVkCcCWcvJr7w8HNYKWP1YxH9iGTIuZG5zY3J5cHQtY2VydC5kbnNjcnkucHQ


## dnscry.pt-naaldwijk-ipv4

DNSCry.pt Naaldwijk - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADTIzLjEzNy4yNDkuMjYgCA4-g3tus39pqm78_CoOc8byRBbLfuc5ceEiFNFWnN4ZMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeS5wdA


## dnscry.pt-naaldwijk-ipv6

DNSCry.pt Naaldwijk - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAGFsyNjAyOmZjMjQ6MTI6OTg3Mzo6YWIxXSAIDj6De26zf2mqbvz8Kg5zxvJEFst-5zlx4SIU0Vac3hkyLmRuc2NyeXB0LWNlcnQuZG5zY3J5LnB0


## dnscry.pt-newyork-ipv4

DNSCry.pt New York - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADDg0LjMzLjI0NS4xMCArmwOBLVc12QqaM0G2TykZIeHlqQpPWlK8YEWtW14L0xkyLmRuc2NyeXB0LWNlcnQuZG5zY3J5LnB0


## dnscry.pt-newyork-ipv6

DNSCry.pt New York - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAF1syYTBjOjhmYzM6NjQwMjo6MTo5ODRdICubA4EtVzXZCpozQbZPKRkh4eWpCk9aUrxgRa1bXgvTGTIuZG5zY3J5cHQtY2VydC5kbnNjcnkucHQ


## dnscry.pt-nuremberg-ipv4

DNSCry.pt Nuremberg - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADTIwMi42MS4yMzYuNjcgr2UzWGeubsFSZXP-_a8P2GA-gsZJ81sKZuhdsgsGqscZMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeS5wdA


## dnscry.pt-nuremberg-ipv6

DNSCry.pt Nuremberg - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAJVsyYTAzOjQwMDA6NWM6NTE6MjRiOTo1MWZmOmZlODA6ZjNhN10gr2UzWGeubsFSZXP-_a8P2GA-gsZJ81sKZuhdsgsGqscZMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeS5wdA


## dnscry.pt-oradea-ipv4

DNSCry.pt Oradea - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADTE5My4zMi44Ny4xMjcgqGZV4dUptaXgY6MX4VJhhQfNcTb8Kd0AVhegLHQzJogZMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeS5wdA


## dnscry.pt-oradea-ipv6

DNSCry.pt Oradea - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAHVsyYTBkOjgxNDA6MDoxMzoyOTE1OmFmOjA6MThdIKhmVeHVKbWl4GOjF-FSYYUHzXE2_CndAFYXoCx0MyaIGTIuZG5zY3J5cHQtY2VydC5kbnNjcnkucHQ


## dnscry.pt-ottoville-ipv4

DNSCry.pt Ottoville - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADzEwNC4yMzQuMjMxLjIzOSBVJyZb_D0SazeybnfWj5DWZ8NUgxii-zg9r-N8VNSWtBkyLmRuc2NyeXB0LWNlcnQuZG5zY3J5LnB0


## dnscry.pt-ottoville-ipv6

DNSCry.pt Ottoville - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAE1syNjAyOmY5NTM6NjoyNTo6YV0gVScmW_w9Ems3sm531o-Q1mfDVIMYovs4Pa_jfFTUlrQZMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeS5wdA


## dnscry.pt-paris-ipv4

DNSCry.pt Paris - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAACzg5LjExNy4yLjE3IAXdC7hGEegKD86br-tVRwZTcJfJZAEFjW4jCV5lzdutGTIuZG5zY3J5cHQtY2VydC5kbnNjcnkucHQ


## dnscry.pt-paris-ipv6

DNSCry.pt Paris - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAHlsyNDAyOmQwYzA6MjI6NmNkMDo0OjQ6NDo1YjgxXSAF3Qu4RhHoCg_Om6_rVUcGU3CXyWQBBY1uIwleZc3brRkyLmRuc2NyeXB0LWNlcnQuZG5zY3J5LnB0


## dnscry.pt-philadelphia-ipv4

DNSCry.pt Philadelphia - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADTE1NC4xNi4xNTkuMjIg2_tLIEpyMKwEhbD7PirfNwPUvZUnTM4z8F8DVkeQI3oZMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeS5wdA


## dnscry.pt-philadelphia-ipv6

DNSCry.pt Philadelphia - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAFVsyNjA0OmJmMDA6MjEwOjEyOjoyXSDb-0sgSnIwrASFsPs-Kt83A9S9lSdMzjPwXwNWR5AjehkyLmRuc2NyeXB0LWNlcnQuZG5zY3J5LnB0


## dnscry.pt-phoenix-ipv4

DNSCry.pt Phoenix - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADjE3My4yNDkuMjAzLjUyIC7cQyu7dTkvrH-MdwA3fzAkI_2dEUnZd3x1ON0tfYubGTIuZG5zY3J5cHQtY2VydC5kbnNjcnkucHQ


## dnscry.pt-phoenix-ipv6

DNSCry.pt Phoenix - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAG1syNjA3OjFlNDA6MToxMGE0OjoxOTpjYTg0XSAu3EMru3U5L6x_jHcAN38wJCP9nRFJ2Xd8dTjdLX2LmxkyLmRuc2NyeXB0LWNlcnQuZG5zY3J5LnB0


## dnscry.pt-portedwards-ipv4

DNSCry.pt Port Edwards - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADzE3Ni4xMTEuMjE5LjEyNiDzuja5nmAyDvA5jakqkuLQEtb245xsAhNwJYDLkKraKhkyLmRuc2NyeXB0LWNlcnQuZG5zY3J5LnB0


## dnscry.pt-portedwards-ipv6

DNSCry.pt Port Edwards - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAGFsyMDAxOjQ3MDoxZjExOjJiYjo6YjIzXSDzuja5nmAyDvA5jakqkuLQEtb245xsAhNwJYDLkKraKhkyLmRuc2NyeXB0LWNlcnQuZG5zY3J5LnB0


## dnscry.pt-portland-ipv4

DNSCry.pt Portland - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADzEwMy4xMjQuMTA2LjIzMyCN5S36eWstGFliH6xl8Mg2gyF99cqzMzgoJfAtWVYJnhkyLmRuc2NyeXB0LWNlcnQuZG5zY3J5LnB0


## dnscry.pt-portland-ipv6

DNSCry.pt Portland - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAIVsyNDAyOmQwYzA6MTY6YTFlNjowOmI4OTM6YmY3OmRkXSCN5S36eWstGFliH6xl8Mg2gyF99cqzMzgoJfAtWVYJnhkyLmRuc2NyeXB0LWNlcnQuZG5zY3J5LnB0


## dnscry.pt-riga-ipv4

DNSCry.pt Riga - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADzE5NS4xMjMuMjEyLjIwMCCKpSwU2DoHr1tktJRs4UIiLfoXBly8F7WmgX74sIHRyhkyLmRuc2NyeXB0LWNlcnQuZG5zY3J5LnB0


## dnscry.pt-riga-ipv6

DNSCry.pt Riga - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAEVsyYTAyOjI3YWM6OjEyNDldIIqlLBTYOgevW2S0lGzhQiIt-hcGXLwXtaaBfviwgdHKGTIuZG5zY3J5cHQtY2VydC5kbnNjcnkucHQ


## dnscry.pt-saltlakecity-ipv4

DNSCry.pt Salt Lake City - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADjEwMy4xMTQuMTYyLjY1IKbTxlVrc12BNolzMCksgqjW75nTqlnHp95UlrGWqm-UGTIuZG5zY3J5cHQtY2VydC5kbnNjcnkucHQ


## dnscry.pt-saltlakecity-ipv6

DNSCry.pt Salt Lake City - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAIVsyNDAyOmQwYzA6MTg6YzhmZjowOmI4OTM6YmY3OmRkXSCm08ZVa3NdgTaJczApLIKo1u-Z06pZx6feVJaxlqpvlBkyLmRuc2NyeXB0LWNlcnQuZG5zY3J5LnB0


## dnscry.pt-sandefjord-ipv4

DNSCry.pt Sandefjord - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADTE5NC4zMi4xMDcuNDggXTsyJ8l_6LJ4TCwKbGyVeIVM1yLzf8sxL2PmKjZIMvcZMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeS5wdA


## dnscry.pt-santaclara-ipv4

DNSCry.pt Santa Clara - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADzE3Ni4xMTEuMjIzLjE2NyCmqAI-1fpR1qtHZyAx3vJJ7SpKXkdmPAnZZ5ga25JckxkyLmRuc2NyeXB0LWNlcnQuZG5zY3J5LnB0


## dnscry.pt-santaclara-ipv6

DNSCry.pt Santa Clara - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAG1syNjA2OjY2ODA6MzU6MTo6NTA2ZDo4Y2UyXSCmqAI-1fpR1qtHZyAx3vJJ7SpKXkdmPAnZZ5ga25JckxkyLmRuc2NyeXB0LWNlcnQuZG5zY3J5LnB0


## dnscry.pt-saopaulo-ipv4

DNSCry.pt Sao Paulo - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADjEwOC4xODEuNjkuMTUzIKai-Qjyp6DgYnQVy1gEzvb3-NTklTiCmy4Afgv7TRJVGTIuZG5zY3J5cHQtY2VydC5kbnNjcnkucHQ


## dnscry.pt-saopaulo-ipv6

DNSCry.pt Sao Paulo - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAHlsyNjA0OjY2MDA6ZmQwMDo5MDo6MWI4YjozYTNjXSCmovkI8qeg4GJ0FctYBM729_jU5JU4gpsuAH4L-00SVRkyLmRuc2NyeXB0LWNlcnQuZG5zY3J5LnB0


## dnscry.pt-seattle-ipv4

DNSCry.pt Seattle - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADTQ1LjE0LjExNS4xMjUgSTZOlEDEMu3NY_DOCJbHzIlEZNFbrckVk2yl5PMnThMZMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeS5wdA


## dnscry.pt-seattle-ipv6

DNSCry.pt Seattle - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAG1syNjA2OjY2ODA6MTk6MTo6NGZiNDo3MWE3XSBJNk6UQMQy7c1j8M4IlsfMiURk0VutyRWTbKXk8ydOExkyLmRuc2NyeXB0LWNlcnQuZG5zY3J5LnB0


## dnscry.pt-singapore-ipv4

DNSCry.pt Singapore - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADTIzLjI3LjEwMS4xOTEgX4DtgHar8KftBn3jpccW4O8BQ2VFctwhwu_rMfmCWsYZMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeS5wdA


## dnscry.pt-singapore-ipv6

DNSCry.pt Singapore - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAFVsyNjA2OmZjNDA6NDAwMzpmOjphXSBfgO2Adqvwp-0GfeOlxxbg7wFDZUVy3CHC7-sx-YJaxhkyLmRuc2NyeXB0LWNlcnQuZG5zY3J5LnB0


## dnscry.pt-singapore02-ipv4

DNSCry.pt Singapore 02 - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADTEwMy4xNzkuNDQuNzMgICxK5c5XamgK_BNMTtSKyEnZM4D44NPAIHddngTPbGUZMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeS5wdA


## dnscry.pt-singapore02-ipv6

DNSCry.pt Singapore 02 - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAF1syNDAzOmNmYzA6MTExNDoxMGU6OmFdICAsSuXOV2poCvwTTE7UishJ2TOA-ODTwCB3XZ4Ez2xlGTIuZG5zY3J5cHQtY2VydC5kbnNjcnkucHQ


## dnscry.pt-sofia-ipv4

DNSCry.pt Sofia - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAACzc5LjEyNC43Ny4zIGjOJralcFGh38dFov6MP6OkkaSPIlSCbku5I7J2NZUfGTIuZG5zY3J5cHQtY2VydC5kbnNjcnkucHQ


## dnscry.pt-sofia-ipv6

DNSCry.pt Sofia - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAFlsyYTAxOjg3NDA6MTo0MDo6OGEyNV0gaM4mtqVwUaHfx0Wi_ow_o6SRpI8iVIJuS7kjsnY1lR8ZMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeS5wdA


## dnscry.pt-spokane-ipv4

DNSCry.pt Spokane - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADTEwNC4zNi44Ni4xODEg_ifyAp41KOphKBVIwROBjWV91n9fuUzlzUqXCIklST0ZMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeS5wdA


## dnscry.pt-spokane-ipv6

DNSCry.pt Spokane - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAFFsyNjA2OmE4YzA6MzoyMDI6OmFdIP4n8gKeNSjqYSgVSMETgY1lfdZ_X7lM5c1KlwiJJUk9GTIuZG5zY3J5cHQtY2VydC5kbnNjcnkucHQ


## dnscry.pt-stockholm-ipv4

DNSCry.pt Stockholm - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADjk1LjE0My4xOTYuMTkwIIe6V4-SeKbsyNllxXsYRoqK7NDU9EtUn7yp48YWeEu9GTIuZG5zY3J5cHQtY2VydC5kbnNjcnkucHQ


## dnscry.pt-stockholm-ipv6

DNSCry.pt Stockholm - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAHFsyYTAzOmQ3ODA6MDoxOTY6OjNlODQ6NTZhZl0gh7pXj5J4puzI2WXFexhGiors0NT0S1SfvKnjxhZ4S70ZMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeS5wdA


## dnscry.pt-sydney-ipv4

DNSCry.pt Sydney - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADDg0LjMzLjE1LjEwMCBq8e6bhSwgentAeVRR__dhXfcSy86CtQPtq0vb_Cl18hkyLmRuc2NyeXB0LWNlcnQuZG5zY3J5LnB0


## dnscry.pt-sydney-ipv6

DNSCry.pt Sydney - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAH1syYTBjOjhmYzE6ODAwNDo1NTM6OjE0NWE6YmJmOV0gavHum4UsIHp7QHlUUf_3YV33EsvOgrUD7atL2_wpdfIZMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeS5wdA


## dnscry.pt-taipeh-ipv4

DNSCry.pt Taipeh - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADjEwMy4xMzEuMTg5LjExIIMLIy-_BnvJTc23i9iX0LlOgTzBwtumxbxntod8ri75GTIuZG5zY3J5cHQtY2VydC5kbnNjcnkucHQ


## dnscry.pt-taipeh-ipv6

DNSCry.pt Taipeh - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAGlsyNDAzOmNmYzA6MTAwNDozNjk6OjViMjFdIIMLIy-_BnvJTc23i9iX0LlOgTzBwtumxbxntod8ri75GTIuZG5zY3J5cHQtY2VydC5kbnNjcnkucHQ


## dnscry.pt-tallinn-ipv4

DNSCry.pt Tallinn - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADTE4NS4xOTQuNTMuMjIgr0WageGep9cjA5yYpY30Z6EsTYHZnSlV-PCfvZssTNcZMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeS5wdA


## dnscry.pt-tallinn-ipv6

DNSCry.pt Tallinn - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAElsyYTA0OjZmMDA6NDo6MTdhXSCvRZqB4Z6n1yMDnJiljfRnoSxNgdmdKVX48J-9myxM1xkyLmRuc2NyeXB0LWNlcnQuZG5zY3J5LnB0


## dnscry.pt-tampa-ipv4

DNSCry.pt Tampa - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADzE2NS4xNDAuMTE3LjI0OCBfK4fFWjW65PRF3_42MZM1Ly9t0ZLHdDA_0uy63rk0zBkyLmRuc2NyeXB0LWNlcnQuZG5zY3J5LnB0


## dnscry.pt-tampa-ipv6

DNSCry.pt Tampa - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAIlsyNjAyOmZjYzA6MjIyMjowOmZmMjQ6YTJjNzoxOWM6MV0gXyuHxVo1uuT0Rd_-NjGTNS8vbdGSx3QwP9Lsut65NMwZMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeS5wdA


## dnscry.pt-taos-ipv4

DNSCry.pt Taos - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADjYzLjEzMy4yMjMuMTM4IIggy47qNTs0s0PnLuDK5UcpSJt_t7XVmTr0tWn1QlqNGTIuZG5zY3J5cHQtY2VydC5kbnNjcnkucHQ


## dnscry.pt-taos-ipv6

DNSCry.pt Taos - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAG1syNjA2OjY2ODA6NTM6MTo6ODQ2YTpiZDc5XSCIIMuO6jU7NLND5y7gyuVHKUibf7e11Zk69LVp9UJajRkyLmRuc2NyeXB0LWNlcnQuZG5zY3J5LnB0


## dnscry.pt-tokyo02-ipv4

DNSCry.pt Tokyo 02 - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADDEwMy4xNzkuNDUuNiDfai5sp1im-BPHwbM1GCnTqn20FIbQfuvvybKsGf0pjhkyLmRuc2NyeXB0LWNlcnQuZG5zY3J5LnB0


## dnscry.pt-tokyo02-ipv6

DNSCry.pt Tokyo 02 - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAE1syYTBhOjYwNDA6OTczZDo6YV0g32oubKdYpvgTx8GzNRgp06p9tBSG0H7r78myrBn9KY4ZMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeS5wdA


## dnscry.pt-toronto-ipv4

DNSCry.pt Toronto - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADDIzLjEzNC44OS4yMyB4hEFNFQU627_7pPlQl-k6TkU1hGlpOEZe7eg7V1tODRkyLmRuc2NyeXB0LWNlcnQuZG5zY3J5LnB0


## dnscry.pt-toronto-ipv6

DNSCry.pt Toronto - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAFVsyNjAyOmZiYTE6ZDAwOjoyMzoxXSB4hEFNFQU627_7pPlQl-k6TkU1hGlpOEZe7eg7V1tODRkyLmRuc2NyeXB0LWNlcnQuZG5zY3J5LnB0


## dnscry.pt-valdivia-ipv4

DNSCry.pt Valdivia - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADTIxNi43My4xNTkuMjYgnpr1thxYT4SkWK38OEbiPOQa3NSVayBN7f8BkMVREC8ZMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeS5wdA


## dnscry.pt-valdivia-ipv6

DNSCry.pt Valdivia - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAFVsyYTA2OmEwMDY6ZDFkMTo6MTE2XSCemvW2HFhPhKRYrfw4RuI85Brc1JVrIE3t_wGQxVEQLxkyLmRuc2NyeXB0LWNlcnQuZG5zY3J5LnB0


## dnscry.pt-vancouver-ipv4

DNSCry.pt Vancouver - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADTEwNC4zNi4xNDguNDYgtfMJd8tAE6dqf5ttVuNplwXDILR4cdhJpktOYiBCEoAZMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeS5wdA


## dnscry.pt-vancouver-ipv6

DNSCry.pt Vancouver - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAIFsyNjA3OjdiMDA6MzAwNDpmZmZmOjphNjhkOjVhMmVdILXzCXfLQBOnan-bbVbjaZcFwyC0eHHYSaZLTmIgQhKAGTIuZG5zY3J5cHQtY2VydC5kbnNjcnkucHQ


## dnscry.pt-vienna-ipv4

DNSCry.pt Vienna - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADTgzLjEzOC41NS4xODYg3kyI1rUYwQymzbrF1c5fYhw1rWmOTm8L6i1aISwm6y4ZMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeS5wdA


## dnscry.pt-vienna-ipv6

DNSCry.pt Vienna - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAHVsyYTBkOmYzMDI6MTEwOjY1MTc6OmJiNDoyMTRdIN5MiNa1GMEMps26xdXOX2IcNa1pjk5vC-otWiEsJusuGTIuZG5zY3J5cHQtY2VydC5kbnNjcnkucHQ


## dnscry.pt-vilnius-ipv4

DNSCry.pt Vilnius - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADTE2Mi4yNTQuODYuMTMg4nDDbNqRwkkkZWTJ5c82d1sbs0NeQCbn-aFldCI2mn4ZMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeS5wdA


## dnscry.pt-vilnius-ipv6

DNSCry.pt Vilnius - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAF1syYTEzOjk0MDE6MDoxOjozZDU4OjFdIOJww2zakcJJJGVkyeXPNndbG7NDXkAm5_mhZXQiNpp-GTIuZG5zY3J5cHQtY2VydC5kbnNjcnkucHQ


## dnscry.pt-warsaw02-ipv4

DNSCry.pt Warsaw 02 - DNSCrypt, no filter, no logs, DNSSEC support (IPv4 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAADjg4LjIxOC4yMDYuMTM3IE42BnReymaCMKWg_FWRMirGsBpqOOzlekDh8UwsfVEQGTIuZG5zY3J5cHQtY2VydC5kbnNjcnkucHQ


## dnscry.pt-warsaw02-ipv6

DNSCry.pt Warsaw 02 - DNSCrypt, no filter, no logs, DNSSEC support (IPv6 server)

https://www.dnscry.pt

sdns://AQcAAAAAAAAAFlsyYTA5OmIyODA6ZmUwMDoyNDo6YV0gTjYGdF7KZoIwpaD8VZEyKsawGmo47OV6QOHxTCx9URAZMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeS5wdA


## dnscrypt.ca-ipv4

Canadian based, unfiltered, DNSSEC validating, and no logs... for your pleasure. https://dnscrypt.ca/

sdns://AQcAAAAAAAAAEzE4NS4xMTEuMTg4LjQ2Ojg0NDMgC-tbTwd-08e_JtBJmgsvjAG9i10itE-LBNCwjTflezQiMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeXB0LmNhLTEtaXB2NA


## dnscrypt.ca-ipv4-doh

Canadian based, unfiltered, DNSSEC validating, and no logs... for your pleasure. https://dnscrypt.ca/

sdns://AgcAAAAAAAAADjE4NS4xMTEuMTg4LjQ2ID8EEe3pxEdwV9V-V4g7HyBbIM3A8yYxKbHuAmmiZ49jEGRuczEuZG5zY3J5cHQuY2EKL2Rucy1xdWVyeQ


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


## dnsfilter

A DNS service provider that offers a secure and private DNS solution. Logs, no DNSSEC, filters some malicious domains.
Homepage: https://www.dnsfilter.com/

sdns://AgAAAAAAAAAADjEwMy4yNDcuMzYuMTUwoMwQYNOcgym2K2-8fQ1t-TCYabmB5-Y5LVzY-kCPTYDmoPf1ryiAHod9ffOivij-FJ8ydKftKfE2_VA845jLqAsNoLNeBZUM-9gln5N1uhAYcLjDxMDsWlKXV-YxZ-neJqnooEROvWe7g_iAezkh6TiskXi4gr1QqtsRIx8ETPXwjffOoOZEumlj4zX-dly5l2sSsQ61QpS0JHd2TMs6OsyjrLL8ICquP7e_BeTIHEGU3KRFEdT5rzBHhuwa5yGECc9ioINVEWRvaC5kbnNmaWx0ZXIuY29tCi9kbnMtcXVlcnk


## dnsforfamily

(DNSCrypt Protocol) (Now supports DNSSEC). Block adult websites, gambling websites, malwares and advertisements.
It also enforces safe search in: Google, YouTube, Bing, DuckDuckGo and Yandex.

Social websites like Facebook and Instagram are not blocked. No DNS queries are logged.

As of 26-May-2022 5.9 million websites are blocked and new websites are added to blacklist daily.
Completely free, no ads or any commercial motive. Operating for 4 years now.

Warning: This server is incompatible with anonymization.

Provided by: https://dnsforfamily.com

sdns://AQMAAAAAAAAADDc4LjQ3LjY0LjE2MSATJeLOABXNSYcSJIoqR5_iUYz87Y4OecMLB84aEAKPrRBkbnNmb3JmYW1pbHkuY29t


## dnsforfamily-doh

(DoH Protocol) (Now supports DNSSEC). Block adult websites, gambling websites, malwares and advertisements.
It also enforces safe search in: Google, YouTube, Bing, DuckDuckGo and Yandex.

Social websites like Facebook and Instagram are not blocked. No DNS queries are logged.

As of 26-May-2022 5.9 million websites are blocked and new websites are added to blacklist daily.
Completely free, no ads or any commercial motive. Operating for 4 years now.

Provided by: https://dnsforfamily.com

sdns://AgMAAAAAAAAADzE2Ny4yMzUuMjM2LjEwN6DMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5qD39a8ogB6HfX3zor4o_hSfMnSn7SnxNv1QPOOYy6gLDaCzXgWVDPvYJZ-TdboQGHC4w8TA7FpSl1fmMWfp3iap6KBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33zqDmRLppY-M1_nZcuZdrErEOtUKUtCR3dkzLOjrMo6yy_CAqrj-3vwXkyBxBlNykRRHU-a8wR4bsGuchhAnPYqCDVRhkbnMtZG9oLmRuc2ZvcmZhbWlseS5jb20KL2Rucy1xdWVyeQ


## dnsforfamily-doh-no-safe-search

(DoH Protocol) (Now supports DNSSEC) Block adult websites, gambling websites, malwares and advertisements.
Unlike other dnsforfamily servers, this one does not enforces safe search. So Google, YouTube, Bing, DuckDuckGo and Yandex are completely accessible without any restriction.

Social websites like Facebook and Instagram are not blocked. No DNS queries are logged.

As of 26-May-2022 5.9 million websites are blocked and new websites are added to blacklist daily.
Completely free, no ads or any commercial motive. Operating for 4 years now.

Warning: This server is incompatible with anonymization.

Provided by: https://dnsforfamily.com

sdns://AgMAAAAAAAAADzE2Ny4yMzUuMjM2LjEwN6DMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5qD39a8ogB6HfX3zor4o_hSfMnSn7SnxNv1QPOOYy6gLDaCzXgWVDPvYJZ-TdboQGHC4w8TA7FpSl1fmMWfp3iap6KBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33zqDmRLppY-M1_nZcuZdrErEOtUKUtCR3dkzLOjrMo6yy_CAqrj-3vwXkyBxBlNykRRHU-a8wR4bsGuchhAnPYqCDVSdkbnMtZG9oLW5vLXNhZmUtc2VhcmNoLmRuc2ZvcmZhbWlseS5jb20KL2Rucy1xdWVyeQ


## dnsforfamily-no-safe-search

(DNSCrypt Protocol) (Now supports DNSSEC) Block adult websites, gambling websites, malwares and advertisements.
Unlike other dnsforfamily servers, this one does not enforces safe search. So Google, YouTube, Bing, DuckDuckGo and Yandex are completely accessible without any restriction.

Social websites like Facebook and Instagram are not blocked. No DNS queries are logged.

As of 26-May-2022 5.9 million websites are blocked and new websites are added to blacklist daily.
Completely free, no ads or any commercial motive. Operating for 4 years now.

Warning: This server is incompatible with anonymization.

Provided by: https://dnsforfamily.com

sdns://AQMAAAAAAAAADzEzNS4xODEuMTkzLjIyMiDrxcZ_hFtGE6tfATvQZYjxgl5pTY_e2cRH_ms8bEWofBBkbnNmb3JmYW1pbHkuY29t


## dnsforfamily-v6

(DNSCrypt Protocol) (Now supports DNSSEC) Block adult websites, gambling websites, malwares and advertisements.
It also enforces safe search in: Google, YouTube, Bing, DuckDuckGo and Yandex.

Social websites like Facebook and Instagram are not blocked. No DNS queries are logged.

As of 26-May-2022 5.9 million websites are blocked and new websites are added to blacklist daily.
Completely free, no ads or any commercial motive. Operating for 4 years now.

Provided by: https://dnsforfamily.com

sdns://AQMAAAAAAAAAF1syYTAxOjRmODoxYzE3OjRkZjg6OjFdIBMl4s4AFc1JhxIkiipHn-JRjPztjg55wwsHzhoQAo-tEGRuc2ZvcmZhbWlseS5jb20


## dnsforge.de

Public DoH resolver running with Pihole for Adblocking (https://dnsforge.de).

Non-logging, AD-filtering, supports DNSSEC. Hosted in Germany.

sdns://AgMAAAAAAAAADDE3Ni45LjkzLjE5OKDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5qD39a8ogB6HfX3zor4o_hSfMnSn7SnxNv1QPOOYy6gLDaCzXgWVDPvYJZ-TdboQGHC4w8TA7FpSl1fmMWfp3iap6KBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33zqDmRLppY-M1_nZcuZdrErEOtUKUtCR3dkzLOjrMo6yy_CAqrj-3vwXkyBxBlNykRRHU-a8wR4bsGuchhAnPYqCDVQtkbnNmb3JnZS5kZQovZG5zLXF1ZXJ5


## dnslow.me

dnslow.me is an open source project, also your advertisement and threat blocking, privacy-first, encrypted DNS.

All DNS requests will be protected with threat-intelligence feeds and randomly distributed to some other DNS resolvers.

More info on the [homepage](https://dnslow.me) and [GitHub](https://github.com/PeterDaveHello/dnslow.me)

sdns://AgAAAAAAAAAAAKDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5qD39a8ogB6HfX3zor4o_hSfMnSn7SnxNv1QPOOYy6gLDaCzXgWVDPvYJZ-TdboQGHC4w8TA7FpSl1fmMWfp3iap6KBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33zqDmRLppY-M1_nZcuZdrErEOtUKUtCR3dkzLOjrMo6yy_CAqrj-3vwXkyBxBlNykRRHU-a8wR4bsGuchhAnPYqCDVQlkbnNsb3cubWUKL2Rucy1xdWVyeQ


## dnspod

A public DNS resolver that supports DoH/DoT in mainland China, provided by dnspod/Tencent-cloud.
Homepage: https://dnspod.cn

Warning: GFW filtering rules are applied by this resolver.

sdns://AgAAAAAAAAAADDE2Mi4xNC4yMS41NgAHZG9oLnB1YgovZG5zLXF1ZXJ5
sdns://AgAAAAAAAAAADTE2Mi4xNC4yMS4xNzgAB2RvaC5wdWIKL2Rucy1xdWVyeQ
sdns://AgAAAAAAAAAADDEyMC41My41My41MwAMMTIwLjUzLjUzLjUzCi9kbnMtcXVlcnk
sdns://AgAAAAAAAAAACjEuMTIuMTIuMTIACjEuMTIuMTIuMTIKL2Rucy1xdWVyeQ
sdns://AgAAAAAAAAAACjEuMTIuMzQuNTYACjEuMTIuMTIuMTIKL2Rucy1xdWVyeQ


## dnswarden-adult-doh

Blocks adult content and enforces safe search on major search engines.
For further customization look here: https://dnswarden.com/customfilter.html

sdns://AgIAAAAAAAAAAKDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5qD39a8ogB6HfX3zor4o_hSfMnSn7SnxNv1QPOOYy6gLDaCzXgWVDPvYJZ-TdboQGHC4w8TA7FpSl1fmMWfp3iap6KBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33zqDmRLppY-M1_nZcuZdrErEOtUKUtCR3dkzLOjrMo6yy_CAqrj-3vwXkyBxBlNykRRHU-a8wR4bsGuchhAnPYqCDVRFkbnMuZG5zd2FyZGVuLmNvbQwvYWR1bHRmaWx0ZXI


## dnswarden-uncensor-dc-swiss

Hosted in switzerland.
For more information look at https://github.com/bhanupratapys/dnswarden or https://dnswarden.com

sdns://AQcAAAAAAAAAFDE4OC4yNDQuMTE3LjExNDoxNDQzIOK0J6XON4YjPmXjlonI5Lx0ZenB5Din7hrX-8uYlDxHKjIuZG5zY3J5cHQtY2VydC5kbnN3YXJkZW4tdW5jZW5zb3JlZC1zd2lzcw


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


## doh-cleanbrowsing-security

Block access to phishing, malware and malicious domains. It does not block adult content.
By https://cleanbrowsing.org/

sdns://AgMAAAAAAAAAAAAVZG9oLmNsZWFuYnJvd3Npbmcub3JnFS9kb2gvc2VjdXJpdHktZmlsdGVyLw


## doh-crypto-sx

DNS-over-HTTPS server. Anycast, no logs, no censorship, DNSSEC.
Globally cached via Cloudflare.
Maintained by Frank Denis.

sdns://AgcAAAAAAAAACzEwNC4yMS42Ljc4AA1kb2guY3J5cHRvLnN4Ci9kbnMtcXVlcnk
sdns://AgcAAAAAAAAADjE3Mi42Ny4xMzQuMTU3AA1kb2guY3J5cHRvLnN4Ci9kbnMtcXVlcnk


## doh-crypto-sx-ipv6

DNS-over-HTTPS server accessible over IPv6. Anycast, no logs, no censorship, DNSSEC.
Globally cached via Cloudflare.
Maintained by Frank Denis.

sdns://AgcAAAAAAAAAGlsyNjA2OjQ3MDA6MzAzNzo6NjgxNTo2NGVdABJkb2gtaXB2Ni5jcnlwdG8uc3gKL2Rucy1xdWVyeQ
sdns://AgcAAAAAAAAAG1syNjA2OjQ3MDA6MzAzNjo6YWM0Mzo4NjlkXQASZG9oLWlwdjYuY3J5cHRvLnN4Ci9kbnMtcXVlcnk


## doh-ibksturm

DoH & DoT Server, No Logging, No Filters, DNSSEC

Running privately by ibksturm in Thurgau, Switzerland

sdns://AgcAAAAAAAAADjIxMy4xOTYuMTkxLjk2IJRq0Q0bVAt-COHOCim4PL_Nyl_hsJ6bi7sOV67HMIVpGGlia3N0dXJtLnN5bm9sb2d5Lm1lOjQ0MwovZG5zLXF1ZXJ5


## doh.appliedprivacy.net

Public DoH resolver operated by the Foundation for Applied Privacy (https://appliedprivacy.net).
Hosted in Vienna, Austria.

Non-logging, non-filtering, supports DNSSEC.

sdns://AgcAAAAAAAAAAKDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5qD39a8ogB6HfX3zor4o_hSfMnSn7SnxNv1QPOOYy6gLDaCzXgWVDPvYJZ-TdboQGHC4w8TA7FpSl1fmMWfp3iap6KBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33zqDmRLppY-M1_nZcuZdrErEOtUKUtCR3dkzLOjrMo6yy_CAqrj-3vwXkyBxBlNykRRHU-a8wR4bsGuchhAnPYqCDVRZkb2guYXBwbGllZHByaXZhY3kubmV0Bi9xdWVyeQ


## doh.ffmuc.net

An open (non-logging, non-filtering, non-censoring) DoH resolver operated by Freifunk Munich with nodes in DE.
https://ffmuc.net/

sdns://AgcAAAAAAAAACjUuMS42Ni4yNTWgzBBg05yDKbYrb7x9DW35MJhpuYHn5jktXNj6QI9NgOag9_WvKIAeh31986K-KP4UnzJ0p-0p8Tb9UDzjmMuoCw2gs14FlQz72CWfk3W6EBhwuMPEwOxaUpdX5jFn6d4mqeigRE69Z7uD-IB7OSHpOKyReLiCvVCq2xEjHwRM9fCN986g5kS6aWPjNf52XLmXaxKxDrVClLQkd3ZMyzo6zKOssvwgKq4_t78F5MgcQZTcpEUR1PmvMEeG7BrnIYQJz2Kgg1UNZG9oLmZmbXVjLm5ldAovZG5zLXF1ZXJ5


## doh.ffmuc.net-2

An open (non-logging, non-filtering, non-censoring) DoH resolver operated by Freifunk Munich with nodes in DE.
https://ffmuc.net/

sdns://AgcAAAAAAAAADjE4NS4xNTAuOTkuMjU1oMwQYNOcgym2K2-8fQ1t-TCYabmB5-Y5LVzY-kCPTYDmoPf1ryiAHod9ffOivij-FJ8ydKftKfE2_VA845jLqAsNoLNeBZUM-9gln5N1uhAYcLjDxMDsWlKXV-YxZ-neJqnooEROvWe7g_iAezkh6TiskXi4gr1QqtsRIx8ETPXwjffOoOZEumlj4zX-dly5l2sSsQ61QpS0JHd2TMs6OsyjrLL8ICquP7e_BeTIHEGU3KRFEdT5rzBHhuwa5yGECc9ioINVDWRvaC5mZm11Yy5uZXQKL2Rucy1xdWVyeQ


## doh.ffmuc.net-v6

An open (non-logging, non-filtering, non-censoring) DoH resolver operated by Freifunk Munich with nodes in DE.
https://ffmuc.net/

sdns://AgcAAAAAAAAAFVsyMDAxOjY3ODplNjg6ZjAwMDo6XaDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5qD39a8ogB6HfX3zor4o_hSfMnSn7SnxNv1QPOOYy6gLDaCzXgWVDPvYJZ-TdboQGHC4w8TA7FpSl1fmMWfp3iap6KBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33zqDmRLppY-M1_nZcuZdrErEOtUKUtCR3dkzLOjrMo6yy_CAqrj-3vwXkyBxBlNykRRHU-a8wR4bsGuchhAnPYqCDVQ1kb2guZmZtdWMubmV0Ci9kbnMtcXVlcnk


## doh.ffmuc.net-v6-2

An open (non-logging, non-filtering, non-censoring) DoH resolver operated by Freifunk Munich with nodes in DE.
https://ffmuc.net/

sdns://AgcAAAAAAAAAFVsyMDAxOjY3ODplZDA6ZjAwMDo6XaDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5qD39a8ogB6HfX3zor4o_hSfMnSn7SnxNv1QPOOYy6gLDaCzXgWVDPvYJZ-TdboQGHC4w8TA7FpSl1fmMWfp3iap6KBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33zqDmRLppY-M1_nZcuZdrErEOtUKUtCR3dkzLOjrMo6yy_CAqrj-3vwXkyBxBlNykRRHU-a8wR4bsGuchhAnPYqCDVQ1kb2guZmZtdWMubmV0Ci9kbnMtcXVlcnk


## doh.tiar.app

Non-Logging DNSCrypt server located in Singapore.
Filters out ads, trackers and malware, supports DNSSEC, provided by id-gmail.

sdns://AQMAAAAAAAAADjE3NC4xMzguMjEuMTI4IO-WgGbo2ZTwZdg-3dMa7u31bYZXRj5KykfN1_6Xw9T2HDIuZG5zY3J5cHQtY2VydC5kbnMudGlhci5hcHA


## doh.tiar.app-doh

Non-Logging DNS-over-HTTPS (HTTP/2 & HTTP/3) server located in Singapore.
Filters out ads, trackers and malware, supports DNSSEC, provided by id-gmail.

sdns://AgMAAAAAAAAADjE3NC4xMzguMjkuMTc1oMwQYNOcgym2K2-8fQ1t-TCYabmB5-Y5LVzY-kCPTYDmoPf1ryiAHod9ffOivij-FJ8ydKftKfE2_VA845jLqAsNoLNeBZUM-9gln5N1uhAYcLjDxMDsWlKXV-YxZ-neJqnooEROvWe7g_iAezkh6TiskXi4gr1QqtsRIx8ETPXwjffOoOZEumlj4zX-dly5l2sSsQ61QpS0JHd2TMs6OsyjrLL8ICquP7e_BeTIHEGU3KRFEdT5rzBHhuwa5yGECc9ioINVDGRvaC50aWFyLmFwcAovZG5zLXF1ZXJ5


## doh.tiarap.org

Non-Logging DNS-over-HTTPS server, cached via Cloudflare.
Filters out ads, trackers and malware, NO ECS, supports DNSSEC.

sdns://AgMAAAAAAAAADDEwNC4yMS42NS42MAAOZG9oLnRpYXJhcC5vcmcKL2Rucy1xdWVyeQ


## doh.tiarap.org-ipv6

Non-Logging DNS-over-HTTPS server (IPv6), cached via Cloudflare.
Filters out ads, trackers and malware, NO ECS, supports DNSSEC.

sdns://AgMAAAAAAAAAG1syNjA2OjQ3MDA6MzAzNDo6NjgxNTo0MTNjXQAOZG9oLnRpYXJhcC5vcmcKL2Rucy1xdWVyeQ


## easymosdns-doh

DoH forwarded runned personally, as a example server of the project [EasyMosdns](https://github.com/pmkol/easymosdns).
No filtering or logs (by the forwarder itself).
No DNSSEC support due to one of the upstream servers (AliDNS) that doesn't support it.
Upstreams are AliDNS and DNSPod for resolving domains in mainland China, and GoogleDNS is for domains.
Cloudflare CDN is used as a front-end for non-China areas, and Mobile CDN is used in China.

Homepage: https://apad.pro/dns-doh/

sdns://AgQAAAAAAAAAAAAMZG9oLmFwYWQucHJvCi9kbnMtcXVlcnk


## faelix-ch-ipv4

An open (non-logging, non-filtering, no ECS) DNSCrypt resolver operated by https://faelix.net/ with IPv4 nodes anycast within AS41495 in Switzerland.

sdns://AQYAAAAAAAAAEzE4NS4xMzQuMTk2LjU1Ojg0NDMgfsvvPi8BgDKNYODh0ewj5Oh32OoJoZNwGgTWs8C-i-EfMi5kbnNjcnlwdC1jZXJ0LnJkbnMuZmFlbGl4Lm5ldA


## faelix-uk-ipv4

An open (non-logging, non-filtering, no ECS) DNSCrypt resolver operated by https://faelix.net/ with IPv4 nodes anycast within AS41495 in the UK.

sdns://AQYAAAAAAAAAEjQ2LjIyNy4yMDAuNTQ6ODQ0MyB-y-8-LwGAMo1g4OHR7CPk6HfY6gmhk3AaBNazwL6L4R8yLmRuc2NyeXB0LWNlcnQucmRucy5mYWVsaXgubmV0


## faelix-uk-ipv6

An open (non-logging, non-filtering, no ECS) DNSCrypt resolver operated by https://faelix.net/ with IPv6 nodes anycast within AS41495 in the UK.

sdns://AQYAAAAAAAAAFFsyYTAxOjllMDA6OjU0XTo4NDQzIH7L7z4vAYAyjWDg4dHsI-Tod9jqCaGTcBoE1rPAvovhHzIuZG5zY3J5cHQtY2VydC5yZG5zLmZhZWxpeC5uZXQ
sdns://AQYAAAAAAAAAFFsyYTAxOjllMDA6OjU1XTo4NDQzIH7L7z4vAYAyjWDg4dHsI-Tod9jqCaGTcBoE1rPAvovhHzIuZG5zY3J5cHQtY2VydC5yZG5zLmZhZWxpeC5uZXQ


## fdn

DoH server in France operated by FDN - French Data Network (non-profit ISP)
https://www.fdn.fr/

sdns://AgcAAAAAAAAADDgwLjY3LjE2OS4xMqDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5qD39a8ogB6HfX3zor4o_hSfMnSn7SnxNv1QPOOYy6gLDaCzXgWVDPvYJZ-TdboQGHC4w8TA7FpSl1fmMWfp3iap6KBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33zqDmRLppY-M1_nZcuZdrErEOtUKUtCR3dkzLOjrMo6yy_CAqrj-3vwXkyBxBlNykRRHU-a8wR4bsGuchhAnPYqCDVQpuczAuZmRuLmZyCi9kbnMtcXVlcnk
sdns://AgcAAAAAAAAADDgwLjY3LjE2OS40MKDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5qD39a8ogB6HfX3zor4o_hSfMnSn7SnxNv1QPOOYy6gLDaCzXgWVDPvYJZ-TdboQGHC4w8TA7FpSl1fmMWfp3iap6KBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33zqDmRLppY-M1_nZcuZdrErEOtUKUtCR3dkzLOjrMo6yy_CAqrj-3vwXkyBxBlNykRRHU-a8wR4bsGuchhAnPYqCDVQpuczEuZmRuLmZyCi9kbnMtcXVlcnk


## fdn-ipv6

DoH server in France operated by FDN - French Data Network (non-profit ISP)
https://www.fdn.fr/

sdns://AgcAAAAAAAAAElsyMDAxOjkxMDo4MDA6OjEyXaDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5qD39a8ogB6HfX3zor4o_hSfMnSn7SnxNv1QPOOYy6gLDaCzXgWVDPvYJZ-TdboQGHC4w8TA7FpSl1fmMWfp3iap6KBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33zqDmRLppY-M1_nZcuZdrErEOtUKUtCR3dkzLOjrMo6yy_CAqrj-3vwXkyBxBlNykRRHU-a8wR4bsGuchhAnPYqCDVQpuczAuZmRuLmZyCi9kbnMtcXVlcnk
sdns://AgcAAAAAAAAAElsyMDAxOjkxMDo4MDA6OjQwXaDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5qD39a8ogB6HfX3zor4o_hSfMnSn7SnxNv1QPOOYy6gLDaCzXgWVDPvYJZ-TdboQGHC4w8TA7FpSl1fmMWfp3iap6KBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33zqDmRLppY-M1_nZcuZdrErEOtUKUtCR3dkzLOjrMo6yy_CAqrj-3vwXkyBxBlNykRRHU-a8wR4bsGuchhAnPYqCDVQpuczEuZmRuLmZyCi9kbnMtcXVlcnk


## ffmuc.net

An open (non-logging, non-filtering, non-censoring) DNSCrypt resolver operated by Freifunk Munich with nodes in DE.
https://ffmuc.net/

sdns://AQcAAAAAAAAADzUuMS42Ni4yNTU6ODQ0MyAH0Hrxz9xdmXadPwJmkKcESWXCdCdseRyu9a7zuQxG-hkyLmRuc2NyeXB0LWNlcnQuZmZtdWMubmV0


## ffmuc.net-v6

An open (non-logging, non-filtering, non-censoring) DNSCrypt resolver operated by Freifunk Munich with nodes in DE.
https://ffmuc.net/

sdns://AQcAAAAAAAAAGlsyMDAxOjY3ODplNjg6ZjAwMDo6XTo4NDQzIAfQevHP3F2Zdp0_AmaQpwRJZcJ0J2x5HK71rvO5DEb6GTIuZG5zY3J5cHQtY2VydC5mZm11Yy5uZXQ


## fluffycat-fr-01

DNSCrypt resolver hosted in Marseille, FR on Oracle Cloud. No logs, no filters, DNSSEC.

sdns://AQcAAAAAAAAAFDEyOS4xNTEuMjI0LjE4Mjo1MzUzICo3YofAHoNsHkixOFHSQXY6G7z1MCrxXxGyfovtQvsGHzIuZG5zY3J5cHQtY2VydC5mbHVmZnljYXQtZnItMDE


## fluffycat-fr-02

DNSCrypt resolver hosted in Marseille, FR on Oracle Cloud. No logs, no filters, DNSSEC.

sdns://AQcAAAAAAAAAFDEyOS4xNTEuMjQzLjE0Mzo1MzUzICaU0eKbEtcmoE0ljHjvADPHNyZBX23wJ4owxhVprIpFHzIuZG5zY3J5cHQtY2VydC5mbHVmZnljYXQtZnItMDI


## google

Google DNS (anycast)

sdns://AgUAAAAAAAAABzguOC44LjggsKKKE4EwvtIbNjGjagI2607EdKSVHowYZtyvD9iPrkkHOC44LjguOAovZG5zLXF1ZXJ5
sdns://AgUAAAAAAAAABzguOC40LjQgsKKKE4EwvtIbNjGjagI2607EdKSVHowYZtyvD9iPrkkHOC44LjQuNAovZG5zLXF1ZXJ5


## google-ipv6

Google DNS (anycast)

sdns://AgUAAAAAAAAAACCwoooTgTC-0hs2MaNqAjbrTsR0pJUejBhm3K8P2I-uSRZbMjAwMTo0ODYwOjQ4NjA6Ojg4ODhdCi9kbnMtcXVlcnk
sdns://AgUAAAAAAAAAACCwoooTgTC-0hs2MaNqAjbrTsR0pJUejBhm3K8P2I-uSRZbMjAwMTo0ODYwOjQ4NjA6Ojg4NDRdCi9kbnMtcXVlcnk
sdns://AgUAAAAAAAAAACCwoooTgTC-0hs2MaNqAjbrTsR0pJUejBhm3K8P2I-uSRRbMjAwMTo0ODYwOjQ4NjA6OjY0XQovZG5zLXF1ZXJ5
sdns://AgUAAAAAAAAAACCwoooTgTC-0hs2MaNqAjbrTsR0pJUejBhm3K8P2I-uSRZbMjAwMTo0ODYwOjQ4NjA6OjY0NjRdCi9kbnMtcXVlcnk


## he

Hurricane Electric DoH server (anycast)

Unknown logging policy.

sdns://AgUAAAAAAAAACzc0LjgyLjQyLjQyoMwQYNOcgym2K2-8fQ1t-TCYabmB5-Y5LVzY-kCPTYDmoPf1ryiAHod9ffOivij-FJ8ydKftKfE2_VA845jLqAsNoLNeBZUM-9gln5N1uhAYcLjDxMDsWlKXV-YxZ-neJqnooEROvWe7g_iAezkh6TiskXi4gr1QqtsRIx8ETPXwjffOoOZEumlj4zX-dly5l2sSsQ61QpS0JHd2TMs6OsyjrLL8ICquP7e_BeTIHEGU3KRFEdT5rzBHhuwa5yGECc9ioINVDG9yZG5zLmhlLm5ldAovZG5zLXF1ZXJ5


## ibksturm

Dnscrypt Server, No Logging, No Filters, DNSSEC, OpenNIC

Running privately by ibksturm in Thurgau, Switzerland

sdns://AQcAAAAAAAAAEzIxMy4xOTYuMTkxLjk2Ojg0NDMgiwvumeI8et789m3naGH-4xzM40t6c2xO_fCxHldawJgYMi5kbnNjcnlwdC1jZXJ0Lmlia3N0dXJt


## iij

DoH server operated by Internet Initiative Japan in Tokyo. Blocks child pornography.
https://www.iij.ad.jp/

sdns://AgEAAAAAAAAACjEwMy4yLjU3LjUgmOPV5TavKVjNL38U9wTvSidtJeM81l8uZfXk8nJ8EzARcHVibGljLmRucy5paWouanAKL2Rucy1xdWVyeQ


## jp.tiar.app

Non-Logging, Non-Filtering DNSCrypt server in Japan.
No ECS, Support DNSSEC

sdns://AQcAAAAAAAAAEjE3Mi4xMDQuOTMuODA6MTQ0MyAyuHY-8b9lNqHeahPAzW9IoXnjiLaZpTeNbVs8TN9UUxsyLmRuc2NyeXB0LWNlcnQuanAudGlhci5hcHA


## jp.tiar.app-doh

Non-Logging, Non-Filtering DNS-over-HTTPS server in Japan.
No ECS, Support DNSSEC

sdns://AgcAAAAAAAAADTE3Mi4xMDQuOTMuODCgzBBg05yDKbYrb7x9DW35MJhpuYHn5jktXNj6QI9NgOag9_WvKIAeh31986K-KP4UnzJ0p-0p8Tb9UDzjmMuoCw2gs14FlQz72CWfk3W6EBhwuMPEwOxaUpdX5jFn6d4mqeigRE69Z7uD-IB7OSHpOKyReLiCvVCq2xEjHwRM9fCN986g5kS6aWPjNf52XLmXaxKxDrVClLQkd3ZMyzo6zKOssvwgKq4_t78F5MgcQZTcpEUR1PmvMEeG7BrnIYQJz2Kgg1ULanAudGlhci5hcHAKL2Rucy1xdWVyeQ


## jp.tiar.app-doh-ipv6

Non-Logging, Non-Filtering DNS-over-HTTPS (IPv6) server in Japan.
No ECS, Support DNSSEC

sdns://AgcAAAAAAAAAIFsyNDAwOjg5MDI6OmYwM2M6OTFmZjpmZWRhOmM1MTRdoMwQYNOcgym2K2-8fQ1t-TCYabmB5-Y5LVzY-kCPTYDmoPf1ryiAHod9ffOivij-FJ8ydKftKfE2_VA845jLqAsNoLNeBZUM-9gln5N1uhAYcLjDxMDsWlKXV-YxZ-neJqnooEROvWe7g_iAezkh6TiskXi4gr1QqtsRIx8ETPXwjffOoOZEumlj4zX-dly5l2sSsQ61QpS0JHd2TMs6OsyjrLL8ICquP7e_BeTIHEGU3KRFEdT5rzBHhuwa5yGECc9ioINVC2pwLnRpYXIuYXBwCi9kbnMtcXVlcnk


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

sdns://AgcAAAAAAAAAG1syNjA2OjQ3MDA6MzAzMDo6YWM0MzphZDNiXQANanAudGlhcmFwLm9yZwovZG5zLXF1ZXJ5


## ksol.io-ns2-dnscrypt-ipv4

DNSCrypt on IPv4 (UDP/TCP). No DoH, doesn't log, doesn't filter, DNSSEC enforced. No EDNS Client-Subnet, padding enabled, as per `dnscrypt-server-docker` default unbound configuration. Location: Hungary

sdns://AQcAAAAAAAAADjE5My4yMDEuMTg4LjQ4IBERKdQJgLSjqCSK99e2f_WRTQzEq9__DeXlQFvxxhZ6GzIuZG5zY3J5cHQtY2VydC5uczIua3NvbC5pbw


## ksol.io-ns2-dnscrypt-ipv6

DNSCrypt on IPv6 (UDP/TCP). No DoH, doesn't log, doesn't filter, DNSSEC enforced. No EDNS Client-Subnet, padding enabled, as per `dnscrypt-server-docker` default unbound configuration. Location: Hungary

sdns://AQcAAAAAAAAAFFsyYTAxOjZlZTA6MTo6MjQxOjFdIBERKdQJgLSjqCSK99e2f_WRTQzEq9__DeXlQFvxxhZ6GzIuZG5zY3J5cHQtY2VydC5uczIua3NvbC5pbw


## libredns

DoH server in Germany. No logging, but no DNS padding and no DNSSEC support.
https://libredns.gr/

sdns://AgYAAAAAAAAADjExNi4yMDIuMTc2LjI2oMwQYNOcgym2K2-8fQ1t-TCYabmB5-Y5LVzY-kCPTYDmoPf1ryiAHod9ffOivij-FJ8ydKftKfE2_VA845jLqAsNoLNeBZUM-9gln5N1uhAYcLjDxMDsWlKXV-YxZ-neJqnooEROvWe7g_iAezkh6TiskXi4gr1QqtsRIx8ETPXwjffOoOZEumlj4zX-dly5l2sSsQ61QpS0JHd2TMs6OsyjrLL8ICquP7e_BeTIHEGU3KRFEdT5rzBHhuwa5yGECc9ioINVD2RvaC5saWJyZWRucy5ncgovZG5zLXF1ZXJ5


## libredns-noads

DoH server in Germany. No logging, but no DNS padding and no DNSSEC support.
no ads version, uses StevenBlack's host list: https://github.com/StevenBlack/hosts

sdns://AgIAAAAAAAAADjExNi4yMDIuMTc2LjI2oMwQYNOcgym2K2-8fQ1t-TCYabmB5-Y5LVzY-kCPTYDmoPf1ryiAHod9ffOivij-FJ8ydKftKfE2_VA845jLqAsNoLNeBZUM-9gln5N1uhAYcLjDxMDsWlKXV-YxZ-neJqnooEROvWe7g_iAezkh6TiskXi4gr1QqtsRIx8ETPXwjffOoOZEumlj4zX-dly5l2sSsQ61QpS0JHd2TMs6OsyjrLL8ICquP7e_BeTIHEGU3KRFEdT5rzBHhuwa5yGECc9ioINVD2RvaC5saWJyZWRucy5ncgQvYWRz


## mullvad-adblock-doh

Same as mullvad-doh but blocks ads and trackers.

sdns://AgMAAAAAAAAACzE5NC4yNDIuMi4zABdhZGJsb2NrLmRucy5tdWxsdmFkLm5ldAovZG5zLXF1ZXJ5


## mullvad-all-doh

Same as mullvad-doh but blocks ads, trackers, malware, adult content, gambling, and social media.

sdns://AgMAAAAAAAAACzE5NC4yNDIuMi45ABNhbGwuZG5zLm11bGx2YWQubmV0Ci9kbnMtcXVlcnk


## mullvad-base-doh

Same as mullvad-doh but blocks ads, trackers, and malware.

sdns://AgMAAAAAAAAACzE5NC4yNDIuMi40ABRiYXNlLmRucy5tdWxsdmFkLm5ldAovZG5zLXF1ZXJ5


## mullvad-doh

Public non-filtering, non-logging (audited), DNSSEC-capable, DNS-over-HTTPS resolver hosted by VPN provider Mullvad.
Anycast IPv4/IPv6 with servers in SE, DE, UK, US, AU, and SG.
https://mullvad.net/en/help/dns-over-https-and-dns-over-tls/

sdns://AgcAAAAAAAAACzE5NC4yNDIuMi4yAA9kbnMubXVsbHZhZC5uZXQKL2Rucy1xdWVyeQ


## mullvad-extend-doh

Same as mullvad-doh but blocks ads, trackers, malware, and social media.

sdns://AgMAAAAAAAAACzE5NC4yNDIuMi41ABhleHRlbmRlZC5kbnMubXVsbHZhZC5uZXQKL2Rucy1xdWVyeQ


## mullvad-family-doh

Same as mullvad-doh but blocks ads, trackers, malware, adult content, and gambling.

sdns://AgMAAAAAAAAACzE5NC4yNDIuMi42ABZmYW1pbHkuZG5zLm11bGx2YWQubmV0Ci9kbnMtcXVlcnk


## nextdns

NextDNS is a cloud-based private DNS service that gives you full control
over what is allowed and what is blocked on the Internet.

DNSSEC, Anycast, Non-logging, NoFilters

https://www.nextdns.io/

sdns://AgcAAAAAAAAACjQ1LjkwLjMwLjAgmjo09yfeubylEAPZzpw5-PJ92cUkKQHCurGkTmNaAhkWYW55Y2FzdC5kbnMubmV4dGRucy5pbwovZG5zLXF1ZXJ5


## nextdns-ipv6

Connects to NextDNS over IPv6.

NextDNS is a cloud-based private DNS service that gives you full control
over what is allowed and what is blocked on the Internet.

DNSSEC, Anycast, Non-logging, NoFilters

https://www.nextdns.io/

sdns://AgcAAAAAAAAADVsyYTA3OmE4YzA6Ol0gmjo09yfeubylEAPZzpw5-PJ92cUkKQHCurGkTmNaAhkWYW55Y2FzdC5kbnMubmV4dGRucy5pbwovZG5zLXF1ZXJ5


## nextdns-ultralow

NextDNS is a cloud-based private DNS service that gives you full control
over what is allowed and what is blocked on the Internet.

https://www.nextdns.io/

To select the server location, the "-ultralow" variant relies on bootstrap servers
instead of anycast.

sdns://AgcAAAAAAAAAACCaOjT3J965vKUQA9nOnDn48n3ZxSQpAcK6saROY1oCGQ5kbnMubmV4dGRucy5pbw8vZG5zY3J5cHQtcHJveHk


## nic.cz

Open, DNSSEC, No-log and No-filter DoH operated by https://nic.cz

sdns://AgcAAAAAAAAADDE4NS40My4xMzUuMaDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5qD39a8ogB6HfX3zor4o_hSfMnSn7SnxNv1QPOOYy6gLDaCzXgWVDPvYJZ-TdboQGHC4w8TA7FpSl1fmMWfp3iap6KBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33zqDmRLppY-M1_nZcuZdrErEOtUKUtCR3dkzLOjrMo6yy_CAqrj-3vwXkyBxBlNykRRHU-a8wR4bsGuchhAnPYqCDVQtvZHZyLm5pYy5jegovZG5zLXF1ZXJ5
sdns://AgcAAAAAAAAACzE5My4xNy40Ny4xoMwQYNOcgym2K2-8fQ1t-TCYabmB5-Y5LVzY-kCPTYDmoPf1ryiAHod9ffOivij-FJ8ydKftKfE2_VA845jLqAsNoLNeBZUM-9gln5N1uhAYcLjDxMDsWlKXV-YxZ-neJqnooEROvWe7g_iAezkh6TiskXi4gr1QqtsRIx8ETPXwjffOoOZEumlj4zX-dly5l2sSsQ61QpS0JHd2TMs6OsyjrLL8ICquP7e_BeTIHEGU3KRFEdT5rzBHhuwa5yGECc9ioINVC29kdnIubmljLmN6Ci9kbnMtcXVlcnk


## nic.cz-ipv6

Open, DNSSEC, No-log and No-filter DoH over IPv6 operated by https://nic.cz

sdns://AgcAAAAAAAAAE1syMDAxOjE0OGY6ZmZmZTo6MV2gzBBg05yDKbYrb7x9DW35MJhpuYHn5jktXNj6QI9NgOag9_WvKIAeh31986K-KP4UnzJ0p-0p8Tb9UDzjmMuoCw2gs14FlQz72CWfk3W6EBhwuMPEwOxaUpdX5jFn6d4mqeigRE69Z7uD-IB7OSHpOKyReLiCvVCq2xEjHwRM9fCN986g5kS6aWPjNf52XLmXaxKxDrVClLQkd3ZMyzo6zKOssvwgKq4_t78F5MgcQZTcpEUR1PmvMEeG7BrnIYQJz2Kgg1ULb2R2ci5uaWMuY3oKL2Rucy1xdWVyeQ
sdns://AgcAAAAAAAAAE1syMDAxOjE0OGY6ZmZmZjo6MV2gzBBg05yDKbYrb7x9DW35MJhpuYHn5jktXNj6QI9NgOag9_WvKIAeh31986K-KP4UnzJ0p-0p8Tb9UDzjmMuoCw2gs14FlQz72CWfk3W6EBhwuMPEwOxaUpdX5jFn6d4mqeigRE69Z7uD-IB7OSHpOKyReLiCvVCq2xEjHwRM9fCN986g5kS6aWPjNf52XLmXaxKxDrVClLQkd3ZMyzo6zKOssvwgKq4_t78F5MgcQZTcpEUR1PmvMEeG7BrnIYQJz2Kgg1ULb2R2ci5uaWMuY3oKL2Rucy1xdWVyeQ


## njalla-doh

Non-logging DoH server in Sweden operated by Njalla.

https://dns.njal.la/

sdns://AgcAAAAAAAAADDk1LjIxNS4xOS41M6DMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5qD39a8ogB6HfX3zor4o_hSfMnSn7SnxNv1QPOOYy6gLDaCzXgWVDPvYJZ-TdboQGHC4w8TA7FpSl1fmMWfp3iap6KBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33zqDmRLppY-M1_nZcuZdrErEOtUKUtCR3dkzLOjrMo6yy_CAqrj-3vwXkyBxBlNykRRHU-a8wR4bsGuchhAnPYqCDVQtkbnMubmphbC5sYQovZG5zLXF1ZXJ5


## plan9dns-fl

Miami Florida, USA. DNSCrypt, no filters, no logs, DNSSEC

Hosted on Vultr, jlongua.github.io/plan9-dns

sdns://AQcAAAAAAAAAEzE0OS4yOC4xMDEuMTE5Ojg0NDMgVaFV4a8StIfx8fnCxDxVlxppqm-hJYyCKqtMtQENnCwkMi5kbnNjcnlwdC1jZXJ0LnBsdXRvbi5wbGFuOS1kbnMuY29t


## plan9dns-fl-doh

Miami Florida, USA. DoH, DoH3 via the Alt-Svc header, no-logs, no-filters, DNSSEC

Hosted on Vultr, jlongua.github.io/plan9-dns, DoT & DoQ supported.

sdns://AgcAAAAAAAAADjE0OS4yOC4xMDEuMTE5IJo6NPcn3rm8pRAD2c6cOfjyfdnFJCkBwrqxpE5jWgIZFHBsdXRvbi5wbGFuOS1kbnMuY29tCi9kbnMtcXVlcnk


## plan9dns-fl-doh-ipv6

Miami Florida, USA. DoH, DoH3 via the Alt-Svc header, no-logs, no-filters, DNSSEC

Hosted on Vultr, jlongua.github.io/plan9-dns, DoT & DoQ supported.

sdns://AgcAAAAAAAAAJ1syMDAxOjE5ZjA6OTAwMjpkZTQ6NTQwMDo0ZmY6ZmUwODo3ZGUzXSCaOjT3J965vKUQA9nOnDn48n3ZxSQpAcK6saROY1oCGRRwbHV0b24ucGxhbjktZG5zLmNvbQovZG5zLXF1ZXJ5


## plan9dns-fl-ipv6

Miami Florida, USA. DNSCrypt, no filters, no logs, DNSSEC

Hosted on Vultr, jlongua.github.io/plan9-dns

sdns://AQcAAAAAAAAALFsyMDAxOjE5ZjA6OTAwMjpkZTQ6NTQwMDo0ZmY6ZmUwODo3ZGUzXTo4NDQzIFWhVeGvErSH8fH5wsQ8VZcaaapvoSWMgiqrTLUBDZwsJDIuZG5zY3J5cHQtY2VydC5wbHV0b24ucGxhbjktZG5zLmNvbQ


## plan9dns-mx

Mexico City, MX. DNSCrypt, no filters, no logs, DNSSEC

Hosted on Vultr, jlongua.github.io/plan9-dns

sdns://AQcAAAAAAAAAEzIxNi4yMzguODAuMjE5Ojg0NDMgKmPCui35rtOj9yk7c07sEtC_Khyo_9_HcpO23GCroNskMi5kbnNjcnlwdC1jZXJ0LmhlbGlvcy5wbGFuOS1kbnMuY29t


## plan9dns-mx-doh

Mexico City, MX. DoH, DoH3 via the Alt-Svc header, no-logs, no-filters, DNSSEC

Hosted on Vultr, jlongua.github.io/plan9-dns, DoT & DoQ supported.

sdns://AgcAAAAAAAAADjIxNi4yMzguODAuMjE5IJo6NPcn3rm8pRAD2c6cOfjyfdnFJCkBwrqxpE5jWgIZFGhlbGlvcy5wbGFuOS1kbnMuY29tCi9kbnMtcXVlcnk


## plan9dns-mx-doh-ipv6

Mexico City, MX. DoH, DoH3 via the Alt-Svc header, no-logs, no-filters, DNSSEC

Hosted on Vultr, jlongua.github.io/plan9-dns, DoT & DoQ supported.

sdns://AgcAAAAAAAAAKFsyMDAxOjE5ZjA6YjQwMDoxZDhjOjU0MDA6NGZmOmZlMTE6YjE1YV0gmjo09yfeubylEAPZzpw5-PJ92cUkKQHCurGkTmNaAhkUaGVsaW9zLnBsYW45LWRucy5jb20KL2Rucy1xdWVyeQ


## plan9dns-mx-ipv6

Mexico City, MX. DNSCrypt, no filters, no logs, DNSSEC

Hosted on Vultr, jlongua.github.io/plan9-dns

sdns://AQcAAAAAAAAALVsyMDAxOjE5ZjA6YjQwMDoxZDhjOjU0MDA6NGZmOmZlMTE6YjE1YV06ODQ0MyAqY8K6Lfmu06P3KTtzTuwS0L8qHKj_38dyk7bcYKug2yQyLmRuc2NyeXB0LWNlcnQuaGVsaW9zLnBsYW45LWRucy5jb20


## plan9dns-nj

Piscataway New Jersey, USA. DNSCrypt, no filters, no logs, DNSSEC

Hosted on Vultr, jlongua.github.io/plan9-dns

sdns://AQcAAAAAAAAAEjIwNy4yNDYuODcuOTY6ODQ0MyCwmQlIDpKk8SiiyrJbPgKhHxCrBJLb8ZWlu6tvr1KvkyQyLmRuc2NyeXB0LWNlcnQua3Jvbm9zLnBsYW45LWRucy5jb20


## plan9dns-nj-doh

Piscataway New Jersey, USA. DoH, DoH3 via the Alt-Svc header, no-logs, no-filters, DNSSEC

Hosted on Vultr, jlongua.github.io/plan9-dns, DoT & DoQ supported.

sdns://AgcAAAAAAAAADTIwNy4yNDYuODcuOTYgmjo09yfeubylEAPZzpw5-PJ92cUkKQHCurGkTmNaAhkUa3Jvbm9zLnBsYW45LWRucy5jb20KL2Rucy1xdWVyeQ


## plan9dns-nj-doh-ipv6

Piscataway New Jersey, USA. DoH, DoH3 via the Alt-Svc header, no-logs, no-filters, DNSSEC

Hosted on Vultr, jlongua.github.io/plan9-dns, DoT & DoQ supported.

sdns://AgcAAAAAAAAAJVsyMDAxOjE5ZjA6NTozYmQ3OjU0MDA6NGZmOmZlMDU6ZGE4M10gmjo09yfeubylEAPZzpw5-PJ92cUkKQHCurGkTmNaAhkUa3Jvbm9zLnBsYW45LWRucy5jb20KL2Rucy1xdWVyeQ


## plan9dns-nj.ipv6

Piscataway New Jersey, USA. DNSCrypt, no filters, no logs, DNSSEC

Hosted on Vultr, jlongua.github.io/plan9-dns

sdns://AQcAAAAAAAAAKlsyMDAxOjE5ZjA6NTozYmQ3OjU0MDA6NGZmOmZlMDU6ZGE4M106ODQ0MyCwmQlIDpKk8SiiyrJbPgKhHxCrBJLb8ZWlu6tvr1KvkyQyLmRuc2NyeXB0LWNlcnQua3Jvbm9zLnBsYW45LWRucy5jb20


## pryv8boi

By pryv8, non Logging, uncensored, DNSSEC - hosted on contabo servers

sdns://AQcAAAAAAAAAEzE2NC42OC4xMjEuMTYyOjQ0NDMgHtKNfXpUMzPyLnXK8DauHpWm1Rqhz7LqwBBmSzdY9BIcMi5kbnNjcnlwdC1jZXJ0LnByeXY4Ym9pLm9yZw


## qihoo360-doh

DoH server runned by Qihoo 360, has logs, GFW filtering rules are applied.
Homepage: https://sdns.360.net

sdns://AgAAAAAAAAAAACBGRinMfizEpf6XTTrC4BNXB3syOlkxKYNaWpRbX7u40Qpkb2guMzYwLmNuCi9kbnMtcXVlcnk


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

Quad9 (anycast) no-dnssec/no-log/no-filter 9.9.9.10 - 149.112.112.10

sdns://AQYAAAAAAAAADTkuOS45LjEwOjg0NDMgZ8hHuMh1jNEgJFVDvnVnRt803x2EwAuMRwNo34Idhj4ZMi5kbnNjcnlwdC1jZXJ0LnF1YWQ5Lm5ldA
sdns://AQYAAAAAAAAAEzE0OS4xMTIuMTEyLjEwOjg0NDMgZ8hHuMh1jNEgJFVDvnVnRt803x2EwAuMRwNo34Idhj4ZMi5kbnNjcnlwdC1jZXJ0LnF1YWQ5Lm5ldA


## quad9-dnscrypt-ip6-filter-ecs-pri

Quad9 (anycast) dnssec/no-log/filter/ecs 2620:fe::11 - 2620:fe::fe:11

sdns://AQMAAAAAAAAAElsyNjIwOmZlOjoxMV06ODQ0MyBnyEe4yHWM0SAkVUO-dWdG3zTfHYTAC4xHA2jfgh2GPhkyLmRuc2NyeXB0LWNlcnQucXVhZDkubmV0
sdns://AQMAAAAAAAAAFVsyNjIwOmZlOjpmZToxMV06ODQ0MyBnyEe4yHWM0SAkVUO-dWdG3zTfHYTAC4xHA2jfgh2GPhkyLmRuc2NyeXB0LWNlcnQucXVhZDkubmV0


## quad9-dnscrypt-ip6-filter-pri

Quad9 (anycast) dnssec/no-log/filter 2620:fe::fe - 2620:fe::9 - 2620:fe::fe:9

sdns://AQMAAAAAAAAAElsyNjIwOmZlOjpmZV06ODQ0MyBnyEe4yHWM0SAkVUO-dWdG3zTfHYTAC4xHA2jfgh2GPhkyLmRuc2NyeXB0LWNlcnQucXVhZDkubmV0
sdns://AQMAAAAAAAAAEVsyNjIwOmZlOjo5XTo4NDQzIGfIR7jIdYzRICRVQ751Z0bfNN8dhMALjEcDaN-CHYY-GTIuZG5zY3J5cHQtY2VydC5xdWFkOS5uZXQ
sdns://AQMAAAAAAAAAFFsyNjIwOmZlOjpmZTo5XTo4NDQzIGfIR7jIdYzRICRVQ751Z0bfNN8dhMALjEcDaN-CHYY-GTIuZG5zY3J5cHQtY2VydC5xdWFkOS5uZXQ


## quad9-dnscrypt-ip6-nofilter-ecs-pri

Quad9 (anycast) no-dnssec/no-log/no-filter/ecs 2620:fe::12 - 2620:fe::fe:12

sdns://AQYAAAAAAAAAElsyNjIwOmZlOjoxMl06ODQ0MyBnyEe4yHWM0SAkVUO-dWdG3zTfHYTAC4xHA2jfgh2GPhkyLmRuc2NyeXB0LWNlcnQucXVhZDkubmV0
sdns://AQYAAAAAAAAAFVsyNjIwOmZlOjpmZToxMl06ODQ0MyBnyEe4yHWM0SAkVUO-dWdG3zTfHYTAC4xHA2jfgh2GPhkyLmRuc2NyeXB0LWNlcnQucXVhZDkubmV0


## quad9-dnscrypt-ip6-nofilter-pri

Quad9 (anycast) no-dnssec/no-log/no-filter 2620:fe::10 - 2620:fe::fe:10

sdns://AQYAAAAAAAAAElsyNjIwOmZlOjoxMF06ODQ0MyBnyEe4yHWM0SAkVUO-dWdG3zTfHYTAC4xHA2jfgh2GPhkyLmRuc2NyeXB0LWNlcnQucXVhZDkubmV0
sdns://AQYAAAAAAAAAFVsyNjIwOmZlOjpmZToxMF06ODQ0MyBnyEe4yHWM0SAkVUO-dWdG3zTfHYTAC4xHA2jfgh2GPhkyLmRuc2NyeXB0LWNlcnQucXVhZDkubmV0


## quad9-doh-ip4-port443-filter-ecs-pri

Quad9 (anycast) dnssec/no-log/filter/ecs 9.9.9.11 - 149.112.112.11

sdns://AgMAAAAAAAAACDkuOS45LjExILAZIHRLu3bJqwU-AeB7fgUORz0g95976kNfr-Q8nSQvE2RuczExLnF1YWQ5Lm5ldDo0NDMKL2Rucy1xdWVyeQ
sdns://AgMAAAAAAAAADjE0OS4xMTIuMTEyLjExILAZIHRLu3bJqwU-AeB7fgUORz0g95976kNfr-Q8nSQvE2RuczExLnF1YWQ5Lm5ldDo0NDMKL2Rucy1xdWVyeQ


## quad9-doh-ip4-port443-filter-pri

Quad9 (anycast) dnssec/no-log/filter 9.9.9.9 - 149.112.112.9 - 149.112.112.112

sdns://AgMAAAAAAAAABzkuOS45LjkgsBkgdEu7dsmrBT4B4Ht-BQ5HPSD3n3vqQ1-v5DydJC8SZG5zOS5xdWFkOS5uZXQ6NDQzCi9kbnMtcXVlcnk
sdns://AgMAAAAAAAAADTE0OS4xMTIuMTEyLjkgsBkgdEu7dsmrBT4B4Ht-BQ5HPSD3n3vqQ1-v5DydJC8SZG5zOS5xdWFkOS5uZXQ6NDQzCi9kbnMtcXVlcnk
sdns://AgMAAAAAAAAADzE0OS4xMTIuMTEyLjExMiCwGSB0S7t2yasFPgHge34FDkc9IPefe-pDX6_kPJ0kLxFkbnMucXVhZDkubmV0OjQ0MwovZG5zLXF1ZXJ5


## quad9-doh-ip4-port443-nofilter-ecs-pri

Quad9 (anycast) no-dnssec/no-log/no-filter/ecs 9.9.9.12 - 149.112.112.12

sdns://AgYAAAAAAAAACDkuOS45LjEyILAZIHRLu3bJqwU-AeB7fgUORz0g95976kNfr-Q8nSQvE2RuczEyLnF1YWQ5Lm5ldDo0NDMKL2Rucy1xdWVyeQ
sdns://AgYAAAAAAAAADjE0OS4xMTIuMTEyLjEyILAZIHRLu3bJqwU-AeB7fgUORz0g95976kNfr-Q8nSQvE2RuczEyLnF1YWQ5Lm5ldDo0NDMKL2Rucy1xdWVyeQ


## quad9-doh-ip4-port443-nofilter-pri

Quad9 (anycast) no-dnssec/no-log/no-filter 9.9.9.10 - 149.112.112.10

sdns://AgYAAAAAAAAACDkuOS45LjEwILAZIHRLu3bJqwU-AeB7fgUORz0g95976kNfr-Q8nSQvE2RuczEwLnF1YWQ5Lm5ldDo0NDMKL2Rucy1xdWVyeQ
sdns://AgYAAAAAAAAADjE0OS4xMTIuMTEyLjEwILAZIHRLu3bJqwU-AeB7fgUORz0g95976kNfr-Q8nSQvE2RuczEwLnF1YWQ5Lm5ldDo0NDMKL2Rucy1xdWVyeQ


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

sdns://AgMAAAAAAAAADVsyNjIwOmZlOjoxMV0gsBkgdEu7dsmrBT4B4Ht-BQ5HPSD3n3vqQ1-v5DydJC8TZG5zMTEucXVhZDkubmV0OjQ0MwovZG5zLXF1ZXJ5
sdns://AgMAAAAAAAAAEFsyNjIwOmZlOjpmZToxMV0gsBkgdEu7dsmrBT4B4Ht-BQ5HPSD3n3vqQ1-v5DydJC8TZG5zMTEucXVhZDkubmV0OjQ0MwovZG5zLXF1ZXJ5


## quad9-doh-ip6-port443-filter-pri

Quad9 (anycast) dnssec/no-log/filter 2620:fe::fe - 2620:fe::9 - 2620:fe::fe:9

sdns://AgMAAAAAAAAADVsyNjIwOmZlOjpmZV0gsBkgdEu7dsmrBT4B4Ht-BQ5HPSD3n3vqQ1-v5DydJC8RZG5zLnF1YWQ5Lm5ldDo0NDMKL2Rucy1xdWVyeQ
sdns://AgMAAAAAAAAADFsyNjIwOmZlOjo5XSCwGSB0S7t2yasFPgHge34FDkc9IPefe-pDX6_kPJ0kLxFkbnMucXVhZDkubmV0OjQ0MwovZG5zLXF1ZXJ5
sdns://AgMAAAAAAAAAD1syNjIwOmZlOjpmZTo5XSCwGSB0S7t2yasFPgHge34FDkc9IPefe-pDX6_kPJ0kLxJkbnM5LnF1YWQ5Lm5ldDo0NDMKL2Rucy1xdWVyeQ


## quad9-doh-ip6-port443-nofilter-ecs-pri

Quad9 (anycast) no-dnssec/no-log/no-filter/ecs 2620:fe::12 - 2620:fe::fe:12

sdns://AgYAAAAAAAAADVsyNjIwOmZlOjoxMl0gsBkgdEu7dsmrBT4B4Ht-BQ5HPSD3n3vqQ1-v5DydJC8TZG5zMTIucXVhZDkubmV0OjQ0MwovZG5zLXF1ZXJ5
sdns://AgYAAAAAAAAAEFsyNjIwOmZlOjpmZToxMl0gsBkgdEu7dsmrBT4B4Ht-BQ5HPSD3n3vqQ1-v5DydJC8TZG5zMTIucXVhZDkubmV0OjQ0MwovZG5zLXF1ZXJ5


## quad9-doh-ip6-port443-nofilter-pri

Quad9 (anycast) no-dnssec/no-log/no-filter 2620:fe::10 - 2620:fe::fe:10

sdns://AgYAAAAAAAAADVsyNjIwOmZlOjoxMF0gsBkgdEu7dsmrBT4B4Ht-BQ5HPSD3n3vqQ1-v5DydJC8TZG5zMTAucXVhZDkubmV0OjQ0MwovZG5zLXF1ZXJ5
sdns://AgYAAAAAAAAAEFsyNjIwOmZlOjpmZToxMF0gsBkgdEu7dsmrBT4B4Ht-BQ5HPSD3n3vqQ1-v5DydJC8TZG5zMTAucXVhZDkubmV0OjQ0MwovZG5zLXF1ZXJ5


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


## restena-doh-ipv4

DNSSEC, No-log and No-filter DoH operated by RESTENA. Homepage: https://www.restena.lu

sdns://AgcAAAAAAAAACzE1OC42NC4xLjI5IGDJBNhGNSOSEa34Gom8fB_myqDD7-nW-TeU7KcLNV6oEWRuc3B1Yi5yZXN0ZW5hLmx1Ci9kbnMtcXVlcnk


## restena-doh-ipv6

DNSSEC, No-log and No-filter DoH (IPv6) operated by RESTENA. Homepage: https://www.restena.lu

sdns://AgcAAAAAAAAAEFsyMDAxOmExODoxOjoyOV0gYMkE2EY1I5IRrfgaibx8H-bKoMPv6db5N5Tspws1XqgRZG5zcHViLnJlc3RlbmEubHUKL2Rucy1xdWVyeQ


## rethinkdns-doh

No-log, No-filter
RethinkDNS, a stub (sky.rethinkdns.com hosted on Cloudflare) and recursive (max.rethinkdns.com hosted on fly.io) resolver
The stub server strips identification parameters from the request and acts as a proxy to another recursive resolver.

sdns://AgYAAAAAAAAAACBdzvEcz84tL6QcR78t69kc0nufblyYal5di10An6SyUBJza3kucmV0aGlua2Rucy5jb20KL2Rucy1xdWVyeQ
sdns://AgYAAAAAAAAAACCaOjT3J965vKUQA9nOnDn48n3ZxSQpAcK6saROY1oCGRJtYXgucmV0aGlua2Rucy5jb20KL2Rucy1xdWVyeQ


## safesurfer

Family safety focused blocklist for over 2 million adult sites, as well as phishing and malware and more.
Free to use, paid for customizing blocking for more categories+sites and viewing usage at my.safesurfer.io. Logs taken for viewing
usage, data never sold - https://safesurfer.io

Warning: this server is incompatible with DNS anonymization.

sdns://AQIAAAAAAAAADzEwNC4xNTUuMjM3LjIyNSAnIH_VEgToNntINABd-f_R0wu-KpwzY55u2_iu2R1A2CAyLmRuc2NyeXB0LWNlcnQuc2FmZXN1cmZlci5jby5ueg


## safesurfer-doh

Family safety focused blocklist for over 2 million adult sites, as well as phishing and malware and more.
Free to use, paid for customizing blocking for more categories+sites and viewing usage at my.safesurfer.io. Logs taken for viewing
usage, data never sold - https://safesurfer.io

sdns://AgAAAAAAAAAAACBW1D4A3rPRi8QazGGnZq98S8bEb0ZyDjFdjqVEtSPc3BFkb2guc2FmZXN1cmZlci5pbwovZG5zLXF1ZXJ5


## saldns01-conoha-ipv4

Hosted on ConoHa VPS Tokyo region. No log. No filter. From experimental [&mu;ODNS project](https://junkurihara.github.io/dns/).

sdns://AQcAAAAAAAAAFDExOC4yNy4xMDguMTQwOjUwNDQzIPL_JdttytAhcocG8vZH5EVk-uQKKose72bD0ji5G-xlIjIuZG5zY3J5cHQtY2VydC5zYWxkbnMwMS50eXBlcS5vcmc


## saldns02-conoha-ipv4

Hosted on ConoHa VPS Tokyo region. No log. No filter. From experimental [&mu;ODNS project](https://junkurihara.github.io/dns/).

sdns://AQcAAAAAAAAAFTEzMy4xMzAuMTE4LjEwMzo1MDQ0MyB7SI0q4_Ff8lFRUCbjPtcAQ3HfdWlLxyGDUUNc3NUZdiIyLmRuc2NyeXB0LWNlcnQuc2FsZG5zMDIudHlwZXEub3Jn


## saldns03-conoha-ipv4

Hosted on ConoHa VPS Tokyo region. No log. No filter. From experimental [&mu;ODNS project](https://junkurihara.github.io/dns/).

sdns://AQcAAAAAAAAAFDEzMy4xMzAuOTguMjUwOjUwNDQzIFl1NfOwMd24kRlr0mXR4rKo-c_jMV7DBUVooDEY1xFeIjIuZG5zY3J5cHQtY2VydC5zYWxkbnMwMy50eXBlcS5vcmc


## scaleway-ams

DNSSEC/Non-logged/Uncensored in Amsterdam - DEV1-S instance donated by Scaleway.com
Maintained by Frank Denis - https://fr.dnscrypt.info

sdns://AQcAAAAAAAAADTUxLjE1LjEyMi4yNTAg6Q3ZfapcbHgiHKLF7QFoli0Ty1Vsz3RXs1RUbxUrwZAcMi5kbnNjcnlwdC1jZXJ0LnNjYWxld2F5LWFtcw


## scaleway-ams-ipv6

DNSSEC/Non-logged/Uncensored in Amsterdam - IPv6 only - DEV1-S instance donated by Scaleway.com
Maintained by Frank Denis - https://fr.dnscrypt.info

sdns://AQcAAAAAAAAAJlsyMDAxOmJjODoxNjQwOjFjZTI6ZGMwMDpmZjpmZTI4OjViMTddIOkN2X2qXGx4Ihyixe0BaJYtE8tVbM90V7NUVG8VK8GQHDIuZG5zY3J5cHQtY2VydC5zY2FsZXdheS1hbXM


## scaleway-fr

DNSSEC/Non-logged/Uncensored in Paris - DEV1-S instance donated by Scaleway.com
Maintained by Frank Denis - https://fr.dnscrypt.info

sdns://AQcAAAAAAAAADjIxMi40Ny4yMjguMTM2IOgBuE6mBr-wusDOQ0RbsV66ZLAvo8SqMa4QY2oHkDJNHzIuZG5zY3J5cHQtY2VydC5mci5kbnNjcnlwdC5vcmc


## scaleway-fr-ipv6

-DNSSEC/Non-logged/Uncensored in Paris - IPv6 only - DEV1-S instance donated by Scaleway.com
Maintained by Frank Denis - https://fr.dnscrypt.info

sdns://AQcAAAAAAAAAJVsyMDAxOmJjODo3MTA6NTgxODpkYzAwOmZmOmZlNWI6M2Y2M10g6AG4TqYGv7C6wM5DRFuxXrpksC-jxKoxrhBjageQMk0fMi5kbnNjcnlwdC1jZXJ0LmZyLmRuc2NyeXB0Lm9yZw


## serbica

Public DNSCrypt server in the Netherlands by https://litepay.ch

sdns://AQcAAAAAAAAAEzE4NS42Ni4xNDMuMTc4OjUzNTMg-Y2MQmGOXiggAEKulN-ITGEn_Kj3TIP1UK1X2wh3o7wXMi5kbnNjcnlwdC1jZXJ0LnNlcmJpY2E


## sfw.scaleway-fr

Uses deep learning to block adult websites. Free, DNSSEC, no logs.
Hosted in Paris, running on a 1-XS server donated by Scaleway.com

Maintained by Frank Denis - https://fr.dnscrypt.info/sfw.html

sdns://AQMAAAAAAAAADzE2My4xNzIuMTgwLjEyNSDfYnO_x1IZKotaObwMhaw_-WRF1zZE9mJygl01WPGh_x8yLmRuc2NyeXB0LWNlcnQuc2Z3LnNjYWxld2F5LWZy


## sth-ads-doh-se

Resolver in Stockholm, Sweden. HTTP/3 DoH server. Non-logging, remove ads and malware, DNSSEC.

sdns://AgMAAAAAAAAADTQ1LjE1My4xODcuOTYAGGRuc3NlLW5vYWRzLmFsZWtiZXJnLm5ldAovZG5zLXF1ZXJ5


## sth-dnscrypt-se

Resolver in Stockholm, Sweden. DNSCrypt server. Non-logging, non-filtering, DNSSEC.

sdns://AQcAAAAAAAAAEjQ1LjE1My4xODcuOTY6NDM0MyAwkzvlkzabRkYs-RrxrcuyTjr9R73mBsx1Y-Ud2o-Whx8yLmRuc2NyeXB0LWNlcnQuc3RoLWRuc2NyeXB0LXNl


## sth-doh-se

Resolver in Stockholm, Sweden. HTTP/3 DoH server. Non-logging, non-filtering, DNSSEC.

sdns://AgcAAAAAAAAADTQ1LjE1My4xODcuOTagzBBg05yDKbYrb7x9DW35MJhpuYHn5jktXNj6QI9NgOag9_WvKIAeh31986K-KP4UnzJ0p-0p8Tb9UDzjmMuoCw2gs14FlQz72CWfk3W6EBhwuMPEwOxaUpdX5jFn6d4mqeigRE69Z7uD-IB7OSHpOKyReLiCvVCq2xEjHwRM9fCN986g5kS6aWPjNf52XLmXaxKxDrVClLQkd3ZMyzo6zKOssvwgKq4_t78F5MgcQZTcpEUR1PmvMEeG7BrnIYQJz2Kgg1USZG5zc2UuYWxla2JlcmcubmV0Ci9kbnMtcXVlcnk


## sth-doh-se-ipv6

Resolver in Stockholm, Sweden. HTTP/3 DoH server. Non-logging, non-filtering, DNSSEC.

sdns://AgcAAAAAAAAAFVsyYTA5OmNkNDI6Zjo0MjViOjoxXaDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5qD39a8ogB6HfX3zor4o_hSfMnSn7SnxNv1QPOOYy6gLDaCzXgWVDPvYJZ-TdboQGHC4w8TA7FpSl1fmMWfp3iap6KBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33zqDmRLppY-M1_nZcuZdrErEOtUKUtCR3dkzLOjrMo6yy_CAqrj-3vwXkyBxBlNykRRHU-a8wR4bsGuchhAnPYqCDVRJkbnNzZS5hbGVrYmVyZy5uZXQKL2Rucy1xdWVyeQ


## switch

Public DoH service provided by SWITCH in Switzerland. Provides protection against malware, but does not block ads.
Homepage: https://www.switch.ch

sdns://AgMAAAAAAAAADTEzMC41OS4zMS4yNDggsBkgdEu7dsmrBT4B4Ht-BQ5HPSD3n3vqQ1-v5DydJC8NMTMwLjU5LjMxLjI0OAovZG5zLXF1ZXJ5
sdns://AgMAAAAAAAAADTEzMC41OS4zMS4yNTEgsBkgdEu7dsmrBT4B4Ht-BQ5HPSD3n3vqQ1-v5DydJC8NMTMwLjU5LjMxLjI1MQovZG5zLXF1ZXJ5


## switch-ipv6

Public DoH (IPv6) service provided by SWITCH in Switzerland. Provides protection against malware, but does not block ads.
Homepage: https://www.switch.ch

sdns://AgMAAAAAAAAAACCwGSB0S7t2yasFPgHge34FDkc9IPefe-pDX6_kPJ0kLxJbMjAwMTo2MjA6MDpmZjo6Ml0KL2Rucy1xdWVyeQ
sdns://AgMAAAAAAAAAACCwGSB0S7t2yasFPgHge34FDkc9IPefe-pDX6_kPJ0kLxJbMjAwMTo2MjA6MDpmZjo6M10KL2Rucy1xdWVyeQ


## tuna-doh-ipv6

DoH server provided by Tsinghua University TUNA Association, located in mainland China, no GFW poisoning yet it has a manual blacklist.

sdns://AgAAAAAAAAAAG1syNDAyOmYwMDA6MTo0MTY6MTAxOjY6Njo2XSACKC69_nqGQOLLyjCQJNWgwRXTVfrgMllhwr12J14JbB1kbnMudHVuYS50c2luZ2h1YS5lZHUuY246ODQ0MwovZG5zLXF1ZXJ5


## uncensoreddns-dk-ipv4

Also known as censurfridns.
DoH, no logs, no filter, unicast hosted in Denmark - https://blog.uncensoreddns.org

sdns://AgYAAAAAAAAADDg5LjIzMy40My43MaDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5qD39a8ogB6HfX3zor4o_hSfMnSn7SnxNv1QPOOYy6gLDaCzXgWVDPvYJZ-TdboQGHC4w8TA7FpSl1fmMWfp3iap6KBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33zqDmRLppY-M1_nZcuZdrErEOtUKUtCR3dkzLOjrMo6yy_CAqrj-3vwXkyBxBlNykRRHU-a8wR4bsGuchhAnPYqCDVRl1bmljYXN0LnVuY2Vuc29yZWRkbnMub3JnCi9kbnMtcXVlcnk


## uncensoreddns-dk-ipv6

Also known as censurfridns.
DoH, no logs, no filter, unicast hosted in Denmark - https://blog.uncensoreddns.org

sdns://AgYAAAAAAAAAElsyYTAxOjNhMDo1Mzo1Mzo6XaDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5qD39a8ogB6HfX3zor4o_hSfMnSn7SnxNv1QPOOYy6gLDaCzXgWVDPvYJZ-TdboQGHC4w8TA7FpSl1fmMWfp3iap6KBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33zqDmRLppY-M1_nZcuZdrErEOtUKUtCR3dkzLOjrMo6yy_CAqrj-3vwXkyBxBlNykRRHU-a8wR4bsGuchhAnPYqCDVRl1bmljYXN0LnVuY2Vuc29yZWRkbnMub3JnCi9kbnMtcXVlcnk


## uncensoreddns-ipv4

Also known as censurfridns.
DoH, no logs, no filter, anycast - https://blog.uncensoreddns.org

sdns://AgYAAAAAAAAADjkxLjIzOS4xMDAuMTAwoMwQYNOcgym2K2-8fQ1t-TCYabmB5-Y5LVzY-kCPTYDmoPf1ryiAHod9ffOivij-FJ8ydKftKfE2_VA845jLqAsNoLNeBZUM-9gln5N1uhAYcLjDxMDsWlKXV-YxZ-neJqnooEROvWe7g_iAezkh6TiskXi4gr1QqtsRIx8ETPXwjffOoOZEumlj4zX-dly5l2sSsQ61QpS0JHd2TMs6OsyjrLL8ICquP7e_BeTIHEGU3KRFEdT5rzBHhuwa5yGECc9ioINVGWFueWNhc3QudW5jZW5zb3JlZGRucy5vcmcKL2Rucy1xdWVyeQ


## uncensoreddns-ipv6

Also known as censurfridns.
DoH, no logs, no filter, anycast - https://blog.uncensoreddns.org

sdns://AgYAAAAAAAAAEVsyMDAxOjY3YzoyOGE0OjpdoMwQYNOcgym2K2-8fQ1t-TCYabmB5-Y5LVzY-kCPTYDmoPf1ryiAHod9ffOivij-FJ8ydKftKfE2_VA845jLqAsNoLNeBZUM-9gln5N1uhAYcLjDxMDsWlKXV-YxZ-neJqnooEROvWe7g_iAezkh6TiskXi4gr1QqtsRIx8ETPXwjffOoOZEumlj4zX-dly5l2sSsQ61QpS0JHd2TMs6OsyjrLL8ICquP7e_BeTIHEGU3KRFEdT5rzBHhuwa5yGECc9ioINVGWFueWNhc3QudW5jZW5zb3JlZGRucy5vcmcKL2Rucy1xdWVyeQ


## userspace-australia

DNSCrypt in Australia (Brisbane & Melbourne) by UserSpace.
No logs | IPv4 | Filtered

sdns://AQIAAAAAAAAAEDEwMy4xNi4xMzEuNzc6NTQgnRNtOLv4IzxEfkbLFOaHa-ncLImdQiP-pS1jaFY5jlUdMi5kbnNjcnlwdC1jZXJ0LnVzZXJzcGFjZS1ibmU
sdns://AQIAAAAAAAAAEjEwMy4yMzYuMTYyLjExOTo1NCBPr5jCD_2geOTMmS5LgQg_v79pgppTm3vLZhe_oahbgR0yLmRuc2NyeXB0LWNlcnQudXNlcnNwYWNlLW1lbA


## userspace-australia-ipv6

DNSCrypt in Australia (Brisbane & Melbourne) by UserSpace.
No logs | IPv6 | Filtered

sdns://AQIAAAAAAAAAJVsyNDA0Ojk0MDA6MTowOjIxNjozZWZmOmZlZjA6MTgwYV06NTQgnRNtOLv4IzxEfkbLFOaHa-ncLImdQiP-pS1jaFY5jlUdMi5kbnNjcnlwdC1jZXJ0LnVzZXJzcGFjZS1ibmU
sdns://AQIAAAAAAAAAJVsyNDA0Ojk0MDA6MzowOjIxNjozZWZmOmZlZTA6N2Y2OV06NTQgT6-Ywg_9oHjkzJkuS4EIP7-_aYKaU5t7y2YXv6GoW4EdMi5kbnNjcnlwdC1jZXJ0LnVzZXJzcGFjZS1tZWw


## v.dnscrypt.uk-ipv4

DNSCrypt, no logs, uncensored, DNSSEC. Hosted in London UK on Digital Ocean
https://www.dnscrypt.uk

sdns://AQcAAAAAAAAADzEwNC4yMzguMTg2LjE5MiDtST2M6teQZk8GPEe3lZojaS18kDY8nkPMtZF75bQe5R0yLmRuc2NyeXB0LWNlcnQudi5kbnNjcnlwdC51aw


## v.dnscrypt.uk-ipv6

DNSCrypt, no logs, uncensored, DNSSEC. Hosted in London UK on Digital Ocean
https://www.dnscrypt.uk

sdns://AQcAAAAAAAAAKFsyMDAxOjE5ZjA6NzQwMjoxNTc0OjU0MDA6MmZmOmZlNjY6MmNmZl0g7Uk9jOrXkGZPBjxHt5WaI2ktfJA2PJ5DzLWRe-W0HuUdMi5kbnNjcnlwdC1jZXJ0LnYuZG5zY3J5cHQudWs


## wikimedia

Wikimedia DNS (formerly called Wikidough), is a caching, recursive,
public resolver service that is run and managed by the Site
Reliability Engineering (Traffic) team at the Foundation.

Wikimedia DNS helps prevent some surveillance and censorship of our
wikis and other websites by securing DNS lookups.

sdns://AgcAAAAAAAAADjE4NS43MS4xMzguMTM4oMwQYNOcgym2K2-8fQ1t-TCYabmB5-Y5LVzY-kCPTYDmoPf1ryiAHod9ffOivij-FJ8ydKftKfE2_VA845jLqAsNoLNeBZUM-9gln5N1uhAYcLjDxMDsWlKXV-YxZ-neJqnooEROvWe7g_iAezkh6TiskXi4gr1QqtsRIx8ETPXwjffOoOZEumlj4zX-dly5l2sSsQ61QpS0JHd2TMs6OsyjrLL8ICquP7e_BeTIHEGU3KRFEdT5rzBHhuwa5yGECc9ioINVEXdpa2ltZWRpYS1kbnMub3JnCi9kbnMtcXVlcnk


## wikimedia-ipv6

Wikimedia DNS over IPv6.

Wikimedia DNS (formerly called Wikidough), is a caching, recursive,
public resolver service that is run and managed by the Site
Reliability Engineering (Traffic) team at the Foundation.

Wikimedia DNS helps prevent some surveillance and censorship of our
wikis and other websites by securing DNS lookups.

sdns://AgcAAAAAAAAAEVsyMDAxOjY3Yzo5MzA6OjFdoMwQYNOcgym2K2-8fQ1t-TCYabmB5-Y5LVzY-kCPTYDmoPf1ryiAHod9ffOivij-FJ8ydKftKfE2_VA845jLqAsNoLNeBZUM-9gln5N1uhAYcLjDxMDsWlKXV-YxZ-neJqnooEROvWe7g_iAezkh6TiskXi4gr1QqtsRIx8ETPXwjffOoOZEumlj4zX-dly5l2sSsQ61QpS0JHd2TMs6OsyjrLL8ICquP7e_BeTIHEGU3KRFEdT5rzBHhuwa5yGECc9ioINVEXdpa2ltZWRpYS1kbnMub3JnCi9kbnMtcXVlcnk


## yandex

Yandex public DNS server (anycast)

sdns://AgUAAAAAAAAACTc3Ljg4LjguMSCoF6cUD2dwqtorNi96I2e3nkHPSJH1ka3xbdOglmOVkQk3Ny44OC44LjEKL2Rucy1xdWVyeQ
sdns://AgUAAAAAAAAACTc3Ljg4LjguOCCoF6cUD2dwqtorNi96I2e3nkHPSJH1ka3xbdOglmOVkQk3Ny44OC44LjgKL2Rucy1xdWVyeQ


## yandex-ipv6

Yandex public DNS server (anycast IPv6)

sdns://AgUAAAAAAAAAE1syYTAyOjZiODo6ZmVlZDpmZl0gqBenFA9ncKraKzYveiNnt55Bz0iR9ZGt8W3ToJZjlZEJNzcuODguOC4xCi9kbnMtcXVlcnk
sdns://AgUAAAAAAAAAF1syYTAyOjZiODowOjE6OmZlZWQ6ZmZdIKgXpxQPZ3Cq2is2L3ojZ7eeQc9IkfWRrfFt06CWY5WRCTc3Ljg4LjguMQovZG5zLXF1ZXJ5


## yandex-safe

Yandex public DNS server with malware filtering (anycast)

sdns://AgEAAAAAAAAACTc3Ljg4LjguMiCoF6cUD2dwqtorNi96I2e3nkHPSJH1ka3xbdOglmOVkQk3Ny44OC44LjIKL2Rucy1xdWVyeQ
sdns://AgEAAAAAAAAACjc3Ljg4LjguODggqBenFA9ncKraKzYveiNnt55Bz0iR9ZGt8W3ToJZjlZEKNzcuODguOC44OAovZG5zLXF1ZXJ5


## yandex-safe-ipv6

Yandex public DNS server with malware filtering (anycast IPv6)

sdns://AgEAAAAAAAAAFFsyYTAyOjZiODo6ZmVlZDpiYWRdIKgXpxQPZ3Cq2is2L3ojZ7eeQc9IkfWRrfFt06CWY5WRCTc3Ljg4LjguMgovZG5zLXF1ZXJ5
sdns://AgEAAAAAAAAAGFsyYTAyOjZiODowOjE6OmZlZWQ6YmFkXSCoF6cUD2dwqtorNi96I2e3nkHPSJH1ka3xbdOglmOVkQk3Ny44OC44LjIKL2Rucy1xdWVyeQ

