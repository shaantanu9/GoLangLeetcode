// Given an integer n, return true if it is a power of four. Otherwise, return false.
// An integer n is a power of four, if there exists an integer x such that n == 4x.
package main

import "fmt"

func main() {
	printResult := powerOfFour(64)
	fmt.Println(printResult)
}

func powerOfFour(n int) bool {
	for n > 0 && n%4 == 0 {
		n /= 4
	}
	return n == 1
}
