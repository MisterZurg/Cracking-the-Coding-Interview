package Rotate_Matrix

import "testing"

var tests = []struct {
	testName string
	input    [][]int
	output   [][]int
}{
	{"1x1",
		[][]int{{1}},
		[][]int{{1}}},
	{"2x2",
		[][]int{{1, 2}, {3, 4}},
		[][]int{{3, 1}, {4, 2}}},
	{"3x3",
		[][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9}},
		[][]int{
			{7, 4, 1},
			{8, 5, 2},
			{9, 6, 3}},
	},
	{"4x4",
		[][]int{
			{5, 1, 9, 11},
			{2, 4, 8, 10},
			{13, 3, 6, 7},
			{15, 14, 12, 16}},
		[][]int{
			{15,13,2,5},
			{14,3,4,1},
			{12,6,8,9},
			{16,7,10,11}},
	},
}

func TestRotateMatrix(t *testing.T) {
	for _,test := range tests{
		want := test.output
		got := rotateMatrixBy90Degrees(test.input)
		for i:= 0; i< len(want); i++ {
			for j:= 0;j<len(want[0]);j++{
				if want[i][j] != got[i][j] {
					t.Fatalf("Test %v failed got %v expected %v", test.testName, got, want)
				}
			}
		}
	}
}
