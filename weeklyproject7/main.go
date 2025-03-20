package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 450, "Space Defense")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	//Initialize player
	player := NewPlayer()

	//Initialize planet
	planet := NewPlanet()

	//Initialize audio
	rl.InitAudioDevice()
	spaceMusic := rl.LoadMusicStream("assets/audio/space.mp3") //Load music Space
	rl.SetMusicVolume(spaceMusic, .05)                         //Set Volume
	rl.PlayMusicStream(spaceMusic)

	for !rl.WindowShouldClose() {
		rl.UpdateMusicStream(spaceMusic)

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		// Updates
		player.MovePlayer()
		player.UpdateProjectiles()

		// Rendering
		planet.DrawPlanet()
		player.DrawPlayer()
		player.DrawProjectiles()

		//Planet health
		planetText := fmt.Sprintf("Planet Health: %d", planet.Health)
		rl.DrawText(planetText, 10, 10, 18, rl.White)

		rl.EndDrawing()
	}

	rl.UnloadMusicStream(spaceMusic)

}
