package option

import "example.com/go-bank-application/utils"

func GetOptions() {
	utils.OutputNextLine("\nWhat do you want to do?")
	utils.OutputNextLine("1. Check Balance")
	utils.OutputNextLine(`2. Deposit Money`)
	utils.OutputNextLine(`3. Withdraw Money`)
	utils.OutputNextLine(`4. Exit`)
}
