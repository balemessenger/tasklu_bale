#!/usr/bin/env bash

VERSION=`${PWD}/scripts/version.sh`
TIME=$(date)

echo "version: ${VERSION}"
docker build --build-arg docker_version=$VERSION -t docker.bale.ai/molana/taskulu:$VERSION -f deploy/Dockerfile .
docker push docker.bale.ai/molana/taskulu:$VERSION