package main

import (
	"fmt"
)

func main() {
	const dividend = 50
	// variable used for showcasing purposes
	divisor := 7
	fmt.Println(divide(dividend, divisor))
}

func divide(x int, y int) (resultFloor, remainder int) {
	remainder = x % y
	resultFloor = (x - remainder) / y
	return
}
