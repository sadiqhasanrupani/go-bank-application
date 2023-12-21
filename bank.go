package main

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

const balanceFileName = "balance.txt"

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

func getBalanceFile() (float64, error) {
	data, err := os.ReadFile(balanceFileName)

	if err != nil {
		return 5000.0, errors.New("Failed to find the balance file.")
	}

	balanceText := string(data)
	balanceFloat, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 5000.0, errors.New("Failed to parse stored balance value.")
	}

	return balanceFloat, nil
}

func writeBalanceToFile(balance float64) {
	amount := fmt.Sprint(balance)

	//^ first param => text's name
	//^ second param => array of byte
	//^ third param => file permissions
	os.WriteFile(balanceFileName, []byte(amount), 0644)
}

func main() {
	accountBalance, err := getBalanceFile()

	if err != nil {
		fmt.Println("\nGot an Error")
		fmt.Println(err)
		fmt.Println("------------------------")
	}

	outputNextLine("\n" + `Welcome to "Go Bank"! 🏦`)

	for {
		outputNextLine("\nWhat do you want to do?")
		outputNextLine("1. Check Balance")
		outputNextLine(`2. Deposit Money`)
		outputNextLine(`3. Withdraw Money`)
		outputNextLine(`4. Exit`)

		choice := getIntUserInput("\nYour choice: ")

		switch choice {
		case 1:
			fmt.Printf("\nYour balance is: %v\n", accountBalance)
		case 2:
			deposit := getFloatUserInput("\nYour Deposit: ")

			if deposit <= 0 {
				outputNextLine("\nInvalid amount. value must be greater than 0.")
				continue
			}

			accountBalance += deposit
			fmt.Printf("\nBalance Updated! New amount is: %v\n", accountBalance)
			writeBalanceToFile(accountBalance)

		case 3:
			withdraw := getFloatUserInput("\nYour Withdraw Money: ")

			t := reflect.TypeOf(withdraw)

			if withdraw <= 0 || t == reflect.TypeOf("") {
				outputNextLine("\nInvalid amount. value must be greater than 0.")
				continue
			}

			if withdraw > accountBalance {
				outputNextLine("\nInvalid amount. You can't withdraw more than you have.")
				continue
			}

			if t == reflect.TypeOf(0) || t == reflect.TypeOf(0.0) {
				accountBalance -= withdraw
			}

			fmt.Printf("\nBalance Updated! New amount is: %v\n", accountBalance)
			writeBalanceToFile(accountBalance)
		default:
			outputNextLine("\nAdieu! 🌟 Return for more adventures! 🚀")
			return
		}

	}
}
