package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go_learn/common"
	"os"
	"path"
)

func LogRequestMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		src, err := getRequestLogFilePath()
		if err != nil {
			common.Logger.Errorln(err)
		}
		defer src.Close()
		common.Logger.SetOutput(src)
		common.Logger.WithFields(logrus.Fields{
			"user_id": c.GetString("user_id"),
		}).Infof("%s %s %d IP:%s",
			c.Request.Method,
			c.Request.RequestURI,
			c.Writer.Status(),
			c.ClientIP(),
		)
	}
}

func getRequestLogFilePath() (*os.File, error) {
	curDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	logDir := path.Join(curDir, "logs")
	err = os.MkdirAll(logDir, 0777)
	if err != nil {
		return nil, err
	}
	logFilePath := path.Join(logDir, "request.log")
	src, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, err
	}
	return src, nil
}
