package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

func DeepCopy(data interface{}, pointer interface{}) {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	e.Encode(data)
	d := gob.NewDecoder(&b)
	d.Decode(pointer)
}

type Address struct {
	Street, City, Country string
}

type Person struct {
	Name    string
	Address *Address
}

var (
	mainAddress = Person{Address: &Address{"calle 26", "Bogota", " Col"}}
	auxAddress  = Person{Address: &Address{"calle 26", "Medallo", " Col"}}
)

func NewPerson(proto Person, name string) *Person {
	res := &Person{}
	DeepCopy(proto, res)
	res.Name = name
	return res
}

func NewMainAddressPerson(name string) *Person {
	return NewPerson(mainAddress, name)
}

func NewAuxAddressPerson(name string) *Person {
	return NewPerson(auxAddress, name)
}

func main() {
	diego := NewMainAddressPerson("Diego")
	dan := NewAuxAddressPerson("Dan")
	fmt.Println(diego, diego.Address)
	fmt.Println(dan, dan.Address)

}
