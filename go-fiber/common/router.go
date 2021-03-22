package common

import "github.com/gofiber/fiber/v2"

func Router(app *fiber.App) {
	// 静态文件目录
	app.Static("/", "dist") // 以主程序的路径为根目录
	// api接口
	api := app.Group("/api")
	{
		api.Get("/info", Info)
		api.Post("/get", Get)
		api.Post("/path", Path)
		api.Post("/office_preview", OfficePreview)
		// api.Post("/local_search", LocalSearch)
		// api.Post("/globle_search", GlobleSearch)
		api.Get("/rebuild/:password", Rebuild)
	}
	app.Get("/d/*", Down)
	app.Get("/root/*", BackHome)
}
