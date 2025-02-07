package game

import (
	"fmt"
	"math/rand"
)

type Result struct {
	Success bool `json:"success"`
	Attempt int  `json:"attempt,omitempty"`
	Level   int  `json:"level,omitempty"`
}

func GameStart(playerName string) (*Result, error) {
	var inputLevel string
	var level int

	_, err := fmt.Scanln(&inputLevel)
	if err != nil {
		fmt.Println("Please select between 1 to 3")
		return nil, err
	}
	switch inputLevel {
	case "1", "easy":
		fmt.Println("Great you have selected easy difficulty level \n Let's start the game")
		level = 10
	case "2", "medium":
		fmt.Println("Great you have selected medium difficulty level \n Let's start the game")
		level = 5
	case "3", "Hard":
		fmt.Println("Great you have selected Hard difficulty level \n Let's start the game")
		level = 3
	default:
		fmt.Println("Please select  between 1 to 3")
	}

	target := rand.Intn(100) + 1

	fmt.Println("Guess a number between 1 to 100", playerName)
	for i := 0; i < level; i++ {

		var guess int
		fmt.Printf("Attempt %d/%d: Enter your guess: ", i+1, level)

		_, err := fmt.Scanln(&guess)
		if err != nil {
			fmt.Println("Invalid input.")
			return nil, err
		}
		if guess < target {
			fmt.Println("Incorrect! The number is greater than", guess, ".")
		} else if guess > target {
			fmt.Println("Incorrect! The number is less than", guess, ".")
		} else {
			fmt.Printf("Congratulations, %s! You guessed the correct number (%d) in %d attempts!\n", playerName, target, i+1)
			return &Result{
				Success: true,
				Attempt: i + 1,
				Level:   level,
			}, nil
		}
	}

	fmt.Printf("Sorry, %s! You ran out of guesses. The correct number was %d.\n", playerName, target)
	return &Result{
		Success: false,
		Level:   level,
	}, nil

}
