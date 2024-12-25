package bucket

// Bucket is a container of tokens.
type Bucket struct {
	// Capacity is the maximum amount of tokens
	// the container can have.
	Capacity uint64

	// Tokens is the current amount of tokens
	// that the bucket has. It can be 0.
	Tokens uint64
}

func (b *Bucket) Add(tokens uint64) {
	if b.Tokens+tokens <= b.Capacity {
		b.Tokens += tokens
	} else {
		b.Tokens = b.Capacity
	}
}

func (b *Bucket) Sub() {
	if b.Tokens > 0 {
		b.Tokens -= 1
	}
}

func (b *Bucket) IsFull() bool {
	return b.Tokens == b.Capacity
}

func (b *Bucket) IsEmpty() bool {
	return b.Tokens == 0
}

func NewBucket(capacity uint64) *Bucket {
	return &Bucket{Capacity: capacity, Tokens: capacity}
}
