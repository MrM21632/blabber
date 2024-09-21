package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		start := time.Now()
		context.Next()
		end := time.Now()

		latency := end.Sub(start)
		req_method := context.Request.Method
		req_uri := context.Request.RequestURI
		status := context.Writer.Status()
		client := context.ClientIP()

		log.WithFields(log.Fields{
			"method":      req_method,
			"uri":         req_uri,
			"status_code": status,
			"latency":     latency,
			"client_ip":   client,
		}).Info("processed incoming request")
	}
}
