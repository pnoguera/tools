#!/usr/bin/env bash

grep -P "^\t+_" tools.go | awk '{print $2}'|tr '"' ' ' | while read i; do go get -v -u $i@latest; done
