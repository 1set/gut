#!/bin/bash

set -e

CHOICE="$1"
TARGET="$2"

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"

# define target
if [[ -z "$TARGET" ]]; then
    TARGET="./y*"
else
    TARGET=("$TARGET")
fi

# prepare environment
printf "\n###### Prepare ######\n"
source "$SCRIPT_DIR"/prepare.sh

printf "\n====== Begin at %s, OS: %s, Mode: %s - %s ======\n" "$(date '+%Y-%m-%d %H:%M:%S %z')" "$OS_NAME" "$CHOICE" "$TARGET"

set -u
COUNT=0
for FOLDER in $TARGET; do
    PACKAGE="${FOLDER##*/}"
    if [[ ! -d "$PACKAGE" ]]; then
        continue
    fi

    if [[ "ci" == "$CHOICE" && 0 -eq $COUNT ]]; then
        printf "\n###### Go Environment ######\n"
        go env
    fi

    printf "\n###### Working on package '%s' ######\n" "$PACKAGE"
    case "$CHOICE" in
    all)
        make fmt PACKAGE="$PACKAGE"
        make build PACKAGE="$PACKAGE"
        make test PACKAGE="$PACKAGE"
        make bench PACKAGE="$PACKAGE"
        make doc PACKAGE="$PACKAGE"
        ;;
    ci)
        make build PACKAGE="$PACKAGE"
        make test PACKAGE="$PACKAGE"
        make bench PACKAGE="$PACKAGE"
        ;;
    dev)
        make fmt PACKAGE="$PACKAGE"
        make testdev PACKAGE="$PACKAGE"
        make benchdev PACKAGE="$PACKAGE"
        ;;
    fastdev)
        make fmt PACKAGE="$PACKAGE"
        make test PACKAGE="$PACKAGE"
        ;;
    *)
        printf "Unknown build option: [%s]\n" "$CHOICE"
        exit 1
        ;;
    esac

    COUNT=$((COUNT + 1))
done
set +u

printf "\n###### Clean Up ######\n"

source "$SCRIPT_DIR"/cleanup.sh

printf "\n====== End at %s, Packages: %d ======\n" "$(date '+%Y-%m-%d %H:%M:%S %z')" "$COUNT"
