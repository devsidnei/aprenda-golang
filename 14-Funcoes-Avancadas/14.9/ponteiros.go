package main

import "fmt"

func inverterSinal(n int) int {
	return n * -1
}

func inverterSinalPonteiro(n *int) {
	*n = *n * -1
}

func main() {

	numero := 10
	inverterSinal(numero)
	inverterSinalPonteiro(&numero)
	fmt.Println(numero)
}
