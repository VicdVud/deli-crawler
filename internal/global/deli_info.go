package global

import "github.com/spf13/viper"

type deliInfo struct {
	Mobile   string // 手机号
	Password string // 密码（加密后）
	SaveDir  string // 下载的Excel保存文件夹路径

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

var Deli = &deliInfo{}

func InitDeliInfo() {
	Deli.Mobile = viper.GetString("deli.mobile")
	Deli.Password = viper.GetString("deli.password")
	Deli.SaveDir = viper.GetString("excel.dir")

	Deli.Cookie.GrUserId = viper.GetString("cookie.gr_user_id")
	Deli.Cookie.Sensorsdata = viper.GetString("cookie.sensorsdata")
	Deli.Cookie.GrwngUid = viper.GetString("cookie.grwng_uid")
	Deli.Cookie.GrSessionId_8 = viper.GetString("cookie.gr_session_id_8")
	Deli.Cookie.CGrSessionId = viper.GetString("cookie.c_gr_session_id")
	Deli.Cookie.CGrSessionId53 = viper.GetString("cookie.c_gr_session_id_53")
	Deli.Cookie.GrSessionId_9 = viper.GetString("cookie.gr_session_id_9")
	Deli.Cookie.GrSessionId_95 = viper.GetString("cookie.gr_session_id_95")
}
