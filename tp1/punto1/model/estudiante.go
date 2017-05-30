package model

import (
	"math/rand"
)

type Estudiante struct {
	ID            int
	Preferencias  []*Hospital
	PreferenciasI map[*Hospital]int
	Propuestas    int
	prefCount     int
	groupH				int
	h							int
}

func NewEstudiante(id, cantH int) Estudiante {
	e := Estudiante{}
	e.ID = id
	e.Preferencias = make([]*Hospital, cantH)
	e.PreferenciasI = make(map[*Hospital]int, cantH)
	return e
}

func (e *Estudiante) Shuffle() {
	for i := range e.Preferencias {
		j := rand.Intn(i + 1)
		e.Preferencias[i], e.Preferencias[j] = e.Preferencias[j], e.Preferencias[i]
		e.PreferenciasI[e.Preferencias[i]] = i
		e.PreferenciasI[e.Preferencias[j]] = j
	}
}

func (e *Estudiante) AddPref(h *Hospital) {
	e.Preferencias[e.prefCount] = h
	e.PreferenciasI[h] = e.prefCount
	e.prefCount++
}

func (e *Estudiante) GetNextProp() int {
	pos := e.Preferencias[e.h].GetPositionAt(e.groupH)
	if pos != -1 {
		e.groupH++
	} else {
		e.h++
		e.groupH = 0
		pos = e.GetNextProp()
	}
	return pos
}
