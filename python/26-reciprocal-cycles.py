#!/usr/bin/env python

def recurring_decimal(divisor):
	if divisor <10:
		divident = 10
	elif divisor >=10 and divisor <100:
		divident = 100
	else:
		divident = 1000
	value_to_check = divident
	string = ''
	for i in xrange(divisor):
		string += str(divident/divisor)
		divident = divident%divisor
		if divident < divisor:
			divident *= 10
			if divident < divisor:
				divident *= 10
				string += '0'
				if divident < divisor:
					divident *= 10
					string += '0'
		if divident == value_to_check:
			return string

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

def gcd(n1,n2):
	remainder = 1
	while remainder != 0:
		remainder = n1%n2
		n1 = n2
		n2 = remainder
	return n1


def lcm(n1,n2):
	return (n1*n2)/gcd(n1,n2)

primes = soe(1000)

d = {n: 0 for n in range(1,1000)}

d[3] = 1

for i in primes[3:]:
	d[i] = len(recurring_decimal(i))

for i in xrange(6,1000):
	if not d[i]:
		if i % 2!= 0 and i%5!= 0:
			for j in primes:
				if i%j == 0:
					number1 = d[j]
					number2 = d[i/j]
					d[i] = lcm(number2,number1)
					break
		else:
			number = i
			while number%2 == 0:
				number = number/2
			while number%5 == 0:
				number = number/5
			d[i] = d[number]

print d.values().index(max(d.values()))+1
