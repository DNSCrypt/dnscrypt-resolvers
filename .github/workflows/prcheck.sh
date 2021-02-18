#! /bin/sh

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
git diff origin/master -- v3 | grep -F '+sdns://' | cut -d'+' -f2- | sort >"$NEW_ENTRIES"
if [ ! -s "$NEW_ENTRIES" ]; then
    echo "No new entries found"
    exit 0
fi

curl -qL https://github.com/jedisct1/dnscrypt-proxy/releases/download/2.0.45/dnscrypt-proxy-linux_x86_64-2.0.45.tar.gz | tar xzvf -
cd linux-x86_64 || exit 1

exit_code=0

CONFIG="test-dnscrypt-proxy.toml"
PIDFILE="dnscrypt-proxy.pid"
LOGFILE="dnscrypt-proxy.log"
while read -r stamp; do
    echo
    echo "* Checking resolver with stamp:"
    echo "$stamp"
    echo
    {
        echo 'listen_addresses = ["127.0.0.1:5300"]'
        echo 'server_names = ["test"]'
        echo '[static."test"]'
        echo "stamp = '${stamp}'"
    } >"$CONFIG"
    ./dnscrypt-proxy -config "$CONFIG" -pidfile "$PIDFILE" -logfile "$LOGFILE" -loglevel 1 &
    sleep 5
    skip_log=false
    if grep -q 'ERROR.*\[.*:.*]:' "$LOGFILE"; then
        echo "(skipping due to IPv6 not being supported by GitHub Actions)"
        skip_log=true
    elif ! ./dnscrypt-proxy -config "$CONFIG" -resolve example.com; then
        echo "** UNABLE TO GET A RESPONSE FROM THE RESOLVER **"
        echo "Bogus stamp: ${stamp}"
        exit_code=1
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
