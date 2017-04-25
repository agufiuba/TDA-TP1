package model

import (
	"container/list"
	"math/rand"
)

type Hospital struct {
	ID           int
	Preferencias map[*Estudiante]int
	Vacantes     int
	Contratos    list.List
}

func NewHospital(id, cantE int) Hospital {
	h := Hospital{}
	h.Vacantes = rand.Intn(int(float64(cantE)*0.3)) + 1
	h.ID = id
	h.Preferencias = make(map[*Estudiante]int, cantE)
	return h
}

func (h *Hospital) AgregarContrato(e *Estudiante) {
	f := h.Contratos.Front()
	l := h.Contratos.Len()
	if f != nil {
		for ; f != nil; f = f.Next() {
			ec := f.Value.(*Estudiante)
			if h.Preferencias[ec] < h.Preferencias[e] {
				h.Contratos.InsertBefore(e, f)
				break
			}
		}
		if l == h.Contratos.Len() {
			h.Contratos.PushBack(e)
		}
	} else {
		h.Contratos.PushFront(e)
	}
}

func (h *Hospital) Shuffle() {
	perm := rand.Perm(len(h.Preferencias))
	i := 0
	for e, _ := range h.Preferencias {
		h.Preferencias[e] = perm[i] + 1
		i++
	}
}
