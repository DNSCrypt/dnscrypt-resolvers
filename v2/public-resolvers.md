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
    format = 'v2'

This list was last updated on 2018-01-30

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


## captnemo-in
Non-logging, provided by Nemo
sdns://AQYAAAAAAAAAEjEzOS41OS40OC4yMjI6NDQzNCAFOt_yxaMpFtga2IpneSwwK6rV0oAyleham9IvhoceEBwyLmRuc2NyeXB0LWNlcnQuY2FwdG5lbW8uaW4g


## cisco
Remove your DNS blind spot
sdns://AQAAAAAAAAAAEjIwOC42Ny4yMjAuMjIwOjQ0MyC3NRFAIG8iXT4r2CLX_WkeocM8yNZmjQy-BL-rykP7eRsyLmRuc2NyeXB0LWNlcnQub3BlbmRucy5jb20


## cisco-familyshield
Block websites not suitable for children
sdns://AQAAAAAAAAAAEjIwOC42Ny4yMjAuMTIzOjQ0MyC3NRFAIG8iXT4r2CLX_WkeocM8yNZmjQy-BL-rykP7eRsyLmRuc2NyeXB0LWNlcnQub3BlbmRucy5jb20


## cisco-ipv6
Cisco OpenDNS IPv6 sandbox
sdns://AQAAAAAAAAAAE1syNjIwOjA6Y2NjOjoyXTo0NDMgtzURQCBvIl0-K9gi1_1pHqHDPMjWZo0MvgS_q8pD-3kbMi5kbnNjcnlwdC1jZXJ0Lm9wZW5kbnMuY29t


## cpunks-ru
Cypherpunks.ru public DNS server
sdns://AQYAAAAAAAAAEjc3LjUxLjE4MS4yMDk6NTM1MyAYOMyj2VMKZjQzXVAFvTdYROOXfuhoK2xVKBK9p40umR4yLmRuc2NyeXB0LWNlcnQuY3lwaGVycHVua3MucnU


## cs-caeast
provided by cryptostorm.is
sdns://AQYAAAAAAAAAEjE2Ny4xMTQuODQuMTMyOjQ0MyAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM


## cs-cawest
provided by cryptostorm.is
sdns://AQYAAAAAAAAAEzE2Mi4yMjEuMjA3LjIyODo0NDMgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz


## cs-cfi
provided by cryptostorm.is
sdns://AQYAAAAAAAAAETIxMi44My4xNzUuMzE6NDQzIDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-ch
provided by cryptostorm.is
sdns://AQYAAAAAAAAAETE4NS42MC4xNDcuNzc6NDQzIDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-de
provided by cryptostorm.is
sdns://AQYAAAAAAAAAEDg0LjE2LjI0MC40Mzo0NDMgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz


## cs-de3
provided by cryptostorm.is
sdns://AQYAAAAAAAAAEjg5LjE2My4yMTQuMTc0OjQ0MyAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM


## cs-dk2
provided by cryptostorm.is
sdns://AQYAAAAAAAAAEzE4NS4yMTIuMTY5LjEzOTo0NDMgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz


## cs-es
provided by cryptostorm.is
sdns://AQYAAAAAAAAAEjE4NS4xNDAuMTE0LjUxOjQ0MyAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM


## cs-fi
provided by cryptostorm.is
sdns://AQYAAAAAAAAAEjE4NS4xMTcuMTE4LjIwOjQ0MyAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM


## cs-fr
provided by cryptostorm.is
sdns://AQYAAAAAAAAAETIxMi4xMjkuNDYuODY6NDQzIDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-fr2
provided by cryptostorm.is
sdns://AQYAAAAAAAAAETIxMi4xMjkuNDYuMzI6NDQzIDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-lv
provided by cryptostorm.is
sdns://AQYAAAAAAAAAETgwLjIzMy4xMzQuNTI6NDQzIDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-md
provided by cryptostorm.is
sdns://AQYAAAAAAAAAETE3Ni4xMjMuMy4yNDk6NDQzIDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-nl
provided by cryptostorm.is
sdns://AQYAAAAAAAAAEjIxMy4xNjMuNjQuMjA4OjQ0MyAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM


## cs-pl
provided by cryptostorm.is
sdns://AQYAAAAAAAAADzUuMTMzLjguMTg3OjQ0MyAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM


## cs-pt
provided by cryptostorm.is
sdns://AQYAAAAAAAAAETEwOS43MS40Mi4yMjg6NDQzIDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-ro
provided by cryptostorm.is
sdns://AQYAAAAAAAAAEDUuMjU0Ljk2LjE5NTo0NDMgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz


## cs-rome
provided by cryptostorm.is
sdns://AQYAAAAAAAAAEjE4NS45NC4xOTMuMjM0OjQ0MyAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM


## cs-uk
provided by cryptostorm.is
sdns://AQYAAAAAAAAAETUuMTAxLjEzNy4yNTE6NDQzIDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-useast2
provided by cryptostorm.is
sdns://AQYAAAAAAAAAEDE5OC43LjU4LjIyNzo0NDMgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz


## cs-usnorth
provided by cryptostorm.is
sdns://AQYAAAAAAAAAEjE3My4yMzQuNTYuMTE1OjQ0MyAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM


## cs-ussouth
provided by cryptostorm.is
sdns://AQYAAAAAAAAADzcwLjMyLjM4LjY3OjQ0MyAxM3KtWVYywkFrhy8Jj4Ub3bllKExsvppPGQlkMNupWh4yLmRuc2NyeXB0LWNlcnQuY3J5cHRvc3Rvcm0uaXM


## cs-ussouth2
provided by cryptostorm.is
sdns://AQYAAAAAAAAAETEwOC42Mi4xOS4xMzE6NDQzIDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## cs-uswest
provided by cryptostorm.is
sdns://AQYAAAAAAAAAEDY0LjEyMC41LjI1MTo0NDMgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz


## cs-uswest3
provided by cryptostorm.is
sdns://AQYAAAAAAAAAEzEwNC4yMzguMTk1LjEzOTo0NDMgMTNyrVlWMsJBa4cvCY-FG925ZShMbL6aTxkJZDDbqVoeMi5kbnNjcnlwdC1jZXJ0LmNyeXB0b3N0b3JtLmlz


## cs-uswest5
provided by cryptostorm.is
sdns://AQYAAAAAAAAAETE3My4yMDguOTUuNzU6NDQzIDEzcq1ZVjLCQWuHLwmPhRvduWUoTGy-mk8ZCWQw26laHjIuZG5zY3J5cHQtY2VydC5jcnlwdG9zdG9ybS5pcw


## d0wn-es-ns1
Server provided by Martin 'd0wn' Albus
sdns://AQcAAAAAAAAAETkxLjE0Mi4yMjAuMjk6NDQzIOsJ6FSu2pcFy0ftaerdQVY2U4LFyI2i42kXO1RHdHUFGzIuZG5zY3J5cHQtY2VydC5lcy5kMHduLmJpeg


## d0wn-fr-ns1
Server provided by Martin 'd0wn' Albus
sdns://AQcAAAAAAAAAETE1MS44MC43LjExNToxMDUzIFioItMp68FPvOtFr0LrL1jHlwrT7TE5fR00hjYjdXJRGzIuZG5zY3J5cHQtY2VydC5mci5kMHduLmJpeg


## d0wn-gr-ns1
Server provided by Martin 'd0wn' Albus
sdns://AQcAAAAAAAAAETg1LjI1LjEwNS4xOTM6NDQzILGcC1xI8lj6C-Rn9F9QvH-YX8VEik-8nVV0WjVXAYAJGzIuZG5zY3J5cHQtY2VydC5nci5kMHduLmJpeg


## d0wn-is-ns1
Server provided by Martin 'd0wn' Albus
sdns://AQcAAAAAAAAAEDM3LjIzNS40OS42MTo0NDMgKyiXTgc6azhyKlvh96AlDFCPqAkjj489dthgmCDXstkbMi5kbnNjcnlwdC1jZXJ0LmlzLmQwd24uYml6


## d0wn-is-ns2
Server provided by Martin 'd0wn' Albus
sdns://AQcAAAAAAAAAETkzLjk1LjIyNi4xNjU6NDQzIIRgNKnGMMo8BKxKkBYl8XqHsqy5QYCsSB4sEOjO3hpvHDIuZG5zY3J5cHQtY2VydC5pczIuZDB3bi5iaXo


## d0wn-lv-ns1
Server provided by Martin 'd0wn' Albus
sdns://AQcAAAAAAAAAEDg5LjExMS4xMy42MDo0NDMgmsNrTHrb59bWl7a_FRwVGrKRjF25EhX4uYZZJjOkpeEbMi5kbnNjcnlwdC1jZXJ0Lmx2LmQwd24uYml6


## d0wn-lv-ns2
Server provided by Martin 'd0wn' Albus
sdns://AQcAAAAAAAAAETE4NS44Ni4xNTEuMjg6NDQzILER-Azj4Bw2zHMJlQCeY1HvCAUDMJ2UF3qjjGeRbQzfHDIuZG5zY3J5cHQtY2VydC5sdjIuZDB3bi5iaXo


## d0wn-lv-ns2-ipv6
Server provided by Martin 'd0wn' Albus
sdns://AQcAAAAAAAAAH1syYTAyOjdhYTA6MTIwMTo6ZjYwZToyNzE5XTo0NDMgsRH4DOPgHDbMcwmVAJ5jUe8IBQMwnZQXeqOMZ5FtDN8cMi5kbnNjcnlwdC1jZXJ0Lmx2Mi5kMHduLmJpeg


## d0wn-mx-ns1
Server provided by Martin 'd0wn' Albus
sdns://AQcAAAAAAAAAEzIwMS4xMzEuMTI2LjIxMjo0NDMgmZ5j8A3nwXE6clYlKj8Je-w60o--KHVpXDfoDT1VT9AbMi5kbnNjcnlwdC1jZXJ0Lm14LmQwd24uYml6


## d0wn-nl-ns4
Server provided by Martin 'd0wn' Albus
sdns://AQcAAAAAAAAAETMxLjIyMC40My4xMDU6NDQzICpPQ4lc5jIwtMV2paw4ZBhNEmlJsWD9sJFmq7TOxgZyHDIuZG5zY3J5cHQtY2VydC5ubDQuZDB3bi5iaXo


## d0wn-se-ns1
Server provided by Martin 'd0wn' Albus
sdns://AQcAAAAAAAAAETk1LjIxNS40NC4xMjQ6NDQzIJ1PdivdJPd6ZLR-D_XGk_2gKjnpj-wM7vJSOl-kA8AyGzIuZG5zY3J5cHQtY2VydC5zZS5kMHduLmJpeg


## d0wn-se-ns1-ipv6
Server provided by Martin 'd0wn' Albus
sdns://AQcAAAAAAAAAHlsyYTAyOjdhYTA6MTYxOTo6NGY1MDphNjldOjQ0MyCdT3Yr3ST3emS0fg_1xpP9oCo56Y_sDO7yUjpfpAPAMhsyLmRuc2NyeXB0LWNlcnQuc2UuZDB3bi5iaXo


## d0wn-se-ns2
Server provided by Martin 'd0wn' Albus
sdns://AQcAAAAAAAAAEDMxLjIyMC41LjE4Njo0NDMgxX3WpBeOreLQwLE4pNQHSt9ZFIiHEVys62CjezSS2hUcMi5kbnNjcnlwdC1jZXJ0LnNlMi5kMHduLmJpeg


## d0wn-tz-ns1
Server provided by Martin 'd0wn' Albus
sdns://AQcAAAAAAAAADzQxLjc5LjY5LjEzOjQ0MyDWBhX70UU7oU52ij5cXLOsF0YTMUY9oK-M998E2k11bRsyLmRuc2NyeXB0LWNlcnQudHouZDB3bi5iaXo


## d0wn-tz-ns1-ipv6
Server provided by Martin 'd0wn' Albus
sdns://AQcAAAAAAAAAHFsyYzBmOmZkYTg6NTo6MmVkMTpkMmVjXTo0NDMg1gYV-9FFO6FOdoo-XFyzrBdGEzFGPaCvjPffBNpNdW0bMi5kbnNjcnlwdC1jZXJ0LnR6LmQwd24uYml6


## d0wn-us-ns4
Server provided by Martin 'd0wn' Albus
sdns://AQcAAAAAAAAAEjEwNy4xODEuMTY4LjUyOjQ0MyDzkl1ToxVmwqzystKKaWc5sGYbjO8bOv3oKA2D1OptfRwyLmRuc2NyeXB0LWNlcnQudXM0LmQwd24uYml6


## d0wn-za-ns1
Server provided by Martin 'd0wn' Albus
sdns://AQcAAAAAAAAAETE2OS4yMzkuMTgxLjM6NDQzIPu9DzqvK7G7zY-TJNX2poznIjiQi5CSz_bSv3yewRNoGzIuZG5zY3J5cHQtY2VydC56YS5kMHduLmJpeg


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
sdns://AQcAAAAAAAAAEDc3LjY2Ljg0LjIzMzo0NDMgN0hVheO50Ij9Ja02sDcB9VIM1kiemt1SFFdJVZ8KmVUlMi5kbnNjcnlwdC1jZXJ0LnJlc29sdmVyMi5kbnNjcnlwdC5ldQ


## dnscrypt.eu-dk-ipv6
Free, non-logged, uncensored. Hosted by Netgroup.
sdns://AQcAAAAAAAAAGFsyMDAxOjE0NDg6MjQzOjpkYzJdOjQ0MyA3SFWF47nQiP0lrTawNwH1UgzWSJ6a3VIUV0lVnwqZVSUyLmRuc2NyeXB0LWNlcnQucmVzb2x2ZXIyLmRuc2NyeXB0LmV1


## dnscrypt.eu-nl
Free, non-logged, uncensored. Hosted by RamNode.
sdns://AQcAAAAAAAAAEjE3Ni41Ni4yMzcuMTcxOjQ0MyBnwA8sIcVUgUXdfLRqJxry65aZMUCjCbYrjRZTEYWcZiUyLmRuc2NyeXB0LWNlcnQucmVzb2x2ZXIxLmRuc2NyeXB0LmV1


## dnscrypt.nl-ns0
Public DNSCrypt server in Amsterdam, the Netherlands
sdns://AQcAAAAAAAAAEDQ1Ljc2LjM1LjIxMjo0NDMgTIT7jAURXfpfl8XtAykTcMeKvNZOFd1TqwjecvuESsofMi5kbnNjcnlwdC1jZXJ0Lm5zMC5kbnNjcnlwdC5ubA


## dnscrypt.nl-ns0-ipv6
Public DNSCrypt server in Amsterdam, the Netherlands
sdns://AQcAAAAAAAAAKlsyMDAxOjE5ZjA6NTAwMTozMGE6NTQwMDpmZjpmZTU4OjcxNDBdOjQ0MyBMhPuMBRFd-l-Xxe0DKRNwx4q81k4V3VOrCN5y-4RKyh8yLmRuc2NyeXB0LWNlcnQubnMwLmRuc2NyeXB0Lm5s


## dns.btr.zone
Free of logs and censors by BTR.Zone
sdns://AQcAAAAAAAAAETk0LjEzMC42Ny4xMzg6NDQzIJna5n7EW4ZZDg1PZTN8BbT6zzVoOCwhDSzSOoTvtH0yHDIuZG5zY3J5cHQtY2VydC5kbnMuYnRyLnpvbmU


## freetsa.org
Non-logged/Uncensored provided by freetsa.org
sdns://AQcAAAAAAAAAEzIwNS4xODUuMTE2LjExNjo1NTMg2P-7QuAxvnp5cwtFVo1Jak6Ky1mqg2b9arkeJyp9FuQbMi5kbnNjcnlwdC1jZXJ0LmZyZWV0c2Eub3Jn


## fvz-anyone
Fusl's public primary OpenNIC Tier2 Anycast DNS Resolver
sdns://AQYAAAAAAAAAEzE4NS4xMjEuMTc3LjE3Nzo0NDMgGmrQoytMWmGmldFTZw1pqxaQP57D9_ZPE-U1oxiyKKUdMi5kbnNjcnlwdC1jZXJ0LmRuc3JlYy5tZW8ud3M


## fvz-anytwo
Fusl's public secondary OpenNIC Tier2 Anycast DNS Resolver
sdns://AQYAAAAAAAAAEzE2OS4yMzkuMjAyLjIwMjo0NDMgGmrQoytMWmGmldFTZw1pqxaQP57D9_ZPE-U1oxiyKKUdMi5kbnNjcnlwdC1jZXJ0LmRuc3JlYy5tZW8ud3M


## ipredator
Public DNSCrypt server in Sweden provided by Ipredator.se
sdns://AQcAAAAAAAAAETE5NC4xMzIuMzIuMzI6NDQzIMRMVmqo1kbEMrEE9T0Alhsy3HHPHAS9nrAT5IDnpHgoHDIuZG5zY3J5cHQtY2VydC5pcHJlZGF0b3Iuc2U


## okturtles
For a surveillance-free world. HTTPS is broken. DNSChain fixes it.
sdns://AQYAAAAAAAAAETIzLjIyNi4yMjcuOTM6NDQzIB2FOVPjT6_QBflMb9HM5jXUEZkEDUjRml01C2p8gXPLHTIuZG5zY3J5cHQtY2VydC5va3R1cnRsZXMuY29t


## opennic-famicoman
OpenNIC DNS server with DNSCrypt support
sdns://AQYAAAAAAAAAEzE0Ni4xODUuMTc2LjM2OjUzNTMguI9IYFUXNpaj0r_g7MdhdRmP4BLhAbT-hpwenEw15082Mi5kbnNjcnlwdC1jZXJ0Lm9wZW5uaWMucGVlcjMuZmFtaWNvbWFuLnBoaWxseW1lc2gubmV0


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

sdns://AQcAAAAAAAAAEjIxMi40Ny4yMjguMTM2OjQ0MyDoAbhOpga_sLrAzkNEW7FeumSwL6PEqjGuEGNqB5AyTR8yLmRuc2NyeXB0LWNlcnQuZnIuZG5zY3J5cHQub3Jn


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
sdns://AQcAAAAAAAAAETEwNy4xNzAuNTcuMzQ6NDQzIOmF8RitTjzGX_IlIBiQxvVYt1taUvVrF8_qwQBci5uqJTIuZG5zY3J5cHQtY2VydC5kbnNjcnlwdC52ZW50cmljbGUudXM


## yandex
Yandex public DNS server (anycast)
sdns://AQQAAAAAAAAAEDc3Ljg4LjguNzg6MTUzNTMg04TAccn3RmKvKszVe13MlxTUB7atNgHhrtwG1W1JYyciMi5kbnNjcnlwdC1jZXJ0LmJyb3dzZXIueWFuZGV4Lm5ldA


## google
Google DNS (anycast)
sdns://AgUAAAAAAAAAACDyXGrcc5eNecJ8nomJCJ-q6eCLTEn6bHic0hWGUwYQaA5kbnMuZ29vZ2xlLmNvbQ0vZXhwZXJpbWVudGFs
