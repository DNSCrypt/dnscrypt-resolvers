# Submitting updates and Adding new servers

### Server Requirements

- Servers from these lists are expected to be reliable, maintained, freely and publicly accessible from anywhere.
  - You should test the server is properly accessible and usable before submitting it.
- It is highly recommended that you submit only servers that you maintain.
  - Once a server has been submitted, we expect you to send updates if its information changes or the service is affected by temporary or recurring outages.
  - If you are submitting a server that you do not operate, make sure that its operator is fine with the additional volume of queries that inclusion in these lists will bring.
- **Frontends to other public services (for example, using Cloudflare or Google as a resolver) are not allowed.**
  - In addition to collecting more data, their value is unclear. To protect client IP addresses, we recommend running proper DNS resolvers.

### Every entry should include:

- **Name**:
  - The unique name that users will configure in their software.
  - If a service is accessible over IPv4 and IPv6, the name of the IPv6 service conventionally has an `-ipv6` suffix.
  - Multiple variants of the same service should share the same prefix. For example: `exampledns`, `exampledns-noads`, and `exampledns-parental-control`.
  - If a server supports both DNSCrypt and DoH, these should be in distinct entries. Typically, the DoH server has the same name as the DNSCrypt server, with a `-doh` suffix.
- **Description**:
  - What makes the server different from other servers? What are its main properties? Where is it located?
  - Some software that uses this list displays only the first line of the description, so try to summarize everything in that line.
  - You may include relevant URLs in the description.
  - Keep the description in plain text. Do not use Markdown; for example, don't use `[https://...](URL)`.
- **DNS stamps** to use for the server.
  - You should edit the stamps to set the proper attributes. This can easily be done using the [Online DNS Stamps calculator](https://dnscrypt.info/stamps/).
  - Read `DNS stamps attributes` for more information.

One entry can include multiple DNS stamps. One or more of DNS stamps will be randomly chosen by client software. Their order is not relevant.

Services accessible over IPv4 and IPv6 must have distinct entries. Do not mix IPv6-only stamps with IPv4 stamps.

The **new server** entry should only be added to the `v3/public-resolvers.md` file. You don't have to edit other files.

### DNS stamps attributes

- **DNSSEC**: The server supports DNSSEC for both upstream and downstream queries.
- **No filter**: Responses received from upstream servers are not blocked or semantically changed.
  - For example, a server that blocks ads cannot have this flag set.
- **No logs**: The server does not store logs.
  - Client IPs and queries may be kept only for rate limiting or abuse control, for no more than a couple of minutes, and cannot be stored permanently.
  - If the logging policy is not known, the `no logs` flag must not be set.

# Protocols

### DNSCrypt relays (DNS anonymizers)

Relays play an important role in DNS privacy. They prevent DNS operators from seeing client IP addresses and make device fingerprinting and query linkability more difficult.

If you are running [`encrypted-dns-server`](https://github.com/DNSCrypt/encrypted-dns-server), either directly or via the [DNSCrypt server Docker image](https://github.com/dnscrypt/dnscrypt-server-docker), you can enable support for anonymized DNS.

At startup, the service prints the stamp of the DNS relay, which looks like this: `sdns://gRIyMTIuNDcuMjI4LjEzNjo0NDM`.

By convention, relays start with an `anon-` prefix.

### DNSCrypt servers

[`encrypted-dns-server`](https://github.com/DNSCrypt/encrypted-dns-server) prints the stamps at startup. Other software may or may not print them.

### DoH servers

If you operate the DoH server, check out the [operational recommendations for DoH servers](https://github.com/DNSCrypt/doh-server?tab=readme-ov-file#operational-recommendations) first.

The [DoH server](https://github.com/DNSCrypt/doh-server) prints the DoH and ODoH DNS stamps at startup. Other software may or may not print them.

If the server has a fixed IP address, enter it in the relevant form field. This allows the server to be used without depending on third-party servers for bootstrapping.

The DoH protocol is fragile, and unless the server frequently switches certificates, it is important to also include [certificate hashes](https://github.com/DNSCrypt/doh-server?tab=readme-ov-file#dns-stamps-and-certificate-hashes).

Before submitting a new entry to the list, take the time to test the server with `dnscrypt-proxy`, both with and without `HTTP/3` if supported by the server.

### Oblivious DoH

Modern DoH servers that allow the client IP address to be hidden by a relay can be added to the `v3/odoh-servers.md` list.

The stamp must have the `Oblivious DoH target` type, and only the server properties need to be set.

The [DoH server](https://github.com/DNSCrypt/doh-server) supports ODoH out of the box. Other software may or may not.

### OpenNIC and servers for parental control

The `opennic.md` and `parental-control.md` files are built automatically from the main `v3/public-resolvers.md` file.

If the server supports the OpenNIC TLD, add it to `v3/opennic.subset`. If the server provides parental controls, add it to `v3/parental-control.md`.


# Before you send a pull request

The availability of the servers and relays is monitored continuously. [The DNS status files](https://download.dnscrypt.info/resolvers-list/status/) are updated continuously with their status.

You can use a tool such as [Uptime Robot](https://uptimerobot.com) to periodically check for the presence of `pass: <your service name>` in these lists.

### Sending a pull request

To make a change to these lists, fork this repository and **open a pull request** after following the guidelines above.

A few checks will run directly from GitHub. If they do not pass, you should inspect the output to determine why the service you added or modified could not be used.

If all the tests pass, the changes will be manually reviewed and eventually merged.
