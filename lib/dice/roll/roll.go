package roll

import "math/rand"

// Roll returns a "roll" of a given `dice`.
//
// Always returns greater than or equal to 1.
func Roll(dice int) int {
	// Add 1 so it 1-int rather than 0-int
	return rand.Intn(dice) + 1
}
