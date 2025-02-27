package main

import (
	"fmt"
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	TileSize = 50
	Rows     = 9
	Cols     = 16
)

// Enemy struct to store position and points
type Enemy struct {
	Position rl.Vector2
	Points   int
}

// Generates a random grid position that is not occupied
func getUniqueGridPosition(existingPositions []rl.Vector2) rl.Vector2 {
	for {
		newPos := rl.NewVector2(float32(rand.Intn(Cols)*TileSize), float32(rand.Intn(Rows)*TileSize))

		// Check if newPos collides with any existing position
		isUnique := true
		for _, pos := range existingPositions {
			if pos.X == newPos.X && pos.Y == newPos.Y {
				isUnique = false
				break
			}
		}

		// If unique, return it
		if isUnique {
			return newPos
		}
	}
}

func main() {
	rl.InitWindow(800, 450, "Turn Based Hero")
	defer rl.CloseWindow()

	rand.Seed(time.Now().UnixNano())

	rl.SetTargetFPS(60)

	playerTex := rl.LoadTexture("textures/slimeski.png")
	enemyTex := rl.LoadTexture("textures/slimeski.png")

	// Initialize player points
	points := 1
	pointsText := fmt.Sprintf("%d", points)

	// Track occupied positions to prevent overlap
	occupiedPositions := []rl.Vector2{}

	// Generate player position (ensuring it doesn't collide with enemies)
	playerPos := getUniqueGridPosition(occupiedPositions)
	occupiedPositions = append(occupiedPositions, playerPos)

	// Generate enemies at unique grid positions
	enemies := []Enemy{}
	for i := 0; i < 5; i++ {
		enemyPos := getUniqueGridPosition(occupiedPositions)
		occupiedPositions = append(occupiedPositions, enemyPos) // Mark as occupied

		enemies = append(enemies, Enemy{
			Position: enemyPos,
			Points:   i + 1,
		})
	}

	// Game over state
	gameOver := false
	gameWon := false

	for !rl.WindowShouldClose() {
		// If the game is over, disable movement
		if !gameOver {
			// Player movement
			if rl.IsKeyPressed(rl.KeyW) {
				playerPos.Y -= TileSize
			}
			if rl.IsKeyPressed(rl.KeyS) {
				playerPos.Y += TileSize
			}
			if rl.IsKeyPressed(rl.KeyA) {
				playerPos.X -= TileSize
			}
			if rl.IsKeyPressed(rl.KeyD) {
				playerPos.X += TileSize
			}

			// Absorb enemies
			for i := len(enemies) - 1; i >= 0; i-- {
				if playerPos.X == enemies[i].Position.X && playerPos.Y == enemies[i].Position.Y {
					if points >= enemies[i].Points {
						points += enemies[i].Points
						pointsText = fmt.Sprintf("%d", points)
						enemies = append(enemies[:i], enemies[i+1:]...)
					} else {
						gameOver = true
					}
				}
			}
			if len(enemies) == 0 {
				gameWon = true
			}
		}

		// Draw everything
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		// Draw player if game is not over
		if !gameOver {
			rl.DrawTextureEx(playerTex, playerPos, 0, 2, rl.Green)
			rl.DrawText(pointsText, int32(playerPos.X), int32(playerPos.Y), 20, rl.Black)
		}

		// Draw enemies
		for _, enemy := range enemies {
			rl.DrawTextureEx(enemyTex, enemy.Position, 0, 2, rl.Red)
			rl.DrawText(fmt.Sprintf("%d", enemy.Points), int32(enemy.Position.X), int32(enemy.Position.Y), 20, rl.Black)
		}

		if rl.IsKeyPressed(rl.KeyQ) {
			break
		}

		if gameOver {
			rl.DrawText("YOU LOST", 325, 180, 30, rl.Black)
			rl.DrawText("Press R to restart", 280, 230, 20, rl.Black)
			rl.DrawText("Press Q to quit", 300, 150, 20, rl.Black)

			// Restart game when R is pressed
			if rl.IsKeyPressed(rl.KeyR) {
				// Reset game state
				points = 1
				pointsText = fmt.Sprintf("%d", points)
				occupiedPositions = []rl.Vector2{}
				playerPos = getUniqueGridPosition(occupiedPositions)
				occupiedPositions = append(occupiedPositions, playerPos)

				enemies = []Enemy{}
				for i := 0; i < 5; i++ {
					enemyPos := getUniqueGridPosition(occupiedPositions)
					occupiedPositions = append(occupiedPositions, enemyPos)

					enemies = append(enemies, Enemy{
						Position: enemyPos,
						Points:   i + 1,
					})
				}
				gameOver = false
			}
		}

		if gameWon {
			rl.DrawText("YOU WON", 325, 180, 30, rl.Black)
			rl.DrawText("Press R to restart", 300, 230, 20, rl.Black)
			rl.DrawText("Press Q to quit", 300, 150, 20, rl.Black)

			// Restart game when R is pressed
			if rl.IsKeyPressed(rl.KeyR) {
				// Reset game state
				points = 1
				pointsText = fmt.Sprintf("%d", points)
				occupiedPositions = []rl.Vector2{}
				playerPos = getUniqueGridPosition(occupiedPositions)
				occupiedPositions = append(occupiedPositions, playerPos)

				enemies = []Enemy{}
				for i := 0; i < 5; i++ {
					enemyPos := getUniqueGridPosition(occupiedPositions)
					occupiedPositions = append(occupiedPositions, enemyPos)

					enemies = append(enemies, Enemy{
						Position: enemyPos,
						Points:   i + 1,
					})
				}
				gameWon = false
			}
		}

		rl.EndDrawing()
	}
}
