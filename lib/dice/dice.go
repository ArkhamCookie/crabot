package dice

import "internal/dice"

// D2 rolls a D2 (1-2)
func D2() int{
	return dice.Roll(2)
}
