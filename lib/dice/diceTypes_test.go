package dice_test

import (
	"testing"

	"crabot/dice"
)

var (
	result int
)

func TestD2(t *testing.T) {
	result = dice.D6(1)

	// TODO: Test that it is a full number

	if result < 0 {
		t.Fatal("dice.D6 was negitive")
	} else if result == 0 {
		t.Fatal("dice.D6 was 0")
	} else if result > 6 {
		t.Fatal("dice.D6 was greater than 6")
	}
}
