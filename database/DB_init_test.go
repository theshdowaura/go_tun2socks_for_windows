package database

import (
	"database/sql"
	"fmt"
	//"github.com/andlabs/ui"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestDatabase(t *testing.T) {
	currentDir, err := os.Getwd()
	if err != nil {
		t.Fatal("无法获取当前目录：", err)
	}

	// 获取上级目录的路径
	parentDir := filepath.Dir(currentDir)
	databaseDir := filepath.Join(parentDir, "database")

	// 创建数据库目录
	err = os.MkdirAll(databaseDir, os.ModePerm)
	if err != nil {
		t.Fatal("无法创建数据库目录：", err)
	}

	// 创建SQLite数据库文件
	dbPath := filepath.Join(databaseDir, "rules.db")
	_, err = os.Create(dbPath)
	if err != nil {
		t.Fatal("无法创建SQLite数据库文件：", err)
	}

	// 打开SQLite数据库连接
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		t.Fatal("无法打开数据库连接:", err)
	}
	defer db.Close()

	// 创建表格
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS rules (id INTEGER PRIMARY KEY AUTOINCREMENT, filename TEXT, ip TEXT)")
	if err != nil {
		t.Fatal("无法创建表格:", err)
	}

	// 遍历rules目录下的文件
	rulesDir := filepath.Join(parentDir, "rules")
	err = filepath.Walk(rulesDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("遍历目录时发生错误：", err)
			return nil
		}

		// 判断是否为文件
		if !info.IsDir() {
			// 读取文件内容
			content, err := ioutil.ReadFile(path)
			if err != nil {
				fmt.Println("读取文件时发生错误：", err)
				return nil
			}

			// 提取文件名和IP地址列表
			filename := filepath.Base(path)
			ips := extractIPs(content)

			// 插入数据
			for _, ip := range ips {
				_, err = db.Exec("INSERT INTO rules (filename, ip) VALUES (?, ?)", filename, ip)
				if err != nil {
					fmt.Println("无法插入数据:", err)
					return nil
				}
			}
		}

		return nil
	})

	if err != nil {
		t.Fatal("遍历目录时发生错误：", err)
	}

	fmt.Println("数据库操作完成")
}
func TestDbOut(t *testing.T) {
	// 打开数据库连接
	db, err := sql.Open("sqlite3", "rules.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 查询符合条件的数据
	rows, err := db.Query("SELECT ip FROM rules WHERE filename = 'Arma3.rules'")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// 导出数据
	for rows.Next() {
		var ip string
		err := rows.Scan(&ip)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(ip)
	}

	// 检查是否有错误
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

// 提取IP地址列表
//
//	func extractIPs(content []byte) []string {
//		var ips []string
//		lines := strings.Split(string(content), "\n")
//		for _, line := range lines {
//			line = strings.TrimSpace(line)
//			if !strings.HasPrefix(line, "#") && len(line) > 0 {
//				ips = append(ips, line)
//			}
//		}
//		return ips
//	}
