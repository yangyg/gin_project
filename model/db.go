package model

import (
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql" // 导入mysql driver
	"github.com/jinzhu/gorm"
)

// Db ...
var Db *gorm.DB

// InitDb initial database
func InitDb(mysqlDSN string) {
	var err error
	Db, err = gorm.Open("mysql", mysqlDSN)
	if err != nil {
		log.Panicln("err:", err.Error())
	}
	Db.LogMode(true)
	//设置连接池
	//空闲
	Db.DB().SetMaxIdleConns(50)
	//打开
	Db.DB().SetMaxOpenConns(100)
	//超时
	Db.DB().SetConnMaxLifetime(time.Second * 30)
	migration()
}

// auto migrate
func migration() {
	Db.AutoMigrate(User{}, Subject{})
}

// CloseDB close database
func CloseDB() {
	Db.Close()
}
