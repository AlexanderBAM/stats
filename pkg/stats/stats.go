package stats

import (
	"github.com/AlexanderBAM/bank/pkg/types"
)

// Avg calculates average amount
func Avg(payments []types.Payment) types.Money {

	sum := types.Money(0)

	var payment types.Payment

	for _, payment = range payments {
		sum += payment.Amount
	}

	numberOfPayments := len(payments)

	average := int(sum) / numberOfPayments

	return types.Money(average)
}

// TotalInCategory calculates the sums of purchases in one category
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	sum := types.Money(0)

	var payment types.Payment

	for _, payment = range payments {
		if category == payment.Category {
			sum += payment.Amount
		}
	}

	return sum
}
