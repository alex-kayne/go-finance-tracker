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

	partialTransactions := transactions[1:4]

	fmt.Printf("Partial Transactions cap: %d\n", cap(partialTransactions))
	fmt.Printf("Partial Transactions len: %d\n", len(partialTransactions))

	//изменится потому что, слайс хранит адрес изначальной ячейки в памяти
	partialTransactions[0] = 9999
	printTransactions(transactions[:])

	newPartialTransactions := transactions[1:3]

	fmt.Printf("New Partial Transactions cap: %d\n", cap(newPartialTransactions))
	fmt.Printf("New Partial Transactions len: %d\n", len(newPartialTransactions))

	partialTransactions[1] = -9999

	fmt.Printf("New Partial Transactions: %d\n", newPartialTransactions)
	fmt.Printf("Transactions: %d\n", transactions)
	fmt.Printf("Partial Transactions: %d\n", partialTransactions)

	allTransactions := transactions[:]

	fmt.Printf("%T\n", allTransactions)

	fmt.Printf("Balance: %d\n", getBalance(transactions[:]))
}
