# parental-control

A set of resolvers blocking popular websites that may not be appropriate
for children.

This is not bulletproof. In particular, websites in languages that are
not English will require additional, local rules.

To use that list, add this to the `[sources]` section of your
`dnscrypt-proxy.toml` configuration file:

    [sources.'parental-control']
    url = 'http://download.dnscrypt.info/resolvers-list/v2/parental-control.md'
    minisign_key = 'RWQf6LRCGA9i53mlYecO4IzT51TGPpvWucNSCh1CBM0QTaLn73Y7GFO3'
    cache_file = 'parental-control.md'

In order to enforce safe search results from Google and Youtube, you may
also want to enable cloaking (`cloaking_rules` in the configuration file).

--

## adguard-dns-family
Adguard DNS with safesearch and adult content blocking
sdns://AQIAAAAAAAAAFDE3Ni4xMDMuMTMwLjEzMjo1NDQzILgxXdexS27jIKRw3C7Wsao5jMnlhvhdRUXWuMm1AFq6ITIuZG5zY3J5cHQuZmFtaWx5Lm5zMS5hZGd1YXJkLmNvbQ


## cisco-familyshield
Block websites not suitable for children
sdns://AQAAAAAAAAAAEjIwOC42Ny4yMjAuMTIzOjQ0MyC3NRFAIG8iXT4r2CLX_WkeocM8yNZmjQy-BL-rykP7eRsyLmRuc2NyeXB0LWNlcnQub3BlbmRucy5jb20


## cleanbrowsing-family
Family filter - Blocks adult content on all languages + mixed content sites + proxies + enforces Safe Search engine on Google, Bing and Youtube.
By https://cleanbrowsing.org/
sdns://AQMAAAAAAAAAFDE4NS4yMjguMTY4LjE2ODo4NDQzILysMvrVQ2kXHwgy1gdQJ8MgjO7w6OmflBjcd2Bl1I8pEWNsZWFuYnJvd3Npbmcub3Jn


## cleanbrowsing-adult
Adult filter - Blocks adult content + enforces Safe Search engine mode.
By https://cleanbrowsing.org/
sdns://AQMAAAAAAAAAEzE4NS4yMjguMTY4LjEwOjg0NDMgvKwy-tVDaRcfCDLWB1AnwyCM7vDo6Z-UGNx3YGXUjykRY2xlYW5icm93c2luZy5vcmc


## doh-cleanbrowsing-family
Family filter - Blocks adult content on all languages + mixed content sites + proxies + enforces Safe Search engine on Google, Bing and Youtube.
By https://cleanbrowsing.org/
sdns://AgMAAAAAAAAAAAAVZG9oLmNsZWFuYnJvd3Npbmcub3JnEy9kb2gvZmFtaWx5LWZpbHRlci8


## doh-cleanbrowsing-adult
Adult filter - Blocks adult content + enforces Safe Search engine mode.
By https://cleanbrowsing.org/
sdns://AgMAAAAAAAAAAAAVZG9oLmNsZWFuYnJvd3Npbmcub3JnEi9kb2gvYWR1bHQtZmlsdGVyLw


## cleanbrowsing-family-ipv6
Family filter - Blocks adult content + mixed content sites + proxies + enforces safe search engine mode.
By https://cleanbrowsing.org/
sdns://AQMAAAAAAAAAFFsyYTBkOjJhMDA6MTo6XTo4NDQzILysMvrVQ2kXHwgy1gdQJ8MgjO7w6OmflBjcd2Bl1I8pEWNsZWFuYnJvd3Npbmcub3Jn


## cleanbrowsing-adult-ipv6
Family filter - Blocks adult content + mixed content sites + proxies + enforces safe search engine mode.
By https://cleanbrowsing.org/
sdns://AQMAAAAAAAAAFVsyYTBkOjJhMDA6MTo6MV06ODQ0MyC8rDL61UNpFx8IMtYHUCfDIIzu8Ojpn5QY3HdgZdSPKRFjbGVhbmJyb3dzaW5nLm9yZw

## sfw.scaleway-fr
Uses deep learning to block adult websites. Free, DNSSEC, no logs.
Hosted in Paris, running on a 1-XS server donated by Scaleway.com
Maintained by Frank Denis

sdns://AQMAAAAAAAAADzE2My4xNzIuMTgwLjEyNSDfYnO_x1IZKotaObwMhaw_-WRF1zZE9mJygl01WPGh_x8yLmRuc2NyeXB0LWNlcnQuc2Z3LnNjYWxld2F5LWZy
