package common

import (
	"alist/alidrive"
	"alist/config"
	"alist/service"
	"fmt"
	"net/url"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/levigross/grequests"
)

func Info(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"meta": fiber.Map{"code": 200, "msg": "success"}, "data": config.Conf.Info})
}

func Info2(c *fiber.Ctx) error {
	type Info struct {
		Url    string `json:"url"`
		Origin string `json:"origin"`
	}
	res, err := grequests.Post("http://httpbin.org/post", &grequests.RequestOptions{
		Data: map[string]string{"name": "lili", "password": "deded"},
		Headers: map[string]string{
			"Content-Type": "application/json",
			"Origin":       "https://aliyundrive.com",
		},
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36",
	})
	if err != nil {
		fmt.Println("请求错误")
	}
	fmt.Println(res.StatusCode)
	fmt.Println(res.String())
	fmt.Println(res.JSON(map[string]string{}))
	return c.SendString("你好info")
}

func Get(c *fiber.Ctx) error {
	type GetReq struct {
		Path     string `json:"path" binding:"required"`
		Password string `json:"password"`
	}
	var get GetReq
	if err := c.BodyParser(&get); err != nil {
		return c.Status(400).JSON(fiber.Map{"msg": "miss data"})
	}
	dir, name := filepath.Split(get.Path)
	file := service.GetFileByDirAndName(dir, name)
	if file == nil {
		return c.Status(400).JSON(fiber.Map{"msg": "path not found", "code": 400})
	}
	down := alidrive.GetDownLoadUrl(file.FileID)
	if down == nil {
		return c.Status(500).JSON(fiber.Map{"msg": "失败"})
	}
	// c.JSON(200, DataResponse(down))		// 待解决
	return c.JSON(fiber.Map{"meta": fiber.Map{"code": 200, "msg": "success"}, "data": down})
}

func Path(c *fiber.Ctx) error {
	type PathItem struct {
		Path     string `json:"path" binding:"required"`
		Password string `json:"password"`
	}
	var res PathItem // res := new(PathItem)
	if err := c.BodyParser(&res); err != nil {
		return c.Status(400).JSON(fiber.Map{"msg": "miss data"})
	}
	// 执行models
	dir, name := filepath.Split(res.Path)
	// fmt.Println(dir, name, "外部函数")
	file := service.GetFileByDirAndName(dir, name)
	// if file == nil {
	// 	return c.JSON(fiber.Map{"msg": "path not found"})
	// }
	// 处理单个文件
	if file.Type == "file" {
		return c.JSON(fiber.Map{"meta": fiber.Map{"code": 200, "msg": "success"}, "data": file})
	}
	// 处理文件夹
	files, err := service.GetFilesByDir(res.Path + "/")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"msg": err.Error()})
	}
	// for i,_ := range *files{

	// }
	return c.JSON(fiber.Map{"meta": fiber.Map{"code": 200, "msg": "success"}, "data": files})
}

func OfficePreview(c *fiber.Ctx) error {
	type OfficePreviewReq struct {
		FileId string `json:"file_id" binding:"required"`
	}
	var req OfficePreviewReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"msg": "miss data"})
	}
	preview := alidrive.GetOfficePreviewUrl(req.FileId)
	return c.JSON(fiber.Map{"meta": fiber.Map{"code": 200, "msg": "success"}, "data": preview})
}

func LocalSearch(c *fiber.Ctx) error {
	return c.SendString("你好local")
}

func GlobleSearch(c *fiber.Ctx) error {
	return c.SendString("你好gbs")
}

func BackHome(c *fiber.Ctx) error {
	return c.Redirect("/")
}

func Rebuild(c *fiber.Ctx) error {
	password := c.Params("password")
	if password != config.Conf.Server.Password {
		return c.JSON(fiber.Map{"msg": "wrong password."})
	}
	if err := service.Clear(); err != nil {
		return c.Status(500).JSON(fiber.Map{"msg": err.Error()})
	}
	if err := service.BuildTree(); err != nil {
		return c.Status(500).JSON(fiber.Map{"msg": err.Error()})
	}
	return c.JSON(fiber.Map{"code": 200, "msg": "success"})
}

func Down(c *fiber.Ctx) error {
	// filePath := c.Params("file")
	filePath, _ := url.QueryUnescape(c.Params("*"))
	// type DownReq struct {
	// 	Password string `form:"pw"`
	// }
	// var down DownReq
	dir, name := filepath.Split(filePath)
	fileModel := service.GetFileByDirAndName(dir, name)
	if fileModel == nil {
		return c.Status(404).JSON(fiber.Map{"msg": "not found", "code": 404})
	}
	// if fileModel.Password != "" && fileModel.Password != down.Password {
	// 	if down.Password == "" {
	// 		return c.JSON(fiber.Map{"msg": "need password"})
	// 	} else {
	// 		return c.JSON(fiber.Map{"msg": "wrong password"})
	// 	}
	// }
	if fileModel.Type == "folder" {
		return c.JSON(fiber.Map{"msg": "无法下载目录"})
	}
	file := alidrive.GetDownLoadUrl(fileModel.FileID)
	if file == nil {
		return c.JSON(fiber.Map{"msg": "下载出错"})
	}
	// fmt.Println(file.Url)
	return c.Status(301).Redirect(file.Url)
}
