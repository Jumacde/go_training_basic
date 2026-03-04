package main

func test_switch() int {
	var a1 = 12
	var a2 = 17
	var a3 = 9
	switch {
	case a1 < a2:
		return a1
	case a1 > a2:
		return a2
	default:
		return a3
	}
}
