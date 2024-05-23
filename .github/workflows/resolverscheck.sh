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

RESOLVERS_LIST="v3/public-resolvers.md"

ERROR_LOG=".error-log"
: >"$ERROR_LOG"

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
            ) >>"$ERROR_LOG"
            return 1
        fi
    fi
}

resolver_name=""
while read -r line; do
    case "$line" in
    \#\#\ *)
        resolver_name=$(echo "$line" | sed 's/^## *//')
        continue
        ;;
    sdns:*)
        try_resolver "$resolver_name" "$line"
        ;;
    esac
done <"$RESOLVERS_LIST"

if [ -s "$ERROR_LOG" ]; then
    exec 1>&2
    echo
    echo "** FAILURES **"
    echo
    cat "$ERROR_LOG"
    exit 0
fi
