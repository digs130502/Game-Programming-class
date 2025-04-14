package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Tree struct {
	X int32
	Y int32
}

func (t *Tree) DrawTree(radius float32, color rl.Color) {
	rl.DrawCircle(t.X, t.Y, radius, color)
}
