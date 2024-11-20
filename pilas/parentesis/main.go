package main

import (
	"fmt"
	"pilas"
)

func check(palabra string) bool {
	stack := pilas.New(10)
	for _, c := range palabra {
		if c == '(' {
			stack.Push(int(c))
		}
		if c == ')' {
			if stack.IsEmpty() {
				return false
			}
			stack.Pop()
		}
	}
	return stack.IsEmpty()
}

func main() {
	fmt.Print("Ingrese la expresion:")

	var cadena string
	fmt.Scanln(&cadena)

	fmt.Println("La expresion es correcta:", check(cadena))
}