package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	var fruitList [4]string

	fruitList[0]="Apple"
	fruitList[1]="Mango"
	fruitList[3]="Grapes"

	fmt.Println("fruits list is", fruitList)
	fmt.Println("fruits list is", len(fruitList))

	var veglist = [5]string{"potatoes", "beans", "mushroom"}
	fmt.Println("veggy list is", len(veglist))


}