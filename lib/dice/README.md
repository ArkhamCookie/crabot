# crabot/dice

A library for rolling common types of dice.

## Examples

### Roll 1 D20

```golang
package main

import (
  "fmt"

  "github.com/ArkhamCookie/crabot/lib/dice"
)

func main() {
  // Roll a D20
  result := dice.D20(1)

  // Print result
  fmt.Println("Rolled a: %d", result)
}
```

### Roll 5 D20

```golang
package main

import (
  "fmt"

  "github.com/ArkhamCookie/crabot/lib/dice"
)

func main() {
  // Roll 5 D20
  result := dice.D20(5)

  // Print result
  fmt.Println("Rolled a: %d", result)
}
```
