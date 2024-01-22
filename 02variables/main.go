package main

import "fmt"

const LoginToken string = "ghabbshh"

func main() {
	var username string = "Abhishek"
	fmt.Println(username)
	fmt.Printf("Variables is of type: %T \n",  username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variables is of type: %T \n",  isLoggedIn)

	var smallValue  uint8= 255
	fmt.Println(smallValue)
	fmt.Printf("Variables is of type: %T \n",  smallValue)

	var smallFloat  float32 = 255.568987689765
	fmt.Println(smallFloat)
	fmt.Printf("Variables is of type: %T \n",  smallFloat)

	//  default values and some aliases

	var anotherVaribales  int
	fmt.Println(anotherVaribales)
	fmt.Println("variables is of type: %T", anotherVaribales)

	// implicit type
	var website = 9999
	fmt.Println(website)


	//  no var style 

	numberOfUser := 0000
	fmt.Println(numberOfUser)
}