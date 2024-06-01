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

-----------
You can help to make the project better

