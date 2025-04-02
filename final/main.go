package main

import rl "github.com/gen2brain/raylib-go/raylib"

// I want to make a typing game. So you will have to type to go from the bottom to top in platforms

func main() {
	rl.InitWindow(800, 450, "Final Game")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	buttons := NewButtons()
	mainMenu := true

	for !rl.WindowShouldClose() {

		if mainMenu {
			buttons.CheckButtons()

			if buttons.Quit.Quit {
				break
			}

			if buttons.Start.Started {
				mainMenu = false
			}
		}

		// --- Rendering ---
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		if mainMenu {
			rl.DrawText("Welcome to my final game!", 270, 150, 20, rl.Black)
			buttons.DrawButtons()
		} else {
			// Game scene here
			rl.DrawText("Game Started!", 350, 200, 20, rl.DarkGreen)
		}

		rl.EndDrawing()
	}
}
