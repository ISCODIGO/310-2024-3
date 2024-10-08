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

func generarResultados(n int) (int, float64, int, int) {
	// crear el slice
	numeros := make([]int, n)

	// capturar los datos
	for i := 0; i < n; i++ {
		fmt.Print("Escriba un numero: ")
		fmt.Scan(&numeros[i])
	}

	// sumar los numeros
	suma := 0
	for i := 0; i < n; i++ {
		suma += numeros[i]
	}

	// calcular el promedio
	promedio := float64(suma) / float64(n)

	// buscar el maximo
	max := numeros[0]
	for i := 1; i < n; i++ {
		if numeros[i] > max {
			max = numeros[i]
		}
	}

	// buscar el minimo
	min := numeros[0]
	for i := 1; i < n; i++ {
		if numeros[i] < min {
			min = numeros[i]
		}
	}

	return suma, promedio, max, min
}

func main() {
	var (
		numero int
		sum int
		prom float64
		max int
		min int
	)
	fmt.Print("Escriba cuantos numeros deben registrarse: ")
	fmt.Scan(&numero)
	sum, prom, max, min = generarResultados(numero)
	fmt.Println("Suma:", sum)
	fmt.Println("Promedio:", prom)
	fmt.Println("Maximo:", max)
	fmt.Println("Minimo:", min)
}