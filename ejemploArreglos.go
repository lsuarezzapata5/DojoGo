package main

import "fmt"
/*
//Recorre el arreglo, suma el valor de todos los elemntos y muestra el total/5
func main() {
	var x [5]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 38

	var total float64 = 0
	for i := 0; i < 5; i++ {
		total += x[i]
	}
	fmt.Println(total / 5)
}*/

/*EJERCICIO: Escribir un programa que encuentre el numero menor en el 
siguiente arreglo y lo imprima:

x := []int{
  48,96,86,68,
  57,82,63,70,
  37,34,83,27,
  19,97, 9,17,
}*/
func main() {
	var x [16]int
	x[0] = 48
	x[1] = 96
	x[2] = 86
	x[3] = 68
	x[4] = 57
	x[5] = 82
	x[6] = 63
	x[7] = 70
	x[8] = 37
	x[9] = 34
	x[10] = 83
	x[11] = 27
	x[12] = 19
	x[13] = 97
	x[14] = 9
	x[15] = 17

	var menor int = x[0]
	for i := 0; i < 15; i++ {
		if x[i]<menor{
			menor = x[i]
		}
	}
	fmt.Println("El menor es: ", menor)
}