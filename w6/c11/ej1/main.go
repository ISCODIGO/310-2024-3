package main

import "fmt"

func revertir1(slice []int) []int {
	// crear un slice del mismo tamano que el original
	size := len(slice)
	copia := make([]int, size)  // esta rellena de ceros

	// recorrer el slice original
	k := size - 1
	for i := 0; i < size; i++ {
		copia[k] = slice[i]
		k -= 1
	}

	return copia
} 

func revertir2(slice []int) []int {
	// usar el mismo slice e ir intercambiando
	// valores

	size := len(slice)
	pos_reversa := size - 1
	mitad := mint(size / 2)
	for i := 0; i < mitad; i++ {
		temp := slice[pos_reversa]
		slice[pos_reversa] = slice[i]
		slice[i] = temp
		pos_reversa -= 1
	}
	return slice
}

func main() {
	prueba := []int{5, 6, 7, 1, 3}
	fmt.Println(revertir2(prueba))

}
