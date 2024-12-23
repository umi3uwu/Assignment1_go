package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/umi3uwu/Assignment1_go/bank"
)

func main() {
	account := &bank.BankAccount{
		AccountNumber: "123456789",
		HolderName:    "John Doe",
		Balance:       0.0,
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Enter a command: Deposit, Withdraw, Get balance, Exit:")
		if !scanner.Scan() {
			break
		}
		command := strings.ToLower(strings.TrimSpace(scanner.Text()))

		switch command {
		case "deposit":
			fmt.Print("Enter amount to deposit: ")
			if !scanner.Scan() {
				break
			}
			amount, err := strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)
			if err != nil || amount <= 0 {
				fmt.Println("Invalid amount. Please enter a positive number.")
				continue
			}
			account.Deposit(amount)

		case "withdraw":
			fmt.Print("Enter amount to withdraw: ")
			if !scanner.Scan() {
				break
			}
			amount, err := strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)
			if err != nil || amount <= 0 {
				fmt.Println("Invalid amount. Please enter a positive number.")
				continue
			}
			account.Withdraw(amount)

		case "get balance":
			account.GetBalance()

		case "exit":
			fmt.Println("Exiting program...")
			return

		default:
			fmt.Println("Invalid command. Please enter one of the following: Deposit, Withdraw, Get balance, Exit.")
		}
	}
}
