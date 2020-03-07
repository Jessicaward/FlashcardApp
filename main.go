package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

const FlashcardFilePath = "flashcards.txt"

type Flashcard struct {
	Definition string
	Answer     string
}

type UserFlashcard struct {
	Definition string
	Answer 	   string
	UserAnswer string
}

func main() {
	quit := false

	for !quit {
		//Flashcards should be loaded on each menu "tick"
		flashcards := GetFlashcards()
		switch GetMenuItem() {
		case 1:
			CreateFlashcard()
		case 2:
			PracticeFlashcards(flashcards)
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

func GetFlashcards() []Flashcard {
	stream := string(ReadFlashcardStreamFromFile())
	return ParseFlashcardsFromString(stream)
}

func PracticeFlashcards(flashcards []Flashcard) {
	//todo: this should keep track of the user's incorrect answers, not just which answer was incorrect
	numberOfQuestions, err := strconv.Atoi(strings.Trim(GetUserInput("Enter the number of questions you would like: "), "\r\n"))
	var correctFlashcards []Flashcard
	var incorrectFlashcards []UserFlashcard

	HandleError(err)

	if numberOfQuestions == 0 {
		return
	}

	for i := 0; i < numberOfQuestions; i++ {
		randomFlashcard := flashcards[rand.Intn(len(flashcards))]
		result, answer := AskQuestion(randomFlashcard)
		if result {
			correctFlashcards = append(correctFlashcards, randomFlashcard)
		} else {
			incorrectFlashcards = append(incorrectFlashcards, ConvertFlashcardToUserFlashcard(randomFlashcard, answer))
		}
	}

	ShowGameReport(correctFlashcards, incorrectFlashcards)
}

func ConvertFlashcardToUserFlashcard(flashcard Flashcard, userAnswer string) UserFlashcard {
	return UserFlashcard {
		Answer: flashcard.Answer,
		Definition: flashcard.Definition,
		UserAnswer: userAnswer,
	}
}

///This returns whether the user correctly answered the question or not
func AskQuestion(flashcard Flashcard) (bool, string) {
	fmt.Println("Definition: ", flashcard.Definition)
	userAnswer := GetUserInput("Answer: ")
	return CheckAnswer(flashcard.Answer, userAnswer), userAnswer
}

func CheckAnswer(correctAnswer string, userAnswer string) bool {
	return strings.EqualFold(strings.TrimSpace(correctAnswer), strings.TrimSpace(userAnswer))
}

func ShowGameReport(correctFlashcards []Flashcard, incorrectFlashcards []UserFlashcard) {
	fmt.Println()
	fmt.Println()
	fmt.Println("~~~~~~~~~~Report~~~~~~~~~~")
	fmt.Println()
	fmt.Println("You got ", len(correctFlashcards), " correct and ", len(incorrectFlashcards), " incorrect.")
	if len(correctFlashcards) > 0 {
		fmt.Println("~~~Correct Answers~~~")
		for _, correct := range correctFlashcards {
			fmt.Println("	Definition: ", correct.Definition)
			fmt.Println("	Answer: ", correct.Answer)
		}
	}
	if len(incorrectFlashcards) > 0 {
		fmt.Println("~~~Incorrect answers~~~")
		for _, incorrect := range incorrectFlashcards {
			fmt.Println("	Definition: ", incorrect.Definition)
			fmt.Println("	Answer: ", incorrect.Answer)
			fmt.Println("   Your answer: ", incorrect.UserAnswer)
		}
	}
	fmt.Println()
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println()
	fmt.Println()
}

func DisplayFlashcards(flashcards []Flashcard) {
	fmt.Println()
	fmt.Println("~~~~~~~~~~~~~~~~~~ F L A S H C A R D S ~~~~~~~~~~~~~~~~~~")
	for _, card := range flashcards {
		fmt.Println(" ")
		for i := 0; i < 25-len(card.Definition); i++ {
			//Formatting - pad left
			fmt.Print(" ")
		}
		fmt.Printf(card.Definition)
		fmt.Print(" means ", card.Answer)
	}
	fmt.Println()
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
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
	HandleError(err)
	_, err = fmt.Fprintln(f, line)
	HandleError(err)
	err = f.Close()
	HandleError(err)
}
