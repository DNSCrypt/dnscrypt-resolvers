# Anonymized DNS relays

Anonymized DNS is a lightweight alternative to Tor and SOCKS proxies,
dedicated to DNS traffic. They hide the client IP address to DNS resolvers,
providing anonymity in addition to confidentiality and integrity.

DNS Anonymization is only compatible with servers supporting the
DNSCrypt protocol.

See the link below for more information:

https://github.com/DNSCrypt/dnscrypt-proxy/wiki/Anonymized-DNS

To use that list, add this to the `[sources]` section of your
`dnscrypt-proxy.toml` configuration file:

    [sources.'relays']
    urls = ['https://raw.githubusercontent.com/DNSCrypt/dnscrypt-resolvers/master/v3/relays.md', 'https://download.dnscrypt.info/resolvers-list/v3/relays.md', 'https://cdn.jsdelivr.net/gh/DNSCrypt/dnscrypt-resolvers@master/v3/relays.md']
    minisign_key = 'RWQf6LRCGA9i53mlYecO4IzT51TGPpvWucNSCh1CBM0QTaLn73Y7GFO3'
    cache_file = 'relays.md'

--


## anon-cipherdns-ct1-za

CipherDNS Cape Town relay.
Anonymized DNSCrypt relay based in Cape Town, South Africa.
Provider information is limited to the resolver-list entry; report issues through the dnscrypt-resolvers project.

sdns://gRMxMDIuMjA5LjIxLjE3Njo4NDQz


## anon-cipherdns-jb1-za

CipherDNS Johannesburg relay.
Anonymized DNSCrypt relay based in Johannesburg, South Africa.
Provider information is limited to the resolver-list entry; report issues through the dnscrypt-resolvers project.

sdns://gRIxMDIuMjE0LjEwLjgyOjg0NDM


## anon-cs-austria

CryptoStorm Vienna, Austria relay.
Anonymized DNS relay operated by CryptoStorm. https://cryptostorm.is/

sdns://gRE5NC4xOTguNDEuMjM1OjQ0Mw


## anon-cs-austria6

CryptoStorm Vienna, Austria relay.
IPv6 endpoint. Anonymized DNS relay operated by CryptoStorm. https://cryptostorm.is/

sdns://gRhbMjAwMTphYzg6Mjk6YTE6OjUzXTo0NDM


## anon-cs-barcelona

CryptoStorm Barcelona, Spain relay.
Anonymized DNS relay operated by CryptoStorm. https://cryptostorm.is/

sdns://gRIzNy4xMjAuMTQyLjExNTo0NDM


## anon-cs-belgium

CryptoStorm Brussels, Belgium relay.
Anonymized DNS relay operated by CryptoStorm. https://cryptostorm.is/

sdns://gREzNy4xMjAuMjM2LjExOjQ0Mw


## anon-cs-belgium6

CryptoStorm Brussels, Belgium relay.
IPv6 endpoint. Anonymized DNS relay operated by CryptoStorm. https://cryptostorm.is/

sdns://gRlbMjAwMTphYzg6Mjc6MTAzOjo1M106NDQz


## anon-cs-berlin

CryptoStorm Berlin, Germany relay.
Anonymized DNS relay operated by CryptoStorm. https://cryptostorm.is/

sdns://gREzNy4xMjAuMjE3Ljc1OjQ0Mw


## anon-cs-berlin6

CryptoStorm Berlin, Germany relay.
IPv6 endpoint. Anonymized DNS relay operated by CryptoStorm. https://cryptostorm.is/

sdns://gRhbMjAwMTphYzg6MzY6NjE6OjUzXTo0NDM


## anon-cs-ch

CryptoStorm Switzerland relay.
Anonymized DNS relay operated by CryptoStorm. https://cryptostorm.is/

sdns://gRMxOTAuMjExLjI1NS4yMjc6NDQz


## anon-cs-czech

CryptoStorm Prague, Czech Republic relay.
Anonymized DNS relay operated by CryptoStorm. https://cryptostorm.is/

sdns://gRMyMTcuMTM4LjIyMC4yNDM6NDQz


## anon-cs-czech6

CryptoStorm Prague, Czech Republic relay.
IPv6 endpoint. Anonymized DNS relay operated by CryptoStorm. https://cryptostorm.is/

sdns://gRhbMjAwMTphYzg6MzM6Nzc6OjUzXTo0NDM


## anon-cs-dc

CryptoStorm Washington, DC relay.
Anonymized DNS relay operated by CryptoStorm. https://cryptostorm.is/

sdns://gRAxOTguNy41OC4yMjc6NDQz


## anon-cs-dc6

CryptoStorm Washington, DC relay.
IPv6 endpoint. Anonymized DNS relay operated by CryptoStorm. https://cryptostorm.is/

sdns://gR9bMjYwNDo5YTAwOjIwMTA6YTBiYjo2Ojo1M106NDQz


## anon-cs-de

CryptoStorm Frankfurt, Germany relay.
Anonymized DNS relay operated by CryptoStorm. https://cryptostorm.is/

sdns://gQ8xNDYuNzAuODIuMzo0NDM


## anon-cs-de6

CryptoStorm Frankfurt, Germany relay.
IPv6 endpoint. Anonymized DNS relay operated by CryptoStorm. https://cryptostorm.is/

sdns://gRhbMmEwZDo1NjAwOjFkOjk6OjUzXTo0NDM


## anon-cs-dus

CryptoStorm Düsseldorf, Germany relay.
Anonymized DNS relay operated by CryptoStorm. https://cryptostorm.is/

sdns://gRI4OS4xNjMuMjIxLjE4MTo0NDM


## anon-cs-dus6

CryptoStorm Düsseldorf, Germany relay.
IPv6 endpoint. Anonymized DNS relay operated by CryptoStorm. https://cryptostorm.is/

sdns://gRtbMjAwMTo0YmEwOmZmZWQ6NzY6OjUzXTo0NDM


## anon-cs-finland

CryptoStorm Finland relay.
Anonymized DNS relay operated by CryptoStorm. https://cryptostorm.is/

sdns://gRE4My4xNDMuMjQyLjQzOjQ0Mw


## anon-cs-finland6

CryptoStorm Finland relay.
IPv6 endpoint. Anonymized DNS relay operated by CryptoStorm. https://cryptostorm.is/

sdns://gRpbMmEwZDo1NjAwOjE0MjoxMTo6NTNdOjQ0Mw


## anon-cs-fl

CryptoStorm Miami, FL relay.
Anonymized DNS relay operated by CryptoStorm. https://cryptostorm.is/

sdns://gRIxNDYuNzAuMjQwLjIwMzo0NDM


## anon-cs-fl6

CryptoStorm Miami, FL relay.
IPv6 endpoint. Anonymized DNS relay operated by CryptoStorm. https://cryptostorm.is/

sdns://gRlbMmEwZDo1NjAwOjY6MTIzOjo1M106NDQz


## anon-cs-fr

CryptoStorm France relay.
Anonymized DNS relay operated by CryptoStorm. https://cryptostorm.is/

sdns://gRExNjMuMTcyLjM0LjU2OjQ0Mw


## anon-cs-ga

CryptoStorm Atlanta, GA relay.
Anonymized DNS relay operated by CryptoStorm. https://cryptostorm.is/

sdns://gRMxMzAuMTk1LjIxMi4yMTE6NDQz


## anon-cs-ga6

CryptoStorm Atlanta, GA relay.
IPv6 endpoint. Anonymized DNS relay operated by CryptoStorm. https://cryptostorm.is/

sdns://gRlbMmEwZDo1NjAwOjE0NTo1Ojo1M106NDQz


## anon-cs-hungary

CryptoStorm Budapest, Hungary relay.
Anonymized DNS relay operated by CryptoStorm. https://cryptostorm.is/

sdns://gRE4Ni4xMDYuNzQuMjE5OjQ0Mw


## anon-cs-hungary6

CryptoStorm Budapest, Hungary relay.
IPv6 endpoint. Anonymized DNS relay operated by CryptoStorm. https://cryptostorm.is/

sdns://gRhbMjAwMTphYzg6MjY6NjE6OjUzXTo0NDM


## anon-cs-il

CryptoStorm Chicago, IL relay.
Anonymized DNS relay operated by CryptoStorm. https://cryptostorm.is/

sdns://gRMxOTUuMjQyLjIxMi4xMzE6NDQz


## anon-cs-il6

CryptoStorm Chicago, IL relay.
IPv6 endpoint. Anonymized DNS relay operated by CryptoStorm. https://cryptostorm.is/

sdns://gRlbMmEwZDo1NjAwOjE0NDoxOjo1M106NDQz


## anon-cs-la

CryptoStorm Los Angeles, CA relay.
Anonymized DNS relay operated by CryptoStorm. https://cryptostorm.is/

sdns://gRMxOTUuMjA2LjEwNC4yMDM6NDQz


## anon-cs-la6

CryptoStorm Los Angeles, CA relay.
IPv6 endpoint. Anonymized DNS relay operated by CryptoStorm. https://cryptostorm.is/

sdns://gRhbMmEwZDo1NjAwOjRmOjU6OjUzXTo0NDM


## anon-cs-london

CryptoStorm London, England relay.
Anonymized DNS relay operated by CryptoStorm. https://cryptostorm.is/

sdns://gRE3OC4xMjkuMjQ4LjY3OjQ0Mw


## anon-cs-london6

CryptoStorm London, England relay.
IPv6 endpoint. Anonymized DNS relay operated by CryptoStorm. https://cryptostorm.is/

sdns://gRtbMjAwMToxYjQwOjUwMDA6YTI6OjUzXTo0NDM


## anon-cs-manchester

CryptoStorm Manchester, England relay.
Anonymized DNS relay operated by CryptoStorm. https://cryptostorm.is/

sdns://gRExOTUuMTIuNDguMTcxOjQ0Mw


## anon-cs-manchester6

CryptoStorm Manchester, England relay.
IPv6 endpoint. Anonymized DNS relay operated by CryptoStorm. https://cryptostorm.is/

sdns://gRhbMjAwMTphYzg6OGI6NjE6OjUzXTo0NDM


## anon-cs-md

CryptoStorm Moldova relay.
Anonymized DNS relay operated by CryptoStorm. https://cryptostorm.is/

sdns://gRExNzYuMTIzLjQuMjMxOjQ0Mw


## anon-cs-milan

CryptoStorm Milan, Italy relay.
Anonymized DNS relay operated by CryptoStorm. https://cryptostorm.is/

sdns://gRMyMTcuMTM4LjIxOS4yMTk6NDQz


## anon-cs-milan6

CryptoStorm Milan, Italy relay.
IPv6 endpoint. Anonymized DNS relay operated by CryptoStorm. https://cryptostorm.is/

sdns://gRhbMjAwMTphYzg6MjQ6YTE6OjUzXTo0NDM


## anon-cs-montreal

CryptoStorm Montreal, Canada relay.
Anonymized DNS relay operated by CryptoStorm. https://cryptostorm.is/

sdns://gRExNzYuMTEzLjc0LjE5OjQ0Mw


## anon-cs-montreal6

CryptoStorm Montreal, Canada relay.
IPv6 endpoint. Anonymized DNS relay operated by CryptoStorm. https://cryptostorm.is/

sdns://gRhbMmEwZDo1NjAwOjE5OjU6OjUzXTo0NDM


## anon-cs-nl

CryptoStorm Netherlands relay.
Anonymized DNS relay operated by CryptoStorm. https://cryptostorm.is/

sdns://gRExODUuMTA3LjgwLjg0OjQ0Mw


## anon-cs-nl6

CryptoStorm Netherlands relay.
IPv6 endpoint. Anonymized DNS relay operated by CryptoStorm. https://cryptostorm.is/

sdns://gRpbMmEwMDoxNzY4OjYwMDE6ODo6NTNdOjQ0Mw


## anon-cs-norway

CryptoStorm Oslo, Norway relay.
Anonymized DNS relay operated by CryptoStorm. https://cryptostorm.is/

sdns://gRI5MS4yMTkuMjE1LjIyNzo0NDM


## anon-cs-norway6

CryptoStorm Oslo, Norway relay.
IPv6 endpoint. Anonymized DNS relay operated by CryptoStorm. https://cryptostorm.is/

sdns://gRhbMjAwMTphYzg6Mzg6OTQ6OjUzXTo0NDM


## anon-cs-nv

CryptoStorm Las Vegas, NV relay.
Anonymized DNS relay operated by CryptoStorm. https://cryptostorm.is/

sdns://gRA3OS4xMTAuNTMuNTE6NDQz


## anon-cs-nv6

CryptoStorm Las Vegas, NV relay.
IPv6 endpoint. Anonymized DNS relay operated by CryptoStorm. https://cryptostorm.is/

sdns://gRhbMmEwZDo1NjAwOjM6MTk6OjUzXTo0NDM


## anon-cs-nyc

CryptoStorm New York City, NY relay.
Anonymized DNS relay operated by CryptoStorm. https://cryptostorm.is/

sdns://gRExNDYuNzAuMTU0LjY3OjQ0Mw


## anon-cs-nyc6

CryptoStorm New York City, NY relay.
IPv6 endpoint. Anonymized DNS relay operated by CryptoStorm. https://cryptostorm.is/

sdns://gRlbMmEwZDo1NjAwOjI0OjU0Ojo1M106NDQz


## anon-cs-ore

CryptoStorm Oregon relay.
Anonymized DNS relay operated by CryptoStorm. https://cryptostorm.is/

sdns://gRExNzkuNjEuMjIzLjQ3OjQ0Mw


## anon-cs-ore6

CryptoStorm Oregon relay.
IPv6 endpoint. Anonymized DNS relay operated by CryptoStorm. https://cryptostorm.is/

sdns://gRdbMjYwNTo2YzgwOjU6ZDo6NTNdOjQ0Mw


## anon-cs-poland

CryptoStorm Warsaw, Poland relay.
Anonymized DNS relay operated by CryptoStorm. https://cryptostorm.is/

sdns://gREzNy4xMjAuMjExLjkxOjQ0Mw


## anon-cs-poland6

CryptoStorm Warsaw, Poland relay.
IPv6 endpoint. Anonymized DNS relay operated by CryptoStorm. https://cryptostorm.is/

sdns://gRlbMmEwZDo1NjAwOjEzOjcxOjo1M106NDQz


## anon-cs-pt

CryptoStorm Portugal relay.
Anonymized DNS relay operated by CryptoStorm. https://cryptostorm.is/

sdns://gRI5MS4yMDUuMjMwLjIyNDo0NDM


## anon-cs-ro

CryptoStorm Romania relay.
Anonymized DNS relay operated by CryptoStorm. https://cryptostorm.is/

sdns://gRExNDYuNzAuNjYuMjI3OjQ0Mw


## anon-cs-sea

CryptoStorm Seattle, WA relay.
Anonymized DNS relay operated by CryptoStorm. https://cryptostorm.is/

sdns://gRA2NC4xMjAuNS4yNTE6NDQz


## anon-cs-sea6

CryptoStorm Seattle, WA relay.
IPv6 endpoint. Anonymized DNS relay operated by CryptoStorm. https://cryptostorm.is/

sdns://gRxbMjYwNzpmNWIyOjE6YTAwYjpiOjo1M106NDQz


## anon-cs-serbia

CryptoStorm Belgrade, Serbia relay.
Anonymized DNS relay operated by CryptoStorm. https://cryptostorm.is/

sdns://gRIzNy4xMjAuMTkzLjIxOTo0NDM


## anon-cs-serbia6

CryptoStorm Belgrade, Serbia relay.
IPv6 endpoint. Anonymized DNS relay operated by CryptoStorm. https://cryptostorm.is/

sdns://gRhbMjAwMTphYzg6N2Q6NDc6OjUzXTo0NDM


## anon-cs-singapore

CryptoStorm Singapore relay.
Anonymized DNS relay operated by CryptoStorm. https://cryptostorm.is/

sdns://gREzNy4xMjAuMTUxLjExOjQ0Mw


## anon-cs-singapore6

CryptoStorm Singapore relay.
IPv6 endpoint. Anonymized DNS relay operated by CryptoStorm. https://cryptostorm.is/

sdns://gRhbMmEwZDo1NjAwOjFmOjc6OjUzXTo0NDM


## anon-cs-swe

CryptoStorm Sweden relay.
Anonymized DNS relay operated by CryptoStorm. https://cryptostorm.is/

sdns://gRMxMjguMTI3LjEwNC4xMDg6NDQz


## anon-cs-sydney

CryptoStorm Sydney, Australia relay.
Anonymized DNS relay operated by CryptoStorm. https://cryptostorm.is/

sdns://gRIzNy4xMjAuMjM0LjI1MTo0NDM


## anon-cs-tokyo

CryptoStorm Tokyo, Japan relay.
Anonymized DNS relay operated by CryptoStorm. https://cryptostorm.is/

sdns://gRAxNDYuNzAuMzEuNDM6NDQz


## anon-cs-tokyo6

CryptoStorm Tokyo, Japan relay.
IPv6 endpoint. Anonymized DNS relay operated by CryptoStorm. https://cryptostorm.is/

sdns://gRhbMjAwMTphYzg6NDA6ZGY6OjUzXTo0NDM


## anon-cs-tx

CryptoStorm Dallas, TX relay.
Anonymized DNS relay operated by CryptoStorm. https://cryptostorm.is/

sdns://gREyMDkuNTguMTQ3LjM2OjQ0Mw


## anon-cs-tx6

CryptoStorm Dallas, TX relay.
IPv6 endpoint. Anonymized DNS relay operated by CryptoStorm. https://cryptostorm.is/

sdns://gR9bMjYwNjo5ODgwOjIxMDA6YTAwNjozOjo1M106NDQz


## anon-cs-vancouver

CryptoStorm Vancouver, Canada relay.
Anonymized DNS relay operated by CryptoStorm. https://cryptostorm.is/

sdns://gRIxOTYuMjQwLjc5LjE2Mzo0NDM


## anon-cs-vancouver6

CryptoStorm Vancouver, Canada relay.
IPv6 endpoint. Anonymized DNS relay operated by CryptoStorm. https://cryptostorm.is/

sdns://gRlbMmEwMjo1NzQwOjI0OjQ1Ojo1M106NDQz


## anon-dnswarden-swiss

DNSWarden Switzerland relay.
Hosted in Switzerland. Maintained by DNSWarden. More information: https://github.com/bhanupratapys/dnswarden and https://dnswarden.com

sdns://gRQxODguMjQ0LjExNy4xMTQ6MTQ0Mw


## anon-kama

kama relay in France.
Anonymized DNS relay hosted in France and maintained by Frank Denis (@jedisct1).

sdns://gQ4xMzcuNzQuMjIzLjIzNA


## anon-saldns01-conoha-ipv4

salDNS Tokyo relay.
Hosted on ConoHa VPS Tokyo region. No logs. From the experimental &mu;ODNS project: https://junkurihara.github.io/dns/.

sdns://gRQxNjMuNDQuMTI0LjIwNDo1MDQ0Mw


## anon-saldns02-conoha-ipv4

salDNS Tokyo relay.
Hosted on ConoHa VPS Tokyo region. No logs. From the experimental &mu;ODNS project: https://junkurihara.github.io/dns/.

sdns://gRUxNjAuMjUxLjIxNC4xNzI6NTA0NDM


## anon-saldns03-conoha-ipv4

salDNS Tokyo relay.
Hosted on ConoHa VPS Tokyo region. No logs. From the experimental &mu;ODNS project: https://junkurihara.github.io/dns/.

sdns://gRQxNjAuMjUxLjE2OC4yNTo1MDQ0Mw


## anon-scaleway

Anonymized DNS relay hosted in France and maintained by Frank Denis (@jedisct1).
Service page: https://www.scaleway.com/

sdns://gRIyMTIuNDcuMjI4LjEzNjo0NDM


## anon-scaleway-ams

Anonymized DNS relay hosted in Amsterdam and maintained by Frank Denis (@jedisct1).
Service page: https://www.scaleway.com/

sdns://gRE1MS4xNS4xMjIuMjUwOjQ0Mw


## anon-scaleway-ams-ipv6

Anonymized DNS relay hosted in Amsterdam and maintained by Frank Denis (@jedisct1).
IPv6 only.
Service page: https://www.scaleway.com/

sdns://gSpbMjAwMTpiYzg6MTY0MDoxY2UyOmRjMDA6ZmY6ZmUyODo1YjE3XTo0NDM


## anon-scaleway2

Anonymized DNS relay hosted in France and maintained by Frank Denis (@jedisct1).
Service page: https://www.scaleway.com/

sdns://gRMxNjMuMTcyLjE4MC4xMjU6NDQz


## anon-serbica

Anonymized DNS relay hosted in Netherlands by https://litepay.ch

sdns://gRMxODUuNjYuMTQzLjE3ODo1MzUz


## anon-tiarap

Anonymized DNS relay hosted in Singapore
Maintained by id-gmail / Tiarap. Service page: https://tiarap.org/

sdns://gRMxNzQuMTM4LjI5LjE3NToxNDQz


## dnscry.pt-anon-adelaide-ipv4

dnscry.pt Adelaide relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ4xNzUuNDUuMTgyLjE3OQ


## dnscry.pt-anon-adelaide-ipv6

dnscry.pt Adelaide relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gRhbMjQwNDo5NDAwOjY4ZjE6MTAwOjo1M10


## dnscry.pt-anon-allendale-ipv4

dnscry.pt Allendale relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ4yMy4xMjkuMTgwLjE5NQ


## dnscry.pt-anon-allendale-ipv6

dnscry.pt Allendale relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gRVbMjYwMjpmN2EzOjA6NTIwMDo6YV0


## dnscry.pt-anon-allentown-ipv4

dnscry.pt Allentown relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ0yMy4xMzcuMjUzLjI0


## dnscry.pt-anon-allentown-ipv6

dnscry.pt Allentown relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gRxbMjYwMjpmYzI0OjE5Ojc0YjA6NTI4NTo6MTJd


## dnscry.pt-anon-amsterdam-ipv4

dnscry.pt Amsterdam relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ4xOTguMTQwLjE0MS40Ng


## dnscry.pt-anon-amsterdam-ipv6

dnscry.pt Amsterdam relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gRZbMmEwMzo5NGUzOjIyMmI6OjEwMzJd


## dnscry.pt-anon-amsterdam02-ipv4

dnscry.pt Amsterdam 02 relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ00NS44Ni4xNjIuMTEw


## dnscry.pt-anon-amsterdam02-ipv6

dnscry.pt Amsterdam 02 relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gRtbMmEwNzplZmMwOjEwMDE6YTVjZTo6YjRiNF0


## dnscry.pt-anon-amsterdam03-ipv4

dnscry.pt Amsterdam 03 relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ0yMy4xMzcuMjQ5LjI2


## dnscry.pt-anon-amsterdam03-ipv6

dnscry.pt Amsterdam 03 relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gRhbMjYwMjpmYzI0OjEyOjk4NzM6OmFiMV0


## dnscry.pt-anon-ashburn-ipv4

dnscry.pt Ashburn relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ43Ny4yNDcuMTI3LjEwNw


## dnscry.pt-anon-ashburn-ipv6

dnscry.pt Ashburn relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gRNbMmEwYTo4ZGMwOmEwNjc6OmFd


## dnscry.pt-anon-atlanta-ipv4

dnscry.pt Atlanta relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ8xNzAuMjQ5LjIzNy4xNTQ


## dnscry.pt-anon-auckland-ipv4

dnscry.pt Auckland relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ4xODUuOTkuMTMzLjExMg


## dnscry.pt-anon-auckland-ipv6

dnscry.pt Auckland relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gRxbMmEwNjoxMjgwOmJlZTE6Mjo6ZWUxMjoyMDhd


## dnscry.pt-anon-baku-ipv4

dnscry.pt Baku relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ0xODAuMTQ5LjQ0LjIy


## dnscry.pt-anon-baku-ipv6

dnscry.pt Baku relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gRNbMmEwMzo5MGMwOjE5NTo6OTFd


## dnscry.pt-anon-bangkok-ipv4

dnscry.pt Bangkok relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ0xMDMuMzguMjUwLjU1


## dnscry.pt-anon-bengaluru-ipv4

dnscry.pt Bengaluru relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ8xNjAuMTkxLjE4Mi4yMTY


## dnscry.pt-anon-bengaluru-ipv6

dnscry.pt Bengaluru relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gRZbMjQwMTpkNGUwOjE6ZjdmZDo6NTNd


## dnscry.pt-anon-bogota-ipv4

dnscry.pt Bogotá relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ0xMDMuNTcuMjUwLjU0


## dnscry.pt-anon-bogota-ipv6

dnscry.pt Bogotá relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gRVbMmEwMzpmODA6NTc6OThiMTo6MV0


## dnscry.pt-anon-bratislava-ipv4

dnscry.pt Bratislava relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ45NS4xMzEuMjAyLjEwNQ


## dnscry.pt-anon-bratislava-ipv6

dnscry.pt Bratislava relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gR1bMmEwNTo1NTAyOjo1OTA2Ojk3Zjg6MmQwZToxXQ


## dnscry.pt-anon-bremen-ipv4

dnscry.pt Bremen relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ04OS4xMDYuNzguMTA2


## dnscry.pt-anon-brisbane-ipv4

dnscry.pt Brisbane relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ40My4yMjQuMTgwLjEzNw


## dnscry.pt-anon-brisbane-ipv6

dnscry.pt Brisbane relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gSJbMjQwNDo5NDAwOjE6MDoyMTY6M2VmZjpmZWY2OjcxOTRd


## dnscry.pt-anon-brussels-ipv4

dnscry.pt Brussels relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ8xOTIuMTIxLjE3MC4xNTE


## dnscry.pt-anon-brussels-ipv6

dnscry.pt Brussels relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gRVbMmEwMzpmODA6MzI6NTJkOTo6MV0


## dnscry.pt-anon-bucharest-ipv4

dnscry.pt Bucharest relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ4xODUuOTMuMjIxLjE2Nw


## dnscry.pt-anon-budapest-ipv4

dnscry.pt Budapest relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ8xOTMuMjAxLjE4NS4xNDY


## dnscry.pt-anon-budapest-ipv6

dnscry.pt Budapest relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gRdbMmEwMTo2ZWUwOjE6OmZmZmY6YmFlXQ


## dnscry.pt-anon-calgary-ipv4

dnscry.pt Calgary relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ0yMy4xMzMuNjQuMTIx


## dnscry.pt-anon-capetown-ipv4

dnscry.pt Cape Town relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ4xMDIuMjE2Ljc5LjIzNw


## dnscry.pt-anon-capetown-ipv6

dnscry.pt Cape Town relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gRlbMmMwZjplZjE4OjlmZmY6MTpiZmY6OmFd


## dnscry.pt-anon-capetown02-ipv4

dnscry.pt Cape Town 02 relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ8xNjAuMTE5LjIzMy4yNDU


## dnscry.pt-anon-chicago-ipv4

dnscry.pt Chicago relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ00NS40MS4yMDQuMjA0


## dnscry.pt-anon-chicago-ipv6

dnscry.pt Chicago relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gR9bMjYwMjpmZWE3OmUwYzplOmJmZjo2OjcwOjE5NGNd


## dnscry.pt-anon-chisinau-ipv4

dnscry.pt Chișinău relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ4xNzYuMTIzLjEwLjEwNQ


## dnscry.pt-anon-copenhagen-ipv4

dnscry.pt Copenhagen relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ4xOTIuMTIxLjExOS4xOQ


## dnscry.pt-anon-copenhagen-ipv6

dnscry.pt Copenhagen relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gSNbMjAwMTo2N2M6YmVjOmI6NDNhOjFhZmY6ZmViMTplYjVkXQ


## dnscry.pt-anon-coventry-ipv4

dnscry.pt Coventry relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ00NS4xNTUuMzcuMTY1


## dnscry.pt-anon-coventry-ipv6

dnscry.pt Coventry relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gRhbMmEwZDpkOGMwOjA6ZjA0Mzo6NjkyN10


## dnscry.pt-anon-dallas-ipv4

dnscry.pt Dallas relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ0yMy4yMzAuMjUzLjk4


## dnscry.pt-anon-dallas-ipv6

dnscry.pt Dallas relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gRNbMjYwMjpmYjk0OjE6Mzk6OmFd


## dnscry.pt-anon-denver-ipv4

dnscry.pt Denver relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ8yMTYuMTIwLjIwMS4xMDU


## dnscry.pt-anon-denver-ipv6

dnscry.pt Denver relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gRhbMjYwNzphNjgwOjY6ZjAxNjo6M2EyNV0


## dnscry.pt-anon-detroit-ipv4

dnscry.pt Detroit relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQw2Ni4xODcuNy4xNDA


## dnscry.pt-anon-detroit-ipv6

dnscry.pt Detroit relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gSNbMjYwNjo2NWMwOjQwOjQ6NWYzOjU0YzQ6OGQxMDo5Yjk4XQ


## dnscry.pt-anon-dhaka-ipv4

dnscry.pt Dhaka relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ0xMDMuMTc0LjUxLjcx


## dnscry.pt-anon-dhaka-ipv6

dnscry.pt Dhaka relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gRVbMjAwMTpkZjE6OGY0MDo1MTo6YV0


## dnscry.pt-anon-dublin-ipv4

dnscry.pt Dublin relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ0xOTQuMjYuMjEzLjE1


## dnscry.pt-anon-dublin-ipv6

dnscry.pt Dublin relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gRVbMmEwOTpjZDQ2OmY6NDI5ZTo6NV0


## dnscry.pt-anon-durham-ipv4

dnscry.pt Durham relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQwzOC40NS42NC4xMTc


## dnscry.pt-anon-durham-ipv6

dnscry.pt Durham relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gR1bMjAwMTo1NTA6NWEwMDo1ZWI6OmRiNTpmYWNlXQ


## dnscry.pt-anon-dusseldorf-ipv4

dnscry.pt Düsseldorf relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ01LjE3NS4xODAuMTcx


## dnscry.pt-anon-ebenecity-ipv4

dnscry.pt Ebène City relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ4xMDIuMjIyLjEwNi45Ng


## dnscry.pt-anon-ebenecity-ipv6

dnscry.pt Ebène City relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gR9bMmMwZjplOGY4OjIwMDA6MjMzOjo0MjU0OmM1YjJd


## dnscry.pt-anon-ebenecity02-ipv4

dnscry.pt Ebène City 02 relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQwxOTYuNDYuNTAuOTM


## dnscry.pt-anon-eygelshoven-ipv4

dnscry.pt Eygelshoven relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQw5My45NS4xMTUuMjE


## dnscry.pt-anon-eygelshoven-ipv6

dnscry.pt Eygelshoven relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ5bMmExMDpjYTgwOjphXQ


## dnscry.pt-anon-flint-ipv4

dnscry.pt Flint relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ8xNDcuMTg5LjE0MC4xMzY


## dnscry.pt-anon-flint-ipv6

dnscry.pt Flint relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gRtbMjYwNjo2NjgwOjI5OjE6OjU4NTk6YTM3Yl0


## dnscry.pt-anon-frankfurt02-ipv4

dnscry.pt Frankfurt 02 relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ00NS4xNDcuNTEuMTIz


## dnscry.pt-anon-fremont02-ipv4

dnscry.pt Fremont 02 relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQwxNjcuODguNDguMTg


## dnscry.pt-anon-gdansk-ipv4

dnscry.pt Gdańsk relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ04Mi4xMTguMjEuMTg5


## dnscry.pt-anon-geneva-ipv4

dnscry.pt Geneva relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQw0NS45MC41OS4xOTM


## dnscry.pt-anon-geneva-ipv6

dnscry.pt Geneva relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gRBbMmEwNTo5NDA2OjphZTFd


## dnscry.pt-anon-grandrapids-ipv4

dnscry.pt Grand Rapids relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ4xODUuMTY1LjQ0LjE2NA


## dnscry.pt-anon-hafnarfjordur-ipv4

dnscry.pt Hafnarfjordur relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ4xOTIuNzEuMjE4LjEyMQ


## dnscry.pt-anon-halifax-ipv4

dnscry.pt Halifax relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQwyMy4xOTEuODAuNDM


## dnscry.pt-anon-halifax-ipv6

dnscry.pt Halifax relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gRVbMjYwMjpmYzFjOmZhMDoyOTo6MV0


## dnscry.pt-anon-hanoi-ipv4

dnscry.pt Hanoi relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ0xMDMuMTk5LjE2Ljkz


## dnscry.pt-anon-helsinki-ipv4

dnscry.pt Helsinki relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ4zNy4yMjguMTI5LjE2MA


## dnscry.pt-anon-helsinki-ipv6

dnscry.pt Helsinki relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gRVbMmEwNjoxNzAwOjE6M2E6OmNiYV0


## dnscry.pt-anon-hongkong-ipv4

dnscry.pt Hong Kong relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQs5Ni45LjIyOC4yNw


## dnscry.pt-anon-hongkong02-ipv4

dnscry.pt Hong Kong 02 relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ4yMTYuMjUwLjk3LjE0OA


## dnscry.pt-anon-hongkong02-ipv6

dnscry.pt Hong Kong 02 relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gRdbMjQwNjplZjgwOjE6MjE3Mzo6YjVmXQ


## dnscry.pt-anon-hongkong03-ipv4

dnscry.pt Hong Kong 03 relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ40NS4xMjMuMTg4LjEyOQ


## dnscry.pt-anon-houston-ipv4

dnscry.pt Houston relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ4yMDkuMTM1LjE3MC41MQ


## dnscry.pt-anon-houston-ipv6

dnscry.pt Houston relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gRRbMjYwMjpmOWYzOjA6Mjo6MTkzXQ


## dnscry.pt-anon-hudiksvall-ipv4

dnscry.pt Hudiksvall relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ05NS4xNDMuMTk2LjE2


## dnscry.pt-anon-ikeja-ipv4

dnscry.pt Ikeja relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ0xNjcuODguNTEuMjQ1


## dnscry.pt-anon-ikeja-ipv6

dnscry.pt Ikeja relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gRhbMmEwMTplMjgxOmFjMDE6ZmQwZDo6MV0


## dnscry.pt-anon-indianapolis-ipv4

dnscry.pt Indianapolis relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ4yMy4xNjguMTM2LjE0NA


## dnscry.pt-anon-indianapolis-ipv6

dnscry.pt Indianapolis relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gRRbMjYwMjpmOWJkOjgwOjExOjphXQ


## dnscry.pt-anon-islamabad-ipv4

dnscry.pt Islamabad relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ4xMDMuOTkuMTMzLjExMA


## dnscry.pt-anon-islamabad-ipv6

dnscry.pt Islamabad relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gRRbMjAwMTpkZjI6ZDQwOjI5OjoyXQ


## dnscry.pt-anon-istanbul-ipv6

dnscry.pt Istanbul relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gRpbMmExMjplMzQyOjMwMDo6ZGFjYTo2M2VhXQ


## dnscry.pt-anon-jacksonville-ipv4

dnscry.pt Jacksonville relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ8xMDQuMjI1LjEyOS4xMDY


## dnscry.pt-anon-jacksonville-ipv6

dnscry.pt Jacksonville relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gRhbMjYwNzphNjgwOjQ6ZjAwMzo6ZWMzMl0


## dnscry.pt-anon-jakarta-ipv4

dnscry.pt Jakarta relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ4xNTEuMjQzLjIyMi45NA


## dnscry.pt-anon-jakarta-ipv6

dnscry.pt Jakarta relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gR5bMjQwNzo2YWMwOjM6NToxMjM0OjQzMjE6ODk6MV0


## dnscry.pt-anon-jena-ipv4

dnscry.pt Jena relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQs4MS43LjExLjI0Ng


## dnscry.pt-anon-johannesburg-ipv4

dnscry.pt Johannesburg relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ8xNjkuMjM5LjEyOC4xMjQ


## dnscry.pt-anon-johannesburg02-ipv4

dnscry.pt Johannesburg 02 relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ8xNjAuMTE5LjIzNC4xNTY


## dnscry.pt-anon-kansascity-ipv4

dnscry.pt Kansas City relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ0yMy4xNTAuNDAuMTIx


## dnscry.pt-anon-kyiv-ipv4

dnscry.pt Kyiv relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ0yMTcuMTIuMjIxLjYx


## dnscry.pt-anon-kyiv-ipv6

dnscry.pt Kyiv relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gRBbMmEwMjoyN2FkOjoyMDFd


## dnscry.pt-anon-kyiv02-ipv6

dnscry.pt Kyiv 02 relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gRVbMmEwMTpmNTAwOjI6MTUwMDo6YV0


## dnscry.pt-anon-lasvegas-ipv4

dnscry.pt Las Vegas relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQs2Ni4xODcuNC4zOQ


## dnscry.pt-anon-libertylake-ipv4

dnscry.pt Liberty Lake relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQwyMy4xODQuNDguMTk


## dnscry.pt-anon-libertylake-ipv6

dnscry.pt Liberty Lake relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gRRbMjYwMjpmYzI0OjIwOjQ4OjphXQ


## dnscry.pt-anon-lima02-ipv4

dnscry.pt Lima 02 relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQw4Ny4xMjEuOTkuMjM


## dnscry.pt-anon-lima02-ipv6

dnscry.pt Lima 02 relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gRNbMmEwMzo5MGMwOjU1NTo6NzJd


## dnscry.pt-anon-lisbon-ipv4

dnscry.pt Lisbon relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQwxNDMuMjAuMTIuMzI


## dnscry.pt-anon-lisbon-ipv6

dnscry.pt Lisbon relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gRRbMmEwZjpjNDQyOjgwMDA6OjMyXQ


## dnscry.pt-anon-lisbon02-ipv4

dnscry.pt Lisbon 02 relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQw5MS4yMDkuMTYuOTg


## dnscry.pt-anon-lisbon02-ipv6

dnscry.pt Lisbon 02 relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ9bMmEwZjpjNDQ0Ojo5OF0


## dnscry.pt-anon-london-ipv4

dnscry.pt London relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQw0NS42Ny44NC4xMzI


## dnscry.pt-anon-london-ipv6

dnscry.pt London relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gRNbMjQwMTo4MzYwOmEyOjQ6OmFd


## dnscry.pt-anon-losangeles-ipv4

dnscry.pt Los Angeles relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ4xMDQuMTU2LjE1NC4xMQ


## dnscry.pt-anon-losangeles-ipv6

dnscry.pt Los Angeles relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gRJbMjYwMjpmN2Y4Ojc6ZDo6YV0


## dnscry.pt-anon-losangeles02-ipv4

dnscry.pt Los Angeles 02 relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ4xMDQuMjAwLjY3LjE5NA


## dnscry.pt-anon-losangeles02-ipv6

dnscry.pt Los Angeles 02 relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gRdbMjYwMjpmZjc1Ojc6Yjc5OjpiNGI0XQ


## dnscry.pt-anon-luxembourg-ipv4

dnscry.pt Luxembourg relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQw0NS44MC4yMDkuNTU


## dnscry.pt-anon-madrid-ipv4

dnscry.pt Madrid relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ01LjEzNC4xMTguMTk4


## dnscry.pt-anon-manchester-ipv4

dnscry.pt Manchester relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ4yMTYuMjQ1LjE0MC4yMA


## dnscry.pt-anon-manchester-ipv6

dnscry.pt Manchester relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gRNbMmEwYTo4ZGMwOjYwNTg6OmFd


## dnscry.pt-anon-manila-ipv4

dnscry.pt Manila relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ0xMDMuMzguMjUxLjYw


## dnscry.pt-anon-marseille-ipv4

dnscry.pt Marseille relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ40NS4xNDAuMTY0LjEyNw


## dnscry.pt-anon-melbourne-ipv4

dnscry.pt Melbourne relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ4xMDMuMTA4LjIyOC4xNQ


## dnscry.pt-anon-miami-ipv4

dnscry.pt Miami relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ4xMjguMjU0LjIwNy41MA


## dnscry.pt-anon-milan-ipv4

dnscry.pt Milan relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ04Mi4xMTguMTYuMTIx


## dnscry.pt-anon-molln-ipv4

dnscry.pt Mölln relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ05MS4xMDguODAuMTU5


## dnscry.pt-anon-molln-ipv6

dnscry.pt Mölln relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gRNbMmEwNTo5MDE6NjoxMDQ4Ojpd


## dnscry.pt-anon-montreal-ipv4

dnscry.pt Montreal relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ8xNDcuMTg5LjEzNi4xODM


## dnscry.pt-anon-montreal-ipv6

dnscry.pt Montreal relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gRpbMjYwNjo2NjgwOjQ1OjE6OmY3OGM6OWIwXQ


## dnscry.pt-anon-moscow-ipv4

dnscry.pt Moscow relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ45My4xODMuMTA2LjIyMg


## dnscry.pt-anon-mumbai02-ipv6

dnscry.pt Mumbai 02 relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gRFbMjYwMjpmYTA4OjU6Ojc1XQ


## dnscry.pt-anon-munich-ipv4

dnscry.pt Munich relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ0xOTQuMzkuMjA1LjEw


## dnscry.pt-anon-munich-ipv6

dnscry.pt Munich relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gRpbMmEwYzo4ZmMwOjE3NDk6NjY6MTg6OjE2XQ


## dnscry.pt-anon-naaldwijk-ipv4

dnscry.pt Naaldwijk relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQs0NS45NS4zOC4yOQ


## dnscry.pt-anon-newcastle-ipv4

dnscry.pt Newcastle relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQs4Mi4yMi4yMC4zNA


## dnscry.pt-anon-newyork-ipv4

dnscry.pt New York relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ4xOTkuMTE5LjEzNy43NA


## dnscry.pt-anon-newyork-ipv6

dnscry.pt New York relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gRJbMjYwMjpmN2Y4OjI6Yzo6YV0


## dnscry.pt-anon-nuremberg-ipv4

dnscry.pt Nuremberg relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ0yMDIuNjEuMjM2LjY3


## dnscry.pt-anon-nuremberg-ipv6

dnscry.pt Nuremberg relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gSVbMmEwMzo0MDAwOjVjOjUxOjI0Yjk6NTFmZjpmZTgwOmYzYTdd


## dnscry.pt-anon-ogden-ipv4

dnscry.pt Ogden relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ4xMDcuMTgyLjE3My44Mw


## dnscry.pt-anon-ogden-ipv6

dnscry.pt Ogden relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gRhbMjYwNzpmMmQ4OjQwMWI6MTA0NTo6YV0


## dnscry.pt-anon-oradea-ipv4

dnscry.pt Oradea relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ8xODUuMjA3LjEyNS4xMDA


## dnscry.pt-anon-oradea-ipv6

dnscry.pt Oradea relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gR1bMmEwZDo4MTQ0OjA6ZjY6MjkxNTphZjowOjE4XQ


## dnscry.pt-anon-ottoville-ipv4

dnscry.pt Ottoville relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ8xMDQuMjM0LjIzMS4yMzk


## dnscry.pt-anon-ottoville-ipv6

dnscry.pt Ottoville relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gRNbMjYwMjpmOTUzOjY6MjU6OmFd


## dnscry.pt-anon-paris-ipv4

dnscry.pt Paris relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQs4OS4xMTcuMi4xNw


## dnscry.pt-anon-paris-ipv6

dnscry.pt Paris relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gR5bMjQwMjpkMGMwOjIyOjZjZDA6NDo0OjQ6NWI4MV0


## dnscry.pt-anon-perth-ipv4

dnscry.pt Perth relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ4yMDMuMjkuMjQwLjI0OQ


## dnscry.pt-anon-perth-ipv6

dnscry.pt Perth relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gSJbMjQwNDo5NDAwOjQ6MDoyMTY6M2VmZjpmZWU2OmE3NjJd


## dnscry.pt-anon-philadelphia-ipv4

dnscry.pt Philadelphia relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ0xNTQuMTYuMTU5LjIy


## dnscry.pt-anon-philadelphia-ipv6

dnscry.pt Philadelphia relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gRVbMjYwNDpiZjAwOjIxMDoxMjo6Ml0


## dnscry.pt-anon-phoenix-ipv4

dnscry.pt Phoenix relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ8xNzMuMjQ5LjIwNS4xOTg


## dnscry.pt-anon-portedwards-ipv4

dnscry.pt Port Edwards relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ8xNzYuMTExLjIxOS4xMjY


## dnscry.pt-anon-portland-ipv4

dnscry.pt Portland relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ8xMDMuMTI0LjEwNi4yMzM


## dnscry.pt-anon-portland-ipv6

dnscry.pt Portland relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gSFbMjQwMjpkMGMwOjE2OmExZTY6MDpiODkzOmJmNzpkZF0


## dnscry.pt-anon-prague-ipv4

dnscry.pt Prague relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ4xOTUuMTIzLjI0NS4xOQ


## dnscry.pt-anon-prague-ipv6

dnscry.pt Prague relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gRBbMmEwNTo5NDAzOjo5OTld


## dnscry.pt-anon-queretaro-ipv4

dnscry.pt Querétaro relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQw4OS4yMjMuODguNzQ


## dnscry.pt-anon-queretaro-ipv6

dnscry.pt Querétaro relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gRRbMmEwMzo5MGMwOjU0NTo6MTFhXQ


## dnscry.pt-anon-redditch-ipv4

dnscry.pt Redditch relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQw0NS42Ny44NS4yMTk


## dnscry.pt-anon-redditch-ipv6

dnscry.pt Redditch relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gRNbMjQwMTo4MzYwOmEzOmU6OmFd


## dnscry.pt-anon-riga-ipv4

dnscry.pt Riga relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ8xOTUuMTIzLjIxMi4yMDA


## dnscry.pt-anon-saltlakecity-ipv4

dnscry.pt Salt Lake City relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ4xMDMuMTE0LjE2Mi42NQ


## dnscry.pt-anon-saltlakecity-ipv6

dnscry.pt Salt Lake City relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gSFbMjQwMjpkMGMwOjE4OmM4ZmY6MDpiODkzOmJmNzpkZF0


## dnscry.pt-anon-sandefjord-ipv4

dnscry.pt Sandefjord relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ0xOTQuMzIuMTA3LjQ4


## dnscry.pt-anon-sandefjord-ipv6

dnscry.pt Sandefjord relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gRVbMmEwMzo5NGUwOjI3MWY6OjViMV0


## dnscry.pt-anon-sanjose-ipv4

dnscry.pt San Jose relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ4xODUuMTA2Ljk2LjIxMA


## dnscry.pt-anon-sanjose-ipv6

dnscry.pt San Jose relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gRtbMjYwNzpmMzU4OjFhOmU6OjhhYTI6OTMzM10


## dnscry.pt-anon-santaclara-ipv4

dnscry.pt Santa Clara relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ8xNzYuMTExLjIyMy4xNjc


## dnscry.pt-anon-santaclara-ipv6

dnscry.pt Santa Clara relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gRtbMjYwNjo2NjgwOjM1OjE6OjUwNmQ6OGNlMl0


## dnscry.pt-anon-seattle-ipv4

dnscry.pt Seattle relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ8yMDkuMTgyLjIyNS4xMDM


## dnscry.pt-anon-seattle-ipv6

dnscry.pt Seattle relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gRhbMjYwNzphNjgwOjk6ZjAwNTo6ODZlN10


## dnscry.pt-anon-seoul-ipv4

dnscry.pt Seoul relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ05Mi4zOC4xMzUuMTI4


## dnscry.pt-anon-seoul02-ipv4

dnscry.pt Seoul 02 relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ8xNTEuMjQ1LjEwNi4xODE


## dnscry.pt-anon-singapore-ipv4

dnscry.pt Singapore relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ4xNTcuMjAuMTA1LjExNQ


## dnscry.pt-anon-singapore-ipv6

dnscry.pt Singapore relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gRVbMjYwNjpmYzQwOjQwMDM6Zjo6YV0


## dnscry.pt-anon-singapore02-ipv4

dnscry.pt Singapore 02 relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ0xMDMuMTc5LjQ0Ljcz


## dnscry.pt-anon-singapore02-ipv6

dnscry.pt Singapore 02 relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gRNbMjQwMTo0NTIwOjExMjI6OmFd


## dnscry.pt-anon-sofia-ipv4

dnscry.pt Sofia relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQs3OS4xMjQuNy40Mw


## dnscry.pt-anon-spokane-ipv4

dnscry.pt Spokane relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ0xMDQuMzYuODYuMTgx


## dnscry.pt-anon-spokane-ipv6

dnscry.pt Spokane relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gRRbMjYwNjphOGMwOjM6MjAyOjphXQ


## dnscry.pt-anon-stockholm02-ipv4

dnscry.pt Stockholm 02 relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ8xODUuMjMxLjEwMC4xMDY


## dnscry.pt-anon-sydney02-ipv4

dnscry.pt Sydney 02 relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ0xOTUuMTE0LjE0Ljc0


## dnscry.pt-anon-sydney02-ipv6

dnscry.pt Sydney 02 relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gRhbMjQwMjo3MzQwOjUwMDA6NjIwMDo6YV0


## dnscry.pt-anon-taipeh-ipv4

dnscry.pt Taipeh relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ8xMDMuMTMxLjE4OS4xOTE


## dnscry.pt-anon-tallinn-ipv4

dnscry.pt Tallinn relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ0xODUuMTk0LjUzLjIy


## dnscry.pt-anon-tampa-ipv4

dnscry.pt Tampa relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ8xNjUuMTQwLjExNy4yNDg


## dnscry.pt-anon-tampa-ipv6

dnscry.pt Tampa relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gRlbMjYwMjpmY2MwOjIyMjI6OWQyZTo6NTNd


## dnscry.pt-anon-taos-ipv4

dnscry.pt Taos relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ42My4xMzMuMjIzLjEzOA


## dnscry.pt-anon-tbilisi-ipv6

dnscry.pt Tbilisi relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gRpbMmExMjplMzQwOjMwMDo6MTc2ODphOTVmXQ


## dnscry.pt-anon-telaviv-ipv4

dnscry.pt Tel Aviv relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQw1LjE4OC4yMjcuMTM


## dnscry.pt-anon-telaviv-ipv6

dnscry.pt Tel Aviv relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gRNbMmEwMzo5MGMwOjFlNzo6Mzld


## dnscry.pt-anon-thessaloniki-ipv4

dnscry.pt Thessaloniki relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQs4NS45MC4xOTcuNw


## dnscry.pt-anon-thessaloniki-ipv6

dnscry.pt Thessaloniki relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gRRbMmExMjo2ZmMzOjgwMDA6OjE5XQ


## dnscry.pt-anon-timisoara-ipv4

dnscry.pt Timișoara relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQw0NS4xMzQuNDguMjU


## dnscry.pt-anon-tirana-ipv4

dnscry.pt Tirana relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ0xODUuNzUuMjQzLjgw


## dnscry.pt-anon-tokyo-ipv4

dnscry.pt Tokyo relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQw0NS42Ny44Ni4xMjM


## dnscry.pt-anon-tokyo-ipv6

dnscry.pt Tokyo relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gRVbMjYwNjpmYzQwOjQwMDI6ZDo6YV0


## dnscry.pt-anon-tokyo02-ipv4

dnscry.pt Tokyo 02 relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQwxMDMuMTc5LjQ1LjY


## dnscry.pt-anon-tokyo02-ipv6

dnscry.pt Tokyo 02 relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gRNbMjQwMTo0NTIwOjMwMWU6OmFd


## dnscry.pt-anon-toronto-ipv4

dnscry.pt Toronto relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ4xNzIuOTMuMTY3LjIxNA


## dnscry.pt-anon-toronto-ipv6

dnscry.pt Toronto relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gRhbMjYwNjo2MDgwOjIwMDE6MTA5OTo6YV0


## dnscry.pt-anon-toronto02-ipv4

dnscry.pt Toronto 02 relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ4xMDMuMTQ0LjE3Ny41Nw


## dnscry.pt-anon-tuusula-ipv4

dnscry.pt Tuusula relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ02NS4yMS4yNTIuMjAx


## dnscry.pt-anon-tuusula-ipv6

dnscry.pt Tuusula relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gRdbMmEwMTo0Zjk6YzAxMTpiODRlOjoxXQ


## dnscry.pt-anon-valdivia-ipv4

dnscry.pt Valdivia relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ0yMTYuNzMuMTU5LjI2


## dnscry.pt-anon-valdivia-ipv6

dnscry.pt Valdivia relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gRVbMmEwNjphMDA2OmQxZDE6OjExNl0


## dnscry.pt-anon-vancouver-ipv4

dnscry.pt Vancouver relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQwyMy4xNTQuODEuOTI


## dnscry.pt-anon-vilnius-ipv4

dnscry.pt Vilnius relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ0xNjIuMjU0Ljg2LjEz


## dnscry.pt-anon-warsaw-ipv4

dnscry.pt Warsaw relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ0xOTUuMy4yMjEuMTYy


## dnscry.pt-anon-yerevan-ipv4

dnscry.pt Yerevan relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQ04NS45MC4yMDcuMTk5


## dnscry.pt-anon-yerevan-ipv6

dnscry.pt Yerevan relay.
IPv6 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gRlbMmEwMzo5MGMwOjVmMToyOTAzOjo1Mzld


## dnscry.pt-anon-zurich-ipv4

dnscry.pt Zürich relay.
IPv4 endpoint. Operated by dnscry.pt. No query logs, no intentional filtering, blocking, redirecting or rewriting, DNSSEC validation.
https://www.dnscry.pt

sdns://gQs5NC4yNi4zNS41Nw

