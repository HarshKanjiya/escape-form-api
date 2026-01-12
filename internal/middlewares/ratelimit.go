package middlewares

import (
	"sync"
	"time"

	"github.com/HarshKanjiya/escape-form-api/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

// RateLimiterConfig
type RateLimiterConfig struct {
	Max        int           // Maximum number of requests
	Expiration time.Duration // Time window
}

type ipRateLimiter struct {
	mu      sync.Mutex
	clients map[string]*clientLimiter
}

type clientLimiter struct {
	count      int
	lastReset  time.Time
	expiration time.Duration
	max        int
}

var limiter = &ipRateLimiter{
	clients: make(map[string]*clientLimiter),
}

// RateLimiter creates a rate limiting middleware
func RateLimiter(config RateLimiterConfig) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ip := c.IP()

		limiter.mu.Lock()
		client, exists := limiter.clients[ip]

		now := time.Now()

		if !exists {
			// Create new client limiter
			client = &clientLimiter{
				count:      1,
				lastReset:  now,
				expiration: config.Expiration,
				max:        config.Max,
			}
			limiter.clients[ip] = client
		} else {
			// Check if we should reset the counter
			if now.Sub(client.lastReset) > client.expiration {
				client.count = 1
				client.lastReset = now
			} else {
				client.count++
			}
		}

		// Check if limit exceeded
		if client.count > client.max {
			limiter.mu.Unlock()
			return utils.Error(c, fiber.StatusTooManyRequests, "Rate limit exceeded. Please try again later.")
		}

		limiter.mu.Unlock()

		// Add rate limit headers
		c.Set("X-RateLimit-Limit", string(rune(config.Max)))
		c.Set("X-RateLimit-Remaining", string(rune(config.Max-client.count)))

		return c.Next()
	}
}

// Periodically remove expired entries from the rate limiter
func CleanupRateLimiter() {
	ticker := time.NewTicker(10 * time.Minute)
	go func() {
		for range ticker.C {
			limiter.mu.Lock()
			now := time.Now()
			for ip, client := range limiter.clients {
				if now.Sub(client.lastReset) > client.expiration {
					delete(limiter.clients, ip)
				}
			}
			limiter.mu.Unlock()
		}
	}()
}
