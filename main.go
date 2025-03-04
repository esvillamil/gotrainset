// Ejercicio de Sets
package main

import (
"fmt"

"github.com/esvillamil/gotrainset/pckset"
)

// main is the entry point of the program.
//
// It initializes a set, adds values to it, checks if values exist in the set,
// removes a value from the set, and checks if the value is still in the set.
//
// It does not return any values.
func main() {
	setA := map[int]bool{}
	fmt.Println("'simplesets.Añadir(1)'")
	simplesets.Añadir(1,setA)
	fmt.Println("'simplesets.Añadir(3)'")
	simplesets.Añadir(3,setA)
	fmt.Println("'simplesets.Existe(1)' -> ", simplesets.Existe(1,setA))
	fmt.Println("'simplesets.Existe(2)' -> ", simplesets.Existe(2,setA))
	fmt.Println("'simplesets.Borrar(1)'")
	simplesets.Borrar(1,setA)
	fmt.Println("'simplesets.Existe(1)' -> ", simplesets.Existe(1,setA))
}
