package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Player struct {
	Pos    rl.Vector2
	Width  float32
	Height float32
	Color  rl.Color
}

func NewPlayer() Player {
	return Player{
		Pos:    rl.NewVector2(float32(rl.GetScreenWidth())/2-15, float32(rl.GetScreenHeight())-45),
		Width:  30,
		Height: 30,
		Color:  rl.Red,
	}
}

func (p *Player) DrawPlayer() {
	rl.DrawRectangle(int32(p.Pos.X), int32(p.Pos.Y), int32(p.Width), int32(p.Height), p.Color)
}

// FIXME:
func CheckInput(b *Buttons) {
	if rl.IsKeyPressed(rl.KeyEscape) {
		b.DrawPauseButtons()
	}
}
