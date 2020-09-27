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

//CategoriesAvg calculates average amount of payments in a category
func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {

	categories := map[types.Category]types.Money{}
	for _, payment := range payments {
		if payment.Status != types.StatusFail {
			categories[payment.Category] += payment.Amount
		}
	}

	counter := map[types.Category]types.Money{}
	for _, payment := range payments {
		if payment.Status != types.StatusFail {
			counter[payment.Category]++
		}
	}

	for key := range categories {
		categories[key] = categories[key] / counter[key]
	}

	return categories
}

//PeriodsDynamic compares two sets of data (maps)
func PeriodsDynamic(first map[types.Category]types.Money, second map[types.Category]types.Money) map[types.Category]types.Money {

	Dynamic := map[types.Category]types.Money{}
	RangeLength := map[types.Category]types.Money{}

	if len(first) >= len(second) {
		RangeLength = first
	} else {
		RangeLength = second
	}

	for key := range RangeLength {
		Dynamic[key] = second[key] - first[key]
	}

	return Dynamic
}
