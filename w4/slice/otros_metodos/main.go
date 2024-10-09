package main

import "fmt"

func main() {
	slice1 := []int{10, 20, 30}

	slice1 = append(slice1, 40)
	slice1 = append(slice1, 50)
	slice1 = append(slice1, 50)
	slice1 = append(slice1, 50)

	fmt.Println("El tamano es", len(slice1), "y la capacidad es", cap(slice1))
	slice2 := make([]string, 11)
	fmt.Println(slice2)
}