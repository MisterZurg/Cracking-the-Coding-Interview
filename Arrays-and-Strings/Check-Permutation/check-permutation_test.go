package main

import "testing"

var tests = []struct{
	word string
	permutation string
	expected bool
}{
	{"cat ", " tAc", true},
	{"god" ,"dog ",  false},
}

func TestIsPermutation(t *testing.T) {
	for _, test := range tests{
		got := isPermutation(test.word, test.permutation)
		if got != test.expected{
			t.Fatalf("expected: %t, got: %v", test.expected, got)
		}
	}
}
