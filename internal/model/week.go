package model

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var DATE_LAYOUT = "2006-01-02"

type Date struct {
	Year  int
	Month int
	Day   int
}

// 记录1-12月份中每个月的天数
var monthDay map[int]int

// 闰年月份天数
var monthDayOfLeapYear map[int]int

func init() {
	monthDay = map[int]int{
		1:  31,
		2:  28,
		3:  31,
		4:  30,
		5:  31,
		6:  30,
		7:  31,
		8:  31,
		9:  30,
		10: 31,
		11: 30,
		12: 31,
	}

	monthDayOfLeapYear = map[int]int{
		1:  31,
		2:  29,
		3:  31,
		4:  30,
		5:  31,
		6:  30,
		7:  31,
		8:  31,
		9:  30,
		10: 31,
		11: 30,
		12: 31,
	}
}

func (d Date) NextDay() Date {
	nd := Date{
		Year:  d.Year,
		Month: d.Month,
		Day:   d.Day + 1,
	}

	checkDate(&nd)
	return nd
}

// FromString 从字符串解析日期
// 日期格式为: "2019-1-1"
func (d *Date) FromString(dateStr string) error {
	dateSlices := strings.Split(dateStr, "-")
	if len(dateSlices) != 3 {
		return errors.New("error date format")
	}

	d.Year, _ = strconv.Atoi(dateSlices[0])
	d.Month, _ = strconv.Atoi(dateSlices[1])
	d.Day, _ = strconv.Atoi(dateSlices[2])
	return nil
}

func (d *Date) ToString() string {
	return fmt.Sprintf("%d-%d-%d", d.Year, d.Month, d.Day)
}

// 判断两个日期是否相等
func (d Date) EqualTo(e Date) bool {
	return d.Year == e.Year && d.Month == e.Month && d.Day == e.Day
}

// 检查日期是否合规
func checkDate(d *Date) {
	var maxDay int
	if isLeapYear(d.Year) {
		maxDay = monthDayOfLeapYear[d.Month]
	} else {
		maxDay = monthDay[d.Month]
	}

	if d.Day > maxDay {
		// 月份+1
		d.Day -= maxDay
		d.Month++
	}

	if d.Month > 12 {
		// 年份+1
		d.Month -= 12
		d.Year++
	}
}

// isLeapYear 检查是否闰年
func isLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || (year%400 == 0)
}

func (d Date) GreaterEqual(other Date) bool {
	if d.Year > other.Year {
		return true
	} else if d.Year < other.Year {
		return false
	}

	if d.Month > other.Month {
		return true
	} else if d.Month < other.Month {
		return false
	}

	if d.Day >= other.Day {
		return true
	}
	return false
}
