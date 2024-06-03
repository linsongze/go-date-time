package gdatetime

import (
	"testing"
	"time"
)

// TestStrftime 测试Strftime函数
func TestStrftime(t *testing.T) {
	// 设定一个具体的时间点：2024年6月3日10点15分30秒
	location, _ := time.LoadLocation("UTC")
	testTime := time.Date(2024, 6, 3, 10, 15, 30, 0, location)
	gdt := Create(testTime)

	// 测试用例列表
	cases := []struct {
		format   string
		expected string
	}{
		{"%Y-%m-%d", "2024-06-03"},             // 测试年月日
		{"%H:%M:%S", "10:15:30"},               // 测试时分秒
		{"%B %d, %Y", "June 03, 2024"},         // 测试月份全称，日期，年份
		{"It is %I:%M %p.", "It is 10:15 AM."}, // 测试12小时制和AM/PM
		{"Weekday: %A", "Weekday: Monday"},     // 测试星期几的全称
		{"%c", "Mon Jun 3 10:15:30 2024"},      // 测试默认的C语言日期时间格式
		{"Year day: %j", "Year day: 155"},      // 测试年中的第几天
	}

	// 循环测试每个用例
	for _, c := range cases {
		got := gdt.Strftime(c.format)
		if got != c.expected {
			t.Errorf("Strftime(&testTime, %q) == %q, want %q", c.format, got, c.expected)
		}
	}
}
