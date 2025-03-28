package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 450, "Midterm Game")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	playerCircle := NewPlayerCircle()
	zone := NewZone()

	for !rl.WindowShouldClose() {

		if !zone.GameOver {
			playerCircle.UpdateCircle()
			zone.CheckMineCreation(playerCircle)
			zone.CheckPlayerCollisions(playerCircle)

			rl.BeginDrawing()
			rl.ClearBackground(rl.RayWhite)
			playerCircle.DrawPlayerCircle()
			zone.DrawMines()

			countText := fmt.Sprintf("Mines placed: %d", zone.Count)
			rl.DrawText(countText, 10, 10, 20, rl.Black)

			rl.EndDrawing()
		} else {
			rl.BeginDrawing()
			rl.ClearBackground(rl.Red)

			gameScore := fmt.Sprintf("Your Score was: %d", zone.Count)
			rl.DrawText("You lost! Press R to restart!", 400, 225, 20, rl.Black)
			rl.DrawText(gameScore, 400, 250, 20, rl.Black)

			if rl.IsKeyPressed(rl.KeyR) {
				zone.GameOver = false
				zone.Mines = []Mine{}
				zone.Count = 0
			}
			rl.EndDrawing()
		}

	}
}
