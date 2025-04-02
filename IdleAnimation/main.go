package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	idleSheet := rl.LoadTexture("sprites/idle.png")
	animation := NewAnimation(rl.NewVector2(200, 200), idleSheet, 4, .15)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		animation.UpdateTime()
		animation.DrawAnimation()

		rl.EndDrawing()
	}
}
