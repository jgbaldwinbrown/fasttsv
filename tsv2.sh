#!/bin/bash

cat "$@" | \
mawk -F "\t" -v OFS="\t" '{print $4, $5, $2}'
