package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in go lang......!!")
	greater()

	greater2()
	result := adder(3,4)

	fmt.Println("Result is ",result)
	proRes , mymessage:= proAdder(2,8,9,6)
	fmt.Println("pro res is:", proRes)
	fmt.Println("pro message is:", mymessage)
}
// function signature
func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}


func proAdder(values ...int)(int, string) {
	total := 0
	for _,val:=range values{
		total+= val
	}
	return total, "hi pro result function"
}
func greater() {
	fmt.Println("Namaste From Golang")
}

func greater2() {
	fmt.Println("Another method")
}
