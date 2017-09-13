#!/usr/bin/env python

import math
from euler import soe

# generates divisors of a given number
def divisors(n):
	divs = [1]
	for i in xrange(2,int(math.sqrt(n))+1):
		if n%i == 0:
			divs.extend([i,n/i])
	return list(set(divs))

# primes under 10000
primes = soe(10000)

# amicable numbers
amiicables = []

#Numbers that are already checked
checked = []

#Start checking for amicable numbers upto 10000
for i in xrange(2,10000):
	if i not in primes and i not in checked:
		da = sum(divisors(i))
		db = sum(divisors(da))
		checked.extend([da,db])
		if i == db:
			if da != db:
				amiicables.extend([i,da])

# printing amicable numbers
print 'sum is {}'.format(sum(amiicables))
