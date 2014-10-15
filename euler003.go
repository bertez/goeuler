package main

// The prime factors of 13195 are 5, 7, 13 and 29.

// What is the largest prime factor of the number 600851475143 ?

import (
	"fmt"
	"math"
)

var num int64 = 600851475143

func is_prime(n int64) bool {
	var i int64 = 2
	for ; i < n; i++ {
		if(n % i == 0) {
			return false
		}
	}
	return true
}

func main() {
	var i int64 = int64(math.Sqrt(float64(num))) //we only need to check up to the square root of the number

	for ; i > 1; i-- {
		if(num % i == 0 && is_prime(i)){
			fmt.Printf("The largest prime factor is %d\n", i)
			break
		}
	}

}