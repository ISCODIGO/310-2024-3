package main

import (
	"fmt"
	"strconv"
)

// Crear una struct
type Punto struct {
	// Atributos
	X int
	Y int
}

// Metodo
func (p Punto) print() string {
	p.X = 40
	return "x=" + strconv.Itoa(p.X) + ", y=" + strconv.Itoa(p.Y)
}

func modificarCoordenadas(p *Punto) {
	p.X = 20
	p.Y = 20
}

func main() {
	p1 := Punto{
		X: 10,
		Y: 10,
	}

	fmt.Println(p1.print())
	fmt.Println("El valor de x se mantiene?", p1.X)

	fmt.Println("Llama a la funcion modificarCoordenadas")
	modificarCoordenadas(&p1)
	fmt.Print(p1)
}