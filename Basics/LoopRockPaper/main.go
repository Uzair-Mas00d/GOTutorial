package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("Loop: ", i)
	}

	for i := 0; i < 9; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Odd: ", i)
	}

	i := 0
	for {
		fmt.Println("Infinite Loop", i)
		i++

		if i == 3 {
			break
		}
	}

	y := 0
	for y < 3 {
		fmt.Println("For as a While loop: ", y)
		y++
	}

	const rounds = 3
	fmt.Println("Lets play rock paper scissors you have ", rounds, "rounds")

	for i := 0; i <= rounds; i++ {
		comptureChoiceNum := rand.Intn(rounds)
		var comptureChoice string
		switch comptureChoiceNum {
		case 0:
			comptureChoice = "rock"
		case 1:
			comptureChoice = "paper"
		case 2:
			comptureChoice = "Scissors"
		}

		var playerChoice string
		fmt.Println("Enter your choice (Rock, Paper, Scissors): ")
		fmt.Scanln(&playerChoice)

		fmt.Printf("Compture chose %s\n", comptureChoice)
		switch {
		case playerChoice == comptureChoice:
			fmt.Println("Its a tie")
		case playerChoice == "Rock" && comptureChoice == "Scissors",
			playerChoice == "Paper" && comptureChoice == "Rock",
			playerChoice == "Scissors" && comptureChoice == "Paper":
			fmt.Println("You win this round")
		default:
			fmt.Println("compture wins")
		}
	}
	fmt.Println("Game Over")
}
