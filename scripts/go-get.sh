#!/bin/bash
set -x

GO_VERSION=1.15

version=$(go version | awk '{ print $3 }')
requirement="go${GO_VERSION/./\.}"
if [[ ! $version =~ $requirement ]]; then
    echo "Go $GO_VERSION required but found $version" >&2
    exit 1
fi

# changing project root directory
pushd $(dirname "$0")/.. >/dev/null
trap "popd >/dev/null" EXIT

go get -u -v "$@"
go mod vendor