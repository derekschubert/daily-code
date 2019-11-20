package main

import (
	"fmt"
	"math"
	"math/rand"
)

/*
  This problem was asked by Google.

  The area of a circle is defined as πr^2.
  Estimate π to 3 decimal places using a Monte Carlo method.

  Hint: The basic equation of a circle is x2 + y2 = r2.
*/

func main() {
	fmt.Println(findPi())
}

type coords struct {
	x float64
	y float64
}

func genRandom() coords {
	return coords{
		x: (rand.Float64() * 2.0) - 1.0,
		y: (rand.Float64() * 2.0) - 1.0,
	}
}

func (c *coords) inCircle() bool {
	a := (c.x * c.x) + (c.y * c.y)
	if a <= 1 {
		return true
	}
	return false
}

func findPi() float64 {
	iterations := 10000000.0
	inside := 0.0

	for i := 0.0; i < iterations; i++ {
		c := genRandom()
		if c.inCircle() {
			inside++
		}
	}

	pi := (inside / iterations) * 4
	pi = math.Floor(pi*1000) / 1000
	return pi
}
