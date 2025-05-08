package main

import (
	"log"
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

//Enemy creation: DONE
//Enemy deletion: DONE (May need to alter detection after adding sprites and animations but to be done LATER)
//Healthbar: DONE
//Player character: DONE
//TODO:
//		Sprites and animations
//		Keyboard typing mechanism
//		Game over screen
//		Different difficulties (Might be able to implement it in the same level. No need for overcomplications with different levels.)
//		Customization. (Text. Color. Sprites and Animations)
//		It's waves based so every level they get a bit faster and more and more. In level 3 they just get infinitely faster

//CURRENTLY DOING:
//		Keyboard typing mechanism
// 		Word assignment: DONE

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
	health := NewHealthBar(&player)
	wordBank, err := LoadWordBank("words.json")
	if err != nil {
		log.Fatalf("Failed to load word bank: %v", err)
	}

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
			player.CheckKeyboardInput()
			zone.CheckEnemyCreation(3, wordBank)
			zone.UpdateEnemies(&player, &health)
			zone.DrawEnemies()
			health.DrawHealthBar()
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
