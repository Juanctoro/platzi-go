package main

import (
	"ProyectoGit/packages"
	"fmt"
)

func multiplicacion(a int) int {
	return a * 2
}

func main() {
	fmt.Println("Primer commit")

	fmt.Println("Para el segundo commit voy a hacer una funcion que calcule el doble de un valor")
	fmt.Println(multiplicacion(3))

	fmt.Println("Crando clase Persona...")
	persona := packages.Persona{
		Nombre: "Juan",
		Edad:   20,
		Altura: 170,
	}

	fmt.Println(persona.Saludar())
}
