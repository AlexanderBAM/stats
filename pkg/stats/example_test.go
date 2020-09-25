package stats

import (
	"fmt"

	"github.com/AlexanderBAM/bank/v2/pkg/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			ID:       2,
			Amount:   10_00,
			Category: "Mobile",
			Status:   types.StatusOk,
		},
		{
			ID:       1,
			Amount:   60_00,
			Category: "Restaurant",
			Status:   types.StatusFail,
		},
		{
			ID:       3,
			Amount:   20_00,
			Category: "Taxi",
			Status:   types.StatusProgress,
		},
	}

	fmt.Println(Avg(payments))

	//Output: 1500
}

func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			ID:       2,
			Amount:   10_00,
			Category: "Mobile",
			Status:   types.StatusOk,
		},
		{
			ID:       1,
			Amount:   50_00,
			Category: "Auto",
			Status:   types.StatusProgress,
		},
		{
			ID:       3,
			Amount:   30_00,
			Category: "Mobile",
			Status:   types.StatusFail,
		},
	}

	fmt.Println(TotalInCategory(payments, "Mobile"))

	//Output: 1000
}
