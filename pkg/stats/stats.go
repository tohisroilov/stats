package stats

import (
	"github.com/tohisroilov/bank/v2/pkg/types"
)

func Avg(payments []types.Payment) types.Money {
	var summa types.Money = 0
	var len types.Money = 0
	for _, value := range payments {
		if value.Status != types.StatusFail{
		len += 1
		summa += value.Amount
	}
}
	return summa / len
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	var summa types.Money = 0
	for _, value := range payments {
		if category == value.Category && value.Status != types.StatusFail{
			summa += value.Amount
		}
	}
	return summa
}