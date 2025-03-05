package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Bar struct {
	Pos rl.Vector2
}

type Ball struct {
	Pos              rl.Vector2
	Vel              rl.Vector2
	Radius           float32
	IgnoreCollisions bool
}

func NewBar() Bar {
	br := Bar{Pos: rl.NewVector2(325, 420)}
	return br
}

func NewBall(bar Bar) Ball {
	bl := Ball{Pos: rl.NewVector2(bar.Pos.X+75, bar.Pos.Y-11), Vel: rl.NewVector2(0, 0), Radius: 10}
	bl.IgnoreCollisions = false
	return bl
}

type PhysicsBody struct {
	Pos              rl.Vector2
	Vel              rl.Vector2
	Radius           float32
	IgnoreCollisions bool
}

func NewPhysicsBody(newPos rl.Vector2, newVel rl.Vector2, newRadius float32) PhysicsBody {
	pb := PhysicsBody{Pos: newPos, Vel: newVel, Radius: newRadius}
	pb.IgnoreCollisions = false
	return pb
}

func (bl *PhysicsBody) Bounce() {
	bl.Vel = rl.Vector2Scale(bl.Vel, -1)
}

func (pb *PhysicsBody) CheckIntersection(otherPb PhysicsBody) bool {
	if rl.Vector2Distance(pb.Pos, otherPb.Pos) <= pb.Radius+otherPb.Radius {
		pb.Bounce()
		return true
	}
	return false
}

func main() {
	rl.InitWindow(800, 450, "Breakout")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	bar := NewBar()
	ball := NewBall(bar)
	ballStuck := true

	for !rl.WindowShouldClose() {
		//Movement controls
		if rl.IsKeyDown(rl.KeyA) {
			bar.Pos.X -= 200 * rl.GetFrameTime()
		}
		if rl.IsKeyDown(rl.KeyD) {
			bar.Pos.X += 200 * rl.GetFrameTime()
		}

		if ballStuck {
			ball.Pos.X = bar.Pos.X + 75
			ball.Pos.Y = bar.Pos.Y - 11

			//Release bar
			if rl.IsKeyPressed(rl.KeySpace) && !rl.IsKeyDown(rl.KeyA) && !rl.IsKeyDown(rl.KeyD) {
				ballStuck = false
				ball.Vel = rl.NewVector2(0, -200) // Initial velocity
			}
			if rl.IsKeyPressed(rl.KeySpace) && rl.IsKeyDown(rl.KeyA) && !rl.IsKeyDown(rl.KeyD) {
				ballStuck = false
				ball.Vel = rl.NewVector2(-300, -200)
			}
			if rl.IsKeyPressed(rl.KeySpace) && !rl.IsKeyDown(rl.KeyA) && rl.IsKeyDown(rl.KeyD) {
				ballStuck = false
				ball.Vel = rl.NewVector2(300, -200)
			}
		} else {
			ball.Pos.X += ball.Vel.X * rl.GetFrameTime()
			ball.Pos.Y += ball.Vel.Y * rl.GetFrameTime()
		}

		if ball.Pos.X >= 800 || ball.Pos.X <= 0 {

		}
		if ball.Pos.Y >= 450 || ball.Pos.Y <= 0 {

		}

		//Rendering
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		rl.DrawRectangle(int32(bar.Pos.X), int32(bar.Pos.Y), 150, 10, rl.Orange)
		rl.DrawCircle(int32(ball.Pos.X), int32(ball.Pos.Y), ball.Radius, rl.RayWhite)

		rl.EndDrawing()
	}
}
