package main

import (
	"fmt"
	"math"
)

var i, j = 12, 17
var formula = "multiplication"

func test_math() {
	fmt.Printf("pi is %g", math.Pi)
}

func addition(x int, y int) int {
	return x + y
}

func sub(i int, j int) int {
	return i - j
}

func mul(i int, j int, formula string) (int, string) {
	return i * j, formula
}
func div(c, d int) int {
	return c / d
}

// try uint.
func mod(e, f uint) uint {
	return e % f
}

// try uint2
func exponentiation(h, g, s, t float64) (float64, float64) {
	pow1 := (math.Pow(h, s))
	pow2 := (math.Pow(g, t))
	return pow1, pow2
}
