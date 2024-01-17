package timestamps

import "strconv"

// Format inputed Unix epoch time into short date time timestamp
func ShortDateTime(time uint32) string {
	timeString := strconv.Itoa(int(time))
	return "<t:" + timeString + ":f>"
}

// Format inputed Unix epoch time into long date time timestamp
func LongDateTime(time uint32) string {
	timeString := strconv.Itoa(int(time))
	return "<t:" + timeString + ":F>"
}

// Format inputed Unix epoch time into short date timestamp
func ShortDate(time uint32) string {
	timeString := strconv.Itoa(int(time))
	return "<t:" + timeString + ":d>"
}

// Format inputed Unix epoch time into long date timestamp
func LongDate(time uint32) string {
	timeString := strconv.Itoa(int(time))
	return "<t:" + timeString + ":D>"
}

// Format inputed Unix epoch time into short time timestamp
func ShortTime(time uint32) string {
	timeString := strconv.Itoa(int(time))
	return "<t:" + timeString + ":t>"
}

// Format inputed Unix epoch time into long time timestamp
func LongTime(time uint32) string {
	timeString := strconv.Itoa(int(time))
	return "<t:" + timeString + ":T>"
}

// Format inputed Unix epoch time into relative timestamp
func Relative(time uint32) string {
	timeString := strconv.Itoa(int(time))
	return "<t:" + timeString + ":R>"
}
