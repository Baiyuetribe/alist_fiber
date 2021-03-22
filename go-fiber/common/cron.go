package common

import (
	"alist/alidrive"
	"alist/config"
	"fmt"

	"github.com/robfig/cron/v3" //定时任务
)

var Cron *cron.Cron // 未定义前，无法直接使用

func refreshToken() {
	alidrive.RefreshToken(config.Conf.AliDrive.RefreshToken) // 更新token获取accessToken---ok
	config.Authorization = config.Bearer + config.Conf.AliDrive.AccessToken
}

func InitCron() {
	Cron = cron.New()
	// _, err := Cron.AddFunc("@every 10s", refreshToken)
	_, err := Cron.AddFunc("@every 2h", refreshToken)
	if err != nil {
		fmt.Println("添加定时任务失败", err)
	}
	Cron.Start()
}
