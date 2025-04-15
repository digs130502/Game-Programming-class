package main

import (
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 450, "Final Game")
	defer rl.CloseWindow()

	rl.SetExitKey(0) // Disable ESC auto-exit
	rl.SetTargetFPS(60)

	rand.Seed(time.Now().UnixNano())

	buttons := NewButtons()
	mainMenu := true
	paused := false
	zone := NewZone()
	player := NewPlayer()

	for !rl.WindowShouldClose() {

		if mainMenu {
			// --- Main Menu Logic ---
			buttons.CheckButtons()

			if buttons.Quit.Clicked {
				break
			}
			if buttons.Start.Clicked {
				mainMenu = false
				buttons.Start.Clicked = false
			}

			rl.BeginDrawing()
			rl.ClearBackground(rl.RayWhite)
			rl.DrawText("Welcome to my final game!", 270, 150, 20, rl.Black)
			buttons.DrawButtons()
			rl.EndDrawing()

		} else if paused {
			// --- Pause Menu Logic ---
			buttons.CheckResumeButtons()

			if buttons.Resume.Clicked {
				paused = false
				buttons.Resume.Clicked = false
			}
			if buttons.Menu.Clicked {
				mainMenu = true
				paused = false
				buttons.Menu.Clicked = false
			}
			if buttons.Quit.Clicked {
				break
			}

			rl.BeginDrawing()
			rl.ClearBackground(rl.LightGray)
			rl.DrawText("Game Paused", 300, 90, 30, rl.Black)
			buttons.DrawPauseButtons()
			rl.EndDrawing()

		} else {
			// --- Gameplay Logic ---
			zone.CheckEnemyCreation(3)
			zone.UpdateEnemies()
			zone.DrawEnemies()
			// Pressing ESC will pause the game
			if rl.IsKeyPressed(rl.KeyEscape) {
				paused = true
			}

			rl.BeginDrawing()
			rl.ClearBackground(rl.RayWhite)
			player.DrawPlayer()
			rl.EndDrawing()
		}
	}
}
