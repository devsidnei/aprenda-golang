package main

import (
	"fmt"
	"time"
)

func apply(fn func(int) int, x int) int {
	return fn(x)
}

func main() {

	anoAtual := time.Now().Year()

	exibirIdade := func(anoNascimento int) string {
		idade := anoAtual - anoNascimento
		return fmt.Sprintf("Sua idade é %d", idade)
	}

	fmt.Println(apply(func(x int) int { return x * x }, 10))

	fmt.Println(exibirIdade(1986))
}
