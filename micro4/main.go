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

	var newCreature Creature = NewRPGCreature("varnon")
	var randCreature Creature = RandomRPGCreature("Goblin")

	fmt.Println("New RPG Creature", "\nName:", newCreature.name, "\nDexterity:", newCreature.dexterity, "\nStrength:", newCreature.strength, "\nIntelligence:", newCreature.intelligence)
	fmt.Println("Random RPG Creature", "\nName:", randCreature.name, "\nDexterity", randCreature.dexterity, "\nStrength:", randCreature.strength, "\nIntelligence:", randCreature.intelligence)

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
