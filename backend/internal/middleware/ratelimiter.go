package middleware

import (
    "net/http"
    "github.com/juju/ratelimit"
    "time"
)

// RateLimiter represents the rate limiter middleware
type RateLimiter struct {
    limiter *ratelimit.Bucket
}

// NewRateLimiter creates a new RateLimiter middleware
func NewRateLimiter(ratePerSec, burst int) *RateLimiter {
    return &RateLimiter{
        limiter: ratelimit.NewBucketWithQuantum(1*time.Second, int64(ratePerSec), int64(burst)),
    }
}

// Middleware is the middleware function to rate limit requests
func (rl *RateLimiter) Middleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if rl.limiter.TakeAvailable(1) == 0 {
            http.Error(w, "Too many requests, please try again later", http.StatusTooManyRequests)
            return
        }
        next.ServeHTTP(w, r)
    })
}
