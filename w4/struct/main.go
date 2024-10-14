package main

import "fmt"

type Engine struct {
    Horsepower int
}

type Car struct {
    Brand  string
    Model  string
    Engine // Embeber el struct Engine
}

func main() {
    myCar := Car{
        Brand:  "Toyota",
        Model:  "Corolla",
        Engine: Engine{Horsepower: 130},
    }
    fmt.Println("Potencia del motor:", myCar.Horsepower)
}