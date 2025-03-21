package main

import (
	"fmt"
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 450, "Flappy Bird")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	rand.Seed(time.Now().UnixNano())

	var playerY float32 = 200
	var playerSpeed float32 = 100
	var pipeX float32 = 750
	var pipeSpeed float32 = 150
	var gap float32 = 80
	var points int = 0
	var gotPoint bool = false
	var pipeHeight float32 = float32(rand.Intn(250) + 50)

	//DONE: Player Movement. Pipe spawning, Gap in between the pipes, score system.
	//To DO: Hitting pipe. Pressing "R" to restart the game

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		rl.DrawRectangle(100, int32(playerY), 50, 50, rl.Orange)
		rl.DrawRectangle(int32(pipeX), 0, 70, int32(pipeHeight), rl.Green)
		rl.DrawRectangle(int32(pipeX), int32(pipeHeight+gap), 70, 450-int32(pipeHeight+gap), rl.Green)

		//Pipe movement
		pipeX -= pipeSpeed * rl.GetFrameTime()

		//Player movement
		if rl.IsKeyDown(rl.KeyW) && playerY > 0 {
			playerY -= playerSpeed * rl.GetFrameTime()
		}
		if rl.IsKeyDown(rl.KeyS) && playerY+50 < 450 {
			playerY += playerSpeed * rl.GetFrameTime()
		}

		//Collision mechanic
		player := rl.Rectangle{X: 100, Y: playerY, Width: 50, Height: 50}
		topPipe := rl.Rectangle{X: pipeX, Y: 0, Width: 70, Height: pipeHeight}
		bottomPipe := rl.Rectangle{X: pipeX, Y: pipeHeight + gap, Width: 70, Height: 450 - (pipeHeight + gap)}

		if rl.CheckCollisionRecs(player, topPipe) || rl.CheckCollisionRecs(player, bottomPipe) {
			playerSpeed = 0
			pipeSpeed = 0
			rl.DrawText("Game Over!", 340, 150, 20, rl.Black)
			rl.DrawText("Press R to restart the game", 250, 180, 20, rl.Black)
			rl.DrawText("Press Q to quit the game", 270, 210, 20, rl.Black)

			if rl.IsKeyDown(rl.KeyR) {
				pipeX = 750
				pipeSpeed = 150
				playerSpeed = 100
				playerY = 200
				points = 0
			}

		}

		//Pipe respawing
		if pipeX <= -50 {
			pipeX = 750
			pipeHeight = float32(rand.Intn(250) + 50)
			gotPoint = false
		}

		//Score system
		if pipeX+70 < 100 && !gotPoint {
			points += 1
			gotPoint = true
		}
		pointsText := fmt.Sprintf("Points: %d", points)
		rl.DrawText(pointsText, 5, 5, 20, rl.Black)

		//Quit game system
		if rl.IsKeyDown(rl.KeyQ) {
			break
		}

		rl.EndDrawing()
	}

}
