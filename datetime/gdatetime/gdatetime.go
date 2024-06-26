package gdatetime

import (
	"errors"
	"github.com/linsongze/go-date-time/datetime/timeconst"
	"github.com/linsongze/go-date-time/datetime/timeunit"
	"time"
)

type GDateTime struct {
	t time.Time
}

// CreateGDateTime create function
func Create(t time.Time) *GDateTime {
	return &GDateTime{t: t}
}

// Now create now Time
func Now() *GDateTime {
	return &GDateTime{t: time.Now()}
}

// Parse dateStr use time layout
func Parse(dateStr string, layout string) (*GDateTime, error) {
	t, err := time.Parse(layout, dateStr)
	if err != nil {
		return nil, err
	}
	return Create(t), nil
}

// FormUnixTimestamp Create GDateTime use unix timestamp
func FormUnixTimestamp(timestamp int64, nano int64) (*GDateTime, error) {
	t := time.Unix(timestamp, nano)
	return Create(t), nil
}

// FormMillisTimestamp Create GDateTime use unix millis timestamp
func FormMillisTimestamp(timestamp int64) (*GDateTime, error) {
	t := time.UnixMilli(timestamp)
	return Create(t), nil
}

// isLeapYear determines if the given year is a leap year.
func isLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || year%400 == 0
}

// DaysInMonth returns the number of days in a given month for a specific year.
func DaysInMonth(year, month int) int {
	switch time.Month(month) {
	case time.January, time.March, time.May, time.July, time.August, time.October, time.December:
		return 31
	case time.April, time.June, time.September, time.November:
		return 30
	case time.February:
		if isLeapYear(year) {
			return 29
		}
		return 28
	default:
		return 0 // invalid month
	}
}

// daysInYear returns the number of days  for a specific year.
func daysInYear(year int) int {
	if isLeapYear(year) {
		return 366
	}
	return 365
}

// Of Obtains an instance of local zone GDateTime from year, month, day, hour, minute, second and nanosecond
func Of(year, month, dayOfMonth, hour, minute, second, nanoOfSecond int) (*GDateTime, error) {
	if year < 0 {
		return nil, errors.New("invalid year")
	}
	if month < 1 || month > 12 {
		return nil, errors.New("month must be between 1 and 12")
	}
	if dayOfMonth < 1 || dayOfMonth > 31 {
		return nil, errors.New("day must be between 1 and 31")
	}
	if hour < 0 || hour > 23 {
		return nil, errors.New("hour must be between 0 and 23")
	}
	if minute < 0 || minute > 59 {
		return nil, errors.New("minute must be between 0 and 59")
	}
	if second < 0 || second > 59 {
		return nil, errors.New("second must be between 0 and 59")
	}
	if nanoOfSecond < 0 || nanoOfSecond > timeconst.MAX_NANO {
		return nil, errors.New("nanosecond must be between 0 and timeconst.MAX_NANO")
	}
	t := time.Date(year, time.Month(month), dayOfMonth, hour, minute, second, nanoOfSecond, time.Local)
	return Create(t), nil
}

// Of2 Obtains an instance of local zone GDateTime from year, month, day, hour, minute, second
func Of2(year, month, dayOfMonth, hour, minute, second int) (*GDateTime, error) {
	return Of(year, month, dayOfMonth, hour, minute, second, 0)
}

// Of3 Obtains an instance of local zone GDateTime from year, month, day, hour, minute
func Of3(year, month, dayOfMonth, hour, minute int) (*GDateTime, error) {
	return Of(year, month, dayOfMonth, hour, minute, 0, 0)
}

func (gdt *GDateTime) ToTime() time.Time {
	return gdt.t
}

// GSecondTimestamp ten-digit number
func (gdt *GDateTime) GetSecondTimestamp() int64 {
	return gdt.t.Unix()
}

// GMillSecondTimestamp GSecondTimestamp thirteen-digit number
func (gdt *GDateTime) GetMillSecondTimestamp() int64 {
	return gdt.t.UnixMilli() / 1000000
}

// GetYear returns the year of the GDateTime.
func (gdt *GDateTime) GetYear() int {
	return gdt.t.Year()
}

// GetMonth returns the month of the GDateTime as an integer (1-12).
func (gdt *GDateTime) GetMonth() int {
	return int(gdt.t.Month())
}

// GetDayOfMonth returns the day of the month of the GDateTime.
func (gdt *GDateTime) GetDayOfMonth() int {
	return gdt.t.Day()
}

// GetDayOfYear returns the day of the year of the GDateTime.
func (gdt *GDateTime) GetDayOfYear() int {
	return gdt.t.YearDay()
}

// GetDayOfWeek returns the day of the week of the GDateTime as an integer (0-6).
// Sunday = 0, Monday = 1, ..., Saturday = 6
func (gdt *GDateTime) GetDayOfWeek() int {
	return int(gdt.t.Weekday())
}

// GetHour return hour of the GDateTime
func (gdt *GDateTime) GetHour() int {
	return gdt.t.Hour()
}

// GetMinute return Minute of the GDateTime
func (gdt *GDateTime) GetMinute() int {
	return gdt.t.Minute()
}

// GetSecond return Second of the GDateTime
func (gdt *GDateTime) GetSecond() int {
	return gdt.t.Second()
}

// GetNano return Nano of the GDateTime
func (gdt *GDateTime) GetNano() int {
	return gdt.t.Nanosecond()
}

// WeekOfYearStartsFromJan1 calculates the week number starting from January 1st, considering the week containing January 1st as the first week of the year.
func (gdt *GDateTime) WeekOfYearStartsFromJan1() int {
	yearStart := time.Date(gdt.t.Year(), time.January, 1, 0, 0, 0, 0, gdt.t.Location())
	yearStartWeekday := yearStart.Weekday()

	daysSinceYearStart := gdt.t.YearDay() - 1
	weekOfYear := (daysSinceYearStart+int(yearStartWeekday))/7 + 1
	return weekOfYear
}

// FirstFullWeekOfYear calculates the week number such that the first full week (Monday through Sunday) that entirely falls within the new year counts as the first week of the year.
func (gdt *GDateTime) FirstFullWeekOfYear() int {
	jan1 := time.Date(gdt.t.Year(), time.January, 1, 0, 0, 0, 0, gdt.t.Location())
	offset := int(time.Monday - jan1.Weekday())
	if offset > 0 {
		offset -= 7
	}
	firstFullWeekMonday := jan1.AddDate(0, 0, offset+7)

	if gdt.t.Before(firstFullWeekMonday) {
		return 0
	}
	daysSinceFirstFullWeek := gdt.t.Sub(firstFullWeekMonday).Hours() / 24
	return int(daysSinceFirstFullWeek)/7 + 1
}

// WeekOfYearISO8601 calculates the week number according to the ISO 8601 standard, where the week containing the first Thursday of the year is considered the first week.
func (gdt *GDateTime) WeekOfYearISO8601() int {
	_, week := gdt.t.ISOWeek()
	return week
}

// WithYear withYear creates a new GDateTime instance with the specified year, keeping other components the same.
func (gdt *GDateTime) WithYear(year int) (*GDateTime, error) {
	if year < 0 {
		return nil, errors.New("invalid year")
	}
	newTime := time.Date(year, gdt.t.Month(), gdt.t.Day(), gdt.t.Hour(), gdt.t.Minute(), gdt.t.Second(), gdt.t.Nanosecond(), gdt.t.Location())
	return Create(newTime), nil
}

// WithMonth withMonth creates a new GDateTime instance with the specified month, keeping other components the same.
func (gdt *GDateTime) WithMonth(month int) (*GDateTime, error) {
	if month < 1 || month > 12 {
		return nil, errors.New("month must be between 1 and 12")
	}
	newTime := time.Date(gdt.t.Year(), time.Month(month), gdt.t.Day(), gdt.t.Hour(), gdt.t.Minute(), gdt.t.Second(), gdt.t.Nanosecond(), gdt.t.Location())
	return Create(newTime), nil
}

// WithDayOfMonth Returns a copy of this GDateTime with the day-of-month altered.
func (gdt *GDateTime) WithDayOfMonth(day int) (*GDateTime, error) {
	if day < 1 || day > DaysInMonth(gdt.t.Year(), int(gdt.t.Month())) {
		return nil, errors.New("invalid day for the month") // invalid day for the month
	}
	newTime := time.Date(gdt.t.Year(), gdt.t.Month(), day, gdt.t.Hour(), gdt.t.Minute(), gdt.t.Second(), gdt.t.Nanosecond(), gdt.t.Location())
	return Create(newTime), nil
}

// WithDayOfYear Returns a copy of this GDateTime with the day-of-year altered.
func (gdt *GDateTime) WithDayOfYear(day int) (*GDateTime, error) {
	if day < 1 || day > daysInYear(gdt.t.Year()) {
		return nil, errors.New("invalid day for the year") // invalid day for the Year
	}
	// Set the new date using the day of the year
	newTime := time.Date(gdt.t.Year(), 1, 1, gdt.t.Hour(), gdt.t.Minute(), gdt.t.Second(), gdt.t.Nanosecond(), gdt.t.Location())
	newTime = newTime.AddDate(0, 0, day-1) // day-1 because January 1 is day 1
	return Create(newTime), nil
}

// WithHour Returns a copy of this GDateTime with the hour-of-day altered.
func (gdt *GDateTime) WithHour(hour int) (*GDateTime, error) {
	if hour < 0 || hour > 23 {
		return nil, errors.New("month must be between 0 and 23") //
	}
	newTime := time.Date(gdt.t.Year(), gdt.t.Month(), gdt.t.Day(), hour, gdt.t.Minute(), gdt.t.Second(), gdt.t.Nanosecond(), gdt.t.Location())
	return Create(newTime), nil
}

// WithMinute Returns a copy of this GDateTime with the minute altered.
func (gdt *GDateTime) WithMinute(minute int) (*GDateTime, error) {
	if minute < 0 || minute > 59 {
		return nil, errors.New("minute must be between 0 and 59") //
	}
	newTime := time.Date(gdt.t.Year(), gdt.t.Month(), gdt.t.Day(), gdt.t.Hour(), minute, gdt.t.Second(), gdt.t.Nanosecond(), gdt.t.Location())
	return Create(newTime), nil
}

// WithSecond Returns a copy of this GDateTime with the second altered.
func (gdt *GDateTime) WithSecond(second int) (*GDateTime, error) {
	if second < 0 || second > 59 {
		return nil, errors.New("second must be between 0 and 59") //
	}
	newTime := time.Date(gdt.t.Year(), gdt.t.Month(), gdt.t.Day(), gdt.t.Hour(), gdt.t.Minute(), second, gdt.t.Nanosecond(), gdt.t.Location())
	return Create(newTime), nil
}

// WithNano Returns a copy of this GDateTime with the nanosecond altered.
func (gdt *GDateTime) WithNano(nano int) (*GDateTime, error) {
	if nano < 0 || nano > timeconst.MAX_NANO {
		return nil, errors.New("nanosecond must be between 0 and timeconst.MAX_NANO") //
	}
	newTime := time.Date(gdt.t.Year(), gdt.t.Month(), gdt.t.Day(), gdt.t.Hour(), gdt.t.Minute(), gdt.t.Second(), nano, gdt.t.Location())
	return Create(newTime), nil
}

// TruncateTo truncates the GDateTime to the specified unit.
func (gdt *GDateTime) TruncateTo(unit timeunit.TimeUnit) *GDateTime {
	var d time.Duration
	switch unit {
	case timeunit.NANOS:
		d = time.Nanosecond
	case timeunit.SECONDS:
		d = time.Second
	case timeunit.MINUTES:
		d = time.Minute
	case timeunit.HOURS:
		d = time.Hour
	default:
		return gdt // If the unit is not recognized, return the original GDateTime
	}
	return Create(gdt.t.Truncate(d))
}

// PlusYears adds the specified number of years to the GDateTime, returning a new GDateTime instance.
func (gdt *GDateTime) PlusYears(years int) *GDateTime {
	newTime := gdt.t.AddDate(years, 0, 0)
	return Create(newTime)
}

// PlusMonths  adds the specified number of months to the GDateTime, returning a new GDateTime instance.
func (gdt *GDateTime) PlusMonths(months int) *GDateTime {
	newTime := gdt.t.AddDate(0, months, 0)
	return Create(newTime)
}

// PlusWeeks adds the specified number of weeks to the GDateTime, returning a new GDateTime instance.
func (gdt *GDateTime) PlusWeeks(weeks int) *GDateTime {
	days := weeks * 7 // Convert weeks to days
	newTime := gdt.t.AddDate(0, 0, days)
	return Create(newTime)
}

// PlusDays adds the specified number of days to the GDateTime, returning a new GDateTime instance.
func (gdt *GDateTime) PlusDays(days int) *GDateTime {
	newTime := gdt.t.AddDate(0, 0, days)
	return Create(newTime)
}

// PlusHours adds the specified number of hours to the GDateTime, returning a new GDateTime instance.
func (gdt *GDateTime) PlusHours(hours int) *GDateTime {
	newTime := gdt.t.Add(time.Duration(hours) * time.Hour)
	return Create(newTime)
}

// PlusMinutes adds the specified number of minutes to the GDateTime, returning a new GDateTime instance.
func (gdt *GDateTime) PlusMinutes(minutes int) *GDateTime {
	newTime := gdt.t.Add(time.Duration(minutes) * time.Minute)
	return Create(newTime)
}

// PlusSeconds adds the specified number of seconds to the GDateTime, returning a new GDateTime instance.
func (gdt *GDateTime) PlusSeconds(seconds int) *GDateTime {
	newTime := gdt.t.Add(time.Duration(seconds) * time.Second)
	return Create(newTime)
}

// PlusNanos adds the specified number of nanoseconds to the GDateTime, returning a new GDateTime instance.
func (gdt *GDateTime) PlusNanos(nanos int) *GDateTime {
	newTime := gdt.t.Add(time.Duration(nanos) * time.Nanosecond)
	return Create(newTime)
}

// Plus adjusts the GDateTime based on the amount and unit specified.
func (gdt *GDateTime) Plus(amountToAdd int, unit timeunit.TimeUnit) *GDateTime {
	switch unit {
	case timeunit.NANOS:
		return gdt.PlusNanos(amountToAdd)
	case timeunit.MICROS:
		// Convert micros to days and nanoseconds
		return gdt.PlusDays(amountToAdd / 86400000000).PlusNanos((amountToAdd % 86400000000) * 1000)
	case timeunit.MILLIS:
		// Convert millis to days and nanoseconds
		return gdt.PlusDays(amountToAdd / 86400000).PlusNanos((amountToAdd % 86400000) * 1000000)
	case timeunit.SECONDS:
		return gdt.PlusSeconds(amountToAdd)
	case timeunit.MINUTES:
		return gdt.PlusMinutes(amountToAdd)
	case timeunit.HOURS:
		return gdt.PlusHours(amountToAdd)
	case timeunit.HALF_DAYS:
		// Convert half-days to days and hours
		return gdt.PlusDays(amountToAdd / 2).PlusHours((amountToAdd % 2) * 12)
	default:
		return gdt // Fallback, no operation if the unit is not recognized
	}
}

// Minus Returns a copy of this date-time with the specified amount subtracted.
func (gdt *GDateTime) Minus(years int, unit timeunit.TimeUnit) *GDateTime {
	return gdt.Plus(-years, unit)
}

// MinusYears subtracts the specified number of years from the GDateTime, returning a new GDateTime instance.
func (gdt *GDateTime) MinusYears(years int) *GDateTime {
	return gdt.PlusYears(-years)
}

// MinusMonths subtracts the specified number of months from the GDateTime, returning a new GDateTime instance.
func (gdt *GDateTime) MinusMonths(months int) *GDateTime {
	return gdt.PlusMonths(-months)
}

// MinusWeeks subtracts the specified number of weeks from the GDateTime, returning a new GDateTime instance.
func (gdt *GDateTime) MinusWeeks(weeks int) *GDateTime {
	return gdt.PlusWeeks(-weeks)
}

// MinusDays subtracts the specified number of days from the GDateTime, returning a new GDateTime instance.
func (gdt *GDateTime) MinusDays(days int) *GDateTime {
	return gdt.PlusDays(-days)
}

// MinusHours subtracts the specified number of hours from the GDateTime, returning a new GDateTime instance.
func (gdt *GDateTime) MinusHours(hours int) *GDateTime {
	newTime := gdt.t.Add(time.Duration(-hours) * time.Hour)
	return Create(newTime)
}

// MinusMinutes subtracts the specified number of minutes from the GDateTime, returning a new GDateTime instance.
func (gdt *GDateTime) MinusMinutes(minutes int) *GDateTime {
	newTime := gdt.t.Add(time.Duration(-minutes) * time.Minute)
	return Create(newTime)
}

// MinusSeconds subtracts the specified number of seconds from the GDateTime, returning a new GDateTime instance.
func (gdt *GDateTime) MinusSeconds(seconds int) *GDateTime {
	newTime := gdt.t.Add(time.Duration(-seconds) * time.Second)
	return Create(newTime)
}

// MinusNanos subtracts the specified number of nanoseconds from the GDateTime, returning a new GDateTime instance.
func (gdt *GDateTime) MinusNanos(nanos int) *GDateTime {
	newTime := gdt.t.Add(time.Duration(-nanos) * time.Nanosecond)
	return Create(newTime)
}

// StartOfMonth returns a new GDateTime instance set to the start of the month of the original GDateTime.
func (gdt *GDateTime) StartOfMonth() *GDateTime {
	startOfMonth := time.Date(gdt.t.Year(), gdt.t.Month(), 1, 0, 0, 0, 0, gdt.t.Location())
	return Create(startOfMonth)
}

// EndOfMonth returns a new GDateTime instance set to the end of the month of the original GDateTime.
func (gdt *GDateTime) EndOfMonth() *GDateTime {
	endOfMonth := time.Date(gdt.t.Year(), gdt.t.Month(), DaysInMonth(gdt.t.Year(), int(gdt.t.Month())), 23, 59, 59, timeconst.MAX_NANO, gdt.t.Location())
	return Create(endOfMonth)
}

// StartOfWeek returns a new GDateTime instance set to the start of the week of the original GDateTime.
func (gdt *GDateTime) StartOfWeek() *GDateTime {
	// Weekday returns Sunday as 0
	offset := int(gdt.t.Weekday())
	startOfWeek := gdt.t.AddDate(0, 0, -offset) // Adjust to start of the week, assuming week starts on Sunday
	startOfWeek = time.Date(startOfWeek.Year(), startOfWeek.Month(), startOfWeek.Day(), 0, 0, 0, 0, startOfWeek.Location())
	return Create(startOfWeek)
}

// EndOfWeek returns a new GDateTime instance set to the end of the week of the original GDateTime.
func (gdt *GDateTime) EndOfWeek() *GDateTime {
	offset := 6 - int(gdt.t.Weekday()) // Assuming week starts on Sunday, end on Saturday
	endOfWeek := gdt.t.AddDate(0, 0, offset)
	endOfWeek = time.Date(endOfWeek.Year(), endOfWeek.Month(), endOfWeek.Day(), 23, 59, 59, timeconst.MAX_NANO, endOfWeek.Location())
	return Create(endOfWeek)
}

// StartOfWeekFromMonday returns a new GDateTime instance set to the start of the week (starting from Monday) of the original GDateTime.
func (gdt *GDateTime) StartOfWeekFromMonday() *GDateTime {
	// Weekday returns Sunday as 0, hence Monday as 1
	offset := (int(gdt.t.Weekday()) + 6) % 7 // Adjust to start of the week, assuming week starts on Monday
	startOfWeek := gdt.t.AddDate(0, 0, -offset)
	startOfWeek = time.Date(startOfWeek.Year(), startOfWeek.Month(), startOfWeek.Day(), 0, 0, 0, 0, startOfWeek.Location())
	return Create(startOfWeek)
}

// EndOfWeekFromMonday returns a new GDateTime instance set to the end of the week (ending on Sunday) of the original GDateTime.
func (gdt *GDateTime) EndOfWeekFromMonday() *GDateTime {
	offset := 7 - int(gdt.t.Weekday()) // Assuming week starts on Monday, end on Sunday
	if offset == 7 {
		offset = 0
	}
	endOfWeek := gdt.t.AddDate(0, 0, offset)
	endOfWeek = time.Date(endOfWeek.Year(), endOfWeek.Month(), endOfWeek.Day(), 23, 59, 59, 999999999, endOfWeek.Location())
	return Create(endOfWeek)
}

// StartOfDay returns a new GDateTime instance set to the start of the day (00:00:00) of the original GDateTime.
func (gdt *GDateTime) StartOfDay() *GDateTime {
	startOfDay := time.Date(gdt.t.Year(), gdt.t.Month(), gdt.t.Day(), 0, 0, 0, 0, gdt.t.Location())
	return Create(startOfDay)
}

// EndOfDay returns a new GDateTime instance set to the end of the day (23:59:59) of the original GDateTime.
func (gdt *GDateTime) EndOfDay() *GDateTime {
	endOfDay := time.Date(gdt.t.Year(), gdt.t.Month(), gdt.t.Day(), 23, 59, 59, timeconst.MAX_NANO, gdt.t.Location())
	return Create(endOfDay)
}

// ConvertToZone change GDateTime zone
func (gdt *GDateTime) ConvertToZone(loc time.Location) *GDateTime {
	t := gdt.t.In(&loc)
	return Create(t)
}

// ResetZoneToDefault reset zone to local
func (gdt *GDateTime) ResetZoneToDefault() *GDateTime {
	t := gdt.t.In(time.Local)
	return Create(t)
}

// EqualDate checks if the year, month, and day of two GDateTime instances are the same.
func (gdt *GDateTime) EqualDate(other *GDateTime) bool {
	// Compare the year, month, and day components of the time.Time field of both GDateTime instances
	return gdt.t.Year() == other.t.Year() && gdt.t.Month() == other.t.Month() && gdt.t.Day() == other.t.Day()
}

// CompareTo compares this GDateTime instance with another to determine the chronological order.
func (gdt *GDateTime) CompareTo(other *GDateTime) int {
	if gdt.t.Equal(other.t) {
		return 0
	} else if gdt.t.After(other.t) {
		return 1
	} else {
		return -1
	}
}

// CompareDate compares the date (year, month, day) components of two GDateTime instances.
func (gdt *GDateTime) CompareDate(other *GDateTime) int {
	if gdt.t.Year() > other.t.Year() {
		return 1
	} else if gdt.t.Year() < other.t.Year() {
		return -1
	}
	if gdt.t.Month() > other.t.Month() {
		return 1
	} else if gdt.t.Month() < other.t.Month() {
		return -1
	}
	if gdt.t.Day() > other.t.Day() {
		return 1
	} else if gdt.t.Day() < other.t.Day() {
		return -1
	}
	return 0
}

// CompareTime compares the time (hour, minute, second) components of two GDateTime instances.
func (gdt *GDateTime) CompareTime(other *GDateTime) int {
	if gdt.t.Hour() > other.t.Hour() {
		return 1
	} else if gdt.t.Hour() < other.t.Hour() {
		return -1
	}
	if gdt.t.Minute() > other.t.Minute() {
		return 1
	} else if gdt.t.Minute() < other.t.Minute() {
		return -1
	}
	if gdt.t.Second() > other.t.Second() {
		return 1
	} else if gdt.t.Second() < other.t.Second() {
		return -1
	}
	return 0
}

// IsBefore checks if this GDateTime instance is before the provided GDateTime instance.
func (gdt *GDateTime) IsBefore(other *GDateTime) bool {
	return gdt.t.Before(other.t)
}

// IsAfter checks if this GDateTime instance is after the provided GDateTime instance.
func (gdt *GDateTime) IsAfter(other *GDateTime) bool {
	return gdt.t.After(other.t)
}

// ToFormatString formats the GDateTime based on time package layout specifier.
func (gdt *GDateTime) ToFormatString(layout string) string {
	return gdt.t.Format(layout)
}

// ToRFC1123 formats the GDateTime according to RFC 1123 format.
func (gdt *GDateTime) ToRFC1123() string {
	return gdt.t.Format(time.RFC1123)
}

// ToISO8601 formats the GDateTime according to ISO 8601 format.
func (gdt *GDateTime) ToISO8601() string {
	return gdt.t.Format(time.RFC3339) // RFC3339 is compatible with ISO 8601 for date-time formats
}

// ToDateTimeString return format yyyy-MM-dd HH:mm:ss
func (gdt *GDateTime) ToDateTimeString() string {
	return gdt.ToFormatString("2006-01-02 15:04:05")
}

// ToDateString return format yyyy-MM-dd
func (gdt *GDateTime) ToDateString() string {
	return gdt.ToFormatString("2006-01-02")
}

// YearsBetween calculates the difference in full years between two GDateTime instances,
// subtracting a year if the end date is before the start date's month and day.
func (gdt *GDateTime) YearsBetween(end *GDateTime) int {
	yearDiff := end.GetYear() - gdt.GetYear()
	if end.GetMonth() < gdt.GetMonth() || (end.GetMonth() == gdt.GetMonth() && end.GetDayOfMonth() < gdt.GetDayOfMonth()) {
		yearDiff--
	}
	return yearDiff
}

// MonthsBetween calculates the difference in full months between two GDateTime instances,
// subtracting a month if the end date's day is before the start date's day.
func (gdt *GDateTime) MonthsBetween(end *GDateTime) int {
	monthDiff := int(end.GetMonth() - gdt.GetMonth() + 12*(end.GetYear()-gdt.GetYear()))
	if end.GetDayOfMonth() < gdt.GetDayOfMonth() {
		monthDiff--
	}
	return monthDiff
}

// DaysBetween calculates the difference in full days between two GDateTime instances,
// based on the actual time difference.
func (gdt *GDateTime) DaysBetween(end *GDateTime) int {
	return int(end.ToTime().Sub(gdt.ToTime()).Hours() / 24)
}

// HoursBetween calculates the difference in hours between two GDateTime instances.
func (gdt *GDateTime) HoursBetween(end *GDateTime) int {
	return int(end.ToTime().Sub(gdt.ToTime()).Hours())
}

// MinutesBetween calculates the difference in minutes between two GDateTime instances.
func (gdt *GDateTime) MinutesBetween(end *GDateTime) int {
	return int(end.ToTime().Sub(gdt.ToTime()).Minutes())
}

// SecondsBetween calculates the difference in seconds between two GDateTime instances.
func (gdt *GDateTime) SecondsBetween(end *GDateTime) int {
	return int(end.ToTime().Sub(gdt.ToTime()).Seconds())
}

// IsWithinRange checks if the GDateTime is within the range specified by start and end.
func (gdt *GDateTime) IsWithinRange(start, end *GDateTime) bool {
	return !gdt.t.Before(start.t) && !gdt.t.After(end.t)
}

// Monday returns the date of Monday for the current week, with weeks starting on Monday.
func (gdt *GDateTime) Monday() *GDateTime {
	return gdt.StartOfWeekFromMonday()
}

// Tuesday returns the date of Tuesday for the current week, with weeks starting on Monday.
func (gdt *GDateTime) Tuesday() *GDateTime {
	return gdt.StartOfWeekFromMonday().PlusDays(1)
}

// Wednesday returns the date of Wednesday for the current week, with weeks starting on Monday.
func (gdt *GDateTime) Wednesday() *GDateTime {
	return gdt.StartOfWeekFromMonday().PlusDays(2)
}

// Thursday returns the date of Thursday for the current week, with weeks starting on Monday.
func (gdt *GDateTime) Thursday() *GDateTime {
	return gdt.StartOfWeekFromMonday().PlusDays(3)
}

// Friday returns the date of Friday for the current week, with weeks starting on Monday.
func (gdt *GDateTime) Friday() *GDateTime {
	return gdt.StartOfWeekFromMonday().PlusDays(4)
}

// Saturday returns the date of Saturday for the current week, with weeks starting on Monday.
func (gdt *GDateTime) Saturday() *GDateTime {
	return gdt.StartOfWeekFromMonday().PlusDays(5)
}

// Sunday returns the date of Sunday for the current week, with weeks starting on Monday.
func (gdt *GDateTime) Sunday() *GDateTime {
	return gdt.StartOfWeekFromMonday().PlusDays(6)
}
