# opennic

Resolvers from the [OpenNIC](https://www.opennic.org/) project.

To use that list, add this to the `[sources]` section of your
`dnscrypt-proxy.toml` configuration file:

    [sources.'opennic']
    urls = ['https://raw.githubusercontent.com/DNSCrypt/dnscrypt-resolvers/master/v2/opennic.md', 'https://download.dnscrypt.info/resolvers-list/v2/opennic.md']
    minisign_key = 'RWQf6LRCGA9i53mlYecO4IzT51TGPpvWucNSCh1CBM0QTaLn73Y7GFO3'
    cache_file = 'opennic.md'

--

## fvz-anyone

Fusl's public primary OpenNIC Tier2 Anycast DNS Resolver

sdns://AQYAAAAAAAAAFDE4NS4xMjEuMTc3LjE3Nzo1MzUzIBpq0KMrTFphppXRU2cNaasWkD-ew_f2TxPlNaMYsiilHTIuZG5zY3J5cHQtY2VydC5kbnNyZWMubWVvLndz

## fvz-anytwo

Fusl's public secondary OpenNIC Tier2 Anycast DNS Resolver

sdns://AQYAAAAAAAAAFDE2OS4yMzkuMjAyLjIwMjo1MzUzIBpq0KMrTFphppXRU2cNaasWkD-ew_f2TxPlNaMYsiilHTIuZG5zY3J5cHQtY2VydC5kbnNyZWMubWVvLndz

## ibksturm

dnscrypt-server (dnscrypt-wrapper - dnsmasq backend), DNSSEC / Non-Logged / Uncensored, OpenNIC
Hosted in Switzerland on a Banana Pi M64 by ibksturm, aka Andreas Ziegler

sdns://AQcAAAAAAAAADzIxNy4xNjIuMjA2LjE3OCB2x0U7IXv1uDMjPD3ypxKY4xhAxu7bxJrlMs0vfeSV0BgyLmRuc2NyeXB0LWNlcnQuaWJrc3R1cm0

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

## publicarray-au2

DNSSEC • OpenNic • Non-logging • Uncensored - hosted on GCP
Maintained by publicarray

sdns://AQcAAAAAAAAAEjM1LjIwMS4yMC4xNzk6ODQ0MyDbu252PopUsAoQmpOFc8eYC4rkr2nWINwVQPMlc8lN8xsyLmRuc2NyeXB0LWNlcnQuZG5zLnNlYnkuaW8

## publicarray-au-doh

DNSSEC • OpenNic • Non-logging • Uncensored - hosted on GCP
Maintained by publicarray

sdns://AgcAAAAAAAAAETM1LjIwMS4yMC4xNzk6NDQzAAtkb2guc2VieS5pbwovZG5zLXF1ZXJ5
