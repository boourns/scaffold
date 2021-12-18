#!/bin/bash

set -e

export PATH=$PATH:`pwd`/cmd/scaffold

for i in model ast; do
  if [ -e $i/script/test.sh ]; then
    pushd $i
    ./script/test.sh
    popd
  fi
done
