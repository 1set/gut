#!/bin/bash

set -e

if [[ ! -z "$TESTRSSDIR" ]] ; then
    chmod -R 700 "$TESTRSSDIR" && rm -fr "$TESTRSSDIR"
fi
