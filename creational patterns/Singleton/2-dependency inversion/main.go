package main

import (
	"fmt"
	"sync"
)

type Database interface {
	GetPopulation(name string) int
}

type singletonDatabase struct {
	capitals map[string]int
}

func (db *singletonDatabase) GetPopulation(name string) int {
	return db.capitals[name]
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

type DummyDatabase struct {
	summyData map[string]int
}

func (db *DummyDatabase) GetPopulation(name string) int {
	if len(db.summyData) == 0 {
		db.summyData = map[string]int{"Bogota": 3000, "Londres": 1000}
	}
	return db.summyData[name]
}

func GetTotalPopulation(cities []string) int {
	result := 0
	for _, city := range cities {
		result += GetSingletonDatabase().GetPopulation(city) // Violates DIP
	}
	return result
}

func GetTotalPopulationEx(db Database, cities []string) int {
	result := 0
	for _, city := range cities {
		result += db.GetPopulation(city) // Violates DIP
	}
	return result
}

func main() {

	names := []string{"Bogota", "Londres"}
	dummy := DummyDatabase{}
	tp := GetTotalPopulationEx(&dummy, names)
	fmt.Println(tp == 4000)
}
