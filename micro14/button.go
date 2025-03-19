package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Button struct {
	X              int32
	Y              int32
	Width          int32
	Height         int32
	Clicked        bool
	mouseOver      bool
	AlternateColor bool
	Color          rl.Color
}

func NewButton() Button {
	return Button{
		X:              300,
		Y:              200,
		Width:          50,
		Height:         50,
		AlternateColor: false,
		Color:          rl.Blue,
	}
}

func DrawButton(b Button) {
	rl.DrawRectangle(b.X, b.Y, b.Width, b.Height, b.Color)
}

func (b *Button) CheckMouseOver() {
	mousePos := rl.GetMousePosition()
	b.mouseOver = false
	if int32(mousePos.X) < b.X || int32(mousePos.X) > b.X+b.Width {
		return
	}
	if int32(mousePos.Y) < b.Y || int32(mousePos.Y) > b.Y+b.Height {
		return
	}
	b.mouseOver = true
}

func (b *Button) IsClicked() bool {
	b.CheckMouseOver()

	if b.mouseOver && rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		b.Clicked = true
		b.AlternateColor = !b.AlternateColor
		return true
	}

	b.Clicked = false
	return false
}

func (b *Button) ChangeColor() {
	if b.AlternateColor {
		b.Color = rl.Green
	} else {
		b.Color = rl.Blue
	}
}
