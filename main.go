package main

import "fmt"

func inputTransactions() (transactions []float32) {
	var input float32
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

func printTransactions(transactions []float32) {
	fmt.Printf("Transactions list: %v\n", transactions)
	var sum float32
	for _, val := range transactions {
		sum += val
	}
	fmt.Printf("Len of transactions: %d\n", len(transactions))
	fmt.Printf("Sum of transactions: %f\n", sum)
}

func main() {
	transactions := inputTransactions()
	printTransactions(transactions)
}
