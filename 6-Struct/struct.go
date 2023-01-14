package main

import (
	"encoding/json"
	"fmt"
)

type Customer struct {
	DocumentnNr int64
	Name        string
	Code        int32
	age         int `json:"age"`
	Address     Address
	//phone      phone
}

type Address struct {
	country      string
	City         string
	neighborhood string
	street       string
	number       string
}

type phone struct {
	area   int8
	number int64
}

func main() {

	var cliente1 Customer
	cliente1.Name = "Sidnei Simeão"
	cliente1.DocumentnNr = 8828926688
	cliente1.Code = 73212
	cliente1.Address.City = "Poços de Caldas"
	cliente1.age = 36

	fmt.Printf("%+v\n", cliente1)

	p := Customer{Name: "Welton Terraplanista", DocumentnNr: 78965464, Code: 654321}
	b, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}
