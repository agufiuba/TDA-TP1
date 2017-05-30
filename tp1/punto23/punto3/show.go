package punto3

import (
	"container/list"
	"strconv"
	"strings"
	"tda/tp1/punto23/model"
)

func ShowK(cfcs *list.List) string {
	s := ""
	i := 1
	for f := cfcs.Front(); f != nil; f = f.Next() {
		cfc, _ := f.Value.(list.List)
		if i != 1 {
			s += "           "
		}
		s += "Grupo " + strconv.Itoa(i) + " => "
		for f2 := cfc.Front(); f2 != nil; f2 = f2.Next() {
			v, _ := f2.Value.(*model.V)
			s += strconv.Itoa(v.ID) + ", "
		}
		s = strings.TrimRight(s, ", ")
		s += "\n"
		i++
	}
	s = strings.TrimRight(s, "\n")
	return s
}
