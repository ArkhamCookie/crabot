package roll_test

import (
	"testing"

	"github.com/ArkhamCookie/crabot/dice/roll"
)

func TestRoll(t *testing.T) {
	output := roll.Roll(1)

	if output != 1 {
		t.Fatalf("`roll.Roll(1)` = %q, wanted 1", output)
	}
}
