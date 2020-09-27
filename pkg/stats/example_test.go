package stats

import (
	"fmt"
	"reflect"
	"testing"

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

func TestCategoriesAvg(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "fun", Amount: 100_00},
		{ID: 2, Category: "auto", Amount: 200_00},
		{ID: 3, Category: "food", Amount: 300_00},
		{ID: 4, Category: "auto", Amount: 400_00},
		{ID: 5, Category: "fun", Amount: 500_00},
		{ID: 6, Category: "auto", Amount: 300_00},
		{ID: 7, Category: "fun", Amount: 600_00},
		{ID: 8, Category: "food", Amount: 800_00},
	}
	expected := map[types.Category]types.Money{
		"auto": 300_00,
		"food": 550_00,
		"fun":  400_00,
	}

	result := CategoriesAvg(payments)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Invalid result, expected: %v, actual: %v", expected, result)
	}
}
