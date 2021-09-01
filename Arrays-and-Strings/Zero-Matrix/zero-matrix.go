package Zero_Matrix

type zero struct {
	row, col int
}

func ZerofiedMatrix(matrix [][]int) [][]int{
	var zeroes []zero
	n := len(matrix)
	m := len(matrix[0])
	// Memorize [i][j] with zeros
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 0 {
				zeroes = append(zeroes, zero{
					row: i,
					col: j,
				})
			}
		}
	}
	// Fill da zeroes
	for _, zeroElem := range zeroes {
		for i := 0; i < n; i++ {
			matrix[i][zeroElem.col] = 0
		}

		for i := 0; i < m; i++ {
			matrix[zeroElem.row][i] = 0
		}
	}
	return matrix
}
