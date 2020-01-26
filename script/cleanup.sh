#!/bin/bash

set -e

# remove ram disk
if [ $OS_NAME == "MACOS" ] || [ $OS_NAME == "LINUX" ]; then
    script/ramdisk.sh destroy GutRamDisk
    script/ramdisk.sh destroy GutReadOnlyDisk
fi

# remove test resources
if [[ ! -z "$TESTRSSDIR" ]]; then
    chmod -R 700 "$TESTRSSDIR" && rm -fr "$TESTRSSDIR"
fi
