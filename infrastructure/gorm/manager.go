package gorm

import (
	"github.com/jinzhu/gorm"
	"time"
	// orm
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

// Database 在中间件中初始化mysql链接,
func Database() {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/restroom?charset=utf8&parseTime=true")
	db.LogMode(true)
	// Error
	if err != nil {
		panic(err)
	}

	//设置连接池
	db.DB().SetMaxIdleConns(20)                  //空闲连接
	db.DB().SetMaxOpenConns(100)                 //最大连接数
	db.DB().SetConnMaxLifetime(time.Second * 30) //超时时间

	// 初始化
	DB = db
	migration()
}

// 数据表结构生成 (仅支持建新表，已存在的表不会修改字段或删除字段)
func migration() {
	DB.AutoMigrate(&ToiletModel{})
}
