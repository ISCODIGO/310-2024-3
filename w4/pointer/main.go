// test the pointer with functions
package main

import "fmt"

func main() {
	x := 10
	fmt.Println("Ejecutando una funcion sin punteros, es una copia de x")
	fmt.Println("x antes", x)
	no_change(x)
	fmt.Println("x despues", x)

	fmt.Println("Ejecutando una funcion con punteros, es una referencia de x")
	fmt.Println("x antes", x)
	change(&x)
	fmt.Println("x despues", x)

	fmt.Println("Valor de x", x)
	fmt.Println("Direccion de x", &x)

    original := 10
    puntero := &original

	fmt.Println("\n*** Descripcion de punteros ***")

    fmt.Println("Variable original:", original)
	fmt.Println("Direccion de la variable original:", &original)
	fmt.Println("Valor del puntero:", puntero)
    fmt.Println("Direccion del puntero:", &puntero)
	fmt.Println("Valor de original a traves del puntero:", *puntero)
}

func no_change(x int) {
	x *= 2
}

func change(x *int) {
	*x *= 2
}