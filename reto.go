package main

import "fmt" 

func main() {
  fmt.Print("Ingrese una palabra: ")
  var input string
  fmt.Scanf("%s", &input)
  	for i := 0; i < len(input); i++{
   	 fmt.Println(input[:i+1])
	}
	for i := len(input)-1; i > 0; i--{
   	 fmt.Println(input[:i])
	}
}