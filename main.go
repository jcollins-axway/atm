package main

import (
	"fmt"

	"github.com/jcollins-axway/atm/atm"
)

func main() {
	// a := atm.InitATM()
	// require initial deposit
	// a.Deposit()

	// preload atm
	a := atm.InitATM(
		atm.WithTwenties(5),
		atm.WithFifties(5),
		atm.WithHundreds(5),
		atm.WithTwoHundreds(5),
		atm.WithFiveHundreds(5),
	)

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
