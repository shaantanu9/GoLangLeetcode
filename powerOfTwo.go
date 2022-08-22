// Given an integer n, return true if it is a power of four. Otherwise, return false.

// An integer n is a power of two, if there exists an integer x such that n == 2x.

package main

import "fmt"

func main() {
	printResult := powerOfTwo(64)
	fmt.Println(printResult)
}

func powerOfTwo(n int) bool {
	for n > 0 && n%2 == 0 {
		n /= 2
	}
	return n == 1
}
