package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Texture        rl.Texture2D
	Pos            rl.Vector2
	Vel            rl.Vector2
	Gravity        float32
	Size           float32
	IsGrounded     bool
	Attacking      bool
	AttackTimer    float32
	HasDealtDamage bool
	Blocking       bool
	Health         int32
	Speed          float32
	Direction      float32
	Color          rl.Color
	AnimationFSM
}

func NewPlayer(p rl.Vector2, d float32, c rl.Color) Player {
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
		Color:        c,
	}
}

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

	if p.Attacking {
		p.AttackTimer -= rl.GetFrameTime()
		if p.AttackTimer <= 0 {
			p.Attacking = false
			p.HasDealtDamage = false
		}
		p.ChangeAnimationState("attack")
		return
	}

	if p.Blocking {
		p.ChangeAnimationState("block")
		return
	}

	if !p.IsGrounded {
		p.ChangeAnimationState("jump")
	} else if p.Vel.X == 0 {
		p.ChangeAnimationState("idle")
	} else {
		p.ChangeAnimationState("walk")
	}
}

func CheckDamage(pl1 *Player, pl2 *Player, h1 *HealthBar1, h2 *HealthBar2) {
	const attackRange = 60.0

	// Player 1 attacking Player 2
	if pl1.Attacking && !pl1.HasDealtDamage {
		inFront := (pl2.Pos.X-pl1.Pos.X)*pl1.Direction > 0
		inRange := rl.Vector2Distance(pl1.Pos, pl2.Pos) <= attackRange

		if inFront && inRange {
			if !pl2.Blocking {
				pl2.Health -= 1
				h2.Pos.X += 20
				h2.Width -= 20
				fmt.Println("Player 2 Health:", pl2.Health)
			} else {
				fmt.Println("Player 2 blocked the attack!")
			}
			pl1.HasDealtDamage = true
		}
	}

	// Player 2 attacking Player 1
	if pl2.Attacking && !pl2.HasDealtDamage {
		inFront := (pl1.Pos.X-pl2.Pos.X)*pl2.Direction > 0
		inRange := rl.Vector2Distance(pl1.Pos, pl2.Pos) <= attackRange

		if inFront && inRange {
			if !pl1.Blocking {
				pl1.Health -= 1
				h1.Width -= 20
				fmt.Println("Player 1 Health:", pl1.Health)
			} else {
				fmt.Println("Player 1 blocked the attack!")
			}
			pl2.HasDealtDamage = true
		}
	}
}

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
	if rl.IsKeyPressed(rl.KeyG) && !p.Attacking {
		p.Attacking = true
		p.AttackTimer = 0.4
		p.HasDealtDamage = false
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
	if rl.IsKeyPressed(rl.KeyL) && !p.Attacking {
		p.Attacking = true
		p.AttackTimer = 0.4
		p.HasDealtDamage = false
	}

	p.Move(float32(playerVelX))
}
