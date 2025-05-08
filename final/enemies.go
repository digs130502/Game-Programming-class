package main

import (
	"fmt"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Zone struct {
	Enemies         []Enemy
	Timer           float32
	Level           int
	LevelTimer      float32
	LevelUpInterval float32
}

type Enemy struct {
	Pos    rl.Vector2
	Vel    rl.Vector2
	Width  int32
	Height int32
	Word   string
}

func NewEnemy(wordBank WordBank, level int) Enemy {
	return Enemy{
		Pos:    rl.NewVector2(float32(rand.Intn(800)), 0),
		Vel:    rl.NewVector2(0, 100),
		Width:  50,
		Height: 50,
		Word:   GetRandomWord(wordBank, level),
	}
}

func NewZone() Zone {
	return Zone{
		Enemies:         []Enemy{},
		Level:           1,
		LevelTimer:      0,
		LevelUpInterval: 10,
	}
}

func (z *Zone) CheckEnemyCreation(n float32, wordBank WordBank) {
	z.Timer += rl.GetFrameTime()

	if z.Timer > n {
		z.Enemies = append(z.Enemies, NewEnemy(wordBank, z.Level))
		fmt.Println("Enemy Created!")
		z.Timer = 0
	}
}

func (z *Zone) UpdateEnemies(pl *Player, h *HealthBar) {
	for i := range z.Enemies {
		vel := rl.Vector2Scale(z.Enemies[i].Vel, rl.GetFrameTime())
		z.Enemies[i].Pos = rl.Vector2Add(z.Enemies[i].Pos, vel)
	}

	for i := len(z.Enemies) - 1; i >= 0; i-- {
		if z.Enemies[i].Pos.Y+float32(z.Enemies[i].Height) >= 450 {
			z.Enemies = append(z.Enemies[:i], z.Enemies[i+1:]...)
			pl.Health -= 1
			h.Width -= 20
		}
	}

	for i := len(z.Enemies) - 1; i >= 0; i-- {
		if pl.Input == z.Enemies[i].Word {
			z.Enemies = append(z.Enemies[:i], z.Enemies[i+1:]...) // delete matched enemy
			pl.Input = ""                                         // reset input
			break                                                 // one word per input
		}
	}

	z.LevelTimer += rl.GetFrameTime()
	if z.LevelTimer >= z.LevelUpInterval && z.Level < 3 {
		z.Level++
		z.LevelTimer = 0
		fmt.Println("Level up! Now at level", z.Level)
	}

}

func (z *Zone) DrawEnemies() {
	for i := range z.Enemies {
		rl.DrawRectangle(int32(z.Enemies[i].Pos.X), int32(z.Enemies[i].Pos.Y), z.Enemies[i].Width, z.Enemies[i].Height, rl.Red)

		textWidth := rl.MeasureText(z.Enemies[i].Word, 20)
		textX := z.Enemies[i].Pos.X + (float32(z.Enemies[i].Width) / 2) - (float32(textWidth) / 2)
		textY := z.Enemies[i].Pos.Y + (float32(z.Enemies[i].Height) + 5)

		rl.DrawText(z.Enemies[i].Word, int32(textX), int32(textY), 20, rl.Black)
	}

	rl.DrawText(fmt.Sprintf("Level: %d", z.Level), 20, 60, 24, rl.Black)
}
