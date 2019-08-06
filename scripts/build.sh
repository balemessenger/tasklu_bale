#!/usr/bin/env bash

VERSION=`${PWD}/scripts/version.sh`
TIME=$(date)

if [ -d build/release ]; then
    mkdir -p build/release
    mkdir -p build/debug
fi

if [ "$1" == "release" ]; then
    echo "Building in release mode"
    go build -o build/release/taskulu_$VERSION -a -installsuffix cgo -ldflags="-X 'taskulu/version.BuildTime=$TIME' -X 'taskulu/version.BuildVersion=$VERSION' -s" main.go
else
    echo "Building in debug mode"
    go build -o build/debug/taskulu_$VERSION -a -v -installsuffix cgo -ldflags="-X 'taskulu/version.BuildTime=$TIME' -X 'taskulu/version.BuildVersion=$VERSION' -s" main.go
fi