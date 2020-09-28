package stats

import (
	"fmt"
	"github.com/tohisroilov/bank/v2/pkg/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			1,
			10_000,
			" ",
			types.StatusOk,
		},
		{
			2,
			15_000,
			" ",
			types.StatusFail,
		},
		{
			2,
			10_000,
			" ",
			types.StatusOk,
		},
	}

	avg := Avg(payments)
	fmt.Println(avg)
	// Output: 10000
}

func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			1,
			10_000,
			"bank",
			types.StatusFail,
		},
		{
			2,
			15_000,
			"car",
			types.StatusOk,
		},
		{
			2,
			5_000,
			"bank",
			types.StatusOk,
		},
	}

	sum := TotalInCategory(payments, "bank")
	fmt.Println(sum)
	// Output: 20000
}