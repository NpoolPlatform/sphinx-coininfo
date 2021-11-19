#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

VERSION=v1.43.0
URL_BASE=https://raw.githubusercontent.com/golangci/golangci-lint
URL=$URL_BASE/$VERSION/install.sh

if [[ ! -f .golangci.yml ]]; then
    echo 'ERROR: missing .golangci.yml in repo root' >&2
    exit 1
fi

if ! command -v gofumpt; then
    go install mvdan.cc/gofumpt@latest
fi

#if ! command -v golangci-lint; then
#    curl -sfL $URL | sh -s $VERSION
#    PATH=$PATH:bin
#fi

if [ ! -f bin/golangci-lint ]; then
    curl -sfL $URL | sh -s $VERSION                                            
fi

PATH=$PATH:bin

golangci-lint version
golangci-lint linters
golangci-lint run "$@"
