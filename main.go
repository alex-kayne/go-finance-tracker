package main

import "fmt"

func inputTransactions() (transactions []int) {
	var input int
	for {
		fmt.Println("Enter amount of transactions or 0 if you want to exit")
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println(err)
			break
		}
		if input == 0 {
			fmt.Println("End of transactions")
			break
		}
		transactions = append(transactions, input)
	}
	return transactions
}

func calculateSum(transactions []int) (transactionsSum int) {
	for _, val := range transactions {
		transactionsSum += val
	}
	return transactionsSum
}

func printTransactions(transactions []int, transactionsSum int) {
	fmt.Printf("Transactions list: %v\n", transactions)
	fmt.Printf("Len of transactions: %d\n", len(transactions))
	fmt.Printf("Sum of transactions: %d\n", transactionsSum)
}

func main() {
	transactions := inputTransactions()
	transactionsSum := calculateSum(transactions)
	printTransactions(transactions, transactionsSum)
}
