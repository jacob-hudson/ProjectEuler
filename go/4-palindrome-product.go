package main


import (
	"fmt" // formatting
	"strconv" // string conversion
)

func isPalindrome(n int64) bool {
	if n < 0 {
		n = -n
	}

	s := strconv.FormatInt(n, 10)
	bound := (len(s) / 2) + 1

	for i := 0; i < bound; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	max := 0
	maxI := 0
	maxJ := 0
	for i := 999; i >= 100; i-- {
		for j := 999; j >= 100; j-- {
			product := i * j
			if max > product {
				break
			}
			if isPalindrome(int64(product)) {
				max = product
				maxI = i
				maxJ = j
				continue
			}
		}
	}
  // output answers
	fmt.Println("Largest palindrome: ", max)
	fmt.Println("Value of i: ", maxI)
	fmt.Println("Value of j: ", maxJ)
}
