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

type WordleState struct {
	word string
	turn int
	done bool
}

func main() {
	playGame()
}

func playGame() {
	words, err := loadWords()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Welcome to Wordle!")
	state := &WordleState{word: randomWord(words)}

	for state.turn < maxTurns && !state.done {
		doTurn(state)
	}

	if !state.done {
		color.Red("You lose! The word was: %s", state.word)
	}
}

func loadWords() ([]string, error) {
	data, err := os.ReadFile("words")
	if err != nil {
		return nil, err
	}
	return strings.Split(string(data), "\n"), nil
}

func randomWord(words []string) string {
	rand.Seed(time.Now().UnixNano())
	return words[rand.Intn(len(words))]
}

func doTurn(state *WordleState) {
	input, err := requestInput(state)
	if err != nil {
		fmt.Println(err)
		state.done = true
		return
	}

	if len(input) != 5 {
		fmt.Println("Invalid, word should be 5 characters long")
		return
	}

	if input == state.word {
		color.Green("You win! The word was: %s", state.word)
		state.done = true
		return
	}

	for index, char := range input {
		if uint8(char) == state.word[index] {
			fmt.Print(color.HiGreenString("%c", char))
		} else if strings.ContainsRune(state.word, char) {
			fmt.Print(color.HiYellowString("%c", char))
		} else {
			fmt.Print(color.WhiteString("%c", char))
		}
	}
	fmt.Println()

	state.turn++
}

func requestInput(state *WordleState) (string, error) {
	fmt.Printf("Turn %d, guess: ", state.turn)

	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		return "", err
	}

	return input, nil
}
