package main
/*

Un programa que solicite N numeros. Luego de capturarlos
hacer los siguientes calculos:
- suma
- promedio
- max
- min

*/

import "fmt"

func generarArray(n int) []int {
	// crear el arreglo
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	return arr
}

func main() {
	fmt.Print("Escriba un numero:")
	var numero int
	fmt.Scan(&numero)
	datos := generarArray(numero)
	fmt.Println(datos)
}