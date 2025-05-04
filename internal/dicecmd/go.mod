module internal/dicecmd

go 1.24.2

require (
	"crabot/dice" v0.0.0
)

require (
	"crabot/dice/roll" v0.0.0 // indirect
)

replace (
	crabot/dice => ../../lib/dice
	crabot/dice/roll => ../../lib/dice/roll
)
