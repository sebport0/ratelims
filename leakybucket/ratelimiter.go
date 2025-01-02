package ratelimiter

import (
	"errors"
	"net/http"
)

type RateLimiter struct{}

func NewRateLimiter(bucketSize int, rate int) (*RateLimiter, error) {
	if bucketSize < 0 {
		return nil, errors.New("bucket size must be >= 0")
	}
	if rate < 1 {
		return nil, errors.New("rate must be >= 1")
	}
	return &RateLimiter{}, nil
}

func (rl *RateLimiter) IsAllowed(r *http.Request) bool {
	return false
}
