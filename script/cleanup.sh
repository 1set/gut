#!/bin/bash

set -e
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"

# remove ram disk
if [[ $OS_NAME == "MACOS" ]]; then
    "$SCRIPT_DIR"/ramdisk.sh destroy GutRamDisk
    "$SCRIPT_DIR"/ramdisk.sh destroy GutReadOnlyDisk
elif [[ $OS_NAME == "LINUX" ]]; then
    sudo "$SCRIPT_DIR"/ramdisk.sh destroy GutRamDisk
    sudo "$SCRIPT_DIR"/ramdisk.sh destroy GutReadOnlyDisk
fi

# remove test resources
if [[ ! -z "$TESTRSSDIR" ]]; then
    chmod -R 700 "$TESTRSSDIR" && rm -fr "$TESTRSSDIR"
    printf "Remove test resource: ${TESTRSSDIR}\n"
fi
