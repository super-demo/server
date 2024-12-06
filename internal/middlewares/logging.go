package middlewares

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func Logger(logger *slog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		t := time.Now()
		w := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: ctx.Writer}
		ctx.Writer = w

		ctx.Next()

		// 200 [POST] – /path
		host := ctx.Request.Host
		responseStatusCode := ctx.Writer.Status()
		uri := ctx.Request.RequestURI
		method := ctx.Request.Method

		body, _ := ioutil.ReadAll(ctx.Request.Body)

		switch {
		case ctx.Writer.Status() >= http.StatusBadGateway:
			logger.Error(
				fmt.Sprintf("%d [%s] – %s", responseStatusCode, method, uri),
				"x-correlation-id", ctx.GetString("correlationID"),
				"x-platform-fid", ctx.GetString("platformID"),
				"x-consumer-id", ctx.GetString("consumerID"),
				"x-consumer-username", ctx.GetString("consumerUsername"),
				"user-agent", ctx.Request.UserAgent(),
				"method", method,
				"host", host,
				"client_ip", ctx.ClientIP(),
				"path", uri,
				"content_type", ctx.ContentType(),
				"request_body", string(body),
				"body_size", ctx.Writer.Size(),
				"response_body", w.body.String(),
				"response_code", ctx.Writer.Status(),
				"response_time", fmt.Sprintf("%v", time.Since(t)),
			)
		default:
			logger.Info(
				fmt.Sprintf("%d [%s] – %s", responseStatusCode, method, uri),
				"x-correlation-id", ctx.GetString("correlationID"),
				"x-platform-fid", ctx.GetString("platformID"),
				"x-consumer-id", ctx.GetString("consumerID"),
				"x-consumer-username", ctx.GetString("consumerUsername"),
				"user-agent", ctx.Request.UserAgent(),
				"method", method,
				"host", host,
				"client_ip", ctx.ClientIP(),
				"path", uri,
				"content_type", ctx.ContentType(),
				"request_body", string(body),
				"body_size", ctx.Writer.Size(),
				"response_body", w.body.String(),
				"response_code", ctx.Writer.Status(),
				"response_time", fmt.Sprintf("%v", time.Since(t)),
			)
		}
	}
}
