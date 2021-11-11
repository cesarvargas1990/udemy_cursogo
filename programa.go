package main

import ("fmt"
	 "os"
	 "strconv"
	 "time"
	)


func main() {
	fmt.Println("***********")
	fmt.Println(os.Args)
	fmt.Println("Hola "+os.Args[1]+" bienvenido al programa en GO")

	edad,err := strconv.Atoi(os.Args[2])
	
	if edad >= 18 && edad < 99  {
		fmt.Println("Eres mayor de edad")
	}else if edad <= 18 {
		fmt.Println("Eres menor de edad o demasiado mayor")
	}else if edad == 99 {
		fmt.Println("Pronto moriras")
	}

	numero,_:= strconv.Atoi(os.Args[2])
	fmt.Println(err)
	if numero%2 == 0 {
		fmt.Println("el numero es par")
	}else{
		fmt.Println("el numero es impar")
	}

	peliculas := []string{"pelicula1","pelicula2","pelicula3"}
	for i := 0; i<len(peliculas); i++ {
		fmt.Println(peliculas[i])
	}
	for _, pelicula := range peliculas {
		fmt.Println(pelicula)
	}

	momento := time.Now()
	hoy := momento.Weekday()
	switch hoy {
		case 0: 
				fmt.Println("Hoy es domingo")
		case 1:
			fmt.Println("Hoy es lunes")
		case 2:
			fmt.Println("Hoy es martes")
		case 3:
			fmt.Println("Hoy es miercoles")
		case 4:
			fmt.Println("Hoy es jueves")
		case 5:
			fmt.Println("Hoy es viernes")
		case 6:
			fmt.Println("Hoy es sabado")
		default:
			fmt.Println("Hoy es otro dia")
			
	}
}