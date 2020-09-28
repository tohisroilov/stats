package stats

import (
	"github.com/tohisroilov/bank/pkg/types"
)

func Avg(payments []types.Payment) types.Money {
	var summa types.Money = 0
	lenght := len(payments)
	for _, value := range payments {
		summa += value.Amount
	}
	return summa / types.Money(lenght)
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	var summa types.Money = 0
	for _, value := range payments {
		if category == value.Category {
			summa += value.Amount
		}
	}
	return summa
}