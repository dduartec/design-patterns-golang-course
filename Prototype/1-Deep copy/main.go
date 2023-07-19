package main

import "fmt"

type Address struct {
	Street, City, Country string
}

func (address *Address) DeepCopy() *Address {
	return &Address{
		Street:  address.Street,
		City:    address.City,
		Country: address.Country,
	}
}

type Person struct {
	Name    string
	Address *Address
	Friend  []string
}

func (p *Person) DeepCopy() *Person {
	q := *p
	q.Address = p.Address.DeepCopy()
	copy(q.Friend, p.Friend)
	return &q
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

	dan := diego.DeepCopy()
	dan.Name = "Daniela"
	dan.Address.City = "Cali"

	fmt.Println(diego, diego.Address)
	fmt.Println(jane, jane.Address)
	fmt.Println(dan, dan.Address)
}
