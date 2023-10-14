module github.com/ArkhamCookie/crabot

go 1.20

require (
	crabot/crabtalk v0.0.0
)

replace (
	crabot/crabtalk => ./lib/crabtalk
)
