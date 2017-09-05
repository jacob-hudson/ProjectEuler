#!/usr/bin/env python

from math import factorial

def main():
    n = 100 # the problem requires 100!
    s = 0
    result = factorial(n)
    # converting result to string, in order to iterate through the digits
    string_result = str(result)
    for number in string_result:
        # converting number to int, in order to calculate the sum
        s += int(number)
    print s

# defines main function in python
if __name__ == '__main__':
    main()
