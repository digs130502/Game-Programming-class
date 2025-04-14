package main

import (
	"fmt"
	"math/rand/v2"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Forest struct {
	Trees           []Tree
	BackgroundColor rl.Color
	TreeColor       rl.Color
	TreeSize        float32
}

func (f *Forest) Generate(seed1, seed2 uint64) {
	pcg := rand.NewPCG(seed1, seed2)
	forestRand := rand.New(pcg)

	f.BackgroundColor = rl.NewColor(
		uint8(forestRand.IntN(255)),
		uint8(forestRand.IntN(255)),
		uint8(forestRand.IntN(255)),
		255,
	)
	f.TreeColor = rl.NewColor(
		uint8(forestRand.IntN(255)),
		uint8(forestRand.IntN(255)),
		uint8(forestRand.IntN(255)),
		255,
	)

	f.TreeSize = forestRand.Float32() * 50

	numTrees := forestRand.IntN(50) + 50

	for i := 0; i < numTrees; i++ {
		f.Trees = append(f.Trees, Tree{forestRand.Int32N(1000), forestRand.Int32N(1000)})
	}
	fmt.Println(len(f.Trees))
	fmt.Println(f.TreeColor)
}

func (f *Forest) DrawForest() {
	rl.ClearBackground(f.BackgroundColor)
	for _, tree := range f.Trees {
		tree.DrawTree(f.TreeSize, f.TreeColor)
	}
}
