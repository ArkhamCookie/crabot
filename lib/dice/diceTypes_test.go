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
	}

	if result == 0 {
		t.Fatal("dice.D2(1) was 0")
	}

	if result > 2 {
		t.Fatal("dice.D2(1) was greater than 2")
	}

	result = dice.D2(2)

	if result < 0 {
		t.Fatal("dice.D2(2) was negative")
	}

	if result == 0 {
		t.Fatal("dice.D2(2) was 0")
	}

	if result < 2 {
		t.Fatal("dice.D2(2) was less than 2")
	}
}

func TestD4(t *testing.T) {
	result = dice.D4(1)

	if result < 0 {
		t.Fatal("dice.D4(1) was negative")
	}

	if result == 0 {
		t.Fatal("dice.D4(1) was 0")
	}

	if result > 2 {
		t.Fatal("dice.D4(1) was greater than 2")
	}

	result = dice.D4(2)

	if result < 0 {
		t.Fatal("dice.D4(2) was negative")
	}

	if result == 0 {
		t.Fatal("dice.D4(2) was 0")
	}

	if result < 2 {
		t.Fatal("dice.D4(2) was less than 2")
	}
}

func TestD6(t *testing.T) {
	result = dice.D6(1)

	if result < 0 {
		t.Fatal("dice.D6(1) was negative")
	}

	if result == 0 {
		t.Fatal("dice.D6(1) was 0")
	}

	if result > 2 {
		t.Fatal("dice.D6(1) was greater than 2")
	}

	result = dice.D6(2)

	if result < 0 {
		t.Fatal("dice.D6(2) was negative")
	}

	if result == 0 {
		t.Fatal("dice.D6(2) was 0")
	}

	if result < 2 {
		t.Fatal("dice.D6(2) was less than 2")
	}
}

func TestD8(t *testing.T) {
	result = dice.D8(1)

	if result < 0 {
		t.Fatal("dice.D8(1) was negative")
	}

	if result == 0 {
		t.Fatal("dice.D8(1) was 0")
	}

	if result > 2 {
		t.Fatal("dice.D8(1) was greater than 2")
	}

	result = dice.D8(2)

	if result < 0 {
		t.Fatal("dice.D8(2) was negative")
	}

	if result == 0 {
		t.Fatal("dice.D8(2) was 0")
	}

	if result < 2 {
		t.Fatal("dice.D8(2) was less than 2")
	}
}

func TestD10(t *testing.T) {
	result = dice.D10(1)

	if result < 0 {
		t.Fatal("dice.D10(1) was negative")
	}

	if result == 0 {
		t.Fatal("dice.D10(1) was 0")
	}

	if result > 2 {
		t.Fatal("dice.D10(1) was greater than 2")
	}

	result = dice.D10(2)

	if result < 0 {
		t.Fatal("dice.D10(2) was negative")
	}

	if result == 0 {
		t.Fatal("dice.D10(2) was 0")
	}

	if result < 2 {
		t.Fatal("dice.D10(2) was less than 2")
	}
}

func TestD12(t *testing.T) {
	result = dice.D12(1)

	if result < 0 {
		t.Fatal("dice.D12(1) was negative")
	}

	if result == 0 {
		t.Fatal("dice.D12(1) was 0")
	}

	if result > 2 {
		t.Fatal("dice.D12(1) was greater than 2")
	}

	result = dice.D12(2)

	if result < 0 {
		t.Fatal("dice.D12(2) was negative")
	}

	if result == 0 {
		t.Fatal("dice.D12(2) was 0")
	}

	if result < 2 {
		t.Fatal("dice.D12(2) was less than 2")
	}
}

func TestD20(t *testing.T) {
	result = dice.D20(1)

	if result < 0 {
		t.Fatal("dice.D20(1) was negative")
	}

	if result == 0 {
		t.Fatal("dice.D20(1) was 0")
	}

	if result > 2 {
		t.Fatal("dice.D20(1) was greater than 2")
	}

	result = dice.D20(2)

	if result < 0 {
		t.Fatal("dice.D20(2) was negative")
	}

	if result == 0 {
		t.Fatal("dice.D20(2) was 0")
	}

	if result < 2 {
		t.Fatal("dice.D20(2) was less than 2")
	}
}

func TestD00(t *testing.T) {
	percentileResults := dice.D00(1)

	for i := range percentileResults {
		if percentileResults[i] < 0 {
			t.Fatal("dice.D00(1) was negative")
		} else if percentileResults[i] > 100 {
			t.Fatal("dice.D00(1) was over 100")
		}
	}

	percentileResults = dice.D00(2)
	for i := range percentileResults {
		if percentileResults[i] < 0 {
			t.Fatal("dice.D00(2) was negative")
		} else if percentileResults[i] > 100 {
			t.Fatal("dice.D00(2) was over 100")
		}
	}
}
