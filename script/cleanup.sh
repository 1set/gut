#!/bin/bash

set -e

# remove ram disk
./ramdisk.sh destroy GutRamDisk
./ramdisk.sh destroy GutReadOnlyDisk

# remove test resources
if [[ ! -z "$TESTRSSDIR" ]] ; then
    chmod -R 700 "$TESTRSSDIR" && rm -fr "$TESTRSSDIR"
fi
