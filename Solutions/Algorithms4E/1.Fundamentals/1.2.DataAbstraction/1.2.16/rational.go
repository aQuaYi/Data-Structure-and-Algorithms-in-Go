package rational

import (
	"errors"
	"fmt"
)

//Rational 是用分数形式表示的有理数
type Rational struct {
	n, d int
}

//NewRational 返回一个Rational类型的实例
func NewRational(n, d int) (Rational, error) {
	if d == 0 {
		return Rational{}, errors.New("NewRational()中使用的d为0。")

	}
	if d < 0 { //为了String()出来的内容好看一些
		d *= -1
		n *= -1
	}
	g := gcd(n, d)
	return Rational{
		n: n / g,
		d: d / g,
	}, nil
}

//Plus 返回两个*Rational相加之和
func (r Rational) Plus(a Rational) Rational {
	n := r.n*a.d + a.n*r.d
	d := r.d * a.d
	result, _ := NewRational(n, d)
	return result
}

//Minus 返回两个*Rational相减之差
func (r Rational) Minus(a Rational) Rational {
	a.n *= -1
	return r.Plus(a)
}

//Times 返回两个Rational的乘积
func (r Rational) Times(a Rational) Rational {
	n := r.n * a.n
	d := r.d * a.d
	result, _ := NewRational(n, d)
	return result
}

//Divides 返回两个Rational的除数
func (r Rational) Divides(a Rational) (Rational, error) {
	if a.n == 0 {
		return Rational{}, errors.New("被除数为0")
	}
	a.n, a.d = a.d, a.n
	return r.Times(a), nil
}

//Equals 判断两个Rational是否相等
func (r Rational) Equals(a Rational) bool {
	if r.n == a.n && r.d == a.d {
		return true
	}
	return false
}

//String 实现了打印接口
func (r Rational) String() string {
	return fmt.Sprintf("%d/%d", r.n, r.d)
}

func gcd(p, q int) int {
	if p < q {
		p, q = q, p
	}
	if q == 0 {
		return p
	}
	return gcd(q, p%q)
}
