package gdatetime

import (
	"fmt"
	"github.com/linsongze/go-date-time/datetime/timeconst"
	"github.com/linsongze/go-date-time/datetime/timeunit"
	"testing"
	"time"
)

// Test the Create function
func TestCreate(t *testing.T) {
	now := time.Now()
	gdt := Create(now)
	if gdt.t != now {
		t.Errorf("Create failed, expected %v, got %v", now, gdt.t)
	}
}

// Test the Now function
func TestNow(t *testing.T) {
	gdt := Now()
	if time.Since(gdt.t) > time.Second {
		t.Errorf("Now is not accurate, got %v", gdt.t)
	}
}

// Test Of method
func TestOf(t *testing.T) {
	gdt, err := Of(2023, 1, 15, 12, 0, 0, 0)
	if err != nil {
		t.Errorf("Of returned error: %v", err)
	}
	expected := time.Date(2023, 1, 15, 12, 0, 0, 0, time.Local)
	if gdt.t != expected {
		t.Errorf("Of failed, expected %v, got %v", expected, gdt.t)
	}
}

// Test Of2 method
func TestOf2(t *testing.T) {
	gdt, err := Of2(2023, 1, 15, 12, 0, 0)
	if err != nil {
		t.Errorf("Of2 returned error: %v", err)
	}
	expected := time.Date(2023, 1, 15, 12, 0, 0, 0, time.Local)
	if gdt.t != expected {
		t.Errorf("Of2 failed, expected %v, got %v", expected, gdt.t)
	}
}

// Test Of3 method
func TestOf3(t *testing.T) {
	gdt, err := Of3(2023, 1, 15, 12, 0)
	if err != nil {
		t.Errorf("Of3 returned error: %v", err)
	}
	expected := time.Date(2023, 1, 15, 12, 0, 0, 0, time.Local)
	if gdt.t != expected {
		t.Errorf("Of3 failed, expected %v, got %v", expected, gdt.t)
	}
}

// Test GetYear method
func TestGetYear(t *testing.T) {
	gdt := Now()
	expectedYear := time.Now().Year()
	if gdt.GetYear() != expectedYear {
		t.Errorf("GetYear failed, expected %v, got %v", expectedYear, gdt.GetYear())
	}
}

// Test GetMonth method
func TestGetMonth(t *testing.T) {
	gdt := Now()
	expectedMonth := int(time.Now().Month())
	if gdt.GetMonth() != expectedMonth {
		t.Errorf("GetMonth failed, expected %v, got %v", expectedMonth, gdt.GetMonth())
	}
}

// Test GetDayOfMonth method
func TestGetDayOfMonth(t *testing.T) {
	gdt := Now()
	expectedDay := time.Now().Day()
	if gdt.GetDayOfMonth() != expectedDay {
		t.Errorf("GetDayOfMonth failed, expected %v, got %v", expectedDay, gdt.GetDayOfMonth())
	}
}

// Test GetDayOfYear method
func TestGetDayOfYear(t *testing.T) {
	gdt := Create(time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC))
	if day := gdt.GetDayOfYear(); day != 1 {
		t.Errorf("GetDayOfYear failed, expected 1, got %d", day)
	}
}

// Test GetDayOfWeek method
func TestGetDayOfWeek(t *testing.T) {
	// 2023-01-01 is a Sunday
	gdt := Create(time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC))
	if day := gdt.GetDayOfWeek(); day != 0 {
		t.Errorf("GetDayOfWeek failed, expected 0 (Sunday), got %d", day)
	}
}

// Test GetHour method
func TestGetHour(t *testing.T) {
	gdt := Create(time.Date(2023, 1, 1, 15, 0, 0, 0, time.UTC))
	if hour := gdt.GetHour(); hour != 15 {
		t.Errorf("GetHour failed, expected 15, got %d", hour)
	}
}

// Test GetMinute method
func TestGetMinute(t *testing.T) {
	gdt := Create(time.Date(2023, 1, 1, 0, 30, 0, 0, time.UTC))
	if minute := gdt.GetMinute(); minute != 30 {
		t.Errorf("GetMinute failed, expected 30, got %d", minute)
	}
}

// Test GetSecond method
func TestGetSecond(t *testing.T) {
	gdt := Create(time.Date(2023, 1, 1, 0, 0, 45, 0, time.UTC))
	if second := gdt.GetSecond(); second != 45 {
		t.Errorf("GetSecond failed, expected 45, got %d", second)
	}
}

// Test GetNano method
func TestGetNano(t *testing.T) {
	nano := 123456789
	gdt := Create(time.Date(2023, 1, 1, 0, 0, 0, nano, time.UTC))
	if nanoGot := gdt.GetNano(); nanoGot != nano {
		t.Errorf("GetNano failed, expected %d, got %d", nano, nanoGot)
	}
}

// Test GSecondTimestamp method
func TestGSecondTimestamp(t *testing.T) {
	gdt := Create(time.Unix(1609459200, 0)) // 2021-01-01 00:00:00 UTC
	expectedTimestamp := int64(1609459200)
	if timestamp := gdt.GetSecondTimestamp(); timestamp != expectedTimestamp {
		t.Errorf("GSecondTimestamp failed, expected %d, got %d", expectedTimestamp, timestamp)
	}
}

// Test GMillSecondTimestamp method
func TestGMillSecondTimestamp(t *testing.T) {
	gdt := Create(time.Unix(0, 1609459200000000)) // 2021-01-01 00:00:00 UTC

	fmt.Println(gdt.GetMillSecondTimestamp())
}
func TestWithYear(t *testing.T) {
	initialTime := time.Date(2023, 3, 1, 0, 0, 0, 0, time.UTC)
	gdt := Create(initialTime)
	// Test valid year
	updatedGdt, err := gdt.WithYear(2025)
	if err != nil {
		t.Errorf("withYear returned error: %v", err)
	}
	if updatedGdt.GetYear() != 2025 {
		t.Errorf("withYear failed, expected %d, got %d", 2025, updatedGdt.GetYear())
	}

	// Test invalid year
	updatedGdt, err = gdt.WithYear(-1)
	if err == nil {
		t.Error("withYear should fail with negative year")
	}
}

func TestWithMonth(t *testing.T) {
	initialTime := time.Date(2023, 1, 31, 0, 0, 0, 0, time.UTC)
	gdt := Create(initialTime)
	// Test valid month
	updatedGdt, err := gdt.WithMonth(12)
	if err != nil {
		t.Errorf("withMonth returned error: %v", err)
	}
	if updatedGdt.GetMonth() != 12 {
		t.Errorf("withMonth failed, expected 12, got %d", updatedGdt.GetMonth())
	}

	// Test invalid month
	updatedGdt, err = gdt.WithMonth(13)
	if err == nil {
		t.Error("withMonth should fail with invalid month")
	}
}

func TestWithDayOfMonth(t *testing.T) {
	initialTime := time.Date(2023, 2, 1, 0, 0, 0, 0, time.UTC) // Non-leap year February
	gdt := Create(initialTime)
	// Test valid day of month
	updatedGdt, err := gdt.WithDayOfMonth(28)
	if err != nil {
		t.Errorf("withDayOfMonth returned error: %v", err)
	}
	if updatedGdt.GetDayOfMonth() != 28 {
		t.Errorf("withDayOfMonth failed, expected 28, got %d", updatedGdt.GetDayOfMonth())
	}

	// Test invalid day of month
	updatedGdt, err = gdt.WithDayOfMonth(30)
	if err == nil {
		t.Error("withDayOfMonth should fail with invalid day for February in a non-leap year")
	}
}

func TestWithDayOfYear(t *testing.T) {
	initialTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC) // Non-leap year
	gdt := Create(initialTime)
	// Test valid day of year
	updatedGdt, err := gdt.WithDayOfYear(365)
	if err != nil {
		t.Errorf("withDayOfYear returned error: %v", err)
	}
	if updatedGdt.GetDayOfYear() != 365 {
		t.Errorf("withDayOfYear failed, expected 365, got %d", updatedGdt.GetDayOfYear())
	}

	// Test invalid day of year (leap year check)
	updatedGdt, err = gdt.WithDayOfYear(366)
	if err == nil {
		t.Error("withDayOfYear should fail with day 366 in a non-leap year")
	}
}

func TestWithHour(t *testing.T) {
	initialTime := time.Date(2023, 10, 10, 12, 0, 0, 0, time.UTC)
	gdt := Create(initialTime)

	// Test valid hour
	updatedGdt, err := gdt.WithHour(23)
	if err != nil {
		t.Errorf("WithHour returned error: %v", err)
	}
	if updatedGdt.GetHour() != 23 {
		t.Errorf("WithHour failed, expected 23, got %d", updatedGdt.GetHour())
	}

	// Test invalid hour
	updatedGdt, err = gdt.WithHour(24)
	if err == nil {
		t.Error("WithHour should fail with invalid hour")
	}
}

func TestWithMinute(t *testing.T) {
	initialTime := time.Date(2023, 10, 10, 12, 30, 0, 0, time.UTC)
	gdt := Create(initialTime)

	// Test valid minute
	updatedGdt, err := gdt.WithMinute(59)
	if err != nil {
		t.Errorf("WithMinute returned error: %v", err)
	}
	if updatedGdt.GetMinute() != 59 {
		t.Errorf("WithMinute failed, expected 59, got %d", updatedGdt.GetMinute())
	}

	// Test invalid minute
	updatedGdt, err = gdt.WithMinute(60)
	if err == nil {
		t.Error("WithMinute should fail with invalid minute")
	}
}

func TestWithSecond(t *testing.T) {
	initialTime := time.Date(2023, 10, 10, 12, 30, 45, 0, time.UTC)
	gdt := Create(initialTime)

	// Test valid second
	updatedGdt, err := gdt.WithSecond(59)
	if err != nil {
		t.Errorf("WithSecond returned error: %v", err)
	}
	if updatedGdt.GetSecond() != 59 {
		t.Errorf("WithSecond failed, expected 59, got %d", updatedGdt.GetSecond())
	}

	// Test invalid second
	updatedGdt, err = gdt.WithSecond(60)
	if err == nil {
		t.Error("WithSecond should fail with invalid second")
	}
}

func TestWithNano(t *testing.T) {
	initialTime := time.Date(2023, 10, 10, 12, 30, 45, 500000000, time.UTC)
	gdt := Create(initialTime)

	// Test valid nanosecond
	updatedGdt, err := gdt.WithNano(999999999)
	if err != nil {
		t.Errorf("WithNano returned error: %v", err)
	}
	if updatedGdt.GetNano() != 999999999 {
		t.Errorf("WithNano failed, expected 999999999, got %d", updatedGdt.GetNano())
	}

	// Test invalid nanosecond
	updatedGdt, err = gdt.WithNano(1000000000)
	if err == nil {
		t.Error("WithNano should fail with invalid nanosecond")
	}
}

func TestPlusYear(t *testing.T) {
	initialTime := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	gdt := Create(initialTime)

	// Add 5 years
	expectedTime := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	updatedGdt := gdt.PlusYears(5)
	if updatedGdt.t != expectedTime {
		t.Errorf("plusYear failed, expected %v, got %v", expectedTime, updatedGdt.t)
	}

	// Subtract 2 years
	expectedTime = time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC)
	updatedGdt = gdt.PlusYears(-2)
	if updatedGdt.t != expectedTime {
		t.Errorf("plusYear with negative failed, expected %v, got %v", expectedTime, updatedGdt.t)
	}
}

func TestPlusMonth(t *testing.T) {
	initialTime := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	gdt := Create(initialTime)

	// Add 6 months
	expectedTime := time.Date(2020, 7, 1, 0, 0, 0, 0, time.UTC)
	updatedGdt := gdt.PlusMonths(6)
	if updatedGdt.t != expectedTime {
		t.Errorf("plusMonth failed, expected %v, got %v", expectedTime, updatedGdt.t)
	}

	// Subtract 1 month
	expectedTime = time.Date(2019, 12, 1, 0, 0, 0, 0, time.UTC)
	updatedGdt = gdt.PlusMonths(-1)
	if updatedGdt.t != expectedTime {
		t.Errorf("plusMonth with negative failed, expected %v, got %v", expectedTime, updatedGdt.t)
	}
}

func TestPlusWeek(t *testing.T) {
	initialTime := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	gdt := Create(initialTime)

	// Add 10 weeks
	expectedTime := time.Date(2020, 3, 11, 0, 0, 0, 0, time.UTC)
	updatedGdt := gdt.PlusWeeks(10)
	if updatedGdt.t != expectedTime {
		t.Errorf("plusWeek failed, expected %v, got %v", expectedTime, updatedGdt.t)
	}

	// Subtract 3 weeks
	expectedTime = time.Date(2019, 12, 11, 0, 0, 0, 0, time.UTC)
	updatedGdt = gdt.PlusWeeks(-3)
	if updatedGdt.t != expectedTime {
		t.Errorf("plusWeek with negative failed, expected %v, got %v", expectedTime, updatedGdt.t)
	}
}

func TestPlusDay(t *testing.T) {
	initialTime := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	gdt := Create(initialTime)

	// Add 30 days
	expectedTime := time.Date(2020, 1, 31, 0, 0, 0, 0, time.UTC)
	updatedGdt := gdt.PlusDays(30)
	if updatedGdt.t != expectedTime {
		t.Errorf("plusDay failed, expected %v, got %v", expectedTime, updatedGdt.t)
	}

	// Subtract 10 days
	expectedTime = time.Date(2019, 12, 22, 0, 0, 0, 0, time.UTC)
	updatedGdt = gdt.PlusDays(-10)
	if updatedGdt.t != expectedTime {
		t.Errorf("plusDay with negative failed, expected %v, got %v", expectedTime, updatedGdt.t)
	}
}

func TestPlusHours(t *testing.T) {
	initialTime := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	gdt := Create(initialTime)

	// Add 5 hours
	expectedTime := time.Date(2020, 1, 1, 5, 0, 0, 0, time.UTC)
	updatedGdt := gdt.PlusHours(5)
	if updatedGdt.t != expectedTime {
		t.Errorf("plusHours failed, expected %v, got %v", expectedTime, updatedGdt.t)
	}

	// Subtract 2 hours
	expectedTime = time.Date(2019, 12, 31, 22, 0, 0, 0, time.UTC)
	updatedGdt = gdt.PlusHours(-2)
	if updatedGdt.t != expectedTime {
		t.Errorf("plusHours with negative failed, expected %v, got %v", expectedTime, updatedGdt.t)
	}
}

func TestPlusMinutes(t *testing.T) {
	initialTime := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	gdt := Create(initialTime)

	// Add 30 minutes
	expectedTime := time.Date(2020, 1, 1, 0, 30, 0, 0, time.UTC)
	updatedGdt := gdt.PlusMinutes(30)
	if updatedGdt.t != expectedTime {
		t.Errorf("plusMinutes failed, expected %v, got %v", expectedTime, updatedGdt.t)
	}

	// Subtract 15 minutes
	expectedTime = time.Date(2019, 12, 31, 23, 45, 0, 0, time.UTC)
	updatedGdt = gdt.PlusMinutes(-15)
	if updatedGdt.t != expectedTime {
		t.Errorf("plusMinutes with negative failed, expected %v, got %v", expectedTime, updatedGdt.t)
	}
}

func TestPlusSeconds(t *testing.T) {
	initialTime := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	gdt := Create(initialTime)

	// Add 90 seconds
	expectedTime := time.Date(2020, 1, 1, 0, 1, 30, 0, time.UTC)
	updatedGdt := gdt.PlusSeconds(90)
	if updatedGdt.t != expectedTime {
		t.Errorf("plusSeconds failed, expected %v, got %v", expectedTime, updatedGdt.t)
	}

	// Subtract 90 seconds
	expectedTime = time.Date(2019, 12, 31, 23, 58, 30, 0, time.UTC)
	updatedGdt = gdt.PlusSeconds(-90)
	if updatedGdt.t != expectedTime {
		t.Errorf("plusSeconds with negative failed, expected %v, got %v", expectedTime, updatedGdt.t)
	}
}

func TestPlusNanos(t *testing.T) {
	initialTime := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	gdt := Create(initialTime)

	// Add 1 billion nanoseconds (1 second)
	expectedTime := time.Date(2020, 1, 1, 0, 0, 1, 0, time.UTC)
	updatedGdt := gdt.PlusNanos(1000000000)
	if updatedGdt.t != expectedTime {
		t.Errorf("plusNanos failed, expected %v, got %v", expectedTime, updatedGdt.t)
	}

	// Subtract 1 billion nanoseconds (1 second)
	expectedTime = time.Date(2019, 12, 31, 23, 59, 59, 0, time.UTC)
	updatedGdt = gdt.PlusNanos(-1000000000)
	if updatedGdt.t != expectedTime {
		t.Errorf("plusNanos with negative failed, expected %v, got %v", expectedTime, updatedGdt.t)
	}
}

func TestPlus(t *testing.T) {
	initialTime := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	gdt := Create(initialTime)

	// Test addition of nanoseconds
	expectedTime := initialTime.Add(time.Duration(100) * time.Nanosecond)
	if got := gdt.Plus(100, timeunit.NANOS); got.t != expectedTime {
		t.Errorf("Plus(NANOS) failed, expected %v, got %v", expectedTime, got.t)
	}

	// Test addition of microseconds
	microsToAdd := 100
	expectedTime = initialTime.Add(time.Duration(microsToAdd*1000) * time.Nanosecond)
	if got := gdt.Plus(microsToAdd, timeunit.MICROS); got.t != expectedTime {
		t.Errorf("Plus(MICROS) failed, expected %v, got %v", expectedTime, got.t)
	}

	// Test addition of milliseconds
	millisToAdd := 100
	expectedTime = initialTime.Add(time.Duration(millisToAdd*1000000) * time.Nanosecond)
	if got := gdt.Plus(millisToAdd, timeunit.MILLIS); got.t != expectedTime {
		t.Errorf("Plus(MILLIS) failed, expected %v, got %v", expectedTime, got.t)
	}

	// Test addition of seconds
	secondsToAdd := 60
	expectedTime = initialTime.Add(time.Duration(secondsToAdd) * time.Second)
	if got := gdt.Plus(secondsToAdd, timeunit.SECONDS); got.t != expectedTime {
		t.Errorf("Plus(SECONDS) failed, expected %v, got %v", expectedTime, got.t)
	}

	// Test addition of minutes
	minutesToAdd := 60
	expectedTime = initialTime.Add(time.Duration(minutesToAdd) * time.Minute)
	if got := gdt.Plus(minutesToAdd, timeunit.MINUTES); got.t != expectedTime {
		t.Errorf("Plus(MINUTES) failed, expected %v, got %v", expectedTime, got.t)
	}

	// Test addition of hours
	hoursToAdd := 24
	expectedTime = initialTime.Add(time.Duration(hoursToAdd) * time.Hour)
	if got := gdt.Plus(hoursToAdd, timeunit.HOURS); got.t != expectedTime {
		t.Errorf("Plus(HOURS) failed, expected %v, got %v", expectedTime, got.t)
	}

	// Test addition of half-days
	halfDaysToAdd := 2
	expectedTime = initialTime.AddDate(0, 0, halfDaysToAdd/2).Add(time.Duration(halfDaysToAdd%2*12) * time.Hour)
	if got := gdt.Plus(halfDaysToAdd, timeunit.HALF_DAYS); got.t != expectedTime {
		t.Errorf("Plus(HALF_DAYS) failed, expected %v, got %v", expectedTime, got.t)
	}
}

func TestMinusYears(t *testing.T) {
	initialTime := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	gdt := Create(initialTime)

	// Subtract 5 years
	expectedTime := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	updatedGdt := gdt.MinusYears(5)
	if updatedGdt.t != expectedTime {
		t.Errorf("minusYears failed, expected %v, got %v", expectedTime, updatedGdt.t)
	}
}

func TestMinusMonths(t *testing.T) {
	initialTime := time.Date(2020, 6, 1, 0, 0, 0, 0, time.UTC)
	gdt := Create(initialTime)

	// Subtract 6 months
	expectedTime := time.Date(2019, 12, 1, 0, 0, 0, 0, time.UTC)
	updatedGdt := gdt.MinusMonths(6)
	if updatedGdt.t != expectedTime {
		t.Errorf("minusMonths failed, expected %v, got %v", expectedTime, updatedGdt.t)
	}
}

func TestMinusWeeks(t *testing.T) {
	initialTime := time.Date(2020, 1, 22, 0, 0, 0, 0, time.UTC)
	gdt := Create(initialTime)

	// Subtract 3 weeks
	expectedTime := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	updatedGdt := gdt.MinusWeeks(3)
	if updatedGdt.t != expectedTime {
		t.Errorf("minusWeeks failed, expected %v, got %v", expectedTime, updatedGdt.t)
	}
}

func TestMinusDays(t *testing.T) {
	initialTime := time.Date(2020, 1, 31, 0, 0, 0, 0, time.UTC)
	gdt := Create(initialTime)

	// Subtract 30 days
	expectedTime := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	updatedGdt := gdt.MinusDays(30)
	if updatedGdt.t != expectedTime {
		t.Errorf("minusDays failed, expected %v, got %v", expectedTime, updatedGdt.t)
	}
}

func TestMinusHours(t *testing.T) {
	initialTime := time.Date(2020, 1, 1, 23, 0, 0, 0, time.UTC)
	gdt := Create(initialTime)

	// Subtract 22 hours
	expectedTime := time.Date(2020, 1, 1, 1, 0, 0, 0, time.UTC)
	updatedGdt := gdt.MinusHours(22)
	if updatedGdt.t != expectedTime {
		t.Errorf("minusHours failed, expected %v, got %v", expectedTime, updatedGdt.t)
	}
}

func TestMinusMinutes(t *testing.T) {
	initialTime := time.Date(2020, 1, 1, 0, 59, 0, 0, time.UTC)
	gdt := Create(initialTime)

	// Subtract 58 minutes
	expectedTime := time.Date(2020, 1, 1, 0, 1, 0, 0, time.UTC)
	updatedGdt := gdt.MinusMinutes(58)
	if updatedGdt.t != expectedTime {
		t.Errorf("minusMinutes failed, expected %v, got %v", expectedTime, updatedGdt.t)
	}
}

func TestMinusSeconds(t *testing.T) {
	initialTime := time.Date(2020, 1, 1, 0, 0, 59, 0, time.UTC)
	gdt := Create(initialTime)

	// Subtract 58 seconds
	expectedTime := time.Date(2020, 1, 1, 0, 0, 1, 0, time.UTC)
	updatedGdt := gdt.MinusSeconds(58)
	if updatedGdt.t != expectedTime {
		t.Errorf("minusSeconds failed, expected %v, got %v", expectedTime, updatedGdt.t)
	}
}

func TestMinusNanos(t *testing.T) {
	initialTime := time.Date(2020, 1, 1, 0, 0, 0, 1000, time.UTC)
	gdt := Create(initialTime)

	// Subtract 500 nanoseconds
	expectedTime := time.Date(2020, 1, 1, 0, 0, 0, 500, time.UTC)
	updatedGdt := gdt.MinusNanos(500)
	if updatedGdt.t != expectedTime {
		t.Errorf("minusNanos failed, expected %v, got %v", expectedTime, updatedGdt.t)
	}
}

func TestStartOfDay(t *testing.T) {
	initialTime := time.Date(2023, 6, 15, 14, 30, 59, 123456789, time.UTC)
	gdt := Create(initialTime)
	expectedTime := time.Date(2023, 6, 15, 0, 0, 0, 0, time.UTC)

	startOfDay := gdt.StartOfDay()
	if !startOfDay.t.Equal(expectedTime) {
		t.Errorf("StartOfDay failed, expected %v, got %v", expectedTime, startOfDay.t)
	}
}

func TestEndOfDay(t *testing.T) {
	initialTime := time.Date(2023, 6, 15, 14, 30, 59, 123456789, time.UTC)
	gdt := Create(initialTime)
	expectedTime := time.Date(2023, 6, 15, 23, 59, 59, timeconst.MAX_NANO, time.UTC)

	endOfDay := gdt.EndOfDay()
	if !endOfDay.t.Equal(expectedTime) {
		t.Errorf("EndOfDay failed, expected %v, got %v", expectedTime, endOfDay.t)
	}
}

func TestConvertToZone(t *testing.T) {
	initialTime := time.Date(2023, 6, 15, 12, 0, 0, 0, time.UTC)
	gdt := Create(initialTime)
	newYorkLocation, _ := time.LoadLocation("America/New_York") // UTC-5

	expectedTime := initialTime.In(newYorkLocation)
	switchedTime := gdt.ConvertToZone(*newYorkLocation)
	if !switchedTime.t.Equal(expectedTime) {
		t.Errorf("ConvertToZone failed, expected %v, got %v", expectedTime, switchedTime.t)
	}
}

func TestResetZoneToDefault(t *testing.T) {
	initialTime := time.Date(2023, 6, 15, 12, 0, 0, 0, time.UTC)
	gdt := Create(initialTime)
	newYorkLocation, _ := time.LoadLocation("America/New_York")
	gdt = gdt.ConvertToZone(*newYorkLocation) // Change time zone to New York

	expectedTime := initialTime.In(time.Local)
	resetTime := gdt.ResetZoneToDefault()
	if !resetTime.t.Equal(expectedTime) {
		t.Errorf("ResetZoneToDefault failed, expected %v, got %v", expectedTime, resetTime.t)
	}
}
func TestEqualDate(t *testing.T) {
	time1 := time.Date(2023, 6, 15, 14, 30, 0, 0, time.UTC)
	time2 := time.Date(2023, 6, 15, 10, 0, 0, 0, time.UTC)
	time3 := time.Date(2023, 6, 16, 14, 30, 0, 0, time.UTC)
	time4 := time.Date(2022, 6, 15, 14, 30, 0, 0, time.UTC)

	gdt1 := Create(time1)
	gdt2 := Create(time2)
	gdt3 := Create(time3)
	gdt4 := Create(time4)

	cases := []struct {
		name     string
		gdtA     *GDateTime
		gdtB     *GDateTime
		expected bool
	}{
		{"Same Day Different Time", gdt1, gdt2, true},
		{"Different Day", gdt1, gdt3, false},
		{"Different Year", gdt1, gdt4, false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if result := tc.gdtA.EqualDate(tc.gdtB); result != tc.expected {
				t.Errorf("%s: expected %v, got %v", tc.name, tc.expected, result)
			}
		})
	}
}
func TestCompareTo(t *testing.T) {
	time1 := time.Date(2023, 6, 15, 14, 30, 0, 0, time.UTC)
	time2 := time.Date(2023, 6, 15, 15, 0, 0, 0, time.UTC)
	time3 := time.Date(2023, 6, 15, 14, 30, 0, 0, time.UTC)

	gdt1 := Create(time1)
	gdt2 := Create(time2)
	gdt3 := Create(time3)

	cases := []struct {
		name     string
		gdtA     *GDateTime
		gdtB     *GDateTime
		expected int
	}{
		{"Equal Times", gdt1, gdt3, 0},
		{"First Earlier Than Second", gdt1, gdt2, -1},
		{"First Later Than Second", gdt2, gdt1, 1},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.gdtA.CompareTo(tc.gdtB)
			if result != tc.expected {
				t.Errorf("%s failed: expected %d, got %d", tc.name, tc.expected, result)
			}
		})
	}
}
func TestIsBeforeAndIsAfter(t *testing.T) {
	time1 := time.Date(2023, 6, 15, 14, 30, 0, 0, time.UTC)
	time2 := time.Date(2023, 6, 15, 15, 30, 0, 0, time.UTC)

	gdt1 := Create(time1)
	gdt2 := Create(time2)

	// gdt1 is before gdt2
	if !gdt1.IsBefore(gdt2) {
		t.Errorf("IsBefore failed: expected true for gdt1(%v) is before gdt2(%v)", gdt1.t, gdt2.t)
	}
	if gdt1.IsAfter(gdt2) {
		t.Errorf("IsAfter failed: expected false for gdt1(%v) is after gdt2(%v)", gdt1.t, gdt2.t)
	}

	// gdt2 is after gdt1
	if !gdt2.IsAfter(gdt1) {
		t.Errorf("IsAfter failed: expected true for gdt2(%v) is after gdt1(%v)", gdt2.t, gdt1.t)
	}
	if gdt2.IsBefore(gdt1) {
		t.Errorf("IsBefore failed: expected false for gdt2(%v) is before gdt1(%v)", gdt2.t, gdt1.t)
	}

	// Same time should be neither before nor after
	gdt3 := Create(time1)
	if gdt1.IsBefore(gdt3) || gdt1.IsAfter(gdt3) {
		t.Errorf("IsBefore/IsAfter failed: expected false for the same time comparison")
	}
}

func TestCompareDate(t *testing.T) {
	time1 := time.Date(2023, 6, 15, 14, 30, 0, 0, time.UTC)
	time2 := time.Date(2023, 6, 16, 14, 30, 0, 0, time.UTC)
	time3 := time.Date(2022, 6, 15, 14, 30, 0, 0, time.UTC)

	gdt1 := Create(time1)
	gdt2 := Create(time2)
	gdt3 := Create(time3)

	if gdt1.CompareDate(gdt1) != 0 {
		t.Errorf("CompareDate same date should return 0")
	}
	if gdt1.CompareDate(gdt2) != -1 {
		t.Errorf("CompareDate earlier year should return -1")
	}
	if gdt2.CompareDate(gdt1) != 1 {
		t.Errorf("CompareDate later year should return 1")
	}
	if gdt1.CompareDate(gdt3) != 1 {
		t.Errorf("CompareDate later year against an earlier year should return 1")
	}
}

func TestCompareTime(t *testing.T) {
	time1 := time.Date(2023, 6, 15, 14, 30, 0, 0, time.UTC)
	time2 := time.Date(2023, 6, 15, 15, 30, 0, 0, time.UTC)
	time3 := time.Date(2023, 6, 15, 14, 29, 59, 0, time.UTC)

	gdt1 := Create(time1)
	gdt2 := Create(time2)
	gdt3 := Create(time3)

	if gdt1.CompareTime(gdt1) != 0 {
		t.Errorf("CompareTime same time should return 0")
	}
	if gdt1.CompareTime(gdt2) != -1 {
		t.Errorf("CompareTime earlier time should return -1")
	}
	if gdt2.CompareTime(gdt1) != 1 {
		t.Errorf("CompareTime later time should return 1")
	}
	if gdt1.CompareTime(gdt3) != 1 {
		t.Errorf("CompareTime later minute against an earlier minute should return 1")
	}
}
func TestFormat(t *testing.T) {
	initialTime := time.Date(2023, 6, 15, 14, 30, 59, 0, time.UTC)
	gdt := Create(initialTime)

	// Example layout (RFC3339)
	expected := "2023-06-15T14:30:59Z"
	result := gdt.Format(time.RFC3339)
	if result != expected {
		t.Errorf("Format failed, expected %v, got %v", expected, result)
	}

	// Custom layout
	expected = "15-06-2023 14:30"
	result = gdt.Format("02-01-2006 15:04")
	if result != expected {
		t.Errorf("Custom Format failed, expected %v, got %v", expected, result)
	}

	// Another custom layout to test day and month names
	expected = "Thursday, Jun 15, 2023"
	result = gdt.Format("Monday, Jan 02, 2006")
	if result != expected {
		t.Errorf("Full date Format failed, expected %v, got %v", expected, result)
	}
}

func TestMonthAndWeekBoundaries(t *testing.T) {
	// Assume we're using a known date for testing: 15th June 2023
	testDate := time.Date(2023, time.June, 15, 12, 0, 0, 0, time.UTC)
	gdt := Create(testDate)

	// Start of Month
	startOfMonth := gdt.StartOfMonth()
	expectedStartOfMonth := time.Date(2023, time.June, 1, 0, 0, 0, 0, time.UTC)
	if startOfMonth.t != expectedStartOfMonth {
		t.Errorf("StartOfMonth failed, expected %v, got %v", expectedStartOfMonth, startOfMonth.t)
	}

	// End of Month
	endOfMonth := gdt.EndOfMonth()
	expectedEndOfMonth := time.Date(2023, time.June, 30, 23, 59, 59, 999999999, time.UTC)
	if endOfMonth.t != expectedEndOfMonth {
		t.Errorf("EndOfMonth failed, expected %v, got %v", expectedEndOfMonth, endOfMonth.t)
	}

	// Start of Week assuming week starts on Sunday
	startOfWeek := gdt.StartOfWeek()
	expectedStartOfWeek := time.Date(2023, time.June, 11, 0, 0, 0, 0, time.UTC) // June 11th is Sunday
	if startOfWeek.t != expectedStartOfWeek {
		t.Errorf("StartOfWeek failed, expected %v, got %v", expectedStartOfWeek, startOfWeek.t)
	}

	// End of Week assuming week ends on Saturday
	endOfWeek := gdt.EndOfWeek()
	expectedEndOfWeek := time.Date(2023, time.June, 17, 23, 59, 59, 999999999, time.UTC) // June 17th is Saturday
	if endOfWeek.t != expectedEndOfWeek {
		t.Errorf("EndOfWeek failed, expected %v, got %v", expectedEndOfWeek, endOfWeek.t)
	}
}

func TestWeekBoundariesFromMonday(t *testing.T) {
	// 定义一个已知的日期用于测试：2023年6月15日，这是一个周四。
	// Define a known date for testing: June 15, 2023, which is a Thursday.
	testDate := time.Date(2023, time.June, 15, 12, 0, 0, 0, time.UTC)
	gdt := Create(testDate)

	// 测试从周一开始的周的开始
	// Test the start of the week starting from Monday
	startOfWeek := gdt.StartOfWeekFromMonday()
	expectedStartOfWeek := time.Date(2023, time.June, 12, 0, 0, 0, 0, time.UTC) // June 12, 2023, is Monday
	if startOfWeek.t != expectedStartOfWeek {
		t.Errorf("StartOfWeekFromMonday failed, expected %v, got %v", expectedStartOfWeek, startOfWeek.t)
	}

	// 测试从周一开始的周的结束
	// Test the end of the week ending on Sunday
	endOfWeek := gdt.EndOfWeekFromMonday()
	expectedEndOfWeek := time.Date(2023, time.June, 18, 23, 59, 59, 999999999, time.UTC) // June 18, 2023, is Sunday
	if endOfWeek.t != expectedEndOfWeek {
		t.Errorf("EndOfWeekFromMonday failed, expected %v, got %v", expectedEndOfWeek, endOfWeek.t)
	}
}
func TestParse(t *testing.T) {
	dateStr := "2023-06-02T15:04:05Z"
	layout := time.RFC3339
	gdt, err := Parse(dateStr, layout)
	if err != nil {
		t.Errorf("Parse failed: %v", err)
	}
	expectedTime := "2023-06-02 15:04:05 +0000 UTC"
	if gdt.ToTime().String() != expectedTime {
		t.Errorf("Expected %v, got %v", expectedTime, gdt.ToTime().String())
	}
}

func TestFormUnixTimestamp(t *testing.T) {
	timestamp := int64(1654172405) // Equivalent to 2022-06-02 20:20:05.5 +0800 CST
	nano := int64(500000000)       // 0.5 second
	gdt, err := FormUnixTimestamp(timestamp, nano)
	location, _ := time.LoadLocation("Asia/Shanghai")
	gdt = gdt.ConvertToZone(*location)
	if err != nil {
		t.Errorf("FormUnixTimestamp failed: %v", err)
	}
	expectedTime := "2022-06-02 20:20:05.5 +0800 CST"
	if gdt.t.String() != expectedTime {
		t.Errorf("Expected %v, got %v", expectedTime, gdt.ToTime().String())
	}
}

func TestFormMillisTimestamp(t *testing.T) {
	millisTimestamp := int64(1654172405000) // Equivalent to 2022-06-02 20:20:05 +0800 CST
	gdt, err := FormMillisTimestamp(millisTimestamp)
	location, _ := time.LoadLocation("Asia/Shanghai")
	gdt = gdt.ConvertToZone(*location)
	if err != nil {
		t.Errorf("FormMillisTimestamp failed: %v", err)
	}
	expectedTime := "2022-06-02 20:20:05 +0800 CST"
	if gdt.ToTime().String() != expectedTime {
		t.Errorf("Expected %v, got %v", expectedTime, gdt.ToTime().String())
	}
}

// Test functions for year, month, day, hour, minute, and second differences
func TestTimeDifferences(t *testing.T) {
	startDate := Create(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC))
	endDate := Create(time.Date(2022, 3, 2, 23, 59, 59, 0, time.UTC))

	// Test YearsBetween
	if diff := startDate.YearsBetween(endDate); diff != 2 {
		t.Errorf("Expected 2 years, got %d", diff)
	}

	// Test MonthsBetween
	if diff := startDate.MonthsBetween(endDate); diff != 26 {
		t.Errorf("Expected 26 months, got %d", diff)
	}

	// Test DaysBetween
	if diff := startDate.DaysBetween(endDate); diff != 791 {
		t.Errorf("Expected 792 days, got %d", diff)
	}

	// Test HoursBetween
	if diff := startDate.HoursBetween(endDate); diff != 19007 {
		t.Errorf("Expected 19047 hours, got %d", diff)
	}

	// Test MinutesBetween
	if diff := startDate.MinutesBetween(endDate); diff != 1140479 {
		t.Errorf("Expected 1142839 minutes, got %d", diff)
	}

	// Test SecondsBetween
	if diff := startDate.SecondsBetween(endDate); diff != 68428799 {
		t.Errorf("Expected 68570399 seconds, got %d", diff)
	}
}
