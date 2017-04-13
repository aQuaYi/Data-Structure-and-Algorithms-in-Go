package matrix

func Dot(x, y []float64) float64 {
	result := 0.
	for i := 0; i < len(x); i++ {
		result += x[i] * y[i]
	}
	return result
}

func Transpose(m [][]float64) [][]float64 {
	result := [][]float64{}
	for i := 0; i < len(m); i++ {
		result = append(result, make([]float64, len(m)))
		for j := 0; j < len(m[i]); j++ {
			result[i][j] = m[j][i]
		}
	}
	return result
}

func Mult(x, y [][]float64) [][]float64 {
	result := [][]float64{}
	y = Transpose(y)
	for i := 0; i < len(x); i++ {
		result = append(result, make([]float64, len(y)))
		for j := 0; j < len(y); j++ {
			result[i][j] = Dot(x[i], y[j])
		}
	}
	return result
}

func MatrixMultArray(m [][]float64, a []float64) []float64 {

}
