package main

import (
	game "GussingGame/Game"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	var playerName string

	fmt.Print("Enter your name: ")
	_, err := fmt.Scanln(&playerName)
	if err != nil {
		fmt.Println("Error reading input. Please restart the game.")
		return
	}
	playerName = strings.ToLower(playerName)
	var results []*game.Result
	for {

		fmt.Printf("Welcome, %s, to the number guessing game!\n", playerName)
		fmt.Println("I am thinking of a number between 1 to 100")
		fmt.Println("**********************************************************************************")
		fmt.Println("Please select the difficulty level:")
		fmt.Println("1. Easy (10 chances)")
		fmt.Println("2. Medium (5 chances)")
		fmt.Println("3. Hard (3 chances)")

		res, err := game.GameStart(playerName)
		if err != nil {
			fmt.Println("Error starting the game:", err)
		}
		fmt.Println(res)

		var playAgain string
		fmt.Println("Do you want to play again")

		_, err = fmt.Scanln(&playAgain)
		if err != nil {
			fmt.Println("Error reading input. Exiting the game.")
			return
		}
	

		results = append(results, res)
		if strings.ToLower(playAgain) != "yes" && strings.ToLower(playAgain) != "y" {
			fmt.Println("Thanks for playing! Goodbye.")
			break
		}

	}

	fmt.Println(results)

	filename := `F:\workspace\Number Gussing Game\result.json`
	file, err := os.ReadFile(filename)
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("file does not exists")
		return
	} else if err != nil {
		return
	}

	var result map[string][]*game.Result
	err = json.Unmarshal(file, &result)
	if err != nil {
		fmt.Println(err)
		return
	}

	nameResult := result[playerName]
	if len(nameResult) == 0 {
		result[playerName] = results
	} else {
		nameResult = append(nameResult, results...)
		result[playerName] = nameResult
	}

	data, err := json.MarshalIndent(result, "", " ")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

}
