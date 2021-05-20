package middleware

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"opencodes/model/db"
	"opencodes/model/service"
	"time"

	"go.uber.org/zap"

	logs "github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
)

// 操作记录
func OperationRecord() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body []byte
		var userId int
		var err error
		if c.Request.Method != http.MethodGet {
			body, err = ioutil.ReadAll(c.Request.Body)
			if err != nil {

			} else {
				c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
			}
		}
		userId = 1
		record := db.SysOperationRecord{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Ip:        c.ClientIP(),
			Method:    c.Request.Method,
			Path:      c.Request.URL.Path,
			Agent:     c.Request.UserAgent(),
			Body:      string(body),
			UserID:    userId,
		}
		writer := responseBodyWriter{
			ResponseWriter: c.Writer,
			body:           &bytes.Buffer{},
		}
		c.Writer = writer
		now := time.Now()
		c.Next()
		latency := time.Now().Sub(now)
		record.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()
		record.Status = c.Writer.Status()
		record.Latency = latency
		record.Resp = writer.body.String()
		if err = service.CreateSysOperationRecord(record); err != nil {
			logs.Error("create operation record error:", zap.Any("err", err))
		}
	}
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}
