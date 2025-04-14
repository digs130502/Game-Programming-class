package main

import (
	"fmt"
	"math/rand/v2"
)

type Potion struct {
	Power  int
	Uses   int
	Effect string
}

func (p *Potion) Generate(seed1, seed2 uint64) {
	pcg := rand.NewPCG(seed1, seed2)
	potionRand := rand.New(pcg)

	potionEffect := potionRand.IntN(4)
	if potionEffect == 0 {
		p.Effect = "Invisibility"
	} else if potionEffect == 1 {
		p.Effect = "Healing"
	} else if potionEffect == 2 {
		p.Effect = "Attack"
	} else {
		p.Effect = "Defense"
	}

	p.Power = potionRand.IntN(5) + 1
	p.Uses = potionRand.IntN(5) + 1

	fmt.Println("Potion Created!")
	fmt.Println("Potion Effect = ", potionEffect, "Potion Power = ", p.Power, "Potion Uses ", p.Uses)
}
