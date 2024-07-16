package response

import "github.com/gin-gonic/gin"

type ResponseData struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	ErrCodeSuccess      = 20001
	ErrCodeParamInvalid = 20003
	ErrCodeInvalidToken = 403
)

var msg = map[int]string{
	ErrCodeSuccess:      "success",
	ErrCodeParamInvalid: "param invalid",
	ErrCodeInvalidToken: "invalid token",
}

func SuccessResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(code, ResponseData{
		Code: code,
		Msg:  msg[code],
		Data: data,
	})
}

func ErrorResponse(c *gin.Context, code int, message string) {
	c.JSON(code, ResponseData{
		Code: code,
		Msg:  msg[code],
		Data: nil,
	})
}
