package utils

import (
	"errors"
	"os"
	"time"
)

func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func GetFileSize(path string) int64 {
	if !exists(path) {
		return 0
	}
	fileInfo, err := os.Stat(path)
	if err != nil {
		return 0
	}
	return fileInfo.Size()
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

/**
 * 获取文件修改时间
 */
func GetFileModTime(path string) (time.Time, error) {
	if !exists(path) {
		return time.Time{}, errors.New("path is not exist")
	}

	fileInfo, err := os.Stat(path)
	if err != nil {
		return time.Time{}, errors.New("fail to get file info.")
	}

	return fileInfo.ModTime(), nil
}
