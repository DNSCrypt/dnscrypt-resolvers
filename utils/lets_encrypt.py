#!/usr/bin/env python3
# /// script
# dependencies = ["cryptography"]
# ///
"""
Hashes that every DoH stamp of a Let's Encrypt server must include.

A server can switch to another Let's Encrypt chain when its certificate is renewed,
so its stamps pin all of them.

CONTRIBUTING.md has the same list for server operators. Keep both in sync.

To check the hashes against the certificates Let's Encrypt publishes:

    uv run utils/lets_encrypt.py
"""

import hashlib
import sys
import urllib.request
from datetime import datetime, timezone
from typing import NamedTuple

from cryptography import x509

LETS_ENCRYPT_CERTS = {
    "https://letsencrypt.org/certs/gen-y/root-x2-by-x1.der":
        "8c5033973ffe4ff1ef0aa92984fc93281ba7a9278ef8bf641226b898d10d4e74",
    "https://letsencrypt.org/certs/gen-y/root-ye-by-x2.der":
        "369d1cf858e309c02aa2824c3763e12c05975e03291bff2bb462567bcbe93f21",
    "https://letsencrypt.org/certs/gen-y/root-yr-by-x1.der":
        "32edbb11cc258fbe33c2cd24a149270f0be6f446f57264b3bd1c25b98533272e",
}

LETS_ENCRYPT_HASHES = [bytes.fromhex(h) for h in LETS_ENCRYPT_CERTS.values()]


class PeerCert(NamedTuple):
    cn: str
    tbs_hash: bytes
    is_ca: bool
    is_self_issued: bool
    organization: str


def lets_encrypt_pins(chains: list[list[PeerCert]], selected: list[bytes]) -> list[bytes] | None:
    """
    Return all the Let's Encrypt hashes if the server uses Let's Encrypt, or None otherwise.

    `selected` holds the hashes picked from the chains.
    If one of them isn't on the list, the list is probably outdated, so this raises an error.
    """
    if not any(c.is_ca and c.organization == "Let's Encrypt" for chain in chains for c in chain):
        return None
    unlisted = [h.hex() for h in selected if h not in LETS_ENCRYPT_HASHES]
    if unlisted:
        raise RuntimeError(
            f"Let's Encrypt chain ends at {', '.join(unlisted)}, "
            "which is missing from lets_encrypt.py; that list may be outdated"
        )
    return LETS_ENCRYPT_HASHES


def main():
    ok = True
    for url, expected in LETS_ENCRYPT_CERTS.items():
        with urllib.request.urlopen(url, timeout=15) as response:
            cert = x509.load_der_x509_certificate(response.read())
        published = hashlib.sha256(cert.tbs_certificate_bytes).hexdigest()
        expires = cert.not_valid_after_utc
        if published != expected:
            status, ok = f"MISMATCH (published hash is {published})", False
        elif expires < datetime.now(timezone.utc):
            status, ok = "EXPIRED", False
        else:
            status = "OK"
        print(f"{status}: {cert.subject.rfc4514_string()}, issued by "
              f"{cert.issuer.rfc4514_string()}, expires {expires.date()}")
    sys.exit(0 if ok else 1)


if __name__ == "__main__":
    main()
