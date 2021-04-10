//Author: Je0azul5

//solo existe el for
package main

import (
	"fmt"
)

func main() {

	var n int
	//primera forma de hacer un for, "Iterativo"
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//segunda forma for ciclo infinito
	for {
		fmt.Println("Digite el numero inicial el for")
		fmt.Scanln(&n)
		if n == 25 {
			fmt.Println("Encontro el numero secreto")
			n = 0
			break
		}
	}

	//Tercera forma de hacer un for
	for n < 15 {
		//los verbos se remplazan en el texto
		fmt.Printf("valor de i->%d, ", n)

		if n == 5 {
			fmt.Println("Uy salta de linea")
			n = n * 2
			// palabra reservada "continue", es utilizada para volver a
			// comenzar el for sin ejecutar el codigo de abajo
			continue
		}
		n++
		fmt.Println("Sos un boludo")
	}

}
