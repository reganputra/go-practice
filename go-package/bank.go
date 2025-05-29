package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const balanceFileName = "account.txt"

func readBalanceFile() (float64, error) {
	read, err := os.ReadFile(balanceFileName)
	if err != nil {
		fmt.Println("Initializing balance to 0.0")
		return 0.0, errors.New("error reading balance file")
	}
	balanceText := string(read)
	convert, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		fmt.Println("Initializing balance to 0.0")
		return 0.0, errors.New("error converting balance to float")
	}
	return convert, nil
}

func writeBalanceToFiles(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(balanceFileName, []byte(balanceText), 0644)
}

func main() {

	var accountBalance, err = readBalanceFile()

	if err != nil {
		fmt.Println(err)
		//panic("Please create an account.txt file with a valid balance.")
	}

	fmt.Println("Welcome to the Bank Management System!")
	for {

		options()

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
