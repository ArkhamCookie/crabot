package roll_test

import (
	"fmt"

	"github.com/ArkhamCookie/crabot/lib/dice/roll"
)

func ExampleRoll() {
	result := roll.Roll(13)

	fmt.Printf("Rolled a %d\n", result)
}
