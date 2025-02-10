package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(800, 450, "raylib [core] exapmle - basic window")

	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		rl.DrawRectangle(0, 0, 50, 50, rl.Black)
		rl.DrawRectangle(375, 200, 50, 50, rl.Green)
		rl.DrawRectangle(750, 0, 50, 50, rl.Blue)
		rl.DrawRectangle(750, 400, 50, 50, rl.Pink)
		rl.DrawRectangle(0, 400, 50, 50, rl.Orange)

		rl.EndDrawing()
	}
}
