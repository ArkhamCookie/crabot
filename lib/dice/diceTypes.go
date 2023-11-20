package dice

import (
	"internal/dice"
)

var (
	result int
)

// D2 rolls a D2 (1-2)
func D2(count int) int {
	for count > 0 {
		result += dice.Roll(2)
		count--
	}
	return result
}

// D4 rolls a D4 (1-4)
func D4(count int) int {
	for count > 0 {
		result += dice.Roll(4)
		count--
	}
	return result
}

// D6 rolls a D6 (1-6)
func D6(count int) int {
	for count > 0 {
		result += dice.Roll(6)
		count--
	}
	return result
}

// D8 rolls a D8 (1-8)
func D8(count int) int {
	for count > 0 {
		result += dice.Roll(8)
		count--
	}
	return result
}

// D10 rolls a D10 (1-10)
func D10(count int) int {
	for count > 0 {
		result += dice.Roll(10)
		count--
	}
	return result
}

// D12 rolls a D12 (1-12)
func D12(count int) int {
	for count > 0 {
		result += dice.Roll(12)
		count--
	}
	return result
}

// D20 rolls a D20 (1-20)
func D20(count int) int {
	for count > 0 {
		result += dice.Roll(20)
		count--
	}
	return result
}

// Can only roll one at a time for the moment
func D00(count int) int {
	for count > 0 {
		a := dice.Roll(10)

		if a == 10 {
			a = 0
		} else {
			a = a * 10
		}

		b := dice.Roll(10)

		if b == 1 && a == 0 {
			return 100
		}

		result = a + b
	}
	return result
}
