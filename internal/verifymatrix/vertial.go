package verifymatrix

/*BuildVertical - Return vertial vector arrange
  input: array to organize
  output: array in vertical*/
func BuildVertical(in []string) (matrixv []string) {
	sizematrix := len(in)
	for y := 0; y < len(in[0]); y++ {
		outputv := string(in[0][y])
		for h := 1; h < sizematrix; h++ {
			if len(in[h]) < y {
				break
			}
			outputv = outputv + string(in[h][y])
		}
		matrixv = append(matrixv, outputv)

	}
	return matrixv

}
