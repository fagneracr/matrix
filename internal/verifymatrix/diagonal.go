package verifymatrix

/*FindDiagonais - Organize array in diagonal
  imput: array string
  output array organized in diogonal, more than 2 elements*/
func FindDiagonais(in []string) (mtdiagonal []string) {
	var mt [][]string
	coll := 0
	for _, i := range in {
		var array []string
		for _, l := range i {
			array = append(array, string(l))
		}
		if len(i) > coll {
			coll = len(i)
		}
		mt = append(mt, array)

	}
	r := len(in)
	for k := 0; k < coll; k++ {
		i := coll - 1
		j := k + 1
		var diagonal string
		for isValid(i, j, r, coll) {
			diagonal = diagonal + mt[i][j]
			i--
			j++
		}
		if len(diagonal) >= 3 {
			mtdiagonal = append(mtdiagonal, diagonal)
		}

	}
	for k := 0; k < r; k++ {
		i := k
		j := 0
		var diagonal string
		for isValid(i, j, r, coll) {
			diagonal = diagonal + mt[i][j]
			i++
			j++
		}
		if len(diagonal) > 3 {
			mtdiagonal = append(mtdiagonal, diagonal)
		}

	}
	return mtdiagonal

}
