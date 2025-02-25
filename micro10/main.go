package main

import (
	"fmt"
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type PhysicsBody struct {
	Pos     rl.Vector2
	Vel     rl.Vector2
	Gravity rl.Vector2
}

func NewPhysicsBody() PhysicsBody {
	pb := PhysicsBody{Pos: rl.NewVector2(100, 200), Vel: rl.NewVector2(0, 0), Gravity: rl.NewVector2(0, 500)}
	return pb
}

type Player struct {
	PhysicsBody //embedding in player
}

func NewPlayer() Player {
	np := Player{}
	np.PhysicsBody = NewPhysicsBody()
	return np
}

func (pb *PhysicsBody) VelocityTick() {
	adjustedVel := rl.Vector2Scale(pb.Vel, rl.GetFrameTime()) //scales the velocity to frame time ok ok
	pb.Pos = rl.Vector2Add(pb.Pos, adjustedVel)               //Adds velocity to position effectively changing the position
}

func (pb *PhysicsBody) GravityTick() {
	adjustedGravity := rl.Vector2Scale(pb.Gravity, rl.GetFrameTime())
	pb.Vel = rl.Vector2Add(pb.Vel, adjustedGravity)
}

func (p *Player) PhysicsUpdate() {
	p.GravityTick()
	p.VelocityTick()
}

func (p Player) DrawPlayer() {
	rl.DrawRectangle(int32(p.Pos.X), int32(p.Pos.Y), 50, 50, rl.Orange)
}

func main() {
	rl.InitWindow(800, 450, "Flappy Bird")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	rand.Seed(time.Now().UnixNano())

	//var playerY float32 = 200
	//var playerSpeed float32 = 100
	var pipeX float32 = 750
	var pipeSpeed float32 = 150
	var gap float32 = 130
	var points int = 0
	var gotPoint bool = false
	var pipeHeight float32 = float32(rand.Intn(250) + 50)

	player := NewPlayer()

	//DONE: Player Movement. Pipe spawning, Gap in between the pipes, score system.
	//To DO: Hitting pipe. Pressing "R" to restart the game

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		rl.DrawRectangle(int32(pipeX), 0, 70, int32(pipeHeight), rl.Green)
		rl.DrawRectangle(int32(pipeX), int32(pipeHeight+gap), 70, 450-int32(pipeHeight+gap), rl.Green)

		//Pipe movement
		pipeX -= pipeSpeed * rl.GetFrameTime()
		player.PhysicsUpdate()
		player.DrawPlayer()

		//Player movement
		if rl.IsKeyPressed(rl.KeySpace) {
			player.Vel = rl.NewVector2(0, -250)
		}

		// Collision mechanic
		playerRect := rl.Rectangle{X: float32(player.Pos.X), Y: float32(player.Pos.Y), Width: 50, Height: 50}
		topPipe := rl.Rectangle{X: pipeX, Y: 0, Width: 70, Height: pipeHeight}
		bottomPipe := rl.Rectangle{X: pipeX, Y: pipeHeight + gap, Width: 70, Height: 450 - (pipeHeight + gap)}

		if rl.CheckCollisionRecs(playerRect, topPipe) || rl.CheckCollisionRecs(playerRect, bottomPipe) {
			pipeSpeed = 0
			player.Vel = rl.NewVector2(0, 0) // Stop player movement
			player.Gravity = rl.NewVector2(0, 0)

			rl.DrawText("Game Over!", 340, 150, 20, rl.Black)
			rl.DrawText("Press R to restart the game", 250, 180, 20, rl.Black)
			rl.DrawText("Press Q to quit the game", 270, 210, 20, rl.Black)

			if rl.IsKeyDown(rl.KeyR) {
				// Reset game state
				pipeX = 750
				pipeSpeed = 150
				player.Pos = rl.NewVector2(100, 200)
				player.Vel = rl.NewVector2(0, 0)
				player.Gravity = rl.NewVector2(0, 500)
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
