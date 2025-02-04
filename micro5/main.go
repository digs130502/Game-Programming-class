package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Creature struct {
	name         string
	dexterity    int
	strength     int
	intelligence int
}

func main() {

	rand.Seed(time.Now().UnixNano())

	var creatures []Creature = CreateRPGRoster(2)

	fmt.Println(creatures)

}

func NewRPGCreature(name string) Creature {

	return Creature{
		name:         name,
		dexterity:    1,
		strength:     1,
		intelligence: 1,
	}

}

func RandomRPGCreature(name string) Creature {

	return Creature{
		name:         name,
		dexterity:    rand.Intn(9) + 1,
		strength:     rand.Intn(9) + 1,
		intelligence: rand.Intn(9) + 1,
	}

}

func CreateRPGRoster(size int) []Creature {

	var dexCount int = 0
	var strCount int = 0
	var intCount int = 0
	var creatures []Creature

	for i := 0; i < size; i++ {
		creatures = append(creatures, RandomRPGCreature("Goblin"))
	}

	for i := 0; i < len(creatures); i++ {
		dexCount += creatures[i].dexterity
		strCount += creatures[i].strength
		intCount += creatures[i].intelligence
	}

	fmt.Println("Total Dexterity:", dexCount, "\nTotal Strength:", strCount, "\nTotal Intelligence:", intCount)

	return creatures
}
