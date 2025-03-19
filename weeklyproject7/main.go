package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(800, 450, "Space Defense")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	player := NewPlayer()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		//Rendering
		player.DrawPlayer()
		rl.DrawCircle(400, 225, 35, rl.Brown)

		//Movement
		player.MovePlayer()
		rl.EndDrawing()
	}
}
