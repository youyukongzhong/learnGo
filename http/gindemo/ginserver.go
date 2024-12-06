package main

import (
	"math/rand"
	"net/http"
	"time"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个默认的 Gin 路由引擎实例，默认包含 Logger 和 Recovery 中间件
	r := gin.Default()

	// 创建一个日志记录器，用于记录程序运行中的重要信息
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err) // 如果日志初始化失败，程序终止
	}

	// 定义一个常量，用于在上下文中存储请求 ID
	const keyRequestId = "requestId"

	// middleware的使用
	// 注册全局中间件，所有路由请求都会执行这些中间件
	r.Use(
		// 第一个中间件：记录请求处理的时间
		func(c *gin.Context) {
			s := time.Now() // 记录请求开始的时间
			c.Next()        // 继续处理请求

			time.Sleep(50 * time.Millisecond) // 模拟延迟，方便观察时间记录（实际环境不需要）
			elapsed := time.Since(s)          // 计算请求处理完成所花费的时间
			// path, status, log latency
			// 记录日志信息：请求路径、状态码以及处理时长
			logger.Info("incoming request",
				zap.String("path", c.Request.URL.Path), // 记录请求路径
				zap.Int("status", c.Writer.Status()),   // 记录返回的 HTTP 状态码
				zap.Duration("elapsed", elapsed))       // 记录请求处理时长
		},
		// 第二个中间件：为每个请求生成一个随机的请求 ID
		func(c *gin.Context) {
			c.Set(keyRequestId, rand.Int()) // 设置一个随机的请求 ID，存储在上下文中

			// 继续处理请求
			c.Next()
		})

	// 定义路由 `/ping`，用于处理 `GET` 请求
	r.GET("/ping", func(c *gin.Context) {
		// 准备返回的 JSON 数据
		h := gin.H{
			"message": "pong", // 返回的消息内容
		}

		// 如果在上下文中存储了请求 ID，则添加到响应中
		if rid, exists := c.Get(keyRequestId); exists {
			h[keyRequestId] = rid // 添加请求 ID 到返回数据中
		}

		// 返回 JSON 格式的响应，HTTP 状态码为 200（OK）
		c.JSON(http.StatusOK, h)
	})

	// 定义路由 `/hello`，用于处理简单的 `GET` 请求
	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "hello")
	})

	// 启动 HTTP 服务，默认监听 0.0.0.0:8080
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
