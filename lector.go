package main

import "fmt"
import "io/ioutil"
import "os"

func main() {
	fmt.Println ("lector")
	


	nuevo_texto := os.Args[1]+"\n"

	fichero2, err := os.OpenFile("fichero.txt",os.O_APPEND|os.O_WRONLY,0777)
	showError(err)

	escribir , err2 := fichero2.WriteString(nuevo_texto)
	fmt.Println(escribir)
	showError(err2)
	fichero2.Close();


	texto,errorDeFichero := ioutil.ReadFile("fichero.txt")
	showError(errorDeFichero)
	fmt.Println(string(texto))

	
	//pasar paramertos en consola asi:
	//go run lector.go "10 otro lenguaje"

	

}

func showError(e error) {
	if (e != nil) {
		panic(e)
	}
}
