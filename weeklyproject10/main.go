package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 450, "Fighting Game")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	isGameOver := false

	pl1 := NewPlayer1()
	pl2 := NewPlayer2()
	h1 := NewHealthBar1(&pl1)
	h2 := NewHealthBar2(&pl2)

	for !rl.WindowShouldClose() {

		if !isGameOver {
			//Updates
			UpdatePlayers(&pl1, &pl2)
			CheckPlayerFloorCollisions(&pl1, &pl2)
			CheckMovement(&pl1, &pl2)
			CheckDamage(&pl1, &pl2, &h1, &h2)

			//Check Game Over
			if pl1.Health <= 0 || pl2.Health <= 0 {
				isGameOver = true
			}

			//Rendering
			rl.BeginDrawing()

			rl.ClearBackground(rl.Black)
			rl.DrawRectangle(0, 350, 800, 100, rl.Orange)
			DrawPlayers(&pl1, &pl2)
			DrawHealthBars(&h1, &h2)

			rl.EndDrawing()
		} else {

			rl.BeginDrawing()

			rl.ClearBackground(rl.Red)

			if pl1.Health > 0 {
				rl.DrawText("Congrats on winning Player 1!", 250, 200, 20, rl.Black)
			} else if pl2.Health > 0 {
				rl.DrawText("Congrats on winning Player 2!", 250, 200, 20, rl.Black)
			}
			rl.DrawText("Press R to restart the game!", 250, 250, 20, rl.Black)

			if rl.IsKeyPressed(rl.KeyR) {
				isGameOver = false
				pl1 = NewPlayer1()
				pl2 = NewPlayer2()
				h1 = NewHealthBar1(&pl1)
				h2 = NewHealthBar2(&pl2)
			}
			rl.EndDrawing()
		}

	}
}
