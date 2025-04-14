package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	forest := Forest{}
	potion := Potion{}

	forest.Generate(420, 69)
	potion.Generate(320, 41)

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		forest.DrawForest()
		rl.EndDrawing()
	}
}
