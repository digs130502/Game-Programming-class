package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(800, 450, "Turn Based Hero")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)
	player := rl.LoadTexture("textures/slimeski.png")
	speed := rl.NewVector2(400, 200)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		rl.DrawTextureEx(player, speed, 0, 2, rl.Green)

		rl.EndDrawing()
	}
}
