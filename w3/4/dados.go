package main

import (
	"fmt"
	"math/rand"
)

func lanzamiento() uint8 {
	cara := rand.Intn(6) + 1
	return uint8(cara)
}

func main() {
	var arr [7]int

	// 100 lanzamientos de dado
	for i := 1; i <= 100; i++ {
		valor := lanzamiento()  // random del [1, 6]
		arr[valor] += 1
	}

	max := 0
	max_pos := -1
	for pos, item := range arr {
		if item > max {
			max = item
			max_pos = pos
		}
	}

	fmt.Println(arr)
	fmt.Println("El mas frecuente es", max_pos)
}