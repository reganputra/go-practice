package main

import (
	"fmt"
	"os"
	"strconv"
)

const balanceFileName = "account.txt"

func readBalanceFile() float64 {
	read, _ := os.ReadFile(balanceFileName)
	balanceText := string(read)
	convert, _ := strconv.ParseFloat(balanceText, 64)

	return convert
}

func writeBalanceToFiles(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(balanceFileName, []byte(balanceText), 0644)
}

func main() {

	var accountBalance = readBalanceFile()

	fmt.Println("Welcome to the Bank Management System!")
	for {

		fmt.Println("What would you like to do today?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Checking balance...")
			fmt.Println(accountBalance)
		case 2:
			var depositAmount float64
			fmt.Print("Enter the amount to deposit: ")
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Deposit amount must be greater than 0.")
				continue
			}

			fmt.Println("Depositing money...")
			accountBalance += depositAmount
			fmt.Println(accountBalance)
			writeBalanceToFiles(accountBalance)
		case 3:
			var withdrawAmount float64
			fmt.Print("Enter the amount to withdraw: ")
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Withdraw amount must be greater than 0.")
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Println("Insufficient funds for this withdrawal.")
			}
			{
				fmt.Println("Withdrawing money...")
				accountBalance -= withdrawAmount
				fmt.Println(withdrawAmount)
				fmt.Println(accountBalance)
				writeBalanceToFiles(accountBalance)
			}
		default:
			fmt.Println("Exiting the system. Thank you for using our services!")
			return
		}

	}
}
