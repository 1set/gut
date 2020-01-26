#!/bin/bash

set -e
# get system temp dir
for TMPDIR in "$TMPDIR" "$TMP" /var/tmp /tmp; do
    test -d "$TMPDIR" && break
done

# set platform name
system_name=$(uname -s)
if [[ $system_name == MINGW64_NT* ]]; then
    platform_name="WINDOWS"
elif [[ $system_name == Linux* ]]; then
    platform_name="LINUX"
elif [[ $system_name == Darwin* ]]; then
    platform_name="MACOS"
else
    platform_name="UNKNOWN"
fi
export OS_NAME="$platform_name"

# create ram disk
if [ $OS_NAME == "MACOS" ] || [ $OS_NAME == "LINUX" ]; then
    script/ramdisk.sh create GutRamDisk 64
    script/ramdisk.sh create GutReadOnlyDisk 16 ReadOnly
    if [[ $OS_NAME == "MACOS" ]]; then
        export RAMDISK_WRITE=/Volumes/GutRamDisk
        export RAMDISK_READONLY=/Volumes/GutReadOnlyDisk
    elif [[ $OS_NAME == "LINUX" ]]; then
        export RAMDISK_WRITE=/mnt/GutRamDisk
        export RAMDISK_READONLY=/mnt/GutReadOnlyDisk
    fi
fi

# uncompress test resource to temp dir
echo "$MSYS"
export MSYS=winsymlinks:nativestrict
export TESTRSSDIR=${TMPDIR%/}/gut_test_resource
chmod -R 700 "$TESTRSSDIR" && rm -fr "$TESTRSSDIR"
unzip -q -o test_resource.zip -d "$TESTRSSDIR"
printf "Uncompress test resource: %s\n\n" "$TESTRSSDIR"

# set permission for ad hoc files
chmod 000 "$TESTRSSDIR"/yos/copy_file/none_perm.txt
chmod 000 "$TESTRSSDIR"/yos/copy_file/output/none_perm.txt

chmod 000 "$TESTRSSDIR"/yos/copy_dir/source/no-perm-dirs/no_perm_dir
chmod 000 "$TESTRSSDIR"/yos/copy_dir/source/no-perm-files/no_perm_file
chmod 000 "$TESTRSSDIR"/yos/copy_dir/output/exist-no-perm-file/one-file-dir/text.txt
chmod 000 "$TESTRSSDIR"/yos/copy_dir/output/exist-no-perm-dir/misc/deep1

chmod 000 "$TESTRSSDIR"/yos/copy_link/source/no_perm_file
chmod 000 "$TESTRSSDIR"/yos/copy_link/output/no_perm_dir

chmod 000 "$TESTRSSDIR"/yos/same_file/set1/none_perm.txt
chmod 000 "$TESTRSSDIR"/yos/same_file/set2/none_perm.txt

chmod 000 "$TESTRSSDIR"/yos/same_dir/source/no-perm-dirs/no_perm_dir
chmod 000 "$TESTRSSDIR"/yos/same_dir/source/no-perm-files/no_perm_file
