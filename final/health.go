package main

import rl "github.com/gen2brain/raylib-go/raylib"

type HealthBar struct {
	Player *Player
	Width  float32
	Height float32
	Pos    rl.Vector2
	Color  rl.Color
}

func NewHealthBar(pl *Player) HealthBar {
	return HealthBar{
		Player: pl,
		Width:  20 * float32(pl.Health),
		Height: 20,
		Pos:    rl.NewVector2(10, 10),
		Color:  rl.Red,
	}
}

func (h *HealthBar) DrawHealthBar() {
	rl.DrawRectangle(int32(h.Pos.X), int32(h.Pos.Y), int32(h.Width), int32(h.Height), h.Color)
}
