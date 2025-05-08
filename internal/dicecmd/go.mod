module internal/dicecmd

go 1.24.2

require (
	github.com/ArkhamCookie/crabot/lib/dice v0.0.0
)

require (
	"crabot/dice/roll" v0.0.0 // indirect
)

replace (
	github.com/ArkhamCookie/crabot/lib/dice => ../../lib/dice
	crabot/dice/roll => ../../lib/dice/roll
)
