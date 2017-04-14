package matrix

import (
	"testing"
)

var (
	x = []float64{1.0, 2.0, 3.0}
	y = []float64{3.0, 0.0, 7.0}

	a = [][]float64{
		[]float64{1., 0., 2.},
		[]float64{0., 1., 3.},
	}

	ax = []float64{7., 11.}

	at = [][]float64{
		[]float64{1., 0.},
		[]float64{0., 1.},
		[]float64{2., 3.},
	}

	b = [][]float64{
		[]float64{1., 0.},
		[]float64{0., 3.},
		[]float64{5., 3.},
	}
	yb = []float64{38., 21.}

	amb = [][]float64{
		[]float64{11., 6.},
		[]float64{15., 12.},
	}
)

func Test_Dot(t *testing.T) {
	newD, _ := Dot(x, y)
	if newD != 24.0 {
		t.Error("Dot()无法通过测试")
	}
}

func Test_Transpose(t *testing.T) {
	aTrans := Transpose(a)
	for i := 0; i < len(at); i++ {
		for j := 0; j < len(at[i]); j++ {
			if aTrans[i][j] != at[i][j] {
				t.Error("Transpose()无法通过测试。")
			}
		}
	}
}

func Test_Mult(t *testing.T) {
	ab, _ := Mult(a, b)
	for i := 0; i < len(ab); i++ {
		for j := 0; j < len(ab[i]); j++ {
			if ab[i][j] != amb[i][j] {
				t.Error("Mult()无法通过测试")
			}
		}
	}
}

func Test_MultArray(t *testing.T) {
	newAX, _ := MultArray(a, x)
	for i := 0; i < len(newAX); i++ {
		if newAX[i] != ax[i] {
			t.Error("MatrixMultArray()无法通过测试。")
		}
	}
}

func Test_ArrayMultMatrix(t *testing.T) {
	newYB, _ := ArrayMultMatrix(y, b)
	for i := 0; i < len(yb); i++ {
		if newYB[i] != yb[i] {
			t.Error("ArrayMultMatrix()无法通过测试")
		}
	}
}
