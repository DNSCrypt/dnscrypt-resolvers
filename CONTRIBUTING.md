# Submitting update and new servers

Servers from these lists are expected to be reliable, maintained and freely and publicly accessible from anywhere.

It's highly recommended to only submit servers that you use yourself on a regular basis. Once a server has been submitted, you're expected to send updates if the server information changes, or if the service is affected by temporary or recurring outages.

If you are submitting a server that you don't operate, make sure that the server is okay with the additional volume of queries that being in these lists is going to bring.

Every entry includes:

- A name: this is the name that users will configure in their software. It has to be unique. If a service is accessible over IPv4 and IPv6, the name of the IPv6 service conventionally has an `-ipv6` suffix. And multiple variants of the same service should share the same prefix (for example: `exampledns`, `exampledns-noads`, `exampledns-parental-control`).
- A description. What makes the server different from other servers? What are the main properties? Where is it located? Some software using that list only display the first line to users, so try to summarize everything in the first line. The description can include links, but don't use Markdown formatting. Keep the description in plain text.
- The DNS stamps to use the server.

An entry can include multiple DNS stamps. One of more of them will be randomly chosen by client software. Their order is not relevant.

Services accessible over IPv4 and IPv6 must have distinct entries. Don't mix IPv6-only stamps with IPv4 stamps.

Frontends to other public services (ex: using Cloudflare or Google as a resolver) are not allowed. Beyond collecting additional data, their value is unclear. To hide client IP addresses, running proper DNS relays is recommended instead.

## DNSCrypt relays (DNS anonymizers)

Relays play an important role in DNS privacy. They prevent DNS operators from seeing client IP addresses, and makes device fingerpriting and query linkability more difficult.

If you are running [`encrypted-dns-server`](https://github.com/DNSCrypt/encrypted-dns-server), either directly or via the [DNSCrypt server Docker image](https://github.com/dnscrypt/dnscrypt-server-docker), you can enable support for anonymized DNS.

At startup, the service prints the stamp of the DNS relay, which looks like this: `sdns://gRIyMTIuNDcuMjI4LjEzNjo0NDM`.

IPv4 and IPv6 anonymizers should be part of distinct entries.

The new relay should only be added to the `v3/relays.md` file.

Per convention, relays start with an `anon-` prefix.

## DNSCrypt servers

[`encrypted-dns-server`](https://github.com/DNSCrypt/encrypted-dns-server) prints the stamps at startup. Other software may or may not print them.

Regardless, you should edit them to set the proper attributes for it. That can be easily done using the [Online DNS Stamps calculator](https://dnscrypt.info/stamps/).

- `DNSSEC` means that the server supports DNSSEC, both for upstream and downstream queries. 
- `no filter` means that responses received from upstream servers are not blocked or semantically changed. For example, a server that blocks ads cannot have that flag set.
- `no logs` means that the server doesn't store logs. Client IPs and queries can only be kept for rate limiting/abuse control, only for a couple minutes, and cannot be permanently stored. If the log policy is not known, the `no logs` flag must not be set.

The new server should only be added to the `v3/public-resolvers.md` file.

## DoH servers

If you are operating the DoH server, check out the [operational recommendations for DoH servers](https://github.com/DNSCrypt/doh-server?tab=readme-ov-file#operational-recommendations) first.

The [DoH server](https://github.com/DNSCrypt/doh-server) prints the DoH and ODoH DNS stamps at startup.
Other software may or may not print them. So, use the [Online DNS Stamps calculator](https://dnscrypt.info/stamps/) to edit or compute them.

If the server has a fixed IP address, enter it in the relevant form field. This will allow the server to be used without depending on 3rd party servers for bootstraping.

The DoH protocol is fragile and unless the server frequently switches certificates, it is important to also include [certificate hashes](https://github.com/DNSCrypt/doh-server?tab=readme-ov-file#dns-stamp-and-certificate-hashes).

Prior to submitting a new entry to that list, take the time to test the server with `dnscrypt-proxy`, with and without `HTTP/3` (if supported by the server).

The new server should only be added to the `v3/public-resolvers.md` file.

If a server supports both DNSCrypt and DoH, these should be in distinct entries. Typically, the DoH server has the same name as the DNSCrypt server, with a `-doh` suffix.

## Oblivious DoH

Modern DoH servers that allow the client IP to be hidden by a relay can be added to the `v3/odoh-servers.md` list.

The stamp must have the `Oblivious DoH target` type, and only the server properties have to be set.

[DoH server](https://github.com/DNSCrypt/doh-server) supports ODoH out of the box. Other software may or may not.

Regardless, you should test that the server is properly accessible via a DoH relay prior to submitting it to the list.

## OpenNIC and servers for parental control

The `opennic.md` and `parental-control.md` files are automatically built from the main `v3/public-resolvers.md` file.

If the server supports the OpenNIC TLD, add it to `v3/opennic.subset`. If the server does parental control, add it to `v3/parental-control.md`.

## Monitoring the service status

The servers and relays are constantly monitored for their availability. [The DNS status files](https://download.dnscrypt.info/resolvers-list/status/) are constantly updated with their status.

You can use a tool such as [Uptime Robot](http://uptimerobot.com) to periodically check the presence of `pass: <your service name>` in these lists.

## Sending a pull request

To make a change to these lists, fork this repository and **open a pull request** after having followed the above guidelines.

A few check will run, directly from GitHub. If they are not passing, you can inspect their output to figure out why the service you added or modified couldn't be used.

If all the tests are passing, the changes will be manually reviewed, and eventually merged.