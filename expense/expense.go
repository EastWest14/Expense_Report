package expense

import (
	"github.com/gAssert"
)

//**************** Expense Command Creation ****************

type Expense struct {
	amount float64
}

func New(amount float64) *Expense {
	gAssert.Assert(amount >= 0, "Expense amount cannot be set to a negative value")
	return &Expense{amount: amount}
}

func (e *Expense) Amount() float64 {
	return e.amount
}
