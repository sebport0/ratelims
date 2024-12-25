// Package ratelimiter provides a naive token bucket algorithm
// implementation to perform rate limiting.
package ratelimiter

import (
	"github.com/sebport0/ratelims/tokenbucket/internal/bucket"
	"github.com/sebport0/ratelims/tokenbucket/internal/refiller"
	"github.com/sebport0/ratelims/tokenbucket/internal/timer"
)

// RateLimiter checks that there are enough tokens
// in its internal bucket to allow new requests
// to pass through. When tokens reachs zero, no more
// requests are allowed to pass until the next refill.
type RateLimiter struct {
	// bucket acts as the internal container for tokens.
	// A bucket has a maximum capacity and can be empty.
	bucket *bucket.Bucket

	// refiller handles the addition of new tokens in
	// the bucket. It can refill N tokens per second.
	refiller *refiller.Refiller
}

// IsAllowed says if a request can pass through,
// if the bucket has enough token, or if it must
// be denied because the bucket is empty.
//
// Each request that is allowed to pass'll take
// one token from the bucket.
func (rl *RateLimiter) IsAllowed() bool {
	if !rl.bucket.IsFull() {
		rl.refill()
	}
	if rl.bucket.IsEmpty() {
		return false
	}
	rl.bucket.Sub()
	return true
}

// refill instructs the rate limiter's refiller to
// add more tokens to the bucket based on the amount
// of seconds that passed since the last refill.
func (rl *RateLimiter) refill() {
	rl.refiller.Refill(rl.bucket)
}

// NewRateLimiter creates a new rate limiter ready for use.
//
// tokens is the amount of tokens that are going to be refilled
// each second.
//
// capacity is the total amount of tokens the bucket can contain.
//
// timer is just an interface to be able to configure time flow.
// In 99.9% of cases the adequate timer to use is timer.StdTimer
// from the tockenbucket package.
func NewRateLimiter(tokens, capacity uint64, timer timer.Timer) *RateLimiter {
	b := bucket.NewBucket(capacity)
	r := refiller.NewRefiller(tokens, timer)

	return &RateLimiter{bucket: b, refiller: r}
}
