package alidrive

import (
	"alist/config"

	"github.com/levigross/grequests"
)

func Post(url string, data map[string]string) *grequests.Response {
	res, err := grequests.Post(url, &grequests.RequestOptions{
		JSON: data,
		Headers: map[string]string{
			"Content-Type":  "application/json",
			"Origin":        "https://aliyundrive.com",
			"authorization": config.Authorization,
		},
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36",
	})
	if err != nil {
		// panic(err)
		return nil
	}
	return res
}

func Post2(url string, data interface{}) *grequests.Response {
	res, err := grequests.Post(url, &grequests.RequestOptions{
		JSON: data,
		Headers: map[string]string{
			"Content-Type":    "application/json",
			"Origin":          "https://www.aliyundrive.com",
			"Accept":          "application/json, text/plain, */*",
			"Accept-Language": "zh-CN,zh;q=0.9,en;q=0.8,zh-TW;q=0.7",
			"Connection":      "keep-alive",
			"authorization":   config.Conf.AliDrive.AccessToken,
		},
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36",
	})
	if err != nil {
		// panic(err)
		return nil
	}
	// fmt.Println(res.StatusCode)

	// fmt.Println(res.String())
	return res
}
