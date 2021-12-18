#!/bin/bash

pushd model
./script/build.sh
popd

go build
