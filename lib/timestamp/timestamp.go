package timestamp

import (
	"fmt"
	"strconv"

	"crabot/unixtime"
)

// Get timestamp from given time in short date format.
func ShortDate(year, month, day, hour, minute, second int, timezone string) (string, error) {
	givenTime, err := unixtime.UnixTime(year, month, day, hour, minute, second, timezone)
	if err != nil {
		return "", err
	}

	timestamp := fmt.Sprintf("<t:%v:d>", givenTime)

	return timestamp, nil
}

// Get timestamp from given time in long date format.
func LongDate(year, month, day, hour, minute, second int, timezone string) (string, error) {
	givenTime, err := unixtime.UnixTime(year, month, day, hour, minute, second, timezone)
	if err != nil {
		return "", err
	}

	timestamp := fmt.Sprintf("<t:%v:D>", givenTime)

	return timestamp, nil
}

// Get timestamp from given time in short time format.
func ShortTime(year, month, day, hour, minute, second int, timezone string) (string, error) {
	givenTime, err := unixtime.UnixTime(year, month, day, hour, minute, second, timezone)
	if err != nil {
		return "", err
	}

	timestamp := fmt.Sprintf("<t:%v:t>", givenTime)

	return timestamp, nil
}

// Get timestamp from given time in long time format.
func LongTime(year, month, day, hour, minute, second int, timezone string) (string, error) {
	givenTime, err := unixtime.UnixTime(year, month, day, hour, minute, second, timezone)
	if err != nil {
		return "", err
	}

	timestamp := fmt.Sprintf("<t:%v:T>", givenTime)

	return timestamp, nil
}

// Get timestamp from given time in short full format.
func ShortFull(year, month, day, hour, minute, second int, timezone string) (string, error) {
	givenTime, err := unixtime.UnixTime(year, month, day, hour, minute, second, timezone)
	if err != nil {
		return "", err
	}

	timestamp := fmt.Sprintf("<t:%v:f>", givenTime)

	return timestamp, nil
}

// Get timestamp from given time in long full format.
func LongFull(year, month, day, hour, minute, second int, timezone string) (string, error) {
	givenTime, err := unixtime.UnixTime(year, month, day, hour, minute, second, timezone)
	if err != nil {
		return "", err
	}

	timestamp := fmt.Sprintf("<t:%v:F>", givenTime)

	return timestamp, nil
}

// Get timestamp from given time in relative time format.
func RelativeTime(year, month, day, hour, minute, second int, timezone string) (string, error) {
	givenTime, err := unixtime.UnixTime(year, month, day, hour, minute, second, timezone)
	if err != nil {
		return "", err
	}

	timestamp := fmt.Sprintf("<t:%v:R>", givenTime)

	return timestamp, nil
}

// Get unix time from given time as string
func Unix(year, month, day, hour, minute, second int, timezone string) (string, error) {
	givenTime, err := unixtime.UnixTime(year, month, day, hour, minute, second, timezone)
	if err != nil {
		return "", err
	}

	unixTimestamp := strconv.FormatInt(givenTime, 10)

	return unixTimestamp, nil
}
