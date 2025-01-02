package bucket

import (
	"net/http"
	"testing"
)

func TestBucketAddsRequestsAndCorrectState(t *testing.T) {
	testCases := []struct {
		desc     string
		capacity int
		Nreqs    int
		isFull   bool
		isEmpty  bool
	}{
		{"one request", 5, 1, false, false},
		{"full", 5, 5, true, false},
		{"empty", 1, 0, false, true},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			b, err := NewBucket(tC.capacity)
			if err != nil {
				t.Fatal(err)
			}
			for i := range tC.Nreqs {
				r, err := http.NewRequest(http.MethodGet, "http://0.0.0.0:8000", nil)
				if err != nil {
					t.Fatalf("request #%d creation failed: %v", i, err)
				}
				err = b.Add(r)
				if err != nil {
					t.Fatal(err)
				}
			}

			if tC.isFull != b.IsFull() {
				t.Fatal("bucket shouldn't be full")
			}

			if tC.isEmpty != b.IsEmpty() {
				t.Fatal("bucket shouldn't be empty")
			}
			if tC.Nreqs != b.Len() {
				t.Fatalf("got %v, expected 1\n", tC.Nreqs)
			}
		})
	}
}

func TestBucketLeaks(t *testing.T) {
	testCases := []struct {
		desc     string
		capacity int
		Nreqs    int
		isFull   bool
		isEmpty  bool
	}{
		{
			desc: "",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
		})
	}
}
