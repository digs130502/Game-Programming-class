package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Player1 struct {
	Texture      rl.Texture2D
	Pos          rl.Vector2
	Vel          rl.Vector2
	Gravity      float32
	Scale        float32
	Rotation     float32
	FrameRec     rl.Rectangle
	CurrentFrame int
}

type Player2 struct {
	Texture      rl.Texture2D
	Pos          rl.Vector2
	Vel          rl.Vector2
	Gravity      float32
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
		Gravity:  70,
		Scale:    3,
		Rotation: 0,
		FrameRec: rl.NewRectangle(0, 0, float32(playerSpr.Width), float32(playerSpr.Height)),
	}
}

func NewPlayer2() Player2 {
	playerSpr := rl.LoadTexture("assets/sprites/player2.png")
	return Player2{
		Texture:  playerSpr,
		Pos:      rl.NewVector2(750, 300),
		Vel:      rl.NewVector2(0, 0),
		Gravity:  70,
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

func UpdatePlayers(pl1 *Player1, pl2 *Player2) {
	//Gravity Updates
	pl1.Pos.Y += pl1.Gravity * rl.GetFrameTime()
	pl2.Pos.Y += pl2.Gravity * rl.GetFrameTime()
}

func CheckPlayerCollisions(pl1 *Player1, pl2 *Player2) {
	//Player 1
	playerHeight := pl1.FrameRec.Height * pl1.Scale
	playerBottom := pl1.Pos.Y + (playerHeight / 2)

	if playerBottom >= 350 {
		pl1.Pos.Y = 350 - (playerHeight / 2)
	}

	//Player 2
	playerHeight2 := pl2.FrameRec.Height * pl2.Scale
	playerBottom2 := pl2.Pos.Y + (playerHeight2 / 2)

	if playerBottom2 >= 350 {
		pl2.Pos.Y = 350 - (playerHeight2 / 2)
	}
}

func CheckMovement(pl1 *Player1, pl2 *Player2) {
	//Player 1 Movement
	if rl.IsKeyDown(rl.KeyD) {
		pl1.Pos.X += 100 * rl.GetFrameTime()
	}
	if rl.IsKeyDown(rl.KeyA) {
		pl1.Pos.X -= 100 * rl.GetFrameTime()
	}
	//FIXME:
	if rl.IsKeyPressed(rl.KeyW) {
		pl1.Pos.Y -= 100 * rl.GetFrameTime()
	}

	//Player 2 Movement
	if rl.IsKeyDown(rl.KeyRight) {
		pl2.Pos.X += 100 * rl.GetFrameTime()
	}
	if rl.IsKeyDown(rl.KeyLeft) {
		pl2.Pos.X -= 100 * rl.GetFrameTime()
	}
	//FIXME:
	if rl.IsKeyDown(rl.KeyUp) {
		pl2.Pos.Y -= 100 * rl.GetFrameTime()
	}
}
