package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	quit := false

	for !quit {
		switch GetMenuItem() {
		case 1:
			CreateFlashcard()
		case 2:
			PracticeFlashcards()
		case 3:
			quit = true
		}
	}
}

func CreateFlashcard() {

}

func PracticeFlashcards() {

}

func GetMenuItem() int {
	PrintMenu()
	input := GetUserInput("Numeric option: ")
	option, err := strconv.Atoi(input)
	HandleError(err)
	return option
}

func PrintMenu() {
	fmt.Println("1. Create new flashcard")
	fmt.Println("2. Practice flashcard")
	fmt.Println("3. Quit program")
}

func GetUserInput(promptText string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(promptText)
	input, _ := reader.ReadString('\n')
	return input
}

func HandleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}
