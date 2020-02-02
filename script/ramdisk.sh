#!/bin/bash

# set -exo pipefail
set -eo pipefail

# Preconditions

function display_help() {
    cat <<END
Usage:
  ./ramdisk.sh create name [size] [access]
  ./ramdisk.sh destroy name
  ./ramdisk.sh reload name

Options:
  name: disk name.
  size: disk size in MiB, default is 128 MiB.
  access: access type, use ReadWrite or ReadOnly, default is ReadWrite.
END
}

system_name=$(uname -s)
case "$system_name" in
Darwin*)
    OS_NAME="MACOS"
    ;;
Linux*)
    OS_NAME="LINUX"
    ;;
*)
    printf "running on unsupported os: $system_name\n"
    exit 1
    ;;
esac

if [[ $# -lt 1 ]]; then
    display_help
    exit 2
fi

case "$1" in
create)
    ACTION="CREATE"
    ;;
destroy)
    ACTION="DESTROY"
    ;;
reload)
    ACTION="RELOAD"
    ;;
*)
    printf "got unknown verb: '$1'\n"
    exit 3
    ;;
esac

DISK_NAME="${2// /}"
[[ -z "$DISK_NAME" ]] && printf "got blank name: '$2'\n" && exit 4

DISK_SIZE_MB=128
if [[ ! -z "$3" ]]; then
    if [[ $3 =~ ^[0-9]+$ ]]; then
        DISK_SIZE_MB=$3
    else
        printf "got invalid size number: '$3'\n" && exit 5
    fi
fi

case "$4" in
ReadWrite)
    ACCESS_TYPE=ReadWrite
    ;;
ReadOnly)
    ACCESS_TYPE=ReadOnly
    ;;
"")
    ACCESS_TYPE=ReadWrite
    ;;
*)
    printf "got invalid access type: $4\n"
    exit 6
    ;;
esac

# macOS operations

function macos_exist_ramdisk() {
    DISK_ID=$(diskutil list | grep -e "APFS Volume" | grep -e "Volume\s*${DISK_NAME}\s*\d" | awk '{ print $7 }')
    if [[ -z "$DISK_ID" ]]; then
        # 1 as failure for non-existance
        return 1
    else
        # 0 as ok for existance
        return 0
    fi
}

function macos_create_ramdisk() {
    if macos_exist_ramdisk; then
        printf "ramdisk '${DISK_NAME}' already exists: ${DISK_ID}\n"
        exit 10
    else
        local DISK_SECTORS=$((2048 * DISK_SIZE_MB))
        DISK_ID=$(hdiutil attach -nomount ram://$DISK_SECTORS)
        diskutil partitionDisk ${DISK_ID} 1 GPTFormat APFS "$DISK_NAME" '100%'
        if ! macos_exist_ramdisk; then
            printf "failed to create ramdisk '${DISK_NAME}'\n"
            exit 11
        fi

        if [[ $ACCESS_TYPE == "ReadWrite" ]]; then
            printf "ramdisk '${DISK_NAME}' just created: ${DISK_ID}\n"
        elif [[ $ACCESS_TYPE == "ReadOnly" ]]; then
            diskutil umount ${DISK_ID}
            diskutil mount readOnly ${DISK_ID}
            printf "read-only ramdisk '${DISK_NAME}' just created: ${DISK_ID}\n"
        fi
    fi
}

function macos_reload_ramdisk() {
    if macos_exist_ramdisk; then
        diskutil umount ${DISK_ID}
        diskutil mount readOnly ${DISK_ID}
        if ! macos_exist_ramdisk; then
            printf "failed to reload ramdisk '${DISK_NAME}'\n"
            exit 21
        fi

        printf "reloaded ramdisk '${DISK_NAME}' for read-only: ${DISK_ID}\n"
    else
        printf "ramdisk '${DISK_NAME}' doesn't exist\n"
        exit 20
    fi
}

function macos_destroy_ramdisk() {
    if macos_exist_ramdisk; then
        diskutil umount ${DISK_ID}
        hdiutil detach ${DISK_ID}
        if macos_exist_ramdisk; then
            printf "failed to detach ramdisk '${DISK_NAME}'\n"
            exit 31
        fi

        printf "ramdisk '${DISK_NAME}' just detached\n"
    else
        printf "ramdisk '${DISK_NAME}' doesn't exist\n"
        exit 30
    fi
}

# Linux operations

function linux_exist_ramdisk() {
    DISK_ID=$(df -t tmpfs | grep -e "/mnt/${DISK_NAME}$" | awk '{ print $6 }')
    if [[ -z "$DISK_ID" ]]; then
        # 1 as failure for non-existance
        return 1
    else
        # 0 as ok for existance
        return 0
    fi
}

function linux_create_ramdisk() {
    if linux_exist_ramdisk; then
        printf "ramdisk '${DISK_NAME}' already exists: ${DISK_ID}\n"
        exit 10
    else
        DISK_ID=/mnt/${DISK_NAME}
        if [[ ! -d "$DISK_ID" ]]; then
            mkdir -p "$DISK_ID"
        fi

        if [[ $ACCESS_TYPE == "ReadWrite" ]]; then
            mount -t tmpfs -o size=${DISK_SIZE_MB}m tmpfs ${DISK_ID}
            printf "ramdisk '${DISK_NAME}' just created: ${DISK_ID}\n"
        elif [[ $ACCESS_TYPE == "ReadOnly" ]]; then
            mount -t tmpfs -o ro,size=${DISK_SIZE_MB}m tmpfs ${DISK_ID}
            printf "read-only ramdisk '${DISK_NAME}' just created: ${DISK_ID}\n"
        fi
    fi
}

function linux_reload_ramdisk() {
    if linux_exist_ramdisk; then
        mount -o remount,ro ${DISK_ID} ${DISK_ID}
        if ! linux_exist_ramdisk; then
            printf "failed to reload ramdisk '${DISK_NAME}'\n"
            exit 21
        fi

        printf "reloaded ramdisk '${DISK_NAME}' for read-only: ${DISK_ID}\n"
    else
        printf "ramdisk '${DISK_NAME}' doesn't exist\n"
        exit 20
    fi
}

function linux_destroy_ramdisk() {
    if linux_exist_ramdisk; then
        umount ${DISK_ID}
        rmdir ${DISK_ID}
        if linux_exist_ramdisk; then
            printf "failed to detach ramdisk '${DISK_NAME}'\n"
            exit 31
        fi

        printf "ramdisk '${DISK_NAME}' just detached\n"
    else
        printf "ramdisk '${DISK_NAME}' doesn't exist\n"
        exit 30
    fi
}

# Main logic
printf "Task: [${OS_NAME}] ${ACTION} "
if [[ $ACTION == "CREATE" ]]; then
    printf "${ACCESS_TYPE} ramdisk '${DISK_NAME}' of ${DISK_SIZE_MB} MiB\n\n"
elif [[ $ACTION == "DESTROY" ]] || [[ $ACTION == "RELOAD" ]]; then
    printf "ramdisk '${DISK_NAME}'\n\n"
fi

if [[ $OS_NAME == "MACOS" ]]; then
    case "$ACTION" in
    CREATE)
        macos_create_ramdisk
        ;;
    DESTROY)
        macos_destroy_ramdisk
        ;;
    RELOAD)
        macos_reload_ramdisk
        ;;
    esac
elif [[ $OS_NAME == "LINUX" ]]; then
    case "$ACTION" in
    CREATE)
        linux_create_ramdisk
        ;;
    DESTROY)
        linux_destroy_ramdisk
        ;;
    RELOAD)
        linux_reload_ramdisk
        ;;
    esac
fi

exit 0
