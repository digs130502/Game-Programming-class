package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(800, 450, "Fighting Game")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	pl1 := NewPlayer1()
	pl2 := NewPlayer2()

	for !rl.WindowShouldClose() {

		//Updates
		UpdatePlayers(&pl1, &pl2)
		CheckPlayerCollisions(&pl1, &pl2)
		CheckMovement(&pl1, &pl2)

		//Rendering
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)
		rl.DrawRectangle(0, 350, 800, 100, rl.Orange)
		DrawPlayers(&pl1, &pl2)

		rl.EndDrawing()
	}
}
