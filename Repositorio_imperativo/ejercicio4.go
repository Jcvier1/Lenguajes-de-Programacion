package main

import (
	"fmt" // no more error
)

type lista_zapatos []calzado

var lZapatos lista_zapatos

type calzado struct  {
	marca string
	precio float32
	talla int
	stock int

}
// type inventario struct {
// 	zapatos   []calzado
// 	stock int
// }

func (l *lista_zapatos) agregarZapato(m string, p float32, t int) {

	if(t > 33 && t < 45){
		aux := l.buscarZapato(m, t)
		if aux == -1 {
			lZapatos = append(lZapatos, calzado{marca: m, precio: p, talla: t, stock: 1})	
		} else {
			(*l)[aux].stock += 1
			(*l)[aux].precio = p
		}
		
	} else {
		fmt.Println("No se puede agregar porque no cumple con la 'talla'")
	}
	
}

func (l *lista_zapatos) buscarZapato(m string, t int) int {
	var result = -1
	var i int
	for i = 0; i < len(*l); i++ {
		if((*l)[i].marca == m && (*l)[i].talla == t){
			result = i
		}
	}
	return result
}




func llenarDatos() {
	lZapatos.agregarZapato("Nike", 150000, 34)
	lZapatos.agregarZapato("Adidas", 40000, 37)
	lZapatos.agregarZapato("Converse", 80000, 38)
	lZapatos.agregarZapato("Puma", 12000, 39)
	lZapatos.agregarZapato("Tin", 100000, 34)
	lZapatos.agregarZapato("Dor", 60000, 40)
	lZapatos.agregarZapato("tenny", 30000, 38)
	lZapatos.agregarZapato("Jordan", 19000, 41)
	lZapatos.agregarZapato("Petr", 15000, 34)
	lZapatos.agregarZapato("Super Teni", 70000, 36)
	
	lZapatos.agregarZapato("Adidas", 40000, 37)


}

/*
	Resultados:
	{Nike 150000 34 1}
	{Adidas 40000 37 2}  
	{Converse 80000 38 1}
	{Puma 12000 39 1}    
	{Tin 100000 34 1}
	{Dor 60000 40 1}
	{tenny 30000 38 1}
	{Jordan 19000 41 1}
	{Petr 15000 34 1}
	{Super Teni 70000 36 1}


*/

// func main() {
// 	llenarDatos()
// 	var i int
// 	for i = 0; i < len(lZapatos); i++ {
// 		fmt.Println(lZapatos[i])
// 	}


// }
