package punto23

import (
	"container/list"
	"strconv"
	"strings"
	"tda/tp1/punto23/model"
)

func DebugGraph(vertices *list.List) string {
	s := ""
	i := 0
	for f := vertices.Front(); f != nil; f = f.Next() {
		if i != 0 {
			s += "           "
		}
		v, _ := f.Value.(*model.V)
		adj := v.Out
		adj = append(adj, v.In...)
		s += strconv.Itoa(v.ID) + " => "
		for _, a := range adj {
			s += strconv.Itoa(a.ID) + ", "
		}
		s = strings.TrimRight(s, ", ")
		s += "\n"
		i++
	}
	s = strings.TrimRight(s, "\n")
	return s
}
