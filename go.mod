module github.com/ArkhamCookie/crabot

go 1.24.2

require (
	crabot/crabtalk v0.0.0
	crabot/dice v0.0.0
	internal/dicecmd v0.0.0
	internal/env v0.0.0
)

require (
	crabot/timestamp v0.0.0
	github.com/bwmarrin/discordgo v0.28.1
)

require (
	crabot/dice/roll v0.0.0 // indirect
	crabot/lastupload v0.0.0 // indirect
	github.com/gSpera/morse v1.1.2 // indirect
	github.com/gorilla/websocket v1.5.3 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	golang.org/x/crypto v0.37.0 // indirect
	golang.org/x/sys v0.32.0 // indirect
)

replace (
	crabot/crabtalk => ./lib/crabtalk
	crabot/dice => ./lib/dice
	crabot/dice/roll => ./lib/dice/roll
	crabot/lastupload => ./lib/lastupload
	crabot/timestamp => ./lib/timestamp
	crabot/timestamp/unixtime => ./lib/timestamp/unixtime
	internal/dicecmd => ./internal/dicecmd
	internal/env => ./internal/env
)
