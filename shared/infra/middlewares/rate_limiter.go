package middlewares

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type IPRateLimiter struct {
	ips           map[string]*rate.Limiter
	mu            *sync.RWMutex
	rateTime      rate.Limit
	requestsCount int
}

func NewIPRateLimiter(rateTime rate.Limit, requestsCount int) *IPRateLimiter {
	return &IPRateLimiter{
		ips:           make(map[string]*rate.Limiter),
		mu:            &sync.RWMutex{},
		rateTime:      rateTime,
		requestsCount: requestsCount,
	}
}

func (rl *IPRateLimiter) AddIP(ip string) *rate.Limiter {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	limiter := rate.NewLimiter(rl.rateTime, rl.requestsCount)
	rl.ips[ip] = limiter
	return limiter
}

func (rl *IPRateLimiter) GetLimiter(ip string) *rate.Limiter {
	rl.mu.Lock()
	limiter, exists := rl.ips[ip]
	if !exists {
		rl.mu.Unlock()
		return rl.AddIP(ip)
	}
	rl.mu.Unlock()
	return limiter
}

func RateLimitMiddleware(rateCount int) gin.HandlerFunc {
	limiter := NewIPRateLimiter(rate.Every(time.Minute), rateCount)

	return func(c *gin.Context) {
		ip := c.ClientIP()
		if !limiter.GetLimiter(ip).Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error":       "rate limit exceeded",
				"retry_after": "60s",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
