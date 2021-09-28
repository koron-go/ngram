/*
Package ngram provides N-gram index.
*/
package ngram

// Index provides N-gram index.
// Its key is keyword in N chars, and value is count of keyword.
type Index map[string]int

// New creates N-gram index on given string.
func New(n int, s string) Index {
	x := Index{}
	x.Add(n, s)
	return x
}

// Add adds N-chars (n) keywords from s.
func (x Index) Add(n int, s string) {
	buf := newBuffer(n)
	for i := range s {
		if b := buf.get(); b >= 0 {
			x[s[b:i]]++
		}
		buf.put(i)
	}
	// put last a N-gram.
	if b := buf.get(); b >= 0 {
		x[s[b:]]++
	}
}

// buffer provides FIFO buffer for byte-index on string.
type buffer struct {
	data []int
	idx  int
}

func newBuffer(n int) buffer {
	data := make([]int, n)
	for i := 0; i < n; i++ {
		data[i] = -1
	}
	return buffer{
		data: data,
		idx:  0,
	}
}

func (b *buffer) get() int {
	return b.data[b.idx]
}

func (b *buffer) put(v int) {
	b.data[b.idx] = v
	b.idx = (b.idx + 1) % len(b.data)
}
