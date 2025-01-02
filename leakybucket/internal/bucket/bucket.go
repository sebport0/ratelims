package bucket

import (
	"errors"
	"time"

	csqueue "github.com/golang-collections/collections/queue"
)

type Bucket struct {
	queue    *csqueue.Queue
	capacity int
	LastLeak time.Time
}

func (b *Bucket) Add(elem any) error {
	if b.IsFull() {
		return errors.New("can't add new elements, bucket is full")
	}
	b.queue.Enqueue(elem)
	return nil
}

func (b *Bucket) IsFull() bool {
	return b.queue.Len() == b.capacity
}

func (b *Bucket) IsEmpty() bool {
	return b.queue.Len() == 0
}

func (b *Bucket) Len() int {
	return b.queue.Len()
}

func NewBucket(capacity int) (*Bucket, error) {
	if capacity < 0 {
		return nil, errors.New("bucket's capacity must be greater or equal to zero")
	}
	q := csqueue.New()
	return &Bucket{queue: q, capacity: capacity}, nil
}
