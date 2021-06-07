#!/bin/bash
set -e

go build tsv.go
./tsv <test.txt | head
time ./rearrange_tsv <test.txt >/dev/null
