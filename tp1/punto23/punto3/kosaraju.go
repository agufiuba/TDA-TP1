package punto3

import (
	"tda/tp1/punto23/model"
	"container/list"
)

func Kosaraju(vertices *list.List, finder map[*model.V]*list.Element) *list.List {
	l := list.List{}
	r := list.List{}
	for f := vertices.Front(); f != nil; f = f.Next() {
		v, _ := f.Value.(*model.V)
		visit(v, &l, vertices, finder)
	}
	aux := list.List{}
	for f := l.Front(); f != nil; f = f.Next() {
		v, _ := f.Value.(*model.V)
		if assign(v, &aux, &l, finder) {
			r.PushBack(aux)
			aux = list.List{}
		}
	}
	return &r
}

func visit(v *model.V, l, vertices *list.List, finder map[*model.V]*list.Element) {
	if !v.Visited {
		v.Visited = true
		for _, u := range v.Out {
			visit(u, l, vertices, finder)
		}
		l.PushFront(v)
		vertices.Remove(finder[v])
	}
}

func assign(v *model.V, l, vertices *list.List, finder map[*model.V]*list.Element) bool {
	if !v.Assigned {
		v.Assigned = true
		for _, u := range v.In {
			assign(u, l, vertices, finder)
		}
		l.PushBack(v)
		vertices.Remove(finder[v])
		return true
	}
	return false
}
