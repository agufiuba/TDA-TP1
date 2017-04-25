package punto2

import (
	"tda/tp1/punto23/model"
	"container/list"
)

func HopcroftTarjan(vertices *list.List, finder map[*model.V]*list.Element) *list.List {
	l := list.List{}
	for f := vertices.Front(); f != nil; f = f.Next() {
		v, _ := f.Value.(*model.V)
		dfs(v, 1, vertices, &l, finder)
	}
	return &l
}

func dfs(v *model.V, depth int, vertices, l *list.List, finder map[*model.V]*list.Element) {
	v.Visited = true
	vertices.Remove(finder[v])
	v.Depth = depth
	v.Low = depth
	childs := 0
	ap := false
	for _, u := range v.Out {
		visit(v, u, depth, &childs, &ap, vertices, l, finder)
	}
	for _, u := range v.In {
		visit(v, u, depth, &childs, &ap, vertices, l, finder)
	}
	if (v.Parent != nil && ap) || (v.Parent == nil && childs > 1) {
		l.PushBack(v)
	}
}

func visit(v, u *model.V, depth int, childs *int, ap *bool, vertices, l *list.List, finder map[*model.V]*list.Element) {
	if !u.Visited {
		u.Parent = v
		dfs(u, depth+1, vertices, l, finder)
		*childs = *childs + 1
		if u.Low >= v.Depth {
			*ap = true
		}
		v.Low = min(v.Low, u.Low)
	} else if u != v.Parent {
		v.Low = min(v.Low, u.Depth)
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
