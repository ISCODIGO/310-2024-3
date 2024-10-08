package main

import "fmt"

func RemoverSlice(slice []int, pos int) []int {
	return append(slice[:pos], slice[pos+1:]...)
}

func main() {
    // Declarar un slice de enteros
	var pos int

	slice := []int {10, 20, 30, 40, 50}
	fmt.Print("Escriba una posicion desde 0 - ", len(slice)-1, ": ")
	fmt.Scan(&pos)
	fmt.Println("Slice original:", slice)
	fmt.Printf("Valor en la posicion %v es %v\n", pos, slice[pos])
	slice = RemoverSlice(slice, pos)
	fmt.Println(slice)
}