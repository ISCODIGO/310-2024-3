/*
Dado un slice de enteros que posiblemente contiene numeros duplicados
crear una funcion que permita eliminar las duplicaciones.

[10 20 10 30 40 40 50] -> [10 20 30 40 50]
repetidos
1
*/

package main

import (
	"fmt"
)

func remover_duplicado(slice []int) []int{
	// crear un slice para almacenar posiciones
	posiciones := make([]int, len(slice))
	// recorrer cada elemento del slice
	limite := len(slice)
	// voy a recorrer hasta el penultimo
	for i := 0; i < limite - 1; i++ {
		// Evitar recorrer un valor repetido
		if posiciones[i] == 1 {
			continue
		}
		// recorrer el resto del slice
		for j := i + 1; j < limite; j++ {
			if slice[i] == slice[j] {
				// editar la posicion con un 1
				posiciones[j] = 1
			}
		}  // fin for-interno
	}  // fin for-externo

	var unicos []int  // slice donde estaran los valores unicos
	for pos, val := range posiciones {
		if val == 0 {
			unicos = append(unicos, slice[pos])
		}
	}
	return unicos
}

func main() {
	entrada := []int {10, 20, 10, 30, 10, 30}
	salida := remover_duplicado(entrada)
	fmt.Println("Entrada:", entrada)
	fmt.Println("Salida:", salida)
}