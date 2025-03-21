package main

import (
	"math"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

//TODO: Make the asteroids spawn away from the planet. Not just randomly inside of the screen boundaries

type Zone struct {
	Asteroids      []Asteroid
	CargoAsteroids []CargoAsteroid
}

type Asteroid struct {
	X         float32
	Y         float32
	VelocityX float32
	VelocityY float32
	Radius    float32
	Color     rl.Color
}

type CargoAsteroid struct {
	X         float32
	Y         float32
	VelocityX float32
	VelocityY float32
	Radius    float32
	Color     rl.Color
}

func NewZone() Zone {
	return Zone{Asteroids: []Asteroid{}, CargoAsteroids: []CargoAsteroid{}}
}

func (z *Zone) NewAsteroid(radius float32) {

	startX := rand.Float32() * 800
	startY := rand.Float32() * 450

	dirX := 400 - startX
	dirY := 225 - startY

	length := float32(math.Sqrt(float64(dirX*dirX + dirY*dirY)))
	velocityX := (dirX / length) * 0.75
	velocityY := (dirY / length) * 0.75

	asteroid := Asteroid{
		X:         startX,
		Y:         startY,
		VelocityX: velocityX,
		VelocityY: velocityY,
		Radius:    radius,
		Color:     rl.NewColor(uint8(rand.Intn(256)), uint8(rand.Intn(256)), uint8(rand.Intn(256)), 255),
	}

	z.Asteroids = append(z.Asteroids, asteroid)
}

func (z *Zone) NewCargoAsteroid(a Asteroid) {
	startX := a.X
	startY := a.Y

	angle := rand.Float32() * 2 * rl.Pi
	velocityX := float32(math.Cos(float64(angle))) * 1
	velocityY := float32(math.Sin(float64(angle))) * 1

	cargoAsteroid := CargoAsteroid{
		X:         startX,
		Y:         startY,
		VelocityX: velocityX,
		VelocityY: velocityY,
		Radius:    5,
		Color:     rl.Yellow,
	}

	z.CargoAsteroids = append(z.CargoAsteroids, cargoAsteroid)
}

func (z *Zone) DrawAsteroid() {
	for _, ast := range z.Asteroids {
		rl.DrawCircle(int32(ast.X), int32(ast.Y), ast.Radius, ast.Color)
	}
}

func (z *Zone) DrawCargoAsteroid() {
	for _, carAst := range z.CargoAsteroids {
		rl.DrawCircle(int32(carAst.X), int32(carAst.Y), carAst.Radius, carAst.Color)
	}
}

func (z *Zone) UpdateAsteroids() {
	for i := range z.Asteroids {
		z.Asteroids[i].X += z.Asteroids[i].VelocityX
		z.Asteroids[i].Y += z.Asteroids[i].VelocityY
	}
}

func (z *Zone) UpdateCargoAsteroids() {
	for i := range z.CargoAsteroids {
		z.CargoAsteroids[i].X += z.CargoAsteroids[i].VelocityX
		z.CargoAsteroids[i].Y += z.CargoAsteroids[i].VelocityY
	}
}

func (z *Zone) SpawnSmallerAsteroids(ast Asteroid) {
	newRadius := 10

	for i := 0; i < 2; i++ {

		angle := rand.Float32() * 2 * rl.Pi
		velocityX := float32(math.Cos(float64(angle))) * 1
		velocityY := float32(math.Sin(float64(angle))) * 1

		smallAsteroid := Asteroid{
			X:         ast.X,
			Y:         ast.Y,
			VelocityX: velocityX,
			VelocityY: velocityY,
			Radius:    float32(newRadius),
			Color:     ast.Color,
		}

		z.Asteroids = append(z.Asteroids, smallAsteroid)
	}
}

/*
// TODO: Check and understand this
// - Add mini asteroid after collision
func (z *Zone) CheckAsteroidCollision(p *Planet, pl *Player) {
	// Load sound
	explosion := rl.LoadSound("assets/audio/explosion.wav")

	// Temporary slice for remaining asteroids
	var newAsteroids []Asteroid
	// Temporary slice for remaining projectiles
	var newProjectiles []Projectile

	// Loop through asteroids
	for _, asteroid := range z.Asteroids {
		planetCollision := false
		projectileDestroyed := false

		// Check collision with planet
		distanceX := asteroid.X - p.X
		distanceY := asteroid.Y - p.Y
		distance := float32(math.Sqrt(float64(distanceX*distanceX + distanceY*distanceY)))

		if distance <= (asteroid.Radius + p.Radius) {
			p.Health -= 1
			planetCollision = true
		}

		for _, proj := range pl.Projectiles {
			distProjX := asteroid.X - proj.X
			distProjY := asteroid.Y - proj.Y
			distProj := float32(math.Sqrt(float64(distProjX*distProjX + distProjY*distProjY)))

			if distProj <= (asteroid.Radius + proj.Radius) {
				projectileDestroyed = true
				rl.PlaySound(explosion)

				if asteroid.Radius > 10 {
					z.SpawnSmallerAsteroids(asteroid)
				} else {
					z.NewCargoAsteroid(asteroid)
					z.NewCargoAsteroid(asteroid)
				}

				// Skip adding this projectile, but BREAK here so we don't continue looping
				continue
			}

			newProjectiles = append(newProjectiles, proj)
		}

		// Keep asteroid if it didn't collide
		if !planetCollision && !projectileDestroyed {
			newAsteroids = append(newAsteroids, asteroid)
		}
	}

	// Update asteroids and projectiles lists
	z.Asteroids = newAsteroids
	pl.Projectiles = newProjectiles // ✅ Now properly removes hit projectiles
}*/

func (z *Zone) CheckAsteroidCollision() {

}

func (z *Zone) CheckCargoAsteroidCollision(pl *Player) {
	//Load sound
	bounce := rl.LoadSound("assets/audio/bounce.wav")

	var newCargoAsteroids []CargoAsteroid

	for _, cargoAsteroid := range z.CargoAsteroids {

		distanceX := cargoAsteroid.X - pl.X
		distanceY := cargoAsteroid.Y - pl.Y
		distance := float32(math.Sqrt(float64(distanceX*distanceX + distanceY*distanceY)))

		playerRadius := (float32(pl.Texture.Width) / 2) * pl.Scale

		if distance <= (cargoAsteroid.Radius + playerRadius) {
			rl.PlaySound(bounce)
			pl.Cargo += 1
			continue
		}

		newCargoAsteroids = append(newCargoAsteroids, cargoAsteroid)
	}

	z.CargoAsteroids = newCargoAsteroids
}
