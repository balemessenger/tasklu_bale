
#!/usr/bin/env bash
#
# Generate all etcd protobuf bindings.
# Run from repository root.
#
set -e

if ! [[ "$0" =~ scripts/genproto.sh ]]; then
	echo "must be run from repository root"
	exit 255
fi

if [[ $(protoc --version | cut -f2 -d' ') != "3.7.1" ]]; then
	echo "could not find protoc 3.7.1, is it installed + in PATH?"
	exit 255
fi


# directories containing protos to be built
DIRS="./api/proto"

# disable go mod
export GO111MODULE=off

# exact version of packages to build
GOGO_PROTO_SHA="1adfc126b41513cc696b209667c8656ea7aac67c"
GRPC_GATEWAY_SHA="92583770e3f01b09a0d3e9bdf64321d8bebd48f2"
SCHWAG_SHA="b7d0fc9aadaaae3d61aaadfc12e4a2f945514912"

# set up self-contained GOPATH for building
mkdir -p ${PWD}/build/_proto
export GOPATH=${PWD}/build/_proto/gopath.proto
export GOBIN=${PWD}/build/_proto/bin
export PATH="${GOBIN}:${PATH}"

GOGOPROTO_ROOT="${GOPATH}/src/github.com/gogo/protobuf"
GOGOPROTO_PATH="${GOGOPROTO_ROOT}:${GOGOPROTO_ROOT}/protobuf"
GRPC_GATEWAY_ROOT="${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway"


if  [[ "$1" == "build" ]]; then
    # Ensure we have the right version of protoc-gen-gogo by building it every time.
    # TODO(jonboulle): vendor this instead of `go get`ting it.
    echo "Building protoc-gen-gogo ..."
    go get -u github.com/gogo/protobuf/{proto,protoc-gen-gogo,gogoproto}
    go get -u golang.org/x/tools/cmd/goimports
    pushd "${GOGOPROTO_ROOT}"
        git reset --hard "${GOGO_PROTO_SHA}"
        make install
    popd

    # generate gateway code
    echo "Building gateway code ..."
    go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
    go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
    pushd "${GRPC_GATEWAY_ROOT}"
        git reset --hard "${GRPC_GATEWAY_SHA}"
        go install ./protoc-gen-grpc-gateway
    popd
fi


echo "Generating protos ..."
for dir in ${DIRS}; do
	pushd "${dir}"

    protoc -I=".:${GRPC_GATEWAY_ROOT}/third_party/googleapis"  --gogofast_out=plugins=grpc,\
Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types:./src \
./*.proto

#		protoc --gofast_out=plugins=grpc:./src -I=".:${GRPC_GATEWAY_ROOT}/third_party/googleapis" ./*.proto
				# shellcheck disable=SC1117
#		protoc --go_out=plugins=grpc:./src ./*.proto
		goimports -w ./src/*.pb.go
	popd
done