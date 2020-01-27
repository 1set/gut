#!/bin/bash

set -e

# remove ram disk
if [[ $OS_NAME == "MACOS" ]]; then
    script/ramdisk.sh destroy GutRamDisk
    script/ramdisk.sh destroy GutReadOnlyDisk
elif [[ $OS_NAME == "LINUX" ]]; then
    sudo script/ramdisk.sh destroy GutRamDisk
    sudo script/ramdisk.sh destroy GutReadOnlyDisk
fi

# remove test resources
if [[ ! -z "$TESTRSSDIR" ]]; then
    chmod -R 700 "$TESTRSSDIR" && rm -fr "$TESTRSSDIR"
fi
