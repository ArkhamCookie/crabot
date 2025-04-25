package dice_test

import (
	"testing"

	"crabot/dice"
)

var (
	result int
)

func TestD2(t *testing.T) {
	result = dice.D2(1)

	if result < 0 {
		t.Fatal("dice.D2(1) was negative")
	} else if result == 0 {
		t.Fatal("dice.D2(1) was 0")
	}

	if result > 2 {
		t.Fatal("dice.D2(1) was greater than 2")
	}

	result = dice.D2(2)

	if result < 0 {
		t.Fatal("dice.D2(2) was negative")
	} else if result == 0 {
		t.Fatal("dice.D2(2) was 0")
	}

	if result < 2 {
		t.Fatal("dice.D2(2) was less than 2")
	}
}

func TestD00(t *testing.T) {
	percentileResults := dice.D00(1)

	for i := range percentileResults {
		if percentileResults[i] < 0 {
			t.Fatal("dice.D00(1) was negative")
		}
	}
}
