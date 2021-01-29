#! /bin/sh

curl -qL https://github.com/jedisct1/dnscrypt-proxy/releases/download/2.0.45/dnscrypt-proxy-linux_x86_64-2.0.45.tar.gz | tar xzvf -
cd linux-x86_64 || exit 1

NEW_ENTRIES="new-entries.txt"
git fetch --all
git diff origin/master -- ../v3 | grep -F '+sdns://' | cut -d'+' -f2- | sort >"$NEW_ENTRIES"
if [ ! -s "$NEW_ENTRIES" ]; then
    echo "No new entries found"
    exit 0
fi

DUPLICATES="duplicates.txt"
uniq -d "$NEW_ENTRIES" >"$DUPLICATES"
if [ -s "$DUPLICATES" ]; then
    echo "** DUPLICATES FOUND **"
    cat "$DUPLICATES"
    exit 1
fi

exit_code=0

CONFIG="test-dnscrypt-proxy.toml"
PIDFILE="dnscrypt-proxy.pid"
LOGFILE="dnscrypt-proxy.log"
while read -r stamp; do
    echo
    echo "* Checking resolver with stamp:"
    echo "$stamp"
    echo
    cp example-dnscrypt-proxy.toml "$CONFIG"
    sed -i -e 's/listen_addresses.*/listen_addresses = ["127.0.0.1:5300"]/' "$CONFIG"
    sed -i -e 's/# *server_names.*/server_names = ["test"]/' "$CONFIG"
    echo '[static."test"]' >>"$CONFIG"
    echo "stamp = '${stamp}'" >>"$CONFIG"
    ./dnscrypt-proxy -config "$CONFIG" -pidfile "$PIDFILE" -logfile "$LOGFILE" -loglevel 1 &
    sleep 5
    if grep -q 'ERROR.*\[.*:.*]:' "$LOGFILE"; then
        echo "(skipping due to IPv6 not being supported by GitHub Actions)"
    elif ! ./dnscrypt-proxy -config "$CONFIG" -resolve example.com; then
        echo "** UNABLE TO GET A RESPONSE FROM THE RESOLVER **"
        echo "Bogus stamp: ${stamp}"
        exit_code=1
    fi
    kill $(cat "$PIDFILE")
    cat "$LOGFILE"
    if grep -v 'ERROR.*\[.*:.*]:' "$LOGFILE" | grep -q 'ERROR|CRITICAL|FATAL'; then
        echo "** ERRORS FOUND **"
        exit_code=1
    fi
    echo "Done!"
    echo
done <"$NEW_ENTRIES"

if [ $exit_code != 0 ]; then
    echo "** ONE OR MORE CHECKS FAILED **"
fi

exit $exit_code
