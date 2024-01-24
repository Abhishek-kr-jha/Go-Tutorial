package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	//  No inheritancde in golang
	// no suoer or parent in go lang

	abhishek := User{"Abhishek", "abhishek@gmail.com", true, 27}
	fmt.Println(abhishek)
	fmt.Printf("abhishek detials are : %+v\n", abhishek)
	fmt.Printf("Name is %v and email is %v", abhishek.Name, abhishek.Email)

}

// defining the strut in go-lang
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
