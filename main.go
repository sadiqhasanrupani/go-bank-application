package main

import (
	"fmt"
	"reflect"

	"example.com/go-bank-application/fileOperation"
	"example.com/go-bank-application/option"
	"example.com/go-bank-application/utils"
	"github.com/Pallinder/go-randomdata"
)

const balanceFileName = "balance.txt"

func main() {
	accountBalance, err := fileOperation.GetFloatFromFile(balanceFileName)

	if err != nil {
		fmt.Println("\nGot an Error")
		fmt.Println(err)
		fmt.Println("------------------------")
	}

	utils.OutputNextLine("\n" + `Welcome to "Go Bank"! 🏦`)
	fmt.Println(`You can reach us 24/7 `, randomdata.PhoneNumber())

	for {
		option.GetOptions()

		choice := utils.GetIntUserInput("\nYour choice: ")

		switch choice {
		case 1:
			fmt.Printf("\nYour balance is: %v\n", accountBalance)
		case 2:
			deposit := utils.GetFloatUserInput("\nYour Deposit: ")

			if deposit <= 0 {
				utils.OutputNextLine("\nInvalid amount. value must be greater than 0.")
				continue
			}

			accountBalance += deposit
			fmt.Printf("\nBalance Updated! New amount is: %v\n", accountBalance)
			fileOperation.WriteFloatToFile(balanceFileName, accountBalance)

		case 3:
			withdraw := utils.GetFloatUserInput("\nYour Withdraw Money: ")

			t := reflect.TypeOf(withdraw)

			if withdraw <= 0 || t == reflect.TypeOf("") {
				utils.OutputNextLine("\nInvalid amount. value must be greater than 0.")
				continue
			}

			if withdraw > accountBalance {
				utils.OutputNextLine("\nInvalid amount. You can't withdraw more than you have.")
				continue
			}

			if t == reflect.TypeOf(0) || t == reflect.TypeOf(0.0) {
				accountBalance -= withdraw
			}

			fmt.Printf("\nBalance Updated! New amount is: %v\n", accountBalance)
			fileOperation.WriteFloatToFile(balanceFileName, accountBalance)
		default:
			utils.OutputNextLine("\nAdieu! 🌟 Return for more adventures! 🚀")
			return
		}

	}
}
