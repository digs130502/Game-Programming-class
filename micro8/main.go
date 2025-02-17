package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(800, 450, "Micro Assignment 8")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	slimeBro := rl.LoadTexture("textures/slimeski.png")
	custom1 := rl.NewColor(100, 234, 100, 255)
	custom2 := rl.NewColor(5, 100, 234, 255)
	custom3 := rl.NewColor(200, 10, 255, 255)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		rl.DrawTextureEx(slimeBro, rl.NewVector2(200, 200), 20, 5, custom1)
		rl.DrawTextureEx(slimeBro, rl.NewVector2(100, 100), -45, 2, custom2)
		rl.DrawTextureEx(slimeBro, rl.NewVector2(650, 0), 90, 4, custom3)

		rl.EndDrawing()
	}
}
