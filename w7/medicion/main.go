// Obtener numeros primos

package main

import (
	"fmt"
	"math"
	"time"
)

func es_primo_rapido(n int) bool {
	if n < 2 {
		return false
	}

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func es_primo_lento(n int) bool {
	if n < 2 {
		return false
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func generar_n_primos_lento(n int) ([]int, time.Duration) {
	fmt.Println("Generando primos lentos")
	t_inicial := time.Now()
	var primos []int
	for i := 0; len(primos) < n; i++ {
		if es_primo_lento(i) {
			primos = append(primos, i)
		}
	}
	return primos, time.Since(t_inicial)
}

func generar_n_primos_rapido(n int) ([]int, time.Duration) {
	fmt.Println("Generando primos rapidos")
	t_inicial := time.Now()
	var primos []int
	for i := 0; len(primos) < n; i++ {
		if es_primo_rapido(i) {
			primos = append(primos, i)
		}
	}
	return primos, time.Since(t_inicial)
}

func main() {
	const N int = 1

	primos1, dif1 := generar_n_primos_rapido(N)
	fmt.Printf("Generados: %d primos en %v\n", len(primos1), dif1)

	primos2, dif2 := generar_n_primos_lento(N)
	fmt.Printf("Generados: %d primos en %v\n", len(primos2), dif2)
}