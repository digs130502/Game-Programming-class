package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Buttons struct {
	Start  Button
	Quit   Button
	Resume Button
	Menu   Button
}

type Button struct {
	X        float32
	Y        float32
	Width    float32
	Height   float32
	Text     string
	TextSize int32
	Colortheme
	Clicked bool
}

type Colortheme struct {
	MainColor  rl.Color
	HoverColor rl.Color
}

func NewColorTheme() Colortheme {
	return Colortheme{
		MainColor:  rl.Brown,
		HoverColor: rl.Red,
	}
}

func NewStartButton() Button {
	cl := NewColorTheme()
	return Button{
		X:          350,
		Y:          200,
		Width:      100,
		Height:     50,
		Text:       "Start Game",
		TextSize:   20,
		Colortheme: cl,
		Clicked:    false,
	}
}

func NewQuitButton() Button {
	cl := NewColorTheme()
	return Button{
		X:          350,
		Y:          275,
		Width:      100,
		Height:     50,
		Text:       "Quit Game",
		TextSize:   20,
		Colortheme: cl,
		Clicked:    false,
	}
}

func NewResumeButton() Button {
	cl := NewColorTheme()
	return Button{
		X:          350,
		Y:          125,
		Width:      100,
		Height:     50,
		Text:       "Resume Game",
		TextSize:   20,
		Colortheme: cl,
		Clicked:    false,
	}
}

func NewMenuButton() Button {
	cl := NewColorTheme()
	return Button{
		X:          350,
		Y:          200,
		Width:      100,
		Height:     50,
		Text:       "Main Menu",
		TextSize:   20,
		Colortheme: cl,
		Clicked:    false,
	}
}

func NewButtons() Buttons {
	return Buttons{
		Start:  NewStartButton(),
		Quit:   NewQuitButton(),
		Resume: NewResumeButton(),
		Menu:   NewMenuButton(),
	}
}

func (b *Buttons) CheckButtons() {
	start := rl.NewRectangle(b.Start.X, b.Start.Y, b.Start.Width, b.Start.Height)
	quit := rl.NewRectangle(b.Quit.X, b.Quit.Y, b.Quit.Width, b.Quit.Height)

	mousePos := rl.GetMousePosition()

	if rl.CheckCollisionPointRec(mousePos, start) && rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		b.Start.Clicked = true
	}
	if rl.CheckCollisionPointRec(mousePos, quit) && rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		b.Quit.Clicked = true
	}
}

func (b *Buttons) CheckResumeButtons() {
	resume := rl.NewRectangle(b.Resume.X, b.Resume.Y, b.Resume.Width, b.Resume.Height)
	menu := rl.NewRectangle(b.Menu.X, b.Menu.Y, b.Menu.Width, b.Menu.Height)
	quit := rl.NewRectangle(b.Quit.X, b.Quit.Y, b.Quit.Width, b.Quit.Height)

	mousePos := rl.GetMousePosition()

	if rl.CheckCollisionPointRec(mousePos, menu) && rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		b.Menu.Clicked = true
	}
	if rl.CheckCollisionPointRec(mousePos, quit) && rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		b.Quit.Clicked = true
	}
	if rl.CheckCollisionPointRec(mousePos, resume) && rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		b.Resume.Clicked = true
	}
}

func (b *Buttons) DrawButtons() {
	rl.DrawRectangle(int32(b.Start.X), int32(b.Start.Y), int32(b.Start.Width), int32(b.Start.Height), b.Start.MainColor)
	rl.DrawRectangle(int32(b.Quit.X), int32(b.Quit.Y), int32(b.Quit.Width), int32(b.Quit.Height), b.Quit.MainColor)
}

// FIXME:
func (b *Buttons) DrawPauseButtons() {
	rl.DrawRectangle(int32(b.Resume.X), int32(b.Resume.Y), int32(b.Resume.Width), int32(b.Resume.Height), b.Resume.MainColor)
	rl.DrawRectangle(int32(b.Menu.X), int32(b.Menu.Y), int32(b.Menu.Width), int32(b.Menu.Height), b.Menu.MainColor)
	rl.DrawRectangle(int32(b.Quit.X), int32(b.Quit.Y), int32(b.Quit.Width), int32(b.Quit.Height), b.Quit.MainColor)
}
