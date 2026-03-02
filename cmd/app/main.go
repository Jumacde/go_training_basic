package main

import "fmt"

func main() {
	//show_text()       // call and extecute show_text method from the hello_world.goo
	//show_rand_digit() // show a random digit
	//test_printfn()
	//test_math() // show pi
	x, y := 7, 17
	result := addition(x, y)
	fmt.Printf("%d + %d = %d\n", x, y, result)

	a, b := hello_world("hello ", "world")
	fmt.Print(a, b)

	// values is dedeclared on the other file
	fmt.Printf("\n%d - %d = %d", i, j, sub(i, j))

	// try combi. diffrent value types on a method and call it.
	result, str := mul(i, j, formula)
	fmt.Printf("\n%s: %d * %d = %d", str, i, j, result)

	c, d := 12, 4
	fmt.Printf("\n%d / %d = %d", c, d, div(c, d))

	// try uint
	var e uint = 17
	var f uint = 4
	fmt.Printf("\n%d mod %d = %d", e, f, mod(e, f))

	// try for loop
	var test uint = 0
	fmt.Printf("\nfor loop. i is added until 10. %d", test_for(uint(test)))
}
