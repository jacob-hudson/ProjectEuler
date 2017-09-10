#!/usr/bin/env python

L = set()
limit = 101
for a in range(2,limit):
    for b in range(2,limit):
        c = a**b
        if c in L:
            pass
        L.add(c)


print len(L)
