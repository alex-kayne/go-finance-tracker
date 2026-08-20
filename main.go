package main

import "fmt"

func printTransactions(transactions []int) {
	fmt.Printf("Transactions: %v\n", transactions)
	fmt.Printf("First transaction: %d\n", transactions[0])
	fmt.Printf("Last transaction: %d\n", transactions[len(transactions)-1])
	fmt.Printf("Transactions count: %d\n", len(transactions))
}

func getBalance(transactions []int) int {
	balance := 0

	for i := 0; i < len(transactions); i++ {
		balance += transactions[i]
	}
	return balance
}

func main() {
	var transactions [5]int
	transactions[0] = 5000
	transactions[1] = -1200
	transactions[2] = -350
	transactions[3] = 2000
	transactions[4] = -500

	transactions[2] = 4085

	printTransactions(transactions[:])

	fmt.Printf("Balance: %d\n", getBalance(transactions[:]))
}
