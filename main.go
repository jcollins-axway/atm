package main

import (
	"fmt"

	"github.com/jcollins-axway/atm/atm"
)

func main() {
	a := atm.InitATM()

	a.Deposit()

	for {
		var action string

		fmt.Print(`
What would you like to do?
	Balance  = b
	Deposit  = d
	Withdraw = w

Enter: `)

		fmt.Scanln(&action)

		switch action {
		case "b":
			a.Balance()
		case "d":
			a.Deposit()
		case "w":
			a.Withdraw()
		default:
			fmt.Println("invalid option, select again")
		}

		fmt.Println("----------------------------------")
	}
}
