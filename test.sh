#!/bin/bash
set -e

go build tsv.go
./tsv <test.txt | head
time ./tsv <test.txt >/dev/null
