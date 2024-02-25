module github.com/ArkhamCookie/crabot

go 1.21

require (
	crabot/crabtalk v0.0.0
	crabot/formatting/markdown v0.0.0
	crabot/formatting/timestamps v0.0.0
	internal/env v0.0.0
)

require github.com/bwmarrin/discordgo v0.27.1

require golang.org/x/crypto v0.17.0 // indirect

replace (
	crabot/crabtalk => ./lib/crabtalk
	crabot/dice => ./lib/dice
	crabot/dice/roll => ./lib/dice/roll
	crabot/formatting => ./lib/formatting
	crabot/formatting/markdown => ./lib/formatting/markdown
	crabot/formatting/timestamps => ./lib/formatting/timestamps
	internal/env => ./internal/env
)
