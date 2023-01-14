package main

import (
	"fmt"
	"strconv"
)

// Parametros e retornos Tipados
func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

/*
importante: é possivel tipar todos parametros, tipadando somente o último
importante: é possivel retornar mais de um valor em funções GO
*/
func calculosMatematicos(n1, n2, n3 int8) (int8, int8) {
	soma := n1 + n2 + n3
	subtracao := n1 - n2 - n3
	return soma, subtracao
}

func main() {
	soma := somar(1, 5)
	fmt.Println(soma)

	var teste = func(txt string) string {
		return txt
	}

	var teste2 = func(txt string) string {
		soma := somar(1, 35)
		return txt + strconv.Itoa(int(soma))
	}

	nome := "Developer"

	// Variaveis criadas em escopo superior são acessiveis nos inferiores ( dalhe JS )
	var teste3 = func(txt string) string {
		return txt + nome
	}

	sum, sub := calculosMatematicos(10, 50, 65)
	fmt.Println("Soma:", sum, "Subtração:", sub)

	resultado := teste("Seu nome é: Sidnei Simeão")
	fmt.Println(resultado)

	resultado2 := teste2("Sua idade é: ")
	fmt.Println(resultado2)

	resultado3 := teste3("Sua profissão é: ")
	fmt.Println(resultado3)

}
