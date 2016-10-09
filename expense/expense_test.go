package expense

import (
	"testing"
)

//**************** Test Expense Command Creation ****************

func TestNew(t *testing.T) {
	amounts := []float64{0.0, 1.0, 2.5, 1000000.0}

	for i, amount := range amounts {
		expense := New(amount)

		if expense == nil {
			t.Errorf("Expense initialized to nil in case: %d", i)
		}
		if expense.amount != amount {
			t.Errorf("Amount set incorrectly. Expected %f, got: %f", amount, expense.amount)
		}
	}
}
