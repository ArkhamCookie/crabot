package timestamp

import (
	"fmt"
	"strconv"
	"time"
)

// Get current time's unix time as a string.
func currentTime() string {
	timeStart := time.Now().Unix()
	convertedTime := strconv.FormatInt(timeStart, 10)

	return convertedTime
}

// Get timestamp for current time in short date format.
func CurrentShortDate() string {
	now := currentTime()

	return fmt.Sprintf("<t:%v:d>", now)
}

// Get timestamp for current time in long date format.
func CurrentLongDate() string {
	now := currentTime()

	return fmt.Sprintf("<t:%v:D>", now)
}

// Get timestamp for current time in short time format.
func CurrentShortTime() string {
	now := currentTime()

	return fmt.Sprintf("<t:%v:t>", now)
}

// Get timestamp for current time in long time format.
func CurrentLongTime() string {
	now := currentTime()

	return fmt.Sprintf("<t:%v:T>", now)
}

// Get timestamp for current time in short full format.
func CurrentShortFull() string {
	now := currentTime()

	return fmt.Sprintf("<t:%v:f>", now)
}

// Get timestamp for current time in long full format.
func CurrentLongFull() string {
	now := currentTime()

	return fmt.Sprintf("<t:%v:F>", now)
}

// Get timestamp for current time in relative time format.
func CurrentRelativeTime() string {
	currentTime := time.Now().Unix()
	convertedTime := strconv.FormatInt(currentTime, 10)

	return fmt.Sprintf("<t:%v:R>", convertedTime)
}

// Get unix timestamp of current time.
func CurrentUnixTimestamp() string {
	return currentTime()
}
