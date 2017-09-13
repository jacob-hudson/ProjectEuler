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
