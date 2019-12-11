package xlsx

import (
	"errors"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/VicdVud/deli-crawler/internal/db"
	"github.com/VicdVud/deli-crawler/internal/logger"
	"github.com/VicdVud/deli-crawler/internal/model"
	"strconv"
	"strings"
)

// ReadAndSave 读取考勤文件，并上传至数据库
// @param path 考勤文件路径
func ReadAndSave(path string) error {
	xlsx, err := excelize.OpenFile(path)
	if err != nil {
		logger.Info(err)
	}

	// 获取第一分页名称
	sheetName := xlsx.GetSheetName(1)

	// 分行读取
	rows := xlsx.GetRows(sheetName)
	for i, row := range rows {
		if i < 2 {
			// 从第三行起读
			continue
		}

		if len(row) < 24 {
			return errors.New("wrong format in xlsx")
		}

		var date string

		// 每一行存储一条记录
		attendance := &model.Attendance{}
		attendance.Name = row[2]
		attendance.Department = row[3]
		date = row[5]
		attendance.Week = row[6]
		attendance.DateType = row[7]
		attendance.ClockIn = row[11]
		attendance.ClockOut = row[12]
		attendance.Duration, _ = strconv.ParseFloat(row[13], 32)
		attendance.Late, _ = strconv.Atoi(row[15])
		attendance.LeaveEarly, _ = strconv.Atoi(row[16])
		attendance.Absent, _ = strconv.Atoi(row[17])

		// 解析日期
		days := strings.Split(date, "-")
		if len(days) < 3 {
			continue
		}
		attendance.Year, _ = strconv.Atoi(days[0])
		attendance.Month, _ = strconv.Atoi(days[1])
		attendance.Day, _ = strconv.Atoi(days[2])

		// 保存至数据库
		err := db.DefaultAttendance.Create(attendance)
		if err != nil {
			return err
		}
	}

	return nil
}
