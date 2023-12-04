package config

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func ReadRuleData() {
	// 获取当前目录
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal("无法获取当前目录：", err)
	}

	// 获取上级目录的路径
	parentDir := filepath.Dir(currentDir)
	filePath := filepath.Join(parentDir, "rules")
	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("无法打开文件：%s\n", err.Error())
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	// 创建一个用于读取文件的Scanner
	scanner := bufio.NewScanner(file)
	var rule []string
	// 逐行读取文件内容
	for scanner.Scan() {
		line := scanner.Text()

		// 忽略以 "#" 开头的注释行
		if strings.HasPrefix(line, "#") {
			continue
		}

		// 将非注释内容导入到 ip_data 的数据接口
		var ipData string
		ipData = line
		rule = append(rule, ipData)

	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("读取文件时发生错误：%s\n", err.Error())
		return
	}
	return
}
