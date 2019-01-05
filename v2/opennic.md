# opennic

Resolvers from the [OpenNIC](https://www.opennic.org/) project.

To use that list, add this to the `[sources]` section of your
`dnscrypt-proxy.toml` configuration file:

    [sources.'opennic']
    urls = ['https://raw.githubusercontent.com/DNSCrypt/dnscrypt-resolvers/master/v2/opennic.md', 'https://download.dnscrypt.info/resolvers-list/v2/opennic.md']
    minisign_key = 'RWQf6LRCGA9i53mlYecO4IzT51TGPpvWucNSCh1CBM0QTaLn73Y7GFO3'
    cache_file = 'opennic.md'

--

## doh-ibksturm

doh-server (nginx - doh-httpproxy - unbound backend), DNSSEC / Non-Logged / Uncensored, OpenNIC and Root DNS-Zone Copy
Hosted in Switzerland on by ibksturm, aka Andreas Ziegler

sdns://AgcAAAAAAAAADzIxNy4xNjIuMjA2LjE3OAAUaWJrc3R1cm0uc3lub2xvZ3kubWUKL2Rucy1xdWVyeQ

## fvz-anyone

Fusl's public primary OpenNIC Tier2 Anycast DNS Resolver

sdns://AQYAAAAAAAAAFDE4NS4xMjEuMTc3LjE3Nzo1MzUzIBpq0KMrTFphppXRU2cNaasWkD-ew_f2TxPlNaMYsiilHTIuZG5zY3J5cHQtY2VydC5kbnNyZWMubWVvLndz

## fvz-anytwo

Fusl's public secondary OpenNIC Tier2 Anycast DNS Resolver

sdns://AQYAAAAAAAAAFDE2OS4yMzkuMjAyLjIwMjo1MzUzIBpq0KMrTFphppXRU2cNaasWkD-ew_f2TxPlNaMYsiilHTIuZG5zY3J5cHQtY2VydC5kbnNyZWMubWVvLndz

## ibksturm

dnscrypt-server (nginx-dnscrypt-wrapper-unbound backend), DNSSEC / Non-Logged / Uncensored, OpenNIC and Root DNS-Zone Copy
Hosted in Switzerland by ibksturm, aka Andreas Ziegler

sdns://AQcAAAAAAAAADzIxNy4xNjIuMjA2LjE3OCD3f8HgokWtd5JfO0look5L3C_Muon2vqbF33rx0-LmDBgyLmRuc2NyeXB0LWNlcnQuaWJrc3R1cm0

## opennic-ethservices

OpenNIC • DNSSEC • 24-hour Logs • AnonymousLogs • NoFilters
Location: Frankfurt, Germany
By ethservices.

sdns://AQUAAAAAAAAAEjk0LjI0Ny40My4yNTQ6NTM1MyDUQmYmXRg576Roac_42Ue6uQtQ664-FvA20PgVt_UIfigyLmRuc2NyeXB0LWNlcnQub3Blbm5pYzEuZXRoLXNlcnZpY2VzLmRl

## opennic-famicoman

OpenNIC • NoLogs • NoFilters
Location: Amsterdam, Netherlands
By famicoman.phillymesh.net

sdns://AQYAAAAAAAAAEzE0Ni4xODUuMTc2LjM2OjUzNTMguI9IYFUXNpaj0r_g7MdhdRmP4BLhAbT-hpwenEw15082Mi5kbnNjcnlwdC1jZXJ0Lm9wZW5uaWMucGVlcjMuZmFtaWNvbWFuLnBoaWxseW1lc2gubmV0

## opennic-tumabox

Public DNS server operated by TumaBox.org

sdns://AQYAAAAAAAAAEjEzMC4yNTUuNzMuOTA6NTM1MyDVkXsRajUxFMI4qpmm6wwofPdoBUGsXb-ooCOeIoxbBg0yLnR1bWFib3gub3Jn

## opennic-tumabox-ipv6

Public DNS server operated by TumaBox.org

sdns://AQYAAAAAAAAAG1syYTAyOmUwMDpmZmZkOjEzOTo6OV06NTM1MyDVkXsRajUxFMI4qpmm6wwofPdoBUGsXb-ooCOeIoxbBg0yLnR1bWFib3gub3Jn

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

## opennic-onic

Public DNS server hosted at MIT (Cambridge, MA, USA) operated by jproulx

sdns://AQcAAAAAAAAADjEyOC41Mi4xMzAuMjA5IBKNsb3hDHyh1SsJH2M-mcGTfRT1-BKwy1s89cvMBHJyIjIuZG5zY3J5cHQtY2VydC5vbmljLmNzYWlsLm1pdC5lZHU

## publicarray-au

DNSSEC • OpenNic • Non-logging • Uncensored - hosted at vultr.com
Maintained by publicarray - https://dns.seby.io

sdns://AQcAAAAAAAAADDQ1Ljc2LjExMy4zMSAIVGh4i6eKXqlF6o9Fg92cgD2WcDvKQJ7v_Wq4XrQsVhsyLmRuc2NyeXB0LWNlcnQuZG5zLnNlYnkuaW8

## publicarray-au-doh

DNSSEC • OpenNic • Non-logging • Uncensored - hosted on GCP
Maintained by publicarray - https://dns.seby.io

sdns://AgcAAAAAAAAADDQ1Ljc2LjExMy4zMSA-GhoPbFPz6XpJLVcIS1uYBwWe4FerFQWHb9g_2j24OBBkb2guc2VieS5pbzo4NDQzCi9kbnMtcXVlcnk

## opennic-userspace

Non-logging OpenNIC resolver in Melbourne, Australia - https://userspace.com.au

sdns://AQYAAAAAAAAAEzEwMy4yMzYuMTYyLjExOTo0NDMgrAN5npeaXgUs0qL88HYBouapH6Vl2B3wcbQae5_HZYgpMi5kbnNjcnlwdC1jZXJ0Lm5zMDMubWVsLnVzZXJzcGFjZS5jb20uYXU

## opennic-userspace-ipv6

Non-logging OpenNIC resolver in Melbourne, Australia - https://userspace.com.au

sdns://AQYAAAAAAAAAJ1syNDA0Ojk0MDA6MzowOjIxNjozZWZmOmZlZTA6N2Y2OV06NTM1MyCsA3mel5peBSzSovzwdgGi5qkfpWXYHfBxtBp7n8dliCkyLmRuc2NyeXB0LWNlcnQubnMwMy5tZWwudXNlcnNwYWNlLmNvbS5hdQ

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
