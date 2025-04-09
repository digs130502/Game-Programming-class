package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Animation struct {
	Name         string
	SpriteSheet  rl.Texture2D
	MaxIndex     int
	CurrentIndex int
	Timer        float32
	SwitchTime   float32
	Loop         bool
}

func (a *Animation) UpdateTime() {
	a.Timer += rl.GetFrameTime()
	if a.Timer > a.SwitchTime {
		a.Timer = 0
		a.CurrentIndex++
	}

	if a.CurrentIndex > a.MaxIndex {
		if a.Loop {
			a.CurrentIndex = 0
		} else {
			a.CurrentIndex = a.MaxIndex
		}
	}
}

type AnimationFSM struct {
	Animations  map[string]Animation
	CurrentAnim Animation
}

func NewAnimationFSM() AnimationFSM {
	return AnimationFSM{
		Animations: make(map[string]Animation),
	}
}

func (a *AnimationFSM) AddAnimation(anim Animation) {
	a.Animations[anim.Name] = anim
}

func (a *AnimationFSM) ChangeAnimationState(newAnim string) {
	if newAnim == a.CurrentAnim.Name {
		return
	}
	anim, ok := a.Animations[newAnim]
	if !ok {
		return
	}
	a.CurrentAnim = anim
	a.CurrentAnim.Reset()
}

func (a *Animation) Reset() {
	a.Timer = 0
	a.CurrentIndex = 0
}

func (a Animation) DrawAnimation(pos rl.Vector2, size float32, direction float32) {
	sourceRect := rl.NewRectangle(float32(32*a.CurrentIndex), 0, 32*direction, 32)
	destRect := rl.NewRectangle(pos.X, pos.Y, size, size)
	rl.DrawTexturePro(a.SpriteSheet, sourceRect, destRect, rl.Vector2Zero(), 0, rl.White)
}

func (a *AnimationFSM) DrawWithFSM(pos rl.Vector2, size float32, direction float32) {
	a.CurrentAnim.UpdateTime()
	a.CurrentAnim.DrawAnimation(pos, size, direction)
}

func NewAnimation(newName string, newSheet rl.Texture2D, spriteNum int, newTime float32) Animation {
	newAnimation := Animation{
		Name:         newName,
		SpriteSheet:  newSheet,
		MaxIndex:     spriteNum - 1,
		CurrentIndex: 0,
		Timer:        0,
		SwitchTime:   newTime,
		Loop:         true,
	}

	return newAnimation
}

type Player struct {
	Pos        rl.Vector2
	Vel        rl.Vector2
	Size       float32
	FeetSize   rl.Vector2
	Color      rl.Color
	Speed      float32
	Direction  float32
	IsGrounded bool
	IsBusy     bool
	AnimationFSM
}

func NewPlayer() Player {
	ani := NewAnimationFSM()
	return Player{
		Pos:          rl.NewVector2(200, 300),
		Vel:          rl.NewVector2(0, 0),
		Size:         4,
		Color:        rl.RayWhite,
		Speed:        200,
		AnimationFSM: ani,
	}
}

func (p *Player) Move(x float32) {
	if x != 0 {
		p.Direction = x
	}
	p.Vel.X = x * p.Speed
}

func (p *Player) UpdatePlayer() {
	// If the current animation is non-looping and has finished, mark player as not busy
	if !p.CurrentAnim.Loop && p.CurrentAnim.CurrentIndex == p.CurrentAnim.MaxIndex {
		p.IsBusy = false
	}

	p.Vel.Y += 800 * rl.GetFrameTime()

	p.Pos = rl.Vector2Add(p.Pos, rl.Vector2Scale(p.Vel, rl.GetFrameTime()))

	if p.Pos.Y >= 300 {
		p.Pos.Y = 300
		p.Vel.Y = 0
		p.IsGrounded = true
	} else {
		p.IsGrounded = false
	}

	// Handle block hold behavior
	if p.CurrentAnim.Name == "block" {
		if !rl.IsKeyDown(rl.KeyF) {
			p.IsBusy = false
		} else if p.CurrentAnim.CurrentIndex == p.CurrentAnim.MaxIndex {
			// Pause timer to freeze frame
			p.CurrentAnim.Timer = 0
		}
		return
	}

	if p.IsBusy {
		return // Don't override current animation if busy (e.g., attacking or blocking)
	}

	if !p.IsGrounded {
		p.ChangeAnimationState("jump")
		return
	}

	if p.Vel.X == 0 {
		p.ChangeAnimationState("idle")
	} else {
		p.ChangeAnimationState("walk")
	}

	if !p.IsGrounded {
		p.AnimationFSM.ChangeAnimationState("jump")
		return
	}

	if p.Vel.X == 0 {
		p.AnimationFSM.ChangeAnimationState("idle")
	} else {
		p.AnimationFSM.ChangeAnimationState("walk")
	}
}

func (p *Player) Jump() {
	if rl.IsKeyPressed(rl.KeySpace) && p.IsGrounded {
		p.Vel.Y = -350
		p.IsGrounded = false
	}
}

func main() {
	rl.InitWindow(800, 450, "Animation Practice")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	player := NewPlayer()
	idleAnimation := NewAnimation("idle", rl.LoadTexture("sprites/p1idle.png"), 3, .5)
	walkAnimation := NewAnimation("walk", rl.LoadTexture("sprites/p1walking.png"), 3, .2)
	jumpAnimation := NewAnimation("jump", rl.LoadTexture("sprites/p1jumping.png"), 5, .2)
	attackAnimation := NewAnimation("attack", rl.LoadTexture("sprites/p1hitting.png"), 4, .035)
	blockAnimation := NewAnimation("block", rl.LoadTexture("sprites/p1blocking.png"), 3, 0.35)
	blockAnimation.Loop = false

	player.AddAnimation(idleAnimation)
	player.AddAnimation(walkAnimation)
	player.AddAnimation(jumpAnimation)
	player.AddAnimation(attackAnimation)
	player.AddAnimation(blockAnimation)
	player.ChangeAnimationState("idle")

	for !rl.WindowShouldClose() {

		//Updates
		creatureVelX := 0.0
		if rl.IsKeyDown(rl.KeyD) {
			creatureVelX = 1
		}
		if rl.IsKeyDown(rl.KeyA) {
			creatureVelX = -1
		}
		if rl.IsKeyDown(rl.KeySpace) {
			player.Jump()
		}
		if rl.IsKeyPressed(rl.KeyF) && !player.IsBusy {
			player.ChangeAnimationState("block")
			player.IsBusy = true
		}

		if rl.IsKeyPressed(rl.KeyG) {
			player.ChangeAnimationState("attack")
			player.IsBusy = true
		}
		player.Move(float32(creatureVelX))
		player.UpdatePlayer()

		//Rendering
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		player.DrawWithFSM(player.Pos, 64, player.Direction)

		rl.EndDrawing()
	}

}
