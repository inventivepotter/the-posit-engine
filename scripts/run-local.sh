#!/bin/bash
set -x

# build
. $(dirname "$0")/build.sh

# run
go run $CMD