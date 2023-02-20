package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/pkg/logger"
	"time"
)

type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w AccessLogWriter) Write(p []byte) (int, error) {
	if n, err := w.body.Write(p); err != nil {
		return n, err
	}
	return w.ResponseWriter.Write(p)
}

func AccessLog() gin.HandlerFunc {
	return func(context *gin.Context) {
		bodyWriter := &AccessLogWriter{ResponseWriter: context.Writer, body: bytes.NewBufferString("")}
		context.Writer = bodyWriter

		beginTime := time.Now().Unix()
		context.Next()
		endTime := time.Now().Unix()

		fields := logger.Fields{
			"request":  context.Request.PostForm.Encode(),
			"response": bodyWriter.body.String(),
		}
		//fields未打印
		global.Logger.WithFields(fields).Infof("access log: method: %s, status_code: %d, begin_time: %d, end_time: %d",
			context.Request.Method,
			bodyWriter.Status(),
			beginTime,
			endTime)
	}
}
