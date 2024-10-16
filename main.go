package main

import "fmt"

func main() {
	//fixed length with[3]
	// var intArr [3]int32
	var intArr [3]int32 =[3]int32{78,4,88}
	fmt.Println(intArr)
	//use & to print the memory allocation of each index
	fmt.Println(&intArr[0])

	//slice

	var intSlice []int32 = []int32{3,5,6}
	fmt.Printf("the lenght of %v with cap %v",len(intSlice),cap(intSlice))
	//value in a array using append method
	intSlice = append(intSlice, 8)
	fmt.Printf("\nthe lenght of %v with cap %v",len(intSlice),cap(intSlice))

	//append with spread operator
	var intSlice2 []int32 = []int32{8,9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)


	//create slice with make fnc      length  n cap
	var intSlice3 []int32 = make([]int32, 3,8)
	fmt.Println(intSlice3)

	//create map

	var myMap map[string]int32 = make(map[string]int32)
	fmt.Println(myMap)

	var myMap2  = map[string]uint32{"Adam":32,"Sarah":32}
	fmt.Println(myMap2["Adam"])
	// if the key doesnt exit, it will return something
	fmt.Println("something")
	fmt.Println(myMap2["Jason"])

	//optional value to check
	var age, ok =myMap2["Josh"]
	if ok{
		fmt.Printf("the age is  %v",age)
	}else{
		fmt.Println("Invalid name")
	}

	//delete value from a map
	//delete(myMap2,"Adam")
	fmt.Println(myMap2)

	for name,age:= range myMap2{
		fmt.Printf("Name: %v, Age: %v\n",name,age)
	}

	//the first value in a for loop will be always the index
	for v,i :=range intArr{
		fmt.Printf("Index: %v,Value: %v \n",i,v)
	}

}
