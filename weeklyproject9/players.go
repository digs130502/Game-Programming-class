package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Player1 struct {
	Texture      rl.Texture2D
	Pos          rl.Vector2
	Vel          rl.Vector2
	Scale        float32
	Rotation     float32
	FrameRec     rl.Rectangle
	CurrentFrame int
}

type Player2 struct {
	Texture      rl.Texture2D
	Pos          rl.Vector2
	Vel          rl.Vector2
	Scale        float32
	Rotation     float32
	FrameRec     rl.Rectangle
	CurrentFrame int
}

func NewPlayer1() Player1 {
	playerSpr := rl.LoadTexture("assets/sprites/player.png")
	return Player1{
		Texture:  playerSpr,
		Pos:      rl.NewVector2(50, 300),
		Vel:      rl.NewVector2(0, 0),
		Scale:    3,
		Rotation: 0,
		FrameRec: rl.NewRectangle(0, 0, float32(playerSpr.Width), float32(playerSpr.Height)),
	}
}

func NewPlayer2() Player2 {
	playerSpr := rl.LoadTexture("assets/sprites/player.png")
	return Player2{
		Texture:  playerSpr,
		Pos:      rl.NewVector2(750, 300),
		Vel:      rl.NewVector2(0, 0),
		Scale:    3,
		Rotation: 0,
		FrameRec: rl.NewRectangle(0, 0, float32(playerSpr.Width), float32(playerSpr.Height)),
	}
}

func DrawPlayers(pl1 *Player1, pl2 *Player2) {
	rl.DrawTexturePro(
		pl1.Texture,
		pl1.FrameRec,
		rl.NewRectangle(pl1.Pos.X, pl1.Pos.Y, pl1.FrameRec.Width*pl1.Scale, pl1.FrameRec.Height*pl1.Scale),
		rl.NewVector2(pl1.FrameRec.Width*pl1.Scale/2, pl1.FrameRec.Height*pl1.Scale/2),
		pl1.Rotation,
		rl.White,
	)
	rl.DrawTexturePro(
		pl2.Texture,
		pl2.FrameRec,
		rl.NewRectangle(pl2.Pos.X, pl2.Pos.Y, pl2.FrameRec.Width*pl2.Scale, pl2.FrameRec.Height*pl2.Scale),
		rl.NewVector2(pl2.FrameRec.Width*pl2.Scale/2, pl2.FrameRec.Height*pl2.Scale/2),
		pl2.Rotation,
		rl.White,
	)
}
