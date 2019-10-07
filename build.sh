#!/bin/bash

PACKAGES=(yrand)

set -eu

for i in "${!PACKAGES[@]}"
do
    PACKAGE="${PACKAGES[$i]}"
    if [ ! -d "$PACKAGE" ]; then
        printf "###### Directory '%s' not existing ######\n" "$PACKAGE"
        exit 1
    fi

    printf "\n###### Working on package '%s' ######\n" "$PACKAGE"
    make build PACKAGE="$PACKAGE"
    make test PACKAGE="$PACKAGE"
    make bench PACKAGE="$PACKAGE"
done
