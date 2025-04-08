package main

import rl "github.com/gen2brain/raylib-go/raylib"

type HealthBar1 struct {
	Player *Player1
	Width  float32
	Height float32
	Pos    rl.Vector2
	Color  rl.Color
}

type HealthBar2 struct {
	Player *Player2
	Width  float32
	Height float32
	Pos    rl.Vector2
	Color  rl.Color
}

func NewHealthBar1(pl1 *Player1) HealthBar1 {
	return HealthBar1{
		Player: pl1,
		Width:  20 * float32(pl1.Health),
		Height: 20,
		Pos:    rl.NewVector2(10, 10),
		Color:  rl.Red,
	}
}

func NewHealthBar2(pl2 *Player2) HealthBar2 {
	return HealthBar2{
		Player: pl2,
		Width:  20 * float32(pl2.Health),
		Height: 20,
		Pos:    rl.NewVector2(590, 10),
		Color:  rl.Red,
	}
}

func DrawHealthBars(h1 *HealthBar1, h2 *HealthBar2) {
	rl.DrawRectangle(int32(h1.Pos.X), int32(h1.Pos.Y), int32(h1.Width), int32(h1.Height), h1.Color)
	rl.DrawRectangle(int32(h2.Pos.X), int32(h2.Pos.Y), int32(h2.Width), int32(h2.Height), h2.Color)
}
