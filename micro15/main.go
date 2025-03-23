package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type PhysicsBody struct {
	Pos     rl.Vector2
	Vel     rl.Vector2
	Gravity rl.Vector2
}

func NewPhysicsBody() PhysicsBody {
	return PhysicsBody{
		Pos:     rl.NewVector2(100, 200),
		Vel:     rl.NewVector2(0, 0),
		Gravity: rl.NewVector2(0, 500),
	}
}

type Player struct {
	PhysicsBody
}

func NewPlayer() Player {
	return Player{PhysicsBody: NewPhysicsBody()}
}

func (pb *PhysicsBody) VelocityTick() {
	adjustedVel := rl.Vector2Scale(pb.Vel, rl.GetFrameTime())
	pb.Pos = rl.Vector2Add(pb.Pos, adjustedVel)
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

type GameState struct {
	Player     Player
	PipeX      float32
	PipeHeight float32
	GotPoint   bool
	Points     int
}

func NewGameState() GameState {
	return GameState{
		Player:     NewPlayer(),
		PipeX:      750,
		PipeHeight: float32(rand.Intn(250) + 50),
		GotPoint:   false,
		Points:     0,
	}
}

func (gs *GameState) Save(filename string) error {
	data, err := json.MarshalIndent(gs, "", "  ")
	if err != nil {
		fmt.Println(err)
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func (gs *GameState) Load(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, gs)
}

func main() {
	rl.InitWindow(800, 450, "Flappy Bird")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)
	rand.Seed(time.Now().UnixNano())

	pipeSpeed := float32(150)
	gap := float32(130)

	game := NewGameState()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		// Pipe rendering
		rl.DrawRectangle(int32(game.PipeX), 0, 70, int32(game.PipeHeight), rl.Green)
		rl.DrawRectangle(int32(game.PipeX), int32(game.PipeHeight+gap), 70, 450-int32(game.PipeHeight+gap), rl.Green)

		// Pipe movement
		game.PipeX -= pipeSpeed * rl.GetFrameTime()

		// Player physics and render
		game.Player.PhysicsUpdate()
		game.Player.DrawPlayer()

		if rl.IsKeyPressed(rl.KeySpace) {
			game.Player.Vel = rl.NewVector2(0, -250)
		}

		// Collision detection
		playerRect := rl.Rectangle{X: game.Player.Pos.X, Y: game.Player.Pos.Y, Width: 50, Height: 50}
		topPipe := rl.Rectangle{X: game.PipeX, Y: 0, Width: 70, Height: game.PipeHeight}
		bottomPipe := rl.Rectangle{X: game.PipeX, Y: game.PipeHeight + gap, Width: 70, Height: 450 - (game.PipeHeight + gap)}

		if rl.CheckCollisionRecs(playerRect, topPipe) || rl.CheckCollisionRecs(playerRect, bottomPipe) {
			pipeSpeed = 0
			game.Player.Vel = rl.NewVector2(0, 0)
			game.Player.Gravity = rl.NewVector2(0, 0)

			rl.DrawText("Game Over!", 340, 150, 20, rl.Black)
			rl.DrawText("Press R to restart", 280, 180, 20, rl.Black)
			rl.DrawText("Press Q to quit", 300, 210, 20, rl.Black)

			if rl.IsKeyPressed(rl.KeyR) {
				game = NewGameState()
				pipeSpeed = 150
			}
		}

		// Pipe reset
		if game.PipeX <= -70 {
			game.PipeX = 750
			game.PipeHeight = float32(rand.Intn(250) + 50)
			game.GotPoint = false
		}

		// Score system
		if game.PipeX+70 < 100 && !game.GotPoint {
			game.Points++
			game.GotPoint = true
		}
		pointsText := fmt.Sprintf("Points: %d", game.Points)
		rl.DrawText(pointsText, 5, 5, 20, rl.Black)

		// Save and load system
		if rl.IsKeyPressed(rl.KeyS) {
			game.Save("savefile.json")
		}
		if rl.IsKeyPressed(rl.KeyL) {
			game.Load("savefile.json")
		}

		if rl.IsKeyDown(rl.KeyQ) {
			break
		}

		rl.EndDrawing()
	}
}
