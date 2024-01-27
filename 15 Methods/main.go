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
	abhishek.GetStatus()
	abhishek.NewMail()
	fmt.Printf("Name is %v and email is %v", abhishek.Name, abhishek.Email)


}

// defining the strut in go-lang
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
	
}


func (u User) GetStatus(){
	fmt.Println("Is user active:", u.Status)

}

func (u User) NewMail(){
	u.Email = "test@go.dev"
	fmt.Println("Email of this user:",u.Email)
}
