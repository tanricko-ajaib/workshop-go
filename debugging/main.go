package main

import (
	"fmt"

	. "github.com/tanricko-ajaib/workshop/debugging/fib"
)

func main() {
	var (
		input  = []int{0, 1, 2, 4, 8, 16}
		output = []int{0, 1, 1, 3, 21, 987}
	)

	for i, n := range input {
		x := Fib(n)

		if x == output[i] {
			fmt.Printf("Fib(%d) = %d\n", n, x)
		} else {
			fmt.Printf("Fib(%d) = %d; want %d\n", n, x, output[i])
		}
	}
}
