package accountsMerge

type UnionFind struct {
	par  []int
	rank []int
}

func NewUnionFind(n int) *UnionFind {
	res := UnionFind{
		par:  make([]int, n),
		rank: make([]int, n),
	}
	for i := range n {
		res.par[i] = i
		res.rank[i] = 1
	}
	return &res
}

func (uf *UnionFind) Find(x int) int {
	for x != uf.par[x] {
		uf.par[x] = uf.par[uf.par[x]]
		x = uf.par[x]
	}
	return x
}

func (uf *UnionFind) Union(x1, x2 int) bool {
	p1, p2 = uf.Find(x1), uf.Find(x2)

	if p1 == p2 {
		return false
	}
}

func accountsMerge(accounts [][]string) [][]string {
}
