package main

import (
	"net/http"

	tokenbucketRateLimiter "github.com/sebport0/ratelims/tokenbucket"
)

func tokenBucketMiddleware(next http.Handler) http.Handler {
	rateLimiter := tokenbucketRateLimiter.NewRateLimiter(
		1, 10, &tokenbucketRateLimiter.StdTimer{},
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !rateLimiter.IsAllowed() {
			tooManyRequests(w)
			return
		}
		next.ServeHTTP(w, r)
	})
}
