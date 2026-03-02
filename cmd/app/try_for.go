package main

func test_for(result uint) uint {
	for i := 0; i < 10; i++ {
		result += uint(i)
	}
	return result
}
