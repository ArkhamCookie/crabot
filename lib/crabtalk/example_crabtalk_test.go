package crabtalk_test

import (
	"fmt"

	"github.com/ArkhamCookie/crabot/lib/crabtalk"
)

func ExampleGet() {
	crab, _ := crabtalk.Get("Hello World")

	fmt.Println(crab)
	// Output: clickclickclickclick click clickclackclickclick clickclackclickclick clackclackclack / clickclackclack clackclackclack clickclackclick clickclackclickclick clackclickclick
}
