package main

import "fmt"

func main() {
	transactions := [5]int{5000, -1200, -350, 2000, -500}
	fmt.Printf("Transactions: %d\n", transactions)
	fmt.Printf("First transaction: %d\n", transactions[0])
	fmt.Printf("Last transaction: %d\n", transactions[len(transactions)-1])
	fmt.Printf("Transactions count: %d\n", len(transactions))

	balance := 0

	for i := 0; i < len(transactions); i++ {
		balance += transactions[i]
	}

	fmt.Printf("Balance: %d\n", balance)
}
