package main

import (
	"fmt"
	"os"
)

func ascendente(palabras []string) {
	file, err := os.Create("ascendente.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
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
}
