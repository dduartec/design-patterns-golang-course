package main

import "fmt"

var overdraftLimit = -500

type BankAccount struct {
	balance int
}

func (b *BankAccount) Deposit(amount int) {
	b.balance += amount
	fmt.Printf("Deposited %d, balance is now %d \n", amount, b.balance)
}

func (b *BankAccount) Withdraw(amount int) bool {
	if b.balance-amount >= overdraftLimit {
		b.balance -= amount
		fmt.Printf("Withdrew %d, balance is now %d \n", amount, b.balance)
		return true
	}
	return false
}

type Command interface {
	Call()
	Undo()
	Succeeded() bool
	SetSucceeded(val bool)
}

type Action int

const (
	Deposit Action = iota
	Withdraw
)

type BankAccountCommand struct {
	account *BankAccount
	action  Action
	amount  int
	succeed bool
}

func NewBankAccountCommand(ba *BankAccount, action Action, amount int) *BankAccountCommand {
	return &BankAccountCommand{account: ba, action: action, amount: amount}
}

func (b *BankAccountCommand) Call() {
	switch b.action {
	case Deposit:
		b.account.Deposit(b.amount)
		b.succeed = true
	case Withdraw:
		b.succeed = b.account.Withdraw(b.amount)
	}
}

func (b *BankAccountCommand) Undo() {
	if !b.succeed {
		return
	}
	switch b.action {
	case Deposit:
		b.account.Withdraw(b.amount)
	case Withdraw:
		b.account.Deposit(b.amount)
	}

}

func (b *BankAccountCommand) Succeeded() bool {
	return b.succeed
}
func (b *BankAccountCommand) SetSucceeded(val bool) {
	b.succeed = val
}

type CompositeBankAccountCommand struct {
	commands []Command
}

func (c *CompositeBankAccountCommand) Call() {
	for _, cmd := range c.commands {
		cmd.Call()
	}
}

func (c *CompositeBankAccountCommand) Undo() {
	for i := range c.commands {
		c.commands[len(c.commands)-i-1].Undo()
	}
}

func (c *CompositeBankAccountCommand) Succeeded() bool {
	for _, cmd := range c.commands {
		if !cmd.Succeeded() {
			return false
		}
	}
	return true
}

func (c *CompositeBankAccountCommand) SetSucceeded(val bool) {
	for _, cmd := range c.commands {
		cmd.SetSucceeded(val)
	}

}

type MoneyTransferCommand struct {
	CompositeBankAccountCommand
	from, to *BankAccount
	amount   int
}

func NewMoneyTransferCommand(from, to *BankAccount, amount int) *MoneyTransferCommand {
	c := &MoneyTransferCommand{from: from, to: to, amount: amount}
	c.commands = append(c.commands, NewBankAccountCommand(from, Withdraw, amount))
	c.commands = append(c.commands, NewBankAccountCommand(to, Deposit, amount))
	return c
}

func (m *MoneyTransferCommand) Call() {
	ok := true
	for _, cmd := range m.commands {
		if ok {
			cmd.Call()
			ok = cmd.Succeeded()
		} else {
			cmd.SetSucceeded(false)
		}
	}
}

func main() {
	from := BankAccount{100}
	to := BankAccount{}
	mtc := NewMoneyTransferCommand(&from, &to, 50)
	mtc.Call()
	fmt.Println(from, to)

	mtc.Undo()
	fmt.Println(from, to)
}
