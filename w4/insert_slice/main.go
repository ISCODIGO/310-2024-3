package main

import "fmt"

func insertar(slice []int, index int, value int) []int {
	slice = append(slice[:index+1], slice[index:]...)
	slice[index] = value
	return slice
  }
func main() {
	// Declarar un slice de enteros
	var pos int
	var nuevo_valor int

	slice := []int{10, 20, 30, 40, 50}
	fmt.Print("Escriba una posicion desde 0 - ", len(slice)-1, ": ")
	fmt.Scan(&pos)
	fmt.Print("Escriba un valor a insertar: ")
	fmt.Scan(&nuevo_valor)
	fmt.Println("Slice original:", slice)
	slice = insertar(slice, pos, nuevo_valor)
	fmt.Println(slice)
}
