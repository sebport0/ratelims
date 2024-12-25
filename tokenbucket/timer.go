package ratelimiter

import "time"

type StdTimer struct{}

func (t *StdTimer) Now() time.Time {
	return time.Now()
}
