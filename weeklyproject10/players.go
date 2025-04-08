package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player1 struct {
	Texture      rl.Texture2D
	Pos          rl.Vector2
	Vel          rl.Vector2
	Gravity      float32
	Scale        float32
	Rotation     float32
	FrameRec     rl.Rectangle
	CurrentFrame int
	IsOnGround   bool
	Attacking    bool
	Blocking     bool
	Health       int32
}

type Player2 struct {
	Texture      rl.Texture2D
	Pos          rl.Vector2
	Vel          rl.Vector2
	Gravity      float32
	Scale        float32
	Rotation     float32
	FrameRec     rl.Rectangle
	CurrentFrame int
	IsOnGround   bool
	Attacking    bool
	Blocking     bool
	Health       int32
}

func NewPlayer1() Player1 {
	playerSpr := rl.LoadTexture("assets/sprites/player.png")
	return Player1{
		Texture:   playerSpr,
		Pos:       rl.NewVector2(50, 290),
		Vel:       rl.NewVector2(0, 0),
		Gravity:   70,
		Scale:     3,
		Rotation:  0,
		FrameRec:  rl.NewRectangle(0, 0, float32(playerSpr.Width), float32(playerSpr.Height)),
		Health:    10,
		Blocking:  false,
		Attacking: false,
	}
}

func NewPlayer2() Player2 {
	playerSpr := rl.LoadTexture("assets/sprites/player2.png")
	return Player2{
		Texture:   playerSpr,
		Pos:       rl.NewVector2(750, 290),
		Vel:       rl.NewVector2(0, 0),
		Gravity:   70,
		Scale:     3,
		Rotation:  0,
		FrameRec:  rl.NewRectangle(0, 0, float32(playerSpr.Width), float32(playerSpr.Height)),
		Health:    10,
		Blocking:  false,
		Attacking: false,
	}
}

func DrawPlayers(pl1 *Player1, pl2 *Player2) {
	rl.DrawTexturePro(
		pl1.Texture,
		pl1.FrameRec,
		rl.NewRectangle(pl1.Pos.X, pl1.Pos.Y, pl1.FrameRec.Width*pl1.Scale, pl1.FrameRec.Height*pl1.Scale),
		rl.Vector2Scale(rl.NewVector2(float32(pl1.Texture.Width)/2, float32(pl1.Texture.Height)/2), pl1.Scale),
		pl1.Rotation,
		rl.White,
	)
	rl.DrawTexturePro(
		pl2.Texture,
		pl2.FrameRec,
		rl.NewRectangle(pl2.Pos.X, pl2.Pos.Y, pl2.FrameRec.Width*pl2.Scale, pl2.FrameRec.Height*pl2.Scale),
		rl.Vector2Scale(rl.NewVector2(float32(pl2.Texture.Width)/2, float32(pl2.Texture.Height)/2), pl2.Scale),
		pl2.Rotation,
		rl.White,
	)
}

func UpdatePlayers(pl1 *Player1, pl2 *Player2) {
	// Player 1 gravity and movement
	pl1.Vel.Y += pl1.Gravity * rl.GetFrameTime()
	pl1.Pos.Y += pl1.Vel.Y

	// Player 2 gravity and movement
	pl2.Vel.Y += pl2.Gravity * rl.GetFrameTime()
	pl2.Pos.Y += pl2.Vel.Y
}

func CheckPlayerFloorCollisions(pl1 *Player1, pl2 *Player2) {
	groundY := float32(350)

	// Player 1
	playerHeight := pl1.FrameRec.Height * pl1.Scale
	playerBottom := pl1.Pos.Y + (playerHeight / 2)
	if playerBottom >= groundY {
		pl1.Pos.Y = groundY - (playerHeight / 2)
		pl1.Vel.Y = 0
		pl1.IsOnGround = true
	} else {
		pl1.IsOnGround = false
	}

	// Player 2
	playerHeight2 := pl2.FrameRec.Height * pl2.Scale
	playerBottom2 := pl2.Pos.Y + (playerHeight2 / 2)
	if playerBottom2 >= groundY {
		pl2.Pos.Y = groundY - (playerHeight2 / 2)
		pl2.Vel.Y = 0
		pl2.IsOnGround = true
	} else {
		pl2.IsOnGround = false
	}
}

func CheckDamage(pl1 *Player1, pl2 *Player2, h1 *HealthBar1, h2 *HealthBar2) {
	const attackRange = 100.0
	// fmt.Println("Distance:", rl.Vector2Distance(pl1.Pos, pl2.Pos))

	// Player 1 attacking Player 2
	if pl1.Attacking && rl.Vector2Distance(pl1.Pos, pl2.Pos) <= attackRange {
		if !pl2.Blocking {
			pl2.Health -= 1
			h2.Pos.X += 20
			h2.Width -= 20
			fmt.Println("Player 2 Health:", pl2.Health)
		} else {
			fmt.Println("Player 2 blocked the attack!")
		}
	}

	// Player 2 attacking Player 1
	if pl2.Attacking && rl.Vector2Distance(pl1.Pos, pl2.Pos) <= attackRange {
		if !pl1.Blocking {
			pl1.Health -= 1
			h1.Width -= 20
			fmt.Println("Player 1 Health:", pl1.Health)
		} else {
			fmt.Println("Player 1 blocked the attack!")
		}
	}
}

func CheckMovement(pl1 *Player1, pl2 *Player2) {

	// Player 1 Movement
	if rl.IsKeyDown(rl.KeyD) {
		pl1.Pos.X += 300 * rl.GetFrameTime()
	}
	if rl.IsKeyDown(rl.KeyA) {
		pl1.Pos.X -= 300 * rl.GetFrameTime()
	}
	if rl.IsKeyPressed(rl.KeyW) && pl1.IsOnGround {
		pl1.Vel.Y = -25
	}
	if rl.IsKeyDown(rl.KeyF) {
		pl1.Blocking = true
	} else {
		pl1.Blocking = false
	}
	if rl.IsKeyPressed(rl.KeyG) {
		pl1.Attacking = true
	} else {
		pl1.Attacking = false
	}

	// Player 2 Movement
	if rl.IsKeyDown(rl.KeyRight) {
		pl2.Pos.X += 300 * rl.GetFrameTime()
	}
	if rl.IsKeyDown(rl.KeyLeft) {
		pl2.Pos.X -= 300 * rl.GetFrameTime()
	}
	if rl.IsKeyPressed(rl.KeyUp) && pl2.IsOnGround {
		pl2.Vel.Y = -25
	}
	if rl.IsKeyDown(rl.KeyK) {
		pl2.Blocking = true
	} else {
		pl2.Blocking = false
	}
	if rl.IsKeyPressed(rl.KeyL) {
		pl2.Attacking = true
	} else {
		pl2.Attacking = false
	}

}
