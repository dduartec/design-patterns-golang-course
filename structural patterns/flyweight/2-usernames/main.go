package main

import (
	"fmt"
	"strings"
)

type User struct {
	FullName string
}

func NewUser(fullname string) *User {
	return &User{fullname}
}

var allNames []string

type BetterUser struct {
	names []uint8
}

func NewBetterUser(fullname string) *BetterUser {
	getOrAdd := func(s string) uint8 {
		for i := range allNames {
			if allNames[i] == s {
				return uint8(i)
			}
		}
		allNames = append(allNames, s)
		return uint8(len(allNames) - 1)
	}
	res := BetterUser{}
	parts := strings.Split(fullname, " ")
	for _, p := range parts {
		res.names = append(res.names, getOrAdd(p))
	}

	return &res
}

func (u *BetterUser) Fullname() string {
	var parts []string
	for _, id := range u.names {
		parts = append(parts, allNames[id])
	}
	return strings.Join(parts, " ")
}

func main() {
	// allocates more memory for the same names
	// john := NewUser("john doe")
	// john2 := NewUser("john smith")
	john := NewBetterUser("john doe")
	john2 := NewBetterUser("john smith")

	fmt.Println(john.Fullname())
	fmt.Println(john2.Fullname())

}
