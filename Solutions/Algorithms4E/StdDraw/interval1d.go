package StdDraw

//Interval1D 是一个直线段
type Interval1D struct {
	lo, hi float64
}

//NewInterval1D 返回*Interval1D类型的值
func NewInterval1D(lo, hi float64) *Interval1D {
	if lo > hi {
		lo, hi = hi, lo
	}
	return &Interval1D{
		lo: lo,
		hi: hi,
	}
}

//Length 返回Interval1D的长度
func (it *Interval1D) Length() float64 {
	return it.hi - it.lo
}

//Contains 返回Interval1D是否包含x值
func (it *Interval1D) Contains(x float64) bool {
	if it.lo < x && x < it.hi {
		return true
	}
	return false
}

//Intersect 返回两个Interval1D实例是否相交
func (it *Interval1D) Intersect(that *Interval1D) bool {
	if it.Contains(that.lo) || it.Contains(that.hi) {
		return true
	}
	return false
}

/* 由于Interval1D只是直线上的间隔，信息不全，无法在2D图像是绘制出来。
func (it *Interval1D) Draw(dc *gg.Context, r, g, b, a, w float64) {
}
*/
