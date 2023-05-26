package algs4

type UF struct {
	id   []int
	n    int
	size map[int]int
}

func NewUF(n int) *UF {
	uf := UF{n: n}
	uf.id = make([]int, n)
	uf.size = map[int]int{}
	// 将读取到的数字存在于一个数组中
	// 每个索引就代表着一个数
	// 索引对应的值就是连通的点的索引值
	// 之后如果不同索引位置上的值相同, 就代表这些索引位置代表的点是连通的
	for i := 0; i < n; i++ {
		uf.id[i] = i
		uf.size[i] = 1
	}
	return &uf
}

func (uf *UF) root(i int) int {
	for i != uf.id[i] {
		// 路径压缩, 使得每一个节点挂到它的祖父辈节点下
		uf.id[i] = uf.id[uf.id[i]]
		i = uf.id[i]
	}
	return i
}

//func (uf *UF) root(i int) int {
//	for i != uf.id[i] {
//		i = uf.id[i]
//	}
//	return i
//}

func (uf *UF) connected(p, q int) bool {
	return uf.root(p) == uf.root(q)
}

// improved quick-union

func (uf *UF) Union(p, q int) {
	i := uf.root(p)
	j := uf.root(q)
	if i == j {
		return
	}
	if uf.size[i] < uf.size[j] {
		uf.id[i] = j
		uf.size[j] += uf.size[i]
	} else {
		uf.id[j] = i
		uf.size[i] += uf.size[j]
	}
}

//func (uf *UF) Union(p, q int) {
//	i := uf.root(p)
//	j := uf.root(q)
//	uf.id[i] = j
//}

// 直接用数组表示连接关系时的代码
//func (uf *UF) connected(p, q int) bool {
//	return uf.root(p) == uf.root(q)
//}

//func (uf *UF) Union(p, q int) {
//	pid := uf.id[p]
//	qid := uf.id[q]
//	for i := 0; i < len(uf.id); i++ {
//		if uf.id[i] == pid {
//			uf.id[i] = qid
//		}
//	}
//}

func (uf *UF) Count() {

}
