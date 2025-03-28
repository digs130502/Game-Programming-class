package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Zone struct {
	Mines    []Mine
	Count    int32
	GameOver bool
}

type PlayerCircle struct {
	Pos    rl.Vector2
	Radius float32
	Color  rl.Color
}

type Mine struct {
	Pos    rl.Vector2
	Radius float32
	Color  rl.Color
}

func NewZone() Zone {
	return Zone{
		Mines:    []Mine{},
		Count:    0,
		GameOver: false,
	}
}

func NewPlayerCircle() PlayerCircle {
	return PlayerCircle{
		Pos:    rl.GetMousePosition(),
		Radius: 50,
		Color:  rl.Blue,
	}
}

func NewMine(pc PlayerCircle) Mine {
	return Mine{
		Pos:    pc.Pos,
		Radius: 25,
		Color:  rl.Black,
	}
}

func (pc *PlayerCircle) DrawPlayerCircle() {
	rl.DrawCircle(int32(pc.Pos.X), int32(pc.Pos.Y), pc.Radius, pc.Color)
}

func (pc *PlayerCircle) UpdateCircle() {
	adjustedPosition := rl.GetMousePosition()
	pc.Pos = adjustedPosition
}

func (z *Zone) CheckMineCreation(pc PlayerCircle) {
	if rl.IsKeyPressed(rl.KeySpace) {
		z.Mines = append(z.Mines, NewMine(pc))
		z.Count += 1
	}
}

func (z *Zone) DrawMines() {
	for _, m := range z.Mines {
		rl.DrawCircle(int32(m.Pos.X), int32(m.Pos.Y), m.Radius, m.Color)
	}
}

func (z *Zone) CheckPlayerCollisions(pc PlayerCircle) {
	for i := 0; i < len(z.Mines); i++ {
		if rl.Vector2Distance(z.Mines[i].Pos, pc.Pos) > z.Mines[i].Radius+pc.Radius {
			z.Mines[i].Color = rl.Red
		}
		if (rl.Vector2Distance(z.Mines[i].Pos, pc.Pos) <= z.Mines[i].Radius+pc.Radius) && z.Mines[i].Color == rl.Red {
			z.GameOver = true
		}
	}
}
