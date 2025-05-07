package timestamp

import (
	"fmt"
)

func UnixShortDate(unixTime int64) string {
	return fmt.Sprintf("<t:%v:d>", unixTime)
}

func UnixLongDate(unixTime int64) string {
	return fmt.Sprintf("<t:%v:D>", unixTime)
}

func UnixShortTime(unixTime int64) string {
	return fmt.Sprintf("<t:%v:t>", unixTime)
}

func UnixLongTime(unixTime int64) string {
	return fmt.Sprintf("<t:%v:T>", unixTime)
}

func UnixShortFull(unixTime int64) string {
	return fmt.Sprintf("<t:%v:f>", unixTime)
}

func UnixLongFull(unixTime int64) string {
	return fmt.Sprintf("<t:%v:F>", unixTime)
}

func UnixRelativeTime(unixTime int64) string {
	return fmt.Sprintf("<t:%v:R>", unixTime)
}
