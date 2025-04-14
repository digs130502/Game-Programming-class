package main

import rl "github.com/gen2brain/raylib-go/raylib"

// I want to make a typing game. So you will have to type to go from the bottom to top in platforms
//TODO: Make a pause game feature where player presses "ESC" and a pause screen is made

func main() {
	rl.InitWindow(800, 450, "Final Game")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	buttons := NewButtons()
	mainMenu := true
	player := NewPlayer()

	for !rl.WindowShouldClose() {

		if mainMenu {
			// Updates
			buttons.CheckButtons()

			//Check if buttons are clicked
			if buttons.Quit.Clicked {
				break
			}

			if buttons.Start.Clicked {
				mainMenu = false
			}

			// Rendering
			rl.BeginDrawing()

			rl.ClearBackground(rl.RayWhite)
			rl.DrawText("Welcome to my final game!", 270, 150, 20, rl.Black)
			buttons.DrawButtons()

			rl.EndDrawing()
		} else {
			// Game scene here
			rl.BeginDrawing()

			rl.ClearBackground(rl.RayWhite)
			player.DrawPlayer()

			rl.EndDrawing()
		}
	}
}
