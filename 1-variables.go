package main

import "fmt"

// AGENDA
// Variable Declaration
// Redeclaration & Shadowing
// Visibility
// Naming Conventions
// Type Conversion

func main() {

}

func variableDeclaration() {
	//CARA 1
	var i int
	i = 42
	i = 27

	//CARA 2
	var j int = 42

	//CARA 3
	k := 42 //Cara paling simple, dan membingungkan terkadang

	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
}
