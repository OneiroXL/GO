package DBBase

import (
	"time"
	"chat/model/base/entity"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB 数据库链接单例
var DB *gorm.DB

// Database 在中间件中初始化mysql链接
func Database(connString string) {
	db, err := gorm.Open("mysql", connString)
	db.LogMode(true)
	// Error
	if err != nil {
		panic(err)
	}
	//设置连接池
	//空闲
	db.DB().SetMaxIdleConns(20)
	//打开
	db.DB().SetMaxOpenConns(100)
	//超时
	db.DB().SetConnMaxLifetime(time.Second * 30)
	//设置全局表名禁用复数
	db.SingularTable(true)

	DB = db

	migration()
}


//执行数据迁移
func migration() {
	// 自动迁移模式
	DB.AutoMigrate(&entity.User{})
}
