package main

import "fmt"

// Funções variaticas sao aquelas que recebem uma quantidade indefinida de parametros
func somar(numeros ...int) int {
	soma := 0
	for _, numero := range numeros {
		soma += numero
	}
	return soma
}

func main() {
	fmt.Println(somar(1, 2, 1, 2, 1, 3, 3, 4, 5, 6, 5, 3))
}
