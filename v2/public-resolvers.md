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
    url = 'http://download.dnscrypt.info/resolvers-list/v2/public-resolvers.md'
    minisign_key = 'RWQf6LRCGA9i53mlYecO4IzT51TGPpvWucNSCh1CBM0QTaLn73Y7GFO3'
    cache_file = 'public-resolvers.md'

This list was last updated on 2018-02-06

--

## adguard-dns-family
Adguard DNS with safesearch and adult content blocking
sdns://AQIAAAAAAAAAFDE3Ni4xMDMuMTMwLjEzMjo1NDQzILgxXdexS27jIKRw3C7Wsao5jMnlhvhdRUXWuMm1AFq6ITIuZG5zY3J5cHQuZmFtaWx5Lm5zMS5hZGd1YXJkLmNvbQ


## adguard-dns
Remove ads and protect your computer from malware
sdns://AQIAAAAAAAAAFDE3Ni4xMDMuMTMwLjEzMDo1NDQzINErR_JS3PLCu_iZEIbq95zkSV2LFsigxDIuUso_OQhzIjIuZG5zY3J5cHQuZGVmYXVsdC5uczEuYWRndWFyZC5jb20


## adguard-dns-family-ipv6
Adguard DNS with safesearch and adult content blocking
sdns://AQIAAAAAAAAAGlsyYTAwOjVhNjA6OmJhZDI6MGZmXTo1NDQzIIwhF6nrwVfW-2QFbwrbwRxdg2c0c8RuJY2bL1fU7jUfITIuZG5zY3J5cHQuZmFtaWx5Lm5zMi5hZGd1YXJkLmNvbQ


## adguard-dns-ipv6
Remove ads and protect your computer from malware
sdns://AQIAAAAAAAAAGVsyYTAwOjVhNjA6OmFkMjowZmZdOjU0NDMggdAC02pMpQxHO3R5ZQ_hLgKzIcthOFYqII5APf3FXpQiMi5kbnNjcnlwdC5kZWZhdWx0Lm5zMi5hZGd1YXJkLmNvbQ


## bn-fr0
Non-logging, uncensored DNS resolver provided by Babylon Network
sdns://AQYAAAAAAAAAETUuMTM1LjY2LjIyMjo1MzUzIIeUBwoUPTXKHKYy57GJMChOrl2v67QB499S6fA3q9GCHzIuZG5zY3J5cHQtY2VydC5iYWJ5bG9uLm5ldHdvcms


## bn-fr0-ipv6
Non-logging, uncensored IPv6 DNS resolver provided by Babylon Network
sdns://AQYAAAAAAAAAHFsyMDAxOjQxZDA6ODo0NDgwOjoyMjJdOjUzNTMgh5QHChQ9NcocpjLnsYkwKE6uXa_rtAHj31Lp8Der0YIfMi5kbnNjcnlwdC1jZXJ0LmJhYnlsb24ubmV0d29yaw


## bn-nl0
Non-logging, uncensored DNS resolver provided by Babylon Network
sdns://AQYAAAAAAAAAEzg3LjI1My4xNTIuMTkwOjUzNTMgh5QHChQ9NcocpjLnsYkwKE6uXa_rtAHj31Lp8Der0YIfMi5kbnNjcnlwdC1jZXJ0LmJhYnlsb24ubmV0d29yaw


## cisco
Remove your DNS blind spot
sdns://AQAAAAAAAAAADjIwOC42Ny4yMjAuMjIwILc1EUAgbyJdPivYItf9aR6hwzzI1maNDL4Ev6vKQ_t5GzIuZG5zY3J5cHQtY2VydC5vcGVuZG5zLmNvbQ


## cisco-familyshield
Block websites not suitable for children
sdns://AQAAAAAAAAAADjIwOC42Ny4yMjAuMTIzILc1EUAgbyJdPivYItf9aR6hwzzI1maNDL4Ev6vKQ_t5GzIuZG5zY3J5cHQtY2VydC5vcGVuZG5zLmNvbQ


## cisco-ipv6
Cisco OpenDNS IPv6 sandbox
sdns://AQAAAAAAAAAAD1syNjIwOjA6Y2NjOjoyXSC3NRFAIG8iXT4r2CLX_WkeocM8yNZmjQy-BL-rykP7eRsyLmRuc2NyeXB0LWNlcnQub3BlbmRucy5jb20


## comodo-02
Comodo Dome Shield (anycast) - https://cdome.comodo.com/shield/
sdns://AQAAAAAAAAAACjguMjAuMjQ3LjIg0sJUqpYcHsoXmZb1X7yAHwg2xyN5q1J-zaiGG-Dgs7AoMi5kbnNjcnlwdC1jZXJ0LnNoaWVsZC0yLmRuc2J5Y29tb2RvLmNvbQ


## cpunks-ru
Cypherpunks.ru public DNS server
sdns://AQYAAAAAAAAAEjc3LjUxLjE4MS4yMDk6NTM1MyAYOMyj2VMKZjQzXVAFvTdYROOXfuhoK2xVKBK9p40umR4yLmRuc2NyeXB0LWNlcnQuY3lwaGVycHVua3MucnU


## cs-caeast
provided by cryptostorm.is
sdns://AQYAAAAAAAAADjE2Ny4xMTQuODQuMTMyIDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-cawest
provided by cryptostorm.is
sdns://AQYAAAAAAAAADzE2Mi4yMjEuMjA3LjIyOCAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM


## cs-cfi
provided by cryptostorm.is
sdns://AQYAAAAAAAAADTIxMi44My4xNzUuMzEgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz


## cs-ch
provided by cryptostorm.is
sdns://AQYAAAAAAAAADTE4NS42MC4xNDcuNzcgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz


## cs-de
provided by cryptostorm.is
sdns://AQYAAAAAAAAADDg0LjE2LjI0MC40MyAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM


## cs-de3
provided by cryptostorm.is
sdns://AQYAAAAAAAAADjg5LjE2My4yMTQuMTc0IDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-dk2
provided by cryptostorm.is
sdns://AQYAAAAAAAAADzE4NS4yMTIuMTY5LjEzOSAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM


## cs-es
provided by cryptostorm.is
sdns://AQYAAAAAAAAADjE4NS4xNDAuMTE0LjUxIDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-fi
provided by cryptostorm.is
sdns://AQYAAAAAAAAADjE4NS4xMTcuMTE4LjIwIDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-fr
provided by cryptostorm.is
sdns://AQYAAAAAAAAADTIxMi4xMjkuNDYuODYgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz


## cs-fr2
provided by cryptostorm.is
sdns://AQYAAAAAAAAADTIxMi4xMjkuNDYuMzIgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz


## cs-lv
provided by cryptostorm.is
sdns://AQYAAAAAAAAADTgwLjIzMy4xMzQuNTIgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz


## cs-md
provided by cryptostorm.is
sdns://AQYAAAAAAAAADTE3Ni4xMjMuMy4yNDkgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz


## cs-nl
provided by cryptostorm.is
sdns://AQYAAAAAAAAADjIxMy4xNjMuNjQuMjA4IDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-pl
provided by cryptostorm.is
sdns://AQYAAAAAAAAACzUuMTMzLjguMTg3IDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-pt
provided by cryptostorm.is
sdns://AQYAAAAAAAAADTEwOS43MS40Mi4yMjggMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz


## cs-ro
provided by cryptostorm.is
sdns://AQYAAAAAAAAADDUuMjU0Ljk2LjE5NSAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM


## cs-rome
provided by cryptostorm.is
sdns://AQYAAAAAAAAADjE4NS45NC4xOTMuMjM0IDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-uk
provided by cryptostorm.is
sdns://AQYAAAAAAAAADTUuMTAxLjEzNy4yNTEgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz


## cs-useast2
provided by cryptostorm.is
sdns://AQYAAAAAAAAADDE5OC43LjU4LjIyNyAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM


## cs-usnorth
provided by cryptostorm.is
sdns://AQYAAAAAAAAADjE3My4yMzQuNTYuMTE1IDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-ussouth
provided by cryptostorm.is
sdns://AQYAAAAAAAAACzcwLjMyLjM4LjY3IDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-ussouth2
provided by cryptostorm.is
sdns://AQYAAAAAAAAADTEwOC42Mi4xOS4xMzEgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz


## cs-uswest
provided by cryptostorm.is
sdns://AQYAAAAAAAAADDY0LjEyMC41LjI1MSAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM


## cs-uswest3
provided by cryptostorm.is
sdns://AQYAAAAAAAAADzEwNC4yMzguMTk1LjEzOSAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM


## cs-uswest5
provided by cryptostorm.is
sdns://AQYAAAAAAAAADTE3My4yMDguOTUuNzUgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz


## d0wn-es-ns1
Server provided by Martin 'd0wn' Albus
sdns://AQcAAAAAAAAADTkxLjE0Mi4yMjAuMjkg6wnoVK7alwXLR-1p6t1BVjZTgsXIjaLjaRc7VEd0dQUbMi5kbnNjcnlwdC1jZXJ0LmVzLmQwd24uYml6


## d0wn-fr-ns1
Server provided by Martin 'd0wn' Albus
sdns://AQcAAAAAAAAAETE1MS44MC43LjExNToxMDUzIFioItMp68FPvOtFr0LrL1jHlwrT7TE5fR00hjYjdXJRGzIuZG5zY3J5cHQtY2VydC5mci5kMHduLmJpeg


## d0wn-gr-ns1
Server provided by Martin 'd0wn' Albus
sdns://AQcAAAAAAAAADTg1LjI1LjEwNS4xOTMgsZwLXEjyWPoL5Gf0X1C8f5hfxUSKT7ydVXRaNVcBgAkbMi5kbnNjcnlwdC1jZXJ0LmdyLmQwd24uYml6


## d0wn-is-ns1
Server provided by Martin 'd0wn' Albus
sdns://AQcAAAAAAAAADDM3LjIzNS40OS42MSArKJdOBzprOHIqW-H3oCUMUI-oCSOPjz122GCYINey2RsyLmRuc2NyeXB0LWNlcnQuaXMuZDB3bi5iaXo


## d0wn-is-ns2
Server provided by Martin 'd0wn' Albus
sdns://AQcAAAAAAAAADTkzLjk1LjIyNi4xNjUghGA0qcYwyjwErEqQFiXxeoeyrLlBgKxIHiwQ6M7eGm8cMi5kbnNjcnlwdC1jZXJ0LmlzMi5kMHduLmJpeg


## d0wn-lv-ns1
Server provided by Martin 'd0wn' Albus
sdns://AQcAAAAAAAAADDg5LjExMS4xMy42MCCaw2tMetvn1taXtr8VHBUaspGMXbkSFfi5hlkmM6Sl4RsyLmRuc2NyeXB0LWNlcnQubHYuZDB3bi5iaXo


## d0wn-lv-ns2
Server provided by Martin 'd0wn' Albus
sdns://AQcAAAAAAAAADTE4NS44Ni4xNTEuMjggsRH4DOPgHDbMcwmVAJ5jUe8IBQMwnZQXeqOMZ5FtDN8cMi5kbnNjcnlwdC1jZXJ0Lmx2Mi5kMHduLmJpeg


## d0wn-lv-ns2-ipv6
Server provided by Martin 'd0wn' Albus
sdns://AQcAAAAAAAAAG1syYTAyOjdhYTA6MTIwMTo6ZjYwZToyNzE5XSCxEfgM4-AcNsxzCZUAnmNR7wgFAzCdlBd6o4xnkW0M3xwyLmRuc2NyeXB0LWNlcnQubHYyLmQwd24uYml6


## d0wn-mx-ns1
Server provided by Martin 'd0wn' Albus
sdns://AQcAAAAAAAAADzIwMS4xMzEuMTI2LjIxMiCZnmPwDefBcTpyViUqPwl77DrSj74odWlcN-gNPVVP0BsyLmRuc2NyeXB0LWNlcnQubXguZDB3bi5iaXo


## d0wn-nl-ns4
Server provided by Martin 'd0wn' Albus
sdns://AQcAAAAAAAAADTMxLjIyMC40My4xMDUgKk9DiVzmMjC0xXalrDhkGE0SaUmxYP2wkWartM7GBnIcMi5kbnNjcnlwdC1jZXJ0Lm5sNC5kMHduLmJpeg


## d0wn-se-ns1
Server provided by Martin 'd0wn' Albus
sdns://AQcAAAAAAAAADTk1LjIxNS40NC4xMjQgnU92K90k93pktH4P9caT_aAqOemP7Azu8lI6X6QDwDIbMi5kbnNjcnlwdC1jZXJ0LnNlLmQwd24uYml6


## d0wn-se-ns1-ipv6
Server provided by Martin 'd0wn' Albus
sdns://AQcAAAAAAAAAGlsyYTAyOjdhYTA6MTYxOTo6NGY1MDphNjldIJ1PdivdJPd6ZLR-D_XGk_2gKjnpj-wM7vJSOl-kA8AyGzIuZG5zY3J5cHQtY2VydC5zZS5kMHduLmJpeg


## d0wn-se-ns2
Server provided by Martin 'd0wn' Albus
sdns://AQcAAAAAAAAADDMxLjIyMC41LjE4NiDFfdakF46t4tDAsTik1AdK31kUiIcRXKzrYKN7NJLaFRwyLmRuc2NyeXB0LWNlcnQuc2UyLmQwd24uYml6


## d0wn-tz-ns1
Server provided by Martin 'd0wn' Albus
sdns://AQcAAAAAAAAACzQxLjc5LjY5LjEzINYGFfvRRTuhTnaKPlxcs6wXRhMxRj2gr4z33wTaTXVtGzIuZG5zY3J5cHQtY2VydC50ei5kMHduLmJpeg


## d0wn-tz-ns1-ipv6
Server provided by Martin 'd0wn' Albus
sdns://AQcAAAAAAAAAGFsyYzBmOmZkYTg6NTo6MmVkMTpkMmVjXSDWBhX70UU7oU52ij5cXLOsF0YTMUY9oK-M998E2k11bRsyLmRuc2NyeXB0LWNlcnQudHouZDB3bi5iaXo


## d0wn-us-ns4
Server provided by Martin 'd0wn' Albus
sdns://AQcAAAAAAAAADjEwNy4xODEuMTY4LjUyIPOSXVOjFWbCrPKy0oppZzmwZhuM7xs6_egoDYPU6m19HDIuZG5zY3J5cHQtY2VydC51czQuZDB3bi5iaXo


## d0wn-za-ns1
Server provided by Martin 'd0wn' Albus
sdns://AQcAAAAAAAAADTE2OS4yMzkuMTgxLjMg-70POq8rsbvNj5Mk1famjOciOJCLkJLP9tK_fJ7BE2gbMi5kbnNjcnlwdC1jZXJ0LnphLmQwd24uYml6


## de.dnsmaschine.net
DNSSEC/Non-logged/Uncensored
Hosted by vultr.com (Frankfurt Germany)
sdns://AQcAAAAAAAAAEzIwOS4yNTAuMjM1LjE3MDo0NDMgz0wbvISl_NVCSe0wDJMS79BAFZoWth1djmhuzv_n3KAiMi5kbnNjcnlwdC1jZXJ0LmRlLmRuc21hc2NoaW5lLm5ldA


## dnscrypt.ca-1
Uncensored DNSSEC validating and log-free
sdns://AQcAAAAAAAAAFDE5OS4xNjcuMTMwLjExODo1MzUzIHT3RVUXvCb3EXflbXKTJ4hscpFbP0YoMD-RDEfDjoJ5HTIuZG5zY3J5cHQtY2VydC5kbnNjcnlwdC5jYS0x


## dnscrypt.ca-2
Uncensored DNSSEC validating and log-free
sdns://AQcAAAAAAAAAFDE5OS4xNjcuMTI4LjExMjo1MzUzIEPVLIJZIpbC22-NSM4iT9zHJibhBvbjiGGT-gCQKWMbHTIuZG5zY3J5cHQtY2VydC5kbnNjcnlwdC5jYS0y


## dnscrypt.ca-3
Uncensored DNSSEC validating and log-free
sdns://AQcAAAAAAAAAEzY5LjE2NS4yMjAuMjIxOjUzNTMgH5iH2E7cxcehI8hiCm1La8uRRaGZT2v2ZkPrLJ4z3lsdMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeXB0LmNhLTM


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
Public DNSCrypt server in Amsterdam, the Netherlands
sdns://AQcAAAAAAAAADDQ1Ljc2LjM1LjIxMiBMhPuMBRFd-l-Xxe0DKRNwx4q81k4V3VOrCN5y-4RKyh8yLmRuc2NyeXB0LWNlcnQubnMwLmRuc2NyeXB0Lm5s


## dnscrypt.nl-ns0-ipv6
Public DNSCrypt server in Amsterdam, the Netherlands
sdns://AQcAAAAAAAAAJlsyMDAxOjE5ZjA6NTAwMTozMGE6NTQwMDpmZjpmZTU4OjcxNDBdIEyE-4wFEV36X5fF7QMpE3DHirzWThXdU6sI3nL7hErKHzIuZG5zY3J5cHQtY2VydC5uczAuZG5zY3J5cHQubmw


## doh-dnscrypt-info
Experimental DNS-over-HTTPS server, currently served by Scaleway and OVH.
Maintained by Frank Denis.

sdns://AgcAAAAAAAAADTM3LjU5LjIzOC4yMTMgwzRA_TfjYt0RwSHqBHwj7OM-D_x-CDgqIHeJHIoN1P0UZG9oLmZyLmRuc2NyeXB0LmluZm8KL2Rucy1xdWVyeQ


## doh-crypto-sx
Experimental DNS-over-HTTPS server, globally cached via Cloudflare.
Maintained by Frank Denis.

sdns://AgcAAAAAAAAAACAd2FCKjFZZBDl8eGRR4I9XYTzzyKcj9vN5_Uw4WLbznw1kb2guY3J5cHRvLnN4Ci9kbnMtcXVlcnk


## freetsa.org
Non-logged/Uncensored provided by freetsa.org
sdns://AQcAAAAAAAAAEzIwNS4xODUuMTE2LjExNjo1NTMg2P-7QuAxvnp5cwtFVo1Jak6Ky1mqg2b9arkeJyp9FuQbMi5kbnNjcnlwdC1jZXJ0LmZyZWV0c2Eub3Jn


## fvz-anyone
Fusl's public primary OpenNIC Tier2 Anycast DNS Resolver
sdns://AQYAAAAAAAAAFDE4NS4xMjEuMTc3LjE3Nzo1MzUzIBpq0KMrTFphppXRU2cNaasWkD-ew_f2TxPlNaMYsiilHTIuZG5zY3J5cHQtY2VydC5kbnNyZWMubWVvLndz


## fvz-anytwo
Fusl's public secondary OpenNIC Tier2 Anycast DNS Resolver
sdns://AQYAAAAAAAAAFDE2OS4yMzkuMjAyLjIwMjo1MzUzIBpq0KMrTFphppXRU2cNaasWkD-ew_f2TxPlNaMYsiilHTIuZG5zY3J5cHQtY2VydC5kbnNyZWMubWVvLndz


## google
Google DNS (anycast)
sdns://AgUAAAAAAAAAACDyXGrcc5eNecJ8nomJCJ-q6eCLTEn6bHic0hWGUwYQaA5kbnMuZ29vZ2xlLmNvbQ0vZXhwZXJpbWVudGFs


## ipredator
Public DNSCrypt server in Sweden provided by Ipredator.se
sdns://AQcAAAAAAAAADTE5NC4xMzIuMzIuMzIgxExWaqjWRsQysQT1PQCWGzLccc8cBL2esBPkgOekeCgcMi5kbnNjcnlwdC1jZXJ0LmlwcmVkYXRvci5zZQ


## opennic-famicoman
OpenNIC DNS server with DNSCrypt support
sdns://AQYAAAAAAAAAEzE0Ni4xODUuMTc2LjM2OjUzNTMguI9IYFUXNpaj0r_g7MdhdRmP4BLhAbT-hpwenEw15082Mi5kbnNjcnlwdC1jZXJ0Lm9wZW5uaWMucGVlcjMuZmFtaWNvbWFuLnBoaWxseW1lc2gubmV0


## opennic-luggs
Public DNS server in Canada operated by Luggs
sdns://AQYAAAAAAAAADTE0Mi40LjIwNC4xMTEgHBl5MxvoI8zPCJp5BpN-XDQQKlasf2Jw4EYlsu3bBOMfMi5kbnNjcnlwdC1jZXJ0Lm5zMy5jYS5sdWdncy5jbw


## opennic-luggs-ipv6
Public DNS server in Canada operated by Luggs
sdns://AQYAAAAAAAAAIVsyNjA3OjUzMDA6MTIwOmE4YToxNDI6NDoyMDQ6MTExXSAcGXkzG-gjzM8ImnkGk35cNBAqVqx_YnDgRiWy7dsE4x8yLmRuc2NyeXB0LWNlcnQubnMzLmNhLmx1Z2dzLmNv


## opennic-onic
Public DNS server hosted at MIT (Cambridge, MA, USA) operated by jproulx
sdns://AQcAAAAAAAAADjEyOC41Mi4xMzAuMjA5IBKNsb3hDHyh1SsJH2M-mcGTfRT1-BKwy1s89cvMBHJyIjIuZG5zY3J5cHQtY2VydC5vbmljLmNzYWlsLm1pdC5lZHU


## opennic-onic-ipv6
Public DNS server hosted at MIT (Cambridge, MA, USA) operated by jproulx
sdns://AQcAAAAAAAAAFlsyMDAxOjQ3MDoxZjA2OjEwYjo6Ml0gEo2xveEMfKHVKwkfYz6ZwZN9FPX4ErDLWzz1y8wEcnIiMi5kbnNjcnlwdC1jZXJ0Lm9uaWMuY3NhaWwubWl0LmVkdQ


## opennic-tumabox
Public DNS server operated by TumaBox.org
sdns://AQYAAAAAAAAAEjEzMC4yNTUuNzMuOTA6NTM1MyDVkXsRajUxFMI4qpmm6wwofPdoBUGsXb-ooCOeIoxbBg0yLnR1bWFib3gub3Jn


## opennic-tumabox-ipv6
Public DNS server operated by TumaBox.org
sdns://AQYAAAAAAAAAG1syYTAyOmUwMDpmZmZkOjEzOTo6OV06NTM1MyDVkXsRajUxFMI4qpmm6wwofPdoBUGsXb-ooCOeIoxbBg0yLnR1bWFib3gub3Jn


## scaleway-fr
DNSSEC/Non-logged/Uncensored - ARM server donated by Scaleway.com
Maintained by Frank Denis - https://fr.dnscrypt.info
This server used to be called `dnscrypt.org-fr`.

sdns://AQcAAAAAAAAADjIxMi40Ny4yMjguMTM2IOgBuE6mBr-wusDOQ0RbsV66ZLAvo8SqMa4QY2oHkDJNHzIuZG5zY3J5cHQtY2VydC5mci5kbnNjcnlwdC5vcmc


## securedns
Uncensored and no logging
sdns://AQcAAAAAAAAAEzE0Ni4xODUuMTY3LjQzOjUzNTMgs6WXaRRXWwSJ4Z-unEPmefryjFcYlwAxf3u0likfsJUcMi5kbnNjcnlwdC1jZXJ0LnNlY3VyZWRucy5ldQ


## securedns-ipv6
Uncensored and no logging
sdns://AQcAAAAAAAAAIVsyYTAzOmIwYzA6MDoxMDEwOjplOWE6MzAwMV06NTM1MyCzpZdpFFdbBInhn66cQ-Z5-vKMVxiXADF_e7SWKR-wlRwyLmRuc2NyeXB0LWNlcnQuc2VjdXJlZG5zLmV1


## soltysiak
Public DNSCrypt server in Poland
sdns://AQcAAAAAAAAAFDE3OC4yMTYuMjAxLjIyMjoyMDUzICXE4YgpFUaXj5wrvbanr6QB7aBRBQhdUwPnGSjAZo8hHTIuZG5zY3J5cHQtY2VydC5zb2x0eXNpYWsuY29t


## soltysiak-ipv6
Public DNSCrypt server in Poland
sdns://AQcAAAAAAAAAGVsyMDAxOjQ3MDo3MDo0ZmY6OjJdOjIwNTMgJcThiCkVRpePnCu9tqevpAHtoFEFCF1TA-cZKMBmjyEdMi5kbnNjcnlwdC1jZXJ0LnNvbHR5c2lhay5jb20


## ventricle.us
Public DNSCrypt resolver provided by Jacob Henner
sdns://AQcAAAAAAAAADTEwNy4xNzAuNTcuMzQg6YXxGK1OPMZf8iUgGJDG9Vi3W1pS9WsXz-rBAFyLm6olMi5kbnNjcnlwdC1jZXJ0LmRuc2NyeXB0LnZlbnRyaWNsZS51cw


## yandex
Yandex public DNS server (anycast)
sdns://AQQAAAAAAAAAEDc3Ljg4LjguNzg6MTUzNTMg04TAccn3RmKvKszVe13MlxTUB7atNgHhrtwG1W1JYyciMi5kbnNjcnlwdC1jZXJ0LmJyb3dzZXIueWFuZGV4Lm5ldA
