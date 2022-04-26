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

//package-level scope
//r := 10 //tidak bisa
var r int = 10

var (
	actorName    string = "Elisabeth Sladen"
	companion    string = "Sarah Jane Smith"
	doctorNumber int    = 3
	season       int    = 11
) //kalau declare value dipackage level, jika unused tidak masalah

var (
	counter int = 0
)

func variableDeclaration() {
	//function-level scope

	//CARA 1
	var i int
	i = 42
	i = 27

	//CARA 2
	var j float32 = 42

	//CARA 3
	k := 42 //Cara paling simple, dan membingungkan terkadang
	l := 42.

	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(l)

	fmt.Printf("%v, %T", k, k)
	fmt.Printf("%v", r)
}
