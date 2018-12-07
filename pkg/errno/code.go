package errno

var (
	// Common errors
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error."}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred whild binding the request body to the struct."}

	ErrValidation = &Errno{Code: 20001, Message: "Validation failed."}
	ErrDatabase   = &Errno{Code: 20002, Message: "Database error."}
	ErrToken      = &Errno{Code: 20003, Message: "Error occurred while signing the JSON web token."}
	ErrRedisConn  = &Errno{Code: 20004, Message: "Redis Conn error."}
	ErrRequestSms = &Errno{Code: 20005, Message: "Error occurred while request sms server"}

	// user errors
	ErrEncrypt           = &Errno{Code: 20101, Message: "Error occurred while encrypting the user password."}
	ErrUserNotFound      = &Errno{Code: 20102, Message: "The user was not found."}
	ErrTokenInvalid      = &Errno{Code: 20103, Message: "The token was invalid."}
	ErrPasswordIncorrect = &Errno{Code: 20104, Message: "The password was incorrect."}
	ErrVcodeNotFound     = &Errno{Code: 20105, Message: "The vcode was not found."}
	ErrUserHasRegist     = &Errno{Code: 20106, Message: "The phone was used to registed"}

	// upload errors
	ErrUploadExt  = &Errno{Code: 20201, Message: "The file ext invalid."}
	ErrUploadFail = &Errno{Code: 20202, Message: "Error occurred while upload file."}

	// admin errors
	ErrAdminHasRegist     = &Errno{Code: 20301, Message: "The username was used to registed"}
	ErrAdminNotFound      = &Errno{Code: 20302, Message: "The admin was not found."}
)
