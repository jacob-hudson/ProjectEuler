#!/usr/bin/env python

def factorial(n):
    if n == 0:
        return 1
    else:
        return n * factorial(n-1)

def main():
    digits = range(10)
    res = []
    n = 999999 # PE uses 1-based indexing

    for i in xrange(len(digits) - 1, -1, -1):
        index, n = divmod(n, factorial(i))
        res.append(digits.pop(index))

    print ''.join(str(i) for i in res)

if __name__ == "__main__":
    main()
