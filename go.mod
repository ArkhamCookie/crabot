module github.com/ArkhamCookie/crabot

go 1.20

require (
	crabot/crabtalk v0.0.0
	github.com/gSpera/morse v1.1.2 // indirect
)

replace (
	crabot/crabtalk => ./lib/crabtalk
)
