package main

import (
	"math"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Zone struct {
	Asteroids []Asteroid
}

type Asteroid struct {
	X         float32
	Y         float32
	VelocityX float32
	VelocityY float32
	Radius    float32
}

func NewZone() Zone {
	return Zone{Asteroids: []Asteroid{}}
}

func (z *Zone) NewAsteroid() {

	// Randomly spawn asteroid anywhere on screen
	startX := rand.Float32() * 800
	startY := rand.Float32() * 450

	// Calculate direction vector toward the center
	dirX := 400 - startX
	dirY := 225 - startY

	// Normalize direction vector (ensures uniform speed)
	length := float32(math.Sqrt(float64(dirX*dirX + dirY*dirY)))
	velocityX := (dirX / length) * 2
	velocityY := (dirY / length) * 2

	asteroid := Asteroid{
		X:         startX,
		Y:         startY,
		VelocityX: velocityX,
		VelocityY: velocityY,
		Radius:    20,
	}

	z.Asteroids = append(z.Asteroids, asteroid)
}

func (z *Zone) DrawAsteroid() {
	for _, ast := range z.Asteroids {
		rl.DrawCircle(int32(ast.X), int32(ast.Y), ast.Radius, rl.Orange)
	}
}

func (z *Zone) UpdateAsteroids() {
	for i := range z.Asteroids {
		z.Asteroids[i].X += z.Asteroids[i].VelocityX
		z.Asteroids[i].Y += z.Asteroids[i].VelocityY
	}
}
