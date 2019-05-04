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

sdns://AgcAAAAAAAAAACA-GhoPbFPz6XpJLVcIS1uYBwWe4FerFQWHb9g_2j24OBRpYmtzdHVybS5zeW5vbG9neS5tZQovZG5zLXF1ZXJ5

## ibksturm

dnscrypt-server (nginx-dnscrypt-wrapper-unbound backend), DNSSEC / Non-Logged / Uncensored, OpenNIC and Root DNS-Zone Copy
Hosted in Switzerland by ibksturm, aka Andreas Ziegler

sdns://AQcAAAAAAAAADzIxNy4xNjIuMjA2LjIyMCAbkeo7E6QrYUfz_2_40sQSRfXIf4wu2U1aexB1dmIC2hgyLmRuc2NyeXB0LWNlcnQuaWJrc3R1cm0

## opennic-ethservices

OpenNIC • DNSSEC • 24-hour Logs • AnonymousLogs • NoFilters
Location: Frankfurt, Germany
By ethservices.

sdns://AQUAAAAAAAAAEjk0LjI0Ny40My4yNTQ6NTM1MyDUQmYmXRg576Roac_42Ue6uQtQ664-FvA20PgVt_UIfigyLmRuc2NyeXB0LWNlcnQub3Blbm5pYzEuZXRoLXNlcnZpY2VzLmRl

## opennic-ethservices2

OpenNIC • DNSSEC • 24-hour Logs • AnonymousLogs • NoFilters
Location: Frankfurt, Germany
By ethservices.

sdns://AQYAAAAAAAAAEzE5NS4xMC4xOTUuMTk1OjUzNTMg8hbE05QkH0WdwNiGcxtcLvFewNj3USVp1A-VL0P77HIoMi5kbnNjcnlwdC1jZXJ0Lm9wZW5uaWMyLmV0aC1zZXJ2aWNlcy5kZQ

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

## publicarray-au

DNSSEC • OpenNic • Non-logging • Uncensored - hosted at vultr.com
Maintained by publicarray - https://dns.seby.io

sdns://AQcAAAAAAAAADDQ1Ljc2LjExMy4zMSAIVGh4i6eKXqlF6o9Fg92cgD2WcDvKQJ7v_Wq4XrQsVhsyLmRuc2NyeXB0LWNlcnQuZG5zLnNlYnkuaW8

## publicarray-au-doh

DNSSEC • OpenNic • Non-logging • Uncensored - hosted on GCP
Maintained by publicarray - https://dns.seby.io

sdns://AgcAAAAAAAAADDQ1Ljc2LjExMy4zMSA-GhoPbFPz6XpJLVcIS1uYBwWe4FerFQWHb9g_2j24OBBkb2guc2VieS5pbzo4NDQzCi9kbnMtcXVlcnk

## opennic-R4SAS

Non-logging OpenNIC resolver in France

sdns://AQYAAAAAAAAAETE1MS44MC4yMjIuNzk6NDQzIO4Y9lZnORlvodxu39dnm6mFruwTRnlmovbEga4Fyw3TIDIuZG5zY3J5cHQtY2VydC5vcGVubmljLmkycGQueHl6
