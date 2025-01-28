package main

import (
	"fmt"
)

func main() {
	var myLevel int
	var xp int

	myLevel, xp = AwardXP(1, 50, 77)

	fmt.Println("My Level:", myLevel, " My current XP: ", xp)
}

func AwardXP(currentLevel int, currentXP int, awardedXP int) (int, int) {

	//In case awarded XP is negative
	if awardedXP <= 0 {
		fmt.Println("ERROR: No negative XP allowed!!!")
		return currentLevel, currentXP
	}

	//Calculate total XP adding current and awarded XP
	var totalXP int = currentXP + awardedXP

	//Add the levels with each 100 XP reached and makes leftover XP
	currentLevel += totalXP / 100
	var leftoverXP int = totalXP % 100

	//Return current level and leftover XP
	return currentLevel, leftoverXP
}
