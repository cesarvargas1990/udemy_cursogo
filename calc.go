package main

import "fmt"

type Gorra struct {
	marca  string
	color  string
	precio float32
	plana  bool
}

func main() {
	//var numero1 int = 10
	//var numero2 int = 6

	/*var gorra_negra = Gorra {
	marca: "Nike",
	color: "Negro",
	precio: 25.50,
	plana: false} */
	var peliculas [3]string;
	peliculas[0] = "la verdad duele"
	peliculas[1] = "ciudadano ejemplar"
	peliculas[2] = "el matarife"

	fmt.Println(peliculas[0])
	fmt.Println(peliculas[1])
	fmt.Println(peliculas[2])
	var gorra_negra = Gorra{"Nike", "Negro", 25.50, false}

	peliculas2 := [3]string{"chuky","matador","azucena"}


	fmt.Println(peliculas2)


		var peliculas3 [3][2]string;
		peliculas3[0][0] = "pelicula 1.1"
		peliculas3[0][1] = "pelicula 1.2"

		peliculas3[1][0] = "pelicula 2.1"
		peliculas3[1][1] = "pelicula 2.2"

		peliculas3[2][0] = "pelicula 3.1"
		peliculas3[2][1] = "pelicula 3.2"

		fmt.Println(peliculas3)

	    peliculas4 := []string{"pelicula1","pelicula2","pelicula3","peliculan"}
		fmt.Println(peliculas4)
		peliculas4 = append(peliculas4,"masacrador")
		peliculas4 = append(peliculas4,"masacrador2")
		fmt.Println(peliculas4[4])
		fmt.Println("###########################")
		fmt.Println(len(peliculas4))
		fmt.Println(peliculas4[0:3])



	fmt.Println(gorra_negra)
	fmt.Println(gorra_negra.marca)
	fmt.Println(gorra_negra.color)

	var numero1 float32 = 10
	var numero2 float32 = 6
	var numero3 float32 = 15
	var numero4 float32 = 6

	// Suma
	fmt.Print("La suma es: ")
	fmt.Println(numero1 + numero2)

	fmt.Println(operacion(numero1, numero2, "suma"))

	// Resta
	fmt.Print("La resta es: ")
	fmt.Println(numero1 - numero2)

	// Multiplicacion
	fmt.Print("La multiplicacion es: ")
	fmt.Println(numero1 * numero2)

	// Division
	fmt.Print("La Division es: ")
	fmt.Println(numero1 / numero2)

	holaMundo()
	holaMundo()

	fmt.Println(operacion(numero1,numero2,"+"))
	fmt.Println(operacion(numero1,numero2,"-"))
	fmt.Println(operacion(numero1,numero2,"*"))
	fmt.Println(operacion(numero1,numero2,"/"))

	fmt.Println("Calculadora 1 *****")
	calculadora(numero1,numero2)
	fmt.Println("Calculadora 2 *****")
	calculadora(numero3,numero4)
	devolverTexto();
	devolverTexto2();
	fmt.Println(devolverTexto3(6));

	fmt.Println(devolverTexto4())

	fmt.Println("pedido 1");
		fmt.Println(gorras(2,"pesos"));


}

func holaMundo() {
	fmt.Println("Hola mundo!!")
}

func operacion(n1 float32,n2 float32,op string) float32 {
	var resultado float32
	if (op=="+") {
		resultado = n1+n2
	}
	if (op=="-") {
		resultado = n1-n2
	}
	if (op=="*") {
		resultado = n1*n2
	}
	if (op=="/") {
		resultado = n1/n2
	}

	return resultado

}

func calculadora (numero1 float32, numero2 float32)  {
		// Suma
		fmt.Print("La suma es: ")
		fmt.Println(operacion(numero1,numero2,"+"))
	
		
	
		// Resta
		fmt.Print("La resta es: ")
		fmt.Println(operacion(numero1,numero2,"-"))
	
		// Multiplicacion
		fmt.Print("La multiplicacion es: ")
		fmt.Println(operacion(numero1,numero2,"*"))
	
		// Division
		fmt.Print("La Division es: ")
		fmt.Println(operacion(numero1,numero2,"/"))

		

}

func devolverTexto () {
	fmt.Println ("texto por defecto")
}

func devolverTexto2() int {
	return 5
}

func devolverTexto3 (dato int) (string , int, string) {
	return "prueba" , dato, "magia"
}

func devolverTexto4 () (dato1 string , numero1 int, dato2 string) {
	dato1 = "cesar"
	numero1 = 5
	dato2 = "prueba"
	return
}

func gorras (cantidad float32, moneda string  ) (string, float32, string) {

	precio := func() float32 {
		return cantidad * 7
	}  

	return "el precio total del pedido es: ", precio(), moneda
}