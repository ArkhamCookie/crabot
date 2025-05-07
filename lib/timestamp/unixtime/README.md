# crabot/timestamp/unixtime

Convert given time variables to unix time.

## Examples

### Basic Usage

```golang
package main

import (
  "fmt"

  "github.com/ArkhamCookie/crabot/lib/timestamp/unixtime"
)

func main() {
  // Get unix time from given variables
  unixTime, _ := unixtime.UnixTime(2025, 12, 12, 12, 12, 00, "UTC")

  fmt.Println(unixTime) // "1765541532"
}
```
