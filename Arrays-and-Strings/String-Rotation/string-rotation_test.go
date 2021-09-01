package String_Rotation

import "testing"

func TestIsRotation(t *testing.T) {
	want := true
	s1 := "waterbottle"
	s2 := "erbottlewat"
	if want != IsRotation(s1, s2) {
		t.Fatalf("Fail")
	}
}
