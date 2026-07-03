module github.com/ArkhamCookie/crabot

go 1.24.2

require (
	crabot/crabtalk v0.0.0
	crabot/dice v0.0.0
	crabot/lastupload v0.0.0
	internal/database v0.0.0
	internal/dicecmd v0.0.0
	internal/env v0.0.0
)

require (
	crabot/timestamp v0.0.0
	github.com/bwmarrin/discordgo v0.28.1
)

require (
	crabot/dice/roll v0.0.0 // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/gSpera/morse v1.1.2 // indirect
	github.com/glebarez/go-sqlite v1.22.0 // indirect
	github.com/google/uuid v1.5.0 // indirect
	github.com/gorilla/websocket v1.5.3 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20230129092748-24d4a6f8daec // indirect
	golang.org/x/crypto v0.37.0 // indirect
	golang.org/x/sys v0.32.0 // indirect
	modernc.org/libc v1.37.6 // indirect
	modernc.org/mathutil v1.6.0 // indirect
	modernc.org/memory v1.7.2 // indirect
	modernc.org/sqlite v1.28.0 // indirect
)

replace (
	crabot/crabtalk => ./lib/crabtalk
	crabot/dice => ./lib/dice
	crabot/dice/roll => ./lib/dice/roll
	crabot/lastupload => ./lib/lastupload
	crabot/timestamp => ./lib/timestamp
	crabot/timestamp/unixtime => ./lib/timestamp/unixtime

	internal/database => ./internal/database
	internal/dicecmd => ./internal/dicecmd
	internal/env => ./internal/env
)
