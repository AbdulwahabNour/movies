package middlewares

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"

	"golang.org/x/time/rate"
)

// RateLimitMiddleware is a middleware function that implements rate limiting
// for incoming requests.
//
// It uses a sync.Map to keep track of clients and their respective rate
// limiters. Each client is associated with an IP address. If a client exceeds
// the rate limit, the middleware will abort the request and return a 429 Too
// Many Requests error.
func (m *MiddleWares) RateLimitMiddleware() gin.HandlerFunc {

	type client struct {
		limiter  *rate.Limiter
		lastSeen time.Time
	}
	var clients sync.Map

	go func() {
		ticker := time.NewTicker(time.Minute)
		defer ticker.Stop()

		for range ticker.C {

			clients.Range(func(key, value interface{}) bool {
				c, ok := value.(*client)
				if !ok {
					return true
				}
				if time.Since(c.lastSeen).Minutes() >= 2.8 {
					clients.Delete(key)
				}

				return true
			})

		}
	}()

	return func(ctx *gin.Context) {
		ip := ctx.ClientIP()

		if limiter, ok := clients.LoadOrStore(ip, &client{limiter: rate.NewLimiter(2, 2), lastSeen: time.Now()}); ok {

			if !limiter.(*client).limiter.Allow() {
				ctx.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
					"error": "Too many requests",
				})
				return
			}

		}

		ctx.Next()
	}
}
