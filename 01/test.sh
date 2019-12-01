#!/bin/bash

set -euxo pipefail

actual=$(FILE="input-test.txt" $(pwd)/01.sh)
expected=$((2+2+654+33583))

if [[ "$expected" != "$actual" ]]; then
    echo "expected=$expected actual=$actual"
    exit 1
fi  

