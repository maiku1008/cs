package cash

import "math"

func cash(input float64) int {
	var coins int
	change := int(math.Round(100 * input))

	for change > 0 {
		switch {
		case change >= 25:
			change -= 25
		case change >= 10:
			change -= 10
		case change >= 5:
			change -= 5
		case change >= 1:
			change -= 1
		}
		coins++
	}
	return coins
}
