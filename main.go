package main

import(
	"fmt"
	"os"
	"strconv"
)
func calcuEdad(AN int)(int) {

	res := 2022 - AN 
	return res
}

func main() {
	
	N := os.Args[1]
	AP := os.Args[2]
	AM := os.Args[3]
	AN, _  := strconv.Atoi(os.Args[4])	

	res := calcuEdad(AN)
		
			fmt.Println("Hola",N,AP,AM,"tienes",res,"años de edad")
			
}