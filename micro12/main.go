package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 450, "Micro Assignment 12")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	rl.InitAudioDevice()

	birds := rl.LoadSound("audio/birds.wav")
	piano := rl.LoadSound("audio/piano.wav")
	woodpecker := rl.LoadSound("audio/woodpecker.wav")
	whoosh := rl.LoadSound("audio/whoosh.wav")
	music := rl.LoadSound("audio/music.wav")

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		rl.DrawRectangle(275, 200, 50, 50, rl.Orange)
		rl.DrawRectangle(335, 200, 50, 50, rl.Orange)
		rl.DrawRectangle(395, 200, 50, 50, rl.Orange)
		rl.DrawRectangle(455, 200, 50, 50, rl.Orange)
		rl.DrawRectangle(515, 200, 50, 50, rl.Orange)

		rl.DrawText("Press keys 1-5 for sounds!", 275, 300, 20, rl.RayWhite)

		if rl.IsKeyPressed(rl.KeyOne) {
			rl.PlaySound(birds)
		}
		if rl.IsKeyPressed(rl.KeyTwo) {
			rl.PlaySound(piano)
		}
		if rl.IsKeyPressed(rl.KeyThree) {
			rl.PlaySound(woodpecker)
		}
		if rl.IsKeyPressed(rl.KeyFour) {
			rl.PlaySound(whoosh)
		}
		if rl.IsKeyPressed(rl.KeyFive) {
			rl.PlaySound(music)
		}

		rl.EndDrawing()
	}
}
