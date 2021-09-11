package variable

import (
	"fmt"
	"strconv"
)

// Global Variable (Pascal case) With in a package
var Val2 = 100

// Package Level Variable (Camel case) With in a package and other package
var myValue int = 100

func Varriable() {

	// Local Variable
	//First Type
	//var val1 int  //Decleration
	//val1 = 100  //Initialzation

	// Second Type
	//var val1 = 100  // Decleration with Initilization

	// Third type

	val1 := 100

	// Type Conversion

	var a int32 = 100
	var b int64 = int64(a)

	//Integer to String
	var str int = 65
	var d string = strconv.Itoa(str)

	//Asky Value
	var f string = string(str)
	//myValue = 50

	//Premitive Datatype

	//Constant variable

	const (
		x = iota
		_
		y = iota
		z
	)

	const i int = 45

	fmt.Print(i)

	fmt.Println("hello world")
	fmt.Println(val1)
	fmt.Println(myValue)
	fmt.Printf("%v , %T", val1, val1)
	fmt.Println("")
	fmt.Printf("%v , %T", b, b)
	fmt.Println("")
	fmt.Printf("%v , %T", d, d)
	fmt.Println("")
	fmt.Printf("%v , %T", f, f)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

}
