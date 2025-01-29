package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Create a game of rock paper scissors with Go and user input through the
command line.

Once you run the game, the program should prompt the player to select from
rock, paper, or scissors by typing in text and hitting enter. The program should
be capable of selecting one as well. If the player enters an invalid selection,
prompt them for a new one. The player should be able to type in “exit” to end
the program.

The game should then compare the responses, and determine who won, or if
there was a tie. The next round should then begin without restarting the
program.

Once this is done, you must add a creative “twist” of your own to this simple
formula. Try to make the basic game of rock paper scissors as fun and
interesting as you can. It is up to you to decide how the specics work and how
the game is structured. Good luck!
*/

func main() {

	//Seed the random variable
	rand.Seed(time.Now().UnixNano())

	//Set up variables
	var playerExit bool = false
	var playerChoice string
	var randomNum int
	var NPCchoice string
	var roundNumber int = 1
	var count int = 0

	//Introductory message
	fmt.Println("Welcome to Rock Paper Scissors! Make your first choice! >:)")
	fmt.Println("Type in your choice! Rock, paper, or scissors? ")
	fmt.Println("Type in \"Exit\" if you don't want to play anymore.")

	//Set up loop to exit only on player typing exit or Exit
	for !playerExit {

		//New round indicator
		fmt.Println("Round:", roundNumber)

		//Set up "twist" if count is greater than or equal to 3 then "super move" is initiated and you skip a round
		if count >= 3 {
			fmt.Println("You've won 3 times in a row! Super move has been initiated! You skip a round.")
			roundNumber += 1
			count = 0
			continue
		}

		//Set up random number for npc choices
		randomNum = rand.Intn(3)

		//Set up conditionals for each npc choice
		if randomNum == 0 {
			NPCchoice = "rock"
		} else if randomNum == 1 {
			NPCchoice = "paper"
		} else if randomNum == 2 {
			NPCchoice = "scissors"
		}

		//Scan for player choice
		fmt.Print("Your choice: ")
		fmt.Scanln(&playerChoice)

		//Conditional in case of wrong answer or type
		if playerChoice != "rock" && playerChoice != "Rock" && playerChoice != "paper" && playerChoice != "Paper" && playerChoice != "scissors" && playerChoice != "Scissors" {
			fmt.Println("Oops! Not an actual answer in the game. Please answer again!")
			continue
		}

		//Set up conditionals
		if playerChoice == NPCchoice {
			fmt.Println("Computer's choice:", NPCchoice)
			fmt.Println("Tie! Next Round! >:)")
			roundNumber += 1
			continue
		} else if (playerChoice == "rock" || playerChoice == "Rock") && NPCchoice == "paper" {
			fmt.Println("Computer's choice:", NPCchoice)
			fmt.Println("You lost! Better luck next round!")
			break
		} else if (playerChoice == "rock" || playerChoice == "Rock") && NPCchoice == "scissors" {
			fmt.Println("Computer's choice:", NPCchoice)
			fmt.Println("You win! Continue to next round!")
			roundNumber += 1
			count += 1
			continue
		} else if (playerChoice == "paper" || playerChoice == "Paper") && NPCchoice == "scissors" {
			fmt.Println("Computer's choice:", NPCchoice)
			fmt.Println("You lost! Better luck next round!")
			break
		} else if (playerChoice == "paper" || playerChoice == "Paper") && NPCchoice == "rock" {
			fmt.Println("Computer's choice:", NPCchoice)
			fmt.Println("You win! Continue to next round!")
			roundNumber += 1
			count += 1
			continue
		} else if (playerChoice == "scissors" || playerChoice == "Scissors") && NPCchoice == "rock" {
			fmt.Println("Computer's choice:", NPCchoice)
			fmt.Println("You lost! Better luck next round!")
			break
		} else if (playerChoice == "scissors" || playerChoice == "Scissors") && NPCchoice == "paper" {
			fmt.Println("Computer's choice:", NPCchoice)
			fmt.Println("You win! Continue to next round!")
			roundNumber += 1
			count += 1
			continue
		}

		if playerChoice == "Exit" || playerChoice == "exit" {
			playerExit = true
		}
	}

}
