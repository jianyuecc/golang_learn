package model

//account
type Account struct {
	deposit float64
	notes   string
	money   float64
	Turnover
}

func (a *Account) SetTurnover(turnover Turnover) {
	a.Turnover = turnover
}

func (a *Account) GetTurnover() Turnover {
	return a.Turnover
}

func (a *Account) Money() float64 {
	return a.money
}

func (a *Account) SetMoney(money float64) {
	a.money = money
}

func (a *Account) Notes() string {
	return a.notes
}

func (a *Account) SetNotes(notes string) {
	a.notes = notes
}

func NewAccount() *Account {
	return &Account{deposit: 1000}
}

func (a *Account) Deposit() float64 {
	return a.deposit
}

func (a *Account) SetDeposit(deposit float64) {
	a.deposit = deposit
}

func init() {

}

type Turnover string

const (
	INCOME Turnover = "收入"
	OUTLAY Turnover = "支出"
)
