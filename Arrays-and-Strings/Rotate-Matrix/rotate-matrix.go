/*
Rotate Matrix:
Given an image represented by an NxN matrix,
where each pixel in the image is 4 bytes,
write a method to rotate the image by 90 degrees.
Can you do this in place?
*/
package Rotate_Matrix

// 1 2 3      7 4 1
// 4 5 6  ->  8 5 2
// 7 8 9      9 6 3
func rotateMatrixBy90Degrees(matrix [][]int) [][]int {
	n := len(matrix)
	// Transpose matrix
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	// Reverse elements
	for i := 0; i < n; i++ {
		for j := 0; j < n / 2; j++ {
			matrix[i][j], matrix[i][n - 1 -j] = matrix[i][n - 1 -j], matrix[i][j]
		}
	}
	return matrix
}
