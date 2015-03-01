#!/bin/bash

../../cmd/scaffold/scaffold -in="user.go" -out="user_model.go" -struct="User" -scaffold="model"

go test

