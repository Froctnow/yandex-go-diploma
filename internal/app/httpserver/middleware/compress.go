package middleware

import (
	"compress/gzip"
	"fmt"

	"github.com/gin-gonic/gin"
)

func CompressMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		headerAcceptEncoding := c.GetHeader("Accept-Encoding")

		if headerAcceptEncoding != "gzip" {
			c.Next()
			return
		}

		gz := gzip.NewWriter(c.Writer)

		c.Writer = &gzipWriter{c.Writer, gz}
		c.Header("Content-Encoding", "gzip")
		defer func() {
			gz.Close()
			c.Header("Content-Length", fmt.Sprint(c.Writer.Size()))
		}()
		c.Next()
	}
}

type gzipWriter struct {
	gin.ResponseWriter
	writer *gzip.Writer
}

func (g *gzipWriter) WriteString(s string) (int, error) {
	g.Header().Del("Content-Length")
	return g.writer.Write([]byte(s))
}

func (g *gzipWriter) Write(data []byte) (int, error) {
	g.Header().Del("Content-Length")
	return g.writer.Write(data)
}

func (g *gzipWriter) WriteHeader(code int) {
	g.Header().Del("Content-Length")
	g.ResponseWriter.WriteHeader(code)
}
