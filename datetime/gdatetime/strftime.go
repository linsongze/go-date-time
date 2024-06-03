package gdatetime

import (
	"fmt"
	"strings"
	"time"
)

var longDayNames = []string{
	"Sunday",
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
}

var shortDayNames = []string{
	"Sun",
	"Mon",
	"Tue",
	"Wed",
	"Thu",
	"Fri",
	"Sat",
}

var shortMonthNames = []string{
	"---",
	"Jan",
	"Feb",
	"Mar",
	"Apr",
	"May",
	"Jun",
	"Jul",
	"Aug",
	"Sep",
	"Oct",
	"Nov",
	"Dec",
}

var longMonthNames = []string{
	"---",
	"January",
	"February",
	"March",
	"April",
	"May",
	"June",
	"July",
	"August",
	"September",
	"October",
	"November",
	"December",
}

func weekNumber(t *time.Time, char int) int {
	weekday := int(t.Weekday())

	if char == 'W' {
		// Monday as the first day of the week
		if weekday == 0 {
			weekday = 6
		} else {
			weekday -= 1
		}
	}

	return (t.YearDay() + 6 - weekday) / 7
}
func formatDateElement(t *time.Time, elem rune) string {
	switch elem {
	case 'a':
		return shortDayNames[t.Weekday()]
	case 'A':
		return longDayNames[t.Weekday()]
	case 'w':
		return fmt.Sprintf("%d", t.Weekday())
	case 'd':
		return fmt.Sprintf("%02d", t.Day())
	case 'b':
		return shortMonthNames[t.Month()]
	case 'B':
		return longMonthNames[t.Month()]
	case 'm':
		return fmt.Sprintf("%02d", t.Month())
	case 'y':
		return fmt.Sprintf("%02d", t.Year()%100)
	case 'Y':
		return fmt.Sprintf("%04d", t.Year())
	case 'H':
		return fmt.Sprintf("%02d", t.Hour())
	case 'I':
		hour := t.Hour()
		if hour == 0 || hour == 12 {
			return "12"
		}
		return fmt.Sprintf("%02d", hour%12)
	case 'p':
		if t.Hour() < 12 {
			return "AM"
		}
		return "PM"
	case 'M':
		return fmt.Sprintf("%02d", t.Minute())
	case 'S':
		return fmt.Sprintf("%02d", t.Second())
	case 'f':
		return fmt.Sprintf("%06d", t.Nanosecond()/1000)
	case 'z':
		return t.Format("-0700")
	case 'Z':
		return t.Format("MST")
	case 'j':
		return fmt.Sprintf("%03d", t.YearDay())
	case 'U':
		return fmt.Sprintf("%02d", weekNumber(t, 'U'))
	case 'W':
		return fmt.Sprintf("%02d", weekNumber(t, 'W'))
	case 'c':
		return t.Format("Mon Jan 2 15:04:05 2006")
	case 'x':
		return fmt.Sprintf("%02d/%02d/%02d", t.Month(), t.Day(), t.Year()%100)
	case 'X':
		return fmt.Sprintf("%02d:%02d:%02d", t.Hour(), t.Minute(), t.Second())
	default:
		return ""
	}
}

// Strftime like C style datez format
func (gtd *GDateTime) Strftime(f string) string {
	var builder strings.Builder
	format := []rune(f)

	for i := 0; i < len(format); i++ {
		if format[i] == '%' && i < len(format)-1 {
			i++
			formatted := formatDateElement(&gtd.t, format[i])
			if format[i] == '%' { // Handle the '%%' case
				builder.WriteRune('%')
			} else {
				builder.WriteString(formatted)
			}
		} else {
			builder.WriteRune(format[i])
		}
	}

	return builder.String()
}
