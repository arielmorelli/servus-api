#!/bin/bash

ARCHS=(amd64 arm64)

# install dependencies
go mod download

# clean up
rm -rf bin
mkdir bin

# Build for all into to the bin folder
for target_arch in ${ARCHS[@]}
do
	env GOOS=linux GOARCH=${target_arch} go build -ldflags "-w" -o bin/servus-api-${target_arch}
done
