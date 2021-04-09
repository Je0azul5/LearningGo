//Author: Je0azul5
package main

//lenguaje extricto con el tiado
//strconv se encarga de las conversiones de strings
import (
	"fmt"
	"strconv"
)

//definiciones de variables
var entero int
var unsignednum uint
var numeroflotante float32
var text string
var bol bool

/*Si se quiere importar las variables a otros paquetes se
definen con la primera letra en mayuscula*/
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
	//llamdo de funciones
	ModTipadoFloat()
	ModTipadoStringEntero()
}

func ModTipadoFloat() {
	//cambio de tipado de float32 a int
	nueva := int(numeroflotante)
	fmt.Println(nueva)
}

//muestra un numero en string
func ModTipadoStringEntero() {
	nn := 12
	//los porcentajes son verbos, es decir un tipo de conversion de numero a texto
	texto := fmt.Sprintf("%d", nn)
	fmt.Println(texto)
	//solo recibe entero generico
	texto = strconv.Itoa(nn)
	fmt.Println(texto)
}
