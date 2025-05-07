package timestamp

import (
	"fmt"
)

// Get timestamp from unix timestamp in short date format.
func UnixShortDate(unixTime int64) string {
	return fmt.Sprintf("<t:%v:d>", unixTime)
}

// Get timestamp from unix timestamp in long date format.
func UnixLongDate(unixTime int64) string {
	return fmt.Sprintf("<t:%v:D>", unixTime)
}

// Get timestamp from unix timestamp in short time format.
func UnixShortTime(unixTime int64) string {
	return fmt.Sprintf("<t:%v:t>", unixTime)
}

// Get timestamp from unix timestamp in long time format.
func UnixLongTime(unixTime int64) string {
	return fmt.Sprintf("<t:%v:T>", unixTime)
}

// Get timestamp from unix timestamp in short full format.
func UnixShortFull(unixTime int64) string {
	return fmt.Sprintf("<t:%v:f>", unixTime)
}

// Get timestamp from unix timestamp in long full format.
func UnixLongFull(unixTime int64) string {
	return fmt.Sprintf("<t:%v:F>", unixTime)
}

// Get timestamp from unix timestamp in relative time format.
func UnixRelativeTime(unixTime int64) string {
	return fmt.Sprintf("<t:%v:R>", unixTime)
}
