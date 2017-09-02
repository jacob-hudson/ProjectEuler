package main

import "fmt"

func main() {

	// variables
	sum := 0

	// counting multiples under 10000
	for i := 0; i < 1000; i++ {

		// using modulus to find multiples of 3 or 5
		if i%3 == 0 || i%5 == 0 {
		
			// adding the value to sum
			sum += i
		}
	}
	
	// results
	fmt.Printf("Answer: %d\n", sum)
}
