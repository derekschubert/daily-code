package main

import (
	"math/rand"
)

/*
	Given a stream of elements too large to store in memory, pick a random element from the stream with uniform probability.
*/

func getRandomElement(xi *[]int) int {
	var el int
	for i, v := range *xi {
		// (1 / (i + 1)) gets us a uniform probability dist
		// make up for open randInt [0, n)
		if rand.Intn(i+1)+1 == 1 {
			el = v
		}
	}
	return el
}
