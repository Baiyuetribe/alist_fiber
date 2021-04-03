package common

import (
	"alist/alidrive"
	"alist/config"
	"alist/service"
	"alist/tools"
	"flag"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// 初始化参数---自动执行
func init() {
	flag.StringVar(&config.Conf.AliDrive.ApiUrl, "api", "https://api.aliyundrive.com/v2", "config api") // 定义配置文件名称
}

func Start() bool {
	printLogo()
	// 检查配置文件
	if !ReadConf("env.yml") { // >>config.yml---此操作同时更新基础配置文件里的内容到config配置里
		return false
	}
	// 初始化阿里云
	if !alidrive.InitAliDrive() {
		return false
	}
	// config.Authorization = config.Bearer + config.Conf.AliDrive.AccessToken
	// 初始化数据库--首次就完成
	config.InitDB()
	if err := service.Clear(); err != nil {
		fmt.Println("清空历史记录失败")
	}
	// 构建目录树
	if err := service.BuildTree(); err != nil {
		fmt.Println("构建目录树失败")
		return false
	}
	fmt.Println("初始化完成")
	// 初始化定时任务
	InitCron()
	// 运行服务
	return true
}

// func InitModel() {
// 	needMigrate := !tools.Exists("alist.db")
// 	fmt.Println(needMigrate)
// 	if needMigrate {
// 		database.InitDB()
// 	}
// }

// read config file
func ReadConf(name string) bool {
	// fmt.Println("读取配置文件...")
	if !tools.Exists(name) {
		fmt.Println("找不到配置文件:", name)
		return false
	}

	confFile, err := ioutil.ReadFile(name) // 文件读取
	if err != nil {
		fmt.Println("读取配置文件时发生错误:", err.Error())
		return false
	}
	// 此处从外围调用
	err = yaml.Unmarshal(confFile, config.Conf) // 入口字节，出口{}
	if err != nil {
		fmt.Printf("加载配置文件时发生错误:%s", err.Error())
		return false
	}
	return true
}

func printLogo() {
	fmt.Print(`

	██████╗  █████╗ ██╗██╗   ██╗██╗   ██╗███████╗
	██╔══██╗██╔══██╗██║╚██╗ ██╔╝██║   ██║██╔════╝
	██████╔╝███████║██║ ╚████╔╝ ██║   ██║█████╗  
	██╔══██╗██╔══██║██║  ╚██╔╝  ██║   ██║██╔══╝  
	██████╔╝██║  ██║██║   ██║   ╚██████╔╝███████╗
	╚═════╝ ╚═╝  ╚═╝╚═╝   ╚═╝    ╚═════╝ ╚══════╝
												 
			  

`)
}
