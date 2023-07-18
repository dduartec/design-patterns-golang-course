package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var entryCount = 0

// Journal handles the entries
type Journal struct {
	entries []string
}

func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s", entryCount, text)
	j.entries = append(j.entries, entry)
	return entryCount
}

func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

/*
Would not follow SRP:

	func (j *Journal) Save()  {
		...
	}
*/

// Separation of concerns: Separate the domain(Journal) from the persistance

type Persistance struct {
	lineSeparator string
}

func (p Persistance) Save(sObj fmt.Stringer, filename string) {
	ioutil.WriteFile(filename, []byte(sObj.String()), 0644)
}

func main() {
	j := &Journal{}
	j.AddEntry("entry 1")
	j.AddEntry("entry 2")
	fmt.Println(j)

	p := Persistance{}

	p.Save(j, "journal.txt")

}
