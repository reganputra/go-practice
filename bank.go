package main

import "fmt"

func main() {

	var accountBalance float64

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
			fmt.Printf("Your current balance is: $%.2f\n", accountBalance)
		case 2:
			var depositAmount float64
			fmt.Print("Enter the amount to deposit: $")
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Deposit amount must be greater than 0.")
				continue
			}

			fmt.Println("Depositing money...")
			accountBalance += depositAmount
			fmt.Printf("Your new balance is: $%.2f\n", accountBalance)
		case 3:
			var withdrawAmount float64
			fmt.Print("Enter the amount to withdraw: $")
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
				fmt.Printf("You withdrew $%.2f\n", withdrawAmount)
				fmt.Printf("Your new balance is: $%.2f\n", accountBalance)
			}
		case 4:
			fmt.Println("Exiting the system. Thank you for using our services!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")

		}

	}
}
