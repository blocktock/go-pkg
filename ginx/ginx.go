package ginx

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ResOK(c *gin.Context) {
	ResSuccess(c, "")
}

type ResponseItem struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

type SimpleResponseItem struct {
	Message string `json:"Message"`
}

func ResSuccess(c *gin.Context, v interface{}) {
	responseItem := ResponseItem{
		Code:    0,
		Message: "",
		Data:    v,
	}
	ResJSON(c, http.StatusOK, responseItem)
}

func ResError(c *gin.Context, code int, msg string) {
	responseItem := ResponseItem{
		Code:    code,
		Message: msg,
	}
	ResJSON(c, http.StatusOK, responseItem)
}

func ResJSON(c *gin.Context, status int, v interface{}) {
	buf, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	c.Data(status, "application/json; charset=utf-8", buf)
	c.Abort()
}
