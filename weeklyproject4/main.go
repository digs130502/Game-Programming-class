package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(800, 450, "Flappy Bird")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	var playerY float32 = 200
	var speed float32 = 100
	var pipeX float32 = 750
	var pipeSpeed float32 = 100

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		rl.DrawRectangle(100, int32(playerY), 50, 50, rl.Orange)
		rl.DrawRectangle(int32(pipeX), 0, 70, 450, rl.Green)
		rl.DrawText("Points: ", 5, 5, 20, rl.Black)

		pipeX -= pipeSpeed * rl.GetFrameTime()

		if rl.IsKeyDown(rl.KeyW) && playerY > 0 {
			playerY -= speed * rl.GetFrameTime()
		}
		if rl.IsKeyDown(rl.KeyS) && playerY+50 < 450 {
			playerY += speed * rl.GetFrameTime()
		}

		if rl.IsKeyDown(rl.KeyR) {
			break
		}

		if pipeX <= -50 {
			pipeX = 750
		}

		rl.EndDrawing()
	}

}
