package stats

import (
	"fmt"
	"github.com/tohisroilov/bank/pkg/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			1,
			10_000,
			" ",
		},
		{
			2,
			15_000,
			" ",
		},
		{
			2,
			5_000,
			" ",
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
		},
		{
			2,
			15_000,
			"car",
		},
		{
			2,
			5_000,
			"bank",
		},
	}

	sum := TotalInCategory(payments, "bank")
	fmt.Println(sum)
	// Output: 15000
}