package main

import "fmt"

func main() {
	fmt.Println("If else in golang")

	loginCount:= 12
	var result string

	if loginCount <10{
		result ="Regular User"
	}else if loginCount >10{
		result = "something else"
	}else{
		result = "Exactly 10 login count"
	}
	fmt.Println(result)

	if 9%2 == 0{
		fmt.Println("Number is even")
	}else{
		fmt.Println("Number is odd")
	}

	if num := 3; num<10{
		fmt.Println("Num is less then 10")
	}else{
		fmt.Println("Num is Not else then 10")
	}

	// if err! = nill{

	// }

	
}