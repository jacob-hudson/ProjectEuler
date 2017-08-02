// this algoritham is slow, I am still learnning go
// I will post a better algoritham later on

package main

import (
      "fmt"
)
// test data below
// prime factors of 13195
// answers - 5,7,13,29 (largest)

func main() {
  // variables
  var target = 600851475143
  var largest = 0

  for i := 2; i < target; i++ {
    // determining if it is a divisor with no remainder
    if target % i == 0 {
      var isPrime = true
      // checking for factors of a factor - determines prime numbers
      for j := 2; j < i; j++ {
        if i % j == 0 {
          isPrime = false
          break;
        }
      }
      if isPrime == true {
        largest = i
      }
    }
  }
  fmt.Println(largest)
}
