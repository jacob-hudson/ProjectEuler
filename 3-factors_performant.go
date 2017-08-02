// this verions only uses the square root of the number
// the largest prime factor has to be smaller than the square root

package main

import "fmt"

// test data below
// prime factors of 13195
// answers - 5,7,13,29 (largest)

func main() {
  // variables
  var target = 600851475143
  var largest = 0
  var factors [2]int

// i * i == sqrt(target), just exposing new viewpoints
  for i := 2; i * i < target; i++ {
    // determining if it is a divisor with no remainder
    if target % i == 0 {
      // stories iteration and it's factor
      factors[0] = i
      factors[1] = target / i
      var isPrime = true
      for k := 0; k < 2; k++ {
        // checking for factors of a factor - determines prime numbers
        for j := 2; j * j < factors[k]; j++ {
          if factors[k] % j == 0 {
            isPrime = false
            break;
          }
        }
        if isPrime && factors[k] > largest {
          largest = factors[k]
        }
      }
    }
  }
  fmt.Println(largest)
}
