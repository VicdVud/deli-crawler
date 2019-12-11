package db

import (
	"database/sql"
	"fmt"
	"github.com/VicdVud/deli-crawler/internal/model"
	"strings"
)

type attendanceDB struct{}

var DefaultAttendance = attendanceDB{}

// Create 增加一条记录
func (a attendanceDB) Create(attendance *model.Attendance) error {
	// 先查询该记录存不存在
	year, month, day := attendance.Date.Date()
	date := fmt.Sprintf("%04d-%02d-%02d", year, month, day)
	rc, _ := a.FindOne(attendance.Name, date)
	if rc != nil {
		// 若存在，判断记录是否有改动
		if attendance.EqualTo(rc) {
			return nil
		}
		// return a.Update(attendance)
		return nil
	}

	// 否则新增
	strSql := "INSERT INTO attendance(" + a.fields() + ") VALUES(?,?,?,?,?,?,?,?,?,?,?)"
	_, err := masterDB.Exec(strSql, strings.TrimSpace(attendance.Name),
		strings.TrimSpace(attendance.Department), attendance.Date,
		strings.TrimSpace(attendance.Week), strings.TrimSpace(attendance.DateType),
		strings.TrimSpace(attendance.ClockIn), strings.TrimSpace(attendance.ClockOut),
		attendance.Duration, attendance.Late,
		attendance.LeaveEarly, attendance.Absent)
	if err != nil {
		return err
	}

	return nil
}

// FindOne 查询一条记录
// @param name 姓名
// @param date 日期，如'2019-01-01'
func (a attendanceDB) FindOne(name string, date string) (*model.Attendance, error) {
	attendance := &model.Attendance{}

	strSql := "SELECT " + a.fields() +
		" FROM attendance WHERE name=? and date=?"
	row := masterDB.QueryRow(strSql, name, date)
	err := a.scanRow(row, attendance)
	if err != nil {
		return nil, err
	}

	return attendance, nil
}

// Update 更新一条记录
// 只更新考勤相关数据
func (a attendanceDB) Update(attendance *model.Attendance) error {
	strSql := "UPDATE attendance SET date_type=?, clock_in=?,clock_out=?,duration=?,late=?,leave_early=?,absent=? where name=? and date=?"
	_, err := masterDB.Exec(strSql, strings.TrimSpace(attendance.DateType),
		strings.TrimSpace(attendance.ClockIn), strings.TrimSpace(attendance.ClockOut),
		attendance.Duration, attendance.Late,
		attendance.LeaveEarly, attendance.Absent,
		strings.TrimSpace(attendance.Name),
		attendance.Date)
	if err != nil {
		return err
	}
	return nil
}

// fields 获取数据库字段列表
func (a attendanceDB) fields() string {
	return "name,department,date,week,date_type,clock_in,clock_out,duration,late,leave_early,absent"
}

// scanRow 将查询结果转为结构体
func (a attendanceDB) scanRow(row *sql.Row, attendance *model.Attendance) error {
	return row.Scan(
		&attendance.Name,
		&attendance.Department,
		&attendance.Date,
		&attendance.Week,
		&attendance.DateType,
		&attendance.ClockIn,
		&attendance.ClockOut,
		&attendance.Duration,
		&attendance.Late,
		&attendance.LeaveEarly,
		&attendance.Absent,
	)
}
