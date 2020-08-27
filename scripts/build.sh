#!/bin/bash
set -x

# changing project root directory
pushd $(dirname "$0")/.. >/dev/null
trap "popd >/dev/null" EXIT

# define vars
CMD="$PWD/cmd/the-posit-engine/"
BIN=${CMD/cmd/bin}

# build
set -e
go build -o "$BIN" "$CMD"