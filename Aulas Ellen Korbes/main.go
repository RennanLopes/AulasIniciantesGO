// Aula 1  - GO

package main

import "fmt"

// Atividade 1
/*
func main() {

	x := 42
	y := "James Bonda"
	z := true

	fmt.Println(x, y, z)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
*/

//Atividade 2
/*
var x int
var y string
var z bool

func main() {

	fmt.Printf("%v\n%v\n%v\n", x, y, z)

}
*/

//Atividade 3
/*
var x int = 42
var y string = "James Bonda"
var z bool = true

func main() {

	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println(s)
}
*/

//Atividade 4
/*
type carros int

var x carros

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x := 42
	fmt.Println(x)
}
*/

//Atividade 4

type carros int

var x carros

var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
	fmt.Println("↑ foi x.\n↓ é y!")
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

}
