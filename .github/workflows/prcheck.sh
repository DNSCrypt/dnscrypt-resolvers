#! /bin/sh

DNSLOOKUP_VERSION=1.10.1

case "$(uname -ms)" in
Darwin\ x86_64) DNSLOOKUP_ARCH=darwin-amd64 ;;
Linux\ x86_64) DNSLOOKUP_ARCH=linux-amd64 ;;
*)
    echo "Unsupported platform" >&2
    exit 2
    ;;
esac

if [ ! -x "${DNSLOOKUP_ARCH}/dnslookup" ]; then
    curl -sL https://github.com/ameshkov/dnslookup/releases/download/v${DNSLOOKUP_VERSION}/dnslookup-${DNSLOOKUP_ARCH}-v${DNSLOOKUP_VERSION}.tar.gz | tar xzpf - || exit 1
fi
PATH="$(pwd)/${DNSLOOKUP_ARCH}:$PATH"

try_resolver() {
    ERROR_LOG_TMP=".errors"
    resolver_name="$1"
    stamp="$2"
    if dnslookup one.net "$stamp" >/dev/null 2>&1; then
        echo "pass: ${resolver_name}"
    elif dnslookup one.net "$stamp" >/dev/null 2>&1; then
        echo "pass: ${resolver_name} (1 retry)"
    elif dnslookup one.net "$stamp" >/dev/null 2>&1; then
        echo "pass: ${resolver_name} (2 retries)"
    elif dnslookup one.net "$stamp" >/dev/null 2>"$ERROR_LOG_TMP"; then
        echo "pass: ${resolver_name} (3 retries)"
    else
        if grep -Eq "(no route|unreachable)" "$ERROR_LOG_TMP"; then
            echo "ipv6: ${resolver_name}"
        else
            (
                echo "* FAILED: ${resolver_name}"
                echo "$stamp"
                cat "$ERROR_LOG_TMP"
                echo
            ) >&2
            return 1
        fi
    fi
}

for aux in v3/parental-control.md v3/opennic.md; do
    grep '^## ' "$aux" | while read -r entry; do
        if ! grep -Fq "$entry" v3/public-resolvers.md; then
            echo "Present in [$aux] but not in public-resolvers.md:"
            echo "$entry"
            exit 1
        fi
    done
done

DUPLICATES="duplicates.txt"
for aux in v3/*.md; do
    (
        grep '^##' "$aux" | tr A-Z a-z
        grep '^sdns://' "$aux"
    ) | sort | uniq -d >"$DUPLICATES"
    if [ -s "$DUPLICATES" ]; then
        echo "** DUPLICATES FOUND in [$aux] **"
        cat "$DUPLICATES"
        exit 1
    fi
done

NEW_ENTRIES="$(pwd)/new-entries.txt"
git fetch --all
git diff origin/master -- $(ls v3/*.md | grep -Ev 'onion|relay|odoh') | grep -F '+sdns://' | cut -d'+' -f2- | sort >"$NEW_ENTRIES"
if [ ! -s "$NEW_ENTRIES" ]; then
    echo "No new entries found"
    exit 0
fi

curl -qL https://github.com/jedisct1/dnscrypt-proxy/releases/download/2.1.5/dnscrypt-proxy-linux_x86_64-2.1.5.tar.gz | tar xzvf -
cd linux-x86_64 || exit 1

exit_code=0

CONFIG="test-dnscrypt-proxy.toml"
PIDFILE="dnscrypt-proxy.pid"
LOGFILE="dnscrypt-proxy.log"
while read -r stamp; do
    echo
    echo ========================
    echo
    echo "* Checking resolver with stamp:"
    echo "$stamp"
    echo

    try_resolver "(new entry)" "$stamp" || exit 1

    {
        echo 'listen_addresses = ["127.0.0.1:5300"]'
        echo 'http3 = true'
        echo 'server_names = ["test"]'
        echo '[static."test"]'
        echo "stamp = '${stamp}'"
    } >"$CONFIG"

    if ! ./dnscrypt-proxy -config "$CONFIG" -show-certs; then
        exit_code=1
    fi
    echo
    echo ---
    echo

    dnssec=false
    if ./dnscrypt-proxy -config "$CONFIG" -list -json | grep -F '"dnssec": true' >/dev/null; then
        dnssec=true
        echo "DNSSEC support is expected"
    else
        echo "DNSSEC support is not expected"
    fi

    ./dnscrypt-proxy -config "$CONFIG" -pidfile "$PIDFILE" -logfile "$LOGFILE" -loglevel 1 &

    sleep 5
    skip_log=false
    if grep -q 'DNSCrypt relay' "$LOGFILE"; then
        echo "(skipping due to IPv6 not being supported by GitHub Actions)"
        skip_log=true
    elif grep -q 'ERROR.*\[.*:.*]:' "$LOGFILE"; then
        echo "(skipping due to relays not being handled by this test)"
        skip_log=true
    elif ! ./dnscrypt-proxy -config "$CONFIG" -resolve example.com; then
        echo "** UNABLE TO GET A RESPONSE FROM THE RESOLVER **"
        echo "Bogus stamp: ${stamp}"
        exit_code=1
    elif $dnssec; then
        if ./dnscrypt-proxy -config "$CONFIG" -resolve -check example.com | grep -F "resolver doesn't support DNSSEC" >/dev/null; then
            echo "** DNSSEC SUPPORT IS EXPECTED BUT NOT DETECTED **"
            exit_code=1
        fi
    fi
    kill $(cat "$PIDFILE")
    if [ "$skip_log" = false ]; then
        cat "$LOGFILE"
        if grep -v 'ERROR.*\[.*:.*]:' "$LOGFILE" | grep -q 'ERROR|CRITICAL|FATAL'; then
            echo "** ERRORS FOUND **"
            exit_code=1
        fi
        echo "Done!"
    fi
    echo
done <"$NEW_ENTRIES"

if [ $exit_code != 0 ]; then
    echo "** ONE OR MORE CHECKS FAILED **"
fi

exit $exit_code
