package service

import (
	"alist/alidrive"
	"alist/config"
	"alist/models"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

func Clear() error {
	return config.DB.Where("1 = 1").Delete(&models.File{}).Error
}

// 创建目录树
func BuildTree() error {
	rootFile := models.File{
		Dir:    "",
		FileID: config.Conf.AliDrive.RootFolder,
		Name:   "root",
		Type:   "folder",
	}

	if err := config.DB.Create(&rootFile).Error; err != nil {
		config.DB.Rollback()
	}

	if err := BuildOne(config.Conf.AliDrive.RootFolder, "root/", config.DB, ""); err != nil {
		config.DB.Rollback()
		fmt.Println("构建失败")
		return err
	}
	return nil
}

func BuildOne(parent string, path string, tx *gorm.DB, parentPassword string) error {
	files := alidrive.GetList(parent, config.Conf.AliDrive.MaxFilesCount, "", "", "")

	if files == nil {
		return nil
	}
	for _, file := range files.Items {
		name := file.Name
		// fmt.Println(name)
		if strings.HasSuffix(name, ".hide") {
			continue
		}
		password := parentPassword
		if strings.Contains(name, ".password-") {
			index := strings.Index(name, ".password-")
			name = file.Name[:index]
			password = file.Name[index+10:]
		}
		newFile := models.File{
			Dir:           path,
			FileExtension: file.FileExtension,
			FileID:        file.FileID,
			Name:          name,
			Type:          file.Type,
			UpdatedAt:     file.UpdatedAt,
			Category:      file.Category,
			ContentType:   file.ContentType,
			Size:          file.Size,
			Password:      password,
		}
		if err := tx.Create(&newFile).Error; err != nil {
			return err
		}
		if file.Type == "folder" {
			if err := BuildOne(file.FileID, fmt.Sprintf("%s%s/", path, name), tx, password); err != nil {
				return err
			}
		}
	}
	return nil
}

func GetFileByDirAndName(dir, name string) *models.File {
	var file models.File
	// fmt.Println(dir, name, "内部函数")
	if err := config.DB.Where("dir = ? AND name = ?", dir, name).First(&file).Error; err != nil {
		// fmt.Println(err)
		// panic(err)
		return nil
	}
	return &file
}

func GetFilesByDir(dir string) (*[]models.File, error) {
	var files []models.File
	if err := config.DB.Where("dir = ?", dir).Find(&files).Error; err != nil {
		return nil, err
	}
	return &files, nil
}
