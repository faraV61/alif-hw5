package commission

import (
	"fmt"
	"math"
	"math/rand"
	"time"
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
func MaskCard(cardNumber string) string {
	last_four := cardNumber[len(cardNumber)-4:]
	return "UZCARD**" + last_four
}
func GetTransaction() int {
	return rand.Intn(900000000) + 100000000
}

func GetCurrentTime() string {
	return time.Now().Format("01.04.2026 15:09")
}
