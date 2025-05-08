package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Words section:
type WordGroup struct {
	Description string
	Words       []string
}

type WordBank map[string]WordGroup

func LoadWordBank(filename string) (WordBank, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var bank WordBank
	if err := json.Unmarshal(data, &bank); err != nil {
		return nil, err
	}

	return bank, nil
}

func GetRandomWord(bank WordBank, level int) string {
	group, ok := bank[fmt.Sprintf("%d", level)]
	if !ok || len(group.Words) == 0 {
		return "default"
	}
	return group.Words[rand.Intn(len(group.Words))]
}

// Keyboard section:
func (pl *Player) CheckKeyboardInput() {
	// Collect all characters typed this frame
	char := rl.GetCharPressed()
	for char != 0 {
		pl.Input += string(char)
		char = rl.GetCharPressed()
	}

	// Handle backspace
	if rl.IsKeyPressed(rl.KeyBackspace) && len(pl.Input) > 0 {
		pl.Input = pl.Input[:len(pl.Input)-1]
	}

	// Display input on screen
	rl.DrawText("Input: "+pl.Input, 20, 40, 24, rl.Green)
}
