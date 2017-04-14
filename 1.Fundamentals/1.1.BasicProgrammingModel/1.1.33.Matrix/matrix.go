package matrix

import (
	"errors"
)

//Dot 是两个向量的乘法
func Dot(x, y []float64) (float64, error) {
	if len(x) != len(y) || len(x) == 0 {
		return 0, errors.New("两个切片的长度为零，或者不一致。")
	}
	result := 0.
	for i := 0; i < len(x); i++ {
		result += x[i] * y[i]
	}
	return result, nil
}

//Transpose 能把矩阵转置
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

//Mult 把两个矩阵相乘
func Mult(a, b [][]float64) ([][]float64, error) {
	var err error
	result := [][]float64{}
	b = Transpose(b)
	for i := 0; i < len(a); i++ {
		result = append(result, make([]float64, len(b)))
		for j := 0; j < len(b); j++ {
			result[i][j], err = Dot(a[i], b[j])
			if err != nil {
				return nil, errors.New("无法完成矩阵的乘积:" + err.Error())
			}
		}
	}
	return result, nil
}

//MultArray 返回矩阵乘以向量的乘积
func MultArray(a [][]float64, x []float64) ([]float64, error) {
	var err error
	result := make([]float64, len(a))
	for i := 0; i < len(a); i++ {
		result[i], err = Dot(a[i], x)
		if err != nil {
			return nil, errors.New("无法利用矩阵乘以向量:" + err.Error())
		}
	}
	return result, nil
}

//ArrayMultMatrix 返回向量乘以矩阵的乘积
func ArrayMultMatrix(y []float64, a [][]float64) ([]float64, error) {
	a = Transpose(a)
	return MultArray(a, y)
}
