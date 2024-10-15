package main

import (
    "fmt"
    "strings"
)

func main() {
    ejemploConcatenacion()
    ejemploLongitud()
    ejemploSubcadena()
    ejemploComparacion()
    ejemploConversionMayusculasMinusculas()
    ejemploReemplazoSubcadena()
    ejemploDividirCadena()
    ejemploPrefijoSufijo()
    ejemploIndiceSubcadena()
    ejemploRecorridoCaracteres()
}

func ejemploConcatenacion() {
    cadena1 := "Hola"
    cadena2 := "Mundo"
    resultado := cadena1 + " " + cadena2
    fmt.Println("Concatenación:", resultado) // Salida: Hola Mundo
}

func ejemploLongitud() {
    cadena := "Golang"
    longitud := len(cadena)
    fmt.Println("Longitud de la cadena:", longitud) // Salida: 6
}

func ejemploSubcadena() {
    cadena := "Aprender Golang"
    subcadena := cadena[9:] // Extrae desde el índice 9 hasta el final
    fmt.Println("Subcadena:", subcadena) // Salida: Golang
}

func ejemploComparacion() {
    cadena1 := "golang"
    cadena2 := "Golang"

    if cadena1 == cadena2 {
        fmt.Println("Comparación: Las cadenas son iguales")
    } else {
        fmt.Println("Comparación: Las cadenas son diferentes") // Salida: Las cadenas son diferentes
    }
}

func ejemploConversionMayusculasMinusculas() {
    cadena := "Aprender Golang"
    fmt.Println("Mayúsculas:", strings.ToUpper(cadena)) // Salida: APRENDER GOLANG
    fmt.Println("Minúsculas:", strings.ToLower(cadena)) // Salida: aprender golang
}

func ejemploReemplazoSubcadena() {
    cadena := "Hola Mundo"
    reemplazada := strings.Replace(cadena, "Mundo", "Golang", 1)
    fmt.Println("Reemplazo de subcadena:", reemplazada) // Salida: Hola Golang
}

func ejemploDividirCadena() {
    cadena := "uno,dos,tres,cuatro"
    partes := strings.Split(cadena, ",")
    fmt.Println("Dividir cadena:", partes) // Salida: [uno dos tres cuatro]
}

func ejemploPrefijoSufijo() {
    cadena := "Golang es genial"

    if strings.HasPrefix(cadena, "Golang") {
        fmt.Println("Prefijo: La cadena comienza con 'Golang'") // Salida: La cadena comienza con 'Golang'
    }

    if strings.HasSuffix(cadena, "genial") {
        fmt.Println("Sufijo: La cadena termina con 'genial'") // Salida: La cadena termina con 'genial'
    }
}

func ejemploIndiceSubcadena() {
    cadena := "Aprender Go es divertido"
    indice := strings.Index(cadena, "Go")
    fmt.Println("Índice de subcadena:", indice) // Salida: El índice de 'Go' es: 9
}

func ejemploRecorridoCaracteres() {
    cadena := "Golang"
    fmt.Println("Recorrido de caracteres:")
    for i, caracter := range cadena {
        fmt.Printf("Índice %d: %c\n", i, caracter)
    }
    // Salida:
    // Índice 0: G
    // Índice 1: o
    // Índice 2: l
    // Índice 3: a
    // Índice 4: n
    // Índice 5: g
}