package main

import "fmt"

func main() {

	// variables:  a, b = 1, c, sum = 0
	a, b, c, sum := 1, 1, 0, 0
	
	// infinite while loop
	for {
		// if c is larger than 4e6, stop
		if c > 4000000 {
			break
		}
		
		// if number is even, add it to the sum
		if c % 2 == 0 {
			sum += c
		}
		
		// adds two previous values to c, eg 2 = 1 + 1, 3 = 1 + 2...
		c = a + b
		
		// assigns c to a, a to b, enabling the loop to progress
		b = a
		a = c
	}
	fmt.Println(sum)
}
