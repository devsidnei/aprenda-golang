package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	if myNumber := ramdomNumber(0, 100); myNumber < 50 {
		fmt.Println("É possivel inicializar variavel dentro da condicional do IF")
	}
}

func ramdomNumber(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max - min)
}
