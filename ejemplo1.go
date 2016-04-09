package main

import "fmt" //Permite cargar las opciones de imprimir, scan
/*
//Dado un numero ingresado, imprime 2*numero
func main() {
  fmt.Print("Ingrese un numero: ")
  var input float64
  fmt.Scanf("%f", &input)

  output := input * 2

  fmt.Println(output)
}
*/
/*EJERCICIO: Cambie el anterior programa para que en vez de capturar
  un decimal, capture un texto y lo imprima.*/

func main() {
  fmt.Print("Ingrese un texto: ")
  var input string
  fmt.Scanf("%s", &input)

  output := input

  fmt.Println(output)
}
