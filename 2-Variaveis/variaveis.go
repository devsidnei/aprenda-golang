package main

import "fmt"

func main() {
	var variavel1 string = "Variável 1 - Tipo explícito"
	variavel2 := "Variável 2 - Tipo implícito"

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "Variável 3 - Conjunto explícito"
		variavel4 string = "Variável 4 - Conjunto explícito"
	)

	fmt.Println(variavel3)
	fmt.Println(variavel4)

	variavel5, variavel6 := "Variavel 5 - Conjunto implicíto", 100.00

	fmt.Println(variavel5)
	fmt.Println(variavel6)

	const constante string = "Constante 1"

	fmt.Println(constante)
}
