package alidrive

import (
	"alist/config"
	"fmt"
)

func InitAliDrive() bool {
	// 初始化阿里云
	if config.Conf.AliDrive.RefreshToken == "" {
		// 重新获取
		fmt.Println("token无效，请重新设置")
		return false
	} else {
		if !RefreshToken(config.Conf.AliDrive.RefreshToken) {
			fmt.Println("token更新失败")
		} // 更新token获取accessToken---ok
		fmt.Println("token更新成功")
		config.Authorization = config.Bearer + config.Conf.AliDrive.AccessToken
	}
	// 更新状态----
	// fmt.Println("请求头", config.Authorization)
	// 初始化用户信息
	GetUserInfo() // 更新User信息--ok
	// fmt.Println(res.String())
	return true
}
