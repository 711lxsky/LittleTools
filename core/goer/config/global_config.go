package config

import "github.com/jinzhu/gorm"

// RunPort 运行端口
var RunPort = ":8228"

// DataBase 全局数据库
var DataBase *gorm.DB

// WorkPath 全局工作路径
var WorkPath string

var ImageDirPath string

var (
	ConnectDB = "sqlite3"
	DBName    = "little_tools.db"
)

var (
	IdentifierLength  = 6
	FileNameMaxLength = 10
	FileNameMinLength = 4
	UserNameMinLength = 3
	UserNameMaxLength = 16
)

var (
	ImageDir = "image"
)
