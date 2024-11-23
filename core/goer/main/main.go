package main

import (
	"github.com/jinzhu/gorm"
	"goer/config"
	myErr "goer/error"
	"goer/util"
	"log"
)

func main() {
	// 工作目录检测并创建必要文件夹
	util.GetWorkDirAndBuildImportantPath()
	// 初始化数据库连接
	util.InitDB()
	// 使用defer确保在函数结束时关闭数据库连接
	defer func(*gorm.DB) {
		err := config.DataBase.Close()
		if err != nil {
			panic(myErr.DataBaseCannotBeCorrectlyClosed + err.Error())
		}
	}(config.DataBase)
	// 初始化gin服务
	engine := util.InitGin()
	InitRouter(engine)
	err := engine.Run(config.RunPort)
	if err != nil {
		log.Fatal(err)
	}
}
