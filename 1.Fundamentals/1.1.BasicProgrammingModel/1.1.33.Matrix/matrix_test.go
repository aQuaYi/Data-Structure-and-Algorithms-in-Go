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

	amb = [][]float64{
		[]float64{11., 6.},
		[]float64{15., 12.},
	}
)

func Test_Dot(t *testing.T) {
	if Dot(x, y) != 24.0 {
		t.Error("Dot()无法通过测试")
	}
}

func Test_Transpose(t *testing.T) {
	aTrans := Transpose(a)
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			if aTrans[i][j] != at[i][j] {
				t.Error("Transpose()无法通过测试。")
			}
		}
	}
}

/*
func Test_Mult(t *testing.T) {
	ab := Mult(a, b)
	for i := 0; i < len(ab); i++ {
		for j := 0; j < len(ab[i]); j++ {
			if ab[i][j] != amb[i][j] {
				t.Error("Mult()无法通过测试")
			}
		}
	}
}
*/
