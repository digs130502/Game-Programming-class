package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Bar struct {
	Pos rl.Vector2
}

type Ball struct {
	Pos rl.Vector2
	Vel rl.Vector2
}

func NewBar() Bar {
	br := Bar{Pos: rl.NewVector2(325, 420)}
	return br
}

func NewBall(bar Bar) Ball {
	bl := Ball{Pos: rl.NewVector2(bar.Pos.X+75, bar.Pos.Y-11), Vel: rl.NewVector2(0, 0)}
	return bl
}

func main() {
	rl.InitWindow(800, 450, "Breakout")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	bar := NewBar()
	ball := NewBall(bar)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		rl.DrawRectangle(int32(bar.Pos.X), int32(bar.Pos.Y), 150, 10, rl.Orange)
		rl.DrawCircle(int32(ball.Pos.X), int32(ball.Pos.Y), 10, rl.RayWhite)

		if rl.IsKeyDown(rl.KeyA) {
			bar.Pos.X -= 200 * rl.GetFrameTime()
		}
		if rl.IsKeyDown(rl.KeyD) {
			bar.Pos.X += 200 * rl.GetFrameTime()
		}

		rl.EndDrawing()
	}
}
