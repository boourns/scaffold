#!/bin/bash

cd test

rm user_sql.go
../../scaffold model -in=user.go

go test
