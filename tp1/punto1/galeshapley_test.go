package punto1

import (
	"tda/tp1/punto1/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	cantE                   = 5
	cantH                   = 10
	estudiantes, hospitales = GenerateArrays(cantE, cantH)
)

func TestGaleShapley(t *testing.T) {
	model.NewHospitales(cantH, cantE)
	// for _, h := range hospitales2 {
	// 	fmt.Println(h.Vacantes)
	// }
	contratos := GaleShapley(estudiantes, hospitales)
	estable := true
	for h, es := range contratos {
		if !estable {
			break
		}
		for f := es.Front(); f != nil; f = f.Next() {
			if !estable {
				break
			}
			e, _ := f.Value.(*model.Estudiante)
			for h2, es2 := range contratos {
				if !estable {
					break
				}
				if e.PreferenciasI[h2] < e.PreferenciasI[h] {
					for f2 := es2.Front(); f2 != nil; f2 = f2.Next() {
						e2 := f2.Value.(*model.Estudiante)
						estable = h2.Preferencias[e2] < h2.Preferencias[e]
						if !estable {
							break
						}
					}
				}
			}
		}
	}
	assert.Equal(t, estable, true, "Estable")
}
