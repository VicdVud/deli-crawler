package global

import (
	"github.com/VicdVud/deli-crawler/internal/model"
	"github.com/spf13/viper"
	"log"
	"strings"
	"time"
)

type deliInfo struct {
	Mobile         string     // 手机号
	Password       string     // 密码（加密后）
	SaveDir        string     // 下载的Excel保存文件夹路径
	FromDate       model.Date // 起始日期
	FromDateString string     // 起始日期字符串
	ToDate         model.Date // 终止日期
	ToDateString   string     // 终止日期字符串

	Cookie struct {
		GrUserId       string
		Sensorsdata    string
		GrwngUid       string
		GrSessionId_8  string
		CGrSessionId   string
		CGrSessionId53 string
		GrSessionId_9  string
		GrSessionId_95 string
	} // Cookie相关参数
}

const (
	TODAY = "today"
)

var Deli = &deliInfo{}

func InitDeliInfo() {
	Deli.Mobile = viper.GetString("deli.mobile")
	Deli.Password = viper.GetString("deli.password")
	Deli.SaveDir = viper.GetString("excel.dir")
	Deli.FromDateString = viper.GetString("date.from")
	Deli.ToDateString = viper.GetString("date.to")

	// 转换全小写
	Deli.FromDateString = strings.ToLower(Deli.FromDateString)
	Deli.ToDateString = strings.ToLower(Deli.ToDateString)

	// 获取今天日期
	dayNow := time.Now()

	if Deli.FromDateString == TODAY {
		Deli.FromDate = model.Date{
			Year:  dayNow.Year(),
			Month: int(dayNow.Month()),
			Day:   dayNow.Day(),
		}
	} else {
		err := Deli.FromDate.FromString(Deli.FromDateString)
		if err != nil {
			log.Fatal(err)
		}
	}

	if Deli.ToDateString == TODAY {
		Deli.ToDate = model.Date{
			Year:  dayNow.Year(),
			Month: int(dayNow.Month()),
			Day:   dayNow.Day(),
		}
	} else {
		err := Deli.ToDate.FromString(Deli.ToDateString)
		if err != nil {
			log.Fatal(err)
		}
	}

	Deli.Cookie.GrUserId = viper.GetString("cookie.gr_user_id")
	Deli.Cookie.Sensorsdata = viper.GetString("cookie.sensorsdata")
	Deli.Cookie.GrwngUid = viper.GetString("cookie.grwng_uid")
	Deli.Cookie.GrSessionId_8 = viper.GetString("cookie.gr_session_id_8")
	Deli.Cookie.CGrSessionId = viper.GetString("cookie.c_gr_session_id")
	Deli.Cookie.CGrSessionId53 = viper.GetString("cookie.c_gr_session_id_53")
	Deli.Cookie.GrSessionId_9 = viper.GetString("cookie.gr_session_id_9")
	Deli.Cookie.GrSessionId_95 = viper.GetString("cookie.gr_session_id_95")
}

// 是否定时获取
func CrawlRegularly() bool {
	return Deli.FromDateString == TODAY
}
