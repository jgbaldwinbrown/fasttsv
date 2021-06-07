#!/usr/bin/env python3

import sys

def rearrange_cols(inconn, outconn, cols):
    for l in inconn:
        sl = l.rstrip('\n').split('\t')
        outl = [sl[x] for x in cols]
        print("\t".join(outl))

def main():
    rearrange_cols(sys.stdin, sys.stdout, [3,4,1])

if __name__ == "__main__":
    main()
