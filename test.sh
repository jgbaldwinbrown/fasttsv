#!/bin/bash
set -e

go build tsv.go
time ./tsv <test.txt >/dev/null
