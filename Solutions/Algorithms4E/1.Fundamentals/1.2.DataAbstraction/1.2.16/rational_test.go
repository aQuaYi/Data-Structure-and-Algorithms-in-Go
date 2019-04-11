package rational

import (
	"testing"
)

var r = Rational{
	n: 2,
	d: 3,
}

var a = Rational{
	n: 3,
	d: 7,
}

var rpa = Rational{
	n: 23,
	d: 21,
}
var rma = Rational{
	n: 5,
	d: 21,
}

var rta = Rational{
	n: 2,
	d: 7,
}

var rda = Rational{
	n: 14,
	d: 9,
}

var zero = Rational{
	n: 0,
	d: 1,
}

func Test_Rational_Equals(t *testing.T) {
	if !r.Equals(r) {
		t.Errorf("Equals()判定%v和%v不相等", r, r)
	}
	if r.Equals(zero) {
		t.Errorf("Equals()判定%v和%v相等", r, zero)
	}
}

func Test_NewRational(t *testing.T) {
	if _, err := NewRational(1, 0); err.Error() != "NewRational()中使用的d为0。" {
		t.Error("NewRational()没能正确处理分母为0的情况。")
	}

	if tr, _ := NewRational(2, 3); !tr.Equals(r) {
		t.Error("把2/3和2/3判定为不相等。")
	}

	if tr, _ := NewRational(4, 6); !tr.Equals(r) {
		t.Error("把4/6和2/3判定为不相等。")
	}

	if r.Equals(a) {
		t.Errorf("把%v和%v判定为相等", t, a)
	}
}

func Test_Rational_Plus(t *testing.T) {
	if !r.Plus(a).Equals(rpa) {
		t.Errorf("%v+%v!=%v", r, a, rpa)
	}
}

func Test_Rational_Minus(t *testing.T) {
	if !r.Minus(a).Equals(rma) {
		t.Errorf("%v-%v!=%v", r, a, rma)
	}
}

func Test_Rational_Times(t *testing.T) {
	if !r.Times(a).Equals(rta) {
		t.Errorf("%v*%v!=%v", r, a, rta)
	}
}
func Test_Rational_Divides(t *testing.T) {
	if _, err := r.Divides(zero); err.Error() != "被除数为0" {
		t.Error("Divides()没有正确处理除以0的情况。")
	}

	trda, err := r.Divides(a)
	if !trda.Equals(rda) || err != nil {
		t.Errorf("%v÷%v!=%v", r, a, rda)
	}
}
