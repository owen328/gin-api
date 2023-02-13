package common

import (
	"github.com/sirupsen/logrus"
)

var Logger = logrus.New()

func init() {
	initLogger()
}
func initLogger() {
	//Logger := logrus.New()
	Logger.SetLevel(logrus.TraceLevel)
	//Logger.SetReportCaller(false)
	Logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05 UTC-07",
	})
}
