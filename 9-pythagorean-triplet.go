package main

import (
	"fmt"
)

func findTriplet() (a, b, c int) {
	for a = 1; a < 500; a++ {
		for b = a + 1; b < 1000; b++ {
			c = 1000 - a - b
			if b > c {
				break
			}
			if a*a+b*b == c*c {
				return
			}
		}
	}
	return
}

func main() {
	a, b, c := findTriplet()
	fmt.Println(a, b, c)
	fmt.Println(a * b * c)
}