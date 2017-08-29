package main

import "fmt"

func main() {
	var n = [1000]int{}
	n[0] = 1
	for i := 0; i < 1000; i++ {
		for j := 0; j < len(n); j++ {
			if n[j] > 0 {
        // bitshifts in Go - https://medium.com/learning-the-g
        // tl;dr << is n[j] * 2, 1 time
				n[j] = n[j] << 1
			}
			if j > 0 && n[j - 1] >= 10 {
				n[j] += int(n[j - 1] / 10)
				n[j - 1] = n[j - 1] % 10
			}
		}
	}
	sum := 0
	for i := 0; i < len(n); i++ {
		sum += n[i]
	}
	fmt.Println(sum)
}
