#!/bin/bash

pushd ..
go build
popd
pushd test
../../scaffold model -in=user.go
popd

