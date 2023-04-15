package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

func Logger() gin.HandlerFunc {
	logger := logrus.New()
	logger.SetLevel(logrus.GetLevel())
	logger.SetFormatter(&logrus.TextFormatter{})
	return func(c *gin.Context) {
		startTime := time.Now().UnixMilli()
		c.Next()
		endTime := time.Now().UnixMilli()
		latencyTime := endTime - startTime
		logger.WithFields(logrus.Fields{
			"code":      c.Writer.Status(),
			"client_ip": c.ClientIP(),
			"url":       c.Request.RequestURI,
			"method":    c.Request.Method,
			"req_time":  latencyTime,
		}).Info("gin")
	}
}
