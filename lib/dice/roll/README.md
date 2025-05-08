# crabot/dice/roll

A library that rolls any type of dice using the [math/rand](https://pkg.go.dev/math/rand) package.

## Examples

### Roll a D13

```golang
package main

import (
  "fmt"

  "github.com/ArkhamCookie/crabot/lib/dice/roll"
)

func main() {
  // Roll the D13
  result := roll.Roll(13)

  // Print the output
  fmt.Printf("Rolled a %d\m", result)
}
```
