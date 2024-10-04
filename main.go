package main

import "fmt"

func main() {

	//* Operator also termed as the dereferencing operator used to declare pointer variable and access the value stored in the address.

	//& operator termed as address operator used to returns the address of a variable or to access the address of a variable to a pointer.

	var p *int32 = new(int32) 
	var i int32
	fmt.Printf("%v\n",p)
	fmt.Printf("%v\n",*p)
	fmt.Printf("%v\n",i)



	var thing1 = [5]float64{1,2,3,4,5}
	fmt.Printf("\nThe memory location of the thing 1 array is : %p",&thing1)
	var result [5]float64 = sqaure(&thing1)
	fmt.Printf("\nThe result is :%v",result)
	fmt.Printf("\nThe value of thing1 is :%v",thing1)

}


func sqaure(thing2 *[5]float64) [5]float64{
	fmt.Printf("\nThe memory location of the thing 2 array is : %p",thing2)
	for i:=range thing2{
		thing2[i]=thing2[i]*thing2[i]
	}
	return *thing2
}
