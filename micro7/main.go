package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(800, 450, "Micro Assignment 7")

	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	var playerX float32 = 400
	var playerY float32 = 225
	var boxSpeed float32 = 100

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		rl.DrawRectangle(int32(playerX), int32(playerY), 50, 50, rl.Black)

		if rl.IsKeyDown(rl.KeyW) && playerY > 0 {
			playerY -= boxSpeed * rl.GetFrameTime()
		}
		if rl.IsKeyDown(rl.KeyS) && playerY+50 < 450 {
			playerY += boxSpeed * rl.GetFrameTime()
		}
		if rl.IsKeyDown(rl.KeyA) && playerX > 0 {
			playerX -= boxSpeed * rl.GetFrameTime()
		}
		if rl.IsKeyDown(rl.KeyD) && playerX+50 < 800 {
			playerX += boxSpeed * rl.GetFrameTime()
		}

		rl.EndDrawing()
	}
}
