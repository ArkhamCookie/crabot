package dice_test

import (
	"testing"

	"crabot/dice"
)

var (
	result int
)

func TestD2(t *testing.T) {
	// Test rolling 1 D2
	result = dice.D2(1)

	// TODO: Test that it is a full number

	if result < 0 {
		t.Fatal("dice.D2 was negitive")
	} else if result == 0 {
		t.Fatal("dice.D2 was 0")
	}

	if result > 2 {
		t.Fatal("dice.D2 was greater than 2 when rolling 1 dice")
	}

	// Test rolling 2 D2
	result = dice.D2(2)

	if result < 0 {
		t.Fatal("dice.D2 was negitive")
	} else if result == 0 {
		t.Fatal("dice.D2 was 0")
	}

	if result < 2 {
		t.Fatal("dice.D2 was less than 2 when rolling 2 dice")
	}
}

func TestD4(t *testing.T) {
	// Test rolling 1 D4
	result = dice.D4(1)

	if result < 0 {
		t.Fatal("dice.D4 was negitive")
	} else if result == 0 {
		t.Fatal("dice.D4 was 0")
	}

	if result < 4 {
		t.Fatal("dice.D4 was greater than 4 when rolling 1 dice")
	}
	
	
}
