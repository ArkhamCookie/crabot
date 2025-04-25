package dice

import (
	"errors"

	"crabot/dice/roll"
)

// CoinFlip flips a coin and get heads or tails.
func CoinFlip() (string, error) {
	// Flip the coin (roll d2).
	value := roll.Roll(2)

	// Set side depending on roll.
	switch value {
	case 1:
		return "heads", nil
	case 2:
		return "tails", nil
	}
	return "", errors.New("issue flipping coin")
}

// CoinCall allows you to 'call' the result of CoinFlip 'in the air'
// by selecting heads or tails.
// If you 'call' it correctly, CoinCall returns true.
func CoinCall(call string) (string, bool, error) {
	// Flip a coin and get the result
	side, err := CoinFlip()
	if err != nil {
		return "", false, err
	}

	// If called correct, then return true.
	if side == call {
		return side, true, nil
	}

	// If called wrong, then return false.
	return side, false, nil
}
