module github.com/ArkhamCookie/crabot

go 1.24.2

require (
	github.com/ArkhamCookie/crabot/lib/crabtalk v0.0.0
	github.com/ArkhamCookie/crabot/lib/lastupload v0.0.0
	github.com/ArkhamCookie/crabot/lib/timestamp v0.0.0
)

require (
	internal/dicecmd v0.0.0
	internal/env v0.0.0
)

require (
	github.com/bwmarrin/discordgo v0.28.1
)

require (
	github.com/ArkhamCookie/crabot/lib/dice v0.0.0 // indirect
	github.com/ArkhamCookie/crabot/lib/dice/roll v0.0.0 // indirect
	github.com/ArkhamCookie/crabot/lib/timestamp/unixtime v0.0.0 // indirect
)

require (
	github.com/gSpera/morse v1.1.2 // indirect
	github.com/gorilla/websocket v1.5.3 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	golang.org/x/crypto v0.37.0 // indirect
	golang.org/x/sys v0.32.0 // indirect
)

replace (
	github.com/ArkhamCookie/crabot/lib/crabtalk => ./lib/crabtalk
	github.com/ArkhamCookie/crabot/lib/dice => ./lib/dice
	github.com/ArkhamCookie/crabot/lib/dice/roll => ./lib/dice/roll
	github.com/ArkhamCookie/crabot/lib/lastupload => ./lib/lastupload
	github.com/ArkhamCookie/crabot/lib/timestamp => ./lib/timestamp
	github.com/ArkhamCookie/crabot/lib/timestamp/unixtime => ./lib/timestamp/unixtime
)

replace (
	internal/dicecmd => ./internal/dicecmd
	internal/env => ./internal/env
)
