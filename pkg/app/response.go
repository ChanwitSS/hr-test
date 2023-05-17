package app

import (
	"hr/pkg/app/e"

	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (g *Gin) Response(httpCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: httpCode,
		Msg:  e.GetMsg(httpCode),
		Data: data,
	})
}
