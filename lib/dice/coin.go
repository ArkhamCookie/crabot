package dice

import "internal/dice"

// Flip a coin and get heads or tails.
// Optionally, you can 'call it in the air' by selecting heads or tails.
// If you 'call' it correctly, Coin function returns true.
func Coin(call string) (string, bool) {
	var side string
	// Flip the coin (roll d2).
	value := dice.Roll(2)

	// Set side depending on roll.
	switch value {
	case 1:
		side = "heads"
	case 2:
		side = "tails"
	}

	// Check if user guessed the side correctly.
	// If they did, then return true.
	if call == side {
		return side, true
	}
	// If called wrong (or not called), then return false.
	return side, false
}
