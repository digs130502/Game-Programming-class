package main

import (
	"fmt"
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

func (z *Zone) CheckAsteroidCollision(pl *Player, p *Planet, explosion rl.Sound) {
	var newAsteroids []Asteroid
	var spawnedAsteroids []Asteroid

	for _, asteroid := range z.Asteroids {
		asteroidDestroyed := false

		// ✅ Check planet collision first
		dxPlanet := asteroid.X - p.X
		dyPlanet := asteroid.Y - p.Y
		distanceToPlanet := float32(math.Sqrt(float64(dxPlanet*dxPlanet + dyPlanet*dyPlanet)))

		if distanceToPlanet <= (asteroid.Radius + p.Radius) {
			p.Health -= 1
			asteroidDestroyed = true
			fmt.Println("asteroid hit planet:", asteroid)
		}

		// ✅ Only check projectile collision if not already destroyed by planet
		if !asteroidDestroyed {
			for i, proj := range pl.Projectiles {
				dx := asteroid.X - proj.X
				dy := asteroid.Y - proj.Y
				distance := float32(math.Sqrt(float64(dx*dx + dy*dy)))

				if distance <= (asteroid.Radius + proj.Radius) {
					rl.PlaySound(explosion)
					asteroidDestroyed = true
					fmt.Println("destroying with projectile:", asteroid)

					if asteroid.Radius > 10 {
						for i := 0; i < 2; i++ {
							angle := rand.Float64() * 2 * math.Pi
							velocityX := float32(math.Cos(angle)) * 1
							velocityY := float32(math.Sin(angle)) * 1

							smallAsteroid := Asteroid{
								X:         asteroid.X,
								Y:         asteroid.Y,
								VelocityX: velocityX,
								VelocityY: velocityY,
								Radius:    10.0,
								Color:     asteroid.Color,
							}
							spawnedAsteroids = append(spawnedAsteroids, smallAsteroid)
						}
					} else {
						z.NewCargoAsteroid(asteroid)
						z.NewCargoAsteroid(asteroid)
					}

					// Remove projectile that hit
					pl.Projectiles = append(pl.Projectiles[:i], pl.Projectiles[i+1:]...)
					break
				}
			}
		}

		if !asteroidDestroyed {
			newAsteroids = append(newAsteroids, asteroid)
		}
	}

	z.Asteroids = append(newAsteroids, spawnedAsteroids...)
}

// TODO: Make the collision better. Player hitbox is too big
func (z *Zone) CheckCargoAsteroidCollision(pl *Player, bounce rl.Sound) {

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
