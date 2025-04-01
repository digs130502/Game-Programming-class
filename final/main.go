package main

import rl "github.com/gen2brain/raylib-go/raylib"

// I want to make a typing game. So you will have to type to go from the bottom to top in platforms

func main() {
	rl.InitWindow(800, 450, "Final Game")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	// Initialize buttons
	buttons := NewButtons()

	// Initialize menu
	mainMenu := true

	for !rl.WindowShouldClose() {

		for mainMenu {
			rl.BeginDrawing()

			rl.ClearBackground(rl.RayWhite)
			rl.DrawText("Welcome to my final game!", 270, 150, 20, rl.Black)
			buttons.DrawButtons()

			rl.EndDrawing()
		}
	}
}
