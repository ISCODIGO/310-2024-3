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
		_, err := fmt.Scan(&numeros[i])
		if err != nil {
			fmt.Println("Ocurrio un error en la lectura")
		}
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

func generarResultadosSinLimite() (int, float64, int, int) {
	// crear el slice
	var numeros []int

	// capturar los datos
	for true {
		var entrada int
		fmt.Print("Leer un numero positivo: ")
		fmt.Scan(&entrada)

		if entrada >= 0 {
			// Agregamos al slice
			numeros = append(numeros, entrada)
		} else {
			break;
		}
	}

	// obtener al tamaño de slice
	n := len(numeros)

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
		//numero int
		sum int
		prom float64
		max int
		min int
	)
	//fmt.Print("Escriba cuantos numeros deben registrarse: ")
	//fmt.Scan(&numero)
	sum, prom, max, min = generarResultadosSinLimite()
	fmt.Println("Suma:", sum)
	fmt.Println("Promedio:", prom)
	fmt.Println("Maximo:", max)
	fmt.Println("Minimo:", min)
}