package main

import "fmt"


func sumaFiltrada(slice []int, valorFiltro int, tipoFiltro int) int {
	total := 0

	// recorrer los elementos
	for _, valor := range slice {
		if tipoFiltro < 0 {
			if valor < valorFiltro {
				total += valor
			}
		} else if tipoFiltro == 0 {
			if valor == valorFiltro {
				total += valor
			}
		} else {
			if valor > valorFiltro {
				total += valor
			}
		}
	}

	return total
}


func main() {
	elementos := []int {1, 2, 4, 8, 16, 32, 64}

	fmt.Println(sumaFiltrada(elementos, 10, 24))
}