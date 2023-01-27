package main

import "fmt"

/*
pode-se nomear o retorno da função
(soma int, sub int)
dentro da função deverão ter valores atribuidos a
todas variaveis nomeadas
*/
func calculosMatemaricos(n1 int, n2 int) (soma int, sub int) {
	soma = n1 + n2
	sub = n1 - n2
	return
}

func main() {
	fmt.Println(calculosMatemaricos(10, 4))

	soma, sub := calculosMatemaricos(45, 9)
	fmt.Println(soma, sub)
}
