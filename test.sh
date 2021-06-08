#!/bin/bash
set -e

go build tsv.go
./tsv <test.txt | head
time ./rearrange_tsv <test.txt >/dev/null
time ./rearrange_tsv <bigtest.txt >/dev/null
