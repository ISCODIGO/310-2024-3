package main

import (
	"fmt"
)

func intercambio(a int, b int) (int, int, bool) {
	return b, a, a == b
}

func main() {
	// var x int
	// var y int
	// var z bool


	x, y, z := intercambio(5, 5)
	fmt.Println(x, y, z)
}