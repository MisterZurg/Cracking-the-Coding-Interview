package One_Way

import "testing"

var tests = []struct {
	word         string
	modification string
	expected     bool
}{
	{"pale", "ple", true},
	{"pales", "pale", true},
	{"pale", "bale", true},
	{"pale", "bae", false},
}

func TestOneWay(t *testing.T) {
	for _, test := range tests {
		got := oneway(test.word, test.modification)
		if got != test.expected{
			t.Fatalf("expected: %t, got: %v", test.expected, got)
		}
	}
}
