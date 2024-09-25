package main

import "fmt"

func nota(valor int) (string, bool) {
	fmt.Println("Valor dentro de la funcion: ", valor)
	/*

	A = [90, 100]
	B = [80, 89]
	C = [70, 79]
	D = [60, 69]
	-------------
	F = [0, 59]

	*/

	if 0 <= valor && valor <= 59 {
		return "F", false
	} else if 60 <= valor && valor <= 69 {
		return "D", true
	} else if 70 <= valor && valor <= 79 {
		return "C", true
	} else if 80 <= valor && valor <= 89 {
		return "B", true
	} else if 90 <= valor && valor <= 100 {
		return "A", true
	} else {
		return "-", false
	}
}

func sumar(valores ...int) int {
	var total int
	for _, valor := range valores {
		total += valor
	}
	return total
}


func main() {
	var x int
	fmt.Scan(&x)
	a, b := nota(x)
	fmt.Println("La nota es", x, "y la calificacion es", a, "aprobo?", b)
	fmt.Println(sumar(1, 2), sumar(1, 2, 3, 4), sumar(1, 2, 3, 4, 5, 6))
}