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

func calculateSum(transactions []int) (int, int, int) {
	income := 0
	outcome := 0
	transactionsSum := 0
	for _, val := range transactions {
		if val < 0 {
			outcome += val
		} else {
			income += val
		}
		transactionsSum += val
	}
	return income, outcome, transactionsSum
}

func printTransactions(income int, outcome int, transactions []int, transactionsSum int) {
	fmt.Printf("Transactions list: %v\n", transactions)
	fmt.Printf("Len of transactions: %d\n", len(transactions))
	fmt.Printf("Sum of transactions: %d\n", transactionsSum)
	fmt.Printf("Income: %d\n", income)
	fmt.Printf("Outcome: %d\n", outcome)
}

func main() {
	transactions := inputTransactions()
	income, outcome, transactionsSum := calculateSum(transactions)
	printTransactions(income, outcome, transactions, transactionsSum)
}
