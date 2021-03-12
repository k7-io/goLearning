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

func TestRingDo(t *testing.T)  {
	r := ring.New(4)
	t.Log(r.Len())
	n := r.Len()
	// Initialize the ring with some integer values
	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next()
	}

	r.Do(func(i interface{}) {
		t.Logf("i:%v", i)
	})

	rr := ring.New(r.Len())
	r.Do(func(i interface{}) {
		// copy r ele to rr
		rr.Value = i
		rr = rr.Next()
	})

	for i:=0; i<rr.Len(); i++ {
		t.Logf("rr:%v", rr.Value)
		rr = rr.Next()
	}

}

func TestRingMov(t *testing.T)  {
	r := ring.New(4)
	t.Log(r.Len())
	n := r.Len()
	// Initialize the ring with some integer values
	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next()
	}
	t.Logf("first:%v", r.Value)
	r = r.Move(3)
	t.Logf("last:%v", r.Value)
}

func TestRingMovV2(t *testing.T)  {
	r := ring.New(4)
	t.Log(r.Len())
	n := r.Len()
	// Initialize the ring with some integer values
	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next()
	}
	t.Logf("first:%v", r.Value)
	index := 5
	r = r.Move(index%n)
	t.Logf("last:%v", r.Value)
}
// todo: link/unlink