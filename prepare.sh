#!/bin/bash

set -e
# get system temp dir
for TMPDIR in "$TMPDIR" "$TMP" /var/tmp /tmp
do
    test -d "$TMPDIR" && break
done

set -eu
# set platform name
system_name=$(uname -s)
if [[ $system_name == MINGW64_NT* ]] ; then
    platform_name="WINDOWS"
elif [[ $system_name == Linux* ]] ; then
    platform_name="LINUX"
elif [[ $system_name == Darwin* ]] ; then
    platform_name="MACOS"
else
    platform_name="UNKNOWN"
fi
export OS_NAME="$platform_name"

# uncompress test resource to temp dir
export MSYS=winsymlinks:native
export TESTRSSDIR=${TMPDIR%/}/gut_test_resource
rm -fr "$TESTRSSDIR"
unzip -o test_resource.zip -d "$TESTRSSDIR"
printf "Uncompress test resource: %s\n\n" "$TESTRSSDIR"
