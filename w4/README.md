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
Dado `[10, 20, 30]`, al remover 20 queda algo similar a `[10] + [30]`
```go
func remover([]int slice, int posicion) []int{
    return append(slice[:posicion], slice[posicion + 1:]...)
}
```
---
# Insertar un elemento
Similar a la eliminacion dada una posicion se hace un corte, donde se inserta el nuevo valor.
Dado `[10, 20, 30]`, insertar 40 en la posicion 1 -> `[10] + 40 + [20, 30]`
```go
func insertar(slice []int, pos int, valor int) []int {
	return append(slice[:pos], valor, slice[pos:]...)
}
```