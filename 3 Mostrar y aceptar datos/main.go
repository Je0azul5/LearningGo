//Author: Je0azul5
package main

// os biblioteca para el sistema operativo
//bufio para la entrada de datos si hay algun problema en la entrada de datos
import (
	"bufio"
	"fmt"
	"os"
)

var num1 int
var num2 int
var rst int
var lgd string

func main() {
	/*
		forma de con error en windows,
		en linux funciona bien
		//peticion de numero
		fmt.Println("Digite un numero: ")
		//numero base 10
		fmt.Scanf("%d", &num1)
		fmt.Println("Digite un numero: ")
		//numero base 10
		fmt.Scanf("%d", &num2)
	*/

	//metodo con puntero primera opcion de lecturra scanln
	fmt.Println("Digite un numero: ")
	//numero base 10
	fmt.Scanln(&num1)

	fmt.Println("Digite un numero: ")
	//numero base 10
	fmt.Scanln(&num2)
	//Printf imprime seguido no por linea
	fmt.Printf("Añada una descrupcion-> ")

	//segundo metodo
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		//lee la entrada
		lgd = scanner.Text()
	}
	rst = num1 + num2
	fmt.Println(lgd, rst)

}
