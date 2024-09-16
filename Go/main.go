package main

import (
	"ProyectoGit/Go/packages"
	"fmt"
)

func multiplicacion(a int) int {
	return a * 2
}

func saludos(a int, canal chan string, persona packages.Persona) {
	// a = cantidad de saludos
	for i := 0; i < a; i++ {
		canal <- persona.Saludar()
	}
	close(canal)
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

	fmt.Println("Vamos con el uso de go rutines")
	canal := make(chan string)

	go saludos(3, canal, persona)

	acomulador := 1
	for saludo := range canal {
		fmt.Printf("Saludo numero %d: %s\n", acomulador, saludo)
		acomulador++
	}
	fmt.Println("Esta parte es de la cabecerera")
	for i := 0; i < 5; i++ {
		fmt.Println("Cabecera")
	}
}
