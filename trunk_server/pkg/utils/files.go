package utils

import (
	"io/fs"
	"path/filepath"
	"strings"
)

//type WalkDirFunc func(path string, d DirEntry, err error) error

// GetFiles 函数用于遍历指定目录下的所有 .doc 文件，并返回一个包含文件路径的切片
func GetFiles(uploadPath, ext string) ([]string, error) {
	files := make([]string, 0)
	err := filepath.WalkDir(uploadPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// 判断是否是文件
		if !d.IsDir() {
			// 获取文件名
			fileName := d.Name()

			// 判断文件扩展名是否为指定扩展名
			if strings.HasSuffix(fileName, ext) {
				// 将文件路径添加到切片中
				files = append(files, path)
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}

func GetFileNameWithoutExt(fileName string) string {
	dotIndex := strings.LastIndex(fileName, ".")
	if dotIndex == -1 {
		return fileName
	}
	return fileName[:dotIndex]
}
