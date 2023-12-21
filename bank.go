package main

import (
	"fmt"
)

func main() {
	var accountBalance float64 = 6000

	for {
		outputNextLine("\n" + `Welcome to "Go Bank"! 🏦`)
		outputNextLine("\nWhat do you want to do?")
		outputNextLine("\n1. Check Balance")
		outputNextLine(`2. Deposit Money`)
		outputNextLine(`3. Withdraw Money`)
		outputNextLine(`4. Exit`)

		choice := getIntUserInput("\nYour choice: ")

		if choice == 1 {
			fmt.Printf("\nYour balance is: %v\n", accountBalance)
		} else if choice == 2 {
			deposit := getFloatUserInput("\nYour Deposit: ")

			if deposit <= 0 {
				outputNextLine("\nInvalid amount. value must be greater than 0.")
				return
			}

			accountBalance += deposit
			fmt.Printf("\nBalance Updated! New amount is: %v\n", accountBalance)
		} else if choice == 3 {
			withdraw := getFloatUserInput("\nYour Withdraw Money: ")
			accountBalance -= withdraw

			if withdraw <= 0 {
				outputNextLine("\nInvalid amount. value must be greater than 0.")
				return
			}

			if withdraw > accountBalance {
				outputNextLine("\nInvalid amount. You can't withdraw more than you have.")
				return
			}

			fmt.Printf("\nBalance Updated! New amount is: %v\n", accountBalance)
		} else if choice == 4 {
			outputNextLine("\nAdieu! 🌟 Return for more adventures! 🚀")
			break
		}
	}
}

func outputNextLine(text string) {
	fmt.Println(text)
}

func outputLine(text string) {
	fmt.Print(text)
}

func getFloatUserInput(message string) float64 {
	var userInput float64
	outputLine(message)
	fmt.Scan(&userInput)

	return userInput
}

func getIntUserInput(message string) int {
	var userInput int
	outputLine(message)
	fmt.Scan(&userInput)

	return userInput
}
