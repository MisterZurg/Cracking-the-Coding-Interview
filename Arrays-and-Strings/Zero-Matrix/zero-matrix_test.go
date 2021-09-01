package Zero_Matrix

import "testing"

//[[1,1,1],[1,0,1],[1,1,1]]
// [[1,0,1],[0,0,0],[1,0,1]]

var tests = []struct {
	matrix   [][]int
	zerofied [][]int
}{
	{
		[][]int{{1,1,1},{1,0,1},{1,1,1}},
		[][]int{{1,0,1},{0,0,0},{1,0,1}},
	},
}

func TestZerofiedMatrix(t *testing.T) {
	for _, test := range tests {
		want := test.zerofied
		got := ZerofiedMatrix(test.matrix)
		for i:= 0; i< len(want); i++ {
			for j:= 0;j<len(want[0]);j++{
				if want[i][j] != got[i][j] {
					t.Fatalf("Test failed got %v expected %v", got, want)
				}
			}
		}
	}
}
