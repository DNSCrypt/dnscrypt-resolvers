# opennic

Resolvers from the [OpenNIC](https://www.opennic.org/) project.

To use that list, add this to the `[sources]` section of your
`dnscrypt-proxy.toml` configuration file:

    [sources.'opennic']
    urls = ['https://raw.githubusercontent.com/DNSCrypt/dnscrypt-resolvers/master/v3/opennic.md', 'https://download.dnscrypt.info/resolvers-list/v3/opennic.md']
    minisign_key = 'RWQf6LRCGA9i53mlYecO4IzT51TGPpvWucNSCh1CBM0QTaLn73Y7GFO3'
    cache_file = 'opennic.md'

--


## doh-ibksturm

doh-server (nginx - doh-httpproxy - unbound backend), DNSSEC / Non-Logged / Uncensored, OpenNIC and Root DNS-Zone Copy
Hosted in Switzerland on by ibksturm, aka Andreas Ziegler

sdns://AgcAAAAAAAAAAKA-GhoPbFPz6XpJLVcIS1uYBwWe4FerFQWHb9g_2j24OCAyhv9lpl-vMghe6hOIw3OLp-N4c8kGzOPEootMwqWJiBRpYmtzdHVybS5zeW5vbG9neS5tZQovZG5zLXF1ZXJ5


## ibksturm

dnscrypt-server (nginx - encrypted-dns - unbound backend), DNSSEC / Non-Logged / Uncensored, OpenNIC and Root DNS-Zone Copy
Hosted in Switzerland by ibksturm, aka Andreas Ziegler

sdns://AQcAAAAAAAAAEjgzLjc4LjEyMC4yMjc6ODQ0MyDBz1dQALBbwmxiH17PmqJWCs6_AH6-yzp_9LIN4LQ57hgyLmRuc2NyeXB0LWNlcnQuaWJrc3R1cm0


## ibksturm-ipv6

dnscrypt-server (nginx - encrypted-dns - unbound backend), DNSSEC / Non-Logged / Uncensored, OpenNIC and Root DNS-Zone Copy
Hosted in Switzerland by ibksturm, aka Andreas Ziegler

sdns://AQcAAAAAAAAALlsyYTAyOjEyMDU6MzRlNzo4ZTMwOmIyNmU6YmZmZjpmZTFkOmUxOWJdOjg0NDMgwc9XUACwW8JsYh9ez5qiVgrOvwB-vss6f_SyDeC0Oe4YMi5kbnNjcnlwdC1jZXJ0Lmlia3N0dXJt


## opennic-R4SAS

DNSSEC - OpenNIC - Non-logging - Uncensored - hosted on ovh.com
Location: Gravelines, France.
Maintained by R4SAS.

sdns://AQcAAAAAAAAADTE1MS44MC4yMjIuNzkgqdYyOk8lgAkmGXUVAs4jHh922d53bIfGu7KKDv_bDk4gMi5kbnNjcnlwdC1jZXJ0Lm9wZW5uaWMuaTJwZC54eXo


## opennic-R4SAS-ipv6

DNSSEC - OpenNIC - Non-logging - Uncensored - hosted on ovh.com
Location: Gravelines, France.
Maintained by R4SAS.

sdns://AQcAAAAAAAAAF1syMDAxOjQ3MDoxZjE1OmI4MDo6NTNdIKnWMjpPJYAJJhl1FQLOIx4fdtned2yHxruyig7_2w5OIDIuZG5zY3J5cHQtY2VydC5vcGVubmljLmkycGQueHl6


## opennic-bongobow

OpenNIC • Non-logging • No DNSSEC
Location: Munich, Germany

Warning: this server is incompatible with DNS anonymization.

sdns://AQYAAAAAAAAAETUuMTg5LjE3MC4xOTY6NDY1IFQ1LFVAO4Luk8QH_cI0RJcNmlbvIr_P-eyQnM0yJDJrKDIuZG5zY3J5cHQtY2VydC5uczE2LmRlLmRucy5vcGVubmljLmdsdWU


## opennic-fische

OpenNIC • Non-logging • DNSSEC
Location: Nurnberg, Germany

sdns://AQcAAAAAAAAAEDc4LjQ3LjI0My4zOjEwNTMgN4CAbUDR-b3uJJMVzfCdL9ivVV7s8wRhifLRPWBfSmQdMi5kbnNjcnlwdC1jZXJ0Lm5zMS5maXNjaGUuaW8


## opennic-iriseden

OpenNIC • Non-logging • DNSSEC
Location: Paris, France
Maintained by iriseden.

sdns://AQcAAAAAAAAAEzYyLjIxMC4xNzcuMTg5OjEwNTMgW8vytBGk6u3kvCpl4q88XjqW-w6JJiJ7QBObcFV7gYAfMi5kbnNjcnlwdC1jZXJ0Lm5zMS5pcmlzZWRlbi5mcg


## opennic-luggs

Public DNS server in Canada operated by Luggs

sdns://AQYAAAAAAAAADTE0Mi40LjIwNC4xMTEgHBl5MxvoI8zPCJp5BpN-XDQQKlasf2Jw4EYlsu3bBOMfMi5kbnNjcnlwdC1jZXJ0Lm5zMy5jYS5sdWdncy5jbw


## opennic-luggs2

Second public DNS server in Canada operated by Luggs

sdns://AQYAAAAAAAAADDE0Mi40LjIwNS40NyC8v7fgUME9okIsALCxrJrWSMXCZLy2FwuIPXGKyG66CR8yLmRuc2NyeXB0LWNlcnQubnM0LmNhLmx1Z2dzLmNv


## opennic-rico4514

OpenNIC • Non-logging • No DNSSEC
Location: Texas, 13, MX

sdns://AQYAAAAAAAAADTE0Mi40LjIwNC4xMTEgHBl5MxvoI8zPCJp5BpN-XDQQKlasf2Jw4EYlsu3bBOMfMi5kbnNjcnlwdC1jZXJ0Lm5zMy5jYS5sdWdncy5jbw


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

sdns://AQcAAAAAAAAADTEzOS45OS4yMjIuNzIgCwVoTw0L4dgal5LC1FbZUtHtLR_rVuV6rVnxO95e4GUbMi5kbnNjcnlwdC1jZXJ0LmRucy5zZWJ5Lmlv


## publicarray-au2-doh

DNSSEC • OpenNIC • Non-logging • Uncensored - hosted on ovh.com.au
Maintained by publicarray - https://dns.seby.io

sdns://AgcAAAAAAAAADTEzOS45OS4yMjIuNzKgPhoaD2xT8-l6SS1XCEtbmAcFnuBXqxUFh2_YP9o9uDggMob_ZaZfrzIIXuoTiMNzi6fjeHPJBszjxKKLTMKliYgRZG9oLTIuc2VieS5pbzo0NDMKL2Rucy1xdWVyeQ

