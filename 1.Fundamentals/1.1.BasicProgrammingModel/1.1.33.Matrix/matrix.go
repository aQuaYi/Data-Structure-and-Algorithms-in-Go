package matrix

func Dot(x, y []float64) float64 {
	result := 0.
	for i := 0; i < len(x); i++ {
		result += x[i] * y[i]
	}
	return result
}

func Transpose(a [][]float64) [][]float64 {
	result := [][]float64{}
	for i := 0; i < len(a[0]); i++ {
		result = append(result, make([]float64, len(a)))
		for j := 0; j < len(a); j++ {
			result[i][j] = a[j][i]
		}
	}
	return result
}

func Mult(a, b [][]float64) [][]float64 {
	result := [][]float64{}
	b = Transpose(b)
	for i := 0; i < len(a); i++ {
		result = append(result, make([]float64, len(b)))
		for j := 0; j < len(b); j++ {
			result[i][j] = Dot(a[i], b[j])
		}
	}
	return result
}

func MatrixMultArray(a [][]float64, x []float64) []float64 {
	result := make([]float64, len(a))
	for i := 0; i < len(a); i++ {
		result[i] = Dot(a[i], x)
	}
	return result
}

func ArrayMultMatrix(y []float64, a [][]float64) []float64 {
	a = Transpose(a)
	return MatrixMultArray(a, y)
}
