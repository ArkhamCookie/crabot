package timestamp

import (
	"fmt"
	"strconv"
	"time"
)

func currentTime() string {
	timeStart := time.Now().Unix()
	convertedTime := strconv.FormatInt(timeStart, 10)

	return convertedTime
}

func CurrentShortDate() string {
	now := currentTime()

	return fmt.Sprintf("<t:%v:d>", now)
}

func CurrentLongDate() string {
	now := currentTime()

	return fmt.Sprintf("<t:%v:D>", now)
}

func CurrentShortTime() string {
	now := currentTime()

	return fmt.Sprintf("<t:%v:t>", now)
}

func CurrentLongTime() string {
	now := currentTime()

	return fmt.Sprintf("<t:%v:T>", now)
}

func CurrentShortFull() string {
	now := currentTime()

	return fmt.Sprintf("<t:%v:f>", now)
}

func CurrentLongFull() string {
	now := currentTime()

	return fmt.Sprintf("<t:%v:F>", now)
}

func CurrentRelativeTime() string {
	currentTime := time.Now().Unix()
	convertedTime := strconv.FormatInt(currentTime, 10)

	return fmt.Sprintf("<t:%v:R>", convertedTime)
}

func CurrentUnixTimestamp() string {
	return currentTime()
}
