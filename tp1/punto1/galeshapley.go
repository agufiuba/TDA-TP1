package punto1

import (
	"tda/tp1/punto1/model"
	"container/list"
	"fmt"
	"strconv"
	"strings"
)

var _ = fmt.Println

func GaleShapley(estudiantes *list.List, hospitales *[]*model.Hospital) map[*model.Hospital]*list.List {
	contratos := make(map[*model.Hospital]*list.List, len(*hospitales))
	f := estudiantes.Front()
	if f != nil {
		e, _ := f.Value.(*model.Estudiante)
		for e.Propuestas < len(*hospitales) {
			h := e.Preferencias[e.Propuestas]
			e.Propuestas++
			if h.Vacantes > h.Contratos.Len() {
				h.AgregarContrato(e)
				contratos[h] = &h.Contratos
				f2 := f.Next()
				estudiantes.Remove(f)
				if f2 == nil {
					break
				}
				f = f2
			} else {
				peorElement := h.Contratos.Front()
				peor := peorElement.Value.(*model.Estudiante)
				if h.Preferencias[peor] > h.Preferencias[e] {
					h.Contratos.Remove(peorElement)
					h.AgregarContrato(e)
					contratos[h] = &h.Contratos
					estudiantes.Remove(f)
					f = estudiantes.PushFront(peor)
				}
			}
			e, _ = f.Value.(*model.Estudiante)
		}
	}
	return contratos
}

func ShowGS(r map[*model.Hospital]*list.List) string {
	s := ""
	i := 0
	for h, es := range r {
		if i != 0 {
			s += "           "
		}
		s += strconv.Itoa(h.ID) + " => "
		for f := es.Front(); f != nil; f = f.Next() {
			e, _ := f.Value.(*model.Estudiante)
			s += strconv.Itoa(e.ID) + ", "
		}
		s = strings.TrimRight(s, ", ")
		s += "\n"
		i++
	}
	s = strings.TrimRight(s, "\n")
	return s
}
