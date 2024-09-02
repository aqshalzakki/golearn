package main

import "fmt"

func main() {

	// if
	// var name string = "Aqshal"
	var name string = "Mulyono"

	// if name == "Aqshal" {
	// 	fmt.Println("Hello", name, "Ganteng")
	// } else if name == "Mulyono" {
	// 	fmt.Println("Hello Mr President", name)
	// } else {
	// 	fmt.Println("Hello", name)
	// }

	// name = "Aqshal"

	// switch name {
	// case "Aqshal":
	// 	fmt.Println("Hello", name, "Ganteng")
	// case "Mulyono":
	// 	fmt.Println("Hello Mr President", name)
	// default:
	// 	fmt.Println("Hello", name)
	// }

	switch length := len(name); length > 7 {
	case true:
		fmt.Println("Nama tidak boleh lebih dari 7")
	case false:
		fmt.Println("Nama aman")
	}
}
