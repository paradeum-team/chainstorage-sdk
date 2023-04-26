package app

import (
	"chainstorage-sdk/common/e"
	"chainstorage-sdk/common/e/code"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response setting gin.JSON
func NewResponse(C *gin.Context, httpCode, errCode int, data interface{}) {
	C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  code.GetMsg(errCode),
		Data: data,
	})
	return
}

func NewOriginResponse(errCode int, data interface{}) Response {
	return Response{Code: errCode, Msg: code.GetMsg(errCode), Data: data}
}
func NewDecorateResponse(C *gin.Context, httpStatusCode int, response Response) {
	C.JSON(httpStatusCode, Response{
		Code: response.Code,
		Msg:  code.GetMsg(response.Code),
		Data: response.Data,
	})
	return
}

/**
 * 根据 bfsError 生成返回响应
 */
func NewSimpleBadResponse(C *gin.Context, bfsError e.BfsError) {
	msg := bfsError.GetNote()
	if len(msg) < 1 {
		msg = code.GetMsg(bfsError.GetSCode())
	}

	httpStatus := bfsError.GetHttpStatus()
	if httpStatus == 0 {
		httpStatus = http.StatusBadRequest
	}

	C.JSON(httpStatus, Response{
		Code: bfsError.GetSCode(),
		Msg:  msg,
		Data: bfsError.GetRaw(),
	})
	return
}

func NewBadResponse(C *gin.Context, httpCode int, bfsError e.BfsError) {
	msg := bfsError.GetNote()
	if len(msg) < 1 {
		msg = code.GetMsg(bfsError.GetSCode())
	}
	C.JSON(httpCode, Response{
		Code: bfsError.GetSCode(),
		Msg:  msg,
		Data: bfsError.GetRaw(),
	})
	return
}
