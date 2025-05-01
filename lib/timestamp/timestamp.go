package timestamp

import (
	"fmt"
	"strconv"

	"crabot/unixtime"
)

func ShortDate(year, month, day, hour, minute, second int, timezone string) (string, error) {
	givenTime, err := unixtime.UnixTime(year, month, day, hour, minute, second, timezone)
	if err != nil {
		return "", err
	}

	timestamp := fmt.Sprintf("<t:%v:d>", givenTime)

	return timestamp, nil
}

func LongDate(year, month, day, hour, minute, second int, timezone string) (string, error) {
	givenTime, err := unixtime.UnixTime(year, month, day, hour, minute, second, timezone)
	if err != nil {
		return "", err
	}

	timestamp := fmt.Sprintf("<t:%v:D>", givenTime)

	return timestamp, nil
}

func ShortTime(year, month, day, hour, minute, second int, timezone string) (string, error) {
	givenTime, err := unixtime.UnixTime(year, month, day, hour, minute, second, timezone)
	if err != nil {
		return "", err
	}

	timestamp := fmt.Sprintf("<t:%v:t>", givenTime)

	return timestamp, nil
}

func LongTime(year, month, day, hour, minute, second int, timezone string) (string, error) {
	givenTime, err := unixtime.UnixTime(year, month, day, hour, minute, second, timezone)
	if err != nil {
		return "", err
	}

	timestamp := fmt.Sprintf("<t:%v:T>", givenTime)

	return timestamp, nil
}

func ShortFull(year, month, day, hour, minute, second int, timezone string) (string, error) {
	givenTime, err := unixtime.UnixTime(year, month, day, hour, minute, second, timezone)
	if err != nil {
		return "", err
	}

	timestamp := fmt.Sprintf("<t:%v:f>", givenTime)

	return timestamp, nil
}

func LongFull(year, month, day, hour, minute, second int, timezone string) (string, error) {
	givenTime, err := unixtime.UnixTime(year, month, day, hour, minute, second, timezone)
	if err != nil {
		return "", err
	}

	timestamp := fmt.Sprintf("<t:%v:F>", givenTime)

	return timestamp, nil
}

func RelativeTime(year, month, day, hour, minute, second int, timezone string) (string, error) {
	givenTime, err := unixtime.UnixTime(year, month, day, hour, minute, second, timezone)
	if err != nil {
		return "", err
	}

	timestamp := fmt.Sprintf("<t:%v:R>", givenTime)

	return timestamp, nil
}

func Unix(year, month, day, hour, minute, second int, timezone string) (string, error) {
	givenTime, err := unixtime.UnixTime(year, month, day, hour, minute, second, timezone)
	if err != nil {
		return "", err
	}

	unixTimestamp := strconv.FormatInt(givenTime, 10)

	return unixTimestamp, nil
}
