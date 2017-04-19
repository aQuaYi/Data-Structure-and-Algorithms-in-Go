package StdDraw

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
