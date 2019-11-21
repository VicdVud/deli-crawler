package global

import "github.com/spf13/viper"

type deliInfo struct {
	Mobile   string // 手机号
	Password string // 密码（加密后）
}

var Deli = &deliInfo{}

func InitDeliInfo() {
	Deli.Mobile = viper.GetString("deli.mobile")
	Deli.Password = viper.GetString("deli.password")
}
