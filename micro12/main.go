package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 450, "Micro Assignment 12")
	defer rl.CloseWindow()

	if !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.EndDrawing()
	}
}
