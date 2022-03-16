package matrix

// If needed, you may define helper functions here.

// AreAdjacent returns true iff a and b are adjacent in lst.
func AreAdjacent(lst []int, a, b int) bool {
	var l int = len(lst)
	if lst == nil {
		return false
	}

	for i := 0; i < l-1; i++ {
		if lst[i] == a && lst[i+1] == b {
			return true
		}
		if lst[i] == b && lst[i+1] == a {
			return true
		}
	}
	return false
}

// Transpose returns the transpose of the 2D matrix mat.
func Transpose(mat [][]int) [][]int {
	if len(mat) == 0 {
		return mat
	}

	ro := len(mat[0])
	co := len(mat)

	t_mat := make([][]int, ro)
	for i := range t_mat {
		t_mat[i] = make([]int, co)
	}

	for i := range t_mat {
		for j := range t_mat[0] {
			t_mat[i][j] = mat[j][i]
		}
	}
	return t_mat

}

// AreNeighbors returns true iff a and b are neighbors in the 2D matrix mat.
func AreNeighbors(mat [][]int, a, b int) bool {
	var l int = len(mat)
	for i := 0; i < l; i++ {
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == a {
				if i-1 >= 0 && mat[i-1][j] == b {
					return true
				} else if i+1 < len(mat) && mat[i+1][j] == b {
					return true
				} else if j-1 >= 0 && mat[i][j-1] == b {
					return true
				} else if j+1 < len(mat[i]) && mat[i][j+1] == b {
					return true
				}
			}
		}
	}
	return false
}
