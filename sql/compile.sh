#!/bin/bash

ego
pushd ../cmd/scaffold
go build
popd
pushd test
./test.sh
popd

