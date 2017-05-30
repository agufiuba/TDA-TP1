package model

import (
	"math/rand"
)

type Hospital struct {
	ID           int
	Preferencias map[*Estudiante]int
	Vacantes     int
	posicionesT  []int
}

func NewHospitales(cantH, cantE int) []*Hospital {
	vacantesDisponibles := cantE
	hospitales := make([]*Hospital, cantH)
	for id := 0; id < cantH; id++ {
		h := Hospital{}
		h.ID = id + 1
		h.Preferencias = make(map[*Estudiante]int, cantE)
		base := cantE - vacantesDisponibles
		if id != (cantH - 1) {
			if vacantesDisponibles > 0 {
				prom := int(float64(vacantesDisponibles) / float64(cantH-id))
				h.Vacantes = prom - 2 + rand.Intn(4)
				if h.Vacantes < 0 {
					h.Vacantes = 0
				}
				if h.Vacantes > vacantesDisponibles {
					h.Vacantes = vacantesDisponibles
				}
				vacantesDisponibles = vacantesDisponibles - h.Vacantes
			}
		} else {
			h.Vacantes = vacantesDisponibles
		}
		h.posicionesT = make([]int, h.Vacantes)
		for i := 0; i < h.Vacantes; i++ {
			h.posicionesT[i] = base + i
		}
		hospitales[id] = &h
	}
	return hospitales
}

func (h *Hospital) Shuffle() {
	perm := rand.Perm(len(h.Preferencias))
	i := 0
	for e, _ := range h.Preferencias {
		h.Preferencias[e] = perm[i] + 1
		i++
	}
}

func (h *Hospital) GetPositionAt(i int) int {
	l := len(h.posicionesT)
	if l > 0 && i < l {
		return h.posicionesT[i]
	}
	return -1
}
