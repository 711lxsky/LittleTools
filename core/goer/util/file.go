package util

import (
	"goer/config"
	"log"
	"os"
	"path/filepath"
)

func GetWorkDirAndBuildImportantPath() {
	// 获取 main.go 文件的绝对路径
	currentFile, err := filepath.Abs("main.go")
	if err != nil {
		log.Fatalf("获取 main.go 绝对路径失败: %v", err)
	}

	// 提取 main.go 文件的上上层目录
	parentDir := filepath.Dir(filepath.Dir(currentFile))
	config.WorkPath = parentDir
	// 创建 img 目录
	config.ImageDirPath = buildPath(parentDir, config.ImageDir)
}

func buildPath(parentDir, targetDir string) string {
	// 目录拼接
	targetDir = filepath.Join(parentDir, targetDir)
	if _, err := os.Stat(targetDir); os.IsNotExist(err) {
		// 目录不存在， 创建
		if err := os.Mkdir(targetDir, os.ModePerm); err != nil {
			log.Fatalf("创建目录失败: %v", err)
		}
		log.Printf("创建目录: %s", targetDir)
	} else if err != nil {
		log.Fatalf("获取目录状态失败: %v", err)
	} else {
		log.Printf("目录已存在: %s", targetDir)
	}
	return targetDir
}

func GenerateNewNameForFile(rawName string) string {
	// 先拿到扩展名
	extensionName := filepath.Ext(rawName)
	// 再生成一个范围长度的随机名称
	fileNewNamePrev, err := GenerateRandomStringWithRangeLength(config.FileNameMinLength, config.FileNameMaxLength)
	if err != nil {
		log.Fatalf("生成随机文件名失败: %v", err)
		return ""
	}
	// 拼接文件名
	return fileNewNamePrev + extensionName
}
