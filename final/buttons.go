package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Buttons struct {
	Start StartButton
	Quit  QuitButton
}

type StartButton struct {
	X        float32
	Y        float32
	Width    float32
	Height   float32
	Text     string
	TextSize int32
	Colortheme
}

type QuitButton struct {
	X        float32
	Y        float32
	Width    float32
	Height   float32
	Text     string
	TextSize int32
	Colortheme
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

func NewStartButton() StartButton {
	cl := NewColorTheme()
	return StartButton{
		X:          350,
		Y:          200,
		Width:      100,
		Height:     50,
		Text:       "Start Game",
		TextSize:   20,
		Colortheme: cl,
	}
}

func NewQuitButton() QuitButton {
	cl := NewColorTheme()
	return QuitButton{
		X:          350,
		Y:          275,
		Width:      100,
		Height:     50,
		Text:       "Quit Game",
		TextSize:   20,
		Colortheme: cl,
	}
}

func NewButtons() Buttons {
	return Buttons{
		Start: NewStartButton(),
		Quit:  NewQuitButton(),
	}
}

func (b *Buttons) DrawButtons() {
	rl.DrawRectangle(int32(b.Start.X), int32(b.Start.Y), int32(b.Start.Width), int32(b.Start.Height), b.Start.MainColor)
	rl.DrawRectangle(int32(b.Quit.X), int32(b.Quit.Y), int32(b.Quit.Width), int32(b.Quit.Height), b.Quit.MainColor)
}
