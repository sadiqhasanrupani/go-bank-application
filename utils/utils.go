package utils

import "fmt"

func OutputNextLine(text string) {
	fmt.Println(text)
}

func OutputLine(text string) {
	fmt.Print(text)
}

func GetFloatUserInput(message string) float64 {
	var userInput float64
	OutputLine(message)
	fmt.Scan(&userInput)

	return userInput
}

func GetIntUserInput(message string) int {
	var userInput int
	OutputLine(message)
	fmt.Scan(&userInput)

	return userInput
}
