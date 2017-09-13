#!/usr/bin/env python

fib_list = [0,1]

i = 2

while 1:

	fib_new = fib_list[i-1] + fib_list[i-2]
	fib_list.append(fib_new)
	if fib_new>10**999:
		print i
		break
	i += 1
