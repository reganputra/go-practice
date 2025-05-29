package main

import (
	"fmt"
	"github.com/Pallinder/go-randomdata"
	"go-package/fileops"
)

const balanceFileName = "account.txt"

func main() {

	var accountBalance, err = fileops.ReadBalanceFile(balanceFileName)

	if err != nil {
		fmt.Println(err)
		//panic("Please create an account.txt file with a valid balance.")
	}

	fmt.Println("Welcome", randomdata.SillyName())
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
			fileops.WriteBalanceToFiles(accountBalance, balanceFileName)
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
				fileops.WriteBalanceToFiles(accountBalance, balanceFileName)
			}
		default:
			fmt.Println("Exiting the system. Thank you for using our services!")
			return
		}

	}
}
