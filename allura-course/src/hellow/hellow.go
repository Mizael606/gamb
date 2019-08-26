package main

import (
	"fmt"
)

func main() {
	iniciarMonitoramento()
	initialise()
	tname := getName()
	tidade := getIdade()
	fmt.Println("Olá , sr(a).\n", tname, "\nSua idade é ", tidade)
}

func initialise() {
	name := "Mel Bianco"
	idade := 26
	fmt.Println("Olá , sr(a).\n", name, "\nSua idade é ", idade)
	fmt.Println("Crie um form como acima:")
}

func getIdade() int {
	var youridade int
	fmt.Println("2-Sua idade:")
	fmt.Scan(&youridade)
	return youridade
}

func getName() string {
	var yourname string
	fmt.Println("1-Seu nome:")
	fmt.Scan(&yourname)
	return yourname
}

func iniciarMonitoramento() {
	fmt.Println("Iniciando aplicação:")
}
