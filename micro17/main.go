package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(800, 600, "Gravity Box Example")
	defer rl.CloseWindow()

	box := Box{
		Pos:   rl.NewVector2(400, 300),
		Vel:   rl.NewVector2(0, 0),
		Size:  rl.NewVector2(50, 50),
		Color: rl.Red,
	}

	blocker1 := Blocker{
		Pos:   rl.NewVector2(250, 150), // Position of the blocker
		Size:  rl.NewVector2(300, 50),
		Color: rl.Gray,
	}
	blocker2 := Blocker{
		Pos:   rl.NewVector2(250, 400), // Position of the blocker
		Size:  rl.NewVector2(300, 50),
		Color: rl.Gray,
	}
	blocker3 := Blocker{
		Pos:   rl.NewVector2(200, 200), // Position of the blocker
		Size:  rl.NewVector2(50, 200),
		Color: rl.Gray,
	}
	blocker4 := Blocker{
		Pos:   rl.NewVector2(550, 200), // Position of the blocker
		Size:  rl.NewVector2(50, 200),
		Color: rl.Gray,
	}

	gravity := rl.NewVector2(0, 0)

	for !rl.WindowShouldClose() {

		// Movement of garvity
		if rl.IsKeyPressed(rl.KeyD) {
			gravity = rl.NewVector2(980, 0)
		}
		if rl.IsKeyPressed(rl.KeyA) {
			gravity = rl.NewVector2(-980, 0)
		}
		if rl.IsKeyPressed(rl.KeyW) {
			gravity = rl.NewVector2(0, -980)
		}
		if rl.IsKeyPressed(rl.KeyS) {
			gravity = rl.NewVector2(0, 980)
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		box.ApplyGravity(gravity)
		box.UpdateBox()

		CheckCollision(&box, blocker1)
		CheckCollision(&box, blocker2)
		CheckCollision(&box, blocker3)
		CheckCollision(&box, blocker4)

		box.DrawBox()
		blocker1.DrawBlocker()
		blocker2.DrawBlocker()
		blocker3.DrawBlocker()
		blocker4.DrawBlocker()

		rl.EndDrawing()
	}
}
