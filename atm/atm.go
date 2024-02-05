package atm

import (
	"fmt"
	"strconv"

	"github.com/jcollins-axway/atm/atm/note"
)

type ATM interface {
	Balance()
	Deposit()
	Withdraw()
}

type atm struct {
	xx    note.Note
	l     note.Note
	c     note.Note
	cc    note.Note
	d     note.Note
	notes []note.Note
}

type Opts func(*atm)

func InitATM(opts ...Opts) ATM {
	a := &atm{
		xx: note.NewNote(20),
		l:  note.NewNote(50),
		c:  note.NewNote(100),
		cc: note.NewNote(200),
		d:  note.NewNote(500),
	}
	a.notes = []note.Note{a.xx, a.l, a.c, a.cc, a.d}

	for _, o := range opts {
		o(a)
	}

	return a
}

func WithTwenties(i int) Opts {
	return func(a *atm) {
		a.xx.Add(i)
	}
}

func WithFifties(i int) Opts {
	return func(a *atm) {
		a.l.Add(i)
	}
}

func WithHundreds(i int) Opts {
	return func(a *atm) {
		a.c.Add(i)
	}
}

func WithTwoHundreds(i int) Opts {
	return func(a *atm) {
		a.cc.Add(i)
	}
}

func WithFiveHundreds(i int) Opts {
	return func(a *atm) {
		a.d.Add(i)
	}
}

func (a *atm) getBalance() int {
	total := 0

	for _, n := range a.notes {
		total += n.GetTotal()
	}
	return total
}

func (a *atm) Balance() {
	bal := a.getBalance()
	balOut := fmt.Sprintf("Your current balance is: $%v\n", bal)
	for _, n := range a.notes {
		balOut += fmt.Sprintf("\t$%v: %v\n", n.GetValue(), n.GetCount())
	}
	fmt.Println(balOut)
}

func (a *atm) Deposit() {
	fmt.Println("Begin deposit:")

	for _, n := range a.notes {
		depositNote(n)
	}
}

func (a *atm) Withdraw() {
	var amountStr string
	var amount int

	fmt.Printf("How much to withdraw? ")
	fmt.Scanln(&amountStr)

	for {
		var err error
		amount, err = strconv.Atoi(amountStr)
		if err != nil {
			fmt.Printf("Please enter a valid amount: ")
			fmt.Scanln(&amountStr)
			continue
		}
		break
	}
	a.withdrawFunds(amount)
}

func (a *atm) withdrawFunds(amount int) {
	cancelAll := []func(){}
	remaining := amount

	for i := len(a.notes) - 1; i >= 0; i-- {
		n := a.notes[i]
		exe, cancel, err := n.StageWithdraw(remaining)
		if err != nil {
			for _, c := range cancelAll {
				c()
			}

			fmt.Printf("INVALID AMOUNT: Insufficient $%v notes", n.GetValue())
			return
		} else if exe != nil {
			remaining -= n.GetStagedTotal()
			cancelAll = append(cancelAll, cancel)
			defer exe()
		}
	}
	fmt.Println("Withdrawing:")
}

func depositNote(n note.Note) int {
	var numStr string
	var number int

	fmt.Printf("How many $%v notes? ", n.GetValue())
	fmt.Scanln(&numStr)

	for {
		var err error
		number, err = strconv.Atoi(numStr)
		if err != nil {
			fmt.Printf("Please enter the number of $%v notes? ", n.GetValue())
			fmt.Scanln(&numStr)
			continue
		}
		if number < 0 {
			fmt.Printf("Please enter a positive number for $%v notes? ", n.GetValue())
			fmt.Scanln(&numStr)
			continue
		}
		break
	}

	n.Add(number)
	return n.GetValue() * number
}
