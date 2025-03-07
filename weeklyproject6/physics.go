package main

import rl "github.com/gen2brain/raylib-go/raylib"

type PhysicsBody struct {
	Pos              rl.Vector2
	Vel              rl.Vector2
	Radius           float32
	IgnoreCollisions bool
}

type Ball struct {
	PhysicsBody
}

type Bar struct {
	PhysicsBody
	Width  float32
	Height float32
}

type Square struct {
	PhysicsBody
	Height float32
	Width  float32
}

func NewBall(bar Bar) Ball {
	return Ball{
		PhysicsBody: PhysicsBody{
			Pos:              rl.NewVector2(bar.Pos.X+75, bar.Pos.Y-11),
			Vel:              rl.NewVector2(0, 0),
			Radius:           10,
			IgnoreCollisions: false,
		},
	}
}

func NewBar() Bar {
	return Bar{
		PhysicsBody: PhysicsBody{
			Pos: rl.NewVector2(325, 420),
			Vel: rl.NewVector2(0, 0),
		},
		Width:  150,
		Height: 10,
	}
}

func NewSquare(pos rl.Vector2, width, height float32) Square {
	return Square{
		PhysicsBody: PhysicsBody{
			Pos:              pos,
			Vel:              rl.NewVector2(0, 0),
			Radius:           0,
			IgnoreCollisions: false,
		},
		Width:  width,
		Height: height,
	}
}

func DrawBall(ball Ball) {
	rl.DrawCircle(int32(ball.PhysicsBody.Pos.X), int32(ball.PhysicsBody.Pos.Y), ball.PhysicsBody.Radius, rl.RayWhite)
}

func DrawBar(bar Bar) {
	rl.DrawRectangle(int32(bar.PhysicsBody.Pos.X), int32(bar.PhysicsBody.Pos.Y), int32(bar.Width), int32(bar.Height), rl.Orange)
}

func DrawSquare(sq Square) {
	rl.DrawRectangle(int32(sq.PhysicsBody.Pos.X), int32(sq.PhysicsBody.Pos.Y), int32(sq.Width), int32(sq.Height), rl.Orange)
}

var squares []Square

func InitSquares(rows, cols int, startX, startY, width, height, padding float32) {
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			posX := startX + float32(col)*(width+padding)
			posY := startY + float32(row)*(height+padding)

			squares = append(squares, NewSquare(rl.NewVector2(posX, posY), width, height))
		}
	}
}

func CheckBallSquareCollision(ball Ball, square Square) bool {
	return ball.Pos.X+ball.Radius >= square.Pos.X &&
		ball.Pos.X-ball.Radius <= square.Pos.X+square.Width &&
		ball.Pos.Y+ball.Radius >= square.Pos.Y &&
		ball.Pos.Y-ball.Radius <= square.Pos.Y+square.Height
}
