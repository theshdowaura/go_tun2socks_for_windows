package database

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

func Database(ThreadCount int) error {
	// 获取当前目录
	currentDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("无法获取当前目录：%v", err)
	}

	// 获取上级目录的路径
	//currentDir := filepath.Dir(currentDir)
	databaseDir := filepath.Join(currentDir, "database")

	// 创建数据库目录
	err = os.MkdirAll(databaseDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("无法创建数据库目录：%v", err)
	}

	// 创建SQLite数据库文件
	dbPath := filepath.Join(databaseDir, "rules.db")
	_, err = os.Create(dbPath)
	if err != nil {
		return fmt.Errorf("无法创建SQLite数据库文件：%v", err)
	}

	// 打开SQLite数据库连接
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return fmt.Errorf("无法打开数据库连接：%v", err)
	}
	defer db.Close()

	// 创建表格
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS rules (id INTEGER PRIMARY KEY AUTOINCREMENT, filename TEXT, ip TEXT)")
	if err != nil {
		return fmt.Errorf("无法创建表格：%v", err)
	}

	// 遍历rules目录下的文件
	rulesDir := filepath.Join(currentDir, "rules")
	files, err := ioutil.ReadDir(rulesDir)
	if err != nil {
		return fmt.Errorf("无法读取目录：%v", err)
	}

	var wg sync.WaitGroup

	threadCount := ThreadCount // 设置线程数量

	var mutex sync.RWMutex // 创建读写锁

	for i := 0; i < threadCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for j := i; j < len(files); j += threadCount {
				file := files[j]

				// 判断是否为文件
				if !file.IsDir() {
					// 读取文件内容
					content, err := ioutil.ReadFile(filepath.Join(rulesDir, file.Name()))
					if err != nil {
						fmt.Println("读取文件时发生错误：", err)
						return
					}

					// 提取文件名和IP地址列表
					filename := file.Name()
					ips := extractIPs(content)

					// 获取写锁
					mutex.Lock()
					// 插入数据
					for _, ip := range ips {
						_, err = db.Exec("INSERT INTO rules (filename, ip) VALUES (?, ?)", filename, ip)
						if err != nil {
							fmt.Println("无法插入数据:", err)
							mutex.Unlock() // 发生错误时释放写锁
							return
						}
					}
					// 释放写锁
					mutex.Unlock()
				}
			}
		}()
	}

	wg.Wait()

	fmt.Println("数据库操作完成")
	return nil
}

// 提取IP地址列表
func extractIPs(content []byte) []string {
	var ips []string
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if !strings.HasPrefix(line, "#") && len(line) > 0 {
			ips = append(ips, line)
		}
	}
	return ips
}
