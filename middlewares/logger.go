package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// all the middleware must have this gin.HandlerFunc as their return type
func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("Output from Logger() - %s - [%s] %s %s %d %s \n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC822),
			param.Method,
			param.Path,
			param.StatusCode,
			param.Latency)
	})
}
