package simplesets

import (

)


// Añadir adds the specified key to the given simple set.
//
// Parameters:
//   - key: the integer key to add.
//   - simpleSet: the map representing the set.
 func Añadir(key int, simpleSet map [int]bool ) {
 	simpleSet[key] = true
}


// Borrar deletes the specified key from the given simpleSet.
//
// Parameters:
//   - key: the integer key to be deleted.
//   - simpleSet: the map from which the key is to be deleted.
//
// Return type(s):
//   None.
 func Borrar( key int , simpleSet map[int] bool)  {
	delete(simpleSet , key)
}

// Existe checks if a given key exists in the simpleSet.
//
// Parameters:
//   - key: the integer key to search for in the simpleSet.
//   - simpleSet: a map of integers to booleans, representing the set.
//
// Returns:
//   - bool: true if the key exists in the simpleSet, false otherwise.
func Existe(key int, simpleSet map[int]bool) bool {

	if _, ok := simpleSet[ key ]; ok {
		return true
	} else {
		return false
	}
}