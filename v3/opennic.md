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

sdns://AgcAAAAAAAAAACA-GhoPbFPz6XpJLVcIS1uYBwWe4FerFQWHb9g_2j24OBRpYmtzdHVybS5zeW5vbG9neS5tZQovZG5zLXF1ZXJ5


## ibksturm

dnscrypt-server (nginx - encrypted-dns - unbound backend), DNSSEC / Non-Logged / Uncensored, OpenNIC and Root DNS-Zone Copy
Hosted in Switzerland by ibksturm, aka Andreas Ziegler

sdns://AQcAAAAAAAAAEDg1LjUuOTMuMjMwOjg0NDMgwc9XUACwW8JsYh9ez5qiVgrOvwB-vss6f_SyDeC0Oe4YMi5kbnNjcnlwdC1jZXJ0Lmlia3N0dXJt


## ibksturm-ipv6

dnscrypt-server (nginx - encrypted-dns - unbound backend), DNSSEC / Non-Logged / Uncensored, OpenNIC and Root DNS-Zone Copy
Hosted in Switzerland by ibksturm, aka Andreas Ziegler

sdns://AQcAAAAAAAAALlsyYTAyOjEyMDU6NTA1NTpkZTYwOmIyNmU6YmZmZjpmZTFkOmUxOWJdOjg0NDMgwc9XUACwW8JsYh9ez5qiVgrOvwB-vss6f_SyDeC0Oe4YMi5kbnNjcnlwdC1jZXJ0Lmlia3N0dXJt


## opennic-R4SAS

DNSSEC - OpenNIC - Non-logging - Uncensored - hosted on ovh.com
Location: Gravelines, France.
Maintained by R4SAS.

sdns://AQcAAAAAAAAAETE1MS44MC4yMjIuNzk6NDQzIKnWMjpPJYAJJhl1FQLOIx4fdtned2yHxruyig7_2w5OIDIuZG5zY3J5cHQtY2VydC5vcGVubmljLmkycGQueHl6


## opennic-R4SAS-ipv6

DNSSEC - OpenNIC - Non-logging - Uncensored - hosted on ovh.com
Location: Gravelines, France.
Maintained by R4SAS.

sdns://AQcAAAAAAAAAG1syMDAxOjQ3MDoxZjE1OmI4MDo6NTNdOjQ0MyCp1jI6TyWACSYZdRUCziMeH3bZ3ndsh8a7sooO_9sOTiAyLmRuc2NyeXB0LWNlcnQub3Blbm5pYy5pMnBkLnh5eg


## opennic-bongobow

OpenNIC • Non-logging • No DNSSEC
Location: Munich, Germany

sdns://AQYAAAAAAAAAETUuMTg5LjE3MC4xOTY6NDY1IFQ1LFVAO4Luk8QH_cI0RJcNmlbvIr_P-eyQnM0yJDJrKDIuZG5zY3J5cHQtY2VydC5uczE2LmRlLmRucy5vcGVubmljLmdsdWU


## opennic-fische

OpenNIC • Non-logging • DNSSEC
Location: Nurnberg, Germany

sdns://AQcAAAAAAAAAEjYyLjIxMC4xODAuNzE6MTA1MyBxLWt8kNHoMqM7vKXCkuZ3PnB32c0qV2I3KGQYtlDKSB8yLmRuc2NyeXB0LWNlcnQubnMyLmlyaXNlZGVuLmZy


## opennic-fische-ipv6

OpenNIC • Non-logging • DNSSEC
Location: Nurnberg, Germany

sdns://AQcAAAAAAAAAHFsyYTAxOjRmODoxYzBjOjgwYzk6OjFdOjEwNTMgcmZXgMxIKLKAtkLUX7t6Lhw7j4-PIqXir5hMytnM-W8dMi5kbnNjcnlwdC1jZXJ0Lm5zMS5maXNjaGUuaW8


## opennic-iriseden

OpenNIC • Non-logging • DNSSEC
Location: Paris, France
Maintained by iriseden.

sdns://AQcAAAAAAAAAHVsyMDAxOmJjODozMmQ3OjMwODo6MjAxXToxMDUzIEUAcwKTPY6tyEQxtfO3rIzEyqN9w7WGPLz7ZsHsx5EGHzIuZG5zY3J5cHQtY2VydC5uczEuaXJpc2VkZW4uZnI


## opennic-iriseden-ipv6

OpenNIC • Non-logging • DNSSEC
Location: Paris, France
Maintained by iriseden.

sdns://AQcAAAAAAAAAHVsyMDAxOmJjODozMmQ3OjMwNzo6MzAxXToxMDUzIJjeEela3WTzMuuZTskr7aOchIg2llSDNRsHfcggITn6HzIuZG5zY3J5cHQtY2VydC5uczIuaXJpc2VkZW4uZnI


## opennic-luggs

Public DNS server in Canada operated by Luggs

sdns://AQYAAAAAAAAADTE0Mi40LjIwNC4xMTEgHBl5MxvoI8zPCJp5BpN-XDQQKlasf2Jw4EYlsu3bBOMfMi5kbnNjcnlwdC1jZXJ0Lm5zMy5jYS5sdWdncy5jbw


## opennic-luggs2

Second public DNS server in Canada operated by Luggs

sdns://AQYAAAAAAAAAEDE0Mi40LjIwNS40Nzo0NDMgvL-34FDBPaJCLACwsaya1kjFwmS8thcLiD1xishuugkfMi5kbnNjcnlwdC1jZXJ0Lm5zNC5jYS5sdWdncy5jbw


## opennic-rico4514

OpenNIC • Non-logging • No DNSSEC
Location: Texas, 13, MX

sdns://AQYAAAAAAAAAETE0Mi40LjIwNC4xMTE6NDQzIBwZeTMb6CPMzwiaeQaTflw0ECpWrH9icOBGJbLt2wTjHzIuZG5zY3J5cHQtY2VydC5uczMuY2EubHVnZ3MuY28


## publicarray-au

DNSSEC • OpenNIC • Non-logging • Uncensored - hosted on vultr.com
Maintained by publicarray - https://dns.seby.io

sdns://AQcAAAAAAAAADDQ1Ljc2LjExMy4zMSAIVGh4i6eKXqlF6o9Fg92cgD2WcDvKQJ7v_Wq4XrQsVhsyLmRuc2NyeXB0LWNlcnQuZG5zLnNlYnkuaW8


## publicarray-au-doh

DNSSEC • OpenNIC • Non-logging • Uncensored - hosted on vultr.com
Maintained by publicarray - https://dns.seby.io

sdns://AgcAAAAAAAAADDQ1Ljc2LjExMy4zMSA-GhoPbFPz6XpJLVcIS1uYBwWe4FerFQWHb9g_2j24OBBkb2guc2VieS5pbzo4NDQzCi9kbnMtcXVlcnk


## publicarray-au2

DNSSEC • OpenNIC • Non-logging • Uncensored - hosted on ovh.com.au
Maintained by publicarray - https://dns.seby.io

sdns://AQcAAAAAAAAADTEzOS45OS4yMjIuNzIgCwVoTw0L4dgal5LC1FbZUtHtLR_rVuV6rVnxO95e4GUbMi5kbnNjcnlwdC1jZXJ0LmRucy5zZWJ5Lmlv


## publicarray-au2-doh

DNSSEC • OpenNIC • Non-logging • Uncensored - hosted on ovh.com.au
Maintained by publicarray - https://dns.seby.io

sdns://AgcAAAAAAAAADTEzOS45OS4yMjIuNzIgPhoaD2xT8-l6SS1XCEtbmAcFnuBXqxUFh2_YP9o9uDgRZG9oLTIuc2VieS5pbzo0NDMKL2Rucy1xdWVyeQ

