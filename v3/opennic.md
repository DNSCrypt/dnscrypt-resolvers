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

BebasDNS default server by BebasID. DNSSEC and OpenNIC supported. Filters ads, tracker, and malware.

sdns://AQMAAAAAAAAAEjEwMy44Ny42OC4xOTQ6ODQ0MyAxXDKkdrOao8ZeLyu7vTnVrT0C7YlPNNf6trdMkje7QR8yLmRuc2NyeXB0LWNlcnQuZG5zLmJlYmFzaWQuY29t


## bebasdns-family

BebasDNS Family Variant by BebasID. DNSSEC and OpenNIC supported. Blocks malicious links, pornography, gambling, and hate websites.

sdns://AQMAAAAAAAAAEjEwMy44Ny42OC4xOTY6ODQ0MyD5k4vgIHmBCZ2DeLtmoDVu1C6nVrRNzSVgZ1T0m0-3rCkyLmRuc2NyeXB0LWNlcnQuaW50ZXJuZXRzZWhhdC5iZWJhc2lkLmNvbQ


## bebasdns-security

BebasDNS Security Variant by BebasID. DNSSEC and OpenNIC supported. Only blocks malicious links.

sdns://AQMAAAAAAAAAEjEwMy44Ny42OC4xOTU6ODQ0MyDxbZzPMadetG2FodrzRfoiJjJi3cxbOsvKAvMyJ09rfiUyLmRuc2NyeXB0LWNlcnQuYW50aXZpcnVzLmJlYmFzaWQuY29t


## bebasdns-unfiltered-doh

BebasDNS by BebasID. DNSSEC and OpenNIC supported. This variant doesn't block anything

sdns://AgcAAAAAAAAADTEwMy44Ny42OC4xOTQAD2Rucy5iZWJhc2lkLmNvbQsvdW5maWx0ZXJlZA


## doh-ibksturm

DoH & DoT Server, No Logging, No Filters, DNSSEC

Running privately by ibksturm in Thurgau, Switzerland

sdns://AgcAAAAAAAAADjIxMy4xOTYuMTkxLjk2IFk-LUmnQCNVVSau-bdCKxnluLFnORtt7l7SkrHKI6dMGGlia3N0dXJtLnN5bm9sb2d5Lm1lOjQ0MwovZG5zLXF1ZXJ5

## ibksturm

Dnscrypt Server, No Logging, No Filters, DNSSEC, OpenNIC

Running privately by ibksturm in Thurgau, Switzerland

sdns://AQcAAAAAAAAAEzIxMy4xOTYuMTkxLjk2Ojg0NDMgiwvumeI8et789m3naGH-4xzM40t6c2xO_fCxHldawJgYMi5kbnNjcnlwdC1jZXJ0Lmlia3N0dXJt


## tirapan-doh-ipv4

Recursive and caching resolver, no-log/filter, DNSSEC, OpenNIC. Homepage: https://www.tirapan.top

sdns://AgcAAAAAAAAADDkxLjEwNy4yMzUuMCARsQLmsfY-UomE1gJfMrE4JB_Ii711GVdNcMmDLVPh6A93d3cudGlyYXBhbi50b3AQL25vbWVuLXF1YWVzaXR1bQ


## tirapan-doh-ipv6

Recursive and caching resolver, no-log/filter, DNSSEC, OpenNIC, IPv6 variant. Homepage: https://www.tirapan.top

sdns://AgcAAAAAAAAAF1syYTAxOjRmODoxYzFlOjhjYmE6OjFdIBGxAuax9j5SiYTWAl8ysTgkH8iLvXUZV01wyYMtU-HoD3d3dy50aXJhcGFuLnRvcBAvbm9tZW4tcXVhZXNpdHVt

