package main;

import (
	"fmt"
)

func main() {
	var suma int = 8 + 9
	var resta int = 6 - 4
	var nombre string = "Cesar Vargas"
	var prueba bool = true
	var prueba2 float32 = 12.25565
	fmt.Println(prueba)
	fmt.Println(prueba2)
	var apellidos string = "vargas benavides"
	pais := "colombia"
	const year int = 2018
	apellidos = "martines"
	fmt.Println("Hola mundo desde go con " + nombre+" "+apellidos+" pais "+pais)
	fmt.Println(suma)
	fmt.Println(resta)
	fmt.Println(year)
}
