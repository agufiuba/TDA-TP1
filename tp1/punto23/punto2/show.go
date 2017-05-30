package punto2

import (
	"container/list"
	"strconv"
	"strings"
	"tda/tp1/punto23/model"
)

func ShowHT(aps *list.List) string {
	s := "Articulation Points => "
	for f := aps.Front(); f != nil; f = f.Next() {
		ap, _ := f.Value.(*model.V)
		s += strconv.Itoa(ap.ID) + ", "
	}
	s = strings.TrimRight(s, ", ")
	return s
}
