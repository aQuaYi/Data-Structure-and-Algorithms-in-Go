package StdDraw

import (
	"math"

	"github.com/fogleman/gg"
)

//Point2D 表示二维上的点
type Point2D struct {
	x, y float64
}

//NewPoint2D 用来生成Point2D点
func NewPoint2D(x, y float64) *Point2D {
	return &Point2D{
		x: x,
		y: y,
	}
}

//X 返回Point2D的X坐标值
func (p *Point2D) X() float64 {
	return p.x
}

//Y 返回Point2D的Y坐标值
func (p *Point2D) Y() float64 {
	return p.y
}

//R 返回Point2D的极坐标R值
func (p *Point2D) R() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

//Theta 返回Point2D的极坐标角度值
func (p *Point2D) Theta() float64 {
	if p.x == 0 && p.y == 0 {
		return 0
	}
	switch {
	case p.x > 0 && p.y >= 0:
		return math.Atan(p.y / p.x)
	case p.x == 0 && p.y > 0:
		return math.Pi / 2
	case p.x < 0:
		return math.Atan(p.y/p.x) + math.Pi
	case p.x == 0 && p.y < 0:
		return math.Pi * 3 / 2
	default:
		return math.Atan(p.y/p.x) + math.Pi*2
	}
}

//DistTo 返回p和that点之间的距离值
func (p *Point2D) DistTo(that *Point2D) float64 {
	x := p.x - that.x
	y := p.y - that.y
	return math.Sqrt(x*x + y*y)
}

//Draw 在指定的画布上绘制点。
func (p *Point2D) Draw(dc *gg.Context, r, g, b, a, radius float64) {
	dc.SetRGBA(r, g, b, a)
	dc.DrawPoint(p.x, p.y, radius)
	dc.Fill()
}
