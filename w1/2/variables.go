package main

import (
	"fmt"
	"math/rand"
)


func main() {
	var x int
	x = rand.Intn(100)
	fmt.Println("Aleatorio: ", x)

	const y int = 100

	z := 134
	w := false
	fmt.Println(z, w)

	a := 3.99999
	b := int(a)
	fmt.Println(a, b)

	c := 10
	var d float32 = float32(c)
	fmt.Println(c, d)
}