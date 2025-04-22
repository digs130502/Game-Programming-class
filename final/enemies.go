package main

import (
	"fmt"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Zone struct {
	Enemies []Enemy
	Timer   float32
}

type Enemy struct {
	Pos    rl.Vector2
	Vel    rl.Vector2
	Width  int32
	Height int32
	Word   string
}

func NewEnemy() Enemy {
	return Enemy{
		Pos:    rl.NewVector2(float32(rand.Intn(800)), 0),
		Vel:    rl.NewVector2(0, 100),
		Width:  50,
		Height: 50,
	}
}

func NewZone() Zone {
	return Zone{
		Enemies: []Enemy{},
	}
}

func (z *Zone) CheckEnemyCreation(n float32) {
	z.Timer += rl.GetFrameTime()

	if z.Timer > n {
		z.Enemies = append(z.Enemies, NewEnemy())
		fmt.Println("Enemy Created!")
		z.Timer = 0
	}
}

func (z *Zone) UpdateEnemies() {
	for i := range z.Enemies {
		vel := rl.Vector2Scale(z.Enemies[i].Vel, rl.GetFrameTime())
		z.Enemies[i].Pos = rl.Vector2Add(z.Enemies[i].Pos, vel)
	}
}

func (z *Zone) DrawEnemies() {
	for i := range z.Enemies {
		rl.DrawRectangle(int32(z.Enemies[i].Pos.X), int32(z.Enemies[i].Pos.Y), z.Enemies[i].Width, z.Enemies[i].Height, rl.Red)
	}
}
