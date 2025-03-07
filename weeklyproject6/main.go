package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 450, "Breakout")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	bar := NewBar()
	ball := NewBall(bar)
	ballStuck := true
	InitSquares(5, 9, 50, 50, 70, 30, 10)

	resetGame := func() {
		bar = NewBar()
		ball = NewBall(bar)
		ballStuck = true
		squares = nil
		InitSquares(5, 9, 50, 50, 70, 30, 10)
	}

	for !rl.WindowShouldClose() {
		if rl.IsKeyDown(rl.KeyA) {
			bar.Pos.X -= 200 * rl.GetFrameTime()
			if ballStuck {
				ball.Pos.X -= 200 * rl.GetFrameTime()
			}
		}
		if rl.IsKeyDown(rl.KeyD) {
			bar.Pos.X += 200 * rl.GetFrameTime()
			if ballStuck {
				ball.Pos.X += 200 * rl.GetFrameTime()
			}
		}

		if rl.IsKeyPressed(rl.KeySpace) && ballStuck {
			ballStuck = false
			ball.Vel = rl.NewVector2(0, -200)
			if rl.IsKeyDown(rl.KeyA) {
				ball.Vel = rl.NewVector2(-300, -200)
			}
			if rl.IsKeyDown(rl.KeyD) {
				ball.Vel = rl.NewVector2(300, -200)
			}
		}

		if !ballStuck {
			ball.Pos.X += ball.Vel.X * rl.GetFrameTime()
			ball.Pos.Y += ball.Vel.Y * rl.GetFrameTime()
		}

		if ball.Pos.X+ball.Radius >= 800 || ball.Pos.X-ball.Radius <= 0 {
			ball.Vel.X *= -1
		}
		if ball.Pos.Y-ball.Radius <= 0 {
			ball.Vel.Y *= -1
		}
		if ball.Pos.Y+ball.Radius >= 450 {
			resetGame()
		}

		if ball.Pos.Y+ball.Radius >= bar.Pos.Y &&
			ball.Pos.X >= bar.Pos.X &&
			ball.Pos.X <= bar.Pos.X+bar.Width &&
			ball.Vel.Y > 0 {

			ball.Vel.Y *= -1
			hitPosition := (ball.Pos.X - bar.Pos.X) / bar.Width
			ball.Vel.X = (hitPosition - 0.5) * 400
		}

		for i := 0; i < len(squares); i++ {
			if CheckBallSquareCollision(ball, squares[i]) {
				square := squares[i]
				squares = append(squares[:i], squares[i+1:]...)

				ballCenterX, ballCenterY := ball.Pos.X, ball.Pos.Y
				squareLeft, squareRight := square.Pos.X, square.Pos.X+square.Width
				squareTop, squareBottom := square.Pos.Y, square.Pos.Y+square.Height

				closestX := max(squareLeft, min(ballCenterX, squareRight))
				closestY := max(squareTop, min(ballCenterY, squareBottom))

				distX := ballCenterX - closestX
				distY := ballCenterY - closestY

				if math.Abs(float64(distX)) > math.Abs(float64(distY)) {
					ball.Vel.X *= -1
				} else {
					ball.Vel.Y *= -1
				}

				speedIncreaseFactor := 1.02
				currentSpeed := rl.Vector2Length(ball.Vel)
				newSpeed := currentSpeed * float32(speedIncreaseFactor)
				ball.Vel = rl.Vector2Scale(rl.Vector2Normalize(ball.Vel), newSpeed)

				break
			}
		}

		if len(squares) == 0 {
			resetGame()
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		DrawBar(bar)
		DrawBall(ball)

		for _, square := range squares {
			DrawSquare(square)
		}

		rl.EndDrawing()
	}
}
