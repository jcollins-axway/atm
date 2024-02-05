package note

import "fmt"

type Note interface {
	Add(int)
	GetCount() int
	GetTotal() int
	GetValue() int
	GetStagedTotal() int
	GetStagedCount() int
	StageWithdraw(int) (func(), func(), error)
}

type note struct {
	value  int
	count  int
	staged int
}

func NewNote(value int) Note {
	return &note{value: value}
}

func (n *note) Add(i int) {
	n.count += i
}

func (n *note) execute() {
	if n.staged == 0 {
		return
	}
	fmt.Printf("$%v:  %v\n", n.value, n.staged)
	n.count -= n.staged
}

func (n *note) cancel() {
	n.staged = 0
}

func (n *note) needed(amount int) int {
	rem := amount % n.GetValue()
	return (amount - rem) / n.GetValue()
}

func (n *note) GetCount() int {
	return n.count
}

func (n *note) GetTotal() int {
	return n.count * n.value
}

func (n *note) GetValue() int {
	return n.value
}

func (n *note) GetStagedCount() int {
	return n.staged
}

func (n *note) GetStagedTotal() int {
	return n.staged * n.value
}

func (n *note) StageWithdraw(amount int) (func(), func(), error) {
	if n.count == 0 {
		return nil, nil, nil
	}
	i := n.needed(amount)
	if n.count < i {
		return nil, nil, fmt.Errorf("insufficient number of bills")
	}
	n.staged = i
	return n.execute, n.cancel, nil
}
