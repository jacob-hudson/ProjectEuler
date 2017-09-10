#!/usr/bin/env python

grand_total = 0

def fifth_power(string):
    total = 0
    length = len(string)
    for j in range(0,length):
        num = (int(string[j])**5)
        total = total + num
    return total

for i in range(10,355000):
    string = str(i)
    x = fifth_power(string)
    if x == i:
        grand_total += x


print grand_total
