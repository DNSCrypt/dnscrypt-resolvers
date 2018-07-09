# opennic

Resolvers from the [OpenNIC](https://www.opennic.org/) project.

To use that list, add this to the `[sources]` section of your
`dnscrypt-proxy.toml` configuration file:

    [sources.'opennic']
    urls = ['https://raw.githubusercontent.com/DNSCrypt/dnscrypt-resolvers/master/v2/opennic.md', 'https://download.dnscrypt.info/resolvers-list/v2/opennic.md']
    cache_file = 'opennic.md'
    minisign_key = "RWQf6LRCGA9i53mlYecO4IzT51TGPpvWucNSCh1CBM0QTaLn73Y7GFO3"

--

## fvz-anyone
OpenNIC • Anycast • NoLogs • NoFilters
Location: Anycast
Fusl's public primary OpenNIC Tier2 Anycast DNS Resolver
sdns://AQYAAAAAAAAAFDE4NS4xMjEuMTc3LjE3Nzo1MzUzIBpq0KMrTFphppXRU2cNaasWkD-ew_f2TxPlNaMYsiilHTIuZG5zY3J5cHQtY2VydC5kbnNyZWMubWVvLndz


## fvz-anytwo
OpenNIC • Anycast • NoLogs • NoFilters
Location: Anycast
Fusl's public secondary OpenNIC Tier2 Anycast DNS Resolver
sdns://AQYAAAAAAAAAFDE2OS4yMzkuMjAyLjIwMjo1MzUzIBpq0KMrTFphppXRU2cNaasWkD-ew_f2TxPlNaMYsiilHTIuZG5zY3J5cHQtY2VydC5kbnNyZWMubWVvLndz


## opennic-luggs
OpenNIC • NoLogs • NoFilters • Whitelisting
Location: Quebec, Canada
By Luggs
sdns://AQYAAAAAAAAAETE0Mi40LjIwNC4xMTE6NDQzIBwZeTMb6CPMzwiaeQaTflw0ECpWrH9icOBGJbLt2wTjHzIuZG5zY3J5cHQtY2VydC5uczMuY2EubHVnZ3MuY28


## opennic-luggs-ipv6
OpenNIC • NoLogs • NoFilters • Whitelisting
Location: Quebec, Canada
By Luggs
sdns://AQYAAAAAAAAAJVsyNjA3OjUzMDA6MTIwOmE4YToxNDI6NDoyMDQ6MTExXTo0NDMgHBl5MxvoI8zPCJp5BpN-XDQQKlasf2Jw4EYlsu3bBOMfMi5kbnNjcnlwdC1jZXJ0Lm5zMy5jYS5sdWdncy5jbw


## opennic-luggs2
OpenNIC • NoLogs • NoFilters • Whitelisting
Location: Quebec, Canada
By Luggs
sdns://AQYAAAAAAAAAEDE0Mi40LjIwNS40Nzo0NDMgvL-34FDBPaJCLACwsaya1kjFwmS8thcLiD1xishuugkfMi5kbnNjcnlwdC1jZXJ0Lm5zNC5jYS5sdWdncy5jbw


## opennic-luggs2-ipv6
OpenNIC • NoLogs • NoFilters • Whitelisting
Location: Quebec, Canada
By Luggs
sdns://AQYAAAAAAAAAJFsyNjA3OjUzMDA6MTIwOmE4YToxNDI6NDoyMDU6NDddOjQ0MyC8v7fgUME9okIsALCxrJrWSMXCZLy2FwuIPXGKyG66CR8yLmRuc2NyeXB0LWNlcnQubnM0LmNhLmx1Z2dzLmNv


## opennic-ethservices
OpenNIC • DNSSEC • 24-hour Logs • AnonymousLogs • NoFilters
Location: Frankfurt, Germany
By ethservices
sdns://AQUAAAAAAAAAEjk0LjI0Ny40My4yNTQ6NTM1MyDUQmYmXRg576Roac_42Ue6uQtQ664-FvA20PgVt_UIfigyLmRuc2NyeXB0LWNlcnQub3Blbm5pYzEuZXRoLXNlcnZpY2VzLmRl


## opennic-famicoman
OpenNIC • NoLogs • NoFilters
Location: Amsterdam, Netherlands
By famicoman.phillymesh.net
sdns://AQYAAAAAAAAAEzE0Ni4xODUuMTc2LjM2OjUzNTMguI9IYFUXNpaj0r_g7MdhdRmP4BLhAbT-hpwenEw15082Mi5kbnNjcnlwdC1jZXJ0Lm9wZW5uaWMucGVlcjMuZmFtaWNvbWFuLnBoaWxseW1lc2gubmV0


## opennic-famicoman-ipv6
OpenNIC • NoLogs • NoFilters
Location: Amsterdam, Netherlands
By famicoman.phillymesh.net
sdns://AQYAAAAAAAAAIVsyYTAzOmIwYzA6MDoxMDEwOjoxYTc6YzAwMV06NTM1MyC4j0hgVRc2lqPSv-Dsx2F1GY_gEuEBtP6GnB6cTDXnTzYyLmRuc2NyeXB0LWNlcnQub3Blbm5pYy5wZWVyMy5mYW1pY29tYW4ucGhpbGx5bWVzaC5uZXQ


## opennic-onic
OpenNIC • DNSSEC • NoLogs • NoFilters
Location: Cambridge, Massachusetts, USA
By jproulx
sdns://AQcAAAAAAAAAEjEyOC41Mi4xMzAuMjA5OjQ0MyASjbG94Qx8odUrCR9jPpnBk30U9fgSsMtbPPXLzARyciIyLmRuc2NyeXB0LWNlcnQub25pYy5jc2FpbC5taXQuZWR1


## ibksturm
OpenNIC • DNSSEC • NoLogs • NoFilters
dnscrypt-server (dnscrypt-wrapper - dnsmasq backend)
Hosted in Switzerland on a Banana Pi M64 by ibksturm, aka Andreas Ziegler
sdns://AQcAAAAAAAAADzIxNy4xNjIuMjA2LjE3OCB2x0U7IXv1uDMjPD3ypxKY4xhAxu7bxJrlMs0vfeSV0BgyLmRuc2NyeXB0LWNlcnQuaWJrc3R1cm0


## publicarray-au
OpenNIC • DNSSEC • NoLogs • NoFilters
Location: Australia (vultr.com)
Maintained by publicarray - https://dns.seby.io
sdns://AQcAAAAAAAAADDQ1Ljc2LjExMy4zMSAIVGh4i6eKXqlF6o9Fg92cgD2WcDvKQJ7v_Wq4XrQsVhsyLmRuc2NyeXB0LWNlcnQuZG5zLnNlYnkuaW8


## publicarray-au2
OpenNIC • DNSSEC • NoLogs • NoFilters
Location: Australia (GCP)
Maintained by publicarray
sdns://AQcAAAAAAAAAEjM1LjIwMS4yMC4xNzk6ODQ0MyDbu252PopUsAoQmpOFc8eYC4rkr2nWINwVQPMlc8lN8xsyLmRuc2NyeXB0LWNlcnQuZG5zLnNlYnkuaW8


## publicarray-au-doh
OpenNIC • DNSSEC • NoLogs • NoFilters
Location: Australia (GCP)
Maintained by publicarray
sdns://AgcAAAAAAAAAETM1LjIwMS4yMC4xNzk6NDQzAAtkb2guc2VieS5pbwovZG5zLXF1ZXJ5
