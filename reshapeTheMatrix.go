func matrixReshape(mat [][]int, r int, c int) [][]int {
	if r*c != len(mat)*len(mat[0]) {
		return mat
	}
	if len(mat) == 0 {
		return mat
	}
	if len(mat[0]) == 0 {
		return mat
	}
	rc := 0
	cc := 0
	nMat := make([][]int, r)
	for i := 0; i < r; i++ {
		nMat[i] = make([]int, c)
	}
	for _, m := range mat {
		for _, n := range m {
			if cc == c {
				if rc == r {
					break
				}
				rc++
				cc = 0
			}
			nMat[rc][cc] = n
			cc++
		}
	}
	return nMat
}

