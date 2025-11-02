package main

import (
	"fmt"
)

// declarando e inicializando variavel a nivel package scope
var y = "olá bom dia"

func main() {
		// declarando e inicializando variavel x a nivel function scope
		x := 10

		fmt.Printf("x: %v, %T\n", x, x)
		fmt.Printf("y: %v, %T\n\n", y, y)

		x, z := 20, 30
		fmt.Printf("x: %v, %T\n", x, x)
		fmt.Printf("z: %v, %T\n", x, z)
}
