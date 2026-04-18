package commission

import (
	"fmt"
	"math"
)

func Calculate(amount float64, isAlif bool) (float64, float64, error) {
	if amount < 500 {
		return 0, 0, fmt.Errorf("сумма слишком мала")
	}
	if amount > 15000000 {
		return 0, 0, fmt.Errorf("превышен лимит ")
	}

	var fee float64
	if isAlif {
		fee = 0
	} else {
		fee = math.Round(amount * 0.0029)
	}

	total := amount + fee
	return fee, total, nil
}
