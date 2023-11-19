package dice_test

import (
	"log"
	"testing"

	"crabot/dice"
)

func TestCoinFlip(t *testing.T) {
	// Flip the coin
	result, err := dice.CoinFlip()
	
	if err != nil {
		log.Fatalln("Error:", err)
	}

	// Confirm it is heads or tails,
	// if it isn't throw an error.
	switch result {
	case "heads":
		return
	case "tails":
		return
	default:
		t.Fatalf(`dice.CoinFlip() = wanted "heads" or "tails" got %s`, result)
	}
}
