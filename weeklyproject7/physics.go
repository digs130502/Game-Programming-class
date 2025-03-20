package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	X           float32
	Y           float32
	Angle       float32
	Scale       float32
	Texture     rl.Texture2D
	Projectiles []Projectile
}

type Projectile struct {
	X         float32
	Y         float32
	VelocityX float32
	VelocityY float32
	Radius    float32
}

type Planet struct {
	Health int32
	X      float32
	Y      float32
	Radius float32
}

// Constructors
func NewPlayer() Player {
	sprite := rl.LoadTexture("assets/textures/spaceship.png")
	return Player{X: 300, Y: 200, Angle: 0, Scale: 2, Texture: sprite, Projectiles: []Projectile{}}
}

func NewPlanet() Planet {
	return Planet{Health: 10, X: 400, Y: 225, Radius: 30}
}

// Drawing Functions
func (p *Player) DrawPlayer() {
	sourceRect := rl.NewRectangle(0, 0, float32(p.Texture.Width), float32(p.Texture.Height))
	destRect := rl.NewRectangle(p.X, p.Y, float32(p.Texture.Width)*p.Scale, float32(p.Texture.Height)*p.Scale)
	origin := rl.Vector2Scale(rl.NewVector2(float32(p.Texture.Width)/2, float32(p.Texture.Height)/2), p.Scale)
	rl.DrawTexturePro(p.Texture, sourceRect, destRect, origin, p.Angle, rl.RayWhite)
}

func (pl *Planet) DrawPlanet() {
	rl.DrawCircle(int32(pl.X), int32(pl.Y), pl.Radius, rl.Brown)
}

func (p *Player) DrawProjectiles() {
	for _, proj := range p.Projectiles {
		rl.DrawCircle(int32(proj.X), int32(proj.Y), proj.Radius, rl.White)
	}
}

// Player shooting function
func (p *Player) Shoot() {
	angleRad := (p.Angle - 90) * (rl.Pi / 180.0)

	velocityX := float32(300 * math.Cos(float64(angleRad)))
	velocityY := float32(300 * math.Sin(float64(angleRad)))

	//Add projectile to the slice
	projectile := Projectile{X: p.X, Y: p.Y, VelocityX: velocityX, VelocityY: velocityY, Radius: 3}
	p.Projectiles = append(p.Projectiles, projectile)
}

// Projectile Update Function
func (p *Player) UpdateProjectiles() {
	for i := range p.Projectiles {
		p.Projectiles[i].X += p.Projectiles[i].VelocityX * rl.GetFrameTime()
		p.Projectiles[i].Y += p.Projectiles[i].VelocityY * rl.GetFrameTime()
	}
}

// Player Movement
func (p *Player) MovePlayer() {
	if rl.IsKeyDown(rl.KeyA) {
		p.X -= 100 * rl.GetFrameTime()
	}
	if rl.IsKeyDown(rl.KeyD) {
		p.X += 100 * rl.GetFrameTime()
	}
	if rl.IsKeyDown(rl.KeyW) {
		p.Y -= 100 * rl.GetFrameTime()
	}
	if rl.IsKeyDown(rl.KeyS) {
		p.Y += 100 * rl.GetFrameTime()
	}

	//Rotation
	if rl.IsKeyDown(rl.KeyQ) {
		p.Angle -= 100 * rl.GetFrameTime()
	}
	if rl.IsKeyDown(rl.KeyR) {
		p.Angle += 100 * rl.GetFrameTime()
	}

	//Shooting
	if rl.IsKeyPressed(rl.KeySpace) {
		p.Shoot()
		shootSound := rl.LoadSound("assets/audio/shoot.wav")
		rl.PlaySound(shootSound)
	}
}
