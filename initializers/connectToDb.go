package initializers

import (
	"gorm.io/driver/sqlite" // 基于 GGO 的 Sqlite 驱动
	// "github.com/glebarez/sqlite" // 纯 Go 实现的 SQLite 驱动, 详情参考： https://github.com/glebarez/sqlite
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	var err error
	// github.com/mattn/go-sqlite3
	DB, err = gorm.Open(sqlite.Open("jwt-go.db"), &gorm.Config{})
	if err != nil {
		panic("fail to connect to database")
	}
}
