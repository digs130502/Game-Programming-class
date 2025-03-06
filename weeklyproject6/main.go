package main

import rl "github.com/gen2brain/raylib-go/raylib"

type PhysicsBody struct {
	Pos              rl.Vector2
	Vel              rl.Vector2
	IgnoreCollisions bool
}

func main() {
	rl.InitWindow(800, 450, "Breakout")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		rl.EndDrawing()
	}
}
