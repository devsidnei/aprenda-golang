package main

import "fmt"

func main() {

	var myMap = make(map[int]string)
	myMap[0] = "Sidnei"

	fmt.Println(myMap)

	myMap2 := map[int]string{
		0: "joão",
	}

	fmt.Println(myMap2)

	var myMap3 = map[string]string{
		"0": "Sauro",
	}

	fmt.Println(myMap3)

	payload := map[string]map[string]string{
		"customer": {
			"name": "Sidnei",
			"age":  "36",
		},
		"address": {
			"neighborhood": "São Jośe",
			"city":         "Poços de Caldas",
			"postalcd":     "3704233",
		},
		"phone": {
			"areacd":    "35",
			"phonearea": "9999999999",
		},
	}

	fmt.Println(payload["customer"]["name"])

}
