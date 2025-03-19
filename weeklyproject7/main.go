package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(800, 450, "Space Defense")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	//Initialize player
	player := NewPlayer()

	//Initialize audio
	rl.InitAudioDevice()
	spaceMusic := rl.LoadMusicStream("assets/audio/space.mp3") //Load music Space
	rl.SetMusicVolume(spaceMusic, .05)                         //Set Volume
	rl.PlayMusicStream(spaceMusic)

	for !rl.WindowShouldClose() {
		rl.UpdateMusicStream(spaceMusic)

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		// Rendering
		rl.DrawCircle(400, 225, 35, rl.Brown)
		player.DrawPlayer()

		// Movement
		player.MovePlayer()

		rl.EndDrawing()
	}

	rl.UnloadMusicStream(spaceMusic)

}
