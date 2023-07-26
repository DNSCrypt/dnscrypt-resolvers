# opennic

Resolvers from the [OpenNIC](https://www.opennic.org/) project.

To use that list, add this to the `[sources]` section of your
`dnscrypt-proxy.toml` configuration file:

    [sources.'opennic']
    urls = ['https://raw.githubusercontent.com/DNSCrypt/dnscrypt-resolvers/master/v3/opennic.md', 'https://download.dnscrypt.info/resolvers-list/v3/opennic.md']
    minisign_key = 'RWQf6LRCGA9i53mlYecO4IzT51TGPpvWucNSCh1CBM0QTaLn73Y7GFO3'
    cache_file = 'opennic.md'

--


## bebasdns

BebasDNS default server by BebasID. DNSSEC supported. Filters ads, tracker, and malware.

sdns://AgMAAAAAAAAADDEwMy44Ny42OC4yNCBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33zg9kbnMuYmViYXNpZC5jb20KL2Rucy1xdWVyeQ


## bebasdns-security

BebasDNS Security Variant by BebasID. DNSSEC supported. Only blocks malicious links.

sdns://AgMAAAAAAAAADDEwMy44Ny42OC4yMyBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33zhVhbnRpdmlydXMuYmViYXNpZC5jb20KL2Rucy1xdWVyeQ


## doh-ibksturm

DoH & DoT Server, No Logging, No Filters, DNSSEC

Running privately by ibksturm in Thurgau, Switzerland

sdns://AgcAAAAAAAAAAKDMEGDTnIMptitvvH0NbfkwmGm5gefmOS1c2PpAj02A5iBETr1nu4P4gHs5Iek4rJF4uIK9UKrbESMfBEz18I33zhRpYmtzdHVybS5zeW5vbG9neS5tZQovZG5zLXF1ZXJ5


## ibksturm

DNSCRYPT V2 Server, No Logging, No Filters, DNSSEC

Running privately by ibksturm in Thurgau, Switzerland

sdns://AQcAAAAAAAAAEzIxMy4xOTYuMTkxLjk2Ojg0NDMgHK0AUhqiLSuBFR07jpBhKvko_oyqyWnot8z4cce7cKkYMi5kbnNjcnlwdC1jZXJ0Lmlia3N0dXJt


## publicarray-au2-doh

DNSSEC • OpenNIC • Non-logging • Uncensored - hosted on ovh.com.au
Maintained by publicarray - https://dns.seby.io

sdns://AgcAAAAAAAAADTEzOS45OS4yMjIuNzKgzBBg05yDKbYrb7x9DW35MJhpuYHn5jktXNj6QI9NgOYgRE69Z7uD-IB7OSHpOKyReLiCvVCq2xEjHwRM9fCN984NZG9oLTIuc2VieS5pbwovZG5zLXF1ZXJ5

