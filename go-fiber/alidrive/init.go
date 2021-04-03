package alidrive

import (
	"alist/config"
	"fmt"
)

func InitAliDrive() bool {
	// 初始化阿里云
	if config.Conf.AliDrive.RefreshToken == "" {
		// 重新获取
		panic("token无效，请重新设置env.yml")
	} else {
		if !RefreshToken(config.Conf.AliDrive.RefreshToken) {
			fmt.Println("token更新失败")
			return false
		}
		fmt.Println("token更新成功")
	}
	// 更新状态----
	// fmt.Println("请求头", config.Authorization)
	// 初始化用户信息
	GetUserInfo() // 更新User信息--ok
	// fmt.Println(res.String())
	return true
}
