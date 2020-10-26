package main

import (
	"fmt"
	"os"
	"sort"
)

type PalabraAscendente []string

func (a PalabraAscendente) Len() int           { return len(a) }
func (a PalabraAscendente) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a PalabraAscendente) Less(i, j int) bool { return a[i] < a[j] }

func ascendente(palabras []string) {
	file, err := os.Create("ascendente.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	sort.Sort(PalabraAscendente(palabras))

	for _, v := range palabras {
		file.WriteString(v)
		file.WriteString("\n")
	}
}

type PalabraDescendente []string

func (a PalabraDescendente) Len() int           { return len(a) }
func (a PalabraDescendente) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a PalabraDescendente) Less(i, j int) bool { return a[i] > a[j] }

func descendente(palabras []string) {
	file, err := os.Create("descendente.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	sort.Sort(PalabraDescendente(palabras))

	for _, v := range palabras {
		file.WriteString(v)
		file.WriteString("\n")
	}
}

func main() {
	var variable string
	var arreglo []string
	for variable != "0" {
		fmt.Println("Inserte string para el slice, presiones 0 para dejar de insertar")
		fmt.Scan(&variable)
		if variable == "0" {
			fmt.Println("Slice: ", arreglo)
		} else {
			arreglo = append(arreglo, variable)
		}
	}
	ascendente(arreglo)
	descendente(arreglo)
}
