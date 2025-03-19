package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Player struct {
	X      float32
	Y      float32
	Height float32
	Width  float32
}

func NewPlayer() Player {
	return Player{X: 300, Y: 200, Height: 25, Width: 25}
}

func (p *Player) DrawPlayer() {
	rl.DrawRectangle(int32(p.X), int32(p.Y), int32(p.Width), int32(p.Height), rl.Beige)
}

// Player Movement
func (p *Player) MovePlayer() {
	if rl.IsKeyDown(rl.KeyA) {
		p.X -= 100 * rl.GetFrameTime()
	}
	if rl.IsKeyDown(rl.KeyD) {
		p.X += 100 * rl.GetFrameTime()
	}
	if rl.IsKeyDown(rl.KeyW) {
		p.Y -= 100 * rl.GetFrameTime()
	}
	if rl.IsKeyDown(rl.KeyS) {
		p.Y += 100 * rl.GetFrameTime()
	}
}
