package ngram_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/koron-go/ngram"
)

func TestNew(t *testing.T) {
	for i, tc := range []struct {
		n   int
		s   string
		exp ngram.Index
	}{
		{2, "abcde", ngram.Index{
			"ab": 1,
			"bc": 1,
			"cd": 1,
			"de": 1,
		}},

		{3, "abcde", ngram.Index{
			"abc": 1,
			"bcd": 1,
			"cde": 1,
		}},

		{2, "こんにちは", ngram.Index{
			"こん": 1,
			"んに": 1,
			"にち": 1,
			"ちは": 1,
		}},
	} {
		x := ngram.New(tc.n, tc.s)
		if d := cmp.Diff(tc.exp, x); d != "" {
			t.Errorf("mismatch: i=%d n=%d s=%q: -want +got\n%s", i, tc.n, tc.s, d)
		}
	}
}

func TestAdd(t *testing.T) {
	for i, tc := range []struct {
		n   int
		s   string
		exp ngram.Index
	}{
		{2, "abcde", ngram.Index{
			"ab": 1,
			"bc": 1,
			"cd": 1,
			"de": 1,
		}},

		{3, "abcde", ngram.Index{
			"abc": 1,
			"bcd": 1,
			"cde": 1,
		}},

		{2, "こんにちは", ngram.Index{
			"こん": 1,
			"んに": 1,
			"にち": 1,
			"ちは": 1,
		}},
	} {
		x := ngram.Index{}
		x.Add(tc.n, tc.s)
		if d := cmp.Diff(tc.exp, x); d != "" {
			t.Errorf("mismatch: i=%d n=%d s=%q: -want +got\n%s", i, tc.n, tc.s, d)
		}
	}
}
