package punto1

import (
	"container/list"
	"sort"
	"strconv"
	"strings"
	"tda/tp1/punto1/model"
)

func DebugGS(estudiantes *list.List, hospitales *[]*model.Hospital) string {
	s := "Preferencias E:\n"
	for f := estudiantes.Front(); f != nil; f = f.Next() {
		s += "                           "
		e, _ := f.Value.(*model.Estudiante)
		s += strconv.Itoa(e.ID) + " => "
		for _, h := range e.Preferencias {
			s += strconv.Itoa(h.ID) + ", "
		}
		s = strings.TrimRight(s, ", ")
		s += "\n"
	}
	s += "           Preferencias H | Vacante H:\n"
	for _, h := range *hospitales {
		s += "                           "
		s += strconv.Itoa(h.ID) + " => "
		for _, kv := range sortPref(h.Preferencias) {
			s += strconv.Itoa(kv.K.ID) + ", "
		}
		s = strings.TrimRight(s, ", ")
		s += " | " + strconv.Itoa(h.Vacantes)
		s += "\n"
	}
	s = strings.TrimRight(s, "\n")
	return s
}

func sortPref(pref map[*model.Estudiante]int) pairList {
	pl := make(pairList, len(pref))
	i := 0
	for k, v := range pref {
		pl[i] = pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

type pair struct {
	K *model.Estudiante
	V int
}

type pairList []pair

func (p pairList) Len() int           { return len(p) }
func (p pairList) Less(i, j int) bool { return p[i].V > p[j].V }
func (p pairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
