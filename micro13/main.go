package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(800, 450, "Micro Assignment 13")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)
}
