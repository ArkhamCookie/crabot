module timestamp

go 1.24.2

require (
	github.com/ArkhamCookie/crabot/lib/timestamp v0.0.0
	github.com/ArkhamCookie/crabot/lib/timestamp/unixtime v0.0.0
)

replace (
	github.com/ArkhamCookie/crabot/lib/timestamp => .
	github.com/ArkhamCookie/crabot/lib/timestamp/unixtime => ./unixtime
)
