package main

import (
	"fmt"
	"math"
)


type Rectangulo struct {
	alto float64
	ancho float64
}

func (r Rectangulo) area() float64 {
	return r.alto * r.ancho
}

func (r *Rectangulo) ampliar(f int) {
	r.alto *= float64(f)
	r.ancho *= float64(f)
}

func (r Rectangulo) cupo(otro Rectangulo) bool {
	min_otro := math.Min(otro.alto, otro.ancho)
	max_otro := math.Max(otro.alto, otro.ancho)
	min_r := math.Min(r.alto, r.ancho)
	max_r := math.Max(r.alto, r.ancho)

	if min_otro <= min_r && max_otro <= max_r {
		return true
	}

	return false
}

func main() {
	rect1 := Rectangulo {
		alto: 20,
		ancho: 20,
	}

	rect2 := Rectangulo {
		alto: 20,
		ancho: 10,
	}

	//factor := 10

	// fmt.Println("El area es:", rect1.area())
	// fmt.Printf("El rectangulo original es %v\n", rect1)
	// rect1.ampliar(factor)
	// fmt.Printf("Ampliado %v veces es: %v", factor, rect1)

	fmt.Println(rect1.cupo(rect2))
}