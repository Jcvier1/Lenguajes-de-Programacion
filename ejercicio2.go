package main

import (
	"fmt" // no more error
)

/* Escriba una función que permita rotar una secuencia de elementos N
posiciones a la izquierda o a la derecha según sea el caso en función al
parámetro que se reciba. Los parámetros deben ser un puntero a un arreglo previamente
creado, la cantidad de movimiento de rotación y la dirección (0 = izq y 1 = der).*/
// [a, b, c], 2, izq -> [c, a, b]

func ejercicio2(lista_numeros *[]string, num int, direc int) {
	var nl = []string{}
	for i, leta := range *lista_numeros { //creamos una copia que va a ser con la que vamos a trabajar
		nl = append(nl, leta)
		_ = i
	}
	for i := 0; i < num; i++ {
		if direc == 0 {
			nl = append(nl, nl[0])
			nl = append(nl[:0], nl[0+1:]...)
		} else if direc == 1 {
			nl = append(nl[:0+1], nl[0:]...)
			a := len(nl) - 1
			nl[0] = nl[a]
			nl = append(nl[:a], nl[a+1:]...)

		} else {
			fmt.Println("the direccion is not valid")
		}
	}
	fmt.Println("Arreglo original: ", *lista_numeros)
	fmt.Println("Arreglo rotado: ", nl)

}

/*
	Resultados:

	-   Arreglo original:  [a b c d e f g]
		Arreglo rotado:  [e f g a b c d]
*/

func main() {

	direc := 1 //1 derecha, 0 izquierda
	var lista_numeros = []string{"a", "b", "c", "d", "e", "f", "g"}
	ejercicio2(&lista_numeros, 3, direc)

}
