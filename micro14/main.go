package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(800, 450, "Micro Assignment 14")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	button := NewButton()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		DrawButton(button)
		button.ChangeColor()
		button.IsClicked()

		rl.EndDrawing()
	}
}
