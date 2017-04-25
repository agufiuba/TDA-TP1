package punto23

import (
	"tda/tp1/punto23/model"
	"bufio"
	"container/list"
	"fmt"
	"log"
	"os"
	"sync"
)

func LoadGraph(fn string) (*list.List, map[*model.V]*list.Element) {
	f, err := os.Open(fn)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var nov, noe int
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%d", &nov)
	vertices := make(map[int]*model.V, nov)
	l := list.List{}
	finder := make(map[*model.V]*list.Element, nov)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for i := 0; i < nov; i++ {
			// TODO: concurrente
			v := model.V{}
			v.ID = i
			vertices[i] = &v
			finder[&v] = l.PushBack(&v)
		}
		wg.Done()
	}()

	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%d", &noe)
	ch := make(chan model.E, noe)
	wg.Wait()
	go func() {
		for scanner.Scan() {
			var from, to int
			fmt.Sscanf(scanner.Text(), "%d %d", &from, &to)
			ch <- model.E{from, to}
		}
		close(ch)
	}()
	wg.Add(2 * noe)
	for e := range ch {
		go func(edge model.E) {
			vertices[edge.From].AddOut(vertices[edge.To])
			wg.Done()
		}(e)
		go func(edge model.E) {
			vertices[edge.To].AddIn(vertices[edge.From])
			wg.Done()
		}(e)
	}
	wg.Wait()
	return &l, finder
}
