#!/usr/bin/env python

last_number = 1001*1001
odd_numbers = xrange(1,last_number+1,2)
i = 0
gap = 1
solution = 1

while odd_numbers[i] != last_number:
	for j in xrange(4):
		i+= gap
		solution += odd_numbers[i]
	gap += 1

print solution
