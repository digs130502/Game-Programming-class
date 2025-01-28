package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	//playerLost variable
	var playerLost bool = false

	//Answer variable to be scanned
	var answer int

	//Loop until player loses
	for !playerLost {

		//Setting up randomization
		var randomNum int = rand.Intn(2)
		var simonNumber int = rand.Intn(100)

		//Assigns random number for simon says or not
		if randomNum == 0 {
			fmt.Println("Simon says:", simonNumber)
		} else {
			fmt.Println(simonNumber)
		}

		//Scanning player's answer
		fmt.Print("Your answer: ")
		fmt.Scanln(&answer)

		// Check the player's input
		if (randomNum == 0 && answer == simonNumber) || (randomNum == 1 && answer != simonNumber) {
			// Correct answer; continue the game
			continue
		} else {
			// Incorrect answer; end the game
			fmt.Println("You lost! Thanks for playing!")
			playerLost = true
		}
	}
}
