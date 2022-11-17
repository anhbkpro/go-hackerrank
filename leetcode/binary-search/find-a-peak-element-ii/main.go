package find_a_peak_element_ii

func FindPeakGrid(mat [][]int) []int {
	startCol, endCol := 0, len(mat[0])-1
	for startCol <= endCol {
		maxRow, midCol := 0, (startCol+endCol)>>1
		for row := 0; row < len(mat); row++ {
			if mat[row][midCol] > mat[maxRow][midCol] {
				maxRow = row
			}
		}
		leftIsBig := midCol-1 >= startCol && mat[maxRow][midCol-1] > mat[maxRow][midCol]
		rightIsBig := midCol+1 <= endCol && mat[maxRow][midCol+1] > mat[maxRow][midCol]
		if !leftIsBig && !rightIsBig {
			return []int{maxRow, midCol}
		} else if rightIsBig {
			startCol = midCol + 1
		} else {
			endCol = midCol - 1
		}
	}
	return []int{-1, -1}
}
