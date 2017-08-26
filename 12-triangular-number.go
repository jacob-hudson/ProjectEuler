package main

import (
    "math"
)

// uint64 - 64 bit unsigned integer

func main() {
    // variables
    divisors := 1
    triangle := uint64(1)

    // number with less than 500 divisors
    for i := 2; divisors < 500; i++ {
        // triangle is the sum of previous triangels
        triangle += uint64(i)
        divisors = num_divisors(triangle)
    }
    println(triangle)
}

func num_divisors(number uint64) int {
    counter := 0
    limit := uint64(math.Sqrt(float64(number)))
    for i := 1; uint64(i) <= limit; i++ {
        if number%uint64(i) == 0 {
            counter += 2
        }
    }
    return counter
}
