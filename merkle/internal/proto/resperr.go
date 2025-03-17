package proto

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type ResponseMsg struct {
	Error string `json:"error,omitempty"`
}

type ResponseSuccessMsg struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func ErrorMsg(c *gin.Context, err string) {
	logrus.WithField("error", err).Info("Error occurred")
	c.AbortWithStatusJSON(http.StatusBadRequest,
		ResponseMsg{
			Error: err,
		})
}

func SuccessMsg(c *gin.Context, code int, msg string, data interface{}) {
	logrus.WithFields(logrus.Fields{
		"code": code,
		"msg":  msg,
		"data": data,
	}).Info("Success message")
	c.AbortWithStatusJSON(http.StatusOK, ResponseSuccessMsg{
		Code:    code,
		Message: msg,
		Data:    data,
	})
}
