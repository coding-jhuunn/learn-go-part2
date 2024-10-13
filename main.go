package main

import "fmt"

func main() {

	a := "hello"
	fmt.Println("Before:", a)
	modString(a)
	fmt.Println("After:", a)

}
func modString(a string) {
	
	a = "modified"
}
