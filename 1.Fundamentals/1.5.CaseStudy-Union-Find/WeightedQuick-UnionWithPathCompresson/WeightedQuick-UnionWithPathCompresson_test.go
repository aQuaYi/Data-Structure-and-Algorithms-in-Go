package wquwpc

import (
	"testing"
)

var u *uf
var U UnionFind

var input = [][]int{
	[]int{4, 3},
	[]int{3, 8},
	[]int{6, 5},
	[]int{9, 4},
	[]int{2, 1},
	[]int{8, 9},
	[]int{5, 0},
	[]int{7, 2},
	[]int{6, 1},
	[]int{1, 0},
	[]int{6, 7},
}

var connectedOne = []int{3, 4, 8, 9}
var connectedTwo = []int{0, 1, 2, 5, 6, 7}

func Test_New(t *testing.T) {
	n := 100
	u = New(n).(*uf)
	if u.count != n {
		t.Error("新生成实例的count出错")
	}
	for i := 0; i < n; i++ {
		if u.id[i] != i || u.sz[i] != 1 {
			t.Error("id或者sz的初始化错误。")
		}
	}
}

func Test_UnionFind_Count(t *testing.T) {
	n := 10
	U = New(n)
	if U.Count() != n {
		t.Error("Count出错")
	}
}

func Test_UnionFind_Find(t *testing.T) {
	n := 100
	u = New(n).(*uf)
	for i := 0; i < n-1; i++ {
		u.id[i] = i + 1
	}

	for i := 0; i < n; i++ {
		if u.Find(i) != n-1 {
			t.Errorf("u.Find(%d)!=%d\n", i, n-1)
		}
	}

	for i := 0; i < n; i++ {
		if u.id[i] != n-1 {
			t.Error("Find的过程中，没能压缩路径\n")
		}
	}
}
func Test_UnionFind_Connected(t *testing.T) {
	n := 100
	u = New(n).(*uf)
	for i := 0; i < n-1; i++ {
		u.id[i] = i + 1
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if !u.Connected(i, j) {
				t.Error("全连通的节点被判定为不连通。")
			}
		}
	}
}

func Test_UnionFind_Union(t *testing.T) {
	n := 10
	U = New(n)
	for _, v := range input {
		U.Union(v[0], v[1])
	}

	for _, v := range connectedOne {
		for _, k := range connectedOne {
			if !U.Connected(v, k) {
				t.Error(v, k, "连通区域的点被判定为不连通")
			}
		}
	}

	for _, v := range connectedTwo {
		for _, k := range connectedTwo {
			if !U.Connected(v, k) {
				t.Error(v, k, "连通区域的点被判定为不连通")
			}
		}
	}

	for _, v := range connectedOne {
		for _, k := range connectedTwo {
			if U.Connected(v, k) {
				t.Error(v, k, "不连通区域的点被判定为不连通")
			}
		}
	}
}
