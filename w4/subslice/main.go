package main

import "fmt"

func main() {
    // Declarar y inicializar un array
    arr := [5]int{10, 20, 30, 40, 50}
    // Crear un sub-slice del array
    subSlice := arr[:3]
    // Imprimir el sub-slice
    fmt.Println("Sub-slice:", subSlice)
}