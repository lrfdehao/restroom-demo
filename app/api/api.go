package api

import (
	"demo/domain/context"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	// SUCCESS 前端指定的请求成功码
	SUCCESS = 0
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendOkResponse(c *gin.Context, data interface{}) {
	// always return http.StatusOK
	c.JSON(http.StatusOK, Response{
		Code:    SUCCESS,
		Message: "",
		Data:    data,
	})
}

func SendOkStringResponse(c *gin.Context, data string) {
	// always return http.StatusOK
	c.String(http.StatusOK, data)
}

func SendFailResponse(c *gin.Context, err *context.ServerError, data ...interface{}) {
	// always return http.StatusOK
	c.JSON(http.StatusOK, Response{
		Code:    err.Key,
		Message: err.DefaultError,
		Data:    data,
	})
}
