#!/usr/bin/env python

def soe(n):
	is_prime = [True]*n
	is_prime[0] = False
	is_prime[1] = False
	for i in xrange(2,int(n**0.5+1)):
		index = i*2
		while index < n:
			is_prime[index] = False
			index = index+i
	prime = []
	for i in xrange(n):
		if is_prime[i] == True:
			prime.append(i)
	return prime

def is_prime(n):
	for i in xrange(2,int(abs(n)**0.5)+1):
		if n%i == 0:
			return False
	return True

primes1000 = soe(1000)
primes = primes1000[:]
largest = 0
for b in primes1000:
	for a in primes1000:
		i = 0
		while 1:
			quadratic = i**2+a*i+b
			if quadratic not in primes:
				if is_prime(quadratic):
					primes.append(quadratic)
				else:
					if i-1 > largest:
						largest = i-1
						ab = a*b
					break
			i += 1
		i = 0

		while 1:
			quadratic = i**2-a*i+b
			if quadratic not in primes:
				if is_prime(quadratic) and quadratic>0:
					primes.append(quadratic)
				else:
					if i-1 > largest:
						largest = i-1
						ab = -1*a*b
					break
			i += 1

print ab
