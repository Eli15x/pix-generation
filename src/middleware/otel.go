package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/attribute"
)

func OTelMetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		elapsed := time.Since(start).Seconds()

		route := c.FullPath()
		if route == "" { route = c.Request.URL.Path }

		attrs := []attribute.KeyValue{
			attribute.String("http.route", route),
			attribute.String("http.method", c.Request.Method),
			attribute.Int("http.status_code", c.Writer.Status()),
		}

		ReqCounter.Add(c, 1, attrs...)
		ReqLatencyHist.Record(c, elapsed, attrs...)
	}
}
