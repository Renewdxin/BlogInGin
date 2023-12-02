package middleware

import (
	"BloginGin/pkg/app"
	"BloginGin/pkg/errcode"
	"BloginGin/pkg/limiter"
	"github.com/gin-gonic/gin"
)

// RateLimiter 返回一个基于令牌桶算法的 Gin 中间件，用于限制请求速率。
// 参数 l 是一个实现了 limiter.LimiterIface 接口的限流器实例。
func RateLimiter(l limiter.LimiterIface) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从上下文中获取唯一标识符作为令牌桶的 key
		key := l.Key(c)

		// 获取令牌桶实例和是否成功的标志
		if bucket, ok := l.GetBucket(key); ok {
			// 尝试获取一个令牌，返回实际获取的令牌数量
			count := bucket.TakeAvailable(1)

			// 如果没有足够的令牌，则认为超过请求速率限制
			if count == 0 {
				// 返回请求速率过高的错误响应
				response := app.NewResponse(c)
				response.ToErrorResponse(errcode.TooManyRequests)
				c.Abort()
				return
			}
		}
		// 请求通过限流检查，继续处理下一个中间件或路由处理程序
		c.Next()
	}
}
