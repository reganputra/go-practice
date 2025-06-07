package cmdmanager

import "fmt"

type CommandManager struct{}

func (cmd CommandManager) ReadLine() ([]string, error) {
	fmt.Println("Please enter your prices. Confirm every price with ENTER")

	var prices []string

	for {
		var price string
		fmt.Print("Price: ")
		fmt.Scan(&price)

		if price == "0" {
			break
		}
		prices = append(prices, price)
	}

	return prices, nil
}

func (cmd CommandManager) WriteResult(data interface{}) error {
	fmt.Println(data)
	return nil
}

func New() CommandManager {
	return CommandManager{}
}
