module github.com/ArkhamCookie/crabot

go 1.20

require (
	crabot/crabtalk v0.0.0
	crabot/formatting/markdown v0.0.0
	internal/env v0.0.0
)

require github.com/bwmarrin/discordgo v0.27.1

require (
	github.com/gSpera/morse v1.1.2 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	golang.org/x/crypto v0.17.0 // indirect
	golang.org/x/sys v0.15.0 // indirect
)

replace (
	crabot/crabtalk => ./lib/crabtalk
	crabot/dice => ./lib/dice
	crabot/dice/roll => ./lib/dice/roll
	crabot/formatting => ./lib/formatting
	crabot/formatting/markdown => ./lib/formatting/markdown
	internal/env => ./internal/env
)
