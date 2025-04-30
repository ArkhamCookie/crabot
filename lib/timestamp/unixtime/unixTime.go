package unixtime

import (
	"errors"
	"time"
)

func UnixTime(year, month, day, hour, minute, second int, timezone string) (int64, error) {
	var err error
	var gottenTimezone *time.Location
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

	gottenTimezone, err = time.LoadLocation(timezone)
	if err != nil {
		return 0, err
	}

	wantedTime := time.Date(year, timeMonth, day, hour, hour, minute, second, gottenTimezone).Unix()

	return wantedTime, nil
}
