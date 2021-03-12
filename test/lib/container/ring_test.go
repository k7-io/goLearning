package container



import (
	"container/ring"
	"testing"
)

func TestRingLen(t *testing.T) {
	// Create a new ring of size 4
	r := ring.New(4)
	t.Log(r.Len())
}

func TestRingIter(t *testing.T)  {
	r := ring.New(4)
	t.Log(r.Len())
	n := r.Len()
	// Initialize the ring with some integer values
	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next()
	}

	for j := 0; j < 2*n; j++ {
		t.Logf("value:%v\n", r.Value)
		r = r.Next()
	}
}

func TestRingReIter(t *testing.T)  {
	r := ring.New(4)
	t.Log(r.Len())
	n := r.Len()
	// Initialize the ring with some integer values
	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next()
	}

	for j := 0; j < n; j++ {
		r = r.Prev()
		t.Logf("value:%v\n", r.Value)
	}
}