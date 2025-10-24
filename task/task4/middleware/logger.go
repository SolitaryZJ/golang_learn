package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 记录请求开始时间
		start := time.Now()

		// 处理请求
		c.Next()

		// 记录请求结束时间
		end := time.Now()

		// 计算请求耗时
		latency := end.Sub(start)

		// 获取请求路径和请求方法
		path := c.Request.URL.Path
		method := c.Request.Method

		// 获取响应状态码
		statusCode := c.Writer.Status()
		fmt.Println("Request:", method, path, "Status:", statusCode, "Latency:", latency)
	}
}
