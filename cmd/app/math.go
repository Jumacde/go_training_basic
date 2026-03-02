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
