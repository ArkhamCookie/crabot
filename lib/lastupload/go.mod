module lastuplaod

go 1.24.2

require (
	github.com/ArkhamCookie/crabot/lib/lastupload v0.0.0
	internal/env v0.0.0
)

replace (
	github.com/ArkhamCookie/crabot/lib/lastupload => .
)
