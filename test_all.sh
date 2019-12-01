#!/bin/bash

set -euxo pipefail

pushd 01
./test.sh
popd
