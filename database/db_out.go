package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func DbOut(game string) []string {
	// 打开数据库连接
	db, err := sql.Open("sqlite3", "database/rules.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 查询符合条件的数据
	rows, err := db.Query("SELECT ip FROM rules WHERE filename = ?", game)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// 导出数据
	var ips []string
	for rows.Next() {
		var ip string
		err := rows.Scan(&ip)
		if err != nil {
			log.Fatal(err)
		}
		ips = append(ips, ip)
	}

	// 检查是否有错误
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return ips
}
