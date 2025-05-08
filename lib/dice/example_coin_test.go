package dice_test

import (
	"fmt"

	"github.com/ArkhamCookie/crabot/lib/dice"
)

func ExmapleCoinFlip() {
	coin, _ := dice.CoinFlip()

	fmt.Printf("Coin was %v", coin)
}

func ExampleCoinCall() {
	coin, correct, _ := dice.CoinCall("heads")

	if correct {
		fmt.Println("You guessed correctly!")
	} else {
		fmt.Println("You guessed wrong.")
	}

	fmt.Printf("coin was %v", coin)
}
