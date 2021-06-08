#!/bin/bash
set -e

go build tsv.go
./tsv <test.txt | head
time ./rearrange_tsv <test.txt >/dev/null
pigz -d -c -p 8 /home/jgbaldwinbrown/Documents/work_stuff/human_transmission_distortion/combo/reannotated_melted_totals.txt.gz | ./rearrange_tsv | head
time ./rearrange_tsv <bigtest.txt >/dev/null
time (pigz -d -c -p 8 /home/jgbaldwinbrown/Documents/work_stuff/human_transmission_distortion/combo/reannotated_melted_totals.txt.gz | ./rearrange_tsv > /dev/null)
