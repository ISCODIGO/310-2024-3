package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hola mundo")
	fmt.Println("2 + 3 =", sumar(2, 3))
}

func sumar(a int, b int) int {
	return a + b
}