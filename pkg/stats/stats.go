package stats

import (
	"github.com/AlexanderBAM/bank/v2/pkg/types"
)

// Avg calculates average amount
func Avg(payments []types.Payment) types.Money {

	sum := types.Money(0)

	var payment types.Payment

	numberOfPayments := 0

	for _, payment = range payments {
		if payment.Status != types.StatusFail {
			sum += payment.Amount
			numberOfPayments++
		}
	}

	// numberOfPayments := len(payments)

	average := int(sum) / numberOfPayments

	return types.Money(average)
}

// TotalInCategory calculates the sums of purchases in one category
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	sum := types.Money(0)

	var payment types.Payment

	for _, payment = range payments {
		if category == payment.Category && payment.Status != types.StatusFail {
			sum += payment.Amount
		}
	}

	return sum
}
