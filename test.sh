#!/bin/bash
set -e

go build tsv.go
./tsv <test2.txt | head
time ./rearrange_tsv <test2.txt >/dev/null
time ./rearrange_tsv <bigtest.txt >/dev/null
