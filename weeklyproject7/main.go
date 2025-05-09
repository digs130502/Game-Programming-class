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

	// Initialize player
	player := NewPlayer()

	//Initialize planet
	planet := NewPlanet()

	// Initialize zone
	zone := NewZone()

	// Initialize audio
	rl.InitAudioDevice()
	spaceMusic := rl.LoadMusicStream("assets/audio/space.mp3") //Load music Space
	rl.SetMusicVolume(spaceMusic, .05)                         //Set Volume
	rl.PlayMusicStream(spaceMusic)
	explosion := rl.LoadSound("assets/audio/explosion.wav")
	bounce := rl.LoadSound("assets/audio/bounce.wav")
	heal := rl.LoadSound("assets/audio/heal.wav")
	shootSound := rl.LoadSound("assets/audio/shoot.wav")

	// Timer
	spawnTimer := 0
	spawnInterval := 180

	// Initialize Game
	gameRunning := planet.CheckGameOver()

	for !rl.WindowShouldClose() {
		rl.UpdateMusicStream(spaceMusic)

		rl.BeginDrawing()
		if gameRunning {
			rl.ClearBackground(rl.Black)

			// Updates
			player.MovePlayer(shootSound)
			player.UpdateProjectiles()
			player.CheckPlanetCollision(&planet, heal)
			zone.UpdateAsteroids()
			zone.UpdateCargoAsteroids()
			zone.CheckAsteroidCollision(&player, &planet, explosion)
			zone.CheckCargoAsteroidCollision(&player, bounce)
			gameRunning = planet.CheckGameOver()

			//Update timer
			spawnTimer++
			if spawnTimer >= spawnInterval {
				zone.NewAsteroid(20)
				spawnTimer = 0
			}

			// Rendering
			planet.DrawPlanet()
			player.DrawPlayer()
			player.DrawProjectiles()
			zone.DrawAsteroid()
			zone.DrawCargoAsteroid()

			//Planet health
			planetText := fmt.Sprintf("Planet Health: %d", planet.Health)
			rl.DrawText(planetText, 10, 10, 18, rl.White)

			//Player cargo
			cargoText := fmt.Sprintf("Player Cargo: %d", player.Cargo)
			rl.DrawText(cargoText, 10, 30, 18, rl.RayWhite)

		} else {
			rl.ClearBackground(rl.Red)
			rl.DrawText("Game Over! Press R to restart.", 400, 225, 20, rl.RayWhite)
			zone.Asteroids = []Asteroid{}
			if rl.IsKeyPressed(rl.KeyR) {
				gameRunning = true
				planet.Health = 10
				spawnTimer = 0
			}
		}

		rl.EndDrawing()
	}

	rl.UnloadMusicStream(spaceMusic)

}
