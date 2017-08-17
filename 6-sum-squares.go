package main

import "fmt"

func main() {
	s1, s2 := 0, 0
// all natural numbers up until 100
	for i := 1; i <= 100; i++ {
    // sum of squares
		s1 += i * i
    // square of sum
		s2 += i
	}
  // square of sum
	s2 = s2 * s2
	diff := s2 - s1
  // output results
	// fmt.Println(s1)
	// fmt.Println(s2)

  // final difference
	fmt.Println(diff)
}
