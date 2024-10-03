package main

import (
	"errors"
	"fmt"
)

func main() {

	printMe("Hello")
}

// dictates what type of params will be passed
func printMe(value string){
	fmt.Println(value)
	result,remainder,err := intDivide(2,1)
	//fmt.Println(err,nil)
	if err!=nil{
		fmt.Printf(err.Error())
	}else if remainder==0{
		fmt.Printf("The result of integer diviison is %v",result)
	}else{
		//other way to print
		//v is for the value
		fmt.Printf("The result of this computation is %v with remainder of %v",result,remainder)

	}
	// first way to print
	//fmt.Println(result,remainder)
}

// dictates what will be the return type of the value of the function
// (int,int) tells that it will be return 2 values
// create error message
// this is a genral design for go
func intDivide(numerator int,denominator int)(int,int,error) {
	var err error
	if denominator==0{
		err = errors.New("cannot dvide byt zero")
		return 0,0,err
	}
	var result int = numerator/denominator
	var remainder int = numerator%denominator
	//return 2 values
	return result,remainder,err
}