package main

import (
	"errors"
	"fmt"
	"math/rand"
	"slices"
)

func sequence() func() int {
	i := 0

	return func() int {
		i++
		return i
	}

}

func fib() func() int {
	a, b := 0, 1

	return func() int {
		a, b = b, a+b

		return a
	}
}

// func rockPaperScissor (){
// 	fmt.Println("Lets play rock paper scissors you have 3 rounds")

// 	for i := 0; i <= 3; i++ {
// 		comptureChoiceNum := rand.Intn(3)
// 		var comptureChoice string
// 		switch comptureChoiceNum {
// 		case 0:
// 			comptureChoice = "rock"
// 		case 1:
// 			comptureChoice = "paper"
// 		case 2:
// 			comptureChoice = "Scissors"
// 		}

// 		var playerChoice string
// 		fmt.Println("Enter your choice (Rock, Paper, Scissors): ")
// 		fmt.Scanln(&playerChoice)

// 		if !slices.Contains([]string{"Rock", "Paper", "Scissors"}, playerChoice){
// 			panic("Invalid input.")
// 		}

// 		fmt.Printf("Compture chose %s\n", comptureChoice)
// 		switch {
// 		case playerChoice == comptureChoice:
// 			fmt.Println("Its a tie")
// 		case playerChoice == "Rock" && comptureChoice == "Scissors",
// 			playerChoice == "Paper" && comptureChoice == "Rock",
// 			playerChoice == "Scissors" && comptureChoice == "Paper":
// 			fmt.Println("You win this round")
// 		default:
// 			fmt.Println("compture wins")
// 		}
// 	}
// 	fmt.Println("Game Over")
// }

func rockPaperScissor () error{
	fmt.Println("Lets play rock paper scissors you have 3 rounds")

	for i := 0; i <= 3; i++ {
		comptureChoiceNum := rand.Intn(3)
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

		if !slices.Contains([]string{"Rock", "Paper", "Scissors"}, playerChoice){
			return errors.New("invalid input")
		}

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

	return nil
}

func main() {
	nextInt := sequence()
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	f := fib()
	fmt.Println(f(), f(), f(), f(), f())

	// rockPaperScissor()

	err := rockPaperScissor()
	if err != nil {
		fmt.Println("Error:", err)
	}
}