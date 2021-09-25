package util

import (
	"fmt"
	"os"
)

func FileOrDirExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func FileOrDirIsNotExists(path string) bool {
	return !FileOrDirExists(path)
}

func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func IsFile(path string) bool {
	return !IsDir(path)
}

func MkDirForce(path string) error {
	if FileOrDirIsNotExists(path) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			fmt.Println("创建文件夹失败,error info:", err)
			return err
		}
		return err
	}
	return nil
}
