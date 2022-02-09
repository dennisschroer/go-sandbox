package main

import (
	"fmt"
	"github.com/fatih/color"
	"math/rand"
	"os"
	"strings"
	"time"
)

const maxTurns = 6

func main() {
	fmt.Println("Welcome to Wordle!")

	// Read words
	data, err := os.ReadFile("words")
	if err != nil {
		fmt.Println(err)
		return
	}
	words := strings.Split(string(data), "\n")

	// Select a random word
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(words))
	word := words[randomIndex]
	//fmt.Printf("Debug; The random index is %d, the word is: %s\n", randomIndex, word)

	// Do turns
	turn := 0
	var input string

	for turn < maxTurns {
		fmt.Printf("Turn %d, guess: ", turn)
		_, err := fmt.Scanln(&input)

		if err != nil {
			fmt.Println(err)
			return
		}

		if len(input) != 5 {
			fmt.Println("Invalid, word should be 5 characters long")
			continue
		}

		if input == word {
			color.Green("You win! The word was: %s", word)
			return
		}

		for index, char := range input {
			if uint8(char) == word[index] {
				fmt.Print(color.HiGreenString("%c", char))
			} else if strings.ContainsRune(word, char) {
				fmt.Print(color.HiYellowString("%c", char))
			} else {
				fmt.Print(color.WhiteString("%c", char))
			}
		}
		fmt.Println()

		turn++
	}

	color.Red("You lose! The word was: %s", word)
}
