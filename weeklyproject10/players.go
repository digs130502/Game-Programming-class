package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Texture    rl.Texture2D
	Pos        rl.Vector2
	Vel        rl.Vector2
	Gravity    float32
	Size       float32
	IsGrounded bool
	Attacking  bool
	Blocking   bool
	Health     int32
	Speed      float32
	Direction  float32
	AnimationFSM
}

func NewPlayer(p rl.Vector2, d float32) Player {
	newAni := NewAnimationFSM()
	return Player{
		Pos:          p,
		Vel:          rl.NewVector2(0, 0),
		Gravity:      70,
		Size:         4,
		Health:       10,
		Blocking:     false,
		Attacking:    false,
		Speed:        200,
		Direction:    d,
		AnimationFSM: newAni,
	}
}

// func DrawPlayers(pl1 *Player1, pl2 *Player2) {
// 	// Player 1
// 	frame1 := pl1.FrameRec
// 	if !pl1.IsFacingRight {
// 		frame1.Width *= -1
// 	}
// 	rl.DrawTexturePro(
// 		pl1.Texture,
// 		frame1,
// 		rl.NewRectangle(pl1.Pos.X, pl1.Pos.Y, frame1.Width*pl1.Scale, frame1.Height*pl1.Scale),
// 		rl.Vector2Scale(rl.NewVector2(float32(pl1.Texture.Width)/2, float32(pl1.Texture.Height)/2), pl1.Scale),
// 		pl1.Rotation,
// 		rl.White,
// 	)

// 	// Player 2
// 	frame2 := pl2.FrameRec
// 	if !pl2.IsFacingRight {
// 		frame2.Width *= -1
// 	}
// 	rl.DrawTexturePro(
// 		pl2.Texture,
// 		frame2,
// 		rl.NewRectangle(pl2.Pos.X, pl2.Pos.Y, frame2.Width*pl2.Scale, frame2.Height*pl2.Scale),
// 		rl.Vector2Scale(rl.NewVector2(float32(pl2.Texture.Width)/2, float32(pl2.Texture.Height)/2), pl2.Scale),
// 		pl2.Rotation,
// 		rl.Red,
// 	)
// }

func UpdatePlayer(p *Player) {
	p.Pos = rl.Vector2Add(p.Pos, rl.Vector2Scale(p.Vel, rl.GetFrameTime()))

	p.Vel.Y += p.Gravity * rl.GetFrameTime()
	p.Pos.Y += p.Vel.Y

	if p.Pos.Y >= 285 {
		p.Pos.Y = 285
		p.Vel.Y = 0
		p.IsGrounded = true
	} else {
		p.IsGrounded = false
	}

	if !p.IsGrounded {
		p.ChangeAnimationState("jump")
	} else if p.Vel.X == 0 {
		p.ChangeAnimationState("idle")
	} else {
		p.ChangeAnimationState("walk")
	}
}

// func CheckDamage(pl1 *Player1, pl2 *Player2, h1 *HealthBar1, h2 *HealthBar2) {
// 	const attackRange = 100.0

// 	// Player 1 attacking Player 2
// 	if pl1.Attacking {
// 		direction := float32(1)
// 		if !pl1.IsFacingRight {
// 			direction = -1
// 		}
// 		if (pl2.Pos.X-pl1.Pos.X)*direction > 0 && rl.Vector2Distance(pl1.Pos, pl2.Pos) <= attackRange {
// 			if !pl2.Blocking {
// 				pl2.Health -= 1
// 				h2.Pos.X += 20
// 				h2.Width -= 20
// 				fmt.Println("Player 2 Health:", pl2.Health)
// 			} else {
// 				fmt.Println("Player 2 blocked the attack!")
// 			}
// 		}
// 	}

// 	// Player 2 attacking Player 1
// 	if pl2.Attacking {
// 		direction := float32(1)
// 		if !pl2.IsFacingRight {
// 			direction = -1
// 		}
// 		if (pl1.Pos.X-pl2.Pos.X)*direction > 0 && rl.Vector2Distance(pl1.Pos, pl2.Pos) <= attackRange {
// 			if !pl1.Blocking {
// 				pl1.Health -= 1
// 				h1.Width -= 20
// 				fmt.Println("Player 1 Health:", pl1.Health)
// 			} else {
// 				fmt.Println("Player 1 blocked the attack!")
// 			}
// 		}
// 	}
// }

func CheckMovement1(p *Player) {

	// Player 1 Movement
	playerVelX := 0.0
	if rl.IsKeyDown(rl.KeyD) {
		playerVelX = 1
	}
	if rl.IsKeyDown(rl.KeyA) {
		playerVelX = -1

	}
	if rl.IsKeyPressed(rl.KeyW) && p.IsGrounded {
		p.Jump()
	}
	if rl.IsKeyDown(rl.KeyF) {
		p.Blocking = true
	} else {
		p.Blocking = false
	}
	if rl.IsKeyPressed(rl.KeyG) {
		p.Attacking = true
	} else {
		p.Attacking = false
	}
	p.Move(float32(playerVelX))
}

func CheckMovement2(p *Player) {

	//Player 2 Movement
	playerVelX := 0.0
	if rl.IsKeyDown(rl.KeyRight) {
		playerVelX = 1
	}
	if rl.IsKeyDown(rl.KeyLeft) {
		playerVelX = -1
	}
	if rl.IsKeyPressed(rl.KeyUp) && p.IsGrounded {
		p.Jump()
	}
	if rl.IsKeyDown(rl.KeyK) {
		p.Blocking = true
	} else {
		p.Blocking = false
	}
	if rl.IsKeyPressed(rl.KeyL) {
		p.Attacking = true
	} else {
		p.Attacking = false
	}
	p.Move(float32(playerVelX))
}
