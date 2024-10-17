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

