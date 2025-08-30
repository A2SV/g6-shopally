package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type ipClient struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

type IPRateLimiter struct {
	limit   rate.Limit
	burst   int
	ttl     time.Duration
	mu      sync.Mutex
	clients map[string]*ipClient
}

func NewIPRateLimiter(limit rate.Limit, burst int, ttl time.Duration) *IPRateLimiter {
	rl := &IPRateLimiter{
		limit:   limit,
		burst:   burst,
		ttl:     ttl,
		clients: make(map[string]*ipClient),
	}
	go rl.cleanupLoop()
	return rl
}

func (r *IPRateLimiter) get(ip string) *rate.Limiter {
	r.mu.Lock()
	defer r.mu.Unlock()

	if c, ok := r.clients[ip]; ok {
		c.lastSeen = time.Now()
		return c.limiter
	}
	lim := rate.NewLimiter(r.limit, r.burst)
	r.clients[ip] = &ipClient{limiter: lim, lastSeen: time.Now()}
	return lim
}

func (r *IPRateLimiter) cleanupLoop() {
	t := time.NewTicker(r.ttl)
	defer t.Stop()
	for range t.C {
		cutoff := time.Now().Add(-r.ttl)
		r.mu.Lock()
		for ip, c := range r.clients {
			if c.lastSeen.Before(cutoff) {
				delete(r.clients, ip)
			}
		}
		r.mu.Unlock()
	}
}

func (r *IPRateLimiter) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		if ip == "" {
			ip = "unknown"
		}
		lim := r.get(ip)
		if !lim.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "too many requests",
			})
			return
		}
		c.Next()
	}
}
