#!/bin/bash

CHOICE="$1"

set -eu

COUNT=0
for FOLDER in ./y*
do
    PACKAGE="${FOLDER##*/}"
    if [ ! -d "$PACKAGE" ]; then
        continue
    fi

    printf "###### Working on package '%s': %s ######\n" "$PACKAGE" "$CHOICE"
    case "$CHOICE" in
    all)
        make build PACKAGE="$PACKAGE"
        make bench PACKAGE="$PACKAGE"
        make cover PACKAGE="$PACKAGE"
        make doc PACKAGE="$PACKAGE"
        ;;
    ci)
        make build PACKAGE="$PACKAGE"
        make test PACKAGE="$PACKAGE"
        make bench PACKAGE="$PACKAGE"
        ;;
    *)
        printf "unknown build option: [%s]\n" "$CHOICE"
        exit 1
        ;;
    esac
    echo ""

    COUNT=$((COUNT+1))
done

printf "====== Handled %d package(s) ======\n" "$COUNT"
