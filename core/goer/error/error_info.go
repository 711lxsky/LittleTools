package error

var (
	CannotConnectDB                 = "Cannot connect to database"
	DataBaseCannotBeCorrectlyClosed = "Data base cannot be correctly closed"
	DataParseError                  = "Data parse error"
	JSONParseError                  = "JSON parse error"
	UnsupportedMediaType            = "Unsupported media type"
	FileParseError                  = "File parse error"
	FileSaveError                   = "File save error"
	DataBaseSaveError               = "Data base save error"
	DataBaseQueryError              = "Data base query error"
	DataBaseDeleteError             = "Data base delete error"
	DataBaseUpdateError             = "Data base update error"
	ConfirmUserPassWordNotSame      = "User confirm password is not correct"
	RegisterUserNameLengthError     = "The length of user's name error"
	UserNameExisted                 = "User name is existed"
	PasswordEncryptError            = "Password encrypt error"
	DataNotExist                    = "Data not exist"
	PasswordError                   = "Password error"
	TokenExpired                    = "token expired"
	TokenInvalid                    = "token invalid"
	TokenClaimsError                = "token claims error"
	TokenAuthorityHeaderMissed      = "token authority header missed"
	ContextError                    = "context error"
	TypeAssertionError              = "type assertion error"
	DataCannotModify                = "data cannot be modified"
)
