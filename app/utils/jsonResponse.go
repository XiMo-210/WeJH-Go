package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wejh-go/app/utils/stateCode"
)

func JsonSuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"data": data,
		"code": stateCode.OK,
		"msg":  "OK",
	})
}

func JsonFailedResponse(c *gin.Context, code stateCode.StateCode, data interface{}) {
	c.JSON(http.StatusBadRequest, gin.H{
		"data": data,
		"code": code,
		"msg":  stateCode.GetStateCodeMsg(code),
	})
}

func JsonErrorResponse(c *gin.Context, err error) {
	code := stateCode.ErrorToStateCode(err)
	c.JSON(http.StatusBadRequest, gin.H{
		"data": nil,
		"code": code,
		"msg":  stateCode.GetStateCodeMsg(code),
	})
}
