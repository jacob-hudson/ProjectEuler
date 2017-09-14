#!/usr/bin/env python

def is_pandigital(n, s=9):
    n = str(n)
    return len(n) == s and not '1234567890'[:s].strip(n)


p = set()
for i in range(2,  60):
    start = 1234 if i < 10 else 123
    for j in range(start, 10000 // i): # // - floor division
        if is_pandigital(str(i) + str(j) + str(i * j)):
            p.add(i * j)

print sum(p)
