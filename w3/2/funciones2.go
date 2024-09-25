package main

import "fmt"

/*
Este ejemplo es una funcion multivariadica, es decir, una funcion que recibe un
numero variable de argumentos.
*/
func sumar(valores ...int) int {
	var total int
	for _, valor := range valores {
		total += valor
	}
	return total
}

func main() {
	fmt.Println(sumar(1, 2), sumar(1, 2, 3, 4), sumar(1, 2, 3, 4, 5, 6))
}