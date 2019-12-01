#!/bin/bash

set -euxo pipefail

FILE=${FILE:-input.txt}

awk '{ sum += int($0/3)-2 } END { print sum }' $FILE

