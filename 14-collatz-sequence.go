package main

import (
	"fmt"
)

func findTerms(n int64) int64 {
	var c int64 = 1

	for {
    // even
		if n%2 == 0 {
			n = n / 2
      // odd
		} else {
			n = (3 * n) + 1
		}
		c++
		if n == 1 {
			return c
		}
	}
}

func main() {
	type answer struct {
		number    int64
		termCount int64
	}
	a := answer{0, 0}

  // all numbers under 1 million
	for i := int64(999999); i > 0; i-- {
		if termCount := findTerms(i); termCount > a.termCount {
			a.termCount = termCount
			a.number = i
		}
	}
	fmt.Println(a.number)
}
