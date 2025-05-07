# crabot/lastupload

Fetch the last upload of a YouTuber.

## Examples

### Get a Channels' ID

```golang
package main

import (
  "fmt"

  "github.com/ArkhamCookie/crabot/lib/lastupload"
)

func main() {
  YOUTUBE_KEY := "" // Insert your API key

  // Fetch channelID
  channelID, err := lastupload.GetChannelID("spiritvoices", YOUTUBE_KEY)
  if err != nil {
    fmt.Printf("[ERROR]: %v\n", err)
  }

  fmt.Println(channelID) // UCi_0J4Zm1qwiYVSGRsA0_Bg
}
```

### Get a Channel's Last Upload Time

```golang
package main

import (
  "fmt"

  "github.com/ArkhamCookie/crabot/lib/lastupload"
)

func main() {
  YOUTUBE_KEY := "" // Insert your API key

  // Get the last upload data
  uploadData, err := lastupload.GetData("spiritvoices", YOUTUBE_KEY)
  if err != nil {
    fmt.Printf("[ERROR]: %v\n", err)
  }

  // Fetch channelID
  date, err := lastupload.GetDate(uploadData, YOUTUBE_KEY)
  if err != nil {
    fmt.Printf("[ERROR]: %v\n", err)
  }

  fmt.Println(date)
}
```
