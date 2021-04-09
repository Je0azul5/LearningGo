package main

import (
	"fmt"
)

var estado bool = true

func main() {
	//Primera estructura de conducion
	if estado {
		fmt.Print("condicion cumplida\n")
	} else {
		fmt.Print("condicion no cumplida\n")
	}

	//se puede asignar un valor dentro del if
	a := 12
	if a = 21; a == 21 {
		fmt.Println("El resultado es 21")
	}

	switch a = 1; a {
	case 1:
		fmt.Println(a + 1)
	case 2:
		fmt.Println(a - 2)
	case 3:
		fmt.Println(a + 3)
	case 4:
		fmt.Println(a + 34)
	case 5:
		fmt.Println(a + 123)

	}

}
