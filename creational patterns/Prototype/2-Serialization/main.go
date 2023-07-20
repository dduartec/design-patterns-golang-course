package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Address struct {
	Street, City, Country string
}

type Person struct {
	Name    string
	Address *Address
	Friend  []string
}

func DeepCopy(data interface{}, pointer interface{}) {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	e.Encode(data)
	d := gob.NewDecoder(&b)
	d.Decode(pointer)
}

func main() {

	diego := Person{
		Name: "Diego",
		Address: &Address{
			Street:  "calle 26",
			City:    "Bogota",
			Country: "Colombia",
		},
	}
	jane := diego
	jane.Name = "Jane"
	jane.Address.City = "Medallo" // overrides the value because is a pointer

	dan := Person{}
	DeepCopy(diego, &dan)
	dan.Name = "Daniela"
	dan.Address.City = "Cali"

	fmt.Println(diego, diego.Address)
	fmt.Println(jane, jane.Address)
	fmt.Println(dan, dan.Address)
}
