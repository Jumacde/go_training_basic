package main

import (
	"fmt"
	"math"
)

func test_printfn() {
	var x = 7
	fmt.Printf("show the number %g by math.Sqrt(%g)", math.Sqrt(float64(x)), x)
}
