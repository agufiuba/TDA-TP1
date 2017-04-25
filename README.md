# Teoría de algoritmos - TP 1
## Instalación
Se necesita tener instalado Golang y tener exportado el path `$GOPATH/bin`.
```
go get https://github.com/agufiuba/tda
```
## Ejecución
### Punto 1 - GaleShapley
#### Generar preferencias de prueba de tamaño `N`
```
tda -e=1 -r -n=N /path/to/export
```
#### Generar preferencias de prueba de tamaño `N` y ejecutar algoritmo
```
tda -e=1 -rx -n=N /path/to/export
```
#### Ejecutar algoritmo a partir de un archivo de preferencias
```
tda -e=1 /path/to/import
```

### Punto 2 - HopcroftTarjan
```
tda -e=2 /path/to/import
```

### Punto 3 - Kosaraju
```
tda -e=3 /path/to/import
```

## Debug y testing
Todos lo modos de ejecución cuentan con la opción debug -d .
El algoritmo de GaleShapley se puede probar de la siguiente manera:
`go test ./... -v`
