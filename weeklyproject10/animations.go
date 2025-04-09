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
	return AnimationFSM{Animations: make(map[string]Animation)}
}

func (a *AnimationFSM) AddAnimation(anim Animation) {
	a.Animations[anim.Name] = anim
}

func (a *AnimationFSM) ChangeAnimationState(newAnim string) {
	if newAnim == a.CurrentAnim.Name {
		return
	}
	_, ok := a.Animations[newAnim]
	if !ok {
		return
	}

	a.CurrentAnim = a.Animations[newAnim]
	a.CurrentAnim.Reset()
}

func (a *AnimationFSM) DrawWithFSM(pos rl.Vector2, size float32, direction float32) {
	a.CurrentAnim.UpdateTime()
	a.CurrentAnim.DrawAnimation(pos, size, direction)
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

func (p *Player) Jump() {
	p.Vel.Y = -25
	p.IsGrounded = false
}

func (p *Player) Move(x float32) {
	if x != 0 {
		p.Direction = x
	}
	p.Vel.X = x * p.Speed
}
