package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 450, "Fighting Game")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	isGameOver := false

	pl1 := NewPlayer1()
	pl2 := NewPlayer2()
	h1 := NewHealthBar1(&pl1)
	h2 := NewHealthBar2(&pl2)

	//Animations
	idleAnimation := NewAnimation("idle", rl.LoadTexture("assets/sprites/p1idle.png"), 3, .5)
	walkAnimation := NewAnimation("walk", rl.LoadTexture("assets/sprites/p1walking.png"), 3, .075)
	jumpAnimation := NewAnimation("jump", rl.LoadTexture("assets/sprites/p1jumping.png"), 5, 0.75)
	jumpAnimation.Loop = false
	attackAnimation := NewAnimation("attack", rl.LoadTexture("assets/sprites/p1hitting.png"), 4, .2)
	blockAnimation := NewAnimation("block", rl.LoadTexture("assets/sprites/p1blocking.png"), 3, .2)
	blockAnimation.Loop = false

	animationFSM := NewAnimationFSM()
	animationFSM.AddAnimation(walkAnimation)
	animationFSM.AddAnimation(idleAnimation)
	animationFSM.AddAnimation(jumpAnimation)
	animationFSM.AddAnimation(attackAnimation)
	animationFSM.AddAnimation(blockAnimation)
	animationFSM.ChangeAnimationState("idle")

	for !rl.WindowShouldClose() {

		if !isGameOver {
			//Updates
			UpdatePlayers(&pl1, &pl2)
			CheckPlayerFloorCollisions(&pl1, &pl2)
			CheckMovement(&pl1, &pl2)
			CheckDamage(&pl1, &pl2, &h1, &h2)

			//Check Game Over
			if pl1.Health <= 0 || pl2.Health <= 0 {
				isGameOver = true
			}

			//Rendering
			rl.BeginDrawing()

			rl.ClearBackground(rl.Black)
			rl.DrawRectangle(0, 350, 800, 100, rl.Orange)
			DrawPlayers(&pl1, &pl2)
			DrawHealthBars(&h1, &h2)

			rl.EndDrawing()
		} else {

			rl.BeginDrawing()

			rl.ClearBackground(rl.Red)

			if pl1.Health > 0 {
				rl.DrawText("Congrats on winning Player 1!", 250, 200, 20, rl.Black)
			} else if pl2.Health > 0 {
				rl.DrawText("Congrats on winning Player 2!", 250, 200, 20, rl.Black)
			}
			rl.DrawText("Press R to restart the game!", 250, 250, 20, rl.Black)

			if rl.IsKeyPressed(rl.KeyR) {
				isGameOver = false
				pl1 = NewPlayer1()
				pl2 = NewPlayer2()
				h1 = NewHealthBar1(&pl1)
				h2 = NewHealthBar2(&pl2)
			}
			rl.EndDrawing()
		}

	}
}
