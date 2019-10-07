#!/bin/bash

set -eu

COUNT=0

for FOLDER in ./y*
do
    PACKAGE="${FOLDER##*/}"
    if [ ! -d "$PACKAGE" ]; then
        continue
    fi

    printf "###### Working on package '%s' ######\n" "$PACKAGE"
    make build PACKAGE="$PACKAGE"
    make test PACKAGE="$PACKAGE"
    make bench PACKAGE="$PACKAGE"
    echo ""

    ((COUNT++))
done

printf "====== Handled %d package(s) ======\n" "$COUNT"
