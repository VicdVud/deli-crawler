package model

import "github.com/VicdVud/deli-crawler/internal/utils"

// 考勤信息
type Attendance struct {
	Name       string  // 姓名
	Department string  // 部门
	Date       string  // 日期
	Week       string  // 星期
	DateType   string  // 日期类型
	ClockIn    string  // 签到时间
	ClockOut   string  // 签退时间
	Duration   float64 // 工作时长
	Late       int     // 迟到时间（分钟）
	LeaveEarly int     // 早退时间（分钟）
	Absent     int     // 旷工时间（小时）
}

func (a *Attendance) EqualTo(e *Attendance) bool {
	return a.Name == e.Name && a.Department == e.Department &&
		a.Week == e.Week && a.DateType == e.DateType &&
		a.ClockIn == e.ClockIn && a.ClockOut == e.ClockOut &&
		utils.Approximately(a.Duration, e.Duration) &&
		a.Late == e.Late && a.LeaveEarly == e.LeaveEarly && a.Absent == e.Absent
}
