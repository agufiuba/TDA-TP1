package main

import (
	"tda/tp1/punto1"
	"tda/tp1/punto23"
	"tda/tp1/punto23/punto2"
	"tda/tp1/punto23/punto3"
	"tda/utils"
	"flag"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func main() {
	tp := flag.Int("tp", 1, "Trabajo Practico")
	ej := flag.Int("e", 0, "Ejercicio")
	r := flag.Bool("r", false, "Generar random")
	n := flag.Int("n", 4, "Cantidad de estudiantes/hospitales")
	rx := flag.Bool("rx", false, "Generar random y ejecutar")
	d := flag.Bool("d", false, "Debug")

	flag.Parse()

	if *d {
		log.SetLevel(log.DebugLevel)
	}

	if *rx {
		*r = true
	}
	if *tp == 1 {
		log.Info("Ejecutando TP1")
		switch *ej {
		case 1:
			log.Info("Ejecutando ejercicio 1 - GaleShapley")
			if *r {
				if *n > 3 {
					if len(flag.Args()) == 1 {
						o := flag.Args()[0]
						ns := strconv.Itoa(*n)
						log.Info("Generando valores aleatorios (" + ns + "x" + ns + ") en '" + o + "'")
						estudiantes, hospitales := punto1.GenerateArrays(*n, *n)
						if *d {
							log.Debug(punto1.DebugGS(estudiantes, hospitales))
						}
						err := utils.SaveToFile(o, punto1.FileContent(estudiantes, hospitales))
						if err != nil {
							log.Fatal(err)
						}
						if *rx {
							log.Info("Aplicando algoritmo GaleShapley")
							log.Info(punto1.ShowGS(punto1.GaleShapley(estudiantes, hospitales)))
						}
					} else {
						log.Fatal("Falta parametro de salida")
					}
				} else {
					log.Fatal("La cantidad de estudiantes/hospitales debe ser mayor a 3")
				}
			} else {
				if len(flag.Args()) == 1 {
					i := flag.Args()[0]
					log.Info("Parseando preferencias desde '" + i + "'")
					estudiantes, hospitales, _ := punto1.LoadPref(i)
					if *d {
						log.Debug(punto1.DebugGS(estudiantes, hospitales))
					}
					log.Info("Aplicando algoritmo GaleShapley")
					log.Info(punto1.ShowGS(punto1.GaleShapley(estudiantes, hospitales)))
				} else {
					log.Fatal("Falta parametro de entrada")
				}
			}
		case 2:
			log.Info("Ejecutando ejercicio 2 - HopcroftTarjan")
			if len(flag.Args()) == 1 {
				i := flag.Args()[0]
				log.Info("Parseando grafo desde '" + i + "'")
				g, finder := punto23.LoadGraph(i)
				if *d {
					log.Debug(punto23.DebugGraph(g))
				}
				log.Info("Aplicando algoritmo HopcroftTarjan")
				aps := punto2.HopcroftTarjan(g, finder)
				log.Info(punto2.ShowHT(aps))
			} else {
				log.Fatal("Falta parametro de entrada")
			}
		case 3:
			log.Info("Ejecutando ejercicio 3 - Kosaraju")
			if len(flag.Args()) == 1 {
				i := flag.Args()[0]
				log.Info("Parseando grafo desde '" + i + "'")
				g, finder := punto23.LoadGraph(i)
				if *d {
					log.Debug(punto23.DebugGraph(g))
				}
				log.Info("Aplicando algoritmo Kosaraju")
				cfcs := punto3.Kosaraju(g, finder)
				log.Info(punto3.ShowK(cfcs))
			} else {
				log.Fatal("Falta parametro de entrada")
			}
		}
	}
}
