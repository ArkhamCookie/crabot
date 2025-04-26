package dicecmd

import (
	"crabot/dice"
	"fmt"
)

func DetermineDiceRoll(diceType string, diceAmount int) int {
	switch diceType {
	case "00":
		fallthrough
	case "100":
		fallthrough
	case "D00":
		fallthrough
	case "D100":
		if diceAmount > 1 {
			fmt.Println("WARNING: Only 1 dice is supported for D00/D100")
		}

		results := dice.D00(1)

		return results[0]
	case "2":
		fallthrough
	case "D2":
		return dice.D2(diceAmount)
	case "4":
		fallthrough
	case "D4":
		return dice.D4(diceAmount)
	case "6":
		fallthrough
	case "D6":
		return dice.D6(diceAmount)
	case "8":
		fallthrough
	case "D8":
		return dice.D8(diceAmount)
	case "10":
		fallthrough
	case "D10":
		return dice.D10(diceAmount)
	case "12":
		fallthrough
	case "D12":
		return dice.D12(diceAmount)
	case "20":
		fallthrough
	case "D20":
		return dice.D20(diceAmount)
	default:
		fmt.Println("WARNING: Unsupported dice type: %s", diceType)

		return dice.D6(diceAmount)
	}
}
