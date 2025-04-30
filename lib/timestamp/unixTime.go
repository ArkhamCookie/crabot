package timestamp

import (
	"errors"
	"strings"
	"time"
)

func UnixTime(year, month, day, hour, minute, second int, timezone string) (int64, error) {
	var timeMonth time.Month

	switch month {
	case 1:
		timeMonth = time.January
	case 2:
		timeMonth = time.February
	case 3:
		timeMonth = time.March
	case 4:
		timeMonth = time.April
	case 5:
		timeMonth = time.May
	case 6:
		timeMonth = time.June
	case 7:
		timeMonth = time.July
	case 8:
		timeMonth = time.August
	case 9:
		timeMonth = time.September
	case 10:
		timeMonth = time.October
	case 11:
		timeMonth = time.November
	case 12:
		timeMonth = time.December
	default:
		return 0, errors.New("invaild month given")
	}

	var location *time.Location

	timezone = strings.ToLower(timezone)

	switch timezone {
	case "utc":
		location = time.UTC
	default:
		return 0, errors.New("unknown timezone given")
	}

	wantedTime := time.Date(year, timeMonth, day, hour, hour, minute, second, location).Unix()

	return wantedTime, nil
}
