package main
// soe - Sieve Of Eratosthenes
import (
	"fmt"
)

func soe(n int) []int {
	// Create a boolean array "prime[0..n]"
	// set all entries as true. A value in prime[i] will
	// finally be false if i is *not* prime, else true.
	integers := make([]bool, n+1)
	for i := 2; i < n+1; i++ {
		integers[i] = true
	}

	for p := 2; p*p <= n; p++ {
		// If integers[p] is not changed, then it is a prime
		if integers[p] == true {
			// Update all multiples of p
			for i := p * 2; i <= n; i += p {
				integers[i] = false
			}
		}
	}

	// return all prime numbers <= n
	var primes []int
	for p := 2; p <= n; p++ {
		if integers[p] == true {
			primes = append(primes, p)
		}
	}

	return primes
}

func main() {
	primes := soe(110000)
	fmt.Println(len(primes))
	if len(primes) > 10001 {
		fmt.Println(primes[10000])
	}
}
