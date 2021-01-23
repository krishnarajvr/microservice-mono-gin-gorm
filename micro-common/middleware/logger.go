package middleware

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

//LoggerToFile Middleware
func LoggerToFile(filePath string, fileName string) gin.HandlerFunc {

	logFile := path.Join(filePath, fileName)

	src, err := os.OpenFile(logFile, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}

	logger := logrus.New()

	logger.Out = src

	logger.SetLevel(logrus.DebugLevel)

	//Set rotatelogs
	logWriter, err := rotatelogs.New(
		//Split file name
		fileName+".%Y%m%d.log",

		//Generate soft chain, point to the latest log file
		rotatelogs.WithLinkName(fileName),

		//Set maximum save time (7 days)
		rotatelogs.WithMaxAge(7*24*time.Hour),

		//Set log cutting interval (1 day)
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	//New hook
	logger.AddHook(lfHook)

	return func(c *gin.Context) {
		//Start time
		startTime := time.Now()

		//Set log file available in the context
		c.Set("log", logger)

		//Process request
		c.Next()

		//End time
		endTime := time.Now()

		//Execution time
		latencyTime := endTime.Sub(startTime)

		//Request method
		reqMethod := c.Request.Method

		reqURI := c.Request.RequestURI

		// status code
		statusCode := c.Writer.Status()

		// request IP
		clientIP := c.ClientIP()

		//Log format
		logger.WithFields(logrus.Fields{
			"statusCode":  statusCode,
			"latencyTime": latencyTime,
			"clientIp":    clientIP,
			"reqMethod":   reqMethod,
			"reqUri":      reqURI,
		}).Info()
	}
}
