package main

// Each new term in the Fibonacci sequence is generated by adding the previous
// two terms. By starting with 1 and 2, the first 10 terms will be:

// 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...

// By considering the terms in the Fibonacci sequence whose values do not
// exceed four million, find the sum of the even-valued terms.

import (
	"fmt"
)

var max int = 4000000
var sum int

func main() {
	a, b := 0,1
	for b < max {
		a, b = b, a + b
		if(b % 2 == 0) {
			sum += b
		}
	}

	fmt.Printf("The sum is: %d\n", sum)
}