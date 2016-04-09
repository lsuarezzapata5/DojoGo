package main

import "fmt" 

func main() {
	p:=Persona{"Luisa", 1.70}
	fmt.Println(p.correr(), "corriendo" )
}

type Persona struct{
	nombre string
	estatura float64
}

func (p *Persona) correr()  string{
	return p.nombre
}


