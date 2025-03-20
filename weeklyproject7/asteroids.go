package main

import (
	"math"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

//TODO: Make the asteroids spawn away from the planet. Not just randomly inside of the screen boundaries

type Zone struct {
	Asteroids []Asteroid
}

type Asteroid struct {
	X         float32
	Y         float32
	VelocityX float32
	VelocityY float32
	Radius    float32
	Color     rl.Color
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
	velocityX := (dirX / length) * 0.75
	velocityY := (dirY / length) * 0.75

	asteroid := Asteroid{
		X:         startX,
		Y:         startY,
		VelocityX: velocityX,
		VelocityY: velocityY,
		Radius:    20,
		Color:     rl.NewColor(uint8(rand.Intn(256)), uint8(rand.Intn(256)), uint8(rand.Intn(256)), 255),
	}

	z.Asteroids = append(z.Asteroids, asteroid)
}

func (z *Zone) DrawAsteroid() {
	for _, ast := range z.Asteroids {
		rl.DrawCircle(int32(ast.X), int32(ast.Y), ast.Radius, ast.Color)
	}
}

func (z *Zone) UpdateAsteroids() {
	for i := range z.Asteroids {
		z.Asteroids[i].X += z.Asteroids[i].VelocityX
		z.Asteroids[i].Y += z.Asteroids[i].VelocityY
	}
}

// TODO: Check and understand this
// - Add mini asteroid after collision
func (z *Zone) CheckAsteroidCollision(p *Planet, pl *Player) {
	//Load sound
	explosion := rl.LoadSound("assets/audio/explosion.wav")
	// Temporary slice to hold asteroids that should remain
	var newAsteroids []Asteroid

	// Loop through asteroids (forward this time, since we won't modify `z.Asteroids` in-place)
	for _, asteroid := range z.Asteroids {
		planetCollision := false
		projectileDestroyed := false

		// Check collision with the planet
		distanceX := asteroid.X - p.X
		distanceY := asteroid.Y - p.Y
		distance := float32(math.Sqrt(float64(distanceX*distanceX + distanceY*distanceY)))

		if distance <= (asteroid.Radius + p.Radius) {
			p.Health -= 1          // Damage planet
			planetCollision = true // Mark asteroid for removal
		}

		// Temporary slice for remaining projectiles
		var newProjectiles []Projectile

		// Check collision with projectiles
		for _, proj := range pl.Projectiles {
			distProjX := asteroid.X - proj.X
			distProjY := asteroid.Y - proj.Y
			distProj := float32(math.Sqrt(float64(distProjX*distProjX + distProjY*distProjY)))

			if distProj <= (asteroid.Radius + proj.Radius) {
				projectileDestroyed = true // Mark asteroid for removal
				rl.PlaySound(explosion)
				continue // Skip adding this projectile (it gets removed)
			}

			// Keep this projectile if it didn't collide
			newProjectiles = append(newProjectiles, proj)
		}

		// Assign updated projectile list
		pl.Projectiles = newProjectiles

		// Keep the asteroid only if it did NOT collide with planet or projectile
		if !planetCollision && !projectileDestroyed {
			newAsteroids = append(newAsteroids, asteroid)
		}
	}

	// Assign updated asteroid list
	z.Asteroids = newAsteroids
}
