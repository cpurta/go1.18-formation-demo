package fuzzing

import (
	"testing"
	"unicode/utf8"

	"github.com/stretchr/testify/require"
)

var input = "The quick brown fox jumped over the lazy dog"

func TestReverse(t *testing.T) {
	reversed := Reverse(input)

	require.Equal(t, "god yzal eht revo depmuj xof nworb kciuq ehT", reversed)
}

func FuzzReverse(f *testing.F) {
	f.Fuzz(func(t *testing.T, orig string) {
		rev := Reverse(orig)
		doubleRev := Reverse(rev)
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}
