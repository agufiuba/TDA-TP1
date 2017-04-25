package model

import (
	"sync"
)

type V struct {
	ID       int
	Visited  bool
	Assigned bool
	CFCGroup int
	Parent   *V
	Low      int
	Depth    int
	IsAP     bool
	Out      []*V
	In       []*V
	M        sync.Mutex
	M2       sync.Mutex
}

func (v *V) AddOut(u *V) {
	v.M.Lock()
	v.Out = append(v.Out, u)
	v.M.Unlock()
}

func (v *V) AddIn(u *V) {
	v.M.Lock()
	v.In = append(v.In, u)
	v.M.Unlock()
}
