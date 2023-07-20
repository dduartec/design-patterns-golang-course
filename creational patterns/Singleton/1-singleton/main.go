package main

import (
	"fmt"
	"sync"
)

type singletonDatabase struct {
	capitals map[string]int
}

func (db *singletonDatabase) GetPopulation(name string) int {
	return db.capitals[name]
}

func GetTotalPipulation(cities []string) int {
	result := 0
	for _, city := range cities {
		result += GetSingletonDatabase().GetPopulation(city)
	}
	return result
}

// sync.Once init() --thread safety
// laziness

var (
	once     sync.Once
	instance *singletonDatabase
)

func GetSingletonDatabase() *singletonDatabase {
	once.Do(func() {
		// load database
		db := singletonDatabase{
			capitals: map[string]int{"Bogota": 8000, "Londres": 12000},
		}
		instance = &db
	})
	return instance
}

func main() {
	db := GetSingletonDatabase()
	b := db.GetPopulation("Bogota")
	fmt.Println(b)
}
