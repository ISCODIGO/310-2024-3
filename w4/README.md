---
marp: true
---

# Otras operaciones de Slice

---
Las siguientes operaciones, de momento, no se pueden establecer con funciones propias de Golang:
- Remover elementos
- Insertar en una posicion dada un elemento.
---
# Remover un elemento
Se hacen dos cortes en el slice a partir de la posicion dada.
Dado `[10, 20, 30, 40]`, al remover 20 queda algo similar a `10 + [30, 40]`
```go
func remover([]int slice, int posicion) []int{
    return append(slice[:posicion], slice[posicion + 1:]...)
}
```
---
# Insertar un elemento
Similar a la eliminacion dada una posicion se hace un corte, donde se inserta el nuevo valor.
Dado `[10, 20, 30]`, insertar 40 en la posicion 1 -> `10 + 40 + [20, 30]`
```go
func insertar(slice []int, pos int, valor int) []int {
	return append(slice[:pos], valor, slice[pos:]...)
}
```
---
# Structs
---
## Como objetos
En Go, los structs se utilizan para crear tipos personalizados que pueden contener datos (similar a los objetos en POO).

---

```go
package main

import "fmt"

// Define a struct (like a class in other languages)
type Car struct {
    Brand string
    Model string
    Year  int
}

func main() {
    // Create an instance of Car
    myCar := Car{
        Brand: "Toyota",
        Model: "Corolla",
        Year:  2020,
    }

    fmt.Println("Car:", myCar.Brand, myCar.Model, myCar.Year)
}
```
---
## Funciones

Go te permite definir métodos en structs, de manera similar a cómo definirías métodos en clases en otros lenguajes.
```go
package main

import "fmt"

func (c Car) Details() string {
    return fmt.Sprintf("Brand: %s, Model: %s, Year: %d", c.Brand, c.Model, c.Year)
}

// Constructor function
func NewCar(brand, model string, year int) Car {
    return Car{Brand: brand, Model: model, Year: year}
}

func main() {
    myCar := Car{Brand: "Toyota", Model: "Corolla", Year: 2020}
    fmt.Println(myCar.Details())

    // usando el constructor
    myCar2 := NewCar("Toyota", "Corolla", 2020)
}

```

---
## Encapsulamiento

Go logra el encapsulamiento utilizando reglas de visibilidad. Los identificadores que comienzan con una letra mayúscula son exportados (públicos), mientras que los que comienzan con una letra minúscula son no exportados (privados).

```go
    type Car struct {
    Brand string // public field
    model string // private field
}

// public method
func (c Car) ShowModel() string {
    return c.model
}
```

---
# Interfaces
---
**Interfaces**
Go utiliza interfaces para implementar polimorfismo. Una interfaz es una colección de firmas de métodos, y cualquier tipo que implemente esos métodos se considera que satisface la interfaz.

---
```go
// Definir la interface
type Vehicle interface {
    Drive() string
}

// Implement la interface para el struct Car
func (c Car) Drive() string {
    return fmt.Sprintf("Driving %s %s", c.Brand, c.Model)
}

// Esta funcion toma un objeto que implementa la interface Vehicle o cualquier tipo que implemente el metodo Drive()
func StartVehicle(v Vehicle) {
    fmt.Println(v.Drive())
}

func main() {
    myCar := Car{Brand: "Toyota", Model: "Corolla", Year: 2020}
    StartVehicle(myCar)
}

