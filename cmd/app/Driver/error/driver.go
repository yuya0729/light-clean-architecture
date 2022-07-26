package myerror

import (
	"fmt"
)

type MyError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (me *MyError) Error() string {
	return fmt.Sprintf(`{"code": %d, "message": %s}`, me.Code, me.Message)
}

// New コンストラクタ
func New(code int, message string) *MyError {
	return &MyError{
		Code:    code,
		Message: message,
	}
}

// {"message": "Not found", "code": 404}
// if err.code == 404 {
// jsonresponse(statusnotfoud, err)
// }
