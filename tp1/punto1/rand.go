package punto1

import (
	"tda/tp1/punto1/model"
	"container/list"
	"math/rand"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func GenerateArrays(cantE, cantH int) (*list.List, *[]*model.Hospital) {
	estudiantes := list.New()
	hospitales := make([]*model.Hospital, 0, cantH)
	for i := 1; i <= cantE; i++ {
		e := model.NewEstudiante(i, cantH)
		estudiantes.PushFront(&e)
	}
	for i := 1; i <= cantH; i++ {
		h := model.NewHospital(i, cantE)
		hospitales = append(hospitales, &h)
	}
	i := 0
	for f := estudiantes.Front(); f != nil; f = f.Next() {
		e := f.Value.(*model.Estudiante)
		for _, h := range hospitales {
			e.AddPref(h)
			h.Preferencias[e] = i
			if i == cantE-1 {
				h.Shuffle()
			}
		}
		i++
		e.Shuffle()
	}
	return estudiantes, &hospitales
}

func FileContent(estudiantes *list.List, hospitales *[]*model.Hospital) string {
	fp := func(prefs map[*model.Estudiante]int, x int) *model.Estudiante {
		var r *model.Estudiante
		for e, p := range prefs {
			if p == x {
				r = e
				break
			}
		}
		return r
	}
	s := ""
	s += strconv.Itoa(estudiantes.Len()) + "\n"
	for f := estudiantes.Front(); f != nil; f = f.Next() {
		e, _ := f.Value.(*model.Estudiante)
		for _, h := range e.Preferencias {
			s += strconv.Itoa(h.ID) + " "
		}
		s += "\n"
	}
	s += strconv.Itoa(len(*hospitales)) + "\n"
	v := ""
	for _, h := range *hospitales {
		for j := 1; j <= len(h.Preferencias); j++ {
			e := fp(h.Preferencias, j)
			s += strconv.Itoa(e.ID) + " "
		}
		v += strconv.Itoa(h.Vacantes) + " "
		s += "\n"
	}
	s += v
	return s
}
