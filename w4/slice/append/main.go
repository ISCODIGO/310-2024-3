package main

import "fmt"

func main() {
    // Declarar un slice de enteros
    var slice []int
    sl2 := []int{100, 200, 300}

    // Agregar elementos usando append
    slice = append(slice, 10)
    slice = append(slice, 20, 30)
    slice = append(slice, sl2...)

    // Imprimir el slice
    fmt.Println("Slice:", slice)

    // Acceder a un elemento
    fmt.Println("Elemento en la posición 1:", slice[1])
}

// [10, 20, 30, 40]