package main

import "fmt"

func nota(valor int) (string, bool) {
	/*

	A = [90, 100]
	B = [80, 89]
	C = [70, 79]
	D = [60, 69]
	-------------
	F = [0, 59]

	*/

	if valor >= 0 && valor <= 59 {
		return "F", false
	} else if valor >= 60 && valor <= 69 {
		return "D", true
	} else if valor >= 70 && valor <= 79 {
		return "C", true
	} else if valor >= 80 && valor <= 89 {
		return "B", true
	} else if valor >= 90 && valor <= 100 {
		return "A", true
	} else {
		return "-", false
	}
}

func main() {
	var x int
	fmt.Scan(&x)
	a, b := nota(x)
	fmt.Println("La nota es", x, "y la calificacion es", a, "aprobo?", b)
}