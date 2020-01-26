package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Flashcard struct {
	Definition string
	Answer     string
}

func main() {
	flashcards := string(ReadFlashcardStreamFromFile())
	ParseFlashcardsFromString(flashcards)

	quit := true

	for !quit {
		switch GetMenuItem() {
		case 1:
			CreateFlashcard()
		case 2:
			PracticeFlashcards()
		case 3:
			DisplayFlashcards()
		case 4:
			quit = true
		}
	}
}

func CreateFlashcard() {

}

func PracticeFlashcards() {

}

func DisplayFlashcards() {

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
	fmt.Println("3. Display all flashcards")
	fmt.Println("4. Quit program")
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

func ReadFlashcardStreamFromFile() []byte {
	file, err := os.Open("flashcards.txt")
	HandleError(err)
	defer file.Close()
	b, err := ioutil.ReadAll(file)
	return b
}

func ParseFlashcardsFromString(input string) []Flashcard {
	lines := strings.Split(input, ",")
	var flashcards []Flashcard
	for line := 0; line < len(lines)-1; line++ {
		flashcards = append(flashcards, ParseSingleFlashcard(lines[line]))
	}
	return flashcards
}

func ParseSingleFlashcard(input string) Flashcard {
	parts := strings.Split(input, "|")
	if parts != nil && len(parts) > 0 {
		return Flashcard{
			Definition: strings.TrimSpace(parts[0]),
			Answer:     strings.TrimSpace(parts[1]),
		}
	}
	fmt.Println("Error: empty string was passed into ParseSingleFlashcard")
	var f Flashcard
	return f
}
