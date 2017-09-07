#!/usr/bin/env python

from math import sqrt

def divisors(n):
	divs = [1]
	for i in xrange(2,int(sqrt(n))+1):
		if n%i == 0:
			divs.extend([i,n/i])
	return list(set(divs))

abundant_numbers = []

for i in xrange(12,28123):
	if sum(divisors(i))>i:
		abundant_numbers.append(i)

non_ab_sum = [x for x in xrange(28123)]

for i in xrange(len(abundant_numbers)):
	for j in xrange(i,28123):
		if abundant_numbers[i] + abundant_numbers[j] < 28123:
			non_ab_sum[abundant_numbers[i] + abundant_numbers[j]] = 0
		else:
			break


print sum(non_ab_sum)
