package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Creature struct {
	SpriteRenderer
	speed float32
	Health
}

type Health struct {
	currentHealth float32
	maxHealth     float32
}

func NewHealth(maxHealth float32) Health {
	return Health{
		currentHealth: maxHealth,
		maxHealth:     maxHealth,
	}
}

func (h *Health) Damage(damage float32) {
	h.currentHealth -= damage
	if h.currentHealth < 0 {
		h.currentHealth = 0
	}
}

func (h *Health) Heal(heal float32) {
	h.currentHealth += heal
	if h.currentHealth > h.maxHealth {
		h.currentHealth = h.maxHealth
	}
}

func NewCreature(pos rl.Vector2, speed, size float32, sprite rl.Texture2D, maxHealth float32, color rl.Color) Creature {
	sr := NewSpriteRenderer(sprite, color, pos)
	return Creature{
		SpriteRenderer: sr,
		speed:          speed,
		Health:         NewHealth(maxHealth),
	}
}

func (c *Creature) Move(offset rl.Vector2) {
	c.Position = rl.Vector2Add(c.Position, rl.Vector2Scale(offset, c.speed*rl.GetFrameTime()))
}
