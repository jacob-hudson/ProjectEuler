#!/usr/bin/env python

def factorial(n):
    if n == 0:
        return 1
    else:
        return n * factorial(n-1)

def main():
    digits = range(10)
    result = []
    n = 999999

    for i in xrange(len(digits) - 1, -1, -1):
        index, n = divmod(n, factorial(i))
        result.append(digits.pop(index))

    print ''.join(str(i) for i in result)

if __name__ == "__main__":
    main()
