package main

import "fmt"

func main() {

	// Inteiros até 255
	var numeroInt8 int8 = -99
	fmt.Println(numeroInt8)

	// Alias para int8
	var numeroByte byte = 255
	fmt.Println(numeroByte)

	// Inteiros positivos e negativos
	var numeroInt64 int64 = -1000000000000000
	fmt.Println(numeroInt64)

	// Somente inteiros positivos
	var numeroUint8 uint8 = 99
	fmt.Println(numeroUint8)

	// Alias de int32 - aspas simples considera valor do caractere na tabela ASC
	var numeroRune rune = 'a'
	fmt.Println(numeroRune)

	// Letras e numeros, sempre entre aspas duplas
	var nome string = "Sidnei Simeão"
	fmt.Println(nome)

	// Somente tru|false
	var estudar bool = true
	fmt.Println(estudar)

	// Tipos possuem valor default
	var dormir bool //false
	fmt.Println(dormir)

	// Tipos possuem valor default
	var carro string // ""
	fmt.Println(carro)

	// Tipos possuem valor default
	var pneus int16 // 0
	fmt.Println(pneus)

}
