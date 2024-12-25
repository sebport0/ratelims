package refiller

import (
	"math"
	"time"

	"github.com/sebport0/ratelims/tokenbucket/internal/bucket"
	"github.com/sebport0/ratelims/tokenbucket/internal/timer"
)

// Refiller controls the addition of new tokens to the bucket.
// Every second, add N tokens.
//
// Tokens to add to the bucket on every tick.
//
// lastRefillEvent tells when was the last time a refill action
// was triggered.
//
// timer is a convenience dependency to have more control over
// the refill events. In practice, most of the times it should
// be a time stdlib wrapper.
type Refiller struct {
	Tokens          uint64
	lastRefillEvent time.Time
	timer           timer.Timer
}

// Refill check the last time a refill was performed and adds
// tokens * seconds since last refill or enough tokens to fill
// the bucket.
func (r *Refiller) Refill(b *bucket.Bucket) {
	now := r.timer.Now()
	timeElapsedSinceLastRefill := now.Sub(r.lastRefillEvent)
	tokensToAdd := float64(r.Tokens * uint64(timeElapsedSinceLastRefill.Seconds()))
	bucketCapacity := float64(b.Capacity)
	b.Add(uint64(math.Min(tokensToAdd, bucketCapacity)))
	r.lastRefillEvent = now
}

func (r *Refiller) LastRefillEvent() time.Time {
	return r.lastRefillEvent
}

func NewRefiller(tokens uint64, timer timer.Timer) *Refiller {
	return &Refiller{
		Tokens:          tokens,
		lastRefillEvent: timer.Now(),
		timer:           timer,
	}
}
