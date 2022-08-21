package main

import (
	"fmt" // no more error
)

/*
    1) Haga un programa que cuente e indique el número de caracteres,
    el número de palabras y el número de líneas de un texto
    cualquiera (obtenido de cualquier forma que considere).
	Considere hacer el cálculo de las tres variables en el mismo proceso.
*/
func ejercicio1(t string) {
	var caracteres = 0
	var lineas = 0
	var palabras = 0
	for i, letra := range t{
		var c = string(letra)
		if(c == "\n"){//los . y , son carac especiales por lo que no se cuentan
			lineas++;
		}
		if(c == "\n" || c == " "){
			palabras++;
		}
		if(c != " " && c != "\n"){
			caracteres++;
		}
		_ = i
	}
	// fmt.Println(len(t))
	fmt.Printf("Caracteres: %d\n", caracteres)
	fmt.Printf("Palabras: %d\n", palabras)
	fmt.Printf("Lineas: %d\n", lineas)

}

/*
	Resultados:

	-   Caracteres: 198
		Palabras: 36
		Lineas: 7   
*/


// func main(){
// 	texto :=`Yo quiero ser llorando el hortelanon
// de la tierra que ocupas y estercolas,
// compañero del alma, tan temprano.

// Alimentando lluvias, caracolas
// y organos mi dolor sin instrumento,
// a las desalentadas amapolas
// dare tu corazón por alimento.`

// 	ejercicio1(texto)

// }