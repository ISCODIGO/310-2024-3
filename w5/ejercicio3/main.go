// Ejercicio: Todo List
// Crea una estructura `Tarea` con campos como titulo, descripcion y completado.
// Utiliza punteros para crear funciones que añadan, marquen como completadas y muestren las tareas.

package main

import "fmt"

// Definición de la estructura Tarea
type Tarea struct {
    titulo       string
    descripcion  string
    completado   bool
}

// Método para marcar una tarea como completada
func (t *Tarea) completar() {
    t.completado = true
}

// Método toString para representar la tarea como una cadena
func (t Tarea) toString(index int) string {
    estado := "[ ]"
    if t.completado {
        estado = "[x]"
    }
    return fmt.Sprintf("Tarea %d: %s %s\nDescripcion: %s\n\n", index+1, estado, t.titulo, t.descripcion)
}

// Función para imprimir todas las tareas
func imprimirTareas(tareas []Tarea) {
    for i, tarea := range tareas {
        fmt.Print(tarea.toString(i))
    }
}

func main() {
    // Crear una lista de tareas
    tareas := []Tarea{
        {titulo: "Comprar leche", descripcion: "Ir al supermercado y comprar leche", completado: false},
        {titulo: "Llamar al doctor", descripcion: "Hacer una cita con el doctor para el chequeo anual", completado: false},
        {titulo: "Pagar facturas", descripcion: "Pagar las facturas de electricidad y agua", completado: false},
    }

    // Imprimir las tareas antes de completarlas
    fmt.Println("Antes de completar las tareas:")
    imprimirTareas(tareas)

    // Marcar algunas tareas como completadas
    tareas[0].completar()
    tareas[2].completar()

    // Imprimir las tareas después de completarlas
    fmt.Println("Después de completar algunas tareas:")
    imprimirTareas(tareas)

    tareas[0].completado = false
    fmt.Println(tareas[0].completado)
}

// salida, error
// if error != nil {
//     fmt.Println(error)
// }