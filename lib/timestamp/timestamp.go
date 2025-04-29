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

func ShortDate() string {
	now := currentTime()

	return fmt.Sprintf("<t:%v:d>", now)
}

func LongDate() string {
	now := currentTime()

	return fmt.Sprintf("<t:%v:D>", now)
}

func ShortTime() string {
	now := currentTime()

	return fmt.Sprintf("<t:%v:t>", now)
}

func LongTime() string {
	now := currentTime()

	return fmt.Sprintf("<t:%v:T>", now)
}

func ShortFull() string {
	now := currentTime()

	return fmt.Sprintf("<t:%v:f>", now)
}

func LongFull() string {
	now := currentTime()

	return fmt.Sprintf("<t:%v:F>", now)
}

func RelativeTime() string {
	currentTime := time.Now().Unix()
	convertedTime := strconv.FormatInt(currentTime, 10)

	return fmt.Sprintf("<t:%v:R>", convertedTime)
}

func UnixTimestamp() string {
	return currentTime()
}
