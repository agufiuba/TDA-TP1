package punto1

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"
	"tda/tp1/punto1/model"
)

var _ = fmt.Println

func GaleShapley(estudiantes *list.List, hospitales *[]*model.Hospital) map[*model.Hospital]*list.List {
	contratos := make(map[*model.Hospital]*list.List, len(*hospitales))
	parejas := make([]*model.Pareja, estudiantes.Len())
	posicionTrans := 0
	for _, h := range *hospitales {
		for i := 0; i < h.Vacantes; i++ {
			parejas[posicionTrans] = &model.Pareja{h, nil, -1}
			posicionTrans++
		}
	}
	f := estudiantes.Front()
	if f != nil {
		e, _ := f.Value.(*model.Estudiante)
		for e.Propuestas < len(parejas) {
			np := e.GetNextProp()
			p := parejas[np]
			e.Propuestas++
			if p.Estudiante == nil {
				p.Estudiante = e
				f2 := f.Next()
				estudiantes.Remove(f)
				if f2 == nil {
					break
				}
				f = f2
			} else {
				actual := p.Estudiante
				if p.Hospital.Preferencias[actual] > p.Hospital.Preferencias[e] {
					p.Estudiante = e
					estudiantes.Remove(f)
					f = estudiantes.PushFront(actual)
				} else {
				}
			}
			e, _ = f.Value.(*model.Estudiante)
		}
	}
	for _, p := range parejas {
		if contratos[p.Hospital] == nil {
			contratos[p.Hospital] = list.New()
		}
		contratos[p.Hospital].PushBack(p.Estudiante)
	}
	return contratos
	// f := estudiantes.Front()
	// if f != nil {
	// 	e, _ := f.Value.(*model.Estudiante)
	// 	for e.Propuestas < len(*hospitales) {
	// 		h := e.Preferencias[e.Propuestas]
	// 		e.Propuestas++
	// 		if h.Vacantes > h.Contratos.Len() {
	// 			h.AgregarContrato(e)
	// 			contratos[h] = &h.Contratos
	// 			f2 := f.Next()
	// 			estudiantes.Remove(f)
	// 			if f2 == nil {
	// 				break
	// 			}
	// 			f = f2
	// 		} else {
	// 			peorElement := h.Contratos.Front()
	// 			peor := peorElement.Value.(*model.Estudiante)
	// 			if h.Preferencias[peor] > h.Preferencias[e] {
	// 				h.Contratos.Remove(peorElement)
	// 				h.AgregarContrato(e)
	// 				contratos[h] = &h.Contratos
	// 				estudiantes.Remove(f)
	// 				f = estudiantes.PushFront(peor)
	// 			}
	// 		}
	// 		e, _ = f.Value.(*model.Estudiante)
	// 	}
	// }
	// return contratos
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
