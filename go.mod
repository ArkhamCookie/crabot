module github.com/ArkhamCookie/crabot

go 1.20

require (
	crabot/crabtalk v0.0.0
	formatting/markdown v0.0.0
)

require github.com/bwmarrin/discordgo v0.27.1

require (
	github.com/gSpera/morse v1.1.2 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	golang.org/x/crypto v0.1.0 // indirect
	golang.org/x/sys v0.1.0 // indirect
)

replace (
	crabot/crabtalk => ./lib/crabtalk
	crabot/formatting => ./lib/formatting
	formatting/markdown => ./lib/formatting/markdown
)
