package dice

import "math/rand"

// Roll *dice* and return the roll.
func Roll(dice int) (int) {
	// Add 1 so it 1-int rather than 0-int
	return rand.Intn(dice) + 1
}
