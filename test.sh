#!/bin/bash
set -e

go build tsv.go
time ./rearrange_tsv <test2.txt >/dev/null
gunzip -c bigtest.txt.gz > bigtest.txt
time ./rearrange_tsv <bigtest.txt >/dev/null
rm bigtest.txt
