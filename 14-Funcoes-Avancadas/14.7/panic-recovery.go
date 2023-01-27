package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovering from panic:", r)
		}
	}()

	fmt.Println("Starting function execution")
	causePanic()
	fmt.Println("Function execution completed")
}

func causePanic() {
	fmt.Println("Panicking")
	panic("Panic caused")
}
