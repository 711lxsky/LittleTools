package config

import "gorm.io/gorm"

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
	NormalClipExpireDays = 3
	UserClipMaxCapacity  = 200
)

var (
	IdentifierLength  = 6
	FileNameMaxLength = 10
	FileNameMinLength = 4
	UserNameMinLength = 3
	UserNameMaxLength = 16
)

var (
	PageDataDefaultNum  = 1
	PageDataDefaultSize = 10
)

var (
	ImageDir = "image"
)

var (
	HeaderContentType = "Content-Type"
	MultipartForm     = "multipart/form-data"
	ApplicationJson   = "application/json"
)

var (
	TokenExpireTimeDays      = 14
	TokenClaimAuthorized     = "authorized"
	TokenClaimUserId         = "userId"
	TokenClaimAuthority      = "authority"
	TokenClaimAuthorityValue = "711lxsky"
	TokenClaimExpireTime     = "expireTime"
	TokenSecret              = "KWrY8RHEjc6^81gd"
	TokenName                = "Authorization"
	TokenHeader              = "Bearer "
)

var (
	EmailSendHost = "smtp.qq.com"
	EmailSendPort = 25
	EmailSendUser = ""
	EmailSendPass = ""
)
