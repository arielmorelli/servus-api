#!/bin/bash

ARCHS=(amd64 arm64 ppc64le ppc64 s390x)


# install dependencies
go mod download

# clean up
rm -rf bin
mkdir bin

# Build for all into to the bin folder
for target_arch in ${ARCHS[@]}
do
	env GOOS=linux GOARCH=${target_arch} go build -o bin/servus-api-${target_arch}
done

