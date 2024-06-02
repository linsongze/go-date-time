# go-date-time
Simple Time Handling Utilities for Go


# Usage
-------------
```
import (
	"fmt"
	"time"

	"github.com/linsongze/go-date-time/datetime/gdatetime" 
)

func main() {
	// 创建当前时间的实例
	// Create an instance for the current time
	now := gdatetime.Now()
	fmt.Println("Current Time:", now.Format("2006-01-02 15:04:05"))
	// Expected: Current Time: <current time>

	// 创建具体时间的实例
	// Create an instance for a specific time
	specificTime, _ := gdatetime.Of(2023, 12, 25, 10, 30, 0, 0)
	fmt.Println("Specific Time:", specificTime.Format("2006-01-02 15:04:05"))
	// Expected: Specific Time: 2023-12-25 10:30:00

	// 获取年、月、日
	// Get year, month, and day
	fmt.Println("Year:", specificTime.GetYear())
	// Expected: Year: 2023
	fmt.Println("Month:", specificTime.GetMonth())
	// Expected: Month: 12
	fmt.Println("Day of Month:", specificTime.GetDayOfMonth())
	// Expected: Day of Month: 25

	// 比较日期
	// Compare dates
	tomorrow := now.PlusDays(1)
	fmt.Println("Tomorrow is after today:", tomorrow.IsAfter(now))
	// Expected: Tomorrow is after today: true

	// 日期加减
	// Date addition and subtraction
	oneWeekLater := now.PlusWeeks(1)
	fmt.Println("One week later:", oneWeekLater.Format("2006-01-02"))
	// Expected: One week later: <date one week from now>

	// 日期的开始与结束
	// Start and end of the day
	startOfDay := now.StartOfDay()
	endOfDay := now.EndOfDay()
	fmt.Println("Start of today:", startOfDay.Format("2006-01-02 15:04:05"))
	// Expected: Start of today: <today's date> 00:00:00
	fmt.Println("End of today:", endOfDay.Format("2006-01-02 15:04:05"))
	// Expected: End of today: <today's date> 23:59:59
	
	// 获取当前月份的开始
	// Get the start of the current month
	startOfMonth := now.StartOfMonth()
	fmt.Println("Start of Month:", startOfMonth.Format("2006-01-02 15:04:05"))
	// Expected: Start of Month: <first day of current month> 00:00:00

	// 获取当前月份的结束
	// Get the end of the current month
	endOfMonth := now.EndOfMonth()
	fmt.Println("End of Month:", endOfMonth.Format("2006-01-02 15:04:05"))
	// Expected: End of Month: <last day of current month> 23:59:59

	// 获取当前周的开始
	// Get the start of the current week
	startOfWeek := now.StartOfWeek()
	fmt.Println("Start of Week:", startOfWeek.Format("2006-01-02 15:04:05"))
	// Expected: Start of Week: <Sunday of the current week> 00:00:00

	// 获取当前周的结束
	// Get the end of the current week
	endOfWeek := now.EndOfWeek()
	fmt.Println("End of Week:", endOfWeek.Format("2006-01-02 15:04:05"))
	// Expected: End of Week: <Saturday of the current week> 23:59:59
	
	// 获取周一开始的当前周的起始日期
	// Get the start of the current week starting from Monday
	startOfWeek := now.StartOfWeekFromMonday()
	fmt.Println("Start of Week from Monday:", startOfWeek.Format("2006-01-02 15:04:05"))
	// 预计结果：Start of Week from Monday: <date of last Monday> 00:00:00

	// 获取以周一开始的当前周的结束日期
	// Get the end of the current week ending on Sunday
	endOfWeek := now.EndOfWeekFromMonday()
	fmt.Println("End of Week from Monday:", endOfWeek.Format("2006-01-02 15:04:05"))
	// 预计结果：End of Week from Monday: <date of next Sunday> 23:59:59

	// 时区转换
	// Time zone conversion
	location, _ := time.LoadLocation("America/New_York")
	timeInNY := now.SwitchZone(*location)
	fmt.Println("Current Time in New York:", timeInNY.Format("2006-01-02 15:04:05 MST"))
	// Expected: Current Time in New York: <time in NY>

	// 重置时区到本地
	// Reset time zone to local
	localTime := timeInNY.ResetZoneToDefault()
	fmt.Println("Time back in local:", localTime.Format("2006-01-02 15:04:05 MST"))
	// Expected: Time back in local: <local time>

	// 比较两个日期的年月日是否相同
	// Compare if two dates are the same in terms of year, month, and day
	sameDay := now.EqualDate(now)
	fmt.Println("Same day comparison:", sameDay)
	// Expected: Same day comparison: true

	// 比较时间顺序
	// Compare chronological order
	compareResult := now.CompareTo(tomorrow)
	fmt.Println("Compare today with tomorrow:", compareResult)
	// Expected: Compare today with tomorrow: -1

	// 另外的日期和时间比较
	// Additional date and time comparison
	compareDateResult := specificTime.CompareDate(now)
	fmt.Println("Compare specific date with today:", compareDateResult)
	// Expected: Compare specific date with today: 1 or -1 depending on the current date

	compareTimeResult := specificTime.CompareTime(now)
	fmt.Println("Compare specific time with current time:", compareTimeResult)
	// Expected: Compare specific time with current time: 1 or -1 depending on the current time

	// 检查日期是否在另一个日期之前或之后
	// Check if a date is before or after another date
	isBefore := specificTime.IsBefore(now)
	fmt.Println("Specific time is before now:", isBefore)
	// Expected: Specific time is before now: true or false

	isAfter := specificTime.IsAfter(now)
	fmt.Println("Specific time is after now:", isAfter)
	// Expected: Specific time is after now: true or false
}
```
# Function List
-----------
```
Create(t time.Time) *GDateTime // Creates a new GDateTime instance. (创建新实例)
Now() *GDateTime // Gets a GDateTime instance representing the current time. (获取当前时间实例)
Of(year, month, dayOfMonth, hour, minute, second, nanoOfSecond int) (*GDateTime, error) // Obtains an instance of GDateTime for a specific year, month, day, hour, minute, second, and nanosecond. (根据具体日期和时间创建实例)
Of2(year, month, dayOfMonth, hour, minute, second int) (*GDateTime, error) // Obtains an instance of GDateTime for a specific year, month, day, hour, minute, and second with nanosecond set to zero. (创建具体日期时间实例，纳秒为0)
Of3(year, month, dayOfMonth, hour, minute int) (*GDateTime, error) // Obtains an instance of GDateTime for a specific year, month, day, hour, and minute with second and nanosecond set to zero. (创建具体日期时间实例，秒和纳秒为0)
Parse(dateStr string, layout string) (*GDateTime, error) // Parses a date string using a specific time layout. (解析日期字符串)
FormUnixTimestamp(timestamp int64, nano int64) (*GDateTime, error) // Creates a GDateTime instance using a Unix timestamp with additional nanoseconds. (使用Unix时间戳创建实例)
FormMillisTimestamp(timestamp int64) (*GDateTime, error) // Creates a GDateTime instance using a Unix timestamp in milliseconds. (使用毫秒级Unix时间戳创建实例)

ToTime() time.Time // Converts a GDateTime instance to a time.Time type. (转换为time.Time类型)
GetSecondTimestamp() int64 // Gets the timestamp in seconds. (获取秒级时间戳)
GetMillSecondTimestamp() int64 // Gets the timestamp in milliseconds. (获取毫秒级时间戳)
GetYear() int // Gets the year. (获取年份)
GetMonth() int // Gets the month. (获取月份)
GetDayOfMonth() int // Gets the day of the month. (获取月中日)
GetDayOfYear() int // Gets the day of the year. (获取年中日)
GetDayOfWeek() int // Gets the day of the week (Sunday = 0). (获取星期几)
GetHour() int // Gets the hour. (获取小时)
GetMinute() int // Gets the minute. (获取分钟)
GetSecond() int // Gets the second. (获取秒)
GetNano() int // Gets the nanosecond. (获取纳秒)

WithYear(year int) (*GDateTime, error) // Sets the year. (设置年份)
WithMonth(month int) (*GDateTime, error) // Sets the month. (设置月份)
WithDayOfMonth(day int) (*GDateTime, error) // Sets the day of the month. (设置月中日)
WithDayOfYear(day int) (*GDateTime, error) // Sets the day of the year. (设置年中日)
WithHour(hour int) (*GDateTime, error) // Sets the hour. (设置小时)
WithMinute(minute int) (*GDateTime, error) // Sets the minute. (设置分钟)
WithSecond(second int) (*GDateTime, error) // Sets the second. (设置秒)
WithNano(nano int) (*GDateTime, error) // Sets the nanosecond. (设置纳秒)

PlusYears(years int) *GDateTime // Adds the specified number of years to the GDateTime. (增加年份)
PlusMonths(months int) *GDateTime // Adds the specified number of months to the GDateTime. (增加月份)
PlusWeeks(weeks int) *GDateTime // Adds the specified number of weeks to the GDateTime. (增加周数)
PlusDays(days int) *GDateTime // Adds the specified number of days to the GDateTime. (增加天数)
PlusHours(hours int) *GDateTime // Adds the specified number of hours to the GDateTime. (增加小时)
PlusMinutes(minutes int) *GDateTime // Adds the specified number of minutes to the GDateTime. (增加分钟)
PlusSeconds(seconds int) *GDateTime // Adds the specified number of seconds to the GDateTime. (增加秒数)
PlusNanos(nanos int) *GDateTime // Adds the specified number of nanoseconds to the GDateTime. (增加纳秒)
Plus(amountToAdd int, unit timeunit.TimeUnit) *GDateTime // Adjusts the time based on the specified amount and unit. (根据数量和单位调整时间)

Minus(years int, unit timeunit.TimeUnit) *GDateTime // Subtracts the specified amount and unit from the time. (根据数量和单位减少时间)
MinusYears(years int) *GDateTime // Subtracts the specified number of years from the GDateTime. (减少年份)
MinusMonths(months int) *GDateTime // Subtracts the specified number of months from the GDateTime. (减少月份)
MinusWeeks(weeks int) *GDateTime // Subtracts the specified number of weeks from the GDateTime. (减少周数)
MinusDays(days int) *GDateTime // Subtracts the specified number of days from the GDateTime. (减少天数)
MinusHours(hours int) *GDateTime // Subtracts the specified number of hours from the GDateTime. (减少小时)
MinusMinutes(minutes int) *GDateTime // Subtracts the specified number of minutes from the GDateTime. (减少分钟)
MinusSeconds(seconds int) *GDateTime // Subtracts the specified number of seconds from the GDateTime. (减少秒数)
MinusNanos(nanos int) *GDateTime // Subtracts the specified number of nanoseconds from the GDateTime. (减少纳秒)

StartOfMonth() *GDateTime // Sets the date to the start of the month. (设置为月初)
EndOfMonth() *GDateTime // Sets the date to the end of the month. (设置为月末)
StartOfWeek() *GDateTime // Sets the date to the start of the week (week starts on Sunday). (设置为周初)
EndOfWeek() *GDateTime // Sets the date to the end of the week (week ends on Saturday). (设置为周末)
StartOfWeekFromMonday() *GDateTime // Sets the date to the start of the week (week starts on Monday). (设置为从周一开始的周初)
EndOfWeekFromMonday() *GDateTime // Sets the date to the end of the week (week ends on Sunday). (设置为从周一开始的周末)
StartOfDay() *GDateTime // Sets the date to the start of the day. (设置为当天开始)
EndOfDay() *GDateTime // Sets the date to the end of the day. (设置为当天结束)

SwitchZone(loc time.Location) *GDateTime // Changes the time zone of the GDateTime. (更改时区)
ResetZoneToDefault() *GDateTime // Resets the time zone to local. (重置为本地时区)

EqualDate(other *GDateTime) bool // Checks if the year, month, and day of two GDateTime instances are the same. (比较年月日是否相同)
CompareTo(other *GDateTime) int // Compares two GDateTime instances to determine their chronological order. (比较两个实例的时间顺序)
CompareDate(other *GDateTime) int // Compares the date components (year, month, day) of two GDateTime instances. (比较两个实例的日期顺序)
CompareTime(other *GDateTime) int // Compares the time components (hour, minute, second) of two GDateTime instances. (比较两个实例的时间顺序)
IsBefore(other *GDateTime) bool // Checks if this GDateTime instance is before the provided GDateTime instance. (判断是否早于另一个实例)
IsAfter(other *GDateTime) bool // Checks if this GDateTime instance is after the provided GDateTime instance. (判断是否晚于另一个实例)

Format(layout string) string // Formats the GDateTime based on the time package layout specifier. (根据格式规范格式化时间)


YearsBetween(end *GDateTime) int // Calculates the full year difference between two dates, adjusting for incomplete year spans. (计算两个日期之间完整年份的差异，考虑不完整的年份差距)
MonthsBetween(end *GDateTime) int // Calculates the full month difference between two dates, adjusting for incomplete month spans. (计算两个日期之间完整月份的差异，考虑不完整的月份差距)
DaysBetween(end *GDateTime) int // Calculates the full day difference between two dates based on actual time difference. (根据实际时间差异计算两个日期之间的天数差异)
HoursBetween(end *GDateTime) int // Calculates the hour difference between two timestamps. (计算两个时间戳之间的小时差)
MinutesBetween(end *GDateTime) int // Calculates the minute difference between two timestamps. (计算两个时间戳之间的分钟差)
SecondsBetween(end *GDateTime) int // Calculates the second difference between two timestamps. (计算两个时间戳之间的秒差)

```

-----------
You can help to make the project better

