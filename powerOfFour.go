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
