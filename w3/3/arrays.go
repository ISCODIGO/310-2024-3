package main

import "fmt"

func main() {
	// declaracion de un array de 5 elementos
	var arr1 [5]int

	// asignacion de valores a un array
	arr1[0] = 11
	arr1[1] = 22
	arr1[2] = 33
	arr1[3] = 44
	arr1[4] = 55

	// declaracion de un array de 5 elementos
	var arr2 = [5]int{11, 22, 33, 44, 55}

	// declaracion de un array de 5 elementos
	arr3 := [...]int{11, 22, 33, 44, 55}

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)

	fmt.Println("El primer elemento del array1 es:", arr1[0])
	fmt.Println("El tamaño del array1 es:", len(arr1))
}