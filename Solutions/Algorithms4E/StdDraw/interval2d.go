package StdDraw

import (
	"github.com/fogleman/gg"
)

//Interval2D 是2D平面上的间隔
type Interval2D struct {
	x, y *Interval1D
}

//Area 返回Interval2D所占的面积大小
func (it *Interval2D) Area() float64 {
	return it.x.Length() * it.y.Length()
}

//Contains 返回Interval2D是否返回p
func (it *Interval2D) Contains(p *Point2D) bool {
	if it.x.Contains(p.x) && it.y.Contains(p.y) {
		return true
	}
	return false
}

//Intersect 判断两个Interval2D是否相交
func (it *Interval2D) Intersect(that *Interval2D) bool {
	if it.x.Intersect(that.x) && it.y.Intersect(that.y) {
		return true
	}
	return false
}

//Draw 在一个*gg.Context图像上绘制这个二维图像
func (it *Interval2D) Draw(dc *gg.Context, r, g, b, a float64) {
	dc.SetRGBA(r, g, b, a)
	dc.DrawRectangle(it.x.lo, it.y.lo, it.x.Length(), it.y.Length())
	dc.Fill()
}
