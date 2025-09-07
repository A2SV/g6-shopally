package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/shopally-ai/pkg/domain"
)

type RateLimiter struct {
	RedisClient *redis.Client
	Limit       int           // max request
	Window      time.Duration // time window
}

func NewRateLimiter(redisAddr string, limit int, window time.Duration) *RateLimiter {
	rdb := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	return &RateLimiter{
		RedisClient: rdb,
		Limit:       int(limit),
		Window:      window,
	}
}

func (rl *RateLimiter) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		deviceID := c.GetHeader("X-Device-ID")

		if deviceID == "" {
			c.JSON(
				http.StatusBadRequest,
				domain.Response{
					Data: nil,
					Error: map[string]interface{}{
						"code":    http.StatusBadRequest,
						"message": "Missing Device ID",
					},
				},
			)
			c.Abort()
			return
		}

		key := fmt.Sprintf("rate:%s", deviceID)

		// Increment counter
		count, err := rl.RedisClient.Incr(ctx, key).Result()
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				domain.Response{
					Data: nil,
					Error: map[string]interface{}{
						"code":    http.StatusInternalServerError,
						"message": "Redis error",
					},
				},
			)
			c.Abort()
			return
		}

		// Set expiry if key was just created
		if count == 1 {
			if err := rl.RedisClient.Expire(ctx, key, rl.Window).Err(); err != nil {
				log.Printf("failed to set expiry for %s: %v", key, err)
			}
		}

		log.Printf("rate limit count for %s: %d (limit %d)", deviceID, count, rl.Limit)

		// Block if over limit
		if int(count) > rl.Limit {
			c.Header("Retry-After", fmt.Sprintf("%.0f", rl.Window.Seconds()))
			c.JSON(
				http.StatusTooManyRequests, // 429
				domain.Response{
					Data: nil,
					Error: map[string]interface{}{
						"code":    http.StatusTooManyRequests,
						"message": "Rate limit exceeded",
					},
				},
			)
			c.Abort()
			return
		}

		c.Next()
	}
}
