package main

import (
	"fmt"
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 450, "Space Defense")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	rand.Seed(time.Now().UnixNano())

	//Initialize player
	player := NewPlayer()

	//Initialize planet
	planet := NewPlanet()

	//Initialize zone
	zone := NewZone()

	//Initialize audio
	rl.InitAudioDevice()
	spaceMusic := rl.LoadMusicStream("assets/audio/space.mp3") //Load music Space
	rl.SetMusicVolume(spaceMusic, .05)                         //Set Volume
	rl.PlayMusicStream(spaceMusic)

	//Initialize Asteroids
	zone.NewAsteroid()
	zone.NewAsteroid()

	for !rl.WindowShouldClose() {
		rl.UpdateMusicStream(spaceMusic)

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		// Updates
		player.MovePlayer()
		player.UpdateProjectiles()
		zone.UpdateAsteroids()
		zone.CheckAsteroidCollision(&planet)

		// Rendering
		planet.DrawPlanet()
		player.DrawPlayer()
		player.DrawProjectiles()
		zone.DrawAsteroid()

		//Planet health
		planetText := fmt.Sprintf("Planet Health: %d", planet.Health)
		rl.DrawText(planetText, 10, 10, 18, rl.White)

		rl.EndDrawing()
	}

	rl.UnloadMusicStream(spaceMusic)

}
