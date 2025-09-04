#! /bin/sh

# Check a single DNS stamp for availability
# Usage: ./check-stamp.sh sdns://...
# Exit codes: 0 = working, 1 = not working or error

DNSCRYPT_PROXY=~/src/dnscrypt-proxy/dnscrypt-proxy/dnscrypt-proxy
CONFIG="/tmp/dnscrypt-proxy-check.toml"
PIDFILE="/tmp/dnscrypt-proxy-check.pid"
LOGFILE="/tmp/dnscrypt-proxy-check.log"

# Check arguments
if [ $# -ne 1 ]; then
    echo "Usage: $0 sdns://..." >&2
    exit 1
fi

STAMP="$1"

# Validate stamp format
if ! echo "$STAMP" | grep -q '^sdns://'; then
    echo "Error: Invalid stamp format. Must start with 'sdns://'" >&2
    exit 1
fi

# Check if dnscrypt-proxy is available
if [ ! -x "$DNSCRYPT_PROXY" ]; then
    echo "Error: dnscrypt-proxy not found at $DNSCRYPT_PROXY" >&2
    exit 1
fi

# Clean up any previous runs
cleanup() {
    if [ -f "$PIDFILE" ]; then
        kill $(cat "$PIDFILE") 2>/dev/null
    fi
    rm -f "$CONFIG" "$PIDFILE" "$LOGFILE"
}
trap cleanup EXIT

# Create config file
{
    echo 'listen_addresses = ["127.0.0.1:5300"]'
    echo 'server_names = ["test-server"]'
    echo 'odoh_servers = true'
    echo 'timeout = 5000'
    echo 'keepalive = 30'
    echo
    echo '[static."test-server"]'
    echo "stamp = '$STAMP'"
} >"$CONFIG"

# Check DNSSEC support
DNSSEC=false
if $DNSCRYPT_PROXY -config "$CONFIG" -list -json 2>/dev/null | grep -F '"dnssec": true' >/dev/null; then
    DNSSEC=true
fi

# Show certificate info (silent mode)
if ! $DNSCRYPT_PROXY -config "$CONFIG" -show-certs >/dev/null 2>&1; then
    echo "Error: Failed to retrieve certificate information" >&2
    exit 1
fi

# Start dnscrypt-proxy
$DNSCRYPT_PROXY -config "$CONFIG" -pidfile "$PIDFILE" -logfile "$LOGFILE" -loglevel 3 &
sleep 1

# Check if process is running
if [ ! -f "$PIDFILE" ] || ! kill -0 $(cat "$PIDFILE") 2>/dev/null; then
    echo "Error: Failed to start dnscrypt-proxy" >&2
    if [ -f "$LOGFILE" ]; then
        tail -n 10 "$LOGFILE" >&2
    fi
    exit 1
fi

# Test resolver with retries
RETRIES=3
SUCCESS=false

for i in $(seq 1 $RETRIES); do
    if $DNSCRYPT_PROXY -config "$CONFIG" -resolve "example.com" >/tmp/resolve-output 2>/dev/null; then
        # Check DNSSEC if expected
        if [ "$DNSSEC" = "true" ]; then
            if grep -F "resolver doesn't support DNSSEC" /tmp/resolve-output >/dev/null; then
                echo "Error: DNSSEC support expected but not detected" >&2
                exit 1
            fi
        fi
        SUCCESS=true
        break
    fi
    [ $i -lt $RETRIES ] && sleep 1
done

# Clean up
kill $(cat "$PIDFILE") 2>/dev/null
rm -f /tmp/resolve-output

# Return result
if [ "$SUCCESS" = "true" ]; then
    echo "OK: Resolver is working"
    exit 0
else
    echo "FAIL: Unable to resolve queries" >&2
    exit 1
fi
