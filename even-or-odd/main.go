package main

import "fmt"

func makeRange(length int) []int {
	a := make([]int, length)

	for i := range a {
		a[i] = i + 1
	}

	return a
}

func IsOdd(number int) bool {
	return number%2 != 0
}

func main() {
	numbers := makeRange(10)

	for _, number := range numbers {
		if IsOdd(number) {
			fmt.Printf("%d is odd\n", number)
		} else {
			fmt.Printf("%d is even\n", number)
		}
	}
}
