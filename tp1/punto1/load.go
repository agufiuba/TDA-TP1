package punto1

import (
	"tda/tp1/punto1/model"
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func LoadPref(fn string) (*list.List, *[]*model.Hospital, error) {
	f, err := os.Open(fn)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()
	var cantE, cantH int
	estudiantes := list.List{}
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%d", &cantE)
	for i := 0; i <= cantE; i++ {
		scanner.Scan()
	}
	fmt.Sscanf(scanner.Text(), "%d", &cantH)
	estudiantesS := make([]*model.Estudiante, cantE)
	for i := 0; i < cantE; i++ {
		e := model.NewEstudiante(i+1, cantH)
		estudiantes.PushFront(&e)
		estudiantesS[i] = &e
	}
	hospitales := make([]*model.Hospital, 0, cantH)
	for i := 1; i <= cantH; i++ {
		h := model.NewHospital(i, cantE)
		hospitales = append(hospitales, &h)
	}
	f.Close()
	f, err = os.Open(fn)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()
	scanner = bufio.NewScanner(f)
	scanner.Scan()
	el := estudiantes.Front()
	e := el.Value.(*model.Estudiante)
	for i := 0; i < cantE; i++ {
		scanner.Scan()
		txt := scanner.Text()
		prefS := strings.Fields(txt)
		for _, p := range prefS {
			pInt, _ := strconv.Atoi(p)
			h := hospitales[pInt-1]
			e.AddPref(h)
		}
		el = el.Next()
		if el != nil {
			e = el.Value.(*model.Estudiante)
		} else {
			break
		}
	}
	scanner.Scan()
	for i := 0; i < cantH; i++ {
		scanner.Scan()
		prefS := strings.Fields(scanner.Text())
		for j, p := range prefS {
			pInt, _ := strconv.Atoi(p)
			hospitales[i].Preferencias[estudiantesS[pInt-1]] = j
		}
	}
	scanner.Scan()
	vacantes := strings.Fields(scanner.Text())
	for i, v := range vacantes {
		vInt, _ := strconv.Atoi(v)
		hospitales[i].Vacantes = vInt
	}
	return &estudiantes, &hospitales, nil
}
