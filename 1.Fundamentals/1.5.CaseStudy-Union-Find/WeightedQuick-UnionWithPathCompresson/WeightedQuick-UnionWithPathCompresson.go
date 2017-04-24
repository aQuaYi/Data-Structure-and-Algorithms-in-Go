//Package wquwpc 是具有路径压缩功能的加权快速连通算法的实现
package wquwpc

//UnionFind 是一个连通发现接口
type UnionFind interface {
	//连接两个节点
	Union(int, int)

	//找到某个节点的根节点
	Find(int) int

	//判断两个节点是否连通
	Connected(int, int) bool

	//返回独立通道的个数
	Count() int
}

type uf struct {
	id    []int
	sz    []int
	count int
}

//New 返回符合UnionFind借口的*uf实例
func New(n int) UnionFind {
	id := make([]int, n)
	sz := make([]int, n)
	for i := 0; i < n; i++ {
		id[i] = i
		sz[i] = 1
	}

	return &uf{
		id:    id,
		sz:    sz,
		count: n,
	}
}

func (u *uf) Count() int {
	return u.count
}

func (u *uf) Find(p int) int {

	pRoot := u.id[p]
	if u.id[pRoot] == pRoot {
		//成立则说明，pRoot为根节点
		//p为根节点或者根节点的子节点，不需要压缩
		return pRoot
	}

	oldP := p
	for p != u.id[p] { //直到P为根节点时，停止
		p = u.id[p]
	}
	for u.id[oldP] != p { //压缩路径上的每一个P点
		pNext := u.id[oldP]
		u.id[oldP] = p
		oldP = pNext
	}
	return p
}

func (u *uf) Connected(p, q int) bool {
	return u.Find(p) == u.Find(q)
}

func (u *uf) Union(p, q int) {
	i, j := u.Find(p), u.Find(q)
	if i == j { //p,q 本来就是连通的
		return
	}

	if u.sz[i] < u.sz[j] { //将小树的根节点指向大树的根节点
		u.id[i] = j
		u.sz[j] += u.sz[i]
	} else {
		u.id[j] = i
		u.sz[i] += u.sz[j]
	}

	u.count--
}
