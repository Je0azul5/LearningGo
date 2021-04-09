//Author: Je0azul5
package main

import "fmt"

//definiciones de variables
var entero int
var unsignednum uint
var numeroflotante float32
var text string
var bol bool

//inicializacion multiple
var uno, dos, tres int

/*
Siempre las numericas inicializan en 0
los strings con cadenas vacias
los bool con false*/

func main() {
	//otra forma de asignacion individual
	texto := "numeroEntero := 3"
	//asignacion multiple
	doce, ss, cuatro, ss2 := 12, "txt en variable", 4, "txt concatenado"
	fmt.Println(texto, "\n->", doce, "->", ss, "->", cuatro, "->", ss2, "\n-algo")

}
