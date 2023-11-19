package dice

import "internal/dice"

// D2 rolls a D2 (1-2)
func D2() int {
	return dice.Roll(2)
}

// D4 rolls a D4 (1-4)
func D4() int {
	return dice.Roll(4)
}

// D6 rolls a D6 (1-6)
func D6() int {
	return dice.Roll(6)
}

// D8 rolls a D8 (1-8)
func D8() int {
	return dice.Roll(8)
}

// D10 rolls a D10 (1-10)
func D10() int {
	return dice.Roll(10)
}

// D12 rolls a D12 (1-12)
func D12() int {
	return dice.Roll(12)
}

// D20 rolls a D20 (1-20)
func D20() int {
	return dice.Roll(20)
}
