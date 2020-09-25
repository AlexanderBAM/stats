package stats

import (
	"fmt"

	"github.com/AlexanderBAM/bank/pkg/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			ID:       2,
			Amount:   10_00,
			Category: "Mobile",
		},
		{
			ID:       1,
			Amount:   60_00,
			Category: "Restaurant",
		},
		{
			ID:       3,
			Amount:   20_00,
			Category: "Taxi",
		},
	}

	fmt.Println(Avg(payments))

	//Output: 3000
}

func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			ID:       2,
			Amount:   10_00,
			Category: "Mobile",
		},
		{
			ID:       1,
			Amount:   50_00,
			Category: "Auto",
		},
		{
			ID:       3,
			Amount:   30_00,
			Category: "Mobile",
		},
	}

	fmt.Println(TotalInCategory(payments, "Mobile"))

	//Output: 4000
}
