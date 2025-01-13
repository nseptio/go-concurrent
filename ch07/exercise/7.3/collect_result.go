/*
In listing 7.13, we use a child goroutine to calculate the factors of one number
and the main() goroutine to work out the factors of the other. Modify this listing
so that, using multiple goroutines, we collect the factors of 10 random numbers.
*/

package main

import (
	"fmt"
	"math/rand/v2"
)

func findFactors(number int) []int {
	result := make([]int, 0)
	for i := 1; i <= number; i++ {
		if number%i == 0 {
			result = append(result, i)
		}
	}
	return result
}

func main() {
	resultChs := make([]chan []int, 10)

	for i := 0; i < 10; i++ {
		resultChs[i] = make(chan []int)
		go func(n int) {
			resultChs[i] <- findFactors(rand.IntN(1_000_000))
		}(i)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(<-resultChs[i])
	}
}
