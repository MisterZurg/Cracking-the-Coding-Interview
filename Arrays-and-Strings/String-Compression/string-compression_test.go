package String_Compression

import "testing"

var tests = []struct {
	word       string
	compressed string
}{
	{"aabcccccaaa", "a2b1c5a3"},
}

func TestStringCompression(t *testing.T) {
	for _, test := range tests {
		got := StringCompression(test.word)
		want := test.compressed
		if got != want {
			t.Fatalf("expected: %v, got: %v", want, got)
		}
	}
}
