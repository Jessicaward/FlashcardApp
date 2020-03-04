package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const FlashcardFilePath = "flashcards.txt"

type Flashcard struct {
	Definition string
	Answer     string
}

func main() {
	stream := string(ReadFlashcardStreamFromFile())
	flashcards := ParseFlashcardsFromString(stream)

	quit := false

	for !quit {
		switch GetMenuItem() {
		case 1:
			CreateFlashcard()
		case 2:
			PracticeFlashcards()
		case 3:
			DisplayFlashcards(flashcards)
		case 4:
			quit = true
		}
	}
}

func CreateFlashcard() {
	fmt.Println()
	definition := GetUserInput("Definition: ")
	answer := GetUserInput("Answer: ")
	WriteFlashcardToFile(Flashcard{
		Definition: definition,
		Answer:     answer,
	})
	fmt.Println()
	fmt.Println("Flashcard created")
	fmt.Println()
}

func PracticeFlashcards() {

}

func DisplayFlashcards(flashcards []Flashcard) {
	fmt.Println()
	fmt.Println("~~~~~~~~F L A S H C A R D S~~~~~~~~")
	for _, card := range flashcards {
		fmt.Println("	", card.Definition, " means ", card.Answer)
	}
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println()
	GetUserInput("Finished?")
}

func GetMenuItem() int {
	PrintMenu()
	input := GetUserInput("Numeric option: ")
	option, err := strconv.Atoi(input[0:1])
	HandleError(err)
	return option
}

func PrintMenu() {
	fmt.Println("=======================")
	fmt.Println("        M E N U        ")
	fmt.Println("=======================")
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
	file, err := os.Open(FlashcardFilePath)
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

func WriteFlashcardToFile(flashcard Flashcard) {
	text := BuildFlashcardString(flashcard)
	AppendLineToFile(FlashcardFilePath, text)
}

func BuildFlashcardString(flashcard Flashcard) string {
	return flashcard.Definition + " | " + flashcard.Answer + ","
}

func AppendLineToFile(filepath string, line string) {
	f, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil{
		HandleError(err)
		return
	}
	_, err = fmt.Fprintln(f, line)
	if err != nil {
		HandleError(err)
		f.Close()
		return
	}

	err = f.Close()
	if err != nil{
		HandleError(err)
		return
	}
}