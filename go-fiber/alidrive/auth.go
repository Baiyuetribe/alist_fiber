package alidrive

import (
	"alist/config"
	"alist/models"
	"fmt"
	"io/ioutil"

	"github.com/levigross/grequests"
	"gopkg.in/yaml.v2"
)

func GetUserInfo() {
	url := config.Conf.AliDrive.ApiUrl + "/user/get"
	res := Post(url, map[string]string{})
	// fmt.Println(res)
	if err := res.JSON(&User); err != nil {
		panic(err)
	}
	// fmt.Println(User)	// 平时打印效果为值，没有key
	// fmt.Println(User.Phone)
}

// 获取jwt-token
func RefreshToken(token string) bool {
	// 刷新tokne
	url := "https://websv.aliyundrive.com/token/refresh"
	// fmt.Println(token, url)
	res, err := grequests.Post(url, &grequests.RequestOptions{
		JSON: map[string]string{"refresh_token": token},
	})
	if err != nil {
		fmt.Println("请求出错")
		return false
	}
	// fmt.Println(res.String())
	type Item struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}
	itmes := new(Item)
	// itmes := &Item{AccessToken: "jdiejdijeidje"}
	// fmt.Println(res.String())
	if err := res.JSON(&itmes); err != nil {
		fmt.Println("参数解析失败")
		fmt.Println(err)
		return false
	}
	config.Conf.AliDrive.AccessToken = itmes.AccessToken   // 新值为：=，非新值使用等号=
	config.Conf.AliDrive.RefreshToken = itmes.RefreshToken // 新值为：=，非新值使用等号=
	config.Authorization = config.Bearer + itmes.AccessToken
	if data, err := yaml.Marshal(config.Conf); err != nil {
		fmt.Println("解析失败")
		return false
	} else {
		if err := ioutil.WriteFile("env.yml", data, 0777); err != nil { //该操作为全覆盖，
			fmt.Println("修改失败")
		}
	}
	return true
}

type Files struct {
	Items      []models.File `json:"items"`
	NextMarker string        `json:"next_marker"`
	Readme     string        `json:"readme"` // Deprecated
	Paths      []Path        `json:"paths"`
}

type Path struct {
	Name   string `json:"name"`
	FileId string `json:"file_id"`
}

// list request bean
type ListReq struct {
	DriveId               string `json:"drive_id"`
	Fields                string `json:"fields"`
	ImageThumbnailProcess string `json:"image_thumbnail_process"`
	ImageUrlProcess       string `json:"image_url_process"`
	Limit                 int    `json:"limit"`
	Marker                string `json:"marker"`
	OrderBy               string `json:"order_by"`
	OrderDirection        string `json:"order_direction"`
	ParentFileId          string `json:"parent_file_id"`
	VideoThumbnailProcess string `json:"video_thumbnail_process"`
}

func GetList(parent string, limit int, marker string, orderBy string, orderDirection string) *Files {
	url := config.Conf.AliDrive.ApiUrl + "/file/list"
	// fmt.Println(url)
	data := ListReq{
		DriveId:               User.DefaultDriveId,
		Fields:                "*",
		ImageThumbnailProcess: config.ImageThumbnailProcess,
		ImageUrlProcess:       config.ImageUrlProcess,
		Limit:                 limit,
		Marker:                marker,
		OrderBy:               orderBy,
		OrderDirection:        orderDirection,
		ParentFileId:          parent,
		VideoThumbnailProcess: config.VideoThumbnailProcess,
	}
	var resp Files
	res := Post2(url, data)
	if err := res.JSON(&resp); err != nil {
		fmt.Println(err)
		return nil
	}
	// fmt.Println(resp)
	return &resp
}

type DownloadResp struct {
	Expiration string `json:"expiration"`
	Method     string `json:"method"`
	Size       int64  `json:"size"`
	Url        string `json:"url"`
	//RateLimit struct{
	//	PartSize int `json:"part_size"`
	//	PartSpeed int `json:"part_speed"`
	//} `json:"rate_limit"`//rate limit
}

// download request bean
type DownloadReq struct {
	DriveId   string `json:"drive_id"`
	FileId    string `json:"file_id"`
	ExpireSec int    `json:"expire_sec"`
	FileName  string `json:"file_name"`
}

// get download_url
func GetDownLoadUrl(fileId string) *DownloadResp {
	url := config.Conf.AliDrive.ApiUrl + "/file/get_download_url"
	data := DownloadReq{
		DriveId:   User.DefaultDriveId,
		FileId:    fileId,
		ExpireSec: 14400,
	}
	var resp DownloadResp
	res := Post2(url, data)
	if res == nil {
		return nil
	}
	if err := res.JSON(&resp); err != nil {
		// fmt.Println(err)
		return nil
	}
	// fmt.Println(resp)
	return &resp
}

// office_preview_url response bean
type OfficePreviewUrlResp struct {
	PreviewUrl  string `json:"preview_url"`
	AccessToken string `json:"access_token"`
}

// get office preview url and token
func GetOfficePreviewUrl(fileId string) *OfficePreviewUrlResp {
	url := config.Conf.AliDrive.ApiUrl + "/file/get_office_preview_url"
	// office_preview_url request bean
	type OfficePreviewUrlReq struct {
		AccessToken string `json:"access_token"`
		DriveId     string `json:"drive_id"`
		FileId      string `json:"file_id"`
	}
	data := OfficePreviewUrlReq{
		AccessToken: config.Conf.AliDrive.AccessToken,
		DriveId:     User.DefaultDriveId,
		FileId:      fileId,
	}
	// fmt.Println(data)
	var resp OfficePreviewUrlResp
	res := Post2(url, data)
	if err := res.JSON(&resp); err != nil {
		fmt.Println(err)
	}
	return &resp
}
