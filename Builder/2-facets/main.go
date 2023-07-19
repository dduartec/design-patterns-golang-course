package main

import "fmt"

type Person struct {
	// address
	StreetAddress, PostCode, City string

	//job
	CompanyName, Position string
	AnnualIncome          int
}

type PersonBuilder struct {
	person *Person
}

func (pb *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*pb}
}

func (pb *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{*pb}
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{&Person{}}
}

type PersonAddressBuilder struct {
	PersonBuilder
}

func (it *PersonAddressBuilder) At(
	streetAddress string) *PersonAddressBuilder {
	it.person.StreetAddress = streetAddress
	return it
}

func (it *PersonAddressBuilder) In(
	city string) *PersonAddressBuilder {
	it.person.City = city
	return it
}

func (it *PersonAddressBuilder) WithPostcode(
	postcode string) *PersonAddressBuilder {
	it.person.PostCode = postcode
	return it
}

type PersonJobBuilder struct {
	PersonBuilder
}

func (pjb *PersonJobBuilder) At(
	companyName string) *PersonJobBuilder {
	pjb.person.CompanyName = companyName
	return pjb
}

func (pjb *PersonJobBuilder) AsA(
	position string) *PersonJobBuilder {
	pjb.person.Position = position
	return pjb
}

func (pjb *PersonJobBuilder) Earning(
	annualIncome int) *PersonJobBuilder {
	pjb.person.AnnualIncome = annualIncome
	return pjb
}

func main() {
	pb := NewPersonBuilder()
	pb.Lives().
		At("calle 26 32 14").
		In("Bogotá").
		WithPostcode("1234123")

	pb.Works().
		At("MELI").
		AsA("Software dev").
		Earning(7000000)

	fmt.Println(pb.person)

}
