package middleware

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/time/rate"
)

var limiter = rate.NewLimiter(10, 1) // 10 requests per second

func RateLimitMiddleware(c *fiber.Ctx) error {
	if !limiter.Allow() {
		return c.Status(429).SendString("Too many requests, please try again later.")
	}

	return c.Next()
}
