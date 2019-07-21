package bank

import "fmt"

var deposits = make(chan int)
var balances = make(chan int)
var withdraw = make(chan int)
var withdrawFailure = make(chan bool)

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) {
	if Balance() > amount {
		withdraw <- amount
	} else {
		withdrawFailure <- true
	}
}

func teller() int {
	var balance int

	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case <-withdrawFailure:
			fmt.Println("Could not complete withdraw operation")
		}
	}
}

func init() {
	go teller()
}
