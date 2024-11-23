package util

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"goer/config"
	myErr "goer/error"
	"goer/model"
)

// InitDB 初始化数据库连接，并返回数据库实例
// 该函数使用了defer来确保在函数结束时关闭数据库连接
// 使用AutoMigrate方法来确保数据库模式与模型匹配
func InitDB() {
	// 打开数据库连接
	var err error
	config.DataBase, err = gorm.Open(config.ConnectDB, config.DBName)
	if err != nil {
		panic(myErr.CannotConnectDB + err.Error())
	}
	// 自动迁移模式， 保持更新到最新
	// 仅创建表， 缺少列和索引， 不会改变现有列的类型或删除未使用的列以保护数据
	config.DataBase.AutoMigrate(
		&model.Clip{},
	)
}
